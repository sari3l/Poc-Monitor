# 2016 List

---
## CVE-2016-9843 (2017-05-23T04:29:00)
> The crc32_big function in crc32.c in zlib 1.2.8 might allow context-dependent attackers to have unspecified impact via vectors involving big-endian CRC calculation.
- [Live-Hack-CVE/CVE-2016-9843](https://github.com/Live-Hack-CVE/CVE-2016-9843)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9843">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9843">

---
## CVE-2016-9842 (2017-05-23T04:29:00)
> The inflateMark function in inflate.c in zlib 1.2.8 might allow context-dependent attackers to have unspecified impact via vectors involving left shifts of negative integers.
- [Live-Hack-CVE/CVE-2016-9842](https://github.com/Live-Hack-CVE/CVE-2016-9842)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9842">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9842">

---
## CVE-2016-9841 (2017-05-23T04:29:00)
> inffast.c in zlib 1.2.8 might allow context-dependent attackers to have unspecified impact by leveraging improper pointer arithmetic.
- [Live-Hack-CVE/CVE-2016-9841](https://github.com/Live-Hack-CVE/CVE-2016-9841)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9841">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9841">

---
## CVE-2016-9840 (2017-05-23T04:29:00)
> inftrees.c in zlib 1.2.8 might allow context-dependent attackers to have unspecified impact by leveraging improper pointer arithmetic.
- [Live-Hack-CVE/CVE-2016-9840](https://github.com/Live-Hack-CVE/CVE-2016-9840)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9840">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9840">

---
## CVE-2016-9299 (2017-01-12T23:59:00)
> The remoting module in Jenkins before 2.32 and LTS before 2.19.3 allows remote attackers to execute arbitrary code via a crafted serialized Java object, which triggers an LDAP query to a third-party server.
- [r00t4dm/Jenkins-CVE-2016-9299](https://github.com/r00t4dm/Jenkins-CVE-2016-9299)	<img alt="forks" src="https://img.shields.io/github/forks/r00t4dm/Jenkins-CVE-2016-9299">	<img alt="stars" src="https://img.shields.io/github/stars/r00t4dm/Jenkins-CVE-2016-9299">

---
## CVE-2016-9192 (2016-12-14T00:59:00)
> A vulnerability in Cisco AnyConnect Secure Mobility Client for Windows could allow an authenticated, local attacker to install and execute an arbitrary executable file with privileges equivalent to the Microsoft Windows operating system SYSTEM account. More Information: CSCvb68043. Known Affected Releases: 4.3(2039) 4.3(748). Known Fixed Releases: 4.3(4019) 4.4(225).
- [serializingme/cve-2016-9192](https://github.com/serializingme/cve-2016-9192)	<img alt="forks" src="https://img.shields.io/github/forks/serializingme/cve-2016-9192">	<img alt="stars" src="https://img.shields.io/github/stars/serializingme/cve-2016-9192">

---
## CVE-2016-9043 (2018-04-24T19:29:00)
> An out of bound write vulnerability exists in the EMF parsing functionality of CorelDRAW X8 (CdrGfx - Corel Graphics Engine (64-Bit) - 18.1.0.661). A specially crafted EMF file can cause a vulnerability resulting in potential code execution. An attacker can send the victim a specific EMF file to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-9043](https://github.com/Live-Hack-CVE/CVE-2016-9043)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9043">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9043">

---
## CVE-2016-9040 (2018-09-07T12:29:00)
> An exploitable denial of service exists in the the Joyent SmartOS OS 20161110T013148Z Hyprlofs file system. The vulnerability is present in the Ioctl system call with the command HYPRLOFSADDENTRIES when used with a 32 bit model. An attacker can cause a buffer to be allocated and never freed. When repeatedly exploit this will result in memory exhaustion, resulting in a full system denial of service.
- [Live-Hack-CVE/CVE-2016-9040](https://github.com/Live-Hack-CVE/CVE-2016-9040)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9040">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9040">

---
## CVE-2016-9038 (2018-04-24T19:29:00)
> An exploitable double fetch vulnerability exists in the SboxDrv.sys driver functionality of Invincea-X 6.1.3-24058. A specially crafted input buffer and race condition can result in kernel memory corruption, which could result in privilege escalation. An attacker needs to execute a special application locally to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-9038](https://github.com/Live-Hack-CVE/CVE-2016-9038)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9038">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9038">

---
## CVE-2016-9037 (2016-12-23T22:59:00)
> An exploitable out-of-bounds array access vulnerability exists in the xrow_header_decode function of Tarantool 1.7.2.0-g8e92715. A specially crafted packet can cause the function to access an element outside the bounds of a global array that is used to determine the type of the specified key's value. This can lead to an out of bounds read within the context of the server. An attacker who exploits this vulnerability can cause a denial of service vulnerability on the server.
- [Live-Hack-CVE/CVE-2016-9037](https://github.com/Live-Hack-CVE/CVE-2016-9037)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9037">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9037">

---
## CVE-2016-9036 (2016-12-23T22:59:00)
> An exploitable incorrect return value vulnerability exists in the mp_check function of Tarantool's Msgpuck library 1.0.3. A specially crafted packet can cause the mp_check function to incorrectly return success when trying to check if decoding a map16 packet will read outside the bounds of a buffer, resulting in a denial of service vulnerability.
- [Live-Hack-CVE/CVE-2016-9036](https://github.com/Live-Hack-CVE/CVE-2016-9036)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9036">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9036">

---
## CVE-2016-9035 (2016-12-14T17:59:00)
> An exploitable buffer overflow exists in the Joyent SmartOS 20161110T013148Z Hyprlofs file system. The vulnerability is present in the Ioctl system call with the command HYPRLOFS_ADD_ENTRIES when dealing with native file systems. An attacker can craft an input that can cause a buffer overflow in the path variable leading to an out of bounds memory access and could result in potential privilege escalation. This vulnerability is distinct from CVE-2016-9033.
- [Live-Hack-CVE/CVE-2016-9035](https://github.com/Live-Hack-CVE/CVE-2016-9035)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9035">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9035">

---
## CVE-2016-9034 (2016-12-14T17:59:00)
> An exploitable buffer overflow exists in the Joyent SmartOS 20161110T013148Z Hyprlofs file system. The vulnerability is present in the Ioctl system call with the command HYPRLOFS_ADD_ENTRIES when dealing with 32-bit file systems. An attacker can craft an input that can cause a buffer overflow in the nm variable leading to an out of bounds memory access and could result in potential privilege escalation. This vulnerability is distinct from CVE-2016-9032.
- [Live-Hack-CVE/CVE-2016-9034](https://github.com/Live-Hack-CVE/CVE-2016-9034)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9034">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9034">

---
## CVE-2016-9033 (2016-12-14T17:59:00)
> An exploitable buffer overflow exists in the Joyent SmartOS 20161110T013148Z Hyprlofs file system. The vulnerability is present in the Ioctl system call with the command HYPRLOFS_ADD_ENTRIES when dealing with native file systems. An attacker can craft an input that can cause a buffer overflow in the path variable leading to an out of bounds memory access and could result in potential privilege escalation. This vulnerability is distinct from CVE-2016-9035.
- [Live-Hack-CVE/CVE-2016-9033](https://github.com/Live-Hack-CVE/CVE-2016-9033)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9033">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9033">

---
## CVE-2016-9032 (2016-12-14T17:59:00)
> An exploitable buffer overflow exists in the Joyent SmartOS 20161110T013148Z Hyprlofs file system. The vulnerability is present in the Ioctl system call with the command HYPRLOFS_ADD_ENTRIES when dealing with native file systems. An attacker can craft an input that can cause a buffer overflow in the nm variable leading to an out of bounds memory access and could result in potential privilege escalation. This vulnerability is distinct from CVE-2016-9034.
- [Live-Hack-CVE/CVE-2016-9032](https://github.com/Live-Hack-CVE/CVE-2016-9032)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9032">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9032">

---
## CVE-2016-9031 (2016-12-14T17:59:00)
> An exploitable integer overflow exists in the Joyent SmartOS 20161110T013148Z Hyprlofs file system. The vulnerability is present in the Ioctl system call with the command HYPRLOFS_ADD_ENTRIES when dealing with 32-bit file systems. An attacker can craft an input that can cause a kernel panic and potentially be leveraged into a full privilege escalation vulnerability. This vulnerability is distinct from CVE-2016-8733.
- [Live-Hack-CVE/CVE-2016-9031](https://github.com/Live-Hack-CVE/CVE-2016-9031)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-9031">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-9031">

---
## CVE-2016-8732 (2018-04-24T19:29:00)
> Multiple security flaws exists in InvProtectDrv.sys which is a part of Invincea Dell Protected Workspace 5.1.1-22303. Weak restrictions on the driver communication channel and additional insufficient checks allow any application to turn off some of the protection mechanisms provided by the Invincea product.
- [Live-Hack-CVE/CVE-2016-8732](https://github.com/Live-Hack-CVE/CVE-2016-8732)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8732">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8732">
- [Live-Hack-CVE/CVE-2016-8732](https://github.com/Live-Hack-CVE/CVE-2016-8732)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8732">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8732">

---
## CVE-2016-8731 (2017-06-21T19:29:00)
> Hard-coded FTP credentials (r:r) are included in the Foscam C1 running firmware 1.9.1.12. Knowledge of these credentials would allow remote access to any cameras found on the internet that do not have port 50021 blocked by an intermediate device.
- [Live-Hack-CVE/CVE-2016-8731](https://github.com/Live-Hack-CVE/CVE-2016-8731)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8731">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8731">

---
## CVE-2016-8730 (2018-04-24T19:29:00)
> An of bound write / memory corruption vulnerability exists in the GIF parsing functionality of Core PHOTO-PAINT X8 18.1.0.661. A specially crafted GIF file can cause a vulnerability resulting in potential memory corruption resulting in code execution. An attacker can send the victim a specific GIF file to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8730](https://github.com/Live-Hack-CVE/CVE-2016-8730)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8730">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8730">

---
## CVE-2016-8728 (2018-04-24T19:29:00)
> An exploitable heap out of bounds write vulnerability exists in the Fitz graphical library part of the MuPDF renderer. A specially crafted PDF file can cause a out of bounds write resulting in heap metadata and sensitive process memory corruption leading to potential code execution. Victim needs to open the specially crafted file in a vulnerable reader in order to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8728](https://github.com/Live-Hack-CVE/CVE-2016-8728)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8728">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8728">

---
## CVE-2016-8726 (2017-04-13T19:59:00)
> An exploitable null pointer dereference vulnerability exists in the Web Application /forms/web_runScript iw_filename functionality of Moxa AWK-3131A Wireless Access Point running firmware 1.1. An HTTP POST request with a blank line in the header will cause a segmentation fault in the web server.
- [Live-Hack-CVE/CVE-2016-8726](https://github.com/Live-Hack-CVE/CVE-2016-8726)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8726">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8726">

---
## CVE-2016-8725 (2017-04-13T19:59:00)
> An exploitable information disclosure vulnerability exists in the Web Application functionality of the Moxa AWK-3131A wireless access point running firmware 1.1. Retrieving a specific URL without authentication can reveal sensitive information to an attacker.
- [Live-Hack-CVE/CVE-2016-8725](https://github.com/Live-Hack-CVE/CVE-2016-8725)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8725">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8725">

---
## CVE-2016-8724 (2017-04-13T19:59:00)
> An exploitable information disclosure vulnerability exists in the serviceAgent functionality of Moxa AWK-3131A Wireless Access Point running firmware 1.1. A specially crafted TCP query will allow an attacker to retrieve potentially sensitive information.
- [Live-Hack-CVE/CVE-2016-8724](https://github.com/Live-Hack-CVE/CVE-2016-8724)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8724">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8724">

---
## CVE-2016-8723 (2017-04-13T19:59:00)
> An exploitable null pointer dereference exists in the Web Application functionality of Moxa AWK-3131A Wireless Access Point running firmware 1.1. Any HTTP GET request not preceded by an '/' will cause a segmentation fault in the web server. An attacker can send any of a multitude of potentially unexpected HTTP get requests to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8723](https://github.com/Live-Hack-CVE/CVE-2016-8723)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8723">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8723">

---
## CVE-2016-8722 (2017-04-13T19:59:00)
> An exploitable Information Disclosure vulnerability exists in the Web Application functionality of Moxa AWK-3131A Series Industrial IEEE 802.11a/b/g/n wireless AP/bridge/client. Retrieving a specific URL without authentication can reveal sensitive information to an attacker.
- [Live-Hack-CVE/CVE-2016-8722](https://github.com/Live-Hack-CVE/CVE-2016-8722)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8722">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8722">

---
## CVE-2016-8721 (2017-04-20T18:59:00)
> An exploitable OS Command Injection vulnerability exists in the web application 'ping' functionality of Moxa AWK-3131A Wireless Access Points running firmware 1.1. Specially crafted web form input can cause an OS Command Injection resulting in complete compromise of the vulnerable device. An attacker can exploit this vulnerability remotely.
- [Live-Hack-CVE/CVE-2016-8721](https://github.com/Live-Hack-CVE/CVE-2016-8721)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8721">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8721">

---
## CVE-2016-8720 (2017-04-13T19:59:00)
> An exploitable HTTP Header Injection vulnerability exists in the Web Application functionality of the Moxa AWK-3131A Wireless Access Point running firmware 1.1. A specially crafted HTTP request can inject a payload in the bkpath parameter which will be copied in to Location header of the HTTP response.
- [Live-Hack-CVE/CVE-2016-8720](https://github.com/Live-Hack-CVE/CVE-2016-8720)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8720">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8720">

---
## CVE-2016-8719 (2017-04-12T19:59:00)
> An exploitable reflected Cross-Site Scripting vulnerability exists in the Web Application functionality of Moxa AWK-3131A Wireless Access Point running firmware 1.1. Specially crafted input, in multiple parameters, can cause a malicious scripts to be executed by a victim.
- [Live-Hack-CVE/CVE-2016-8719](https://github.com/Live-Hack-CVE/CVE-2016-8719)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8719">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8719">

---
## CVE-2016-8718 (2017-04-12T19:59:00)
> An exploitable Cross-Site Request Forgery vulnerability exists in the Web Application functionality of Moxa AWK-3131A Wireless Access Point running firmware 1.1. A specially crafted form can trick a client into making an unintentional request to the web server which will be treated as an authentic request.
- [Live-Hack-CVE/CVE-2016-8718](https://github.com/Live-Hack-CVE/CVE-2016-8718)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8718">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8718">

---
## CVE-2016-8717 (2018-04-02T17:29:00)
> An exploitable Use of Hard-coded Credentials vulnerability exists in the Moxa AWK-3131A Wireless Access Point running firmware 1.1. The device operating system contains an undocumented, privileged (root) account with hard-coded credentials, giving attackers full control of affected devices.
- [Live-Hack-CVE/CVE-2016-8717](https://github.com/Live-Hack-CVE/CVE-2016-8717)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8717">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8717">

---
## CVE-2016-8716 (2017-04-12T19:59:00)
> An exploitable Cleartext Transmission of Password vulnerability exists in the Web Application functionality of Moxa AWK-3131A Wireless Access Point running firmware 1.1. The Change Password functionality of the Web Application transmits the password in cleartext. An attacker capable of intercepting this traffic is able to obtain valid credentials.
- [Live-Hack-CVE/CVE-2016-8716](https://github.com/Live-Hack-CVE/CVE-2016-8716)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8716">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8716">

---
## CVE-2016-8715 (2017-02-28T15:59:00)
> An exploitable heap corruption vulnerability exists in the loadTrailer functionality of Iceni Argus version 6.6.05. A specially crafted PDF file can cause a heap corruption resulting in arbitrary code execution. An attacker can send/provide a malicious PDF file to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8715](https://github.com/Live-Hack-CVE/CVE-2016-8715)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8715">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8715">

---
## CVE-2016-8714 (2017-03-10T10:59:00)
> An exploitable buffer overflow vulnerability exists in the LoadEncoding functionality of the R programming language version 3.3.0. A specially crafted R script can cause a buffer overflow resulting in a memory corruption. An attacker can send a malicious R script to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8714](https://github.com/Live-Hack-CVE/CVE-2016-8714)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8714">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8714">

---
## CVE-2016-8713 (2017-02-10T17:59:00)
> A remote out of bound write / memory corruption vulnerability exists in the PDF parsing functionality of Nitro Pro 10.5.9.9. A specially crafted PDF file can cause a vulnerability resulting in potential memory corruption. An attacker can send the victim a specific PDF file to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8713](https://github.com/Live-Hack-CVE/CVE-2016-8713)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8713">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8713">
- [Live-Hack-CVE/CVE-2016-8713](https://github.com/Live-Hack-CVE/CVE-2016-8713)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8713">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8713">

---
## CVE-2016-8712 (2017-04-13T19:59:00)
> An exploitable nonce reuse vulnerability exists in the Web Application functionality of Moxa AWK-3131A Wireless AP running firmware 1.1. The device uses one nonce for all session authentication requests and only changes the nonce if the web application has been idle for 300 seconds.
- [Live-Hack-CVE/CVE-2016-8712](https://github.com/Live-Hack-CVE/CVE-2016-8712)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8712">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8712">

---
## CVE-2016-8711 (2017-02-10T17:59:00)
> A potential remote code execution vulnerability exists in the PDF parsing functionality of Nitro Pro 10. A specially crafted PDF file can cause a vulnerability resulting in potential code execution. An attacker can send the victim a specific PDF file to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8711](https://github.com/Live-Hack-CVE/CVE-2016-8711)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8711">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8711">

---
## CVE-2016-8710 (2017-01-26T21:59:00)
> An exploitable heap write out of bounds vulnerability exists in the decoding of BPG images in Libbpg library. A crafted BPG image decoded by libbpg can cause an integer underflow vulnerability causing an out of bounds heap write leading to remote code execution. This vulnerability can be triggered via attempting to decode a crafted BPG image using Libbpg.
- [Live-Hack-CVE/CVE-2016-8710](https://github.com/Live-Hack-CVE/CVE-2016-8710)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8710">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8710">

---
## CVE-2016-8709 (2017-02-10T17:59:00)
> A remote out of bound write / memory corruption vulnerability exists in the PDF parsing functionality of Nitro Pro 10. A specially crafted PDF file can cause a vulnerability resulting in potential memory corruption. An attacker can send the victim a specific PDF file to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8709](https://github.com/Live-Hack-CVE/CVE-2016-8709)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8709">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8709">

---
## CVE-2016-8707 (2016-12-23T22:59:00)
> An exploitable out of bounds write exists in the handling of compressed TIFF images in ImageMagicks's convert utility. A crafted TIFF document can lead to an out of bounds write which in particular circumstances could be leveraged into remote code execution. The vulnerability can be triggered through any user controlled TIFF that is handled by this functionality.
- [Live-Hack-CVE/CVE-2016-8707](https://github.com/Live-Hack-CVE/CVE-2016-8707)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8707">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8707">

---
## CVE-2016-8390 (2018-06-04T19:29:00)
> An exploitable out of bounds write vulnerability exists in the parsing of ELF Section Headers of Hopper Disassembler 3.11.20. A specially crafted ELF file can cause attacker controlled pointer arithmetic resulting in a partially controlled out of bounds write. An attacker can craft an ELF file with specific section headers to trigger this vulnerability.
- [Live-Hack-CVE/CVE-2016-8390](https://github.com/Live-Hack-CVE/CVE-2016-8390)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8390">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8390">

---
## CVE-2016-8389 (2017-02-28T15:59:00)
> An exploitable integer-overflow vulnerability exists within Iceni Argus. When it attempts to convert a malformed PDF to XML, it will attempt to convert each character from a font into a polygon and then attempt to rasterize these shapes. As the application attempts to iterate through the rows and initializing the polygon shape in the buffer, it will write outside of the bounds of said buffer. This can lead to code execution under the context of the account running it.
- [Live-Hack-CVE/CVE-2016-8389](https://github.com/Live-Hack-CVE/CVE-2016-8389)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8389">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8389">

---
## CVE-2016-8388 (2017-02-28T15:59:00)
> An exploitable arbitrary heap-overwrite vulnerability exists within Iceni Argus. When it attempts to convert a malformed PDF to XML, it will explicitly trust an index within the specific font object and use it to write the font's name to a single object within an array of objects.
- [Live-Hack-CVE/CVE-2016-8388](https://github.com/Live-Hack-CVE/CVE-2016-8388)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-8388">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-8388">

---
## CVE-2016-7892 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable use after free vulnerability in the TextField class. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7892](https://github.com/Live-Hack-CVE/CVE-2016-7892)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7892">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7892">
- [Live-Hack-CVE/CVE-2016-7892](https://github.com/Live-Hack-CVE/CVE-2016-7892)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7892">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7892">

---
## CVE-2016-7890 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have security bypass vulnerability in the implementation of the same origin policy.
- [Live-Hack-CVE/CVE-2016-7890](https://github.com/Live-Hack-CVE/CVE-2016-7890)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7890">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7890">
- [Live-Hack-CVE/CVE-2016-7890](https://github.com/Live-Hack-CVE/CVE-2016-7890)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7890">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7890">

---
## CVE-2016-7881 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable use after free vulnerability in the MovieClip class when handling conversion to an object. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7881](https://github.com/Live-Hack-CVE/CVE-2016-7881)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7881">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7881">
- [Live-Hack-CVE/CVE-2016-7881](https://github.com/Live-Hack-CVE/CVE-2016-7881)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7881">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7881">

---
## CVE-2016-7880 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable use after free vulnerability when setting the length property of an array object. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7880](https://github.com/Live-Hack-CVE/CVE-2016-7880)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7880">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7880">
- [Live-Hack-CVE/CVE-2016-7880](https://github.com/Live-Hack-CVE/CVE-2016-7880)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7880">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7880">

---
## CVE-2016-7879 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable use after free vulnerability in the NetConnection class when handling an attached script object. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7879](https://github.com/Live-Hack-CVE/CVE-2016-7879)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7879">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7879">
- [Live-Hack-CVE/CVE-2016-7879](https://github.com/Live-Hack-CVE/CVE-2016-7879)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7879">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7879">

---
## CVE-2016-7878 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable use after free vulnerability in the PSDK's MediaPlayer class. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7878](https://github.com/Live-Hack-CVE/CVE-2016-7878)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7878">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7878">
- [Live-Hack-CVE/CVE-2016-7878](https://github.com/Live-Hack-CVE/CVE-2016-7878)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7878">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7878">

---
## CVE-2016-7877 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable use after free vulnerability in the Action Message Format serialization (AFM0). Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7877](https://github.com/Live-Hack-CVE/CVE-2016-7877)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7877">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7877">

---
## CVE-2016-7876 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable memory corruption vulnerability in the Clipboard class related to data handling functionality. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7876](https://github.com/Live-Hack-CVE/CVE-2016-7876)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7876">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7876">

---
## CVE-2016-7875 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable integer overflow vulnerability in the BitmapData class. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7875](https://github.com/Live-Hack-CVE/CVE-2016-7875)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7875">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7875">

---
## CVE-2016-7874 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable memory corruption vulnerability in the NetConnection class when handling the proxy types. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7874](https://github.com/Live-Hack-CVE/CVE-2016-7874)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7874">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7874">

---
## CVE-2016-7873 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable memory corruption vulnerability in the PSDK class related to ad policy functionality method. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7873](https://github.com/Live-Hack-CVE/CVE-2016-7873)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7873">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7873">
- [Live-Hack-CVE/CVE-2016-7873](https://github.com/Live-Hack-CVE/CVE-2016-7873)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7873">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7873">

---
## CVE-2016-7872 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable use after free vulnerability in the MovieClip class related to objects at multiple presentation levels. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7872](https://github.com/Live-Hack-CVE/CVE-2016-7872)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7872">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7872">
- [Live-Hack-CVE/CVE-2016-7872](https://github.com/Live-Hack-CVE/CVE-2016-7872)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7872">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7872">

---
## CVE-2016-7871 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable memory corruption vulnerability in the Worker class. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7871](https://github.com/Live-Hack-CVE/CVE-2016-7871)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7871">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7871">
- [Live-Hack-CVE/CVE-2016-7871](https://github.com/Live-Hack-CVE/CVE-2016-7871)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7871">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7871">

---
## CVE-2016-7870 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable buffer overflow / underflow vulnerability in the RegExp class for specific search strategies. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7870](https://github.com/Live-Hack-CVE/CVE-2016-7870)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7870">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7870">
- [Live-Hack-CVE/CVE-2016-7870](https://github.com/Live-Hack-CVE/CVE-2016-7870)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7870">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7870">

---
## CVE-2016-7869 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable buffer overflow / underflow vulnerability in the RegExp class related to backtrack search functionality. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7869](https://github.com/Live-Hack-CVE/CVE-2016-7869)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7869">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7869">
- [Live-Hack-CVE/CVE-2016-7869](https://github.com/Live-Hack-CVE/CVE-2016-7869)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7869">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7869">

---
## CVE-2016-7868 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable buffer overflow / underflow vulnerability in the RegExp class related to alternation functionality. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7868](https://github.com/Live-Hack-CVE/CVE-2016-7868)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7868">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7868">
- [Live-Hack-CVE/CVE-2016-7868](https://github.com/Live-Hack-CVE/CVE-2016-7868)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7868">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7868">

---
## CVE-2016-7867 (2016-12-15T06:59:00)
> Adobe Flash Player versions 23.0.0.207 and earlier, 11.2.202.644 and earlier have an exploitable buffer overflow / underflow vulnerability in the RegExp class related to bookmarking in searches. Successful exploitation could lead to arbitrary code execution.
- [Live-Hack-CVE/CVE-2016-7867](https://github.com/Live-Hack-CVE/CVE-2016-7867)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7867">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7867">
- [Live-Hack-CVE/CVE-2016-7867](https://github.com/Live-Hack-CVE/CVE-2016-7867)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7867">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7867">

---
## CVE-2016-7440 (2016-12-13T16:59:00)
> The C software implementation of AES Encryption and Decryption in wolfSSL (formerly CyaSSL) before 3.9.10 makes it easier for local users to discover AES keys by leveraging cache-bank timing differences.
- [Live-Hack-CVE/CVE-2016-7440](https://github.com/Live-Hack-CVE/CVE-2016-7440)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7440">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7440">

---
## CVE-2016-7052 (2016-09-26T19:59:00)
> crypto/x509/x509_vfy.c in OpenSSL 1.0.2i allows remote attackers to cause a denial of service (NULL pointer dereference and application crash) by triggering a CRL operation.
- [Live-Hack-CVE/CVE-2016-7052](https://github.com/Live-Hack-CVE/CVE-2016-7052)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-7052">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-7052">

---
## CVE-2016-6989 (2016-10-13T20:00:00)
> Adobe Flash Player before 18.0.0.382 and 19.x through 23.x before 23.0.0.185 on Windows and OS X and before 11.2.202.637 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4273, CVE-2016-6982, CVE-2016-6983, CVE-2016-6984, CVE-2016-6985, CVE-2016-6986, and CVE-2016-6990.
- [Live-Hack-CVE/CVE-2016-6989](https://github.com/Live-Hack-CVE/CVE-2016-6989)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6989">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6989">

---
## CVE-2016-6983 (2016-10-13T19:59:00)
> Adobe Flash Player before 18.0.0.382 and 19.x through 23.x before 23.0.0.185 on Windows and OS X and before 11.2.202.637 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4273, CVE-2016-6982, CVE-2016-6984, CVE-2016-6985, CVE-2016-6986, CVE-2016-6989, and CVE-2016-6990.
- [Live-Hack-CVE/CVE-2016-6983](https://github.com/Live-Hack-CVE/CVE-2016-6983)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6983">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6983">
- [Live-Hack-CVE/CVE-2016-6983](https://github.com/Live-Hack-CVE/CVE-2016-6983)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6983">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6983">

---
## CVE-2016-6982 (2016-10-13T19:59:00)
> Adobe Flash Player before 18.0.0.382 and 19.x through 23.x before 23.0.0.185 on Windows and OS X and before 11.2.202.637 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4273, CVE-2016-6983, CVE-2016-6984, CVE-2016-6985, CVE-2016-6986, CVE-2016-6989, and CVE-2016-6990.
- [Live-Hack-CVE/CVE-2016-6982](https://github.com/Live-Hack-CVE/CVE-2016-6982)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6982">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6982">
- [Live-Hack-CVE/CVE-2016-6982](https://github.com/Live-Hack-CVE/CVE-2016-6982)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6982">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6982">

---
## CVE-2016-6981 (2016-10-13T19:59:00)
> Use-after-free vulnerability in Adobe Flash Player before 18.0.0.382 and 19.x through 23.x before 23.0.0.185 on Windows and OS X and before 11.2.202.637 on Linux allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-6987.
- [Live-Hack-CVE/CVE-2016-6981](https://github.com/Live-Hack-CVE/CVE-2016-6981)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6981">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6981">
- [Live-Hack-CVE/CVE-2016-6981](https://github.com/Live-Hack-CVE/CVE-2016-6981)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6981">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6981">

---
## CVE-2016-6924 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, and CVE-2016-6922.
- [Live-Hack-CVE/CVE-2016-6924](https://github.com/Live-Hack-CVE/CVE-2016-6924)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6924">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6924">
- [Live-Hack-CVE/CVE-2016-6924](https://github.com/Live-Hack-CVE/CVE-2016-6924)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6924">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6924">

---
## CVE-2016-6922 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-6922](https://github.com/Live-Hack-CVE/CVE-2016-6922)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6922">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6922">

---
## CVE-2016-6306 (2016-09-26T19:59:00)
> The certificate parser in OpenSSL before 1.0.1u and 1.0.2 before 1.0.2i might allow remote attackers to cause a denial of service (out-of-bounds read) via crafted certificate operations, related to s3_clnt.c and s3_srvr.c.
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-6306](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6306)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6306">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6306">

---
## CVE-2016-6304 (2016-09-26T19:59:00)
> Multiple memory leaks in t1_lib.c in OpenSSL before 1.0.1u, 1.0.2 before 1.0.2i, and 1.1.0 before 1.1.0a allow remote attackers to cause a denial of service (memory consumption) via large OCSP Status Request extensions.
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-6304](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6304)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6304">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6304">

---
## CVE-2016-6302 (2016-09-16T05:59:00)
> The tls_decrypt_ticket function in ssl/t1_lib.c in OpenSSL before 1.1.0 does not consider the HMAC size during validation of the ticket length, which allows remote attackers to cause a denial of service via a ticket that is too short.
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-6302](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6302)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6302">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-6302">

---
## CVE-2016-6210 (2017-02-13T17:59:00)
> sshd in OpenSSH before 7.3, when SHA256 or SHA512 are used for user password hashing, uses BLOWFISH hashing on a static password when the username does not exist, which allows remote attackers to enumerate users by leveraging the timing difference between responses when a large password is provided.
- [goomdan/CVE-2016-6210-exploit](https://github.com/goomdan/CVE-2016-6210-exploit)	<img alt="forks" src="https://img.shields.io/github/forks/goomdan/CVE-2016-6210-exploit">	<img alt="stars" src="https://img.shields.io/github/stars/goomdan/CVE-2016-6210-exploit">
- [hackingyseguridad/ssha](https://github.com/hackingyseguridad/ssha)	<img alt="forks" src="https://img.shields.io/github/forks/hackingyseguridad/ssha">	<img alt="stars" src="https://img.shields.io/github/stars/hackingyseguridad/ssha">
- [Tardcircus/CVE2016-6210](https://github.com/Tardcircus/CVE2016-6210)	<img alt="forks" src="https://img.shields.io/github/forks/Tardcircus/CVE2016-6210">	<img alt="stars" src="https://img.shields.io/github/stars/Tardcircus/CVE2016-6210">
- [calebshortt/opensshd_user_enumeration](https://github.com/calebshortt/opensshd_user_enumeration)	<img alt="forks" src="https://img.shields.io/github/forks/calebshortt/opensshd_user_enumeration">	<img alt="stars" src="https://img.shields.io/github/stars/calebshortt/opensshd_user_enumeration">
- [justlce/CVE-2016-6210-Exploit](https://github.com/justlce/CVE-2016-6210-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/justlce/CVE-2016-6210-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/justlce/CVE-2016-6210-Exploit">
- [bassitone/OpenSSH-User-Enumeration](https://github.com/bassitone/OpenSSH-User-Enumeration)	<img alt="forks" src="https://img.shields.io/github/forks/bassitone/OpenSSH-User-Enumeration">	<img alt="stars" src="https://img.shields.io/github/stars/bassitone/OpenSSH-User-Enumeration">
- [samh4cks/CVE-2016-6210-OpenSSH-User-Enumeration](https://github.com/samh4cks/CVE-2016-6210-OpenSSH-User-Enumeration)	<img alt="forks" src="https://img.shields.io/github/forks/samh4cks/CVE-2016-6210-OpenSSH-User-Enumeration">	<img alt="stars" src="https://img.shields.io/github/stars/samh4cks/CVE-2016-6210-OpenSSH-User-Enumeration">

---
## CVE-2016-6207 (2016-08-12T15:59:00)
> Integer overflow in the _gdContributionsAlloc function in gd_interpolation.c in GD Graphics Library (aka libgd) before 2.2.3 allows remote attackers to cause a denial of service (out-of-bounds memory write or memory consumption) via unspecified vectors.
- [Live-Hack-CVE/CVE-2016-6207](https://github.com/Live-Hack-CVE/CVE-2016-6207)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-6207">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-6207">

---
## CVE-2016-6187 (2016-08-06T20:59:00)
> The apparmor_setprocattr function in security/apparmor/lsm.c in the Linux kernel before 4.6.5 does not validate the buffer size, which allows local users to gain privileges by triggering an AppArmor setprocattr hook.
- [Milo-D/CVE-2016-6187_LPE](https://github.com/Milo-D/CVE-2016-6187_LPE)	<img alt="forks" src="https://img.shields.io/github/forks/Milo-D/CVE-2016-6187_LPE">	<img alt="stars" src="https://img.shields.io/github/stars/Milo-D/CVE-2016-6187_LPE">
- [vnik5287/cve-2016-6187-poc](https://github.com/vnik5287/cve-2016-6187-poc)	<img alt="forks" src="https://img.shields.io/github/forks/vnik5287/cve-2016-6187-poc">	<img alt="stars" src="https://img.shields.io/github/stars/vnik5287/cve-2016-6187-poc">

---
## CVE-2016-5734 (2016-07-03T01:59:00)
> phpMyAdmin 4.0.x before 4.0.10.16, 4.4.x before 4.4.15.7, and 4.6.x before 4.6.3 does not properly choose delimiters to prevent use of the preg_replace e (aka eval) modifier, which might allow remote attackers to execute arbitrary PHP code via a crafted string, as demonstrated by the table search-and-replace implementation.
- [miko550/CVE-2016-5734-docker](https://github.com/miko550/CVE-2016-5734-docker)	<img alt="forks" src="https://img.shields.io/github/forks/miko550/CVE-2016-5734-docker">	<img alt="stars" src="https://img.shields.io/github/stars/miko550/CVE-2016-5734-docker">
- [HKirito/phpmyadmin4.4_cve-2016-5734](https://github.com/HKirito/phpmyadmin4.4_cve-2016-5734)	<img alt="forks" src="https://img.shields.io/github/forks/HKirito/phpmyadmin4.4_cve-2016-5734">	<img alt="stars" src="https://img.shields.io/github/stars/HKirito/phpmyadmin4.4_cve-2016-5734">
- [msharm33/CVE2016-5734](https://github.com/msharm33/CVE2016-5734)	<img alt="forks" src="https://img.shields.io/github/forks/msharm33/CVE2016-5734">	<img alt="stars" src="https://img.shields.io/github/stars/msharm33/CVE2016-5734">
- [KosukeShimofuji/CVE-2016-5734](https://github.com/KosukeShimofuji/CVE-2016-5734)	<img alt="forks" src="https://img.shields.io/github/forks/KosukeShimofuji/CVE-2016-5734">	<img alt="stars" src="https://img.shields.io/github/stars/KosukeShimofuji/CVE-2016-5734">

---
## CVE-2016-5399 (2017-04-21T20:59:00)
> The bzread function in ext/bz2/bz2.c in PHP before 5.5.38, 5.6.x before 5.6.24, and 7.x before 7.0.9 allows remote attackers to cause a denial of service (out-of-bounds write) or execute arbitrary code via a crafted bz2 archive.
- [Live-Hack-CVE/CVE-2016-5399](https://github.com/Live-Hack-CVE/CVE-2016-5399)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-5399">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-5399">

---
## CVE-2016-5386 (2016-07-19T02:00:00)
> The net/http package in Go through 1.6 does not attempt to address RFC 3875 section 4.1.18 namespace conflicts and therefore does not protect CGI applications from the presence of untrusted client data in the HTTP_PROXY environment variable, which might allow remote attackers to redirect a CGI application's outbound HTTP traffic to an arbitrary proxy server via a crafted Proxy header in an HTTP request, aka an "httpoxy" issue.
- [Live-Hack-CVE/CVE-2016-5386](https://github.com/Live-Hack-CVE/CVE-2016-5386)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-5386">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-5386">

---
## CVE-2016-5385 (2016-07-19T02:00:00)
> PHP through 7.0.8 does not attempt to address RFC 3875 section 4.1.18 namespace conflicts and therefore does not protect applications from the presence of untrusted client data in the HTTP_PROXY environment variable, which might allow remote attackers to redirect an application's outbound HTTP traffic to an arbitrary proxy server via a crafted Proxy header in an HTTP request, as demonstrated by (1) an application that makes a getenv('HTTP_PROXY') call or (2) a CGI configuration of PHP, aka an "httpoxy" issue.
- [Live-Hack-CVE/CVE-2016-5385](https://github.com/Live-Hack-CVE/CVE-2016-5385)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-5385">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-5385">

---
## CVE-2016-5195 (2016-11-10T21:59:00)
> Race condition in mm/gup.c in the Linux kernel 2.x through 4.x before 4.8.3 allows local users to gain privileges by leveraging incorrect handling of a copy-on-write (COW) feature to write to a read-only memory mapping, as exploited in the wild in October 2016, aka "Dirty COW."
- [fei9747/CVE-2016-5195](https://github.com/fei9747/CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/fei9747/CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/fei9747/CVE-2016-5195">
- [flux10n/dirtycow](https://github.com/flux10n/dirtycow)	<img alt="forks" src="https://img.shields.io/github/forks/flux10n/dirtycow">	<img alt="stars" src="https://img.shields.io/github/stars/flux10n/dirtycow">
- [KaviDk/dirtyCow](https://github.com/KaviDk/dirtyCow)	<img alt="forks" src="https://img.shields.io/github/forks/KaviDk/dirtyCow">	<img alt="stars" src="https://img.shields.io/github/stars/KaviDk/dirtyCow">
- [1equeneRise/scumjr9](https://github.com/1equeneRise/scumjr9)	<img alt="forks" src="https://img.shields.io/github/forks/1equeneRise/scumjr9">	<img alt="stars" src="https://img.shields.io/github/stars/1equeneRise/scumjr9">
- [malinthag62/The-exploitation-of-Dirty-Cow-CVE-2016-5195](https://github.com/malinthag62/The-exploitation-of-Dirty-Cow-CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/malinthag62/The-exploitation-of-Dirty-Cow-CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/malinthag62/The-exploitation-of-Dirty-Cow-CVE-2016-5195">
- [gurpreetsinghsaluja/dirtycow](https://github.com/gurpreetsinghsaluja/dirtycow)	<img alt="forks" src="https://img.shields.io/github/forks/gurpreetsinghsaluja/dirtycow">	<img alt="stars" src="https://img.shields.io/github/stars/gurpreetsinghsaluja/dirtycow">
- [passionchenjianyegmail8/scumjrs](https://github.com/passionchenjianyegmail8/scumjrs)	<img alt="forks" src="https://img.shields.io/github/forks/passionchenjianyegmail8/scumjrs">	<img alt="stars" src="https://img.shields.io/github/stars/passionchenjianyegmail8/scumjrs">
- [TotallyNotAHaxxer/CVE-2016-5195](https://github.com/TotallyNotAHaxxer/CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/TotallyNotAHaxxer/CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/TotallyNotAHaxxer/CVE-2016-5195">
- [vinspiert/scumjrs](https://github.com/vinspiert/scumjrs)	<img alt="forks" src="https://img.shields.io/github/forks/vinspiert/scumjrs">	<img alt="stars" src="https://img.shields.io/github/stars/vinspiert/scumjrs">
- [scumjr/dirtycow-vdso](https://github.com/scumjr/dirtycow-vdso)	<img alt="forks" src="https://img.shields.io/github/forks/scumjr/dirtycow-vdso">	<img alt="stars" src="https://img.shields.io/github/stars/scumjr/dirtycow-vdso">
- [r1is/CVE-2022-0847](https://github.com/r1is/CVE-2022-0847)	<img alt="forks" src="https://img.shields.io/github/forks/r1is/CVE-2022-0847">	<img alt="stars" src="https://img.shields.io/github/stars/r1is/CVE-2022-0847">
- [ellietoulabi/Dirty-Cow](https://github.com/ellietoulabi/Dirty-Cow)	<img alt="forks" src="https://img.shields.io/github/forks/ellietoulabi/Dirty-Cow">	<img alt="stars" src="https://img.shields.io/github/stars/ellietoulabi/Dirty-Cow">
- [th3-5had0w/DirtyCOW-PoC](https://github.com/th3-5had0w/DirtyCOW-PoC)	<img alt="forks" src="https://img.shields.io/github/forks/th3-5had0w/DirtyCOW-PoC">	<img alt="stars" src="https://img.shields.io/github/stars/th3-5had0w/DirtyCOW-PoC">
- [KasunPriyashan/Y2S1-Project-Linux-Exploitaion-using-CVE-2016-5195-Vulnerability](https://github.com/KasunPriyashan/Y2S1-Project-Linux-Exploitaion-using-CVE-2016-5195-Vulnerability)	<img alt="forks" src="https://img.shields.io/github/forks/KasunPriyashan/Y2S1-Project-Linux-Exploitaion-using-CVE-2016-5195-Vulnerability">	<img alt="stars" src="https://img.shields.io/github/stars/KasunPriyashan/Y2S1-Project-Linux-Exploitaion-using-CVE-2016-5195-Vulnerability">
- [nazgul6092/2nd-Year-Project-01-Linux-Exploitation-using-CVE-20166-5195](https://github.com/nazgul6092/2nd-Year-Project-01-Linux-Exploitation-using-CVE-20166-5195)	<img alt="forks" src="https://img.shields.io/github/forks/nazgul6092/2nd-Year-Project-01-Linux-Exploitation-using-CVE-20166-5195">	<img alt="stars" src="https://img.shields.io/github/stars/nazgul6092/2nd-Year-Project-01-Linux-Exploitation-using-CVE-20166-5195">
- [arttnba3/CVE-2016-5195](https://github.com/arttnba3/CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/arttnba3/CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/arttnba3/CVE-2016-5195">
- [firefart/dirtycow](https://github.com/firefart/dirtycow)	<img alt="forks" src="https://img.shields.io/github/forks/firefart/dirtycow">	<img alt="stars" src="https://img.shields.io/github/stars/firefart/dirtycow">
- [DanielEbert/CVE-2016-5195](https://github.com/DanielEbert/CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/DanielEbert/CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/DanielEbert/CVE-2016-5195">
- [timwr/CVE-2016-5195](https://github.com/timwr/CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/timwr/CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/timwr/CVE-2016-5195">
- [zakariamaaraki/Dirty-COW-CVE-2016-5195-](https://github.com/zakariamaaraki/Dirty-COW-CVE-2016-5195-)	<img alt="forks" src="https://img.shields.io/github/forks/zakariamaaraki/Dirty-COW-CVE-2016-5195-">	<img alt="stars" src="https://img.shields.io/github/stars/zakariamaaraki/Dirty-COW-CVE-2016-5195-">
- [sandeeparth07/CVE-2016_5195-vulnarability](https://github.com/sandeeparth07/CVE-2016_5195-vulnarability)	<img alt="forks" src="https://img.shields.io/github/forks/sandeeparth07/CVE-2016_5195-vulnarability">	<img alt="stars" src="https://img.shields.io/github/stars/sandeeparth07/CVE-2016_5195-vulnarability">
- [dulanjaya23/Dirty-Cow-CVE-2016-5195-](https://github.com/dulanjaya23/Dirty-Cow-CVE-2016-5195-)	<img alt="forks" src="https://img.shields.io/github/forks/dulanjaya23/Dirty-Cow-CVE-2016-5195-">	<img alt="stars" src="https://img.shields.io/github/stars/dulanjaya23/Dirty-Cow-CVE-2016-5195-">
- [shanuka-ashen/Dirty-Cow-Explanation-CVE-2016-5195-](https://github.com/shanuka-ashen/Dirty-Cow-Explanation-CVE-2016-5195-)	<img alt="forks" src="https://img.shields.io/github/forks/shanuka-ashen/Dirty-Cow-Explanation-CVE-2016-5195-">	<img alt="stars" src="https://img.shields.io/github/stars/shanuka-ashen/Dirty-Cow-Explanation-CVE-2016-5195-">
- [imust6226/dirtcow](https://github.com/imust6226/dirtcow)	<img alt="forks" src="https://img.shields.io/github/forks/imust6226/dirtcow">	<img alt="stars" src="https://img.shields.io/github/stars/imust6226/dirtcow">
- [jas502n/CVE-2016-5195](https://github.com/jas502n/CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/jas502n/CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/jas502n/CVE-2016-5195">
- [Brucetg/DirtyCow-EXP](https://github.com/Brucetg/DirtyCow-EXP)	<img alt="forks" src="https://img.shields.io/github/forks/Brucetg/DirtyCow-EXP">	<img alt="stars" src="https://img.shields.io/github/stars/Brucetg/DirtyCow-EXP">
- [xpcmdshell/derpyc0w](https://github.com/xpcmdshell/derpyc0w)	<img alt="forks" src="https://img.shields.io/github/forks/xpcmdshell/derpyc0w">	<img alt="stars" src="https://img.shields.io/github/stars/xpcmdshell/derpyc0w">
- [acidburnmi/CVE-2016-5195-master](https://github.com/acidburnmi/CVE-2016-5195-master)	<img alt="forks" src="https://img.shields.io/github/forks/acidburnmi/CVE-2016-5195-master">	<img alt="stars" src="https://img.shields.io/github/stars/acidburnmi/CVE-2016-5195-master">
- [titanhp/Dirty-COW-CVE-2016-5195-Testing](https://github.com/titanhp/Dirty-COW-CVE-2016-5195-Testing)	<img alt="forks" src="https://img.shields.io/github/forks/titanhp/Dirty-COW-CVE-2016-5195-Testing">	<img alt="stars" src="https://img.shields.io/github/stars/titanhp/Dirty-COW-CVE-2016-5195-Testing">
- [arbll/dirtycow](https://github.com/arbll/dirtycow)	<img alt="forks" src="https://img.shields.io/github/forks/arbll/dirtycow">	<img alt="stars" src="https://img.shields.io/github/stars/arbll/dirtycow">
- [NguyenCongHaiNam/Research-CVE-2016-5195](https://github.com/NguyenCongHaiNam/Research-CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/NguyenCongHaiNam/Research-CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/NguyenCongHaiNam/Research-CVE-2016-5195">
- [LinuxKernelContent/DirtyCow](https://github.com/LinuxKernelContent/DirtyCow)	<img alt="forks" src="https://img.shields.io/github/forks/LinuxKernelContent/DirtyCow">	<img alt="stars" src="https://img.shields.io/github/stars/LinuxKernelContent/DirtyCow">
- [idhyt/androotzf](https://github.com/idhyt/androotzf)	<img alt="forks" src="https://img.shields.io/github/forks/idhyt/androotzf">	<img alt="stars" src="https://img.shields.io/github/stars/idhyt/androotzf">
- [EDLLT/CVE-2016-5195-master](https://github.com/EDLLT/CVE-2016-5195-master)	<img alt="forks" src="https://img.shields.io/github/forks/EDLLT/CVE-2016-5195-master">	<img alt="stars" src="https://img.shields.io/github/stars/EDLLT/CVE-2016-5195-master">
- [ASUKA39/CVE-2016-5195](https://github.com/ASUKA39/CVE-2016-5195)	<img alt="forks" src="https://img.shields.io/github/forks/ASUKA39/CVE-2016-5195">	<img alt="stars" src="https://img.shields.io/github/stars/ASUKA39/CVE-2016-5195">
- [sakilahamed/Linux-Kernel-Exploit-LAB](https://github.com/sakilahamed/Linux-Kernel-Exploit-LAB)	<img alt="forks" src="https://img.shields.io/github/forks/sakilahamed/Linux-Kernel-Exploit-LAB">	<img alt="stars" src="https://img.shields.io/github/stars/sakilahamed/Linux-Kernel-Exploit-LAB">
- [B1ackCat/cve-2016-5195-DirtyCOW](https://github.com/B1ackCat/cve-2016-5195-DirtyCOW)	<img alt="forks" src="https://img.shields.io/github/forks/B1ackCat/cve-2016-5195-DirtyCOW">	<img alt="stars" src="https://img.shields.io/github/stars/B1ackCat/cve-2016-5195-DirtyCOW">

---
## CVE-2016-5180 (2016-10-03T15:59:00)
> Heap-based buffer overflow in the ares_create_query function in c-ares 1.x before 1.12.0 allows remote attackers to cause a denial of service (out-of-bounds write) or possibly execute arbitrary code via a hostname with an escaped trailing dot.
- [Live-Hack-CVE/CVE-2016-5180](https://github.com/Live-Hack-CVE/CVE-2016-5180)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-5180">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-5180">

---
## CVE-2016-4655 (2016-08-25T21:59:00)
> The kernel in Apple iOS before 9.3.5 allows attackers to obtain sensitive information from memory via a crafted app.
- [0xyf77/CVE-2016-4655](https://github.com/0xyf77/CVE-2016-4655)	<img alt="forks" src="https://img.shields.io/github/forks/0xyf77/CVE-2016-4655">	<img alt="stars" src="https://img.shields.io/github/stars/0xyf77/CVE-2016-4655">
- [liangle1986126z/jndok](https://github.com/liangle1986126z/jndok)	<img alt="forks" src="https://img.shields.io/github/forks/liangle1986126z/jndok">	<img alt="stars" src="https://img.shields.io/github/stars/liangle1986126z/jndok">
- [tomitokics/br0ke](https://github.com/tomitokics/br0ke)	<img alt="forks" src="https://img.shields.io/github/forks/tomitokics/br0ke">	<img alt="stars" src="https://img.shields.io/github/stars/tomitokics/br0ke">
- [Cryptiiiic/skybreak](https://github.com/Cryptiiiic/skybreak)	<img alt="forks" src="https://img.shields.io/github/forks/Cryptiiiic/skybreak">	<img alt="stars" src="https://img.shields.io/github/stars/Cryptiiiic/skybreak">
- [jndok/PegasusX](https://github.com/jndok/PegasusX)	<img alt="forks" src="https://img.shields.io/github/forks/jndok/PegasusX">	<img alt="stars" src="https://img.shields.io/github/stars/jndok/PegasusX">

---
## CVE-2016-4508 (2016-07-06T14:59:00)
> Cross-site scripting (XSS) vulnerability in Rexroth Bosch BLADEcontrol-WebVIS 3.0.2 and earlier allows remote attackers to inject arbitrary web script or HTML via unspecified vectors.
- [Live-Hack-CVE/CVE-2016-4508](https://github.com/Live-Hack-CVE/CVE-2016-4508)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4508">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4508">

---
## CVE-2016-4507 (2016-07-06T14:59:00)
> SQL injection vulnerability in Rexroth Bosch BLADEcontrol-WebVIS 3.0.2 and earlier allows remote authenticated users to execute arbitrary SQL commands via unspecified vectors.
- [Live-Hack-CVE/CVE-2016-4507](https://github.com/Live-Hack-CVE/CVE-2016-4507)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4507">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4507">

---
## CVE-2016-4437 (2016-06-07T14:06:00)
> Apache Shiro before 1.2.5, when a cipher key has not been configured for the "remember me" feature, allows remote attackers to execute arbitrary code or bypass intended access restrictions via an unspecified request parameter.
- [WSIDFY/-WSIDFY-CVE-2016-4437](https://github.com/WSIDFY/-WSIDFY-CVE-2016-4437)	<img alt="forks" src="https://img.shields.io/github/forks/WSIDFY/-WSIDFY-CVE-2016-4437">	<img alt="stars" src="https://img.shields.io/github/stars/WSIDFY/-WSIDFY-CVE-2016-4437">
- [XuCcc/VulEnv](https://github.com/XuCcc/VulEnv)	<img alt="forks" src="https://img.shields.io/github/forks/XuCcc/VulEnv">	<img alt="stars" src="https://img.shields.io/github/stars/XuCcc/VulEnv">
- [zhzyker/vulmap](https://github.com/zhzyker/vulmap)	<img alt="forks" src="https://img.shields.io/github/forks/zhzyker/vulmap">	<img alt="stars" src="https://img.shields.io/github/stars/zhzyker/vulmap">
- [4nth0ny1130/shisoserial](https://github.com/4nth0ny1130/shisoserial)	<img alt="forks" src="https://img.shields.io/github/forks/4nth0ny1130/shisoserial">	<img alt="stars" src="https://img.shields.io/github/stars/4nth0ny1130/shisoserial">
- [m3terpreter/CVE-2016-4437](https://github.com/m3terpreter/CVE-2016-4437)	<img alt="forks" src="https://img.shields.io/github/forks/m3terpreter/CVE-2016-4437">	<img alt="stars" src="https://img.shields.io/github/stars/m3terpreter/CVE-2016-4437">
- [bkfish/Awesome_shiro](https://github.com/bkfish/Awesome_shiro)	<img alt="forks" src="https://img.shields.io/github/forks/bkfish/Awesome_shiro">	<img alt="stars" src="https://img.shields.io/github/stars/bkfish/Awesome_shiro">
- [pizza-power/CVE-2016-4437](https://github.com/pizza-power/CVE-2016-4437)	<img alt="forks" src="https://img.shields.io/github/forks/pizza-power/CVE-2016-4437">	<img alt="stars" src="https://img.shields.io/github/stars/pizza-power/CVE-2016-4437">
- [xk-mt/CVE-2016-4437](https://github.com/xk-mt/CVE-2016-4437)	<img alt="forks" src="https://img.shields.io/github/forks/xk-mt/CVE-2016-4437">	<img alt="stars" src="https://img.shields.io/github/stars/xk-mt/CVE-2016-4437">

---
## CVE-2016-4432 (2016-06-01T20:59:00)
> The AMQP 0-8, 0-9, 0-91, and 0-10 connection handling in Apache Qpid Java before 6.0.3 might allow remote attackers to bypass authentication and consequently perform actions via vectors related to connection state logging.
- [Live-Hack-CVE/CVE-2016-4432](https://github.com/Live-Hack-CVE/CVE-2016-4432)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4432">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4432">

---
## CVE-2016-4287 (2016-09-14T18:59:00)
> Integer overflow in Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code via unspecified vectors.
- [Live-Hack-CVE/CVE-2016-4287](https://github.com/Live-Hack-CVE/CVE-2016-4287)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4287">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4287">
- [Live-Hack-CVE/CVE-2016-4287](https://github.com/Live-Hack-CVE/CVE-2016-4287)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4287">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4287">

---
## CVE-2016-4285 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4285](https://github.com/Live-Hack-CVE/CVE-2016-4285)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4285">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4285">

---
## CVE-2016-4284 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4284](https://github.com/Live-Hack-CVE/CVE-2016-4284)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4284">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4284">

---
## CVE-2016-4283 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4284, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4283](https://github.com/Live-Hack-CVE/CVE-2016-4283)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4283">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4283">

---
## CVE-2016-4282 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4282](https://github.com/Live-Hack-CVE/CVE-2016-4282)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4282">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4282">
- [Live-Hack-CVE/CVE-2016-4282](https://github.com/Live-Hack-CVE/CVE-2016-4282)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4282">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4282">

---
## CVE-2016-4281 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4281](https://github.com/Live-Hack-CVE/CVE-2016-4281)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4281">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4281">
- [Live-Hack-CVE/CVE-2016-4281](https://github.com/Live-Hack-CVE/CVE-2016-4281)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4281">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4281">

---
## CVE-2016-4280 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4276, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4280](https://github.com/Live-Hack-CVE/CVE-2016-4280)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4280">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4280">
- [Live-Hack-CVE/CVE-2016-4280](https://github.com/Live-Hack-CVE/CVE-2016-4280)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4280">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4280">

---
## CVE-2016-4277 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to bypass intended access restrictions and obtain sensitive information via unspecified vectors, a different vulnerability than CVE-2016-4271 and CVE-2016-4278.
- [Live-Hack-CVE/CVE-2016-4277](https://github.com/Live-Hack-CVE/CVE-2016-4277)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4277">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4277">
- [Live-Hack-CVE/CVE-2016-4277](https://github.com/Live-Hack-CVE/CVE-2016-4277)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4277">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4277">
- [Live-Hack-CVE/CVE-2016-4271](https://github.com/Live-Hack-CVE/CVE-2016-4271)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4271">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4271">
- [Live-Hack-CVE/CVE-2016-4278](https://github.com/Live-Hack-CVE/CVE-2016-4278)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4278">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4278">

---
## CVE-2016-4276 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4275, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4276](https://github.com/Live-Hack-CVE/CVE-2016-4276)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4276">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4276">

---
## CVE-2016-4275 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4274, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4275](https://github.com/Live-Hack-CVE/CVE-2016-4275)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4275">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4275">

---
## CVE-2016-4274 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-4275, CVE-2016-4276, CVE-2016-4280, CVE-2016-4281, CVE-2016-4282, CVE-2016-4283, CVE-2016-4284, CVE-2016-4285, CVE-2016-6922, and CVE-2016-6924.
- [Live-Hack-CVE/CVE-2016-4274](https://github.com/Live-Hack-CVE/CVE-2016-4274)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4274">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4274">

---
## CVE-2016-4273 (2016-10-13T19:59:00)
> Adobe Flash Player before 18.0.0.382 and 19.x through 23.x before 23.0.0.185 on Windows and OS X and before 11.2.202.637 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-6982, CVE-2016-6983, CVE-2016-6984, CVE-2016-6985, CVE-2016-6986, CVE-2016-6989, and CVE-2016-6990.
- [Live-Hack-CVE/CVE-2016-4273](https://github.com/Live-Hack-CVE/CVE-2016-4273)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4273">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4273">
- [Live-Hack-CVE/CVE-2016-4273](https://github.com/Live-Hack-CVE/CVE-2016-4273)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4273">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4273">

---
## CVE-2016-4271 (2016-09-14T18:59:00)
> Adobe Flash Player before 18.0.0.375 and 19.x through 23.x before 23.0.0.162 on Windows and OS X and before 11.2.202.635 on Linux allows attackers to bypass intended access restrictions and obtain sensitive information via unspecified vectors, a different vulnerability than CVE-2016-4277 and CVE-2016-4278, aka a "local-with-filesystem Flash sandbox bypass" issue.
- [Live-Hack-CVE/CVE-2016-4271](https://github.com/Live-Hack-CVE/CVE-2016-4271)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4271">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4271">
- [Live-Hack-CVE/CVE-2016-4278](https://github.com/Live-Hack-CVE/CVE-2016-4278)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4278">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4278">

---
## CVE-2016-4163 (2016-06-16T14:59:00)
> Adobe Flash Player before 18.0.0.352 and 19.x through 21.x before 21.0.0.242 on Windows and OS X and before 11.2.202.621 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-1096, CVE-2016-1098, CVE-2016-1099, CVE-2016-1100, CVE-2016-1102, CVE-2016-1104, CVE-2016-4109, CVE-2016-4111, CVE-2016-4112, CVE-2016-4113, CVE-2016-4114, CVE-2016-4115, CVE-2016-4120, CVE-2016-4160, CVE-2016-4161, and CVE-2016-4162.
- [Live-Hack-CVE/CVE-2016-4163](https://github.com/Live-Hack-CVE/CVE-2016-4163)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4163">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4163">
- [Live-Hack-CVE/CVE-2016-4163](https://github.com/Live-Hack-CVE/CVE-2016-4163)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4163">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4163">

---
## CVE-2016-4162 (2016-06-16T14:59:00)
> Adobe Flash Player before 18.0.0.352 and 19.x through 21.x before 21.0.0.242 on Windows and OS X and before 11.2.202.621 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-1096, CVE-2016-1098, CVE-2016-1099, CVE-2016-1100, CVE-2016-1102, CVE-2016-1104, CVE-2016-4109, CVE-2016-4111, CVE-2016-4112, CVE-2016-4113, CVE-2016-4114, CVE-2016-4115, CVE-2016-4120, CVE-2016-4160, CVE-2016-4161, and CVE-2016-4163.
- [Live-Hack-CVE/CVE-2016-4162](https://github.com/Live-Hack-CVE/CVE-2016-4162)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4162">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4162">

---
## CVE-2016-4161 (2016-06-16T14:59:00)
> Adobe Flash Player before 18.0.0.352 and 19.x through 21.x before 21.0.0.242 on Windows and OS X and before 11.2.202.621 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-1096, CVE-2016-1098, CVE-2016-1099, CVE-2016-1100, CVE-2016-1102, CVE-2016-1104, CVE-2016-4109, CVE-2016-4111, CVE-2016-4112, CVE-2016-4113, CVE-2016-4114, CVE-2016-4115, CVE-2016-4120, CVE-2016-4160, CVE-2016-4162, and CVE-2016-4163.
- [Live-Hack-CVE/CVE-2016-4161](https://github.com/Live-Hack-CVE/CVE-2016-4161)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4161">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4161">
- [Live-Hack-CVE/CVE-2016-4161](https://github.com/Live-Hack-CVE/CVE-2016-4161)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4161">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4161">

---
## CVE-2016-4160 (2016-06-16T14:59:00)
> Adobe Flash Player before 18.0.0.352 and 19.x through 21.x before 21.0.0.242 on Windows and OS X and before 11.2.202.621 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-1096, CVE-2016-1098, CVE-2016-1099, CVE-2016-1100, CVE-2016-1102, CVE-2016-1104, CVE-2016-4109, CVE-2016-4111, CVE-2016-4112, CVE-2016-4113, CVE-2016-4114, CVE-2016-4115, CVE-2016-4120, CVE-2016-4161, CVE-2016-4162, and CVE-2016-4163.
- [Live-Hack-CVE/CVE-2016-4160](https://github.com/Live-Hack-CVE/CVE-2016-4160)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4160">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4160">
- [Live-Hack-CVE/CVE-2016-4160](https://github.com/Live-Hack-CVE/CVE-2016-4160)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4160">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4160">

---
## CVE-2016-4121 (2016-06-16T14:59:00)
> Use-after-free vulnerability in Adobe Flash Player before 18.0.0.352 and 19.x through 21.x before 21.0.0.242 on Windows and OS X and before 11.2.202.621 on Linux allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-1097, CVE-2016-1106, CVE-2016-1107, CVE-2016-1108, CVE-2016-1109, CVE-2016-1110, CVE-2016-4108, and CVE-2016-4110.
- [Live-Hack-CVE/CVE-2016-4121](https://github.com/Live-Hack-CVE/CVE-2016-4121)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4121">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4121">

---
## CVE-2016-4120 (2016-06-16T14:59:00)
> Adobe Flash Player before 18.0.0.352 and 19.x through 21.x before 21.0.0.242 on Windows and OS X and before 11.2.202.621 on Linux allows attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-1096, CVE-2016-1098, CVE-2016-1099, CVE-2016-1100, CVE-2016-1102, CVE-2016-1104, CVE-2016-4109, CVE-2016-4111, CVE-2016-4112, CVE-2016-4113, CVE-2016-4114, CVE-2016-4115, CVE-2016-4160, CVE-2016-4161, CVE-2016-4162, and CVE-2016-4163.
- [Live-Hack-CVE/CVE-2016-4120](https://github.com/Live-Hack-CVE/CVE-2016-4120)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4120">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4120">
- [Live-Hack-CVE/CVE-2016-4120](https://github.com/Live-Hack-CVE/CVE-2016-4120)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-4120">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-4120">

---
## CVE-2016-3958 (2016-05-23T19:59:00)
> Untrusted search path vulnerability in Go before 1.5.4 and 1.6.x before 1.6.1 on Windows allows local users to gain privileges via a Trojan horse DLL in the current working directory, related to use of the LoadLibrary function.
- [Live-Hack-CVE/CVE-2016-3958](https://github.com/Live-Hack-CVE/CVE-2016-3958)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-3958">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-3958">

---
## CVE-2016-3955 (2016-07-03T21:59:00)
> The usbip_recv_xbuff function in drivers/usb/usbip/usbip_common.c in the Linux kernel before 4.5.3 allows remote attackers to cause a denial of service (out-of-bounds write) or possibly have unspecified other impact via a crafted length value in a USB/IP packet.
- [Live-Hack-CVE/CVE-2016-3955](https://github.com/Live-Hack-CVE/CVE-2016-3955)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-3955">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-3955">

---
## CVE-2016-3861 (2016-09-11T21:59:00)
> LibUtils in Android 4.x before 4.4.4, 5.0.x before 5.0.2, 5.1.x before 5.1.1, 6.x before 2016-09-01, and 7.0 before 2016-09-01 mishandles conversions between Unicode character encodings with different encoding widths, which allows remote attackers to execute arbitrary code or cause a denial of service (heap-based buffer overflow) via a crafted file, aka internal bug 29250543.
- [dropk1ck/CVE-2016-3861](https://github.com/dropk1ck/CVE-2016-3861)	<img alt="forks" src="https://img.shields.io/github/forks/dropk1ck/CVE-2016-3861">	<img alt="stars" src="https://img.shields.io/github/stars/dropk1ck/CVE-2016-3861">

---
## CVE-2016-3714 (2016-05-05T18:59:00)
> The (1) EPHEMERAL, (2) HTTPS, (3) MVG, (4) MSL, (5) TEXT, (6) SHOW, (7) WIN, and (8) PLT coders in ImageMagick before 6.9.3-10 and 7.x before 7.0.1-1 allow remote attackers to execute arbitrary code via shell metacharacters in a crafted image, aka "ImageTragick."
- [JoshMorrison99/CVE-2016-3714](https://github.com/JoshMorrison99/CVE-2016-3714)	<img alt="forks" src="https://img.shields.io/github/forks/JoshMorrison99/CVE-2016-3714">	<img alt="stars" src="https://img.shields.io/github/stars/JoshMorrison99/CVE-2016-3714">
- [MrrRaph/pandagik](https://github.com/MrrRaph/pandagik)	<img alt="forks" src="https://img.shields.io/github/forks/MrrRaph/pandagik">	<img alt="stars" src="https://img.shields.io/github/stars/MrrRaph/pandagik">
- [shelld3v/RCE-python-oneliner-payload](https://github.com/shelld3v/RCE-python-oneliner-payload)	<img alt="forks" src="https://img.shields.io/github/forks/shelld3v/RCE-python-oneliner-payload">	<img alt="stars" src="https://img.shields.io/github/stars/shelld3v/RCE-python-oneliner-payload">
- [artfreyr/wp-imagetragick](https://github.com/artfreyr/wp-imagetragick)	<img alt="forks" src="https://img.shields.io/github/forks/artfreyr/wp-imagetragick">	<img alt="stars" src="https://img.shields.io/github/stars/artfreyr/wp-imagetragick">
- [mike-williams/imagetragick-poc](https://github.com/mike-williams/imagetragick-poc)	<img alt="forks" src="https://img.shields.io/github/forks/mike-williams/imagetragick-poc">	<img alt="stars" src="https://img.shields.io/github/stars/mike-williams/imagetragick-poc">
- [jpeanut/ImageTragick-CVE-2016-3714-RShell](https://github.com/jpeanut/ImageTragick-CVE-2016-3714-RShell)	<img alt="forks" src="https://img.shields.io/github/forks/jpeanut/ImageTragick-CVE-2016-3714-RShell">	<img alt="stars" src="https://img.shields.io/github/stars/jpeanut/ImageTragick-CVE-2016-3714-RShell">
- [ImageTragick/PoCs](https://github.com/ImageTragick/PoCs)	<img alt="forks" src="https://img.shields.io/github/forks/ImageTragick/PoCs">	<img alt="stars" src="https://img.shields.io/github/stars/ImageTragick/PoCs">
- [chusiang/CVE-2016-3714.ansible.role](https://github.com/chusiang/CVE-2016-3714.ansible.role)	<img alt="forks" src="https://img.shields.io/github/forks/chusiang/CVE-2016-3714.ansible.role">	<img alt="stars" src="https://img.shields.io/github/stars/chusiang/CVE-2016-3714.ansible.role">
- [Hood3dRob1n/CVE-2016-3714](https://github.com/Hood3dRob1n/CVE-2016-3714)	<img alt="forks" src="https://img.shields.io/github/forks/Hood3dRob1n/CVE-2016-3714">	<img alt="stars" src="https://img.shields.io/github/stars/Hood3dRob1n/CVE-2016-3714">
- [tommiionfire/CVE-2016-3714](https://github.com/tommiionfire/CVE-2016-3714)	<img alt="forks" src="https://img.shields.io/github/forks/tommiionfire/CVE-2016-3714">	<img alt="stars" src="https://img.shields.io/github/stars/tommiionfire/CVE-2016-3714">
- [jackdpeterson/imagick_secure_puppet](https://github.com/jackdpeterson/imagick_secure_puppet)	<img alt="forks" src="https://img.shields.io/github/forks/jackdpeterson/imagick_secure_puppet">	<img alt="stars" src="https://img.shields.io/github/stars/jackdpeterson/imagick_secure_puppet">

---
## CVE-2016-3709 (2022-07-28T17:15:00)
> Possible cross-site scripting vulnerability in libxml after commit 960f0e2.
- [Live-Hack-CVE/CVE-2016-3709](https://github.com/Live-Hack-CVE/CVE-2016-3709)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-3709">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-3709">

---
## CVE-2016-3189 (2016-06-30T17:59:00)
> Use-after-free vulnerability in bzip2recover in bzip2 1.0.6 allows remote attackers to cause a denial of service (crash) via a crafted bzip2 file, related to block ends set to before the start of the block.
- [Live-Hack-CVE/CVE-2016-3189](https://github.com/Live-Hack-CVE/CVE-2016-3189)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-3189">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-3189">

---
## CVE-2016-2338 (2022-09-29T03:15:00)
> An exploitable heap overflow vulnerability exists in the Psych::Emitter start_document function of Ruby. In Psych::Emitter start_document function heap buffer "head" allocation is made based on tags array length. Specially constructed object passed as element of tags array can increase this array size after mentioned allocation and cause heap overflow.
- [SpiralBL0CK/CVE-2016-2338-nday](https://github.com/SpiralBL0CK/CVE-2016-2338-nday)	<img alt="forks" src="https://img.shields.io/github/forks/SpiralBL0CK/CVE-2016-2338-nday">	<img alt="stars" src="https://img.shields.io/github/stars/SpiralBL0CK/CVE-2016-2338-nday">

---
## CVE-2016-2182 (2016-09-16T05:59:00)
> The BN_bn2dec function in crypto/bn/bn_print.c in OpenSSL before 1.1.0 does not properly validate division results, which allows remote attackers to cause a denial of service (out-of-bounds write and application crash) or possibly have unspecified other impact via unknown vectors.
- [Live-Hack-CVE/CVE-2016-2182](https://github.com/Live-Hack-CVE/CVE-2016-2182)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2182">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2182">

---
## CVE-2016-2181 (2016-09-16T05:59:00)
> The Anti-Replay feature in the DTLS implementation in OpenSSL before 1.1.0 mishandles early use of a new epoch number in conjunction with a large sequence number, which allows remote attackers to cause a denial of service (false-positive packet drops) via spoofed DTLS records, related to rec_layer_d1.c and ssl3_record.c.
- [Live-Hack-CVE/CVE-2016-2181](https://github.com/Live-Hack-CVE/CVE-2016-2181)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2181">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2181">

---
## CVE-2016-2180 (2016-08-01T02:59:00)
> The TS_OBJ_print_bio function in crypto/ts/ts_lib.c in the X.509 Public Key Infrastructure Time-Stamp Protocol (TSP) implementation in OpenSSL through 1.0.2h allows remote attackers to cause a denial of service (out-of-bounds read and application crash) via a crafted time-stamp file that is mishandled by the "openssl ts" command.
- [Live-Hack-CVE/CVE-2016-2180](https://github.com/Live-Hack-CVE/CVE-2016-2180)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2180">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2180">
- [Live-Hack-CVE/CVE-2016-2180](https://github.com/Live-Hack-CVE/CVE-2016-2180)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2180">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2180">

---
## CVE-2016-2178 (2016-06-20T01:59:00)
> The dsa_sign_setup function in crypto/dsa/dsa_ossl.c in OpenSSL through 1.0.2h does not properly ensure the use of constant-time operations, which makes it easier for local users to discover a DSA private key via a timing side-channel attack.
- [Live-Hack-CVE/CVE-2016-2178](https://github.com/Live-Hack-CVE/CVE-2016-2178)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2178">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2178">
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-2178](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2178)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2178">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2178">

---
## CVE-2016-2177 (2016-06-20T01:59:00)
> OpenSSL through 1.0.2h incorrectly uses pointer arithmetic for heap-buffer boundary checks, which might allow remote attackers to cause a denial of service (integer overflow and application crash) or possibly have unspecified other impact by leveraging unexpected malloc behavior, related to s3_srvr.c, ssl_sess.c, and t1_lib.c.
- [Live-Hack-CVE/CVE-2016-2177](https://github.com/Live-Hack-CVE/CVE-2016-2177)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2177">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2177">

---
## CVE-2016-2176 (2016-05-05T01:59:00)
> The X509_NAME_oneline function in crypto/x509/x509_obj.c in OpenSSL before 1.0.1t and 1.0.2 before 1.0.2h allows remote attackers to obtain sensitive information from process stack memory or cause a denial of service (buffer over-read) via crafted EBCDIC ASN.1 data.
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-2176](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2176)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2176">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2176">

---
## CVE-2016-2125 (2018-10-31T20:29:00)
> It was found that Samba before versions 4.5.3, 4.4.8, 4.3.13 always requested forwardable tickets when using Kerberos authentication. A service to which Samba authenticated using Kerberos could subsequently use the ticket to impersonate Samba to other services or domain users.
- [Live-Hack-CVE/CVE-2016-2125](https://github.com/Live-Hack-CVE/CVE-2016-2125)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2125">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2125">

---
## CVE-2016-2123 (2018-11-01T13:29:00)
> A flaw was found in samba versions 4.0.0 to 4.5.2. The Samba routine ndr_pull_dnsp_name contains an integer wrap problem, leading to an attacker-controlled memory overwrite. ndr_pull_dnsp_name parses data from the Samba Active Directory ldb database. Any user who can write to the dnsRecord attribute over LDAP can trigger this memory corruption. By default, all authenticated LDAP users can write to the dnsRecord attribute on new DNS objects. This makes the defect a remote privilege escalation.
- [Live-Hack-CVE/CVE-2016-2123](https://github.com/Live-Hack-CVE/CVE-2016-2123)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2123">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2123">

---
## CVE-2016-2119 (2016-07-07T15:59:00)
> libcli/smb/smbXcli_base.c in Samba 4.x before 4.2.14, 4.3.x before 4.3.11, and 4.4.x before 4.4.5 allows man-in-the-middle attackers to bypass a client-signing protection mechanism, and consequently spoof SMB2 and SMB3 servers, via the (1) SMB2_SESSION_FLAG_IS_GUEST or (2) SMB2_SESSION_FLAG_IS_NULL flag.
- [Live-Hack-CVE/CVE-2016-2119](https://github.com/Live-Hack-CVE/CVE-2016-2119)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2119">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2119">

---
## CVE-2016-2118 (2016-04-12T23:59:00)
> The MS-SAMR and MS-LSAD protocol implementations in Samba 3.x and 4.x before 4.2.11, 4.3.x before 4.3.8, and 4.4.x before 4.4.2 mishandle DCERPC connections, which allows man-in-the-middle attackers to perform protocol-downgrade attacks and impersonate users by modifying the client-server data stream, aka "BADLOCK."
- [Live-Hack-CVE/CVE-2016-2118](https://github.com/Live-Hack-CVE/CVE-2016-2118)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2118">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2118">
- [nickanderson/cfengine-CVE-2016-2118](https://github.com/nickanderson/cfengine-CVE-2016-2118)	<img alt="forks" src="https://img.shields.io/github/forks/nickanderson/cfengine-CVE-2016-2118">	<img alt="stars" src="https://img.shields.io/github/stars/nickanderson/cfengine-CVE-2016-2118">

---
## CVE-2016-2109 (2016-05-05T01:59:00)
> The asn1_d2i_read_bio function in crypto/asn1/a_d2i_fp.c in the ASN.1 BIO implementation in OpenSSL before 1.0.1t and 1.0.2 before 1.0.2h allows remote attackers to cause a denial of service (memory consumption) via a short invalid encoding.
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-2109](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2109)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2109">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-2109">

---
## CVE-2016-2108 (2016-05-05T01:59:00)
> The ASN.1 implementation in OpenSSL before 1.0.1o and 1.0.2 before 1.0.2c allows remote attackers to execute arbitrary code or cause a denial of service (buffer underflow and memory corruption) via an ANY field in crafted serialized data, aka the "negative zero" issue.
- [Live-Hack-CVE/CVE-2016-2108](https://github.com/Live-Hack-CVE/CVE-2016-2108)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2108">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2108">

---
## CVE-2016-2107 (2016-05-05T01:59:00)
> The AES-NI implementation in OpenSSL before 1.0.1t and 1.0.2 before 1.0.2h does not consider memory allocation during a certain padding check, which allows remote attackers to obtain sensitive cleartext information via a padding-oracle attack against an AES CBC session. NOTE: this vulnerability exists because of an incorrect fix for CVE-2013-0169.
- [Live-Hack-CVE/CVE-2016-2107](https://github.com/Live-Hack-CVE/CVE-2016-2107)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2107">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2107">
- [FiloSottile/CVE-2016-2107](https://github.com/FiloSottile/CVE-2016-2107)	<img alt="forks" src="https://img.shields.io/github/forks/FiloSottile/CVE-2016-2107">	<img alt="stars" src="https://img.shields.io/github/stars/FiloSottile/CVE-2016-2107">
- [tmiklas/docker-cve-2016-2107](https://github.com/tmiklas/docker-cve-2016-2107)	<img alt="forks" src="https://img.shields.io/github/forks/tmiklas/docker-cve-2016-2107">	<img alt="stars" src="https://img.shields.io/github/stars/tmiklas/docker-cve-2016-2107">

---
## CVE-2016-2105 (2016-05-05T01:59:00)
> Integer overflow in the EVP_EncodeUpdate function in crypto/evp/encode.c in OpenSSL before 1.0.1t and 1.0.2 before 1.0.2h allows remote attackers to cause a denial of service (heap memory corruption) via a large amount of binary data.
- [Live-Hack-CVE/CVE-2016-2105](https://github.com/Live-Hack-CVE/CVE-2016-2105)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2105">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2105">
- [Live-Hack-CVE/CVE-2016-2105](https://github.com/Live-Hack-CVE/CVE-2016-2105)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-2105">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-2105">

---
## CVE-2016-2098 (2016-04-07T23:59:00)
> Action Pack in Ruby on Rails before 3.2.22.2, 4.x before 4.1.14.2, and 4.2.x before 4.2.5.2 allows remote attackers to execute arbitrary Ruby code by leveraging an application's unrestricted use of the render method.
- [Shakun8/CVE-2016-2098](https://github.com/Shakun8/CVE-2016-2098)	<img alt="forks" src="https://img.shields.io/github/forks/Shakun8/CVE-2016-2098">	<img alt="stars" src="https://img.shields.io/github/stars/Shakun8/CVE-2016-2098">
- [j4k0m/CVE-2016-2098](https://github.com/j4k0m/CVE-2016-2098)	<img alt="forks" src="https://img.shields.io/github/forks/j4k0m/CVE-2016-2098">	<img alt="stars" src="https://img.shields.io/github/stars/j4k0m/CVE-2016-2098">
- [Debalinax64/CVE-2016-2098](https://github.com/Debalinax64/CVE-2016-2098)	<img alt="forks" src="https://img.shields.io/github/forks/Debalinax64/CVE-2016-2098">	<img alt="stars" src="https://img.shields.io/github/stars/Debalinax64/CVE-2016-2098">
- [DanielCodex/CVE-2016-2098-my-first-exploit](https://github.com/DanielCodex/CVE-2016-2098-my-first-exploit)	<img alt="forks" src="https://img.shields.io/github/forks/DanielCodex/CVE-2016-2098-my-first-exploit">	<img alt="stars" src="https://img.shields.io/github/stars/DanielCodex/CVE-2016-2098-my-first-exploit">
- [its-arun/CVE-2016-2098](https://github.com/its-arun/CVE-2016-2098)	<img alt="forks" src="https://img.shields.io/github/forks/its-arun/CVE-2016-2098">	<img alt="stars" src="https://img.shields.io/github/stars/its-arun/CVE-2016-2098">
- [3rg1s/CVE-2016-2098](https://github.com/3rg1s/CVE-2016-2098)	<img alt="forks" src="https://img.shields.io/github/forks/3rg1s/CVE-2016-2098">	<img alt="stars" src="https://img.shields.io/github/stars/3rg1s/CVE-2016-2098">
- [0x00-0x00/CVE-2016-2098](https://github.com/0x00-0x00/CVE-2016-2098)	<img alt="forks" src="https://img.shields.io/github/forks/0x00-0x00/CVE-2016-2098">	<img alt="stars" src="https://img.shields.io/github/stars/0x00-0x00/CVE-2016-2098">
- [Alejandro-MartinG/rails-PoC-CVE-2016-2098](https://github.com/Alejandro-MartinG/rails-PoC-CVE-2016-2098)	<img alt="forks" src="https://img.shields.io/github/forks/Alejandro-MartinG/rails-PoC-CVE-2016-2098">	<img alt="stars" src="https://img.shields.io/github/stars/Alejandro-MartinG/rails-PoC-CVE-2016-2098">
- [CyberDefenseInstitute/PoC_CVE-2016-2098_Rails42](https://github.com/CyberDefenseInstitute/PoC_CVE-2016-2098_Rails42)	<img alt="forks" src="https://img.shields.io/github/forks/CyberDefenseInstitute/PoC_CVE-2016-2098_Rails42">	<img alt="stars" src="https://img.shields.io/github/stars/CyberDefenseInstitute/PoC_CVE-2016-2098_Rails42">
- [hderms/dh-CVE_2016_2098](https://github.com/hderms/dh-CVE_2016_2098)	<img alt="forks" src="https://img.shields.io/github/forks/hderms/dh-CVE_2016_2098">	<img alt="stars" src="https://img.shields.io/github/stars/hderms/dh-CVE_2016_2098">

---
## CVE-2016-2004 (2016-04-21T11:00:00)
> HPE Data Protector before 7.03_108, 8.x before 8.15, and 9.x before 9.06 allow remote attackers to execute arbitrary code via unspecified vectors related to lack of authentication.  NOTE: this vulnerability exists because of an incomplete fix for CVE-2014-2623.
- [marcocarolasec/CVE-2016-2004-Exploit](https://github.com/marcocarolasec/CVE-2016-2004-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/marcocarolasec/CVE-2016-2004-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/marcocarolasec/CVE-2016-2004-Exploit">

---
## CVE-2016-20018 (2022-12-19T09:15:00)
> Knex Knex.js through 2.3.0 has a limited SQL injection vulnerability that can be exploited to ignore the WHERE clause of a SQL query.
- [Live-Hack-CVE/CVE-2016-20018](https://github.com/Live-Hack-CVE/CVE-2016-20018)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-20018">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-20018">
- [Live-Hack-CVE/CVE-2016-20018](https://github.com/Live-Hack-CVE/CVE-2016-20018)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-20018">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-20018">

---
## CVE-2016-20017 (2022-10-19T05:15:00)
> D-Link DSL-2750B devices before 1.05 allow remote unauthenticated command injection via the login.cgi cli parameter, as exploited in the wild in 2016 through 2022.
- [Live-Hack-CVE/CVE-2016-20017](https://github.com/Live-Hack-CVE/CVE-2016-20017)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-20017">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-20017">

---
## CVE-2016-20016 (2022-10-19T05:15:00)
> MVPower CCTV DVR models, including TV-7104HE 1.8.4 115215B9 and TV7108HE, contain a web shell that is accessible via a /shell URI. A remote unauthenticated attacker can execute arbitrary operating system commands as root. This vulnerability has also been referred to as the "JAWS webserver RCE" because of the easily identifying HTTP response server field. Other firmware versions, at least from 2014 through 2019, can be affected. This was exploited in the wild in 2017 through 2022.
- [Live-Hack-CVE/CVE-2016-20016](https://github.com/Live-Hack-CVE/CVE-2016-20016)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-20016">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-20016">

---
## CVE-2016-20015 (2022-09-20T18:15:00)
> In the ebuild package through smokeping-2.7.3-r1 for SmokePing on Gentoo, the initscript allows the smokeping user to gain ownership of any file, allowing for the smokeping user to gain root privileges. There is a race condition involving /var/lib/smokeping and chown.
- [Live-Hack-CVE/CVE-2016-20015](https://github.com/Live-Hack-CVE/CVE-2016-20015)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-20015">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-20015">
- [Live-Hack-CVE/CVE-2016-20015](https://github.com/Live-Hack-CVE/CVE-2016-20015)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-20015">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-20015">

---
## CVE-2016-20012 (2021-09-15T20:15:00)
> OpenSSH through 8.7 allows remote attackers, who have a suspicion that a certain combination of username and public key is known to an SSH server, to test whether this suspicion is correct. This occurs because a challenge is sent only when that combination could be valid for a login session. NOTE: the vendor does not recognize user enumeration as a vulnerability for this product
- [aztec-eagle/cve-2016-20012](https://github.com/aztec-eagle/cve-2016-20012)	<img alt="forks" src="https://img.shields.io/github/forks/aztec-eagle/cve-2016-20012">	<img alt="stars" src="https://img.shields.io/github/stars/aztec-eagle/cve-2016-20012">

---
## CVE-2016-1907 (2016-01-19T05:59:00)
> The ssh_packet_read_poll2 function in packet.c in OpenSSH before 7.1p2 allows remote attackers to cause a denial of service (out-of-bounds read and application crash) via crafted network traffic.
- [Live-Hack-CVE/CVE-2016-1907](https://github.com/Live-Hack-CVE/CVE-2016-1907)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1907">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1907">

---
## CVE-2016-1669 (2016-05-14T21:59:00)
> The Zone::New function in zone.cc in Google V8 before 5.0.71.47, as used in Google Chrome before 50.0.2661.102, does not properly determine when to expand certain memory allocations, which allows remote attackers to cause a denial of service (buffer overflow) or possibly have unspecified other impact via crafted JavaScript code.
- [Live-Hack-CVE/CVE-2016-1669](https://github.com/Live-Hack-CVE/CVE-2016-1669)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1669">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1669">

---
## CVE-2016-1531 (2016-04-07T23:59:00)
> Exim before 4.86.2, when installed setuid root, allows local users to gain privileges via the perl_startup argument.
- [k4u5h41/CVE-2016-1531](https://github.com/k4u5h41/CVE-2016-1531)	<img alt="forks" src="https://img.shields.io/github/forks/k4u5h41/CVE-2016-1531">	<img alt="stars" src="https://img.shields.io/github/stars/k4u5h41/CVE-2016-1531">

---
## CVE-2016-15020 (2023-01-16T11:15:00)
> A vulnerability was found in liftkit database up to 2.13.1. It has been classified as critical. This affects the function processOrderBy of the file src/Query/Query.php. The manipulation leads to sql injection. Upgrading to version 2.13.2 is able to address this issue. The name of the patch is 42ec8f2b22e0b0b98fb5b4444ed451c1b21d125a. It is recommended to upgrade the affected component. The associated identifier of this vulnerability is VDB-218391.
- [Live-Hack-CVE/CVE-2016-15020](https://github.com/Live-Hack-CVE/CVE-2016-15020)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15020">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15020">

---
## CVE-2016-15019 (2023-01-15T19:15:00)
> A vulnerability was found in tombh jekbox. It has been rated as problematic. This issue affects some unknown processing of the file lib/server.rb. The manipulation leads to exposure of information through directory listing. The attack may be initiated remotely. The name of the patch is 64eb2677671018fc08b96718b81e3dbc83693190. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-218375.
- [Live-Hack-CVE/CVE-2016-15019](https://github.com/Live-Hack-CVE/CVE-2016-15019)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15019">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15019">

---
## CVE-2016-15018 (2023-01-15T19:15:00)
> A vulnerability was found in krail-jpa up to 0.9.1. It has been classified as critical. This affects an unknown part. The manipulation leads to sql injection. Upgrading to version 0.9.2 is able to address this issue. The name of the patch is c1e848665492e21ef6cc9be443205e36b9a1f6be. It is recommended to upgrade the affected component. The identifier VDB-218373 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2016-15018](https://github.com/Live-Hack-CVE/CVE-2016-15018)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15018">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15018">

---
## CVE-2016-15017 (2023-01-10T15:15:00)
> A vulnerability has been found in fabarea media_upload and classified as critical. This vulnerability affects the function getUploadedFileList of the file Classes/Service/UploadFileService.php. The manipulation leads to pathname traversal. Upgrading to version 0.9.0 is able to address this issue. The name of the patch is b25d42a4981072321c1a363311d8ea2a4ac8763a. It is recommended to upgrade the affected component. VDB-217786 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2016-15017](https://github.com/Live-Hack-CVE/CVE-2016-15017)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15017">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15017">

---
## CVE-2016-15016 (2023-01-08T18:15:00)
> A vulnerability was found in mrtnmtth joomla_mod_einsatz_stats up to 0.2. It has been classified as critical. This affects the function getStatsByType of the file helper.php. The manipulation of the argument year leads to sql injection. Upgrading to version 0.3 is able to address this issue. The name of the patch is 27c1b443cff45c81d9d7d926a74c76f8b6ffc6cb. It is recommended to upgrade the affected component. The identifier VDB-217653 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2016-15016](https://github.com/Live-Hack-CVE/CVE-2016-15016)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15016">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15016">

---
## CVE-2016-15015 (2023-01-08T18:15:00)
> A vulnerability, which was classified as problematic, was found in viafintech Barzahlen Payment Module PHP SDK up to 2.0.0. Affected is the function verify of the file src/Webhook.php. The manipulation leads to observable timing discrepancy. Upgrading to version 2.0.1 is able to address this issue. The name of the patch is 3e7d29dc0ca6c054a6d6e211f32dae89078594c1. It is recommended to upgrade the affected component. VDB-217650 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2016-15015](https://github.com/Live-Hack-CVE/CVE-2016-15015)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15015">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15015">

---
## CVE-2016-15014 (2023-01-07T20:15:00)
> A vulnerability has been found in CESNET theme-cesnet up to 1.x and classified as problematic. Affected by this vulnerability is an unknown functionality of the file cesnet/core/lostpassword/templates/resetpassword.php. The manipulation leads to insufficiently protected credentials. Attacking locally is a requirement. Upgrading to version 2.0.0 is able to address this issue. The name of the patch is 2b857f2233ce5083b4d5bc9bfc4152f933c3e4a6. It is recommended to upgrade the affected component. The identifier VDB-217633 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2016-15014](https://github.com/Live-Hack-CVE/CVE-2016-15014)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15014">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15014">

---
## CVE-2016-15013 (2023-01-07T20:15:00)
> A vulnerability was found in ForumHulp searchresults. It has been rated as critical. Affected by this issue is the function list_keywords of the file event/listener.php. The manipulation of the argument word leads to sql injection. The name of the patch is dd8a312bb285ad9735a8e1da58e9e955837b7322. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217628.
- [Live-Hack-CVE/CVE-2016-15013](https://github.com/Live-Hack-CVE/CVE-2016-15013)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15013">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15013">

---
## CVE-2016-15012 (2023-01-07T13:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability was found in forcedotcom SalesforceMobileSDK-Windows up to 4.x. It has been rated as critical. This issue affects the function ComputeCountSql of the file SalesforceSDK/SmartStore/Store/QuerySpec.cs. The manipulation leads to sql injection. Upgrading to version 5.0.0 is able to address this issue. The name of the patch is 83b3e91e0c1e84873a6d3ca3c5887eb5b4f5a3d8. It is recommended to upgrade the affected component. The associated identifier of this vulnerability is VDB-217619. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2016-15012](https://github.com/Live-Hack-CVE/CVE-2016-15012)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15012">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15012">

---
## CVE-2016-15011 (2023-01-06T10:15:00)
> A vulnerability classified as problematic was found in e-Contract dssp up to 1.3.1. Affected by this vulnerability is the function checkSignResponse of the file dssp-client/src/main/java/be/e_contract/dssp/client/SignResponseVerifier.java. The manipulation leads to xml external entity reference. Upgrading to version 1.3.2 is able to address this issue. The name of the patch is ec4238349691ec66dd30b416ec6eaab02d722302. It is recommended to upgrade the affected component. The identifier VDB-217549 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2016-15011](https://github.com/Live-Hack-CVE/CVE-2016-15011)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15011">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15011">

---
## CVE-2016-15010 (2023-01-05T09:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability classified as problematic was found in University of Cambridge django-ucamlookup up to 1.9.1. Affected by this vulnerability is an unknown functionality of the component Lookup Handler. The manipulation leads to cross site scripting. The attack can be launched remotely. Upgrading to version 1.9.2 is able to address this issue. The name of the patch is 5e25e4765637ea4b9e0bf5fcd5e9a922abee7eb3. It is recommended to upgrade the affected component. The identifier VDB-217441 was assigned to this vulnerability. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2016-15010](https://github.com/Live-Hack-CVE/CVE-2016-15010)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15010">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15010">

---
## CVE-2016-15009 (2023-01-05T09:15:00)
> A vulnerability classified as problematic has been found in OpenACS bug-tracker. Affected is an unknown function of the file lib/nav-bar.adp of the component Search. The manipulation leads to cross-site request forgery. It is possible to launch the attack remotely. The name of the patch is aee43e5714cd8b697355ec3bf83eefee176d3fc3. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217440.
- [Live-Hack-CVE/CVE-2016-15009](https://github.com/Live-Hack-CVE/CVE-2016-15009)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15009">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15009">

---
## CVE-2016-15008 (2023-01-04T10:15:00)
> A vulnerability was found in oxguy3 coebot-www and classified as problematic. This issue affects the function displayChannelCommands/displayChannelQuotes/displayChannelAutoreplies/showChannelHighlights/showChannelBoir of the file js/channel.js. The manipulation leads to cross site scripting. The attack may be initiated remotely. The name of the patch is c1a6c44092585da4236237e0e7da94ee2996a0ca. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-217355.
- [Live-Hack-CVE/CVE-2016-15008](https://github.com/Live-Hack-CVE/CVE-2016-15008)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15008">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15008">

---
## CVE-2016-15007 (2023-01-02T19:15:00)
> A vulnerability was found in Centralized-Salesforce-Dev-Framework. It has been declared as problematic. Affected by this vulnerability is the function SObjectService of the file src/classes/SObjectService.cls of the component SOQL Handler. The manipulation of the argument orderDirection leads to injection. The name of the patch is db03ac5b8a9d830095991b529c067a030a0ccf7b. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-217195.
- [Live-Hack-CVE/CVE-2016-15007](https://github.com/Live-Hack-CVE/CVE-2016-15007)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15007">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15007">

---
## CVE-2016-15006 (2023-01-02T08:15:00)
> A vulnerability, which was classified as problematic, has been found in enigmaX up to 2.2. This issue affects the function getSeed of the file main.c of the component Scrambling Table Handler. The manipulation leads to predictable seed in pseudo-random number generator (prng). The attack may be initiated remotely. Upgrading to version 2.3 is able to address this issue. The name of the patch is 922bf90ca14a681629ba0b807a997a81d70225b5. It is recommended to upgrade the affected component. The identifier VDB-217181 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2016-15006](https://github.com/Live-Hack-CVE/CVE-2016-15006)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15006">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15006">

---
## CVE-2016-15005 (2022-12-27T22:15:00)
> CSRF tokens are generated using math/rand, which is not a cryptographically secure rander number generation, making predicting their values relatively trivial and allowing an attacker to bypass CSRF protections which relatively few requests.
- [Live-Hack-CVE/CVE-2016-15005](https://github.com/Live-Hack-CVE/CVE-2016-15005)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-15005">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-15005">

---
## CVE-2016-12615 ()
> 
- [gk0d/CVE-2016-12615-POC-EXP](https://github.com/gk0d/CVE-2016-12615-POC-EXP)	<img alt="forks" src="https://img.shields.io/github/forks/gk0d/CVE-2016-12615-POC-EXP">	<img alt="stars" src="https://img.shields.io/github/stars/gk0d/CVE-2016-12615-POC-EXP">

---
## CVE-2016-10993 (2019-09-17T15:15:00)
> The ScoreMe theme through 2016-04-01 for WordPress has XSS via the s parameter.
- [varelsecurity/CVE-2016-10993](https://github.com/varelsecurity/CVE-2016-10993)	<img alt="forks" src="https://img.shields.io/github/forks/varelsecurity/CVE-2016-10993">	<img alt="stars" src="https://img.shields.io/github/stars/varelsecurity/CVE-2016-10993">

---
## CVE-2016-10924 (2019-08-22T14:15:00)
> The ebook-download plugin before 1.2 for WordPress has directory traversal.
- [rvizx/CVE-2016-10924](https://github.com/rvizx/CVE-2016-10924)	<img alt="forks" src="https://img.shields.io/github/forks/rvizx/CVE-2016-10924">	<img alt="stars" src="https://img.shields.io/github/stars/rvizx/CVE-2016-10924">
- [LGenAgul/Wordpress-ebook-CVE-2016-10924](https://github.com/LGenAgul/Wordpress-ebook-CVE-2016-10924)	<img alt="forks" src="https://img.shields.io/github/forks/LGenAgul/Wordpress-ebook-CVE-2016-10924">	<img alt="stars" src="https://img.shields.io/github/stars/LGenAgul/Wordpress-ebook-CVE-2016-10924">

---
## CVE-2016-10541 (2018-05-31T20:29:00)
> The npm module "shell-quote" 1.6.0 and earlier cannot correctly escape ">" and "<" operator used for redirection in shell. Applications that depend on shell-quote may also be vulnerable. A malicious user could perform code injection.
- [Live-Hack-CVE/CVE-2016-10541](https://github.com/Live-Hack-CVE/CVE-2016-10541)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10541">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10541">
- [Live-Hack-CVE/CVE-2016-10541](https://github.com/Live-Hack-CVE/CVE-2016-10541)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10541">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10541">

---
## CVE-2016-10229 (2017-04-04T05:59:00)
> udp.c in the Linux kernel before 4.5 allows remote attackers to execute arbitrary code via UDP traffic that triggers an unsafe second checksum calculation during execution of a recv system call with the MSG_PEEK flag.
- [Live-Hack-CVE/CVE-2016-10229](https://github.com/Live-Hack-CVE/CVE-2016-10229)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10229">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10229">

---
## CVE-2016-10228 (2017-03-02T01:59:00)
> The iconv program in the GNU C Library (aka glibc or libc6) 2.31 and earlier, when invoked with multiple suffixes in the destination encoding (TRANSLATE or IGNORE) along with the -c option, enters an infinite loop when processing invalid multi-byte input sequences, leading to a denial of service.
- [Live-Hack-CVE/CVE-2016-10228](https://github.com/Live-Hack-CVE/CVE-2016-10228)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10228">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10228">
- [Live-Hack-CVE/CVE-2016-10228](https://github.com/Live-Hack-CVE/CVE-2016-10228)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10228">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10228">

---
## CVE-2016-10191 (2017-02-09T15:59:00)
> Heap-based buffer overflow in libavformat/rtmppkt.c in FFmpeg before 2.8.10, 3.0.x before 3.0.5, 3.1.x before 3.1.6, and 3.2.x before 3.2.2 allows remote attackers to execute arbitrary code by leveraging failure to check for RTMP packet size mismatches.
- [KaviDk/Heap-Over-Flow-with-CVE-2016-10191](https://github.com/KaviDk/Heap-Over-Flow-with-CVE-2016-10191)	<img alt="forks" src="https://img.shields.io/github/forks/KaviDk/Heap-Over-Flow-with-CVE-2016-10191">	<img alt="stars" src="https://img.shields.io/github/stars/KaviDk/Heap-Over-Flow-with-CVE-2016-10191">

---
## CVE-2016-1019 (2016-04-07T10:59:00)
> Adobe Flash Player 21.0.0.197 and earlier allows remote attackers to cause a denial of service (application crash) or possibly execute arbitrary code via unspecified vectors, as exploited in the wild in April 2016.
- [Live-Hack-CVE/CVE-2016-1019](https://github.com/Live-Hack-CVE/CVE-2016-1019)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1019">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1019">
- [KaviDk/Heap-Over-Flow-with-CVE-2016-10191](https://github.com/KaviDk/Heap-Over-Flow-with-CVE-2016-10191)	<img alt="forks" src="https://img.shields.io/github/forks/KaviDk/Heap-Over-Flow-with-CVE-2016-10191">	<img alt="stars" src="https://img.shields.io/github/stars/KaviDk/Heap-Over-Flow-with-CVE-2016-10191">

---
## CVE-2016-1010 (2016-03-12T15:59:00)
> Integer overflow in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-0963 and CVE-2016-0993.
- [Live-Hack-CVE/CVE-2016-1010](https://github.com/Live-Hack-CVE/CVE-2016-1010)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1010">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1010">
- [Live-Hack-CVE/CVE-2016-1010](https://github.com/Live-Hack-CVE/CVE-2016-1010)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1010">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1010">

---
## CVE-2016-10033 (2016-12-30T19:59:00)
> The mailSend function in the isMail transport in PHPMailer before 5.2.18 might allow remote attackers to pass extra parameters to the mail command and consequently execute arbitrary code via a \" (backslash double quote) in a crafted Sender property.
- [zeeshanbhattined/exploit-CVE-2016-10033](https://github.com/zeeshanbhattined/exploit-CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/zeeshanbhattined/exploit-CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/zeeshanbhattined/exploit-CVE-2016-10033">
- [j4k0m/CVE-2016-10033](https://github.com/j4k0m/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/j4k0m/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/j4k0m/CVE-2016-10033">
- [cved-sources/cve-2016-10033](https://github.com/cved-sources/cve-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/cved-sources/cve-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/cved-sources/cve-2016-10033">
- [opsxcq/exploit-CVE-2016-10033](https://github.com/opsxcq/exploit-CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/opsxcq/exploit-CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/opsxcq/exploit-CVE-2016-10033">
- [0x00-0x00/CVE-2016-10033](https://github.com/0x00-0x00/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/0x00-0x00/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/0x00-0x00/CVE-2016-10033">
- [awidardi/opsxcq-cve-2016-10033](https://github.com/awidardi/opsxcq-cve-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/awidardi/opsxcq-cve-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/awidardi/opsxcq-cve-2016-10033">
- [pedro823/cve-2016-10033-45](https://github.com/pedro823/cve-2016-10033-45)	<img alt="forks" src="https://img.shields.io/github/forks/pedro823/cve-2016-10033-45">	<img alt="stars" src="https://img.shields.io/github/stars/pedro823/cve-2016-10033-45">
- [liusec/WP-CVE-2016-10033](https://github.com/liusec/WP-CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/liusec/WP-CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/liusec/WP-CVE-2016-10033">
- [qwertyuiop12138/CVE-2016-10033](https://github.com/qwertyuiop12138/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/qwertyuiop12138/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/qwertyuiop12138/CVE-2016-10033">
- [chipironcin/CVE-2016-10033](https://github.com/chipironcin/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/chipironcin/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/chipironcin/CVE-2016-10033">
- [Bajunan/CVE-2016-10033](https://github.com/Bajunan/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/Bajunan/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/Bajunan/CVE-2016-10033">
- [GeneralTesler/CVE-2016-10033](https://github.com/GeneralTesler/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/GeneralTesler/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/GeneralTesler/CVE-2016-10033">
- [paralelo14/CVE_2016-10033](https://github.com/paralelo14/CVE_2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/paralelo14/CVE_2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/paralelo14/CVE_2016-10033">
- [Zenexer/safeshell](https://github.com/Zenexer/safeshell)	<img alt="forks" src="https://img.shields.io/github/forks/Zenexer/safeshell">	<img alt="stars" src="https://img.shields.io/github/stars/Zenexer/safeshell">
- [CAOlvchonger/CVE-2016-10033](https://github.com/CAOlvchonger/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/CAOlvchonger/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/CAOlvchonger/CVE-2016-10033">
- [eb613819/CTF_CVE-2016-10033](https://github.com/eb613819/CTF_CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/eb613819/CTF_CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/eb613819/CTF_CVE-2016-10033">
- [ElnurBDa/CVE-2016-10033](https://github.com/ElnurBDa/CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/ElnurBDa/CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/ElnurBDa/CVE-2016-10033">
- [Astrowmist/POC-CVE-2016-10033](https://github.com/Astrowmist/POC-CVE-2016-10033)	<img alt="forks" src="https://img.shields.io/github/forks/Astrowmist/POC-CVE-2016-10033">	<img alt="stars" src="https://img.shields.io/github/stars/Astrowmist/POC-CVE-2016-10033">

---
## CVE-2016-1002 (2016-03-12T15:59:00)
> Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allow attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-0960, CVE-2016-0961, CVE-2016-0962, CVE-2016-0986, CVE-2016-0989, CVE-2016-0992, and CVE-2016-1005.
- [Live-Hack-CVE/CVE-2016-1002](https://github.com/Live-Hack-CVE/CVE-2016-1002)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1002">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1002">
- [Live-Hack-CVE/CVE-2016-1002](https://github.com/Live-Hack-CVE/CVE-2016-1002)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1002">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1002">

---
## CVE-2016-10012 (2017-01-05T02:59:00)
> The shared memory manager (associated with pre-authentication compression) in sshd in OpenSSH before 7.4 does not ensure that a bounds check is enforced by all compilers, which might allows local users to gain privileges by leveraging access to a sandboxed privilege-separation process, related to the m_zback and m_zlib data structures.
- [Live-Hack-CVE/CVE-2016-10012](https://github.com/Live-Hack-CVE/CVE-2016-10012)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10012">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10012">
- [Live-Hack-CVE/CVE-2016-10012](https://github.com/Live-Hack-CVE/CVE-2016-10012)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10012">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10012">

---
## CVE-2016-10011 (2017-01-05T02:59:00)
> authfile.c in sshd in OpenSSH before 7.4 does not properly consider the effects of realloc on buffer contents, which might allow local users to obtain sensitive private-key information by leveraging access to a privilege-separated child process.
- [Live-Hack-CVE/CVE-2016-10011](https://github.com/Live-Hack-CVE/CVE-2016-10011)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10011">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10011">
- [Live-Hack-CVE/CVE-2016-10011](https://github.com/Live-Hack-CVE/CVE-2016-10011)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10011">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10011">

---
## CVE-2016-10010 (2017-01-05T02:59:00)
> sshd in OpenSSH before 7.4, when privilege separation is not used, creates forwarded Unix-domain sockets as root, which might allow local users to gain privileges via unspecified vectors, related to serverloop.c.
- [Live-Hack-CVE/CVE-2016-10010](https://github.com/Live-Hack-CVE/CVE-2016-10010)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10010">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10010">
- [Live-Hack-CVE/CVE-2016-10010](https://github.com/Live-Hack-CVE/CVE-2016-10010)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10010">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10010">

---
## CVE-2016-10009 (2017-01-05T02:59:00)
> Untrusted search path vulnerability in ssh-agent.c in ssh-agent in OpenSSH before 7.4 allows remote attackers to execute arbitrary local PKCS#11 modules by leveraging control over a forwarded agent-socket.
- [Live-Hack-CVE/CVE-2016-10009](https://github.com/Live-Hack-CVE/CVE-2016-10009)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-10009">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-10009">

---
## CVE-2016-1000027 (2020-01-02T23:15:00)
> Pivotal Spring Framework through 5.3.16 suffers from a potential remote code execution (RCE) issue if used for Java deserialization of untrusted data. Depending on how the library is implemented within a product, this issue may or not occur, and authentication may be required. NOTE: the vendor's position is that untrusted data is not an intended use case. The product's behavior will not be changed because some users rely on deserialization of trusted data.
- [Live-Hack-CVE/CVE-2016-1000027](https://github.com/Live-Hack-CVE/CVE-2016-1000027)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-1000027">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-1000027">

---
## CVE-2016-0999 (2016-03-12T15:59:00)
> Use-after-free vulnerability in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-0987, CVE-2016-0988, CVE-2016-0990, CVE-2016-0991, CVE-2016-0994, CVE-2016-0995, CVE-2016-0996, CVE-2016-0997, CVE-2016-0998, and CVE-2016-1000.
- [Live-Hack-CVE/CVE-2016-0999](https://github.com/Live-Hack-CVE/CVE-2016-0999)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0999">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0999">
- [Live-Hack-CVE/CVE-2016-0999](https://github.com/Live-Hack-CVE/CVE-2016-0999)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0999">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0999">

---
## CVE-2016-0996 (2016-03-12T15:59:00)
> Use-after-free vulnerability in the setInterval method in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code via crafted arguments, a different vulnerability than CVE-2016-0987, CVE-2016-0988, CVE-2016-0990, CVE-2016-0991, CVE-2016-0994, CVE-2016-0995, CVE-2016-0997, CVE-2016-0998, CVE-2016-0999, and CVE-2016-1000.
- [Live-Hack-CVE/CVE-2016-0996](https://github.com/Live-Hack-CVE/CVE-2016-0996)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0996">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0996">
- [Live-Hack-CVE/CVE-2016-0996](https://github.com/Live-Hack-CVE/CVE-2016-0996)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0996">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0996">

---
## CVE-2016-0995 (2016-03-12T15:59:00)
> Use-after-free vulnerability in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-0987, CVE-2016-0988, CVE-2016-0990, CVE-2016-0991, CVE-2016-0994, CVE-2016-0996, CVE-2016-0997, CVE-2016-0998, CVE-2016-0999, and CVE-2016-1000.
- [Live-Hack-CVE/CVE-2016-0995](https://github.com/Live-Hack-CVE/CVE-2016-0995)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0995">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0995">
- [Live-Hack-CVE/CVE-2016-0995](https://github.com/Live-Hack-CVE/CVE-2016-0995)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0995">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0995">

---
## CVE-2016-0994 (2016-03-12T15:59:00)
> Use-after-free vulnerability in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code by using the actionCallMethod opcode with crafted arguments, a different vulnerability than CVE-2016-0987, CVE-2016-0988, CVE-2016-0990, CVE-2016-0991, CVE-2016-0995, CVE-2016-0996, CVE-2016-0997, CVE-2016-0998, CVE-2016-0999, and CVE-2016-1000.
- [Live-Hack-CVE/CVE-2016-0994](https://github.com/Live-Hack-CVE/CVE-2016-0994)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0994">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0994">
- [Live-Hack-CVE/CVE-2016-0994](https://github.com/Live-Hack-CVE/CVE-2016-0994)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0994">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0994">

---
## CVE-2016-0993 (2016-03-12T15:59:00)
> Integer overflow in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-0963 and CVE-2016-1010.
- [Live-Hack-CVE/CVE-2016-0993](https://github.com/Live-Hack-CVE/CVE-2016-0993)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0993">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0993">
- [Live-Hack-CVE/CVE-2016-0993](https://github.com/Live-Hack-CVE/CVE-2016-0993)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0993">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0993">

---
## CVE-2016-0989 (2016-03-12T15:59:00)
> Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allow attackers to execute arbitrary code or cause a denial of service (memory corruption) via unspecified vectors, a different vulnerability than CVE-2016-0960, CVE-2016-0961, CVE-2016-0962, CVE-2016-0986, CVE-2016-0992, CVE-2016-1002, and CVE-2016-1005.
- [Live-Hack-CVE/CVE-2016-0989](https://github.com/Live-Hack-CVE/CVE-2016-0989)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0989">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0989">
- [Live-Hack-CVE/CVE-2016-0989](https://github.com/Live-Hack-CVE/CVE-2016-0989)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0989">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0989">

---
## CVE-2016-0987 (2016-03-12T15:59:00)
> Use-after-free vulnerability in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-0988, CVE-2016-0990, CVE-2016-0991, CVE-2016-0994, CVE-2016-0995, CVE-2016-0996, CVE-2016-0997, CVE-2016-0998, CVE-2016-0999, and CVE-2016-1000.
- [Live-Hack-CVE/CVE-2016-0987](https://github.com/Live-Hack-CVE/CVE-2016-0987)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0987">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0987">
- [Live-Hack-CVE/CVE-2016-0987](https://github.com/Live-Hack-CVE/CVE-2016-0987)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0987">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0987">

---
## CVE-2016-0985 (2016-02-10T20:59:00)
> Adobe Flash Player before 18.0.0.329 and 19.x and 20.x before 20.0.0.306 on Windows and OS X and before 11.2.202.569 on Linux, Adobe AIR before 20.0.0.260, Adobe AIR SDK before 20.0.0.260, and Adobe AIR SDK & Compiler before 20.0.0.260 allow attackers to execute arbitrary code by leveraging an unspecified "type confusion."
- [Live-Hack-CVE/CVE-2016-0985](https://github.com/Live-Hack-CVE/CVE-2016-0985)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0985">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0985">

---
## CVE-2016-0963 (2016-03-12T15:59:00)
> Integer overflow in Adobe Flash Player before 18.0.0.333 and 19.x through 21.x before 21.0.0.182 on Windows and OS X and before 11.2.202.577 on Linux, Adobe AIR before 21.0.0.176, Adobe AIR SDK before 21.0.0.176, and Adobe AIR SDK & Compiler before 21.0.0.176 allows attackers to execute arbitrary code via unspecified vectors, a different vulnerability than CVE-2016-0993 and CVE-2016-1010.
- [Live-Hack-CVE/CVE-2016-0963](https://github.com/Live-Hack-CVE/CVE-2016-0963)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0963">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0963">
- [Live-Hack-CVE/CVE-2016-0963](https://github.com/Live-Hack-CVE/CVE-2016-0963)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0963">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0963">

---
## CVE-2016-0800 (2016-03-01T20:59:00)
> The SSLv2 protocol, as used in OpenSSL before 1.0.1s and 1.0.2 before 1.0.2g and other products, requires a server to send a ServerVerify message before establishing that a client possesses certain plaintext RSA data, which makes it easier for remote attackers to decrypt TLS ciphertext data by leveraging a Bleichenbacher RSA padding oracle, aka a "DROWN" attack.
- [Live-Hack-CVE/CVE-2016-0800](https://github.com/Live-Hack-CVE/CVE-2016-0800)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0800">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0800">
- [Live-Hack-CVE/CVE-2016-0800](https://github.com/Live-Hack-CVE/CVE-2016-0800)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0800">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0800">
- [clic-kbait/A2SV--SSL-VUL-Scan](https://github.com/clic-kbait/A2SV--SSL-VUL-Scan)	<img alt="forks" src="https://img.shields.io/github/forks/clic-kbait/A2SV--SSL-VUL-Scan">	<img alt="stars" src="https://img.shields.io/github/stars/clic-kbait/A2SV--SSL-VUL-Scan">

---
## CVE-2016-0799 (2016-03-03T20:59:00)
> The fmtstr function in crypto/bio/b_print.c in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g improperly calculates string lengths, which allows remote attackers to cause a denial of service (overflow and out-of-bounds read) or possibly have unspecified other impact via a long string, as demonstrated by a large amount of ASN.1 data, a different vulnerability than CVE-2016-2842.
- [Live-Hack-CVE/CVE-2016-0799](https://github.com/Live-Hack-CVE/CVE-2016-0799)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0799">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0799">

---
## CVE-2016-0798 (2016-03-03T20:59:00)
> Memory leak in the SRP_VBASE_get_by_user implementation in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g allows remote attackers to cause a denial of service (memory consumption) by providing an invalid username in a connection attempt, related to apps/s_server.c and crypto/srp/srp_vfy.c.
- [Live-Hack-CVE/CVE-2016-0798](https://github.com/Live-Hack-CVE/CVE-2016-0798)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0798">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0798">
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-0798](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0798)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0798">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0798">

---
## CVE-2016-0797 (2016-03-03T20:59:00)
> Multiple integer overflows in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g allow remote attackers to cause a denial of service (heap memory corruption or NULL pointer dereference) or possibly have unspecified other impact via a long digit string that is mishandled by the (1) BN_dec2bn or (2) BN_hex2bn function, related to crypto/bn/bn.h and crypto/bn/bn_print.c.
- [Live-Hack-CVE/CVE-2016-0797](https://github.com/Live-Hack-CVE/CVE-2016-0797)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0797">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0797">
- [Live-Hack-CVE/CVE-2016-0797](https://github.com/Live-Hack-CVE/CVE-2016-0797)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0797">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0797">
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797">
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0797">

---
## CVE-2016-0792 (2016-04-07T23:59:00)
> Multiple unspecified API endpoints in Jenkins before 1.650 and LTS before 1.642.2 allow remote authenticated users to execute arbitrary code via serialized data in an XML file, related to XStream and groovy.util.Expando.
- [Aviksaikat/CVE-2016-0792](https://github.com/Aviksaikat/CVE-2016-0792)	<img alt="forks" src="https://img.shields.io/github/forks/Aviksaikat/CVE-2016-0792">	<img alt="stars" src="https://img.shields.io/github/stars/Aviksaikat/CVE-2016-0792">
- [R0B1NL1N/java-deserialization-exploits](https://github.com/R0B1NL1N/java-deserialization-exploits)	<img alt="forks" src="https://img.shields.io/github/forks/R0B1NL1N/java-deserialization-exploits">	<img alt="stars" src="https://img.shields.io/github/stars/R0B1NL1N/java-deserialization-exploits">
- [jpiechowka/jenkins-cve-2016-0792](https://github.com/jpiechowka/jenkins-cve-2016-0792)	<img alt="forks" src="https://img.shields.io/github/forks/jpiechowka/jenkins-cve-2016-0792">	<img alt="stars" src="https://img.shields.io/github/stars/jpiechowka/jenkins-cve-2016-0792">

---
## CVE-2016-0778 (2016-01-14T22:59:00)
> The (1) roaming_read and (2) roaming_write functions in roaming_common.c in the client in OpenSSH 5.x, 6.x, and 7.x before 7.1p2, when certain proxy and forward options are enabled, do not properly maintain connection file descriptors, which allows remote servers to cause a denial of service (heap-based buffer overflow) or possibly have unspecified other impact by requesting many forwardings.
- [Live-Hack-CVE/CVE-2016-0778](https://github.com/Live-Hack-CVE/CVE-2016-0778)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0778">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0778">
- [Live-Hack-CVE/CVE-2016-0778](https://github.com/Live-Hack-CVE/CVE-2016-0778)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0778">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0778">

---
## CVE-2016-0777 (2016-01-14T22:59:00)
> The resend_bytes function in roaming_common.c in the client in OpenSSH 5.x, 6.x, and 7.x before 7.1p2 allows remote servers to obtain sensitive information from process memory by requesting transmission of an entire buffer, as demonstrated by reading a private key.
- [Live-Hack-CVE/CVE-2016-0777](https://github.com/Live-Hack-CVE/CVE-2016-0777)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0777">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0777">
- [Live-Hack-CVE/CVE-2016-0777](https://github.com/Live-Hack-CVE/CVE-2016-0777)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0777">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0777">

---
## CVE-2016-0728 (2016-02-08T03:59:00)
> The join_session_keyring function in security/keys/process_keys.c in the Linux kernel before 4.4.1 mishandles object references in a certain error case, which allows local users to gain privileges or cause a denial of service (integer overflow and use-after-free) via crafted keyctl commands. <a href="http://cwe.mitre.org/data/definitions/190.html">CWE-190: Integer Overflow or Wraparound</a> <br />

<a href="http://cwe.mitre.org/data/definitions/416.html">CWE-416: Use After Free</a>
- [tndud042713/cve](https://github.com/tndud042713/cve)	<img alt="forks" src="https://img.shields.io/github/forks/tndud042713/cve">	<img alt="stars" src="https://img.shields.io/github/stars/tndud042713/cve">
- [th30d00r/Linux-Vulnerability-CVE-2016-0728-and-Exploit](https://github.com/th30d00r/Linux-Vulnerability-CVE-2016-0728-and-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/th30d00r/Linux-Vulnerability-CVE-2016-0728-and-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/th30d00r/Linux-Vulnerability-CVE-2016-0728-and-Exploit">
- [nardholio/cve-2016-0728](https://github.com/nardholio/cve-2016-0728)	<img alt="forks" src="https://img.shields.io/github/forks/nardholio/cve-2016-0728">	<img alt="stars" src="https://img.shields.io/github/stars/nardholio/cve-2016-0728">
- [sugarvillela/CVE](https://github.com/sugarvillela/CVE)	<img alt="forks" src="https://img.shields.io/github/forks/sugarvillela/CVE">	<img alt="stars" src="https://img.shields.io/github/stars/sugarvillela/CVE">
- [hal0taso/CVE-2016-0728](https://github.com/hal0taso/CVE-2016-0728)	<img alt="forks" src="https://img.shields.io/github/forks/hal0taso/CVE-2016-0728">	<img alt="stars" src="https://img.shields.io/github/stars/hal0taso/CVE-2016-0728">
- [sibilleg/exploit_cve-2016-0728](https://github.com/sibilleg/exploit_cve-2016-0728)	<img alt="forks" src="https://img.shields.io/github/forks/sibilleg/exploit_cve-2016-0728">	<img alt="stars" src="https://img.shields.io/github/stars/sibilleg/exploit_cve-2016-0728">
- [bittorrent3389/cve-2016-0728](https://github.com/bittorrent3389/cve-2016-0728)	<img alt="forks" src="https://img.shields.io/github/forks/bittorrent3389/cve-2016-0728">	<img alt="stars" src="https://img.shields.io/github/stars/bittorrent3389/cve-2016-0728">
- [neuschaefer/cve-2016-0728-testbed](https://github.com/neuschaefer/cve-2016-0728-testbed)	<img alt="forks" src="https://img.shields.io/github/forks/neuschaefer/cve-2016-0728-testbed">	<img alt="stars" src="https://img.shields.io/github/stars/neuschaefer/cve-2016-0728-testbed">
- [fochess/cve_2016_0728](https://github.com/fochess/cve_2016_0728)	<img alt="forks" src="https://img.shields.io/github/forks/fochess/cve_2016_0728">	<img alt="stars" src="https://img.shields.io/github/stars/fochess/cve_2016_0728">
- [mfer/cve_2016_0728](https://github.com/mfer/cve_2016_0728)	<img alt="forks" src="https://img.shields.io/github/forks/mfer/cve_2016_0728">	<img alt="stars" src="https://img.shields.io/github/stars/mfer/cve_2016_0728">
- [sunnyjiang/cve_2016_0728](https://github.com/sunnyjiang/cve_2016_0728)	<img alt="forks" src="https://img.shields.io/github/forks/sunnyjiang/cve_2016_0728">	<img alt="stars" src="https://img.shields.io/github/stars/sunnyjiang/cve_2016_0728">
- [googleweb/CVE-2016-0728](https://github.com/googleweb/CVE-2016-0728)	<img alt="forks" src="https://img.shields.io/github/forks/googleweb/CVE-2016-0728">	<img alt="stars" src="https://img.shields.io/github/stars/googleweb/CVE-2016-0728">
- [bjzz/cve_2016_0728_exploit](https://github.com/bjzz/cve_2016_0728_exploit)	<img alt="forks" src="https://img.shields.io/github/forks/bjzz/cve_2016_0728_exploit">	<img alt="stars" src="https://img.shields.io/github/stars/bjzz/cve_2016_0728_exploit">
- [kennetham/cve_2016_0728](https://github.com/kennetham/cve_2016_0728)	<img alt="forks" src="https://img.shields.io/github/forks/kennetham/cve_2016_0728">	<img alt="stars" src="https://img.shields.io/github/stars/kennetham/cve_2016_0728">
- [isnuryusuf/cve_2016_0728](https://github.com/isnuryusuf/cve_2016_0728)	<img alt="forks" src="https://img.shields.io/github/forks/isnuryusuf/cve_2016_0728">	<img alt="stars" src="https://img.shields.io/github/stars/isnuryusuf/cve_2016_0728">
- [idl3r/cve-2016-0728](https://github.com/idl3r/cve-2016-0728)	<img alt="forks" src="https://img.shields.io/github/forks/idl3r/cve-2016-0728">	<img alt="stars" src="https://img.shields.io/github/stars/idl3r/cve-2016-0728">
- [sidrk01/cve-2016-0728](https://github.com/sidrk01/cve-2016-0728)	<img alt="forks" src="https://img.shields.io/github/forks/sidrk01/cve-2016-0728">	<img alt="stars" src="https://img.shields.io/github/stars/sidrk01/cve-2016-0728">

---
## CVE-2016-0705 (2016-03-03T20:59:00)
> Double free vulnerability in the dsa_priv_decode function in crypto/dsa/dsa_ameth.c in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g allows remote attackers to cause a denial of service (memory corruption) or possibly have unspecified other impact via a malformed DSA private key.
- [Live-Hack-CVE/CVE-2016-0705](https://github.com/Live-Hack-CVE/CVE-2016-0705)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0705">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0705">
- [nidhi7598/OPENSSL_1.0.1g_CVE-2016-0705](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0705)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0705">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2016-0705">
- [hshivhare67/OpenSSL_1.0.1g_CVE-2016-0705](https://github.com/hshivhare67/OpenSSL_1.0.1g_CVE-2016-0705)	<img alt="forks" src="https://img.shields.io/github/forks/hshivhare67/OpenSSL_1.0.1g_CVE-2016-0705">	<img alt="stars" src="https://img.shields.io/github/stars/hshivhare67/OpenSSL_1.0.1g_CVE-2016-0705">

---
## CVE-2016-0703 (2016-03-02T11:59:00)
> The get_client_master_key function in s2_srvr.c in the SSLv2 implementation in OpenSSL before 0.9.8zf, 1.0.0 before 1.0.0r, 1.0.1 before 1.0.1m, and 1.0.2 before 1.0.2a accepts a nonzero CLIENT-MASTER-KEY CLEAR-KEY-LENGTH value for an arbitrary cipher, which allows man-in-the-middle attackers to determine the MASTER-KEY value and decrypt TLS ciphertext data by leveraging a Bleichenbacher RSA padding oracle, a related issue to CVE-2016-0800.
- [Live-Hack-CVE/CVE-2016-0703](https://github.com/Live-Hack-CVE/CVE-2016-0703)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0703">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0703">

---
## CVE-2016-0702 (2016-03-03T20:59:00)
> The MOD_EXP_CTIME_COPY_FROM_PREBUF function in crypto/bn/bn_exp.c in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g does not properly consider cache-bank access times during modular exponentiation, which makes it easier for local users to discover RSA keys by running a crafted application on the same Intel Sandy Bridge CPU core as a victim and leveraging cache-bank conflicts, aka a "CacheBleed" attack.
- [Live-Hack-CVE/CVE-2016-0702](https://github.com/Live-Hack-CVE/CVE-2016-0702)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2016-0702">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2016-0702">
- [Trinadh465/OpenSSL-1_0_1g_CVE-2016-0702](https://github.com/Trinadh465/OpenSSL-1_0_1g_CVE-2016-0702)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/OpenSSL-1_0_1g_CVE-2016-0702">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/OpenSSL-1_0_1g_CVE-2016-0702">

---
## CVE-2016-0451 (2016-01-21T03:00:00)
> Unspecified vulnerability in the Oracle GoldenGate component in Oracle GoldenGate 11.2 and 12.1.2 allows remote attackers to affect confidentiality, integrity, and availability via unknown vectors, a different vulnerability than CVE-2016-0452. Per Oracle:  The CVSS score is 10.0 only on Windows for Database versions prior to 12c. The CVSS is 7.5 (Confidentiality, Integrity and Availability is "Partial+") for Database 12c on Windows and for all versions of Database on Linux, Unix and other platforms.
- [rwincey/Oracle-GoldenGate---CVE-2016-0451](https://github.com/rwincey/Oracle-GoldenGate---CVE-2016-0451)	<img alt="forks" src="https://img.shields.io/github/forks/rwincey/Oracle-GoldenGate---CVE-2016-0451">	<img alt="stars" src="https://img.shields.io/github/stars/rwincey/Oracle-GoldenGate---CVE-2016-0451">

---
## CVE-2016-0010 (2016-01-13T05:59:00)
> Microsoft Office 2007 SP3, Office 2010 SP2, Office 2013 SP1, Office 2013 RT SP1, Office 2016, Excel for Mac 2011, PowerPoint for Mac 2011, Word for Mac 2011, Excel 2016 for Mac, PowerPoint 2016 for Mac, Word 2016 for Mac, and Word Viewer allow remote attackers to execute arbitrary code via a crafted Office document, aka "Microsoft Office Memory Corruption Vulnerability."
- [Sunqiz/CVE-2016-0010-reproduction](https://github.com/Sunqiz/CVE-2016-0010-reproduction)	<img alt="forks" src="https://img.shields.io/github/forks/Sunqiz/CVE-2016-0010-reproduction">	<img alt="stars" src="https://img.shields.io/github/stars/Sunqiz/CVE-2016-0010-reproduction">
