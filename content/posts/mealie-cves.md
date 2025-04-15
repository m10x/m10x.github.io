---
title: All Your Recipe Are Belong to Us (Part 3/3) - Broken Access Controls Leading to Privilege Escalation and More in Mealie
date: 2025-03-25T04:00:56+01:00
toc: false
images: 
tags:
  - CVE
  - Web
  - English
---

I used 4 hours of my free time (not counting the Responsible Disclosure and Blog Posts...) to "speed pentest" the three biggest and most popular (measured by Github stars) open-source cooking recipe managers. 

Among them is the recipe manager that I personally use and that gave me the idea for this project: [Mealie](https://github.com/mealie-recipes/mealie). Mealie offers "Recipe Management For The Modern Household" and had >7500 stars at the time of testing. Since Mealie introduced new features and many code changes with version 2.0.0 only a month ago (10/22/2024), I thought now would be a good time to test it. I found 4 different Broken Access Control vulnerabilities that can be used for privilege escalation. All 4 vulnerabilities were found in version 2.2.0, although they are most likely present in Mealie since version 2.0.0 or even earlier.

## Overview of the Vulnerabilities
| CVE | Name                                                                               | CVSS Score      |
| ----- | ---------------------------------------------------------------------------------- | ------ |
| [CVE-2024-55073](https://www.cve.org/CVERecord?id=CVE-2024-55073) | Users can edit their own profile in order to give themselves more permissions or to change their household | [7.6 High](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:U/C:L/I:H/A:L) |
| [CVE-2024-55072](https://www.cve.org/CVERecord?id=CVE-2024-55072) | Group managers can give themselves more permissions                                          | [5.4 Medium](https://www.first.org/cvss/calculator/3.1#CVSS:3.0/AV:N/AC:L/PR:L/UI:N/S:U/C:N/I:L/A:L) |
| [CVE-2024-55071](https://www.cve.org/CVERecord?id=CVE-2024-55071) | Not fixed yet :)          | [4.2 Medium](https://www.first.org/cvss/calculator/3.1#CVSS:3.0/AV:N/AC:H/PR:L/UI:N/S:U/C:N/I:L/A:L) |
| [CVE-2024-55070](https://www.cve.org/CVERecord?id=CVE-2024-55070) | Users can share recipes of other groups          | [3.1 Low](https://www.first.org/cvss/calculator/3.0#CVSS:3.1/AV:N/AC:H/PR:L/UI:N/S:U/C:L/I:N/A:N) |

## Remediation

The three disclosed vulnerabilities have been fixed in version [2.5.0] (https://github.com/mealie-recipes/mealie/releases/tag/v2.5.0). One undisclosed vulnerability has not yet been fixed.

## Vulnerabilities in Detail

### [CVE-2024-55073] Users can edit their own profile in order to give themselves more permissions or to change their household (7.6 High)
Users are allowed to edit their own profile:

![Mealie1](/media/2025/03/mealie1.png)

Upon clicking "Update", the following API call is sent:

![Mealie2](/media/2025/03/mealie2.png)

As one can see, the API call contains not only the Username, Full Name and Email, but also a few more attributes such as multiple permissions (canXXX) as well as the householdId the user is belonging to. The current user is not allowed to Invite, Manage nor Organize. However, the "cans" can be changed to true, as well as the householdId set to another one:

![Mealie3](/media/2025/03/mealie3.png)

Mealie succesfully updates those values. Both of these actions can normally only be done by an administrator.
Thus, we've escalated our privileges as well as changed our household. The following screenshot shows, that User2 now belong to Household1 instead of Household2 and that User2 is allowed to Manage the household's members

![Mealie4](/media/2025/03/mealie4.png)

Getting the id of another household in order to switch to another household is trivial, as one can see all householdIds of the current group.

![Mealie41](/media/2025/03/mealie41.png)

On a positive note, it was not possible to escalate privileges to administrator or to change to another group.  
Nonetheless, we are able to manage households and their members, to invite Members, to organize as well as to take a look/modify/delete other households receipes and shopping lists by switching to that household.

### [CVE-2024-55072] Group managers can give themselves more permissions (5.4 Medium)
Group managers are not allowed to change their own permissions:

![Mealie5](/media/2025/03/mealie5.png)

When a group manager changes a permission of another group member, an API call is issued to the API endpoint `PUT /api/households/permissions`, which contains the userId of the other group member. If this userId is swapped with the group manager's own userId, the group manager can change their own permissions.

![Mealie6](/media/2025/03/mealie6.png)

Thus they escalated their privileges:

![Mealie61](/media/2025/03/mealie61.png)


### [CVE-2024-55071] Not fixed yet :) (5.2 Medium)

...

### [CVE-2024-55070] Users can share recipes of other groups (3.1 Low)

Users can only see recipes created in their group, unless a recipe is shared and they know the link.
A user can share recipes of his group, leading to a call to the API endpoint `POST /api/shared/recipes`. By interchanging the recipeId with the recipeId of a recipe which belongs to another group, they are able to share the recipe:

![Mealie7](/media/2025/03/mealie7.png)

Thus, they are able to view the recipe, too

![Mealie8](/media/2025/03/mealie8.png)

The users of the other group cannot see that the recipe is being shared. 

![Mealie9](/media/2025/03/mealie9.png)

Normally, the frontend lists all shares of a recipe, but it does not show shares which were issued by users not belonging to that group. Here is an example of what this would  normally look like:

![Mealie81](/media/2025/03/mealie81.png)


## Timeline
| Date | Event |
| - | - |
| 2024-11-21 | Discovered the vulnerabilities |
| 2024-11-22 | Reported the vulnerabilities |
| 2024-11-23 | Maintainer acknowledged the vulnerabilities |
| 2024-11-23 | CVEs requested |
| 2024-11-27 | CVEs were reserved |
| 2025-01-22 | 3 out of 4 vulnerabilities have been fixed |
