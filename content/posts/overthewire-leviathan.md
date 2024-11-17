---
title: "OverTheWire - Leviathan"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

Level 0

root@kali:~# ssh leviathan.labs.overthewire.org -p 2223 -l leviathan0  
[...]

leviathan0@leviathan:~$ ls  
leviathan0@leviathan:~$ ls -alh  
total 24K  
drwxr-xr-x 3 root root 4.0K Oct 29 2018 .  
drwxr-xr-x 10 root root 4.0K Oct 29 2018 ..  
drwxr-x--- 2 leviathan1 leviathan0 4.0K Oct 29 2018 .backup  
-rw-r--r-- 1 root root 220 May 15 2017 .bash_logout  
-rw-r--r-- 1 root root 3.5K May 15 2017 .bashrc  
-rw-r--r-- 1 root root 675 May 15 2017 .profile  
leviathan0@leviathan:~$ cat .backup  
cat: .backup: Is a directory  
leviathan0@leviathan:~$ cd .backup/  
leviathan0@leviathan:~/.backup$ ls  
bookmarks.html  
leviathan0@leviathan:~/.backup$ ls -alh  
total 140K  
drwxr-x--- 2 leviathan1 leviathan0 4.0K Oct 29 2018 .  
drwxr-xr-x 3 root root 4.0K Oct 29 2018 ..  
-rw-r----- 1 leviathan1 leviathan0 131K Oct 29 2018 bookmarks.html  
leviathan0@leviathan:~/.backup$ cat bookmarks.html | grep pass  
<DT><A HREF="http://leviathan.labs.overthewire.org/passwordus.html | This will be fixed later, the password for leviathan1 is rioGegei8m" ADD_DATE="1155384634" LAST_CHARSET="ISO-8859-1" ID="rdf:#$2wIU71">password to leviathan1</A>  
leviathan0@leviathan:~/.backup$

Level 0 zu Level 1