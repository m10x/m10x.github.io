---
title: "HackThis - Alle Main Level Walkthrough | Tipps + Anleitung"
date: 2018-01-23T20:39:56+01:00
toc: false
images:
tags:
  - hackthis
  - ctf
  - german
---

Walkthrough | Tipps + Anleitung zu allen Main Levels 1-10 von [HackThis](https://www.hackthis.co.uk/levels/Main). Die Main Level gehören zu dein einfacheren von [HackThis](https://www.hackthis.co.uk/levels/Main).

## **Video**

[![Kurzes Video Walkthrough ohne Erklärunge](http://img.youtube.com/vi/G7O8Zh9dPbk/0.jpg)](http://www.youtube.com/watch?v=G7O8Zh9dPbk)

## **Level 1**

Drücke **Strg + u** um den Quelltext der Seite anzuzeigen. Achte auf Kommentare im Quelltext.

Zuerst sehen wir uns den Quelltext an, indem wir **Strg + u** drücken.  
In Zeile 29 finden wir folgendes:  **<!-- username: in, password: out -->**  
Einfacher geht es wohl kaum.

## **Level 2**

Suche in dem Quelltext nach einem bestimmten Begriff.

Nutzername und Passwort sind diesmal wieder im Quelltext zu finden. Diesmal aber nicht direkt am Anfang des Quelltextes.  
Die richtigen Zeilen finden wir schnell, wenn wir z.B. nach **levelform** suchen.  
In Zeile 868 und 870 können wir sehen, dass jeweils hinter dem **user** und **password** Label, der richtige Nutzername und das richtige Password in schwarz stehen.

```html
<label for="user">Username:</label> <span style="color: #000000">resu</span>
<label for="user">Password:</label> <span style="color: #000000">ssap</span>
```

## **Level 3**

Suche nach Javascript im Quelltext und sieh dir die Funktionen an.

Wieder stehen Nutzername und Passwort als Klartext im Quelltext.  
In Zeile 41 sehen wir eine interessante Javascript-Funktion.

```js
if(document.getElementById('user').value == 'heaven' && document.getElementById('pass').value == 'hell') { }
else { e.preventDefault(); alert('Incorrect login') }
```

## **Level 4**

Die Lösung bei diesem Level ist in einem verstecken Feld (= hidden field) zu finden.  
Suche in dem Quelltext doch mal nach **hidden**.

Durchsuchen wir den Quelltext doch zuerst wieder nach **level-form**.  
Diesmal ist ein verstecktes Feld zu sehen mit dem Wert **../../extras/ssap.xml**.  
Besuchen wir doch mal [https://www.hackthis.co.uk/levels/extras/ssap.xml](https://www.hackthis.co.uk/levels/extras/ssap.xmlhttps://www.hackthis.co.uk/levels/extras/ssap.xml)...

## **Level 5**

Sieh dir wieder den Quelltext bei der Klasse **level-form** an.

Zuerst suchen wir wieder im Quelltext nach **level-form**. Von Zeile 867 bis 871 ist eine Javascript-Funktion zu sehen.

```javascript
var pass;
pass=prompt("Password","");
if (pass=="9286jas") {
window.location.href="/levels/main/5?pass=9286jas";
}
```

Das Passwort ist also **9286jas**.

## **Level 6**

Du musst einen Weg senden ein POST-Request zu senden, mit dem Parameter **user=Ronald**.  
Alternativ kannst du auch die Form verändern (F12).

Hier gibt es zwei verschiedene Lösungsmöglichkeiten.  
Zum einen können wir ein POST-Request senden mit dem Parameter **user=Ronald**.  
Um sehr einfach ein POST request zu machen, kann man ein Browser-Plugin dafür benutzen.  
Firefox Quantum: [New Hackbar](https://addons.mozilla.org/de/firefox/addon/new-hackbar/)  
Firefox Älter: [Hackbar](https://addons.mozilla.org/de/firefox/addon/hackbar/)  
Google Chrome: [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop)  
Ich verwende Firefox Quantum mit dem "New Hackbar" Plugin dafür benutzen, was du benutzt, ist natürlich dir überlassen.

Eine weitere Möglichkeit ist es den Inspektor von Firefox ( oder auch Google Chrome ) zu benutzen. Drücke **F12** oder **Rechtsklick** > **Element untersuchen** um diesen zu öffnen.  
Nun können wir bei dem Listen-Feld den aktuell ausgewählten Benutzernamen mit einem Doppelklick bearbeiten und **Ronald** einsetzen.  
Danach einfach auf **Submit** klicken.

## **Level 7**

Die Seite mit der Text-Datei kann nicht von Suchmaschinen gefunden werden, da Bots ausgeschlossen wurden.  
Wie wurde das wohl gemacht?

Der Tipp ist folgender: Die Seite mit der Text-Datei kann nicht von Suchmaschinen gefunden werden, da Bots ausgeschlossen wurden.  
Sehen wir uns doch mal die [robots.txt](https://www.hackthis.co.uk/robots.txt) Datei an, welche bestimmte User-Agenten ausschließen kann, z.B. auch die Bots von Suchmaschinen.

```shell
[...]
Disallow: /levels/extras/userpass.txt
[...]
```

Das sieht doch nach etwas aus mit dem wir was anfangen können.  
Wenn wir [https://www.hackthis.co.uk/levels/extras/userpass.txt](https://www.hackthis.co.uk/levels/extras/userpass.txt) öffnen finden wir

Nutzername: **48w3756**  
Passwort: **u3qh458**

## **Level 8**

Es ist wieder ein verstecktes Feld vorhanden. Rechne die Binär Zahl zu Hexadezimal um.

Es ist wieder ein verstecktes Feld vorhanden. Suche nach **hidden** um es zu finden.  
Der Wert des Feldes leitet uns zu [https://www.hackthis.co.uk/levels/extras/secret.txt](https://www.hackthis.co.uk/levels/extras/secret.txt) weiter.

```shell
1011 0000 0000 1011
1111 1110 1110 1101
```

Was wir gefunden haben sind 2 Binär Zahlen. Rechnen wir diese doch in Hexadezimal um.

```shell
1011 0000 0000 1011 = B00B
1111 1110 1110 1101 = FEED
```

## **Level 9**

Bei diesem Level geht es darum die richtige Email herauszufinden...  
Oder deine eigene Email zur richtigen zu machen.

Hier gibt es wieder verschiedene Möglichkeiten.  
Wenn wir uns den Quelltext ansehen können wir ein verstecktes Feld mit dem Wert **admin@hackthis.co.uk** finden. Welche die "richtige" Email ist um das Level zu bestehen.

Allerdings können wir den Wert dieses Feldes auch mit dem Inspektor ( F12 bei Firefox / Chrome ) bearbeiten, sodass unsere eigene Email die richtige ist.

Ein weiterer Weg ist es mit Burp unseren HTTP-Verkehr abzufangen und dann mit Hilfe von Burp den Wert von **email2** zu unserer Email zu verändern.

## **Level 10**

Suche nach einem versteckten Feld und finde eine Möglichkeit die Hashes zu cracken.

Im Quelltext können wir wieder ein verstecktes Feld finden, mit dem Wert **level10pass.txt**.  
Bei [https://www.hackthis.co.uk/levels/extras/level10pass.txt](https://www.hackthis.co.uk/levels/extras/level10pass.txt) finden wir 2 Hashes.

**69bfe1e6e44821df7f8a0927bd7e61ef208fdb25deaa4353450bc3fb904abd52**  
und  
**f1abe1b083d12d181ae136cfc75b8d18a8ecb43ac4e9d1a36d6a9c75b6016b61**

Mit Hilfe von [https://crackstation.net/](https://crackstation.net/) können wir diese beiden Hashes cracken lassen.  
Das Ergebnis ist:

**carl**  
**guess**

Somit haben wir alle Main Level von Hack This geschafft! :D  
Vielen Dank für's lesen.