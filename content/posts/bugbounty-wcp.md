---
title: 403 Help is Forbidden - Web Cache Poisoning in the Wild
date: 2025-08-08T04:00:56+01:00
toc: true
images: 
tags:
  - Bugbounty
  - Web
  - English
---

## TLDR

Wrote a [Web Cache Poisoning scanner](https://github.com/Hackmanit/Web-Cache-Vulnerability-Scanner). Ran it against bug bounties. Found **high hundreds** of Cache Poisoned Denial of Services (affecting e.g. popular SaaS products, financial and health websites, as well as European governmental websites) as well as two Cached Reflected XSS (= Stored XSS). Cached Poisoned Denial of Service didn't generate much interest, but XSS earned me good money.

## Prolog
In 2021, I chose web cache poisoning as the topic for [my bachelor's thesis](https://hackmanit.de/images/download/thesis/Automated-Scanning-for-Web-Cache-Poisoning-Vulnerabilities.pdf). I chose it because of the excellent papers [Practical Web Cache Poisoning](https://portswigger.net/research/practical-web-cache-poisoning) (2018, James Kettle), [Your Cache Has Fallen: Cache-Poisoned Denial-of-Service Attack](https://cpdos.org/) (2019, Hoai Viet Nguyen, Luigi Lo Iacono, and Hannes Federrath)​, and [Web Cache Entanglement: Novel Pathways to Poisoning](https://portswigger.net/research/web-cache-entanglement) (2020, James Kettle)​. I gathered all the known web cache poisoning techniques, sorted them into categories, and bundeled them into a scanner: the [Web Cache Vulnerability Scanner (wcvs)](https://github.com/Hackmanit/Web-Cache-Vulnerability-Scanner). Additionally, I scanned 51 of the top 1000 websites for web cache poisoning. However, the results weren't that great. There were too many false positives, and only 11 instances of non-malicious cached content injections.

Nevertheless, throughout the years, I maintained the scanner, fixing bugs and adding and improving techniques. The positive feedback was motivating, and a few bug bounty hunters thanked me for helping them earn good money using the scanner. Every once in a while, I thought about running the scanner against some bug bounties again. This finally led me to my next free time project: Improving the scanner and running it against bug bounties!

## Automation is Key

## Results

## Conclusion
