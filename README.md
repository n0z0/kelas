# kelas
Kelas Pentest Indonesia

# Introduction to Ethical Hacking and Iinformation Security
## Ethical Hacking

<iframe width="560" height="315"
    src="https://www.youtube.com/embed/olqJhwbBZdA"
    title="Ethical Hacking Explained in 2 Minutes"
    frameborder="0"
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
    allowfullscreen>
</iframe>

<iframe width="560" height="315"
  src="https://www.youtube.com/embed/XLvPpirlmEs"
  title="What Is Ethical Hacking? (Ethical Hacking In 8 Minutes)"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen>
</iframe>

## Elements of Information Security

<iframe width="560" height="315"
  src="https://www.youtube.com/embed/6jli-yKbk0A"
  title="What Are The Elements Of Information Security?"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen>
</iframe>

## Cyber Kill Chain Methodology

<iframe width="560" height="315"
    src="https://www.youtube.com/embed/s6XgCJ141Ww"
    title="The Cyber Kill Chain Explained: Map & Analyze Cyber Attacks!"
    frameborder="0"
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
    allowfullscreen>
</iframe>

## Risk Management

<iframe width="560" height="315"
    src="https://www.youtube.com/embed/mJIhpbzaF-o"
    title="The Ultimate Guide To Risk Management in Cybersecurity"
    frameborder="0"
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
    allowfullscreen>
</iframe>

## Incident Management

<iframe width="560" height="315"
  src="https://www.youtube.com/embed/mpsCsmM0vVQ"
  title="How Incident Response Works in Cybersecurity"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen>
</iframe>

## Information Assurance

<iframe width="560" height="315"
  src="https://www.youtube.com/embed/RtmuuWg_dkQ"
  title="What is Information Assurance"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen>
</iframe>

## MITRE ATT&CK Framework

<iframe width="560" height="315"
    src="https://www.youtube.com/embed/rA7NWxP3I8M"
    title="YouTube video player"
    frameborder="0"
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
    allowfullscreen>
</iframe>


# Understanding Footprinting and Reconnaisance in Cybersecurity
## Reconnaissance
## Hybrid Footprinting
## Footprinting
## Active Footprinting
## Passive Footprinting

# Network Scanning and Enumeration in Information Security

[NetScan](./netscan/)

## Standar Procedures for Network Scanning
## Performing Scanning Beyond IDS and Firewall
## Information Security Controls
## Relevant Laws and Regulations
## Introduction to Network Scanning
## Practical Exercises
## Enumeration and Countermeasures

# Understanding Vulnerability Analysis in Cybersecurity
## Best Practices for Vulnerability Analysis
## Vulnerability Assessment Tools
## Vulnerability Testing Research
## Vulnerability Analysis
## Introduction to Vulnerabilities
## Vulnerability Hunting
## Case Studies and Examples

# Understanding Web Server Security and Attack Mitigation
## Common Web Server Attacks
## Web Server Security Tools
## Website Defacement
## Web Server Attack Tools
## Patch Management
## Patch Management Tools
## Web Server Attack Methodology
## Web Server Operation

# Understanding Web Application Security and Attack Mitigation
## Webhooks and Web Shell
## Web Application Hacking Methodology
## Web Application Security Risks
## Web API
## Web API Hacking Methodology
## Web Application Security
## Web Application Architecture
## Web Application Threats

## Day 2 Slides: OWASP Top 10 Threats & Mitigation
- Slides: `webmitigation/slide/day2_owasp_top10_threats_mitigation.tex`
- PDF build output: `webmitigation/slide/day2_owasp_top10_threats_mitigation.pdf`

### Compile LaTeX to PDF
- Windows (MiKTeX)
  - Install MiKTeX (ensure on-demand packages allowed)
  - Open PowerShell, run:
    - `cd webmitigation/slide`
    - `pdflatex -interaction=nonstopmode -halt-on-error day2_owasp_top10_threats_mitigation.tex`
    - Run twice to resolve references
- Linux (TeX Live)
  - Install TeX Live full or required packages (`beamer`, `pgf`, `pgf-pie`, `listings`, `fontawesome5`)
  - `cd webmitigation/slide && pdflatex -interaction=nonstopmode -halt-on-error day2_owasp_top10_threats_mitigation.tex`
- Docker (isolated build)
  - `docker run --rm -v %CD%/webmitigation/slide:/work -w /work texlive/texlive pdflatex -interaction=nonstopmode -halt-on-error day2_owasp_top10_threats_mitigation.tex`
  - On Linux/macOS replace `%CD%` with `$PWD`

### Run Labs (Docker Compose)
- Prereqs: Docker Desktop, browser (Chrome/Firefox), optional OWASP ZAP/Burp
- Start lab targets:
  - `cd webmitigation/slide`
  - `docker compose up -d` (see “Docker Compose: Lab Mandiri” slide)
- Access targets:
  - DVWA: `http://localhost:8080`
  - Juice Shop: `http://localhost:3000`
  - WebGoat: `http://localhost:8081/WebGoat`

### Run Go Sample Service (Security Headers / Rate Limit / SSRF)
- Create a small Go service or reuse snippets in slides (A05, A07, A10)
- Quick run example:
  - `go run .` (ensure `go.mod` exists)
  - Test with `curl -I http://localhost:PORT` and verify headers
- Security checks for Go modules:
  - `go mod tidy && go list -m -u all`
  - `govulncheck ./...` and `go mod verify`

### Use Proxy (ZAP/Burp) with Browser + ESM
- Set browser proxy to ZAP/Burp listener (e.g., `127.0.0.1:8080`)
- Import ZAP/Burp CA cert into the browser to intercept HTTPS
- Open target app via browser and exercise flows from slides
- In ZAP: use Baseline or Active Scan against `http://localhost:...`
- In Burp: define Target Scope, use Repeater/Intruder for lab steps

# Understanding Session Hijacking Attacks and Countermeasures
## Spoofing
## Session Hijacking Tools
## Session Hijacking Detection Methods
## Session Hijacking Prevention Tools
## Session Fixation Attack
## Network Level Session Hijacking
## Session Hijacking
## Session Replay Attacks
## Application-Level Session Hijacking
## Man-in-the-Browser Attack
## Types of Session Hijacking

# Understanding SQL Injection Attacks and Mitigation
## SQL Injection Methodology
## SQL Injection
## Blind SQL Injection
## SQL Injection Detection Tools
## SQL Injection Tools
## Types of SQL Injection
## Signature Evasion Techniques

# Evading IDS and Firewalls: Techniques and Countermeasures
## Introduction to Evasion Techniques
## Firewall, IDS Rules, and Patch Management
## Bypassing Firewall Rules Using Tunneling
## Bypassing Windows IDS
## Bypassing Windows Firewall
## Bypassing Antivirus

# System Hacking Techniques and Countermeasures
## Introduction to System Hacking
## Offline & Online Password Cracking
## Privilege Escalation
## Privilege Escalation in Linux Machines
## Buffer Overflow
## Clearing Windows and Linux Machine Logs

# Understanding Malware Threats and Mitigation Strategies
## Anti-Trojan and Antivirus Software
## Exploit Kits and Exploits
## Malware Analysis
## Virus Detection Methods
## Malware Overview
## Trojan
## Virus
## Malware Detection Tools
## Ransomware

# Understanding Denial-of-Service (DoS) and Distributed Denial-of-Service (DDoS) Attacks
## DoS/DDoS Protection Tools
## DoS/DDoS Attack Detection Techniques
## DoS/DDoS Attack Techniques
## DoS and DDoS Attacks Overview
## Botnets
## DoS/DDoS Attack Tools

# Introduction to Cryptography and Encryption Techniques
## Email Encryption
## Encryption Algorithms
## Public Key Infrastructure (PKI)
## Cryptography Tools
## Cryptography Overview
## MD5 and MD6 Hash Calculators
## Disk Encryption
