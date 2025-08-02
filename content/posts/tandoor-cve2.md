---
title: Continuous Checks are Important - Privilege Escalation in Tandoor Recipes
date: 2025-08-02T04:00:56+01:00
toc: false
images: 
tags:
  - CVE
  - Web
  - English
---

Tandoor Recipes 2.0.0-alpha-1 is vulnerable to privilege escalation. This is due to the rework of the API, which resulted in the User Profile API Endpoint containing two boolean values indicating whether a user is staff or administrative. Consequently, any user can escalate their privileges to the highest level.

This vulnerable version was released soon after I published the second part of my series, 'All Your Recipe Are Belong to Us', in which I tested - among others - Tandoor Recipes for vulnerabilities (https://m10x.de/posts/2025/01/all-your-recipe-are-belong-to-us-part-2/3-server-side-template-injection-rce-arbitrary-file-read-and-unrestricted-file-upload-stored-xss-in-tandoor-recipes/). This highlights the importance of continuous checks!

## Overview of the Vulnerabilities
| CVE | Name                                                                               | CVSS Score      |
| ----- | ---------------------------------------------------------------------------------- | ------ |
| CVE-2025-PENDING | Privilege Escalation | [8.8 High](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:I/C:H/I:H/A:H) |

## Remediation

The maintainer reacted quickly and professionally. The vulnerability has been fixed in Tandoor Recipes version [2.0.0-alpha-2](https://github.com/TandoorRecipes/recipes/releases/tag/2.0.0-alpha-2).

## Vulnerabilities in Detail

### [CVE-2025-PENDING] Privilege Escalation (8.8 High)

Users are allowed to update their names:

![Tandoor1](/media/2025/08/tandoor1.png)

In this case the following API Call was sent:

![Tandoor2](/media/2025/08/tandoor2.png)

The regular user is able to modify the parameters is_staff and is_superuser to true, in order to grant themself those privileges:

![Tandoor3](/media/2025/08/tandoor3.png)

Thus, the user has now the staff and admin privileges (just as already revealed by the previous api response)

![Tandoor4](/media/2025/08/tandoor4.png)

## Timeline
| Date | Event |
| - | - |
| 2025-04-07 | Discovered and reported the vulnerability |
| 2025-04-09 | Maintainer acknowledged the vulnerabilities thankfully |
| 2025-04-18 | Maintainer fixed the privilege escalation in version 2.0.0-alpha-2 |