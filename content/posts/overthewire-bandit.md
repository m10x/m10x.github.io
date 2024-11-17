---
title: "OverTheWire Bandit WriteUp - Alle Level"
date: 2018-02-13T20:38:56+01:00
toc: false
images:
tags:
  - overthewire
  - ctf
  - german
---

## **Einleitung Wargames**

[OverTheWire.org](http://overthewire.org/wargames/) bietet über 15 verschiedene Wargames an, welche verschiedene Schwierigskeitgrade haben. Wenn du den Begriff "Wargame" im Bezug zu Cyber Security zum ersten Mal hörst, fragst du dich bestimmt: "Was ist eigentlich ein Wargame?".  
Wargames sind Herausforderungen, entweder online oder offline als Virtuelle Maschine, in denen man versucht Sicherheitslücken auszunutzen und/oder Zugriff zu etwas zu erlangen. Sie sind teilweise sehr unterschiedlich was die benötigten Fähigkeiten, die Herausforderung und die Vorgehensweise anbelangt. Wargames werden von verschiedenen Internetseiten angeboten und haben teilweise eine große Community. Eine Liste von Wargame Anbietern kannst du [hier](https://github.com/apsdehal/awesome-ctf#wargames) finden.

## **Einleitung Bandit Wargame**

[Bandit](http://overthewire.org/wargames/bandit/) ist, wie schon erwähnt, eines der Wargames von [OverTheWire.org](http://overthewire.org/wargames/)  
Es ist ein Wargame für Beginner / Einsteiger und perfekt für welche, die noch keine Erfahrung mit Wargames haben. Es bringt einem die Grundlagen bei, die man benötigt, um andere Wargames erfolgreich zu absolvieren. Grundkenntnisse in Linux / Bash sind vorausgesetzt. Falls ihr bei einem Level nicht weiter kommen solltet, könnt ihr euch ansehen, wie ich das Level gelöst habe oder auch euren Ansatz mit meinem vergleichen. Bei [Bandit](http://overthewire.org/wargames/bandit/) geht es darum sich mit einem vorgegebenen Benutzernamen und Passwort über SSH sich zu einem Linux Server zu verbinden. Nun ist es die Aufgabe das Passwort herauszufinden, um sich als ein anderer Benutzer anzumelden.

## **Level 0**

[Hier](http://overthewire.org/wargames/bandit/) kommst du zu der Internetseite von Bandit. Alles Relevante werde ich dir aber auch hier mitteilen.

Als erstes müssen wir uns über [SSH](https://wiki.ubuntuusers.de/SSH/) mit dem Spiel verbinden. Die Adresse des Host ist **bandit.labs.overthewire.org** und der Port **2220**. Der Benutzername ist **bandit0** und das Passwort ebenfalls **bandit0**.

Linux:  
Um die SSH Verbindung zum Spiel herzustellen, benutzen wir folgendes:

```shell
ssh bandit0@bandit.labs.overthewire.org -p 2220
```

In diesem Fall haben wir uns als Benutzer **bandit0** angemeldet.

Um die Verbindung zu trennen um uns nachher als nächster Benutzer anzumelden, schreiben wir **exit**.

## **Level 0 -> Level 1**

Nun müssen wir versuchen das Passwort für den nächsten Benutzer **bandit1** herauszufinden.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort können wir in einer Datei namens **readme** finden, welche im Homeverzeichnis liegt.

Als erstes benutzen wir den Befehl **ls** um herauszufinden, welche Dateien im aktuellen Verzeichnis zu finden sind.

```shell
bandit0@bandit:~$ ls
readme
```

Wir haben nun also die Datei **readme** gefunden, in welcher das Passwort gespeichert ist.  
Jetzt können wir **cat** benutzen, um den Inhalt der Datei auszugeben.

```shell
bandit0@bandit:~$ cat readme
boJ9jbbUNNfktd78OOpsqOltutMc3MY1
```

## **Level 1 -> Level 2**

Der Benutzername, mit dem wir uns für das nächste Level anmelden müssen, ist logischerweise **bandit1**. Das Passwort haben wir gerade herausgefunden. Es lautet: **boJ9jbbUNNfktd78OOpsqOltutMc3MY1**  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort liegt in einer Datei namens **-**, welche im Homeverzeichnis liegt.

Wir benutzen erst einmal wieder **ls** um alle Dateien im aktuellen Verzeichnis anzeigen zu lassen.

```shell
bandit1@bandit:~$ ls
-
```

Wir können nun aber nicht analog zum vorherigen Level **cat -** benutzen, um den Inhalt der Datei **-** auszugeben, da **-** ebenfalls dazu benutzt wird, um Optionen für Befehle anzugeben.  
Stattdessen können wir **cat ./-** benutzen. **./** bezeichnet den aktuellen Ordner, in dem man sich befindet.

```shell
bandit1@bandit:~$ cat ./-
CV1DtqXWVFXTvM2F0k09SHz0YwRINYA9
```

Alternativ können wir auch **cat < -** benutzen. **<** wird, unter Anderem, benutzt um den Inhalt einer Datei an einen Befehl umzuleiten. Weiter Informationen zu Umleitungen kannst du [hier](https://wiki.ubuntuusers.de/Shell/Umleitungen/) finden.

```shell
bandit1@bandit:~$ cat < -
CV1DtqXWVFXTvM2F0k09SHz0YwRINYA9
```

## **Level 2 -> Level 3**

Wir melden uns nun als **bandit2** mit dem Passwort **CV1DtqXWVFXTvM2F0k09SHz0YwRINYA9** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort ist in einer Datei namens **spaces in this filename**, welche im Heimverzeichnis liegt.

Wir benutzen erst einmal wieder **ls**, um alle Dateien im aktuellen Verzeichnis anzeigen zu lassen.

```shell
bandit2@bandit:~$ ls
spaces in this filename
```

Hier haben wir jetzt verschiedene Möglichkeiten.  
Die einfachste Möglichkeit ist **cat "spaces in this filename"** oder **cat 'spaces in this filename'**.  
Alternativ geht auch **cat spaces\ in\ this\ filename**.

```shell
bandit2@bandit:~$ cat "spaces in this filename"
UmHadQclWmgdLOKQ3YNgjWxGoRMb5luK
```

## **Level 3 -> Level 4**

Wir melden uns nun als **bandit3** mit dem Passwort **UmHadQclWmgdLOKQ3YNgjWxGoRMb5luK** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in einer versteckten Datei (= hidden file), in dem **inhere** Ordner.

Wir fangen wieder mit dem alt bewährten **ls** an, um alle Dateien im aktuellen Verzeichnis anzeigen zu lassen.

```shell
bandit3@bandit:~$ ls
inhere
```

Wir sehen nun den Ordner **inhere**. Um diesen zu öffnen, benutzen wir den Befehl **cd**

```shell
bandit3@bandit:~$ cd inhere
bandit3@bandit:~/inhere$
```

Wenn wir nun wieder **ls** benutzen, wird uns nichts angezeigt, da die Datei, welche wir suchen, versteckt wurde.  
Deswegen benutzen wir die Option **-a** von dem Befehl **ls**. Das sorgt dafür, dass auch alle versteckten Dateien und Ordner angezeigt werden.

```shell
bandit3@bandit:~/inhere$ ls -a
. .. .hidden
```

Da wir jetzt wissen, dass die Datei **.hidden** heißt und in dem aktuellen Ordner hier liegt, können wir den Inhalt wie gewohnt mit **cat** ausgeben lassen

```shell
bandit3@bandit:~/inhere$ cat .hidden
pIwrPrtPN36QITSp3EQaw936yaFoFgAB
```

## **Level 4 -> Level 5**

Wir melden uns nun als **bandit4** mit dem Passwort **pIwrPrtPN36QITSp3EQaw936yaFoFgAB** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in der einzigen Datei, im **inhere** Ordner, welche für Menschen lesbar (= human-readable) ist.

Es kommt wieder zuerst das altbewährte **ls** zum Einsatz und danach **cd** um den Ordner **inhere** zu öffnen

```shell
bandit4@bandit:~$ ls
inhere
bandit4@bandit:~$ cd inhere
bandit4@bandit:~/inhere$
```

Jetzt wieder **ls** um die Dateien im Ordner sehen zu können

```shell
bandit4@bandit:~/inhere$ ls
-file00 -file02 -file04 -file06 -file08
-file01 -file03 -file05 -file07 -file09
```

Es befinden sich 10 Dateien in dem Ordner. Man könnte sich nun jede einzelne ansehen, aber das würde zu lange dauern und zu aufwändig sein.  
Mit **file** können wir uns anzeigen lassen was für eine Art von Inhalt eine Datei enthält. **./*** bezeichnet alle Dateien die im aktuellen Ordner sind.

```shell
bandit4@bandit:~/inhere$ file ./*
./-file00: data
./-file01: data
./-file02: data
./-file03: data
./-file04: data
./-file05: data
./-file06: data
./-file07: ASCII text
./-file08: data
./-file09: data
```

**./-file07: ASCII text** sieht viel versprechend aus. Sehen wir uns die Datei mal genauer an mit Hilfe von **cat**.

```shell
bandit4@bandit:~/inhere$ cat ./-file07
koReBOKuIDDepwhWk7jZC0RTdopnAYKh
```

## **Level 5 -> Level 6**

Wir melden uns nun als **bandit5** mit dem Passwort **koReBOKuIDDepwhWk7jZC0RTdopnAYKh** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort ist wieder in einer Datei innerhalb des **inhere** Ordners.  
Diese Datei hat folgende Eigenschaften:  
- Sie ist für Menschen lesbar (= human-readable)  
- Sie ist 1033 Bytes groß  
- Sie ist nicht ausführbar (= not executable)

Als erstes begeben wir uns wieder in den **inhere** Ordner und schauen uns dessen Inhalt an.

```shell
bandit5@bandit:~$ ls
inhere
bandit5@bandit:~$ cd inhere/
bandit5@bandit:~/inhere$ ls
maybehere00 maybehere04 maybehere08 maybehere12 maybehere16
maybehere01 maybehere05 maybehere09 maybehere13 maybehere17
maybehere02 maybehere06 maybehere10 maybehere14 maybehere18
maybehere03 maybehere07 maybehere11 maybehere15 maybehere19
```

Im **inhere** Ordner befinden sich also 20 weitere Ordner. Um die richtige Datei, welche das Passwort beinhaltet, schnell und einfach zu finden, können wir den Befehl **find** benutzen.  
Wir benutzen den Befehl find mit folgenden Optionen:  
**-type f**, da wir eine normale Datei suchen. **f** steht in diesem Fall einfach nur für **file**.  
**-readable**, da die Datei für Menschen lesbar ist (= human-readable)  
**! -executable**, da die Datei nicht ausführbar ist (= non executable). Das **!** verneint in diesem Fall die Option **-executable**.  
**-size 1033c**, da die Datei 1033 Bytes groß ist. **c** gibt in diesem Fall an, dass die Größe in Bytes gemeint ist.  
Und natürlich auch **./**, was einfach nur angibt, dass die Datei irgendwo innerhalb des aktuellen Ordners (**inhere**) zu finden ist.

```shell
bandit5@bandit:~/inhere$ find ./ -type f -readable ! -executable -size 1033c
./maybehere07/.file2
```

Nun müssen wir nur noch den Inhalt der Datei **.file2**, welche sich im Ordner **maybehere07** befindet, ausgeben.

```shell
bandit5@bandit:~/inhere$ cat ./maybehere07/.file2
DXjZPULLxYr17uwoI01bNLQbtFemEgo7
```

## **Level 6 -> Level 7**

Wir melden uns nun als **bandit6** mit dem Passwort **DXjZPULLxYr17uwoI01bNLQbtFemEgo7** an.  
Der Tipp zu diesem Level ist wie folgt:  
Die Datei, in welcher sich das Passwort befindet, ist irgendwo auf dem Server und hat folgende Eigenschaften:  
- Sie ist im Besitz von dem Benutzer **bandit7**  
- Sie ist im Besitz von der Gruppe **bandit6**  
- Sie ist 33 Bytes groß

Hier kommt, wie im vorherigen Level, wieder der Befehl **find** zum Einsatz.  
Diesmal mit folgenden Optionen:  
**/**, damit der gesamte Server durchsucht wird  
**-user bandit7**, da der Benutzer **bandit7** als Besitzer eingetragen ist  
**-group bandit6**, da ebenfalls die Gruppe **bandit6** als Besitzer eingetragen ist  
**-size 33c**, da die Datei 33 Bytes groß ist. **c** gibt in diesem Fall an, dass die Größe in Bytes gemeint ist.  
**2>/dev/null**, diese Option ist dafür da jede Fehlermeldung umzuleiten, damit nicht jede einzelne wie z.B. "Zugriff verweigert" in unserem Terminal ausgegeben wird. Weitere Informationen zu Umleitungen kannst du [hier](https://wiki.archlinux.de/title/Umleitungen#Ausgabeumleitung) finden.

```shell
bandit6@bandit:~/inhere$ find / -user bandit7 -group bandit6 -size 33c 2>/dev/null
/var/lib/dpkg/info/bandit7.password
```

Jetzt wo wir die Datei gefunden haben, müssen wir wieder nur ihren Inhalt ausgeben.

```shell
bandit6@bandit:~$ cat /var/lib/dpkg/info/bandit7.password
HKBPTKQnIay4Fw76bEy8PVxKEDQRKTzs
```

## **Level 7 -> Level 8**

Wir melden uns nun als **bandit7** mit dem Passwort **HKBPTKQnIay4Fw76bEy8PVxKEDQRKTzs** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in der Datei **data.txt** neben dem Wort **millionth**.

Mithilfe von [grep](https://wiki.ubuntuusers.de/grep/) können wir Input nach bestimmten Zeichen und Wörtern durchsuchen und nur Zeilen darstellen lassen, welche diese enthalten. Mithilfe von **>** leiten wir den Output von **cat** zu [grep](https://wiki.ubuntuusers.de/grep/) um. Mehr Informationen zu [Umleitungen](https://wiki.ubuntuusers.de/Shell/Umleitungen/#Umleiten-der-Ausgabe-mit) kannst du [hier](https://wiki.ubuntuusers.de/Shell/Umleitungen/#Umleiten-der-Ausgabe-mit) finden.

```shell
bandit7@bandit:~$ ls
data.txt
bandit7@bandit:~$ cat data.txt | grep millionth
millionth cvX2JJa4CFALtqS87jk27qwqGhBM9plV
```

## **Level 8 -> Level 9**

Wir melden uns nun als **bandit8** mit dem Passwort **cvX2JJa4CFALtqS87jk27qwqGhBM9plV** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in der Datei **data.txt** in der einzigen Zeile, die ein einziges Mal vorkommt.

Für dieses Level benutzen wir wie gewohnt:  
- cat  
- den [Pipe-Operator |](https://wiki.ubuntuusers.de/Shell/Umleitungen/#Der-Pipe-Operator), dieser Leitet den Output eines Befehls direkt zu einem anderen Befehl weiter  
- [sort](https://wiki.ubuntuusers.de/sort/), mithilfe von [sort](https://wiki.ubuntuusers.de/sort/) können wir Dateien zeilenweise sortieren.  
- [uniq](https://wiki.ubuntuusers.de/uniq/), mithilfe von [uniq](https://wiki.ubuntuusers.de/uniq/) können wir Dateien ohne doppelte Zeilen ausgeben. Die Option **-u** sorgt dafür, dass nur Zeilen ausgegeben werden, welche nicht mehrmals vorkommen.

```shell
bandit8@bandit:~$ ls
data.txt
bandit8@bandit:~$ cat data.txt | sort | uniq -u
UsvVyFSfZZWbi6wgC7dAFyFuR6jQQUhR
```

**cat** gibt den Inhalt der Datei **data.txt** an **sort** weiter.  
**sort** sortiert diesen und gibt ihn an **uniq -u** weiter.  
**uniq -u** gibt dann die nur die Zeilen aus, welche nicht mehrmals vorkommen

## **Level 9 -> Level 10**

Wir melden uns nun als **bandit9** mit dem Passwort **UsvVyFSfZZWbi6wgC7dAFyFuR6jQQUhR** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in der Datei **data.txt** in einer der wenigen für Menschen lesbaren (= human-readable) Strings und beginnt mit mehreren **=**.

**data.txt** ist diesmal eine Binär-Datei. Wenn wir [cat](https://wiki.ubuntuusers.de/cat/) benutzen sehen wir nur ein unleserliches Wirrwarr. Stattdessen benutzen wir [strings](https://en.wikipedia.org/wiki/Strings_(Unix)). Strings gibt nur Zeilen mit "druckbaren"/"lesbaren" Charakteren aus.

```shell
bandit9@bandit:~$ ls
data.txt
bandit9@bandit:~$ strings data.txt | grep '=='
========== theP`
========== password
L========== isA
========== truKLdjsbJ5g7yyJ2X2R0o3a5HQJFuLk
```

[su_box title="Möchtest du benachrichtigt werden, wenn ein neuer Artikel veröffentlicht wurde?" style="noise" box_color="#1874CD" title_color="#FFFFFF" radius="20" class=""][email-subscribers namefield="YES" desc="" group="Public"][/su_box]

## **Level 10 -> Level 11**

Wir melden uns nun als **bandit10** mit dem Passwort **truKLdjsbJ5g7yyJ2X2R0o3a5HQJFuLk** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in der Datei **data.txt**, welche [Base64](https://de.wikipedia.org/wiki/Base64) enkodierte Daten enthält.

Mithilfe von **base64 -d** können wir ganz einfach Strings und Dateien, welche mit [Base64](https://de.wikipedia.org/wiki/Base64) enkodiert wurden, dekodieren.

```shell
bandit10@bandit:~$ ls
data.txt
bandit10@bandit:~$ base64 -d data.txt 
The password is IFukwKGsFW8MOq3IRFqrxE1hxTNEbUPR
```

## **Level 11 -> Level 12**

Wir melden uns nun als **bandit11** mit dem Passwort **IFukwKGsFW8MOq3IRFqrxE1hxTNEbUPR** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in der Datei **data.txt**, allerdings wurden jeweils alle Kleinbuchstaben (a-z) und Großbuchstaben (A-Z) um 13 Positionen rotiert.

Hier wurde die bekannte [ROT13 Caesar-Verschlüsselung](https://de.wikipedia.org/wiki/ROT13) benutzt. Diese können wir mithilfe von [tr](https://wiki.ubuntuusers.de/tr/) oder Python rückgängig machen.

```shell
bandit11@bandit:~$ ls
data.txt

bandit11@bandit:~$ cat data.txt
Gur cnffjbeq vf 5Gr8L4qetPEsPk8htqjhRK8XSP6x2RHh

bandit11@bandit:~$ cat data.txt | tr a-zA-Z n-za-mN-ZA-M
The password is 5Te8Y4drgCRfCx8ugdwuEX8KFC6k2EUu

bandit11@bandit:~$ python -c 'print "Gur cnffjbeq vf 5Gr8L4qetPEsPk8htqjhRK8XSP6x2RHh".decode("rot13")'
The password is 5Te8Y4drgCRfCx8ugdwuEX8KFC6k2EUu
```

## **Level 12 -> Level 13**

Wir melden uns nun als **bandit12** mit dem Passwort **5Te8Y4drgCRfCx8ugdwuEX8KFC6k2EUu** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort befindet sich in der Datei **data.txt**, welche ein Hexdump einer Datei ist, die mehrmals komprimiert wurde. Für dieses Level empfiehlt es sich, unter **/tmp** ein Verzeichnis zu erstellen, die Datei dort hinzukopieren und umzubenennen.

In diesem Level müssen wir Dateien mehrmals mithilfe von [tar](https://wiki.ubuntuusers.de/tar/), [gzip](https://wiki.ubuntuusers.de/gzip/) und [bzip2](https://wiki.ubuntuusers.de/bzip2/) entpacken. Mithilfe von [file](https://de.wikipedia.org/wiki/File#Benutzung), finden wir heraus, womit die Dateien komprimiert wurden.

```shell
bandit12@bandit:~$ ls
data.txt
bandit12@bandit:~$ cat data.txt 
00000000: 1f8b 0808 ecf2 445a 0203 6461 7461 322e ......DZ..data2.
00000010: 6269 6e00 0149 02b6 fd42 5a68 3931 4159 bin..I...BZh91AY
00000020: 2653 5930 3e1b 4000 0014 ffff dde3 2b6d &SY0>.@.......+m
[...]
bandit12@bandit:~$ mkdir /tmp/m10x
bandit12@bandit:~$ cp data.txt /tmp/m10x/data.txt
bandit12@bandit:~$ cd /tmp/m10x/

bandit12@bandit:/tmp/m10x$ xxd -r data.txt > data
bandit12@bandit:/tmp/m10x$ file data
data: gzip compressed data, was "data2.bin", last modified: Thu Dec 28 13:34:36 2017, max compression, from Unix
bandit12@bandit:/tmp/m10x$ mv data data.gz
bandit12@bandit:/tmp/m10x$ gzip -d data.gz
bandit12@bandit:/tmp/m10x$ ls
data data.txt

bandit12@bandit:/tmp/m10x$ file data
data: bzip2 compressed data, block size = 900k
bandit12@bandit:/tmp/m10x$ mv data data.bz2
bandit12@bandit:/tmp/m10x$ bzip2 -d data.bz2 
bandit12@bandit:/tmp/m10x$ ls
data data.txt

bandit12@bandit:/tmp/m10x$ file data
data: gzip compressed data, was "data4.bin", last modified: Thu Dec 28 13:34:36 2017, max compression, from Unix
bandit12@bandit:/tmp/m10x$ mv data data.gz
bandit12@bandit:/tmp/m10x$ gzip -d data.gz 
bandit12@bandit:/tmp/m10x$ ls
data data.txt

bandit12@bandit:/tmp/m10x$ file data
data: POSIX tar archive (GNU)
bandit12@bandit:/tmp/m10x$ mv data data.tar
bandit12@bandit:/tmp/m10x$ tar -xvf data.tar 
data5.bin
bandit12@bandit:/tmp/m10x$ ls
data.tar data.txt data5.bin

bandit12@bandit:/tmp/m10x$ file data5.bin 
data5.bin: POSIX tar archive (GNU)
bandit12@bandit:/tmp/m10x$ tar -xvf data5.bin
data6.bin

bandit12@bandit:/tmp/m10x$ file data6.bin 
data6.bin: bzip2 compressed data, block size = 900k
bandit12@bandit:/tmp/m10x$ mv data6.bin data6.bz
bandit12@bandit:/tmp/m10x$ bzip2 -d data6.bz
bandit12@bandit:/tmp/m10x$ ls
data.tar data.txt data5.bin data6

bandit12@bandit:/tmp/m10x$ file data6
data6: POSIX tar archive (GNU)
bandit12@bandit:/tmp/m10x$ mv data6 data6.tar
bandit12@bandit:/tmp/m10x$ tar -xvf data6.tar
data8.bin

bandit12@bandit:/tmp/m10x$ file data8.bin
data8.bin: gzip compressed data, was "data9.bin", last modified: Thu Dec 28 13:34:36 2017, max compression, from Unix
bandit12@bandit:/tmp/m10x$ mv data8.bin data8.gz
bandit12@bandit:/tmp/m10x$ gzip -d data8.gz
bandit12@bandit:/tmp/m10x$ ls
data.tar data.txt data5.bin data6.tar data8

bandit12@bandit:/tmp/m10x$ file data8
data8: ASCII text
bandit12@bandit:/tmp/m10x$ cat data8
The password is 8ZjyCRiBWFYkneahHwxCv3wb2a1ORpYL
```

## **Level 13 -> Level 14**

Wir melden uns nun als **bandit13** mit dem Passwort **8ZjyCRiBWFYkneahHwxCv3wb2a1ORpYL** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort für das nächste Level befindet sich in  **/etc/bandit_pass/bandit14** und kann nur von dem Benutzer **bandit14** gelesen werden. Allerdings bekommst du den [privaten SSH key](https://de.wikipedia.org/wiki/Secure_Shell#Authentifizierung), mit welchem du dich als **bandit14** einloggen kannst.

In diesem Level bekommen wir den [privaten SSH key](https://de.wikipedia.org/wiki/Secure_Shell#Authentifizierung) für das nächste Level. Mit **ssh -i** können wir den [privaten SSH key](https://de.wikipedia.org/wiki/Secure_Shell#Authentifizierung) um uns als **bandit14** mit **localhost** zu verbinden.

```shell
bandit13@bandit:~$ ls
sshkey.private
bandit13@bandit:~$ ssh -i sshkey.private bandit14@localhost
bandit14@bandit:~$
```

## **Level 14 -> Level 15**

Wir melden uns nun als **bandit14** mit dem Passwort **4wcYUJFw0k0XLShlDzztnTBHiqxU3b3e** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort für das nächste Level bekommst du, indem du das Passwort des derzeitigen Levels zu dem Port 30000 auf dem lokalen Host einreichst.

Bei diesem Level senden wir mithilfe von [nc ( netcat )](https://wiki.ubuntuusers.de/netcat/) das Passwort an **localhost 3000**.

```shell
bandit14@bandit:~$ cat /etc/bandit_pass/bandit14
4wcYUJFw0k0XLShlDzztnTBHiqxU3b3e
bandit14@bandit:~$ echo 4wcYUJFw0k0XLShlDzztnTBHiqxU3b3e | nc -v localhost 30000
nc: connect to localhost port 30000 (tcp) failed: Connection refused
Connection to localhost 30000 port [tcp/*] succeeded!
Correct!
BfMYroe26WYalil77FoDi9qh59eK5xNr
```

## **Level 15 -> Level 16**

Wir melden uns nun als **bandit15** mit dem Passwort **BfMYroe26WYalil77FoDi9qh59eK5xNr** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort für das nächste Level erhältst du, indem du das Passwort des derzeitigen Levels zu dem Port 30001 auf dem lokal Host einreichst, während du [SSL-Verschlüsselung](https://de.wikipedia.org/wiki/Transport_Layer_Security) benutzt.

Hier benutzen wir **openssl s_client -connect localhost:30001** um eine SSL Verbindung mit **localhost** auf dem Port **30001** herzustellen.

```shell
bandit15@bandit:~$ echo BfMYroe26WYalil77FoDi9qh59eK5xNr | openssl s_client -quiet -connect localhost:30001
depth=0 CN = bandit
verify error:num=18:self signed certificate
verify return:1
depth=0 CN = bandit
verify return:1
Correct!
cluFn7wTiGryunymYOu4RcffSxQluehd
```

## **Level 16 -> Level 17**

Wir melden uns nun als **bandit16** mit dem Passwort **cluFn7wTiGryunymYOu4RcffSxQluehd** an.  
Der Tipp zu diesem Level ist wie folgt:  
Die Zugangsdaten für das nächste Level erhältst du, indem du das Passwort des derzeitigen Levels zu einem **Port** auf dem **lokalen Host** in der Reichweite von **31000 bis 32000** einreichst. Zuerst musst du herausfinden auf welchem dieser Ports ein Server läuft und danach welcher **SSL** benutzt und welcher nicht. Es gibt nur **einen** Server der dir die Zugangsdaten zurücksendet, die anderen senden nur zurück, was du ihnen gesendet hast.

Mit [nmap](https://wiki.ubuntuusers.de/nmap/) können wir nach offenen Ports scannen. Durch die Option **-A** ermittelt [nmap](https://wiki.ubuntuusers.de/nmap/) unter anderem, welche Programme bei den offenen Ports laufen. **-p** legt fest welche Ports gescannt werden sollen. Wenn man **-p** weglässt, werden nur Ports gescannt, welche üblicherweise von Programmen benutzt werden.

```shell
bandit16@bandit:~$ nmap -A -p 31000-32000 localhost
[...]
PORT      STATE SERVICE      VERSION
31046/tcp open  echo
31518/tcp open  ssl/echo
[...]
31691/tcp open  echo
31790/tcp open  ssl/unknown
[...]
31960/tcp open  echo
[...]
```

Nur der Port **31790** scheint kein Echo zurückzugeben. Der nächste Schritt ist analog zum vorherigen Level.

```shell
bandit16@bandit:~$ echo cluFn7wTiGryunymYOu4RcffSxQluehd | openssl s_client -quiet -connect localhost:31790
depth=0 CN = bandit
verify error:num=18:self signed certificate
verify return:1
depth=0 CN = bandit
verify return:1
Correct!
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAvmOkuifmMg6HL2YPIOjon6iWfbp7c3jx34YkYWqUH57SUdyJ
imZzeyGC0gtZPGujUSxiJSWI/oTqexh+cAMTSMlOJf7+BrJObArnxd9Y7YT2bRPQ
[...]
dxviW8+TFVEBl1O4f7HVm6EpTscdDxU+bCXWkfjuRb7Dy9GOtt9JPsX8MBTakzh3
vBgsyi/sN3RqRBcGU40fOoZyfAMT8s1m/uYv52O6IgeuZ/ujbjY=
-----END RSA PRIVATE KEY-----
```

Als nächstes erstellen wir ein neues Verzeichnis **/tmp/shhkey/**. Wir kopieren den SSH-Key und leiten ihn von **echo** zu der Datei **sshkey.private** um, welche dadurch erstellt wird.

```shell
bandit16@bandit:~$ mkdir /tmp/sshkey/
bandit16@bandit:~$ echo "-----BEGIN RSA PRIVATE KEY-----
> MIIEogIBAAKCAQEAvmOkuifmMg6HL2YPIOjon6iWfbp7c3jx34YkYWqUH57SUdyJ
> imZzeyGC0gtZPGujUSxiJSWI/oTqexh+cAMTSMlOJf7+BrJObArnxd9Y7YT2bRPQ
[...]
> dxviW8+TFVEBl1O4f7HVm6EpTscdDxU+bCXWkfjuRb7Dy9GOtt9JPsX8MBTakzh3
> vBgsyi/sN3RqRBcGU40fOoZyfAMT8s1m/uYv52O6IgeuZ/ujbjY=
> -----END RSA PRIVATE KEY-----
> " > /tmp/sshkey/sshkey.private
```

Jetzt können wir uns wie bei Level 13 mithilfe der Datei **sshkey.private** eine Verbindung aufbauen.

```shell
bandit16@bandit:~$ cd /tmp/sshkey
bandit16@bandit:/tmp/sshkey$ ssh -i sshkey.private bandit17@localhost
[...]
This is a OverTheWire game server. More information on http://www.overthewire.org/wargames
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@ WARNING: UNPROTECTED PRIVATE KEY FILE! @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
Permissions 0664 for 'sshkey.private' are too open.
It is required that your private key files are NOT accessible by others.
This private key will be ignored.
Load key "sshkey.private": bad permissions
bandit17@localhost's password:
packet_write_wait: Connection to 127.0.0.1 port 22: Broken pipe
```

Allerdings müssen wir noch die Berechtigungen der Datei ändern, sodass andere diesen nicht lesen können.

```shell
bandit16@bandit:/tmp/sshkey$ ls -la
total 720
drwxrwxr-x 2 bandit16 bandit16 4096 Feb 15 21:50 .
drwxrwx-wt 1929 root root 724992 Feb 15 21:55 ..
-rw-rw-r-- 1 bandit16 bandit16 1676 Feb 15 21:50 sshkey.private
```

Mithilfe von [chmod](https://wiki.ubuntuusers.de/chmod/) können wir Berechtigungen für Dateien ändern. Wir setzen die Berechtigungen so, dass die Datei nur gelesen werden darf, und dies auch nur von dem aktuellen Nutzer **bandit16**.

```shell
bandit16@bandit:/tmp/sshkey$ chmod 400 sshkey.private
bandit16@bandit:/tmp/sshkey$ ls -la
total 720
drwxrwxr-x 2 bandit16 bandit16 4096 Feb 15 21:50 .
drwxrwx-wt 1929 root root 724992 Feb 15 21:56 ..
-r-------- 1 bandit16 bandit16 1676 Feb 15 21:50 sshkey.private
bandit16@bandit:/tmp/sshkey$ ssh -i sshkey.private bandit17@localhost
```

## **Level 17 -> Level 18**

Wir melden uns nun als **bandit17** mit dem Passwort **xLYVMN9WE5zQ5vHacb0sZEVqbrp7nBTn** an.  
Der Tipp zu diesem Level ist wie folgt:  
In dem Heimverzeichnis befinden sich 2 Dateien: **passwords.old** und **passwords.new**. Das Passwort für das nächste Level befindet sich in **passwords.new** und ist die einzige Zeile, welche sich von **passwords.old** unterscheidet.

Mit [diff](https://wiki.ubuntuusers.de/diff/) können wir den Inhalt von zwei Dateien vergleichen.

```shell
bandit17@bandit:~$ ls
passwords.new passwords.old
bandit17@bandit:~$ diff passwords.new passwords.old
42c42
< kfBf3eYk5BPBRzwjqutbbfE887SVc5Yd
---
> 6vcSC74ROI95NqkKaeEC2ABVMDX9TyUr
```

## **Level 18 -> Level 19**

Wir melden uns nun als **bandit18** mit dem Passwort **kfBf3eYk5BPBRzwjqutbbfE887SVc5Yd** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort für das nächste Level befindet sich in einer Datei im Heimverzeichnis namens **readme**. Allerdings hat jemand [.bashrc](https://wiki.ubuntuusers.de/Bash/bashrc/) so modifiziert, dass du ausgeloggt wirst, wenn du dich über SSH einloggst.

Da wir uns nur für den Bruchteil einer Sekunde anmelden können, müssen wir einen anderen Weg kommen, um an das Password innerhalb der Datei **readme** zu kommen.  
Indem wir **cat readme** an den SSH-Verbindungsbefehl hängen, wird dieser ausgeführt, nachdem wir das richtige Password eingegeben haben.

```shell
root@kali:~# ssh bandit18@bandit.labs.overthewire.org -p 2220
[...]
Welcome to OverTheWire!
[...]
Byebye !
Connection to bandit.labs.overthewire.org closed.
root@kali:~#

root@kali:~# ssh bandit18@bandit.labs.overthewire.org -p 2220 cat readme
This is a OverTheWire game server. More information on http://www.overthewire.org/wargames
bandit18@bandit.labs.overthewire.org's password: 
IueksS7Ubh8G3DCwVzrTd8rAVOwq3M5x
root@kali:~#
```

## **Level 19 -> Level 20**

Wir melden uns nun als **bandit19** mit dem Passwort **IueksS7Ubh8G3DCwVzrTd8rAVOwq3M5x** an.  
Der Tipp zu diesem Level ist wie folgt:  
Um Zugang zu dem nächsten Level zu erhalten, solltest du die [setuid](https://de.wikipedia.org/wiki/Setuid) Binärdatei im Heimverzeichnis benutzen. Führe sie ohne Argumente aus, um herauszufinden wie man sie benutzt. Das Passwort für das nächste Level kann wie gewohnt in **/etc/bandit_pass** gefunden werden.

```shell
bandit19@bandit:~$ ls 
bandit20-do
```

Wir haben die Binärdatei **bandit20-do** im Heimverzeichnis. Führen wir diese mal aus.

```shell
bandit19@bandit:~$ ./bandit20-do 
Run a command as another user.
Example: ./bandit20-do id
```

Anscheinend können wir dadurch Befehle als ein anderer Benutzer ausführen. Benutzen wir mal den Befehl der uns vorgeschlagen wird.

```shell
bandit19@bandit:~$ ./bandit20-do id
uid=11019(bandit19) gid=11019(bandit19) euid=11020(bandit20) groups=11019(bandit19)
bandit19@bandit:~$ ./bandit20-do whoami
bandit20
```

Wir können durch die Binärdatei also Befehle als **bandit20** ausführen lassen. Dadurch können wir die Datei **/etc/bandit_pass/bandit20** auslesen lassen, welche das Passwort für **bandit20** enthält.

```shell
bandit19@bandit:~$ ./bandit20-do cat /etc/bandit_pass/bandit20
GbKksEFF4yrVs6il55v6gwY5aVje5f0j
```

## **Level 20 -> Level 21**

Wir melden uns nun als **bandit20** mit dem Passwort **GbKksEFF4yrVs6il55v6gwY5aVje5f0j** an.  
Der Tipp zu diesem Level ist wie folgt:  
Im Heimverzeichnis befindet sich eine [setuid](https://de.wikipedia.org/wiki/Setuid) Binärdatei welche folgendes macht: Sie stellt eine Verbindung zum lokalen Host auf dem Port auf, welchen du in der Kommandozeile als Argument spezifizierst. Sie liest dann den Text denn sie über diese Verbindung gesendet bekommt und vergleicht diesen mit dem Passwort des derzeitigen Levels. Wenn das Passwort korrekt ist, sendet es das Passwort für das nächste Level zurück.

```shell
bandit20@bandit:~$ ls
suconnect
```

Diesmal haben wir die Binärdatei **suconnect** im Heimverzeichnis.

```shell
bandit20@bandit:~$ ./suconnect 
Usage: ./suconnect <portnumber>
This program will connect to the given port on localhost using TCP. If it receives the correct password from the other side, the next password is transmitted back.
```

Indem wir **&** an das Ende eines Befehls hängen, wird dieser im Hintergrund ausgeführt und wir können weitere Befehle eingeben. Durch **nc -l 4444 &** wartet **netcat** im Hintergrund auf eine Verbindung auf Port **4444**.

```shell
bandit20@bandit:~$ nc -l 4444 &
[1] 20058
```

Nun können wir die Binärdatei auch als Hintergrundprozess ausführen lassen und geben den Port **4444** an.

```shell
bandit20@bandit:~$ ./suconnect 4444 &
[2] 20308
```

Durch **fg %1** können wir den Prozess 1 wieder in der Vordergrund holen, in diesem Fall ist das unser Netcat-Listener, zu welchem **suconnect** eine Verbindung aufgebaut hat.

```shell
bandit20@bandit:~$ fg %1
nc -l 4444
```

Jetzt müssen wir nur noch das Passwort des derzeitigen Levels eingeben und Enter drücken. Netcat sendet dieses dann an **suconnect**.

```shell
GbKksEFF4yrVs6il55v6gwY5aVje5f0j
Read: GbKksEFF4yrVs6il55v6gwY5aVje5f0j
Password matches, sending next password
gE269g2h3mw3pwgrj0Ha9Uoqen1c9DGr
[2]- Done ./suconnect 4444
```

## **Level 21 -> Level 22**

Wir melden uns nun als **bandit21** mit dem Passwort **gE269g2h3mw3pwgrj0Ha9Uoqen1c9DGr** an.  
Der Tipp zu diesem Level ist wie folgt:  
Ein Programm wird in regelmäßigen Abständen von [cron](https://de.wikipedia.org/wiki/Cron), dem zeitbasierten Job-Steuerer ausgeführt. Sieh in **/etc/cron.d/** nach der Konfiguration und welches Kommando ausgeführt wird.

Sehen wir uns als erstes an, was wir in **/etc/cron.d** finden können.

```shell
bandit21@bandit:~$ ls /etc/cron.d
cronjob_bandit22 cronjob_bandit23 cronjob_bandit24 popularity-contest
```

**cronjob_bandit22** sieht nach dem aus wonach wir suchen.

```shell
bandit21@bandit:~$ cat /etc/cron.d/cronjob_bandit22
@reboot bandit22 /usr/bin/cronjob_bandit22.sh &> /dev/null
* * * * * bandit22 /usr/bin/cronjob_bandit22.sh &> /dev/null
```

Der Cronjob **cronjob_bandit22** führt regelmäßig das Bash-Skript **/usr/bin/cronjob_bandit22.sh** aus. Sehen wir uns dieses an.

```shell
bandit21@bandit:~$ cat /usr/bin/cronjob_bandit22.sh
#!/bin/bash
chmod 644 /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
```

Das Bash-Skript ändern die Zugriffsrechte der Datei **/tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv**. Was sich in dieser wohl befindet?

```shell
cat /etc/bandit_pass/bandit22 > /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
bandit21@bandit:~$ cat /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
Yk7owGAcWjwMVRwrTesJEwB7WVOiILLI
```

## **Level 22 -> Level 23**

Wir melden uns nun als **bandit22** mit dem Passwort **Yk7owGAcWjwMVRwrTesJEwB7WVOiILLI** an.  
Der Tipp zu diesem Level ist wie folgt:  
Ein Programm wird in regelmäßigen Abständen von [cron](https://de.wikipedia.org/wiki/Cron), dem zeitbasierten Job-Steuerer ausgeführt. Sieh in **/etc/cron.d/** nach der Konfiguration und welches Kommando ausgeführt wird.  
Sich [Shell-Skripts](https://wiki.ubuntuusers.de/Shell/Bash-Skripting-Guide_f%C3%BCr_Anf%C3%A4nger/) anzusehen, welche von anderen Leuten geschrieben wurden, ist eine hilfreiche Fähigkeit. Das Skript in diesem Level ist bewusst so gemacht, dass man es einfach lesen kann. Wenn du Probleme damit hast, zu verstehen was es macht, führe es aus und sieh dir die Debuginformationen an, welche es ausgibt.

Sehen wir uns zuerst wieder die Cronjobs an.

```shell
bandit22@bandit:~$ ls /etc/cron.d
cronjob_bandit22 cronjob_bandit23 cronjob_bandit24 popularity-contest
bandit22@bandit:~$ cat /etc/cron.d/cronjob_bandit23
@reboot bandit23 /usr/bin/cronjob_bandit23.sh &> /dev/null
* * * * * bandit23 /usr/bin/cronjob_bandit23.sh &> /dev/null
```

Der Cronjob **cronjob_bandit23** führt regelmäßig das Bash-Skript **/usr/bin/cronjob_bandit23.sh** aus.

```shell
bandit22@bandit:~$ cat /usr/bin/cronjob_bandit23.sh 
#!/bin/bash

myname=$(whoami)
mytarget=$(echo I am user $myname | md5sum | cut -d ' ' -f 1)
echo "Copying passwordfile /etc/bandit_pass/$myname to /tmp/$mytarget"
cat /etc/bandit_pass/$myname > /tmp/$mytarget
```

Das Skript macht folgendes. In der letzten Zeile kopiert es den Inhalt der Datei **/etc/bandit_pass/bandit23**, welche das Passwort von **bandit23** enthält, nach **/tmp/$mytarget**.  
Die Variable **$mytarget** setzt sich dadurch zusammen, dass von dem String **I am user bandit23** zuerst eine [md5sum](https://wiki.ubuntuusers.de/md5sum/) gebildet wird. Durch [cut](https://wiki.ubuntuusers.de/cut/) mit der Option **-d ' '** werden zuerst alle Leerzeichen entfernt und dann durch die Option **-f 1** wird das erste Feld ausgewählt.  
Um den Dateinamen herauszufinden, müssen wir das Echo Kommando, welches das Skript ausführt, einfach selber ausführen.

```shell
bandit22@bandit:~$ echo I am user bandit23 | md5sum | cut -d ' ' -f 1
8ca319486bfbbc3663ea0fbe81326349
```

Jetzt wo wir den Dateinamen wissen, können wir den Inhalt wie gewohnt ausgeben lassen.

```shell
bandit22@bandit:~$ cat /tmp/8ca319486bfbbc3663ea0fbe81326349
jc1udXuA1tiHqjIsL8yaapX5XIAI6i0n
```

## **Level 23 -> Level 24**

Wir melden uns nun als **bandit23** mit dem Passwort **jc1udXuA1tiHqjIsL8yaapX5XIAI6i0n** an.  
Der Tipp zu diesem Level ist wie folgt:  
Ein Programm wird in regelmäßigen Abständen von [cron](https://de.wikipedia.org/wiki/Cron), dem zeitbasierten Job-Steuerer ausgeführt. Sieh in **/etc/cron.d/** nach der Konfiguration und welches Kommando ausgeführt wird.  
Dieses Level erfordert von dir dein erstes eigenes Skript zu erstellen. Dies ist ein großer Schritt und du solltest Stolz auf dich sein, wenn du dieses Level schaffst!  
Denke daran, dass dein Skript gelöscht wird sobald es ausgeführt wird. Also behalte besser eine Kopie davon.

Sehen wir uns zuerst wieder die Cronjobs an.

```shell
bandit23@bandit:~$ ls /etc/cron.d
cronjob_bandit22 cronjob_bandit23 cronjob_bandit24 popularity-contest

bandit23@bandit:~$ cat /etc/cron.d/cronjob_bandit24
@reboot bandit24 /usr/bin/cronjob_bandit24.sh &> /dev/null
* * * * * bandit24 /usr/bin/cronjob_bandit24.sh &> /dev/null
```

```shell
bandit23@bandit:~$ cat /usr/bin/cronjob_bandit24.sh 
#!/bin/bash

myname=$(whoami)
cd /var/spool/$myname
echo "Executing and deleting all scripts in /var/spool/$myname:"
for i in * .*;
do
if [ "$i" != "." -a "$i" != ".." ];
then
echo "Handling $i"
timeout -s 9 60 ./$i
rm -f ./$i
fi
done
```

Das Skript führt also alle Skripts in **/var/spool/bandit24** aus und löscht diese danach. Erstellen wir uns unter **tmp** ein neuen Verzeichnis, in welchem wir uns ein Skript erstellen können.

```shell
bandit23@bandit:~$ mkdir /tmp/script
bandit23@bandit:~$ cd /tmp/script
```

Mit **vi givemepass.sh** erstellen wir die Bash-Datei **givemepass.sh** und öffnen diese in dem Texteditor [vi](https://de.wikipedia.org/wiki/Vi).

```shell
bandit23@bandit:/tmp/script$ vi givemepass.sh

#!/bin/bash
cat /etc/bandit_pass/bandit24 >> /tmp/script/pass
```

Unser Skript liest den Inhalt der Datei **/etc/bandit_pass/bandit24** aus und sendet diesen an die Datei **/tmp/script/pass**. Jetzt müssen wir noch die Zugriffsrechte unseres Skripts ändern, damit dieses auch vom Skript von **bandit24** ausgeführt werden kann. Außerdem ändern wir auch noch die Zugriffsrechte des von uns erstellten Verzeichnisses **/tmp/script**, damit **bandit24** auch die Rechte hat dort die Datei **pass** zu erstellen.

```shell
bandit23@bandit:/tmp/script$ chmod 777 givemepass.sh 
bandit23@bandit:/tmp/script$ chmod 777 /tmp/script
```

Jetzt können wir unser Skript in das Verzeichnis **/var/spool/bandit24** kopieren, damit es dort von **bandit24** ausgeführt wird.

```shell
bandit23@bandit:/tmp/script$ cp givemepass.sh /var/spool/bandit24
```

Nun müssen wir nur noch warten bis der Cronjob ausgeführt wird.

```shell
bandit23@bandit:/tmp/script$ ls
givemepass.sh pass
bandit23@bandit:/tmp/script$ cat pass
UoMYTrfrBFHyQXmg6gzctqAwOmw1IohZ
```

## **Level 24 -> Level 25**

Wir melden uns nun als **bandit24** mit dem Passwort **UoMYTrfrBFHyQXmg6gzctqAwOmw1IohZ** an.  
Der Tipp zu diesem Level ist wie folgt:  
Ein [Daemon](https://de.wikipedia.org/wiki/Daemon) hört Port **30002** ab und gibt dir dass Passwort für **bandit25**, wenn du ihm das Passwort für **bandit24** und einen geheimen 4-stelligen Zahlencode gibst. Es gibt keinen anderen Weg, als alle der 10000 Möglichkeiten des Zahlencodes durchzugehen. Ein selbstgeschriebenes Skript kann dir dabei helfen. Dies nennt man [Brute-Forcing](https://de.wikipedia.org/wiki/Brute-Force-Methode).

In diesem Level müssen wir ein kleines Brute-Force Skript erstellen. Ich habe Python dafür benutzt.

```shell
bandit24@bandit:~$ mkdir /tmp/brute
bandit24@bandit:~$ cd /tmp/brute
bandit24@bandit:/tmp/brute$ vi brute.py
```

```python
import socket
pin = 0
passwd = 'UoMYTrfrBFHyQXmg6gzctqAwOmw1IohZ '
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('localhost', 30002))
s.recv(1024)
while (pin < 10000):
   print '[+] Versuche: ' + str(pin)
   s.sendall(passwd + str(pin) + '\n')
   data = s.recv(1024)
   print data
   pin += 1
```

```shell
bandit24@bandit:/tmp/brute$ python brute.py
[...]
[+] Versuche: 5440
Correct!
The password of user bandit25 is uNG9O58gUE7snukf3bvZ0rxhtnjzSGzG
```

## **Level 25 -> Level 26**

Wir melden uns nun als **bandit25** mit dem Passwort **uNG9O58gUE7snukf3bvZ0rxhtnjzSGzG** an.  
Der Tipp zu diesem Level ist wie folgt:  
Sich als **bandit26** einzuloggen von **bandit25** aus sollte einfach sein... Allerdings ist die Shell für Benutzer **bandit26** nicht **/bin/bash**, sondern etwas anderes. Finde heraus was es stattdessen ist, wie es funktioniert und wie man daraus herausbrechen kann.

Wie schon bei zwei vorherigen Leveln benutzen wir den SSH-Schlüssel um eine SSH-Verbindung aufzubauen.

```shell
bandit25@bandit:~$ ls
bandit26.sshkey
bandit25@bandit:~$ ssh -i bandit26.sshkey bandit26@localhost 
[...]
Welcome to OverTheWire!
[...]
Connection to localhost closed.
```

Die Verbindung wird sofort geschlossen. Sehen wir uns mal **/etc/passwd** an und suchen da nach **bandit26**.

```shell
bandit25@bandit:~$ cat /etc/passwd | grep bandit26
bandit26:x:11026:11026:bandit level 26:/home/bandit26:/usr/bin/showtext
```

**bandit26** hat nicht **/bin/bash** angegeben, wodurch wir eine Shell bekommen hätten, sondern **/usr/bin/showtext**. Sehen wir uns diese Datei mal an.

```shell
bandit25@bandit:~$ cat /usr/bin/showtext
#!/bin/sh
export TERM=linux
more ~/text.txt
exit 0
```

Hier müssen wir einen Trick anwenden. Und zwar müssen wir das Terminal kleiner machen, damit nicht mehr der gesamte Text angezeigt werden kann. Dadurch wird das **more** getriggert und wir werden nicht sofort herausgeworfen. :)

[![Terminal klein](https://imgur.com/mdcY2gQ.jpg)](https://imgur.com/mdcY2gQ)

Durch das drücken von **v** können wir den Text-Editor [vim](https://de.wikipedia.org/wiki/Vim) öffnen. Jetzt können wir folgendes eingeben um uns das Passwort ausgeben zu lassen:

```vim
:e /etc/bandit_pass/bandit26
```

**Enter** drücken:

```vim
5czgV9L3Xx8JPOyRbXh6lQbmIOWvPT6Z
```

Alternativ können wir uns auch eine Shell geben durch:

```vim
:set shell=/bin/bash
```

und danach

```vim
:shell
```

Nun haben wir eine Shell.

```shell
bandit26@bandit:~$ cat /etc/bandit_pass/bandit26
5czgV9L3Xx8JPOyRbXh6lQbmIOWvPT6Z
```

## **Level 26 -> Level 27**

Wir melden uns nun als **bandit26** mit dem Passwort **5czgV9L3Xx8JPOyRbXh6lQbmIOWvPT6Z** an.  
Der Tipp zu diesem Level ist wie folgt:  
Denkt beim anmelden per SSH an das was wir im letzen Level herausgefunden haben...

Damit die Verbindung nicht sofort geschlossen wird, müssen wir die letzten Schritte des vorherigen Levels wiederholen. Das Terminal klein ziehen, **v** drücken um [vim](https://de.wikipedia.org/wiki/Vim) zu öffnen und dann:

```vim
:set shell=/bin/bash
:shell
```

Wir haben nun wieder eine Shell. :)

```shell
bandit26@bandit:~$ ls -l
total 12
-rwsr-x--- 1 bandit27 bandit26 7296 Oct 16  2018 bandit27-do
-rw-r----- 1 bandit26 bandit26  258 Oct 16  2018 text.txt

```

Wir haben hier also 2 Dateien. **text.txt** enthält nur das Banner. Führen wir also mal **bandit27-do** aus.

```shell
bandit26@bandit:~$ ./bandit27-do 
Run a command as another user.
  Example: ./bandit27-do id
  ```

Wie wir bei **ls -l** schon sehen konnten, hat diese Binary das [setuid bit](https://www.hackingarticles.in/linux-privilege-escalation-using-suid-binaries/) gesetzt! Wir können hiermit also Befehle als ein anderer User ausführen. In diesem Falle ist dies **bandit27**!

```shell
bandit26@bandit:~$ ./bandit27-do cat /etc/bandit_pass/bandit27
3ba3118a22e93127a4ed485be72ef5ea
```

Wir können nun also einfach das Passwort für **bandit27** auslesen!

## **Level 27 -> Level 28**

Wir melden uns nun als **bandit27** mit dem Passwort **3ba3118a22e93127a4ed485be72ef5ea** an.  
Der Tipp zu diesem Level ist wie folgt:  
Das Passwort für das nächste Level könnt ihr in dem [Git repository](https://de.atlassian.com/git/tutorials/setting-up-a-repository/git-clone) **ssh://bandit27-git@localhost/home/bandit27-git/repo** finden. Das Passwort für dieses repository ist dasselbe wie für den User **bandit27**.

Erstellen wir uns erstmal einen Ordner innerhalb von **tmp**, in diesen wir die repository klonen können.  
Das Passwort ist dasselbe wie vom User **bandit 27**, also **3ba3118a22e93127a4ed485be72ef5ea**

```shell
bandit27@bandit:~$ mkdir /tmp/m10x2/
bandit27@bandit:~$ cd /tmp/m10x2/
bandit27@bandit:/tmp/m10x2$ git clone ssh://bandit27-git@localhost/home/bandit27-git/repo

Cloning into 'repo'...
Could not create directory '/home/bandit27/.ssh'.
The authenticity of host 'localhost (127.0.0.1)' can't be established.
ECDSA key fingerprint is SHA256:98UL0ZWr85496EtCRkKlo20X3OPnyPSB5tB5RPbhczc.
Are you sure you want to continue connecting (yes/no)? yes
Failed to add the host to the list of known hosts (/home/bandit27/.ssh/known_hosts).
This is a OverTheWire game server. 
More information on http://www.overthewire.org/wargames
bandit27-git@localhost's password: 

remote: Counting objects: 3, done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 0 (delta 0)
Receiving objects: 100% (3/3), done.
```

Die Suche nach dem Passwort ist nun sehr einfach...

```shell
bandit27@bandit:/tmp/m10x2$ lsrepo
bandit27@bandit:/tmp/m10x2$ cd repo
bandit27@bandit:/tmp/m10x2/repo$ ls -alh
total 16K
drwxr-sr-x 3 bandit27 root 4.0K May 13 10:32 .
drwxr-sr-x 3 bandit27 root 4.0K May 13 10:32 ..
drwxr-sr-x 8 bandit27 root 4.0K May 13 10:32 .git
-rw-r--r-- 1 bandit27 root 68 May 13 10:32 README
bandit27@bandit:/tmp/m10x2/repo$ cat README
The password to the next level is: 0ef186ac70e04ea33b4c1853d2526fa2
```

## **Level 28 -> Level 29**

Wir melden uns nun als **bandit28** mit dem Passwort **0ef186ac70e04ea33b4c1853d2526fa2** an.  
Der Tipp zu diesem Level ist wie folgt:  
Diesmal ist das Passwort zensiert... Können wir dies rückgängig machen?

Der Anfang ist analog zum vorherigen Level

```shell
bandit28@bandit:~$ mkdir /tmp/m10x3
bandit28@bandit:~$ cd /tmp/m10x3
bandit28@bandit:/tmp/m10x3$ git clone ssh://bandit28-git@localhost/home/bandit28-git/repo

Cloning into 'repo'...
Could not create directory '/home/bandit28/.ssh'.
The authenticity of host 'localhost (127.0.0.1)' can't be established.
ECDSA key fingerprint is SHA256:98UL0ZWr85496EtCRkKlo20X3OPnyPSB5tB5RPbhczc.
Are you sure you want to continue connecting (yes/no)? yes
Failed to add the host to the list of known hosts (/home/bandit28/.ssh/known_hosts).
This is a OverTheWire game server. 
More information on http://www.overthewire.org/wargames

bandit28-git@localhost's password: 
remote: Counting objects: 9, done.
remote: Compressing objects: 100% (6/6), done.
remote: Total 9 (delta 2), reused 0 (delta 0)
Receiving objects: 100% (9/9), done.
Resolving deltas: 100% (2/2), done.

bandit28@bandit:/tmp/m10x3$ cd repo
bandit28@bandit:/tmp/m10x3/repo$ ls 
README.md

bandit28@bandit:/tmp/m10x3/repo$ cat README.md

# Bandit NotesSome notes for level29 of bandit.

## credentials
- username: bandit29
- password: xxxxxxxxxx
```

Das Passwort ist diese mal zensiert worden.  
Wenn wir Glück haben, gibt es aber eine ältere Version der repository, in welcher das Passwort noch nicht zensiert war.

Mit dem Befehl **git log** können wir uns die Änderungen an der repository anzeigen lassen.

```shell
bandit28@bandit:/tmp/m10x3/repo$ git log
commit 073c27c130e6ee407e12faad1dd3848a110c4f95
Author: Morla Porla <morla@overthewire.org>
Date: Tue Oct 16 14:00:39 2018 +0200
fix info leak

commit 186a1038cc54d1358d42d468cdc8e3cc28a93fcb
Author: Morla Porla <morla@overthewire.org>
Date: Tue Oct 16 14:00:39 2018 +0200
add missing data

commit b67405defc6ef44210c53345fc953e6a21338cc7
Author: Ben Dover <noone@overthewire.org>
Date: Tue Oct 16 14:00:39 2018 +0200
initial commit of README.md
```

Bei dem letzten commit wurde ein **info leak** gefixt.  
Nehmen wir uns vorherigen commit vor, bei dem wohl der info leak noch vorhanden ist.

```shell
bandit28@bandit:/tmp/m10x3/repo$ git checkout 186a1038cc54d1358d42d468cdc8e3cc28a93fcb
Note: checking out '186a1038cc54d1358d42d468cdc8e3cc28a93fcb'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by performing another checkout.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -b with the checkout command again. 
Example:git checkout -b <new-branch-name>

HEAD is now at 186a103... add missing data
```

Wir haben nun eine ältere Version des Projektes vor uns. Lesen wir noch einmal die **README.md** Datei.

```shell
bandit28@bandit:/tmp/m10x3/repo$ ls
README.md
bandit28@bandit:/tmp/m10x3/repo$ cat README.md 
# Bandit Notes
Some notes for level29 of bandit.

## credentials
- username: bandit29
- password: bbc96594b4e001778eee9975372716b2
```

## **Level 29 -> Level 30**

Wir melden uns nun als **bandit29** mit dem Passwort **bbc96594b4e001778eee9975372716b2** an.  
Der Tipp zu diesem Level ist wie folgt:

Dieses Mal müssen wir das Passwort in einer anderen [Branch](https://git-scm.com/book/de/v1/Git-Branching-Was-ist-ein-Branch%3F) finden!

Der Anfang ist wieder analog...

```shell
[...]
bandit29@bandit:/tmp/m10x4/repo$ cat README.md
# Bandit Notes
Some notes for bandit30 of bandit.

## credentials
- username: bandit30
- password: <no passwords in production!>
```

Es steht dort, dass anscheinend keine Passwörter in der Produktionsbranch benutzt werden.  
Wir befinden uns in der Branch **master**. Mal sehen welche anderen wir finden können...

```shell
bandit29@bandit:/tmp/m10x4/repo$ git branch
* master

bandit29@bandit:/tmp/m10x4/repo$ git branch -r
origin/HEAD -> origin/master
origin/dev
origin/master
origin/sploits-dev
```

die **dev** Branch sieht vielversprechend aus...

```shell
bandit29@bandit:/tmp/m10x4/repo$ git checkout dev
Branch dev set up to track remote branch dev from origin.
Switched to a new branch 'dev'

bandit29@bandit:/tmp/m10x4/repo$ cat README.md
# Bandit Notes
Some notes for bandit30 of bandit.

## credentials
- username: bandit30
- password: 5b90576bedb2cc04c86a9e924ce42faf
```

## **Level 30 -> Level 31**

Wir melden uns nun als **bandit30** mit dem Passwort **5b90576bedb2cc04c86a9e924ce42faf** an.  
Der Tipp zu diesem Level ist wie folgt:

Bei einem Tag dieser repository wurde eine Dateiname verändert.

Der Anfang ist wieder analog...

```shell
bandit30@bandit:/tmp/m10x5/repo$ cat README.md
just an epmty file... muahaha
```

Dieses mal haben wir keine **Commits** oder anderen **Branches** die uns weiterhelfen.  
Allerdings haben wir einen **tag**. **Tags** sind quasi **Branches** ohne eine **commit history**.  
Es ist ein **tag** vorhanden, und zwar mit dem Namen **secret**.  
Mit dem Befehl **git show --name-only** können wir uns dann die Namen der Dateien anzeigen lassen, welche umbenannt wurden.

```shell
bandit30@bandit:/tmp/m10x5/repo$ git tag
secret
bandit30@bandit:/tmp/m10x5/repo$ git show --name-only secret
47e603bb428404d265f59c42920d81e5
```

Eine Datei hatte wohl das Passwort als Namen. :)

## **Level 31 -> Level 32**

Wir melden uns nun als **bandit31** mit dem Passwort **47e603bb428404d265f59c42920d81e5** an.  
Der Tipp zu diesem Level ist wie folgt:

README.md enthält die nötigen Anweisungen, die man vornehmen muss.

Der Anfang ist wieder analog...

```shell
bandit31@bandit:/tmp/m10x6/repo$ ls -alh
total 20K
drwxr-sr-x 3 bandit31 root 4.0K May 13 11:42 .
drwxr-sr-x 3 bandit31 root 4.0K May 13 11:42 ..
drwxr-sr-x 8 bandit31 root 4.0K May 13 11:42 .git
-rw-r--r-- 1 bandit31 root 6 May 13 11:42 .gitignore
-rw-r--r-- 1 bandit31 root 147 May 13 11:42 README.md

bandit31@bandit:/tmp/m10x6/repo$ cat .gitignore
*.txt

bandit31@bandit:/tmp/m10x6/repo$ cat README.md 
This time your task is to push a file to the remote repository.

Details:
File name: key.txt
Content: 'May I come in?'
Branch: master
```

Erstellen wir wie gefordert die Text Datei und fügen sie zu der repository hinzu.

```shell
bandit31@bandit:/tmp/m10x6/repo$ touch key.txt
bandit31@bandit:/tmp/m10x6/repo$ echo "May I come in?" > key.txt
bandit31@bandit:/tmp/m10x6/repo$ git add key.txt
The following paths are ignored by one of your .gitignore files:key.txt
Use -f if you really want to add them.
```

**.gitignore** verhindert, dass wir irgendwelche **.txt** Dateien hinzufügen können! Um dies zu umgehen, können wir entweder **-f** als Parameter nutzen, oder **.gitignore** löschen.

```shell
bandit31@bandit:/tmp/m10x6/repo$ rm .gitignore 
bandit31@bandit:/tmp/m10x6/repo$ git add key.txt
bandit31@bandit:/tmp/m10x6/repo$ git commit -m "file upload"
[master 0a1e81b] file upload
1 file changed, 1 insertion(+)
create mode 100644 key.txt
```

**"file upload"** ist in diesem Fall nur ein Kommentar, bei dem wir mitteilen können was geändert wurde.  
Der **commit** ist jetzt vorbereitet und wir müssen ihn nur noch **push**en.

```shell
bandit31@bandit:/tmp/m10x6/repo$ git push origin master

Could not create directory '/home/bandit31/.ssh'.
The authenticity of host 'localhost (127.0.0.1)' can't be established.
ECDSA key fingerprint is SHA256:98UL0ZWr85496EtCRkKlo20X3OPnyPSB5tB5RPbhczc.
Are you sure you want to continue connecting (yes/no)? yes
Failed to add the host to the list of known hosts (/home/bandit31/.ssh/known_hosts).
This is a OverTheWire game server. 
More information on http://www.overthewire.org/wargames

bandit31-git@localhost's password: 
Counting objects: 3, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 321 bytes | 0 bytes/s, done.
Total 3 (delta 0), reused 0 (delta 0)

remote: ### Attempting to validate files... ####
remote: 
remote: .oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.
remote: 
remote: Well done! Here is the password for the next level:
remote: 56a9bf19c63d650ce78e6ec0354ee45e
remote: 
remote: .oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.
remote: To ssh://localhost/home/bandit31-git/repo
! [remote rejected] master -> master (pre-receive hook declined)
error: failed to push some refs to 'ssh://bandit31-git@localhost/home/bandit31-git/repo'
```

## **Level 32 -> Level 33**

Wir melden uns nun als **bandit32** mit dem Passwort **56a9bf19c63d650ce78e6ec0354ee45e** an.  
Der Tipp zu diesem Level ist wie folgt:

Das vorherige Level war nun erstmal das letzte git Level. Hier müssen wir wieder eine Shell Escape machen!

Wir sind in einer sh shell gefangen!  
Wir können nun einfach $0 eingeben um eine Standard Bash shell zu erhalten.

```shell
WELCOME TO THE UPPERCASE SHELL>> $0
$ cat /etc/bandit_pass/bandit33
c9c3199ddf4121b10cf581a98d51caee
```

## **Level 33 -> Level 34**

Wir melden uns nun als **bandit33** mit dem Passwort **c9c3199ddf4121b10cf581a98d51caee** an.  
Der Tipp zu diesem Level ist wie folgt:

Wird ergänzt, wenn Level verfügbar ist

In diesem Level befindet sich nur die Datei **README.txt**.

```shell
bandit26@bandit:~$ ls
README.txt text.txt
bandit26@bandit:~$ cat README.txt 
Congratulations on solving the last level of this game!

At this moment, there are no more levels to play in this game. However, we are constantly working
on new levels and will most likely expand this game with more levels soon.
Keep an eye out for an announcement on our usual communication channels!
In the meantime, you could play some of our other wargames.

If you have an idea for an awesome new level, please let us know!
```