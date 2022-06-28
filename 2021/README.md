# 2021 List

---
## CVE-2021-46422 (2022-04-27T13:15:00)
> Telesquare SDT-CW3B1 1.1.0 is affected by an OS command injection vulnerability that allows a remote attacker to execute OS commands without any authentication.
- [Chocapikk/CVE-2021-46422](https://github.com/Chocapikk/CVE-2021-46422)

---
## CVE-2021-44228 (2021-12-10T10:15:00)
> Apache Log4j2 2.0-beta9 through 2.15.0 (excluding security releases 2.12.2, 2.12.3, and 2.3.1) JNDI features used in configuration, log messages, and parameters do not protect against attacker controlled LDAP and other JNDI related endpoints. An attacker who can control log messages or log message parameters can execute arbitrary code loaded from LDAP servers when message lookup substitution is enabled. From log4j 2.15.0, this behavior has been disabled by default. From version 2.16.0 (along with 2.12.2, 2.12.3, and 2.3.1), this functionality has been completely removed. Note that this vulnerability is specific to log4j-core and does not affect log4net, log4cxx, or other Apache Logging Services projects.
- [aws-samples/kubernetes-log4j-cve-2021-44228-node-agent](https://github.com/aws-samples/kubernetes-log4j-cve-2021-44228-node-agent)
- [intel-xeon/CVE-2021-44228---detection-with-PowerShell](https://github.com/intel-xeon/CVE-2021-44228---detection-with-PowerShell)

---
## CVE-2021-43229 (2021-12-15T15:15:00)
> Windows NTFS Elevation of Privilege Vulnerability This CVE ID is unique from CVE-2021-43230, CVE-2021-43231.
- [Citizen13X/CVE-2021-43229](https://github.com/Citizen13X/CVE-2021-43229)

---
## CVE-2021-42697 (2021-11-02T22:15:00)
> Akka HTTP 10.1.x before 10.1.15 and 10.2.x before 10.2.7 can encounter stack exhaustion while parsing HTTP headers, which allows a remote attacker to conduct a Denial of Service attack by sending a User-Agent header with deeply nested comments.
- [cxosmo/CVE-2021-42697](https://github.com/cxosmo/CVE-2021-42697)

---
## CVE-2021-41773 (2021-10-05T09:15:00)
> A flaw was found in a change made to path normalization in Apache HTTP Server 2.4.49. An attacker could use a path traversal attack to map URLs to files outside the directories configured by Alias-like directives. If files outside of these directories are not protected by the usual default configuration "require all denied", these requests can succeed. If CGI scripts are also enabled for these aliased pathes, this could allow for remote code execution. This issue is known to be exploited in the wild. This issue only affects Apache 2.4.49 and not earlier versions. The fix in Apache HTTP Server 2.4.50 was found to be incomplete, see CVE-2021-42013.
- [pwn3z/CVE-2021-41773-Apache-RCE](https://github.com/pwn3z/CVE-2021-41773-Apache-RCE)

---
## CVE-2021-40903 (2022-06-17T14:15:00)
> A vulnerability in Antminer Monitor 0.50.0 exists because of backdoor or misconfiguration inside a settings file in flask server. Settings file has a predefined secret string, which would be randomly generated, however it is static.
- [vulnz/CVE-2021-40903](https://github.com/vulnz/CVE-2021-40903)

---
## CVE-2021-40650 (2022-06-14T10:15:00)
> In Connx Version 6.2.0.1269 (20210623), a cookie can be issued by the application and not have the secure flag set.
- [l00neyhacker/CVE-2021-40650](https://github.com/l00neyhacker/CVE-2021-40650)

---
## CVE-2021-40649 (2022-06-14T10:15:00)
> In Connx Version 6.2.0.1269 (20210623), a cookie can be issued by the application and not have the HttpOnly flag set.
- [l00neyhacker/CVE-2021-40649](https://github.com/l00neyhacker/CVE-2021-40649)

---
## CVE-2021-40444 (2021-09-15T12:15:00)
> Microsoft MSHTML Remote Code Execution Vulnerability
- [klezVirus/CVE-2021-40444](https://github.com/klezVirus/CVE-2021-40444)

---
## CVE-2021-39409 (2022-06-24T17:15:00)
> A vulnerability exists in Online Student Rate System v1.0 that allows any user to register as an administrator without needing to be authenticated.
- [StefanDorresteijn/CVE-2021-39409](https://github.com/StefanDorresteijn/CVE-2021-39409)

---
## CVE-2021-39408 (2022-06-24T17:15:00)
> Cross Site Scripting (XSS) vulnerability exists in Online Student Rate System 1.0 via the page parameter on the index.php file
- [StefanDorresteijn/CVE-2021-39408](https://github.com/StefanDorresteijn/CVE-2021-39408)

---
## CVE-2021-38314 (2021-09-02T17:15:00)
> The Gutenberg Template Library & Redux Framework plugin <= 4.2.11 for WordPress registered several AJAX actions available to unauthenticated users in the `includes` function in `redux-core/class-redux-core.php` that were unique to a given site but deterministic and predictable given that they were based on an md5 hash of the site URL with a known salt value of '-redux' and an md5 hash of the previous hash with a known salt value of '-support'. These AJAX actions could be used to retrieve a list of active plugins and their versions, the site's PHP version, and an unsalted md5 hash of site’s `AUTH_KEY` concatenated with the `SECURE_AUTH_KEY`.
- [c0ff33b34n/CVE-2021-38314](https://github.com/c0ff33b34n/CVE-2021-38314)

---
## CVE-2021-3560 (2022-02-16T19:15:00)
> It was found that polkit could be tricked into bypassing the credential checks for D-Bus requests, elevating the privileges of the requestor to the root user. This flaw could be used by an unprivileged local attacker to, for example, create a new local administrator. The highest threat from this vulnerability is to data confidentiality and integrity as well as system availability.
- [UNICORDev/exploit-CVE-2021-3560](https://github.com/UNICORDev/exploit-CVE-2021-3560)

---
## CVE-2021-34527 (2021-07-02T22:15:00)
> Windows Print Spooler Remote Code Execution Vulnerability
- [cyb3rpeace/CVE-2021-34527](https://github.com/cyb3rpeace/CVE-2021-34527)

---
## CVE-2021-3156 (2021-01-26T21:15:00)
> Sudo before 1.9.5p2 contains an off-by-one error that can result in a heap-based buffer overflow, which allows privilege escalation to root via "sudoedit -s" and a command-line argument that ends with a single backslash character.
- [arvindshima/CVE-2021-3156](https://github.com/arvindshima/CVE-2021-3156)
- [arvindshima/CVE-2021-3156](https://github.com/arvindshima/CVE-2021-3156)

---
## CVE-2021-26855 (2021-03-03T00:15:00)
> Microsoft Exchange Server Remote Code Execution Vulnerability This CVE ID is unique from CVE-2021-26412, CVE-2021-26854, CVE-2021-26857, CVE-2021-26858, CVE-2021-27065, CVE-2021-27078.
- [raheel0x01/CVE-2021-26855](https://github.com/raheel0x01/CVE-2021-26855)
- [kh4sh3i/ProxyLogon](https://github.com/kh4sh3i/ProxyLogon)
- [TheDudeD6/ExchangeSmash](https://github.com/TheDudeD6/ExchangeSmash)
- [gobysec/GobyVuls](https://github.com/gobysec/GobyVuls)
- [1342486672/Flangvik](https://github.com/1342486672/Flangvik)
- [hakivvi/proxylogon](https://github.com/hakivvi/proxylogon)
- [zhzyker/vulmap](https://github.com/zhzyker/vulmap)
- [dwisiswant0/proxylogscan](https://github.com/dwisiswant0/proxylogscan)
- [thau0x01/poc_proxylogon](https://github.com/thau0x01/poc_proxylogon)
- [hosch3n/ProxyVulns](https://github.com/hosch3n/ProxyVulns)
- [yaoxiaoangry3/Flangvik](https://github.com/yaoxiaoangry3/Flangvik)
- [herwonowr/exprolog](https://github.com/herwonowr/exprolog)
- [cryptolakk/ProxyLogon-Mass-RCE](https://github.com/cryptolakk/ProxyLogon-Mass-RCE)
- [RickGeex/ProxyLogon](https://github.com/RickGeex/ProxyLogon)
- [SCS-Labs/HAFNIUM-Microsoft-Exchange-0day](https://github.com/SCS-Labs/HAFNIUM-Microsoft-Exchange-0day)
- [Nick-Yin12/106362522](https://github.com/Nick-Yin12/106362522)
- [Udyz/Proxylogon](https://github.com/Udyz/Proxylogon)
- [shacojx/CVE_2021_26855_SSRF](https://github.com/shacojx/CVE_2021_26855_SSRF)
- [shacojx/CVE-2021-26855-exploit-Exchange](https://github.com/shacojx/CVE-2021-26855-exploit-Exchange)
- [Flangvik/SharpProxyLogon](https://github.com/Flangvik/SharpProxyLogon)
- [mekhalleh/exchange_proxylogon](https://github.com/mekhalleh/exchange_proxylogon)
- [r0ckysec/CVE-2021-26855_Exchange](https://github.com/r0ckysec/CVE-2021-26855_Exchange)
- [praetorian-inc/proxylogon-exploit](https://github.com/praetorian-inc/proxylogon-exploit)
- [hictf/CVE-2021-26855-CVE-2021-27065](https://github.com/hictf/CVE-2021-26855-CVE-2021-27065)
- [Immersive-Labs-Sec/ProxyLogon](https://github.com/Immersive-Labs-Sec/ProxyLogon)
- [shacojx/Scan-Vuln-CVE-2021-26855](https://github.com/shacojx/Scan-Vuln-CVE-2021-26855)
- [DCScoder/Exchange_IOC_Hunter](https://github.com/DCScoder/Exchange_IOC_Hunter)
- [TaroballzChen/ProxyLogon-CVE-2021-26855-metasploit](https://github.com/TaroballzChen/ProxyLogon-CVE-2021-26855-metasploit)
- [p0wershe11/ProxyLogon](https://github.com/p0wershe11/ProxyLogon)
- [SotirisKar/CVE-2021-26855](https://github.com/SotirisKar/CVE-2021-26855)

---
## CVE-2021-25003 ()
> 
- [biulove0x/CVE-2021-25003](https://github.com/biulove0x/CVE-2021-25003)

---
## CVE-2021-22204 (2021-04-23T18:15:00)
> Improper neutralization of user data in the DjVu file format in ExifTool versions 7.44 and up allows arbitrary code execution when parsing the malicious image
- [GatoGamer1155/CVE-2021-22204](https://github.com/GatoGamer1155/CVE-2021-22204)
- [UNICORDev/exploit-CVE-2021-22204](https://github.com/UNICORDev/exploit-CVE-2021-22204)
- [ChispasHK/MiruCTF](https://github.com/ChispasHK/MiruCTF)
- [tuhin81/CVE-2021-22204-exiftool](https://github.com/tuhin81/CVE-2021-22204-exiftool)
- [0xBruno/CVE-2021-22204](https://github.com/0xBruno/CVE-2021-22204)
- [trganda/CVE-2021-22204](https://github.com/trganda/CVE-2021-22204)
- [Asaad27/CVE-2021-22204-RSE](https://github.com/Asaad27/CVE-2021-22204-RSE)
- [ph-arm/CVE-2021-22204-Gitlab](https://github.com/ph-arm/CVE-2021-22204-Gitlab)
- [AssassinUKG/CVE-2021-22204](https://github.com/AssassinUKG/CVE-2021-22204)
- [Konstantinos-Papanagnou/CMSpit](https://github.com/Konstantinos-Papanagnou/CMSpit)
- [PenTestical/CVE-2021-22204](https://github.com/PenTestical/CVE-2021-22204)
- [bilkoh/POC-CVE-2021-22204](https://github.com/bilkoh/POC-CVE-2021-22204)
- [convisolabs/CVE-2021-22204-exiftool](https://github.com/convisolabs/CVE-2021-22204-exiftool)
- [se162xg/CVE-2021-22204](https://github.com/se162xg/CVE-2021-22204)

---
## CVE-2021-22005 (2021-09-23T12:15:00)
> The vCenter Server contains an arbitrary file upload vulnerability in the Analytics service. A malicious actor with network access to port 443 on vCenter Server may exploit this issue to execute code on vCenter Server by uploading a specially crafted file.
- [InventorMAO/cve-2021-22005](https://github.com/InventorMAO/cve-2021-22005)
