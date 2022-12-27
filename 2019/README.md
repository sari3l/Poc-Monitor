# 2019 List

---
## CVE-2019-9978 (2019-03-24T15:29:00)
> The social-warfare plugin before 3.5.3 for WordPress has stored XSS via the wp-admin/admin-post.php?swp_debug=load_options swp_url parameter, as exploited in the wild in March 2019. This affects Social Warfare and Social Warfare Pro.
- [caique-garbim/CVE-2019-9978_Exploit](https://github.com/caique-garbim/CVE-2019-9978_Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/caique-garbim/CVE-2019-9978_Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/caique-garbim/CVE-2019-9978_Exploit">
- [hash3liZer/CVE-2019-9978](https://github.com/hash3liZer/CVE-2019-9978)	<img alt="forks" src="https://img.shields.io/github/forks/hash3liZer/CVE-2019-9978">	<img alt="stars" src="https://img.shields.io/github/stars/hash3liZer/CVE-2019-9978">
- [cved-sources/cve-2019-9978](https://github.com/cved-sources/cve-2019-9978)	<img alt="forks" src="https://img.shields.io/github/forks/cved-sources/cve-2019-9978">	<img alt="stars" src="https://img.shields.io/github/stars/cved-sources/cve-2019-9978">
- [mpgn/CVE-2019-9978](https://github.com/mpgn/CVE-2019-9978)	<img alt="forks" src="https://img.shields.io/github/forks/mpgn/CVE-2019-9978">	<img alt="stars" src="https://img.shields.io/github/stars/mpgn/CVE-2019-9978">
- [KTN1990/CVE-2019-9978](https://github.com/KTN1990/CVE-2019-9978)	<img alt="forks" src="https://img.shields.io/github/forks/KTN1990/CVE-2019-9978">	<img alt="stars" src="https://img.shields.io/github/stars/KTN1990/CVE-2019-9978">

---
## CVE-2019-9947 (2019-03-23T18:29:00)
> An issue was discovered in urllib2 in Python 2.x through 2.7.16 and urllib in Python 3.x through 3.7.3. CRLF injection is possible if the attacker controls a url parameter, as demonstrated by the first argument to urllib.request.urlopen with \r\n (specifically in the path component of a URL that lacks a ? character) followed by an HTTP header or a Redis command. This is similar to the CVE-2019-9740 query string issue. This is fixed in: v2.7.17, v2.7.17rc1, v2.7.18, v2.7.18rc1; v3.5.10, v3.5.10rc1, v3.5.8, v3.5.8rc1, v3.5.8rc2, v3.5.9; v3.6.10, v3.6.10rc1, v3.6.11, v3.6.11rc1, v3.6.12, v3.6.9, v3.6.9rc1; v3.7.4, v3.7.4rc1, v3.7.4rc2, v3.7.5, v3.7.5rc1, v3.7.6, v3.7.6rc1, v3.7.7, v3.7.7rc1, v3.7.8, v3.7.8rc1, v3.7.9.
- [Live-Hack-CVE/CVE-2019-9947](https://github.com/Live-Hack-CVE/CVE-2019-9947)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9947">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9947">

---
## CVE-2019-9855 (2019-09-06T19:15:00)
> LibreOffice is typically bundled with LibreLogo, a programmable turtle vector graphics script, which can execute arbitrary python commands contained with the document it is launched from. LibreOffice also has a feature where documents can specify that pre-installed scripts can be executed on various document script events such as mouse-over, etc. Protection was added to block calling LibreLogo from script event handers. However a Windows 8.3 path equivalence handling flaw left LibreOffice vulnerable under Windows that a document could trigger executing LibreLogo via a Windows filename pseudonym. This issue affects: Document Foundation LibreOffice 6.2 versions prior to 6.2.7; 6.3 versions prior to 6.3.1.
- [Live-Hack-CVE/CVE-2019-9855](https://github.com/Live-Hack-CVE/CVE-2019-9855)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9855">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9855">

---
## CVE-2019-9851 (2019-08-15T22:15:00)
> LibreOffice is typically bundled with LibreLogo, a programmable turtle vector graphics script, which can execute arbitrary python commands contained with the document it is launched from. Protection was added, to address CVE-2019-9848, to block calling LibreLogo from document event script handers, e.g. mouse over. However LibreOffice also has a separate feature where documents can specify that pre-installed scripts can be executed on various global script events such as document-open, etc. In the fixed versions, global script event handlers are validated equivalently to document script event handlers. This issue affects: Document Foundation LibreOffice versions prior to 6.2.6.
- [Live-Hack-CVE/CVE-2019-9851](https://github.com/Live-Hack-CVE/CVE-2019-9851)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9851">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9851">

---
## CVE-2019-9850 (2019-08-15T22:15:00)
> LibreOffice is typically bundled with LibreLogo, a programmable turtle vector graphics script, which can execute arbitrary python commands contained with the document it is launched from. LibreOffice also has a feature where documents can specify that pre-installed scripts can be executed on various document script events such as mouse-over, etc. Protection was added, to address CVE-2019-9848, to block calling LibreLogo from script event handers. However an insufficient url validation vulnerability in LibreOffice allowed malicious to bypass that protection and again trigger calling LibreLogo from script event handlers. This issue affects: Document Foundation LibreOffice versions prior to 6.2.6.
- [Live-Hack-CVE/CVE-2019-9850](https://github.com/Live-Hack-CVE/CVE-2019-9850)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9850">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9850">

---
## CVE-2019-9766 (2019-03-14T09:29:00)
> Stack-based buffer overflow in Free MP3 CD Ripper 2.6, when converting a file, allows user-assisted remote attackers to execute arbitrary code via a crafted .mp3 file.
- [zeronohacker/CVE-2019-9766](https://github.com/zeronohacker/CVE-2019-9766)	<img alt="forks" src="https://img.shields.io/github/forks/zeronohacker/CVE-2019-9766">	<img alt="stars" src="https://img.shields.io/github/stars/zeronohacker/CVE-2019-9766">
- [moonheadobj/CVE-2019-9766](https://github.com/moonheadobj/CVE-2019-9766)	<img alt="forks" src="https://img.shields.io/github/forks/moonheadobj/CVE-2019-9766">	<img alt="stars" src="https://img.shields.io/github/stars/moonheadobj/CVE-2019-9766">

---
## CVE-2019-9740 (2019-03-13T03:29:00)
> An issue was discovered in urllib2 in Python 2.x through 2.7.16 and urllib in Python 3.x through 3.7.3. CRLF injection is possible if the attacker controls a url parameter, as demonstrated by the first argument to urllib.request.urlopen with \r\n (specifically in the query string after a ? character) followed by an HTTP header or a Redis command. This is fixed in: v2.7.17, v2.7.17rc1, v2.7.18, v2.7.18rc1; v3.5.10, v3.5.10rc1, v3.5.8, v3.5.8rc1, v3.5.8rc2, v3.5.9; v3.6.10, v3.6.10rc1, v3.6.11, v3.6.11rc1, v3.6.12, v3.6.9, v3.6.9rc1; v3.7.4, v3.7.4rc1, v3.7.4rc2, v3.7.5, v3.7.5rc1, v3.7.6, v3.7.6rc1, v3.7.7, v3.7.7rc1, v3.7.8, v3.7.8rc1, v3.7.9.
- [Live-Hack-CVE/CVE-2019-9740](https://github.com/Live-Hack-CVE/CVE-2019-9740)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9740">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9740">

---
## CVE-2019-9721 (2019-03-12T09:29:00)
> A denial of service in the subtitle decoder in FFmpeg 3.2 and 4.1 allows attackers to hog the CPU via a crafted video file in Matroska format, because handle_open_brace in libavcodec/htmlsubtitles.c has a complex format argument to sscanf.
- [Live-Hack-CVE/CVE-2019-9721](https://github.com/Live-Hack-CVE/CVE-2019-9721)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9721">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9721">

---
## CVE-2019-9634 (2019-03-08T15:29:00)
> Go through 1.12 on Windows misuses certain LoadLibrary functionality, leading to DLL injection.
- [Live-Hack-CVE/CVE-2019-9634](https://github.com/Live-Hack-CVE/CVE-2019-9634)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9634">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9634">

---
## CVE-2019-9592 (2019-03-06T16:29:00)
> A reflected Cross-site scripting (XSS) vulnerability in ShoreTel Connect ONSITE 19.45.1602.0 allows remote attackers to inject arbitrary web script or HTML via the url parameter.
- [Live-Hack-CVE/CVE-2019-9592](https://github.com/Live-Hack-CVE/CVE-2019-9592)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9592">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9592">

---
## CVE-2019-9591 (2019-03-06T16:29:00)
> A reflected Cross-site scripting (XSS) vulnerability in ShoreTel Connect ONSITE before 19.49.1500.0 allows remote attackers to inject arbitrary web script or HTML via the brandUrl parameter.
- [Live-Hack-CVE/CVE-2019-9591](https://github.com/Live-Hack-CVE/CVE-2019-9591)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9591">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9591">

---
## CVE-2019-9193 (2019-04-01T21:30:00)
> ** DISPUTED ** In PostgreSQL 9.3 through 11.2, the "COPY TO/FROM PROGRAM" function allows superusers and users in the 'pg_execute_server_program' group to execute arbitrary code in the context of the database's operating system user. This functionality is enabled by default and can be abused to run arbitrary operating system commands on Windows, Linux, and macOS. NOTE: Third parties claim/state this is not an issue because PostgreSQL functionality for ‘COPY TO/FROM PROGRAM’ is acting as intended. References state that in PostgreSQL, a superuser can execute commands as the server user without using the ‘COPY FROM PROGRAM’.
- [chromanite/CVE-2019-9193-PostgreSQL-9.3-11.7](https://github.com/chromanite/CVE-2019-9193-PostgreSQL-9.3-11.7)	<img alt="forks" src="https://img.shields.io/github/forks/chromanite/CVE-2019-9193-PostgreSQL-9.3-11.7">	<img alt="stars" src="https://img.shields.io/github/stars/chromanite/CVE-2019-9193-PostgreSQL-9.3-11.7">
- [b4keSn4ke/CVE-2019-9193](https://github.com/b4keSn4ke/CVE-2019-9193)	<img alt="forks" src="https://img.shields.io/github/forks/b4keSn4ke/CVE-2019-9193">	<img alt="stars" src="https://img.shields.io/github/stars/b4keSn4ke/CVE-2019-9193">
- [wkjung0624/cve-2019-9193](https://github.com/wkjung0624/cve-2019-9193)	<img alt="forks" src="https://img.shields.io/github/forks/wkjung0624/cve-2019-9193">	<img alt="stars" src="https://img.shields.io/github/stars/wkjung0624/cve-2019-9193">

---
## CVE-2019-9139 (2019-04-25T18:29:00)
> DaviewIndy 8.98.7 and earlier versions have a Integer overflow vulnerability, triggered when the user opens a malformed PDF file that is mishandled by Daview.exe. Attackers could exploit this and arbitrary code execution.
- [Live-Hack-CVE/CVE-2019-9139](https://github.com/Live-Hack-CVE/CVE-2019-9139)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-9139">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-9139">

---
## CVE-2019-9053 (2019-03-26T17:29:00)
> An issue was discovered in CMS Made Simple 2.2.8. It is possible with the News module, through a crafted URL, to achieve unauthenticated blind time-based SQL injection via the m1_idlist parameter.
- [ELIZEUOPAIN/CVE-2019-9053-CMS-Made-Simple-2.2.10---SQL-Injection-Exploit](https://github.com/ELIZEUOPAIN/CVE-2019-9053-CMS-Made-Simple-2.2.10---SQL-Injection-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/ELIZEUOPAIN/CVE-2019-9053-CMS-Made-Simple-2.2.10---SQL-Injection-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/ELIZEUOPAIN/CVE-2019-9053-CMS-Made-Simple-2.2.10---SQL-Injection-Exploit">
- [zmiddle/Simple_CMS_SQLi](https://github.com/zmiddle/Simple_CMS_SQLi)	<img alt="forks" src="https://img.shields.io/github/forks/zmiddle/Simple_CMS_SQLi">	<img alt="stars" src="https://img.shields.io/github/stars/zmiddle/Simple_CMS_SQLi">
- [k4u5h41/CVE-2019-9053](https://github.com/k4u5h41/CVE-2019-9053)	<img alt="forks" src="https://img.shields.io/github/forks/k4u5h41/CVE-2019-9053">	<img alt="stars" src="https://img.shields.io/github/stars/k4u5h41/CVE-2019-9053">
- [Matthsh/SQLi-correction](https://github.com/Matthsh/SQLi-correction)	<img alt="forks" src="https://img.shields.io/github/forks/Matthsh/SQLi-correction">	<img alt="stars" src="https://img.shields.io/github/stars/Matthsh/SQLi-correction">
- [e-renna/CVE-2019-9053](https://github.com/e-renna/CVE-2019-9053)	<img alt="forks" src="https://img.shields.io/github/forks/e-renna/CVE-2019-9053">	<img alt="stars" src="https://img.shields.io/github/stars/e-renna/CVE-2019-9053">
- [maraspiras/46635.py](https://github.com/maraspiras/46635.py)	<img alt="forks" src="https://img.shields.io/github/forks/maraspiras/46635.py">	<img alt="stars" src="https://img.shields.io/github/stars/maraspiras/46635.py">
- [padsalatushal/CVE-2019-9053](https://github.com/padsalatushal/CVE-2019-9053)	<img alt="forks" src="https://img.shields.io/github/forks/padsalatushal/CVE-2019-9053">	<img alt="stars" src="https://img.shields.io/github/stars/padsalatushal/CVE-2019-9053">
- [SUNNYSAINI01001/46635.py_CVE-2019-9053](https://github.com/SUNNYSAINI01001/46635.py_CVE-2019-9053)	<img alt="forks" src="https://img.shields.io/github/forks/SUNNYSAINI01001/46635.py_CVE-2019-9053">	<img alt="stars" src="https://img.shields.io/github/stars/SUNNYSAINI01001/46635.py_CVE-2019-9053">
- [pedrojosenavasperez/CVE-2019-9053-Python3](https://github.com/pedrojosenavasperez/CVE-2019-9053-Python3)	<img alt="forks" src="https://img.shields.io/github/forks/pedrojosenavasperez/CVE-2019-9053-Python3">	<img alt="stars" src="https://img.shields.io/github/stars/pedrojosenavasperez/CVE-2019-9053-Python3">

---
## CVE-2019-8985 (2019-02-21T19:29:00)
> On Netis WF2411 with firmware 2.1.36123 and other Netis WF2xxx devices (possibly WF2411 through WF2880), there is a stack-based buffer overflow that does not require authentication. This can cause denial of service (device restart) or remote code execution. This vulnerability can be triggered by a GET request with a long HTTP "Authorization: Basic" header that is mishandled by user_auth->user_ok in /bin/boa.
- [Ler2sq/CVE-2019-8985](https://github.com/Ler2sq/CVE-2019-8985)	<img alt="forks" src="https://img.shields.io/github/forks/Ler2sq/CVE-2019-8985">	<img alt="stars" src="https://img.shields.io/github/stars/Ler2sq/CVE-2019-8985">

---
## CVE-2019-8292 (2019-10-01T20:15:00)
> Online Store System v1.0 delete_product.php doesn't check to see if a user authtenticated or has administrative rights allowing arbitrary product deletion.
- [Live-Hack-CVE/CVE-2019-8292](https://github.com/Live-Hack-CVE/CVE-2019-8292)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-8292">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-8292">

---
## CVE-2019-7672 (2019-06-05T19:29:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. The flash version of the web interface contains a hard-coded username and password, which may allow an authenticated attacker to escalate privileges.
- [Live-Hack-CVE/CVE-2019-7672](https://github.com/Live-Hack-CVE/CVE-2019-7672)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7672">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7672">

---
## CVE-2019-7671 (2019-06-05T19:29:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. Parameters sent to scripts are not properly sanitized before being returned to the user, which may allow an attacker to execute arbitrary code in a user’s browser session in context of an affected site.
- [Live-Hack-CVE/CVE-2019-7671](https://github.com/Live-Hack-CVE/CVE-2019-7671)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7671">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7671">

---
## CVE-2019-7670 (2019-07-01T19:15:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. The application incorrectly neutralizes special elements that could modify the intended OS command when it is sent to a downstream component, which could allow attackers to execute commands directly on the operating system.
- [Live-Hack-CVE/CVE-2019-7670](https://github.com/Live-Hack-CVE/CVE-2019-7670)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7670">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7670">

---
## CVE-2019-7669 (2019-07-01T19:15:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. Improper validation of file extensions when uploading files could allow a remote authenticated attacker to upload and execute malicious applications within the application’s web root with root privileges.
- [Live-Hack-CVE/CVE-2019-7669](https://github.com/Live-Hack-CVE/CVE-2019-7669)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7669">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7669">

---
## CVE-2019-7667 (2019-07-01T19:15:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. The application generates database backup files with a predictable name, and an attacker can use brute force to identify the database backup file name. A malicious actor can exploit this issue to download the database file and disclose login information, which can allow the attacker to bypass authentication and have full access to the system.
- [Live-Hack-CVE/CVE-2019-7667](https://github.com/Live-Hack-CVE/CVE-2019-7667)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7667">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7667">

---
## CVE-2019-7666 (2019-07-01T19:15:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. The application allows improper authentication using the MD5 hash value of the password, which may allow an attacker with access to the database to login as admin without decrypting the password.
- [Live-Hack-CVE/CVE-2019-7666](https://github.com/Live-Hack-CVE/CVE-2019-7666)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7666">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7666">

---
## CVE-2019-7655 (2020-01-29T16:15:00)
> Wowza Streaming Engine 4.8.0 and earlier from multiple authenticated XSS vulnerabilities via the (1) customList%5B0%5D.value field in enginemanager/server/serversetup/edit_adv.htm of the Server Setup configuration or the (2) host field in enginemanager/j_spring_security_check of the login form. This issue was resolved in Wowza Streaming Engine 4.8.5.
- [Live-Hack-CVE/CVE-2019-7655](https://github.com/Live-Hack-CVE/CVE-2019-7655)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7655">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7655">

---
## CVE-2019-7654 (2020-01-29T16:15:00)
> Wowza Streaming Engine 4.8.0 and earlier suffers from multiple CSRF vulnerabilities. For example, an administrator, by following a link, can be tricked into making unwanted changes such as adding another admin user via enginemanager/server/user/edit.htm in the Server->Users component. This issue was resolved in Wowza Streaming Engine 4.8.5.
- [Live-Hack-CVE/CVE-2019-7654](https://github.com/Live-Hack-CVE/CVE-2019-7654)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7654">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7654">

---
## CVE-2019-7281 (2019-07-01T19:15:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. An unauthenticated user can send unverified HTTP requests, which may allow the attacker to perform certain actions with administrative privileges if a logged-in user visits a malicious website.
- [Live-Hack-CVE/CVE-2019-7281](https://github.com/Live-Hack-CVE/CVE-2019-7281)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7281">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7281">

---
## CVE-2019-7280 (2019-07-01T19:15:00)
> Prima Systems FlexAir, Versions 2.3.38 and prior. The session-ID is of an insufficient length and can be exploited by brute force, which may allow a remote attacker to obtain a valid session and bypass authentication.
- [Live-Hack-CVE/CVE-2019-7280](https://github.com/Live-Hack-CVE/CVE-2019-7280)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7280">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7280">

---
## CVE-2019-7213 (2019-04-24T15:29:00)
> SmarterTools SmarterMail 16.x before build 6985 allows directory traversal. An authenticated user could delete arbitrary files or could create files in new folders in arbitrary locations on the mail server. This could lead to command execution on the server for instance by putting files inside the web directories.
- [secunnix/CVE-2019-7213](https://github.com/secunnix/CVE-2019-7213)	<img alt="forks" src="https://img.shields.io/github/forks/secunnix/CVE-2019-7213">	<img alt="stars" src="https://img.shields.io/github/stars/secunnix/CVE-2019-7213">

---
## CVE-2019-7061 (2019-05-23T18:29:00)
> Adobe Acrobat and Reader versions 2019.010.20098 and earlier, 2019.010.20098 and earlier, 2017.011.30127 and earlier version, and 2015.006.30482 and earlier have an out-of-bounds read vulnerability. Successful exploitation could lead to information disclosure .
- [Live-Hack-CVE/CVE-2019-7061](https://github.com/Live-Hack-CVE/CVE-2019-7061)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-7061">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-7061">

---
## CVE-2019-6832 (2019-09-17T20:15:00)
> A CWE-287: Authentication vulnerability exists in spaceLYnk (all versions before 2.4.0) and Wiser for KNX (all versions before 2.4.0 - formerly known as homeLYnk), which could cause loss of control when an attacker bypasses the authentication.
- [Live-Hack-CVE/CVE-2019-6832](https://github.com/Live-Hack-CVE/CVE-2019-6832)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6832">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6832">

---
## CVE-2019-6827 (2019-07-15T21:15:00)
> A CWE-787: Out-of-bounds Write vulnerability exists in Interactive Graphical SCADA System (IGSS), Version 14 and prior, which could cause a software crash when data in the mdb database is manipulated.
- [Live-Hack-CVE/CVE-2019-6827](https://github.com/Live-Hack-CVE/CVE-2019-6827)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6827">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6827">

---
## CVE-2019-6825 (2019-07-15T21:15:00)
> A CWE-427: Uncontrolled Search Path Element vulnerability exists in ProClima (all versions prior to version 8.0.0) which could allow a malicious DLL file, with the same name of any resident DLLs inside the software installation, to execute arbitrary code in all versions of ProClima prior to version 8.0.0.
- [Live-Hack-CVE/CVE-2019-6825](https://github.com/Live-Hack-CVE/CVE-2019-6825)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6825">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6825">

---
## CVE-2019-6824 (2019-07-15T21:15:00)
> A CWE-119: Buffer Errors vulnerability exists in ProClima (all versions prior to version 8.0.0) which allows an unauthenticated, remote attacker to execute arbitrary code on the targeted system in all versions of ProClima prior to version 8.0.0.
- [Live-Hack-CVE/CVE-2019-6824](https://github.com/Live-Hack-CVE/CVE-2019-6824)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6824">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6824">

---
## CVE-2019-6823 (2019-07-15T21:15:00)
> A CWE-94: Code Injection vulnerability exists in ProClima (all versions prior to version 8.0.0) which could allow an unauthenticated, remote attacker to execute arbitrary code on the targeted system in all versions of ProClima prior to version 8.0.0.
- [Live-Hack-CVE/CVE-2019-6823](https://github.com/Live-Hack-CVE/CVE-2019-6823)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6823">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6823">

---
## CVE-2019-6822 (2019-07-15T21:15:00)
> A Use After Free: CWE-416 vulnerability exists in Zelio Soft 2, V5.2 and earlier, which could cause remote code execution when opening a specially crafted Zelio Soft 2 project file.
- [Live-Hack-CVE/CVE-2019-6822](https://github.com/Live-Hack-CVE/CVE-2019-6822)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6822">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6822">

---
## CVE-2019-6812 (2019-05-22T20:29:00)
> A CWE-798 use of hardcoded credentials vulnerability exists in BMX-NOR-0200H with firmware versions prior to V1.7 IR 19 which could cause a confidentiality issue when using FTP protocol.
- [Live-Hack-CVE/CVE-2019-6812](https://github.com/Live-Hack-CVE/CVE-2019-6812)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6812">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6812">

---
## CVE-2019-6757 (2019-06-03T19:29:00)
> This vulnerability allows remote attackers to execute arbitrary code on vulnerable installations of Foxit Reader 9.4.16811. User interaction is required to exploit this vulnerability in that the target must visit a malicious page or open a malicious file. The specific flaw exists within ConvertToPDF_x86.dll. The issue results from the lack of validating the existence of an object prior to performing operations on the object. An attacker can leverage this vulnerability to execute code in the context of the current process. Was ZDI-CAN-7696.
- [Live-Hack-CVE/CVE-2019-6757](https://github.com/Live-Hack-CVE/CVE-2019-6757)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6757">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6757">

---
## CVE-2019-6756 (2019-06-03T19:29:00)
> This vulnerability allows remote attackers to disclose sensitive information on vulnerable installations of Foxit PhantomPDF 9.4.0.16811. User interaction is required to exploit this vulnerability in that the target must visit a malicious page or open a malicious file. The specific flaw exists within the parsing of HTML files. The issue results from the lack of validating the existence of an object prior to performing operations on the object. An attacker can leverage this in conjunction with other vulnerabilities to execute code in the context of the current process. Was ZDI-CAN-7769.
- [Live-Hack-CVE/CVE-2019-6756](https://github.com/Live-Hack-CVE/CVE-2019-6756)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6756">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6756">

---
## CVE-2019-6755 (2019-06-03T19:29:00)
> This vulnerability allows remote attackers to execute arbitrary code on vulnerable installations of Foxit Reader 9.3.10826. User interaction is required to exploit this vulnerability in that the target must visit a malicious page or open a malicious file. The specific flaw exists within ConvertToPDF_x86.dll. The issue results from the lack of proper validation of user-supplied data, which can result in a write past the end of an allocated object. An attacker can leverage this vulnerability to execute code in the context of the current process. Was ZDI-CAN-7613.
- [Live-Hack-CVE/CVE-2019-6755](https://github.com/Live-Hack-CVE/CVE-2019-6755)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6755">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6755">

---
## CVE-2019-6747 (2019-06-03T19:29:00)
> This vulnerability allows remote attackers to execute arbitrary code on vulnerable installations of Foxit Studio Photo 3.6.6. User interaction is required to exploit this vulnerability in that the target must visit a malicious page or open a malicious file. The specific flaw exists within the handling of EZI files. The issue results from the lack of proper validation of user-supplied data, which can result in a write past the end of an allocated structure. An attacker can leverage this vulnerability to execute code in the context of the current process. Was ZDI-CAN-7636.
- [Live-Hack-CVE/CVE-2019-6747](https://github.com/Live-Hack-CVE/CVE-2019-6747)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6747">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6747">

---
## CVE-2019-6741 (2019-06-03T19:29:00)
> This vulnerability allows remote attackers to execute arbitrary code on vulnerable installations of Samsung Galaxy S9 prior to January 2019 Security Update (SMR-JAN-2019 - SVE-2018-13467). User interaction is required to exploit this vulnerability in that the target must connect to a wireless network. The specific flaw exists within the captive portal. By manipulating HTML, an attacker can force a page redirection. An attacker can leverage this vulnerability to execute code in the context of the current process. Was ZDI-CAN-7476.
- [Live-Hack-CVE/CVE-2019-6741](https://github.com/Live-Hack-CVE/CVE-2019-6741)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6741">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6741">

---
## CVE-2019-6737 (2019-06-03T18:29:00)
> This vulnerability allows remote attackers to execute arbitrary code on vulnerable installations of Bitdefender SafePay 23.0.10.34. User interaction is required to exploit this vulnerability in that the target must visit a malicious page or open a malicious file. The specific flaw exists within the processing of TIScript. The issue lies in the handling of the openFile method, which allows for an arbitrary file write with attacker controlled data. An attacker can leverage this vulnerability execute code in the context of the current process. Was ZDI-CAN-7247.
- [Live-Hack-CVE/CVE-2019-6737](https://github.com/Live-Hack-CVE/CVE-2019-6737)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6737">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6737">

---
## CVE-2019-6706 (2019-01-23T19:29:00)
> Lua 5.3.5 has a use-after-free in lua_upvaluejoin in lapi.c. For example, a crash outcome might be achieved by an attacker who is able to trigger a debug.upvaluejoin call in which the arguments have certain relationships.
- [Live-Hack-CVE/CVE-2019-6706](https://github.com/Live-Hack-CVE/CVE-2019-6706)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6706">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6706">

---
## CVE-2019-6447 (2019-01-16T14:29:00)
> The ES File Explorer File Manager application through 4.1.9.7.4 for Android allows remote attackers to read arbitrary files or execute applications via TCP port 59777 requests on the local Wi-Fi network. This TCP port remains open after the ES application has been launched once, and responds to unauthenticated application/json data over HTTP.
- [VinuKalana/CVE-2019-6447-Android-Vulnerability-in-ES-File-Explorer](https://github.com/VinuKalana/CVE-2019-6447-Android-Vulnerability-in-ES-File-Explorer)	<img alt="forks" src="https://img.shields.io/github/forks/VinuKalana/CVE-2019-6447-Android-Vulnerability-in-ES-File-Explorer">	<img alt="stars" src="https://img.shields.io/github/stars/VinuKalana/CVE-2019-6447-Android-Vulnerability-in-ES-File-Explorer">
- [Osuni-99/CVE-2019-6447](https://github.com/Osuni-99/CVE-2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/Osuni-99/CVE-2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/Osuni-99/CVE-2019-6447">
- [k4u5h41/CVE-2019-6447](https://github.com/k4u5h41/CVE-2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/k4u5h41/CVE-2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/k4u5h41/CVE-2019-6447">
- [Kayky-cmd/CVE-2019-6447--.](https://github.com/Kayky-cmd/CVE-2019-6447--.)	<img alt="forks" src="https://img.shields.io/github/forks/Kayky-cmd/CVE-2019-6447--.">	<img alt="stars" src="https://img.shields.io/github/stars/Kayky-cmd/CVE-2019-6447--.">
- [volysandro/cve_2019-6447](https://github.com/volysandro/cve_2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/volysandro/cve_2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/volysandro/cve_2019-6447">
- [febinrev/CVE-2019-6447-ESfile-explorer-exploit](https://github.com/febinrev/CVE-2019-6447-ESfile-explorer-exploit)	<img alt="forks" src="https://img.shields.io/github/forks/febinrev/CVE-2019-6447-ESfile-explorer-exploit">	<img alt="stars" src="https://img.shields.io/github/stars/febinrev/CVE-2019-6447-ESfile-explorer-exploit">
- [julio-cfa/POC-ES-File-Explorer-CVE-2019-6447](https://github.com/julio-cfa/POC-ES-File-Explorer-CVE-2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/julio-cfa/POC-ES-File-Explorer-CVE-2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/julio-cfa/POC-ES-File-Explorer-CVE-2019-6447">
- [fs0c131y/ESFileExplorerOpenPortVuln](https://github.com/fs0c131y/ESFileExplorerOpenPortVuln)	<img alt="forks" src="https://img.shields.io/github/forks/fs0c131y/ESFileExplorerOpenPortVuln">	<img alt="stars" src="https://img.shields.io/github/stars/fs0c131y/ESFileExplorerOpenPortVuln">
- [N3H4L/CVE-2019-6447](https://github.com/N3H4L/CVE-2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/N3H4L/CVE-2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/N3H4L/CVE-2019-6447">
- [SandaRuFdo/ES-File-Explorer-Open-Port-Vulnerability---CVE-2019-6447](https://github.com/SandaRuFdo/ES-File-Explorer-Open-Port-Vulnerability---CVE-2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/SandaRuFdo/ES-File-Explorer-Open-Port-Vulnerability---CVE-2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/SandaRuFdo/ES-File-Explorer-Open-Port-Vulnerability---CVE-2019-6447">
- [Chethine/EsFileExplorer-CVE-2019-6447](https://github.com/Chethine/EsFileExplorer-CVE-2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/Chethine/EsFileExplorer-CVE-2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/Chethine/EsFileExplorer-CVE-2019-6447">
- [KasunPriyashan/CVE-2019_6447-ES-File-Explorer-Exploitation](https://github.com/KasunPriyashan/CVE-2019_6447-ES-File-Explorer-Exploitation)	<img alt="forks" src="https://img.shields.io/github/forks/KasunPriyashan/CVE-2019_6447-ES-File-Explorer-Exploitation">	<img alt="stars" src="https://img.shields.io/github/stars/KasunPriyashan/CVE-2019_6447-ES-File-Explorer-Exploitation">
- [vino-theva/CVE-2019-6447](https://github.com/vino-theva/CVE-2019-6447)	<img alt="forks" src="https://img.shields.io/github/forks/vino-theva/CVE-2019-6447">	<img alt="stars" src="https://img.shields.io/github/stars/vino-theva/CVE-2019-6447">
- [KaviDk/CVE-2019-6447-in-Mobile-Application](https://github.com/KaviDk/CVE-2019-6447-in-Mobile-Application)	<img alt="forks" src="https://img.shields.io/github/forks/KaviDk/CVE-2019-6447-in-Mobile-Application">	<img alt="stars" src="https://img.shields.io/github/stars/KaviDk/CVE-2019-6447-in-Mobile-Application">

---
## CVE-2019-6182 (2019-09-03T19:15:00)
> A stored CSV Injection vulnerability was reported in Lenovo XClarity Administrator (LXCA) versions prior to 2.5.0 that could allow an administrative user to store malformed data in LXCA Jobs and Event Log data, that could result in crafted formulas stored in an exported CSV file. The crafted formula is not executed on LXCA itself.
- [Live-Hack-CVE/CVE-2019-6182](https://github.com/Live-Hack-CVE/CVE-2019-6182)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6182">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6182">

---
## CVE-2019-6181 (2019-09-03T19:15:00)
> A reflected cross-site scripting (XSS) vulnerability was reported in Lenovo XClarity Administrator (LXCA) versions prior to 2.5.0 that could allow a crafted URL, if visited, to cause JavaScript code to be executed in the user's web browser. The JavaScript code is not executed on LXCA itself.
- [Live-Hack-CVE/CVE-2019-6181](https://github.com/Live-Hack-CVE/CVE-2019-6181)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6181">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6181">

---
## CVE-2019-6180 (2019-09-03T19:15:00)
> A stored cross-site scripting (XSS) vulnerability was reported in Lenovo XClarity Administrator (LXCA) versions prior to 2.5.0 that could allow an administrative user to cause JavaScript code to be stored in LXCA which may then be executed in the user's web browser. The JavaScript code is not executed on LXCA itself.
- [Live-Hack-CVE/CVE-2019-6180](https://github.com/Live-Hack-CVE/CVE-2019-6180)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6180">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6180">

---
## CVE-2019-6179 (2019-09-03T19:15:00)
> An XML External Entity (XXE) processing vulnerability was reported in Lenovo XClarity Administrator (LXCA) prior to version 2.5.0 , Lenovo XClarity Integrator (LXCI) for Microsoft System Center prior to version 7.7.0, and Lenovo XClarity Integrator (LXCI) for VMWare vCenter prior to version 6.1.0 that could allow information disclosure.
- [Live-Hack-CVE/CVE-2019-6179](https://github.com/Live-Hack-CVE/CVE-2019-6179)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6179">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6179">

---
## CVE-2019-6178 (2019-08-19T16:15:00)
> An information leakage vulnerability in Iomega and LenovoEMC NAS products could allow disclosure of some device details such as Share names through the device API when Personal Cloud is enabled. This does not allow read, write, delete, or any other access to the underlying file systems and their contents.
- [Live-Hack-CVE/CVE-2019-6178](https://github.com/Live-Hack-CVE/CVE-2019-6178)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6178">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6178">

---
## CVE-2019-6177 (2019-08-21T20:15:00)
> A vulnerability reported in Lenovo Solution Center version 03.12.003, which is no longer supported, could allow log files to be written to non-standard locations, potentially leading to privilege escalation. Lenovo ended support for Lenovo Solution Center and recommended that customers migrate to Lenovo Vantage or Lenovo Diagnostics in April 2018.
- [Live-Hack-CVE/CVE-2019-6177](https://github.com/Live-Hack-CVE/CVE-2019-6177)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6177">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6177">

---
## CVE-2019-6169 (2019-06-26T14:15:00)
> A vulnerability reported in Lenovo Service Bridge before version 4.1.0.1 could allow unencrypted downloads over FTP.
- [Live-Hack-CVE/CVE-2019-6169](https://github.com/Live-Hack-CVE/CVE-2019-6169)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6169">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6169">

---
## CVE-2019-6168 (2019-06-26T14:15:00)
> A vulnerability reported in Lenovo Service Bridge before version 4.1.0.1 could allow remote code execution.
- [Live-Hack-CVE/CVE-2019-6168](https://github.com/Live-Hack-CVE/CVE-2019-6168)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6168">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6168">

---
## CVE-2019-6167 (2019-06-26T14:15:00)
> A vulnerability reported in Lenovo Service Bridge before version 4.1.0.1 could allow remote code execution.
- [Live-Hack-CVE/CVE-2019-6167](https://github.com/Live-Hack-CVE/CVE-2019-6167)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6167">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6167">

---
## CVE-2019-6166 (2019-06-26T14:15:00)
> A vulnerability reported in Lenovo Service Bridge before version 4.1.0.1 could allow cross-site request forgery.
- [Live-Hack-CVE/CVE-2019-6166](https://github.com/Live-Hack-CVE/CVE-2019-6166)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6166">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6166">

---
## CVE-2019-6142 (2019-11-05T21:15:00)
> It has been reported that XSS is possible in Forcepoint Email Security, versions 8.5 and 8.5.3. It is strongly recommended that you apply the relevant hotfix in order to remediate this issue.
- [Live-Hack-CVE/CVE-2019-6142](https://github.com/Live-Hack-CVE/CVE-2019-6142)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6142">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6142">

---
## CVE-2019-6002 (2019-07-26T14:15:00)
> Cross-site scripting vulnerability in Central Dogma 0.17.0 to 0.40.1 allows remote attackers to inject arbitrary web script or HTML via unspecified vectors.
- [Live-Hack-CVE/CVE-2019-6002](https://github.com/Live-Hack-CVE/CVE-2019-6002)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-6002">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-6002">

---
## CVE-2019-5924 (2019-03-12T22:29:00)
> Cross-site request forgery (CSRF) vulnerability in Smart Forms 2.6.15 and earlier allows remote attackers to hijack the authentication of administrators via a specially crafted page.
- [Live-Hack-CVE/CVE-2019-5924](https://github.com/Live-Hack-CVE/CVE-2019-5924)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5924">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5924">

---
## CVE-2019-5891 (2019-04-01T16:29:00)
> An issue was discovered in OverIT Geocall 6.3 before build 2:346977. An unauthenticated servlet allows an attacker to obtain a cookie of an authenticated user, and login to the web application.
- [Live-Hack-CVE/CVE-2019-5891](https://github.com/Live-Hack-CVE/CVE-2019-5891)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5891">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5891">

---
## CVE-2019-5890 (2019-04-01T16:29:00)
> An issue was discovered in OverIT Geocall 6.3 before build 2:346977. Weak authentication and session management allows an authenticated user to obtain access to the Administrative control panel and execute administrative functions.
- [Live-Hack-CVE/CVE-2019-5890](https://github.com/Live-Hack-CVE/CVE-2019-5890)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5890">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5890">

---
## CVE-2019-5889 (2019-04-01T16:29:00)
> An log-management directory traversal issue was discovered in OverIT Geocall 6.3 before build 2:346977.
- [Live-Hack-CVE/CVE-2019-5889](https://github.com/Live-Hack-CVE/CVE-2019-5889)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5889">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5889">

---
## CVE-2019-5888 (2019-04-01T16:29:00)
> Multiple XSS vulnerabilities were discovered in OverIT Geocall 6.3 before build 2:346977.
- [Live-Hack-CVE/CVE-2019-5888](https://github.com/Live-Hack-CVE/CVE-2019-5888)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5888">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5888">

---
## CVE-2019-5815 (2019-12-11T01:15:00)
> Type confusion in xsltNumberFormatGetMultipleLevel prior to libxslt 1.1.33 could allow attackers to potentially exploit heap corruption via crafted XML data.
- [Live-Hack-CVE/CVE-2019-5815](https://github.com/Live-Hack-CVE/CVE-2019-5815)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5815">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5815">

---
## CVE-2019-5797 (2022-09-29T02:15:00)
> Double free in DOMStorage in Google Chrome prior to 73.0.3683.75 allowed a remote attacker to potentially exploit heap corruption via a crafted HTML page.
- [Live-Hack-CVE/CVE-2019-5797](https://github.com/Live-Hack-CVE/CVE-2019-5797)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5797">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5797">

---
## CVE-2019-5477 (2019-08-16T16:15:00)
> A command injection vulnerability in Nokogiri v1.10.3 and earlier allows commands to be executed in a subprocess via Ruby's `Kernel.open` method. Processes are vulnerable only if the undocumented method `Nokogiri::CSS::Tokenizer#load_file` is being called with unsafe user input as the filename. This vulnerability appears in code generated by the Rexical gem versions v1.0.6 and earlier. Rexical is used by Nokogiri to generate lexical scanner code for parsing CSS queries. The underlying vulnerability was addressed in Rexical v1.0.7 and Nokogiri upgraded to this version of Rexical in Nokogiri v1.10.4.
- [Live-Hack-CVE/CVE-2019-5477](https://github.com/Live-Hack-CVE/CVE-2019-5477)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5477">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5477">

---
## CVE-2019-5456 (2019-07-30T21:15:00)
> SMTP MITM refers to a malicious actor setting up an SMTP proxy server between the UniFi Controller version <= 5.10.21 and their actual SMTP server to record their SMTP credentials for malicious use later.
- [Live-Hack-CVE/CVE-2019-5456](https://github.com/Live-Hack-CVE/CVE-2019-5456)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5456">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5456">

---
## CVE-2019-5455 (2019-07-30T21:15:00)
> Bypassing lock protection exists in Nextcloud Android app 3.6.0 when creating a multi-account and aborting the process.
- [Live-Hack-CVE/CVE-2019-5455](https://github.com/Live-Hack-CVE/CVE-2019-5455)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5455">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5455">

---
## CVE-2019-5420 (2019-03-27T14:29:00)
> A remote code execution vulnerability in development mode Rails <5.2.2.1, <6.0.0.beta3 can allow an attacker to guess the automatically generated development mode secret token. This secret token can be used in combination with other Rails internals to escalate to a remote code execution exploit.
- [laffray/ruby-RCE-CVE-2019-5420-](https://github.com/laffray/ruby-RCE-CVE-2019-5420-)	<img alt="forks" src="https://img.shields.io/github/forks/laffray/ruby-RCE-CVE-2019-5420-">	<img alt="stars" src="https://img.shields.io/github/stars/laffray/ruby-RCE-CVE-2019-5420-">
- [PenTestical/CVE-2019-5420](https://github.com/PenTestical/CVE-2019-5420)	<img alt="forks" src="https://img.shields.io/github/forks/PenTestical/CVE-2019-5420">	<img alt="stars" src="https://img.shields.io/github/stars/PenTestical/CVE-2019-5420">
- [mpgn/Rails-doubletap-RCE](https://github.com/mpgn/Rails-doubletap-RCE)	<img alt="forks" src="https://img.shields.io/github/forks/mpgn/Rails-doubletap-RCE">	<img alt="stars" src="https://img.shields.io/github/stars/mpgn/Rails-doubletap-RCE">
- [trickstersec/CVE-2019-5420](https://github.com/trickstersec/CVE-2019-5420)	<img alt="forks" src="https://img.shields.io/github/forks/trickstersec/CVE-2019-5420">	<img alt="stars" src="https://img.shields.io/github/stars/trickstersec/CVE-2019-5420">
- [cved-sources/cve-2019-5420](https://github.com/cved-sources/cve-2019-5420)	<img alt="forks" src="https://img.shields.io/github/forks/cved-sources/cve-2019-5420">	<img alt="stars" src="https://img.shields.io/github/stars/cved-sources/cve-2019-5420">
- [CyberSecurityUP/CVE-2019-5420-POC](https://github.com/CyberSecurityUP/CVE-2019-5420-POC)	<img alt="forks" src="https://img.shields.io/github/forks/CyberSecurityUP/CVE-2019-5420-POC">	<img alt="stars" src="https://img.shields.io/github/stars/CyberSecurityUP/CVE-2019-5420-POC">
- [scumdestroy/CVE-2019-5420.rb](https://github.com/scumdestroy/CVE-2019-5420.rb)	<img alt="forks" src="https://img.shields.io/github/forks/scumdestroy/CVE-2019-5420.rb">	<img alt="stars" src="https://img.shields.io/github/stars/scumdestroy/CVE-2019-5420.rb">
- [mmeza-developer/CVE-2019-5420-RCE](https://github.com/mmeza-developer/CVE-2019-5420-RCE)	<img alt="forks" src="https://img.shields.io/github/forks/mmeza-developer/CVE-2019-5420-RCE">	<img alt="stars" src="https://img.shields.io/github/stars/mmeza-developer/CVE-2019-5420-RCE">
- [j4k0m/CVE-2019-5420](https://github.com/j4k0m/CVE-2019-5420)	<img alt="forks" src="https://img.shields.io/github/forks/j4k0m/CVE-2019-5420">	<img alt="stars" src="https://img.shields.io/github/stars/j4k0m/CVE-2019-5420">
- [Eremiel/CVE-2019-5420](https://github.com/Eremiel/CVE-2019-5420)	<img alt="forks" src="https://img.shields.io/github/forks/Eremiel/CVE-2019-5420">	<img alt="stars" src="https://img.shields.io/github/stars/Eremiel/CVE-2019-5420">
- [AnasTaoutaou/CVE-2019-5420](https://github.com/AnasTaoutaou/CVE-2019-5420)	<img alt="forks" src="https://img.shields.io/github/forks/AnasTaoutaou/CVE-2019-5420">	<img alt="stars" src="https://img.shields.io/github/stars/AnasTaoutaou/CVE-2019-5420">
- [knqyf263/CVE-2019-5420](https://github.com/knqyf263/CVE-2019-5420)	<img alt="forks" src="https://img.shields.io/github/forks/knqyf263/CVE-2019-5420">	<img alt="stars" src="https://img.shields.io/github/stars/knqyf263/CVE-2019-5420">

---
## CVE-2019-5418 (2019-03-27T14:29:00)
> There is a File Content Disclosure vulnerability in Action View <5.2.2.1, <5.1.6.2, <5.0.7.2, <4.2.11.1 and v3 where specially crafted accept headers can cause contents of arbitrary files on the target system's filesystem to be exposed.
- [kailing0220/CVE-2019-5418](https://github.com/kailing0220/CVE-2019-5418)	<img alt="forks" src="https://img.shields.io/github/forks/kailing0220/CVE-2019-5418">	<img alt="stars" src="https://img.shields.io/github/stars/kailing0220/CVE-2019-5418">
- [W01fh4cker/Serein](https://github.com/W01fh4cker/Serein)	<img alt="forks" src="https://img.shields.io/github/forks/W01fh4cker/Serein">	<img alt="stars" src="https://img.shields.io/github/stars/W01fh4cker/Serein">
- [mpgn/Rails-doubletap-RCE](https://github.com/mpgn/Rails-doubletap-RCE)	<img alt="forks" src="https://img.shields.io/github/forks/mpgn/Rails-doubletap-RCE">	<img alt="stars" src="https://img.shields.io/github/stars/mpgn/Rails-doubletap-RCE">
- [mpgn/CVE-2019-5418](https://github.com/mpgn/CVE-2019-5418)	<img alt="forks" src="https://img.shields.io/github/forks/mpgn/CVE-2019-5418">	<img alt="stars" src="https://img.shields.io/github/stars/mpgn/CVE-2019-5418">
- [random-robbie/CVE-2019-5418](https://github.com/random-robbie/CVE-2019-5418)	<img alt="forks" src="https://img.shields.io/github/forks/random-robbie/CVE-2019-5418">	<img alt="stars" src="https://img.shields.io/github/stars/random-robbie/CVE-2019-5418">
- [ztgrace/CVE-2019-5418-Rails3](https://github.com/ztgrace/CVE-2019-5418-Rails3)	<img alt="forks" src="https://img.shields.io/github/forks/ztgrace/CVE-2019-5418-Rails3">	<img alt="stars" src="https://img.shields.io/github/stars/ztgrace/CVE-2019-5418-Rails3">
- [Bad3r/RailroadBandit](https://github.com/Bad3r/RailroadBandit)	<img alt="forks" src="https://img.shields.io/github/forks/Bad3r/RailroadBandit">	<img alt="stars" src="https://img.shields.io/github/stars/Bad3r/RailroadBandit">
- [NotoriousRebel/RailRoadBandit](https://github.com/NotoriousRebel/RailRoadBandit)	<img alt="forks" src="https://img.shields.io/github/forks/NotoriousRebel/RailRoadBandit">	<img alt="stars" src="https://img.shields.io/github/stars/NotoriousRebel/RailRoadBandit">
- [takeokunn/CVE-2019-5418](https://github.com/takeokunn/CVE-2019-5418)	<img alt="forks" src="https://img.shields.io/github/forks/takeokunn/CVE-2019-5418">	<img alt="stars" src="https://img.shields.io/github/stars/takeokunn/CVE-2019-5418">
- [brompwnie/CVE-2019-5418-Scanner](https://github.com/brompwnie/CVE-2019-5418-Scanner)	<img alt="forks" src="https://img.shields.io/github/forks/brompwnie/CVE-2019-5418-Scanner">	<img alt="stars" src="https://img.shields.io/github/stars/brompwnie/CVE-2019-5418-Scanner">
- [omarkurt/CVE-2019-5418](https://github.com/omarkurt/CVE-2019-5418)	<img alt="forks" src="https://img.shields.io/github/forks/omarkurt/CVE-2019-5418">	<img alt="stars" src="https://img.shields.io/github/stars/omarkurt/CVE-2019-5418">

---
## CVE-2019-5114 (2019-10-25T18:15:00)
> An exploitable SQL injection vulnerability exists in the authenticated portion of YouPHPTube 7.6. Specially crafted web requests can cause SQL injections. An attacker can send a web request with parameters containing SQL injection attacks to trigger this vulnerability, potentially allowing exfiltration of the database, user credentials and,in certain configuration, access the underlying operating system.
- [Live-Hack-CVE/CVE-2019-5114](https://github.com/Live-Hack-CVE/CVE-2019-5114)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-5114">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-5114">

---
## CVE-2019-4571 (2019-09-25T20:15:00)
> IBM Content Navigator 3.0CD is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 166721.
- [Live-Hack-CVE/CVE-2019-4571](https://github.com/Live-Hack-CVE/CVE-2019-4571)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4571">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4571">

---
## CVE-2019-4566 (2019-09-24T14:15:00)
> IBM Security Key Lifecycle Manager 3.0 and 3.0.1 stores user credentials in plain in clear text which can be read by a local user. IBM X-Force ID: 166627.
- [Live-Hack-CVE/CVE-2019-4566](https://github.com/Live-Hack-CVE/CVE-2019-4566)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4566">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4566">

---
## CVE-2019-4565 (2019-09-20T16:15:00)
> IBM Security Key Lifecycle Manager 3.0 and 3.0.1 does not require that users should have strong passwords by default, which makes it easier for attackers to compromise user accounts. IBM X-Force ID: 166626.
- [Live-Hack-CVE/CVE-2019-4565](https://github.com/Live-Hack-CVE/CVE-2019-4565)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4565">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4565">

---
## CVE-2019-4515 (2019-09-24T14:15:00)
> IBM Security Key Lifecycle Manager 3.0 and 3.0.1 is vulnerable to cross-site request forgery which could allow an attacker to execute malicious and unauthorized actions transmitted from a user that the website trusts. IBM X-Force ID: 165137.
- [Live-Hack-CVE/CVE-2019-4515](https://github.com/Live-Hack-CVE/CVE-2019-4515)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4515">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4515">

---
## CVE-2019-4514 (2019-10-04T14:15:00)
> IBM Security Key Lifecycle Manager 2.6, 2.7, 3.0, and 3.0.1 discloses sensitive information to unauthorized users. The information can be used to mount further attacks on the system. IBM X-Force ID: 165136.
- [Live-Hack-CVE/CVE-2019-4514](https://github.com/Live-Hack-CVE/CVE-2019-4514)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4514">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4514">

---
## CVE-2019-4505 (2019-09-20T16:15:00)
> IBM WebSphere Application Server 7.0, 8.0, 8.5, and 9.0 Network Deployment could allow a remote attacker to obtain sensitive information, caused by sending a specially-crafted URL. This can lead the attacker to view any file in a certain directory. IBM X-Force ID: 164364.
- [Live-Hack-CVE/CVE-2019-4505](https://github.com/Live-Hack-CVE/CVE-2019-4505)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4505">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4505">

---
## CVE-2019-4494 (2019-10-01T15:15:00)
> IBM Jazz Reporting Service (JRS) 6.0, 6.0.1, 6.0.2, 6.0.3, 6.0.4, 6.0.5, 6.0.6, and 6.0.6.1 is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 164115.
- [Live-Hack-CVE/CVE-2019-4494](https://github.com/Live-Hack-CVE/CVE-2019-4494)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4494">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4494">

---
## CVE-2019-4477 (2019-09-17T19:15:00)
> IBM WebSphere Application Server 7.0, 8.0, 8.5, and 9.0 could allow a user with access to audit logs to obtain sensitive information, caused by improper handling of command line options. IBM X-Force ID: 163997.
- [Live-Hack-CVE/CVE-2019-4477](https://github.com/Live-Hack-CVE/CVE-2019-4477)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4477">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4477">

---
## CVE-2019-4456 (2019-07-30T14:15:00)
> IBM Daeja ViewONE Professional, Standard & Virtual 5.0.5 and 5.0.6 is vulnerable to an XML External Entity Injection (XXE) attack when processing XML data. A remote attacker could exploit this vulnerability to expose sensitive information or consume memory resources. IBM X-Force ID: 163620.
- [Live-Hack-CVE/CVE-2019-4456](https://github.com/Live-Hack-CVE/CVE-2019-4456)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4456">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4456">

---
## CVE-2019-4423 (2019-09-30T16:15:00)
> IBM Sterling File Gateway 2.2.0.0 through 6.0.1.0 could allow a remote attacker to traverse directories on the system. An attacker could send a specially-crafted URL request containing "dot dot" sequences (/../) to view arbitrary files on the system. IBM X-Force ID: 162769.
- [Live-Hack-CVE/CVE-2019-4423](https://github.com/Live-Hack-CVE/CVE-2019-4423)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4423">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4423">

---
## CVE-2019-4284 (2019-08-05T14:15:00)
> IBM Cloud Private 2.1.0 , 3.1.0, 3.1.1, and 3.1.2 could allow a local privileged user to obtain sensitive OIDC token that is printed to log files, which could be used to log in to the system as another user. IBM X-Force ID: 160512.
- [Live-Hack-CVE/CVE-2019-4284](https://github.com/Live-Hack-CVE/CVE-2019-4284)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4284">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4284">

---
## CVE-2019-4280 (2019-09-30T16:15:00)
> IBM Sterling File Gateway 2.2.0.0 through 6.0.1.0 displays sensitive information in HTTP requests which could be used in further attacks against the system. IBM X-Force ID: 160503.
- [Live-Hack-CVE/CVE-2019-4280](https://github.com/Live-Hack-CVE/CVE-2019-4280)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4280">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4280">
- [Live-Hack-CVE/CVE-2019-4280](https://github.com/Live-Hack-CVE/CVE-2019-4280)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4280">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4280">

---
## CVE-2019-4275 (2019-08-02T14:15:00)
> IBM Jazz for Service Management 1.1.3, 1.1.3.1, and 1.1.3.2 could allow an unauthorized local user to create unique catalog names that could cause a denial of service. IBM X-Force ID: 160296.
- [Live-Hack-CVE/CVE-2019-4275](https://github.com/Live-Hack-CVE/CVE-2019-4275)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4275">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4275">

---
## CVE-2019-4271 (2019-09-17T19:15:00)
> IBM WebSphere Application Server 7.0, 8.0, 8.5, and 9.0 Admin console is vulnerable to a Client-side HTTP parameter pollution vulnerability. IBM X-Force ID: 160243.
- [Live-Hack-CVE/CVE-2019-4271](https://github.com/Live-Hack-CVE/CVE-2019-4271)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4271">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4271">
- [Live-Hack-CVE/CVE-2019-4271](https://github.com/Live-Hack-CVE/CVE-2019-4271)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4271">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4271">

---
## CVE-2019-4270 (2019-09-17T19:15:00)
> IBM WebSphere Application Server 7.0, 8.0, 8.5, and 9.0 Admin Console is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 160203.
- [Live-Hack-CVE/CVE-2019-4270](https://github.com/Live-Hack-CVE/CVE-2019-4270)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4270">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4270">
- [Live-Hack-CVE/CVE-2019-4270](https://github.com/Live-Hack-CVE/CVE-2019-4270)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4270">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4270">

---
## CVE-2019-4268 (2019-09-17T19:15:00)
> IBM WebSphere Application Server 7.0, 8.0, 8.5, and 9.0 could allow a remote attacker to traverse directories on the system. An attacker could send a specially-crafted URL containing "dot dot" sequences (/../) to view arbitrary files on the system. IBM X-Force ID: 160201.
- [Live-Hack-CVE/CVE-2019-4268](https://github.com/Live-Hack-CVE/CVE-2019-4268)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4268">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4268">
- [Live-Hack-CVE/CVE-2019-4268](https://github.com/Live-Hack-CVE/CVE-2019-4268)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4268">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4268">

---
## CVE-2019-4267 (2019-07-22T14:15:00)
> The IBM Spectrum Protect 7.1 and 8.1 Backup-Archive Client is vulnerable to a buffer overflow. This could allow execution of arbitrary code on the local system or the application to crash. IBM X-Force ID: 160200.
- [Live-Hack-CVE/CVE-2019-4267](https://github.com/Live-Hack-CVE/CVE-2019-4267)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4267">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4267">
- [Live-Hack-CVE/CVE-2019-4267](https://github.com/Live-Hack-CVE/CVE-2019-4267)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4267">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4267">

---
## CVE-2019-4249 (2019-06-27T14:15:00)
> IBM Rational Collaborative Lifecycle Management 6.0 through 6.0.6.1 is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 159647.
- [Live-Hack-CVE/CVE-2019-4249](https://github.com/Live-Hack-CVE/CVE-2019-4249)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4249">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4249">

---
## CVE-2019-4246 (2019-10-01T15:15:00)
> IBM Daeja ViewONE Virtual 5.0 through 5.0.6 could expose internal parameters to ViewONE clients that could be used in further attacks against the system. IBM X-Force ID: 159521.
- [Live-Hack-CVE/CVE-2019-4246](https://github.com/Live-Hack-CVE/CVE-2019-4246)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4246">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4246">

---
## CVE-2019-4241 (2019-06-26T15:15:00)
> IBM PureApplication System 2.2.3.0 through 2.2.5.3 could allow an authenticated user with local access to bypass authentication and obtain administrative access. IBM X-Force ID: 159467.
- [Live-Hack-CVE/CVE-2019-4241](https://github.com/Live-Hack-CVE/CVE-2019-4241)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4241">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4241">

---
## CVE-2019-4237 (2019-07-01T15:15:00)
> A Cross-Frame Scripting vulnerability in IBM InfoSphere Information Server 11.3, 11.5, and 11.7 can allow an attacker to load the vulnerable application inside an HTML iframe tag on a malicious page. IBM X-Force ID: 159419.
- [Live-Hack-CVE/CVE-2019-4237](https://github.com/Live-Hack-CVE/CVE-2019-4237)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4237">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4237">

---
## CVE-2019-4236 (2019-07-22T14:15:00)
> A IBM Spectrum Protect 7.l client backup or archive operation running for an HP-UX VxFS object is silently skipping Access Control List (ACL) entries from backup or archive if there are more than twelve ACL entries associated with the object in total. As a result, it could allow a local attacker to restore or retrieve the object with incorrect ACL entries. IBM X-Force ID: 159418.
- [Live-Hack-CVE/CVE-2019-4236](https://github.com/Live-Hack-CVE/CVE-2019-4236)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4236">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4236">

---
## CVE-2019-4235 (2019-06-26T15:15:00)
> IBM PureApplication System 2.2.3.0 through 2.2.5.3 does not require that users should have strong passwords by default, which makes it easier for attackers to compromise user accounts. IBM X-Force ID: 159417.
- [Live-Hack-CVE/CVE-2019-4235](https://github.com/Live-Hack-CVE/CVE-2019-4235)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4235">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4235">

---
## CVE-2019-4234 (2019-06-26T15:15:00)
> IBM PureApplication System 2.2.3.0 through 2.2.5.3 weakness in the implementation of locking feature in pattern editor. An attacker by intercepting the subsequent requests can bypass business logic to modify the pattern to unlocked state. IBM X-Force ID: 159416.
- [Live-Hack-CVE/CVE-2019-4234](https://github.com/Live-Hack-CVE/CVE-2019-4234)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4234">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4234">

---
## CVE-2019-4231 (2019-12-20T17:15:00)
> IBM Cognos Analytics 11.0 and 11.1 is vulnerable to cross-site request forgery which could allow an attacker to execute malicious and unauthorized actions transmitted from a user that the website trusts. IBM X-Force ID: 159356.
- [Live-Hack-CVE/CVE-2019-4231](https://github.com/Live-Hack-CVE/CVE-2019-4231)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4231">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4231">
- [Live-Hack-CVE/CVE-2019-4231](https://github.com/Live-Hack-CVE/CVE-2019-4231)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4231">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4231">

---
## CVE-2019-4227 (2019-10-04T14:15:00)
> IBM MQ 8.0.0.4 - 8.0.0.12, 9.0.0.0 - 9.0.0.6, 9.1.0.0 - 9.1.0.2, and 9.1.0 - 9.1.2 AMQP Listeners could allow an unauthorized user to conduct a session fixation attack due to clients not being disconnected as they should. IBM X-Force ID: 159352.
- [Live-Hack-CVE/CVE-2019-4227](https://github.com/Live-Hack-CVE/CVE-2019-4227)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4227">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4227">

---
## CVE-2019-4212 (2019-07-25T15:15:00)
> IBM QRadar SIEM 7.2 and 7.3 is vulnerable to cross-site request forgery which could allow an attacker to execute malicious and unauthorized actions transmitted from a user that the website trusts. IBM X-Force ID: 159132.
- [Live-Hack-CVE/CVE-2019-4212](https://github.com/Live-Hack-CVE/CVE-2019-4212)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4212">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4212">

---
## CVE-2019-4211 (2019-07-17T14:15:00)
> IBM QRadar SIEM 7.2 and 7.3 is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 159131.
- [Live-Hack-CVE/CVE-2019-4211](https://github.com/Live-Hack-CVE/CVE-2019-4211)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4211">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4211">

---
## CVE-2019-4194 (2019-07-17T14:15:00)
> IBM Jazz for Service Management 1.1.3, 1.1.3.1, and 1.1.3.2 is missing function level access control that could allow a user to delete authorized resources. IBM X-Force ID: 159033.
- [Live-Hack-CVE/CVE-2019-4194](https://github.com/Live-Hack-CVE/CVE-2019-4194)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4194">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4194">
- [Live-Hack-CVE/CVE-2019-4194](https://github.com/Live-Hack-CVE/CVE-2019-4194)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4194">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4194">

---
## CVE-2019-4186 (2019-09-05T15:15:00)
> IBM Jazz for Service Management 1.1.3 is vulnerable to HTTP header injection, caused by incorrect trust in the HTTP Host header during caching. By sending a specially crafted HTTP GET request, a remote attacker could exploit this vulnerability to inject arbitrary HTTP headers, which will allow the attacker to conduct various attacks against the vulnerable system, including cross-site scripting, cache poisoning or session hijacking. IBM X-force ID: 158976.
- [Live-Hack-CVE/CVE-2019-4186](https://github.com/Live-Hack-CVE/CVE-2019-4186)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4186">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4186">
- [Live-Hack-CVE/CVE-2019-4186](https://github.com/Live-Hack-CVE/CVE-2019-4186)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4186">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4186">

---
## CVE-2019-4175 (2019-09-17T19:15:00)
> IBM Cognos Controller 10.3.0, 10.3.1, 10.4.0, and 10.4.1 uses weaker than expected cryptographic algorithms that could allow an attacker to decrypt highly sensitive information. IBM X-Force ID: 158880.
- [Live-Hack-CVE/CVE-2019-4175](https://github.com/Live-Hack-CVE/CVE-2019-4175)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4175">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4175">
- [Live-Hack-CVE/CVE-2019-4175](https://github.com/Live-Hack-CVE/CVE-2019-4175)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4175">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4175">

---
## CVE-2019-4171 (2019-09-17T19:15:00)
> IBM Cognos Controller 10.3.0, 10.3.1, 10.4.0, and 10.4.1 does not set the secure attribute on authorization tokens or session cookies. This could allow an attacker to obtain sensitive information using man in the middle techniques. IBM X-Force ID: 158876.
- [Live-Hack-CVE/CVE-2019-4171](https://github.com/Live-Hack-CVE/CVE-2019-4171)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4171">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4171">
- [Live-Hack-CVE/CVE-2019-4171](https://github.com/Live-Hack-CVE/CVE-2019-4171)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4171">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4171">

---
## CVE-2019-4163 (2019-07-31T17:15:00)
> IBM StoreIQ 7.6.0.0. through 7.6.0.18 could allow an authenticated user to obtain sensitive information that a privileged user should only be allowed to view. IBM X-Force ID: 158696.
- [Live-Hack-CVE/CVE-2019-4163](https://github.com/Live-Hack-CVE/CVE-2019-4163)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4163">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4163">

---
## CVE-2019-4147 (2019-09-16T19:15:00)
> IBM Sterling File Gateway 2.2.0.0 through 6.0.1.0 is vulnerable to SQL injection. A remote attacker could send specially-crafted SQL statements, which could allow the attacker to view, add, modify or delete information in the back-end database. IBM X-Force ID: 158413.
- [Live-Hack-CVE/CVE-2019-4147](https://github.com/Live-Hack-CVE/CVE-2019-4147)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4147">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4147">
- [Live-Hack-CVE/CVE-2019-4147](https://github.com/Live-Hack-CVE/CVE-2019-4147)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4147">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4147">

---
## CVE-2019-4141 (2019-09-27T14:15:00)
> IBM MQ 7.1.0.0 - 7.1.0.9, 7.5.0.0 - 7.5.0.9, 8.0.0.0 - 8.0.0.11, 9.0.0.0 - 9.0.0.6, 9.1.0.0 - 9.1.0.2, and 9.1.1 - 9.1.2 is vulnerable to a denial of service attack caused by a memory leak in the clustering code. IBM X-Force ID: 158337.
- [Live-Hack-CVE/CVE-2019-4141](https://github.com/Live-Hack-CVE/CVE-2019-4141)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4141">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4141">
- [Live-Hack-CVE/CVE-2019-4141](https://github.com/Live-Hack-CVE/CVE-2019-4141)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4141">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4141">

---
## CVE-2019-4135 (2019-06-25T16:15:00)
> IBM Security Access Manager 9.0.1 through 9.0.6 is affected by a security vulnerability that could allow authenticated users to impersonate other users. IBM X-Force ID: 158331.
- [Live-Hack-CVE/CVE-2019-4135](https://github.com/Live-Hack-CVE/CVE-2019-4135)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4135">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4135">

---
## CVE-2019-4134 (2019-07-02T15:15:00)
> IBM Planning Analytics 2.0 is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 158281.
- [Live-Hack-CVE/CVE-2019-4134](https://github.com/Live-Hack-CVE/CVE-2019-4134)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4134">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4134">

---
## CVE-2019-4133 (2019-08-29T15:15:00)
> IBM Cloud Automation Manager 3.1.2 could allow a malicious user on the client side (with access to client computer) to run a custom script. IBM X-Force ID: 158278.
- [Live-Hack-CVE/CVE-2019-4133](https://github.com/Live-Hack-CVE/CVE-2019-4133)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4133">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4133">

---
## CVE-2019-4132 (2019-08-29T15:15:00)
> IBM Cloud Automation Manager 3.1.2 could allow a user to be impropertly redirected and obtain sensitive information rather than receive a 404 error message. IBM X-Force ID: 158274.
- [Live-Hack-CVE/CVE-2019-4132](https://github.com/Live-Hack-CVE/CVE-2019-4132)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4132">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4132">

---
## CVE-2019-4120 (2019-08-20T20:15:00)
> IBM Cloud Private 3.1.1 and 3.1.2 is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 158146.
- [Live-Hack-CVE/CVE-2019-4120](https://github.com/Live-Hack-CVE/CVE-2019-4120)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4120">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4120">

---
## CVE-2019-4118 (2019-07-11T20:15:00)
> IBM Multicloud Manager 3.1.0, 3.1.1, and 3.1.2 ibm-mcm-chart could allow a local attacker with admin privileges to obtain highly sensitive information upon deployment. IBM X-Force ID: 158144.
- [Live-Hack-CVE/CVE-2019-4118](https://github.com/Live-Hack-CVE/CVE-2019-4118)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4118">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4118">

---
## CVE-2019-4117 (2019-08-20T19:15:00)
> IBM Cloud Private 3.1.1 and 3.1.2 is vulnerable to cross-site request forgery which could allow an attacker to execute malicious and unauthorized actions transmitted from a user that the website trusts. IBM X-Force ID: 158116.
- [Live-Hack-CVE/CVE-2019-4117](https://github.com/Live-Hack-CVE/CVE-2019-4117)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4117">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4117">

---
## CVE-2019-4116 (2019-07-25T15:15:00)
> IBM Cloud Private 2.1.0, 3.1.0, and 3.1.1 could disclose highly sensitive information in installer logs that could be use for further attacks against the system. IBM X-Force ID: 158115.
- [Live-Hack-CVE/CVE-2019-4116](https://github.com/Live-Hack-CVE/CVE-2019-4116)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4116">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4116">

---
## CVE-2019-4115 (2019-09-30T16:15:00)
> IBM WebSphere eXtreme Scale 8.6 Admin API is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 158113.
- [Live-Hack-CVE/CVE-2019-4115](https://github.com/Live-Hack-CVE/CVE-2019-4115)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4115">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4115">

---
## CVE-2019-4112 (2019-09-30T16:15:00)
> IBM WebSphere eXtreme Scale 8.6 Admin Console allows web pages to be stored locally which can be read by another user on the system. IBM X-Force ID: 158105.
- [Live-Hack-CVE/CVE-2019-4112](https://github.com/Live-Hack-CVE/CVE-2019-4112)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4112">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4112">

---
## CVE-2019-4109 (2019-09-30T16:15:00)
> IBM WebSphere eXtreme Scale 8.6 Admin Console could allow a remote attacker to hijack the clicking action of the victim. By persuading a victim to visit a malicious Web site, a remote attacker could exploit this vulnerability to hijack the victim's click actions and possibly launch further attacks against the victim. IBM X-Force ID: 158102.
- [Live-Hack-CVE/CVE-2019-4109](https://github.com/Live-Hack-CVE/CVE-2019-4109)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4109">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4109">

---
## CVE-2019-4106 (2019-09-30T16:15:00)
> IBM WebSphere eXtreme Scale 8.6 Admin Console is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 158099.
- [Live-Hack-CVE/CVE-2019-4106](https://github.com/Live-Hack-CVE/CVE-2019-4106)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4106">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4106">

---
## CVE-2019-4086 (2019-09-17T19:15:00)
> IBM Cloud Application Performance Management 8.1.4 could allow a remote attacker to hijack the clicking action of the victim. By persuading a victim to visit a malicious Web site, a remote attacker could exploit this vulnerability to hijack the victim's click actions and possibly launch further attacks against the victim. IBM X-Force ID: 157509.
- [Live-Hack-CVE/CVE-2019-4086](https://github.com/Live-Hack-CVE/CVE-2019-4086)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4086">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4086">

---
## CVE-2019-4084 (2019-06-27T14:15:00)
> IBM Jazz Foundation products (IBM Rational Collaborative Lifecycle Management 6.0 through 6.0.6.1) could allow an authenticated user to obtain sensitive information from CLM Applications that could be used in further attacks against the system. IBM X-Force ID: 157384.
- [Live-Hack-CVE/CVE-2019-4084](https://github.com/Live-Hack-CVE/CVE-2019-4084)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4084">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4084">

---
## CVE-2019-4083 (2019-06-27T14:15:00)
> IBM Jazz Foundation products (IBM Rational Collaborative Lifecycle Management 6.0 through 6.0.6.1) is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 157383.
- [Live-Hack-CVE/CVE-2019-4083](https://github.com/Live-Hack-CVE/CVE-2019-4083)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4083">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4083">

---
## CVE-2019-4073 (2019-04-25T15:29:00)
> IBM Sterling B2B Integrator Standard Edition 6.0.0.0 and 6.0.0.1 is vulnerable to cross-site scripting. This vulnerability allows users to embed arbitrary JavaScript code in the Web UI thus altering the intended functionality potentially leading to credentials disclosure within a trusted session. IBM X-Force ID: 157107.
- [Live-Hack-CVE/CVE-2019-4073](https://github.com/Live-Hack-CVE/CVE-2019-4073)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4073">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4073">

---
## CVE-2019-4072 (2019-05-09T15:29:00)
> IBM Tivoli Storage Productivity Center (IBM Spectrum Control Standard Edition 5.2.1 through 5.2.17) allows users to remain idle within the application even when a user has logged out. Utilizing the application back button users can remain logged in as the current user for a short period of time, therefore users are presented with information for Spectrum Control Application. IBM X-Force ID: 157064.
- [Live-Hack-CVE/CVE-2019-4072](https://github.com/Live-Hack-CVE/CVE-2019-4072)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4072">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4072">

---
## CVE-2019-4054 (2019-07-17T14:15:00)
> IBM QRadar SIEM 7.2 and 7.3 could allow a local user to obtain sensitive information when exporting content that could aid an attacker in further attacks against the system. IBM X-Force ID: 156563.
- [Live-Hack-CVE/CVE-2019-4054](https://github.com/Live-Hack-CVE/CVE-2019-4054)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4054">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4054">

---
## CVE-2019-4048 (2019-06-06T01:29:00)
> IBM Maximo Asset Management 7.6 could allow a physical user of the system to obtain sensitive information from a previous user of the same machine. IBM X-Force ID: 156311.
- [Live-Hack-CVE/CVE-2019-4048](https://github.com/Live-Hack-CVE/CVE-2019-4048)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4048">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4048">

---
## CVE-2019-4047 (2019-04-29T17:29:00)
> IBM Jazz Reporting Service (JRS) 6.0.6 could allow an authenticated user to access the execution log files as a guest user, and obtain the information of the server execution. IBM X-Force ID: 156243.
- [Live-Hack-CVE/CVE-2019-4047](https://github.com/Live-Hack-CVE/CVE-2019-4047)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-4047">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-4047">

---
## CVE-2019-3939 (2019-04-30T21:29:00)
> Crestron AM-100 with firmware 1.6.0.2 and AM-101 with firmware 2.7.0.2 use default credentials admin/admin and moderator/moderator for the web interface. An unauthenticated, remote attacker can use these credentials to gain privileged access to the device.
- [Live-Hack-CVE/CVE-2019-3939](https://github.com/Live-Hack-CVE/CVE-2019-3939)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3939">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3939">

---
## CVE-2019-3938 (2019-04-30T21:29:00)
> Crestron AM-100 with firmware 1.6.0.2 and AM-101 with firmware 2.7.0.2 stores usernames, passwords, and other configuration options in the file generated via the "export configuration" feature. The configuration file is encrypted using the awenc binary. The same binary can be used to decrypt any configuration file since all the encryption logic is hard coded. A local attacker can use this vulnerability to gain access to devices username and passwords.
- [Live-Hack-CVE/CVE-2019-3938](https://github.com/Live-Hack-CVE/CVE-2019-3938)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3938">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3938">

---
## CVE-2019-3935 (2019-04-30T21:29:00)
> Crestron AM-100 with firmware 1.6.0.2 and AM-101 with firmware 2.7.0.2 allows anyone to act as a moderator to a slide show via crafted HTTP POST requests to conference.cgi. A remote, unauthenticated attacker can use this vulnerability to start, stop, and disconnect active slideshows.
- [Live-Hack-CVE/CVE-2019-3935](https://github.com/Live-Hack-CVE/CVE-2019-3935)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3935">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3935">

---
## CVE-2019-3928 (2019-04-30T21:29:00)
> Crestron AM-100 with firmware 1.6.0.2 and AM-101 with firmware 2.7.0.2 allow any user to obtain the presentation passcode via the iso.3.6.1.4.1.3212.100.3.2.7.4 OIDs. A remote, unauthenticated attacker can use this vulnerability to access a restricted presentation or to become the presenter.
- [Live-Hack-CVE/CVE-2019-3928](https://github.com/Live-Hack-CVE/CVE-2019-3928)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3928">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3928">

---
## CVE-2019-3886 (2019-04-04T16:29:00)
> An incorrect permissions check was discovered in libvirt 4.8.0 and above. The readonly permission was allowed to invoke APIs depending on the guest agent, which could lead to potentially disclosing unintended information or denial of service by causing libvirt to block.
- [Live-Hack-CVE/CVE-2019-3886](https://github.com/Live-Hack-CVE/CVE-2019-3886)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3886">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3886">

---
## CVE-2019-3865 (2020-06-22T19:15:00)
> A vulnerability was found in quay-2, where a stored XSS vulnerability has been found in the super user function of quay. Attackers are able to use the name field of service key to inject scripts and make it run when admin users try to change the name.
- [Live-Hack-CVE/CVE-2019-3865](https://github.com/Live-Hack-CVE/CVE-2019-3865)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3865">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3865">

---
## CVE-2019-3692 (2020-01-24T09:15:00)
> The packaging of inn on SUSE Linux Enterprise Server 11; openSUSE Factory, Leap 15.1 allows local attackers to escalate from user inn to root via symlink attacks. This issue affects: SUSE Linux Enterprise Server 11 inn version 2.4.2-170.21.3.1 and prior versions. openSUSE Factory inn version 2.6.2-2.2 and prior versions. openSUSE Leap 15.1 inn version 2.5.4-lp151.2.47 and prior versions.
- [Live-Hack-CVE/CVE-2019-3692](https://github.com/Live-Hack-CVE/CVE-2019-3692)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3692">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3692">

---
## CVE-2019-3638 (2019-09-12T16:15:00)
> Reflected Cross Site Scripting vulnerability in Administrators web console in McAfee Web Gateway (MWG) 7.8.x prior to 7.8.2.13 allows remote attackers to collect sensitive information or execute commands with the MWG administrator's credentials via tricking the administrator to click on a carefully constructed malicious link.
- [Live-Hack-CVE/CVE-2019-3638](https://github.com/Live-Hack-CVE/CVE-2019-3638)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3638">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3638">

---
## CVE-2019-3635 (2019-08-14T17:15:00)
> Exfiltration of Data in McAfee Web Gateway (MWG) 7.8.2.x prior to 7.8.2.12 allows attackers to obtain sensitive data via crafting a complex webpage that will trigger the Web Gateway to block the user accessing an iframe.
- [Live-Hack-CVE/CVE-2019-3635](https://github.com/Live-Hack-CVE/CVE-2019-3635)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-3635">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-3635">

---
## CVE-2019-3560 (2019-04-29T16:29:00)
> An improperly performed length calculation on a buffer in PlaintextRecordLayer could lead to an infinite loop and denial-of-service based on user input. This issue affected versions of fizz prior to v2019.03.04.00.
- [ahaShiyu/CVE-2019-3560](https://github.com/ahaShiyu/CVE-2019-3560)	<img alt="forks" src="https://img.shields.io/github/forks/ahaShiyu/CVE-2019-3560">	<img alt="stars" src="https://img.shields.io/github/stars/ahaShiyu/CVE-2019-3560">

---
## CVE-2019-2983 (2019-10-16T18:15:00)
> Vulnerability in the Java SE, Java SE Embedded product of Oracle Java SE (component: Serialization). Supported versions that are affected are Java SE: 7u231, 8u221, 11.0.4 and 13; Java SE Embedded: 8u221. Difficult to exploit vulnerability allows unauthenticated attacker with network access via multiple protocols to compromise Java SE, Java SE Embedded. Successful attacks of this vulnerability can result in unauthorized ability to cause a partial denial of service (partial DOS) of Java SE, Java SE Embedded. Note: This vulnerability applies to Java deployments, typically in clients running sandboxed Java Web Start applications or sandboxed Java applets (in Java SE 8), that load and run untrusted code (e.g., code that comes from the internet) and rely on the Java sandbox for security. This vulnerability can also be exploited by using APIs in the specified Component, e.g., through a web service which supplies data to the APIs. CVSS 3.0 Base Score 3.7 (Availability impacts). CVSS Vector: (CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:L).
- [Live-Hack-CVE/CVE-2019-2983](https://github.com/Live-Hack-CVE/CVE-2019-2983)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-2983">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-2983">

---
## CVE-2019-2981 (2019-10-16T18:15:00)
> Vulnerability in the Java SE, Java SE Embedded product of Oracle Java SE (component: JAXP). Supported versions that are affected are Java SE: 7u231, 8u221, 11.0.4 and 13; Java SE Embedded: 8u221. Difficult to exploit vulnerability allows unauthenticated attacker with network access via multiple protocols to compromise Java SE, Java SE Embedded. Successful attacks of this vulnerability can result in unauthorized ability to cause a partial denial of service (partial DOS) of Java SE, Java SE Embedded. Note: This vulnerability applies to Java deployments, typically in clients running sandboxed Java Web Start applications or sandboxed Java applets (in Java SE 8), that load and run untrusted code (e.g., code that comes from the internet) and rely on the Java sandbox for security. This vulnerability can also be exploited by using APIs in the specified Component, e.g., through a web service which supplies data to the APIs. CVSS 3.0 Base Score 3.7 (Availability impacts). CVSS Vector: (CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:L).
- [Live-Hack-CVE/CVE-2019-2981](https://github.com/Live-Hack-CVE/CVE-2019-2981)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-2981">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-2981">

---
## CVE-2019-2964 (2019-10-16T18:15:00)
> Vulnerability in the Java SE, Java SE Embedded product of Oracle Java SE (component: Concurrency). Supported versions that are affected are Java SE: 7u231, 8u221, 11.0.4 and 13; Java SE Embedded: 8u221. Difficult to exploit vulnerability allows unauthenticated attacker with network access via multiple protocols to compromise Java SE, Java SE Embedded. Successful attacks of this vulnerability can result in unauthorized ability to cause a partial denial of service (partial DOS) of Java SE, Java SE Embedded. Note: This vulnerability can only be exploited by supplying data to APIs in the specified Component without using Untrusted Java Web Start applications or Untrusted Java applets, such as through a web service. CVSS 3.0 Base Score 3.7 (Availability impacts). CVSS Vector: (CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:L).
- [Live-Hack-CVE/CVE-2019-2964](https://github.com/Live-Hack-CVE/CVE-2019-2964)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-2964">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-2964">

---
## CVE-2019-2745 (2019-07-23T23:15:00)
> Vulnerability in the Java SE component of Oracle Java SE (subcomponent: Security). Supported versions that are affected are Java SE: 7u221, 8u212 and 11.0.3. Difficult to exploit vulnerability allows unauthenticated attacker with logon to the infrastructure where Java SE executes to compromise Java SE. Successful attacks of this vulnerability can result in unauthorized access to critical data or complete access to all Java SE accessible data. Note: This vulnerability applies to Java deployments, typically in clients running sandboxed Java Web Start applications or sandboxed Java applets (in Java SE 8), that load and run untrusted code (e.g., code that comes from the internet) and rely on the Java sandbox for security. This vulnerability can also be exploited by using APIs in the specified Component, e.g., through a web service which supplies data to the APIs. CVSS 3.0 Base Score 5.1 (Confidentiality impacts). CVSS Vector: (CVSS:3.0/AV:L/AC:H/PR:N/UI:N/S:U/C:H/I:N/A:N).
- [Live-Hack-CVE/CVE-2019-2745](https://github.com/Live-Hack-CVE/CVE-2019-2745)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-2745">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-2745">

---
## CVE-2019-2729 (2019-06-19T23:15:00)
> Vulnerability in the Oracle WebLogic Server component of Oracle Fusion Middleware (subcomponent: Web Services). Supported versions that are affected are 10.3.6.0.0, 12.1.3.0.0 and 12.2.1.3.0. Easily exploitable vulnerability allows unauthenticated attacker with network access via HTTP to compromise Oracle WebLogic Server. Successful attacks of this vulnerability can result in takeover of Oracle WebLogic Server. CVSS 3.0 Base Score 9.8 (Confidentiality, Integrity and Availability impacts). CVSS Vector: (CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H).
- [Luchoane/CVE-2019-2729_creal](https://github.com/Luchoane/CVE-2019-2729_creal)	<img alt="forks" src="https://img.shields.io/github/forks/Luchoane/CVE-2019-2729_creal">	<img alt="stars" src="https://img.shields.io/github/stars/Luchoane/CVE-2019-2729_creal">
- [0xn0ne/weblogicScanner](https://github.com/0xn0ne/weblogicScanner)	<img alt="forks" src="https://img.shields.io/github/forks/0xn0ne/weblogicScanner">	<img alt="stars" src="https://img.shields.io/github/stars/0xn0ne/weblogicScanner">
- [pizza-power/weblogic-CVE-2019-2729-POC](https://github.com/pizza-power/weblogic-CVE-2019-2729-POC)	<img alt="forks" src="https://img.shields.io/github/forks/pizza-power/weblogic-CVE-2019-2729-POC">	<img alt="stars" src="https://img.shields.io/github/stars/pizza-power/weblogic-CVE-2019-2729-POC">
- [waffl3ss/CVE-2019-2729](https://github.com/waffl3ss/CVE-2019-2729)	<img alt="forks" src="https://img.shields.io/github/forks/waffl3ss/CVE-2019-2729">	<img alt="stars" src="https://img.shields.io/github/stars/waffl3ss/CVE-2019-2729">
- [dr0op/WeblogicScan](https://github.com/dr0op/WeblogicScan)	<img alt="forks" src="https://img.shields.io/github/forks/dr0op/WeblogicScan">	<img alt="stars" src="https://img.shields.io/github/stars/dr0op/WeblogicScan">
- [ruthlezs/CVE-2019-2729-Exploit](https://github.com/ruthlezs/CVE-2019-2729-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/ruthlezs/CVE-2019-2729-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/ruthlezs/CVE-2019-2729-Exploit">
- [black-mirror/Weblogic](https://github.com/black-mirror/Weblogic)	<img alt="forks" src="https://img.shields.io/github/forks/black-mirror/Weblogic">	<img alt="stars" src="https://img.shields.io/github/stars/black-mirror/Weblogic">

---
## CVE-2019-25089 (2022-12-27T12:15:00)
> A vulnerability has been found in Morgawr Muon 0.1.1 and classified as problematic. Affected by this vulnerability is an unknown functionality of the file src/muon/handler.clj. The manipulation leads to insufficiently random values. The attack can be launched remotely. Upgrading to version 0.2.0-indev is able to address this issue. The name of the patch is c09ed972c020f759110c707b06ca2644f0bacd7f. It is recommended to upgrade the affected component. The identifier VDB-216877 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2019-25089](https://github.com/Live-Hack-CVE/CVE-2019-25089)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-25089">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-25089">

---
## CVE-2019-25088 (2022-12-27T10:15:00)
> A vulnerability was found in ytti Oxidized Web. It has been classified as problematic. Affected is an unknown function of the file lib/oxidized/web/views/conf_search.haml. The manipulation of the argument to_research leads to cross site scripting. It is possible to launch the attack remotely. The name of the patch is 55ab9bdc68b03ebce9280b8746ef31d7fdedcc45. It is recommended to apply a patch to fix this issue. VDB-216870 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2019-25088](https://github.com/Live-Hack-CVE/CVE-2019-25088)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-25088">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-25088">

---
## CVE-2019-25087 (2022-12-27T09:15:00)
> A vulnerability was found in RamseyK httpserver. It has been rated as critical. This issue affects the function ResourceHost::getResource of the file src/ResourceHost.cpp of the component URI Handler. The manipulation of the argument uri leads to path traversal: '../filedir'. The attack may be initiated remotely. The name of the patch is 1a0de56e4dafff9c2f9c8f6b130a764f7a50df52. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-216863.
- [Live-Hack-CVE/CVE-2019-25087](https://github.com/Live-Hack-CVE/CVE-2019-25087)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-25087">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-25087">

---
## CVE-2019-25086 (2022-12-27T09:15:00)
> A vulnerability was found in IET-OU Open Media Player up to 1.5.0. It has been declared as problematic. This vulnerability affects the function webvtt of the file application/controllers/timedtext.php. The manipulation of the argument ttml_url leads to cross site scripting. The attack can be initiated remotely. Upgrading to version 1.5.1 is able to address this issue. The name of the patch is 3f39f2d68d11895929c04f7b49b97a734ae7cd1f. It is recommended to upgrade the affected component. VDB-216862 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2019-25086](https://github.com/Live-Hack-CVE/CVE-2019-25086)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-25086">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-25086">

---
## CVE-2019-25084 (2022-12-25T18:15:00)
> A vulnerability, which was classified as problematic, has been found in Hide Files on GitHub up to 2.x. This issue affects the function addEventListener of the file extension/options.js. The manipulation leads to cross site scripting. The attack may be initiated remotely. Upgrading to version 3.0.0 is able to address this issue. The name of the patch is 9de0c57df81db1178e0e79431d462f6d9842742e. It is recommended to upgrade the affected component. The associated identifier of this vulnerability is VDB-216767.
- [Live-Hack-CVE/CVE-2019-25084](https://github.com/Live-Hack-CVE/CVE-2019-25084)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-25084">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-25084">
- [Live-Hack-CVE/CVE-2019-25084](https://github.com/Live-Hack-CVE/CVE-2019-25084)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-25084">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-25084">

---
## CVE-2019-25013 (2021-01-04T18:15:00)
> The iconv feature in the GNU C Library (aka glibc or libc6) through 2.32, when processing invalid multi-byte input sequences in the EUC-KR encoding, may have a buffer over-read.
- [Live-Hack-CVE/CVE-2019-25013](https://github.com/Live-Hack-CVE/CVE-2019-25013)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-25013">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-25013">

---
## CVE-2019-20892 (2020-06-25T10:15:00)
> net-snmp before 5.8.1.pre1 has a double free in usm_free_usmStateReference in snmplib/snmpusm.c via an SNMPv3 GetBulk request. NOTE: this affects net-snmp packages shipped to end users by multiple Linux distributions, but might not affect an upstream release.
- [Live-Hack-CVE/CVE-2019-20892](https://github.com/Live-Hack-CVE/CVE-2019-20892)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-20892">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-20892">

---
## CVE-2019-20790 (2020-04-27T14:15:00)
> OpenDMARC through 1.3.2 and 1.4.x, when used with pypolicyd-spf 2.0.2, allows attacks that bypass SPF and DMARC authentication in situations where the HELO field is inconsistent with the MAIL FROM field.
- [Live-Hack-CVE/CVE-2019-20790](https://github.com/Live-Hack-CVE/CVE-2019-20790)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-20790">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-20790">

---
## CVE-2019-20446 (2020-02-02T14:15:00)
> In xml.rs in GNOME librsvg before 2.46.2, a crafted SVG file with nested patterns can cause denial of service when passed to the library for processing. The attacker constructs pattern elements so that the number of final rendered objects grows exponentially.
- [Live-Hack-CVE/CVE-2019-20446](https://github.com/Live-Hack-CVE/CVE-2019-20446)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-20446">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-20446">

---
## CVE-2019-20172 (2019-12-31T03:15:00)
> Kernel/VM/MemoryManager.cpp in SerenityOS before 2019-12-30 does not reject syscalls with pointers into the kernel-only virtual address space, which allows local users to gain privileges by overwriting a return address that was found on the kernel stack.
- [Live-Hack-CVE/CVE-2019-20172](https://github.com/Live-Hack-CVE/CVE-2019-20172)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-20172">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-20172">

---
## CVE-2019-19966 (2019-12-25T04:15:00)
> In the Linux kernel before 5.1.6, there is a use-after-free in cpia2_exit() in drivers/media/usb/cpia2/cpia2_v4l.c that will cause denial of service, aka CID-dea37a972655.
- [Live-Hack-CVE/CVE-2019-19966](https://github.com/Live-Hack-CVE/CVE-2019-19966)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19966">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19966">

---
## CVE-2019-19953 (2019-12-24T01:15:00)
> In GraphicsMagick 1.4 snapshot-20191208 Q8, there is a heap-based buffer over-read in the function EncodeImage of coders/pict.c.
- [Live-Hack-CVE/CVE-2019-19953](https://github.com/Live-Hack-CVE/CVE-2019-19953)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19953">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19953">

---
## CVE-2019-19951 (2019-12-24T01:15:00)
> In GraphicsMagick 1.4 snapshot-20190423 Q8, there is a heap-based buffer overflow in the function ImportRLEPixels of coders/miff.c.
- [Live-Hack-CVE/CVE-2019-19951](https://github.com/Live-Hack-CVE/CVE-2019-19951)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19951">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19951">

---
## CVE-2019-19950 (2019-12-24T01:15:00)
> In GraphicsMagick 1.4 snapshot-20190403 Q8, there is a use-after-free in ThrowException and ThrowLoggedException of magick/error.c.
- [Live-Hack-CVE/CVE-2019-19950](https://github.com/Live-Hack-CVE/CVE-2019-19950)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19950">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19950">

---
## CVE-2019-19949 (2019-12-24T01:15:00)
> In ImageMagick 7.0.8-43 Q16, there is a heap-based buffer over-read in the function WritePNGImage of coders/png.c, related to Magick_png_write_raw_profile and LocaleNCompare.
- [Live-Hack-CVE/CVE-2019-19949](https://github.com/Live-Hack-CVE/CVE-2019-19949)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19949">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19949">

---
## CVE-2019-19948 (2019-12-24T01:15:00)
> In ImageMagick 7.0.8-43 Q16, there is a heap-based buffer overflow in the function WriteSGIImage of coders/sgi.c.
- [Live-Hack-CVE/CVE-2019-19948](https://github.com/Live-Hack-CVE/CVE-2019-19948)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19948">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19948">

---
## CVE-2019-19947 (2019-12-24T00:15:00)
> In the Linux kernel through 5.4.6, there are information leaks of uninitialized memory to a USB device in the drivers/net/can/usb/kvaser_usb/kvaser_usb_leaf.c driver, aka CID-da2311a6385c.
- [Live-Hack-CVE/CVE-2019-19947](https://github.com/Live-Hack-CVE/CVE-2019-19947)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19947">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19947">

---
## CVE-2019-19945 (2020-03-16T18:15:00)
> uhttpd in OpenWrt through 18.06.5 and 19.x through 19.07.0-rc2 has an integer signedness error. This leads to out-of-bounds access to a heap buffer and a subsequent crash. It can be triggered with an HTTP POST request to a CGI script, specifying both "Transfer-Encoding: chunked" and a large negative Content-Length value.
- [delicateByte/CVE-2019-19945_Test](https://github.com/delicateByte/CVE-2019-19945_Test)	<img alt="forks" src="https://img.shields.io/github/forks/delicateByte/CVE-2019-19945_Test">	<img alt="stars" src="https://img.shields.io/github/stars/delicateByte/CVE-2019-19945_Test">

---
## CVE-2019-19920 (2019-12-22T18:15:00)
> sa-exim 4.2.1 allows attackers to execute arbitrary code if they can write a .cf file or a rule. This occurs because Greylisting.pm relies on eval (rather than direct parsing and/or use of the taint feature). This issue is similar to CVE-2018-11805.
- [Live-Hack-CVE/CVE-2019-19920](https://github.com/Live-Hack-CVE/CVE-2019-19920)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19920">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19920">
- [Live-Hack-CVE/CVE-2019-19920](https://github.com/Live-Hack-CVE/CVE-2019-19920)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19920">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19920">

---
## CVE-2019-19918 (2019-12-20T20:15:00)
> Lout 3.40 has a heap-based buffer overflow in the srcnext() function in z02.c.
- [Live-Hack-CVE/CVE-2019-19918](https://github.com/Live-Hack-CVE/CVE-2019-19918)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19918">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19918">
- [Live-Hack-CVE/CVE-2019-19918](https://github.com/Live-Hack-CVE/CVE-2019-19918)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19918">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19918">

---
## CVE-2019-19787 (2019-12-13T16:15:00)
> ATasm 1.06 has a stack-based buffer overflow in the get_signed_expression() function in setparse.c via a crafted .m65 file.
- [Live-Hack-CVE/CVE-2019-19787](https://github.com/Live-Hack-CVE/CVE-2019-19787)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19787">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19787">

---
## CVE-2019-19786 (2019-12-13T16:15:00)
> ATasm 1.06 has a stack-based buffer overflow in the parse_expr() function in setparse.c via a crafted .m65 file.
- [Live-Hack-CVE/CVE-2019-19786](https://github.com/Live-Hack-CVE/CVE-2019-19786)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19786">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19786">

---
## CVE-2019-19785 (2019-12-13T16:15:00)
> ATasm 1.06 has a stack-based buffer overflow in the to_comma() function in asm.c via a crafted .m65 file.
- [Live-Hack-CVE/CVE-2019-19785](https://github.com/Live-Hack-CVE/CVE-2019-19785)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19785">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19785">

---
## CVE-2019-19781 (2019-12-27T14:15:00)
> An issue was discovered in Citrix Application Delivery Controller (ADC) and Gateway 10.5, 11.1, 12.0, 12.1, and 13.0. They allow Directory Traversal.
- [Vulnmachines/Ctirix_RCE-CVE-2019-19781](https://github.com/Vulnmachines/Ctirix_RCE-CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/Vulnmachines/Ctirix_RCE-CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/Vulnmachines/Ctirix_RCE-CVE-2019-19781">
- [hackingyseguridad/nmap](https://github.com/hackingyseguridad/nmap)	<img alt="forks" src="https://img.shields.io/github/forks/hackingyseguridad/nmap">	<img alt="stars" src="https://img.shields.io/github/stars/hackingyseguridad/nmap">
- [k-fire/CVE-2019-19781-exploit](https://github.com/k-fire/CVE-2019-19781-exploit)	<img alt="forks" src="https://img.shields.io/github/forks/k-fire/CVE-2019-19781-exploit">	<img alt="stars" src="https://img.shields.io/github/stars/k-fire/CVE-2019-19781-exploit">
- [VladRico/CVE-2019-19781](https://github.com/VladRico/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/VladRico/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/VladRico/CVE-2019-19781">
- [pwn3z/CVE-2019-19781-Citrix](https://github.com/pwn3z/CVE-2019-19781-Citrix)	<img alt="forks" src="https://img.shields.io/github/forks/pwn3z/CVE-2019-19781-Citrix">	<img alt="stars" src="https://img.shields.io/github/stars/pwn3z/CVE-2019-19781-Citrix">
- [mpgn/CVE-2019-19781](https://github.com/mpgn/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/mpgn/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/mpgn/CVE-2019-19781">
- [LeapBeyond/cve_2019_19781](https://github.com/LeapBeyond/cve_2019_19781)	<img alt="forks" src="https://img.shields.io/github/forks/LeapBeyond/cve_2019_19781">	<img alt="stars" src="https://img.shields.io/github/stars/LeapBeyond/cve_2019_19781">
- [DanielWep/CVE-NetScalerFileSystemCheck](https://github.com/DanielWep/CVE-NetScalerFileSystemCheck)	<img alt="forks" src="https://img.shields.io/github/forks/DanielWep/CVE-NetScalerFileSystemCheck">	<img alt="stars" src="https://img.shields.io/github/stars/DanielWep/CVE-NetScalerFileSystemCheck">
- [cisagov/check-cve-2019-19781](https://github.com/cisagov/check-cve-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/cisagov/check-cve-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/cisagov/check-cve-2019-19781">
- [andripwn/CVE-2019-19781](https://github.com/andripwn/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/andripwn/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/andripwn/CVE-2019-19781">
- [w4fz5uck5/CVE-2019-19781-CitrixRCE](https://github.com/w4fz5uck5/CVE-2019-19781-CitrixRCE)	<img alt="forks" src="https://img.shields.io/github/forks/w4fz5uck5/CVE-2019-19781-CitrixRCE">	<img alt="stars" src="https://img.shields.io/github/stars/w4fz5uck5/CVE-2019-19781-CitrixRCE">
- [jamesjguthrie/Shitrix-CVE-2019-19781](https://github.com/jamesjguthrie/Shitrix-CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/jamesjguthrie/Shitrix-CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/jamesjguthrie/Shitrix-CVE-2019-19781">
- [qiong-qi/CVE-2019-19781-poc](https://github.com/qiong-qi/CVE-2019-19781-poc)	<img alt="forks" src="https://img.shields.io/github/forks/qiong-qi/CVE-2019-19781-poc">	<img alt="stars" src="https://img.shields.io/github/stars/qiong-qi/CVE-2019-19781-poc">
- [SharpHack/CVE-2019-19781](https://github.com/SharpHack/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/SharpHack/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/SharpHack/CVE-2019-19781">
- [yukar1z0e/CVE-2019-19781](https://github.com/yukar1z0e/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/yukar1z0e/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/yukar1z0e/CVE-2019-19781">
- [Roshi99/Remote-Code-Execution-Exploit-for-Citrix-Application-Delivery-Controller-and-Citrix-Gateway-CVE-201](https://github.com/Roshi99/Remote-Code-Execution-Exploit-for-Citrix-Application-Delivery-Controller-and-Citrix-Gateway-CVE-201)	<img alt="forks" src="https://img.shields.io/github/forks/Roshi99/Remote-Code-Execution-Exploit-for-Citrix-Application-Delivery-Controller-and-Citrix-Gateway-CVE-201">	<img alt="stars" src="https://img.shields.io/github/stars/Roshi99/Remote-Code-Execution-Exploit-for-Citrix-Application-Delivery-Controller-and-Citrix-Gateway-CVE-201">
- [darren646/CVE-2019-19781POC](https://github.com/darren646/CVE-2019-19781POC)	<img alt="forks" src="https://img.shields.io/github/forks/darren646/CVE-2019-19781POC">	<img alt="stars" src="https://img.shields.io/github/stars/darren646/CVE-2019-19781POC">
- [citrix/ioc-scanner-CVE-2019-19781](https://github.com/citrix/ioc-scanner-CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/citrix/ioc-scanner-CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/citrix/ioc-scanner-CVE-2019-19781">
- [mandiant/ioc-scanner-CVE-2019-19781](https://github.com/mandiant/ioc-scanner-CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/mandiant/ioc-scanner-CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/mandiant/ioc-scanner-CVE-2019-19781">
- [nmanzi/webcvescanner](https://github.com/nmanzi/webcvescanner)	<img alt="forks" src="https://img.shields.io/github/forks/nmanzi/webcvescanner">	<img alt="stars" src="https://img.shields.io/github/stars/nmanzi/webcvescanner">
- [L4r1k/CitrixNetscalerTriageScript](https://github.com/L4r1k/CitrixNetscalerTriageScript)	<img alt="forks" src="https://img.shields.io/github/forks/L4r1k/CitrixNetscalerTriageScript">	<img alt="stars" src="https://img.shields.io/github/stars/L4r1k/CitrixNetscalerTriageScript">
- [L4r1k/CitrixNetscalerAnalysis](https://github.com/L4r1k/CitrixNetscalerAnalysis)	<img alt="forks" src="https://img.shields.io/github/forks/L4r1k/CitrixNetscalerAnalysis">	<img alt="stars" src="https://img.shields.io/github/stars/L4r1k/CitrixNetscalerAnalysis">
- [onSec-fr/CVE-2019-19781-Forensic](https://github.com/onSec-fr/CVE-2019-19781-Forensic)	<img alt="forks" src="https://img.shields.io/github/forks/onSec-fr/CVE-2019-19781-Forensic">	<img alt="stars" src="https://img.shields.io/github/stars/onSec-fr/CVE-2019-19781-Forensic">
- [x1sec/CVE-2019-19781](https://github.com/x1sec/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/x1sec/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/x1sec/CVE-2019-19781">
- [j81blog/ADC-19781](https://github.com/j81blog/ADC-19781)	<img alt="forks" src="https://img.shields.io/github/forks/j81blog/ADC-19781">	<img alt="stars" src="https://img.shields.io/github/stars/j81blog/ADC-19781">
- [0xams/citrixvulncheck](https://github.com/0xams/citrixvulncheck)	<img alt="forks" src="https://img.shields.io/github/forks/0xams/citrixvulncheck">	<img alt="stars" src="https://img.shields.io/github/stars/0xams/citrixvulncheck">
- [zenturacp/cve-2019-19781-web](https://github.com/zenturacp/cve-2019-19781-web)	<img alt="forks" src="https://img.shields.io/github/forks/zenturacp/cve-2019-19781-web">	<img alt="stars" src="https://img.shields.io/github/stars/zenturacp/cve-2019-19781-web">
- [RaulCalvoLaorden/CVE-2019-19781](https://github.com/RaulCalvoLaorden/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/RaulCalvoLaorden/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/RaulCalvoLaorden/CVE-2019-19781">
- [Azeemering/CVE-2019-19781-DFIR-Notes](https://github.com/Azeemering/CVE-2019-19781-DFIR-Notes)	<img alt="forks" src="https://img.shields.io/github/forks/Azeemering/CVE-2019-19781-DFIR-Notes">	<img alt="stars" src="https://img.shields.io/github/stars/Azeemering/CVE-2019-19781-DFIR-Notes">
- [x1sec/citrix-honeypot](https://github.com/x1sec/citrix-honeypot)	<img alt="forks" src="https://img.shields.io/github/forks/x1sec/citrix-honeypot">	<img alt="stars" src="https://img.shields.io/github/stars/x1sec/citrix-honeypot">
- [trhacknon/CVE-2019-19781](https://github.com/trhacknon/CVE-2019-19781)	<img alt="forks" src="https://img.shields.io/github/forks/trhacknon/CVE-2019-19781">	<img alt="stars" src="https://img.shields.io/github/stars/trhacknon/CVE-2019-19781">
- [34zY/APT-Backpack](https://github.com/34zY/APT-Backpack)	<img alt="forks" src="https://img.shields.io/github/forks/34zY/APT-Backpack">	<img alt="stars" src="https://img.shields.io/github/stars/34zY/APT-Backpack">

---
## CVE-2019-19725 (2019-12-11T18:16:00)
> sysstat through 12.2.0 has a double free in check_file_actlst in sa_common.c.
- [Live-Hack-CVE/CVE-2019-19725](https://github.com/Live-Hack-CVE/CVE-2019-19725)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19725">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19725">

---
## CVE-2019-19221 (2019-11-21T23:15:00)
> In Libarchive 3.4.0, archive_wstring_append_from_mbs in archive_string.c has an out-of-bounds read because of an incorrect mbrtowc or mbtowc call. For example, bsdtar crashes via a crafted archive.
- [Live-Hack-CVE/CVE-2019-19221](https://github.com/Live-Hack-CVE/CVE-2019-19221)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19221">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19221">

---
## CVE-2019-19054 (2019-11-18T06:15:00)
> A memory leak in the cx23888_ir_probe() function in drivers/media/pci/cx23885/cx23888-ir.c in the Linux kernel through 5.3.11 allows attackers to cause a denial of service (memory consumption) by triggering kfifo_alloc() failures, aka CID-a7b2df76b42b.
- [Live-Hack-CVE/CVE-2019-19054](https://github.com/Live-Hack-CVE/CVE-2019-19054)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-19054">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-19054">

---
## CVE-2019-18901 (2020-03-02T16:15:00)
> A UNIX Symbolic Link (Symlink) Following vulnerability in the mysql-systemd-helper of the mariadb packaging of SUSE Linux Enterprise Server 12, SUSE Linux Enterprise Server 15 allows local attackers to change the permissions of arbitrary files to 0640. This issue affects: SUSE Linux Enterprise Server 12 mariadb versions prior to 10.2.31-3.25.1. SUSE Linux Enterprise Server 15 mariadb versions prior to 10.2.31-3.26.1.
- [Live-Hack-CVE/CVE-2019-18901](https://github.com/Live-Hack-CVE/CVE-2019-18901)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18901">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18901">

---
## CVE-2019-18897 (2020-03-02T16:15:00)
> A UNIX Symbolic Link (Symlink) Following vulnerability in the packaging of salt of SUSE Linux Enterprise Server 12, SUSE Linux Enterprise Server 15; openSUSE Factory allows local attackers to escalate privileges from user salt to root. This issue affects: SUSE Linux Enterprise Server 12 salt-master version 2019.2.0-46.83.1 and prior versions. SUSE Linux Enterprise Server 15 salt-master version 2019.2.0-6.21.1 and prior versions. openSUSE Factory salt-master version 2019.2.2-3.1 and prior versions.
- [Live-Hack-CVE/CVE-2019-18897](https://github.com/Live-Hack-CVE/CVE-2019-18897)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18897">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18897">

---
## CVE-2019-18845 (2019-11-09T18:15:00)
> The MsIo64.sys and MsIo32.sys drivers in Patriot Viper RGB before 1.1 allow local users (including low integrity processes) to read and write to arbitrary memory locations, and consequently gain NT AUTHORITY\SYSTEM privileges, by mapping \Device\PhysicalMemory into the calling process via ZwOpenSection and ZwMapViewOfSection.
- [Exploitables/CVE-2019-18845](https://github.com/Exploitables/CVE-2019-18845)	<img alt="forks" src="https://img.shields.io/github/forks/Exploitables/CVE-2019-18845">	<img alt="stars" src="https://img.shields.io/github/stars/Exploitables/CVE-2019-18845">
- [hfiref0x/KDU](https://github.com/hfiref0x/KDU)	<img alt="forks" src="https://img.shields.io/github/forks/hfiref0x/KDU">	<img alt="stars" src="https://img.shields.io/github/stars/hfiref0x/KDU">
- [kkent030315/MsIoExploit](https://github.com/kkent030315/MsIoExploit)	<img alt="forks" src="https://img.shields.io/github/forks/kkent030315/MsIoExploit">	<img alt="stars" src="https://img.shields.io/github/stars/kkent030315/MsIoExploit">

---
## CVE-2019-18683 (2019-11-04T16:15:00)
> An issue was discovered in drivers/media/platform/vivid in the Linux kernel through 5.3.8. It is exploitable for privilege escalation on some Linux distributions where local users have /dev/video0 access, but only if the driver happens to be loaded. There are multiple race conditions during streaming stopping in this driver (part of the V4L2 subsystem). These issues are caused by wrong mutex locking in vivid_stop_generating_vid_cap(), vivid_stop_generating_vid_out(), sdr_cap_stop_streaming(), and the corresponding kthreads. At least one of these race conditions leads to a use-after-free.
- [Limesss/cve-2019-18683](https://github.com/Limesss/cve-2019-18683)	<img alt="forks" src="https://img.shields.io/github/forks/Limesss/cve-2019-18683">	<img alt="stars" src="https://img.shields.io/github/stars/Limesss/cve-2019-18683">
- [sanjana123-cloud/CVE-2019-18683](https://github.com/sanjana123-cloud/CVE-2019-18683)	<img alt="forks" src="https://img.shields.io/github/forks/sanjana123-cloud/CVE-2019-18683">	<img alt="stars" src="https://img.shields.io/github/stars/sanjana123-cloud/CVE-2019-18683">

---
## CVE-2019-18634 (2020-01-29T18:15:00)
> In Sudo before 1.8.26, if pwfeedback is enabled in /etc/sudoers, users can trigger a stack-based buffer overflow in the privileged sudo process. (pwfeedback is a default setting in Linux Mint and elementary OS; however, it is NOT the default for upstream and many other packages, and would exist only if enabled by an administrator.) The attacker needs to deliver a long string to the stdin of getln() in tgetpass.c.
- [ptef/CVE-2019-18634](https://github.com/ptef/CVE-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/ptef/CVE-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/ptef/CVE-2019-18634">
- [TheJoyOfHacking/saleemrashid-sudo-cve-2019-18634](https://github.com/TheJoyOfHacking/saleemrashid-sudo-cve-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/TheJoyOfHacking/saleemrashid-sudo-cve-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/TheJoyOfHacking/saleemrashid-sudo-cve-2019-18634">
- [saleemrashid/sudo-cve-2019-18634](https://github.com/saleemrashid/sudo-cve-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/saleemrashid/sudo-cve-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/saleemrashid/sudo-cve-2019-18634">
- [aesophor/CVE-2019-18634](https://github.com/aesophor/CVE-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/aesophor/CVE-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/aesophor/CVE-2019-18634">
- [N1et/CVE-2019-18634](https://github.com/N1et/CVE-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/N1et/CVE-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/N1et/CVE-2019-18634">
- [Y3A/CVE-2019-18634](https://github.com/Y3A/CVE-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/Y3A/CVE-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/Y3A/CVE-2019-18634">
- [edsonjt81/sudo-cve-2019-18634](https://github.com/edsonjt81/sudo-cve-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/edsonjt81/sudo-cve-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/edsonjt81/sudo-cve-2019-18634">
- [Plazmaz/CVE-2019-18634](https://github.com/Plazmaz/CVE-2019-18634)	<img alt="forks" src="https://img.shields.io/github/forks/Plazmaz/CVE-2019-18634">	<img alt="stars" src="https://img.shields.io/github/stars/Plazmaz/CVE-2019-18634">

---
## CVE-2019-18413 (2019-10-24T18:15:00)
> In TypeStack class-validator 0.10.2, validate() input validation can be bypassed because certain internal attributes can be overwritten via a conflicting name. Even though there is an optional forbidUnknownValues parameter that can be used to reduce the risk of this bypass, this option is not documented and thus most developers configure input validation in the vulnerable default manner. With this vulnerability, attackers can launch SQL Injection or XSS attacks by injecting arbitrary malicious input. NOTE: a software maintainer agrees with the "is not documented" finding but suggests that much of the responsibility for the risk lies in a different product.
- [Live-Hack-CVE/CVE-2019-18413](https://github.com/Live-Hack-CVE/CVE-2019-18413)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18413">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18413">

---
## CVE-2019-18391 (2019-12-23T16:15:00)
> A heap-based buffer overflow in the vrend_renderer_transfer_write_iov function in vrend_renderer.c in virglrenderer through 0.8.0 allows guest OS users to cause a denial of service via VIRGL_CCMD_RESOURCE_INLINE_WRITE commands.
- [Live-Hack-CVE/CVE-2019-18391](https://github.com/Live-Hack-CVE/CVE-2019-18391)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18391">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18391">

---
## CVE-2019-18390 (2019-12-23T16:15:00)
> An out-of-bounds read in the vrend_blit_need_swizzle function in vrend_renderer.c in virglrenderer through 0.8.0 allows guest OS users to cause a denial of service via VIRGL_CCMD_BLIT commands.
- [Live-Hack-CVE/CVE-2019-18390](https://github.com/Live-Hack-CVE/CVE-2019-18390)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18390">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18390">

---
## CVE-2019-18389 (2019-12-23T16:15:00)
> A heap-based buffer overflow in the vrend_renderer_transfer_write_iov function in vrend_renderer.c in virglrenderer through 0.8.0 allows guest OS users to cause a denial of service, or QEMU guest-to-host escape and code execution, via VIRGL_CCMD_RESOURCE_INLINE_WRITE commands.
- [Live-Hack-CVE/CVE-2019-18389](https://github.com/Live-Hack-CVE/CVE-2019-18389)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18389">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18389">

---
## CVE-2019-18388 (2019-12-23T16:15:00)
> A NULL pointer dereference in vrend_renderer.c in virglrenderer through 0.8.0 allows guest OS users to cause a denial of service via malformed commands.
- [Live-Hack-CVE/CVE-2019-18388](https://github.com/Live-Hack-CVE/CVE-2019-18388)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18388">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18388">

---
## CVE-2019-18265 (2022-11-30T23:15:00)
> Digital Alert Systems’ DASDEC software prior to version 4.1 contains a cross-site scripting (XSS) vulnerability that allows remote attackers to inject arbitrary web script or HTML via the SSH username, username field of the login page, or via the HTTP host header. The injected content is stored in logs and rendered when viewed in the web application.
- [Live-Hack-CVE/CVE-2019-18265](https://github.com/Live-Hack-CVE/CVE-2019-18265)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-18265">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-18265">

---
## CVE-2019-17662 (2019-10-16T18:15:00)
> ThinVNC 1.0b1 is vulnerable to arbitrary file read, which leads to a compromise of the VNC server. The vulnerability exists even when authentication is turned on during the deployment of the VNC server. The password for authentication is stored in cleartext in a file that can be read via a ../../ThinVnc.ini directory traversal attack vector.
- [Tamagaft/CVE-2019-17662](https://github.com/Tamagaft/CVE-2019-17662)	<img alt="forks" src="https://img.shields.io/github/forks/Tamagaft/CVE-2019-17662">	<img alt="stars" src="https://img.shields.io/github/stars/Tamagaft/CVE-2019-17662">
- [k4is3r13/Bash-Script-CVE-2019-17662](https://github.com/k4is3r13/Bash-Script-CVE-2019-17662)	<img alt="forks" src="https://img.shields.io/github/forks/k4is3r13/Bash-Script-CVE-2019-17662">	<img alt="stars" src="https://img.shields.io/github/stars/k4is3r13/Bash-Script-CVE-2019-17662">
- [rajendrakumaryadav/CVE-2019-17662-Exploit](https://github.com/rajendrakumaryadav/CVE-2019-17662-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/rajendrakumaryadav/CVE-2019-17662-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/rajendrakumaryadav/CVE-2019-17662-Exploit">
- [whokilleddb/CVE-2019-17662](https://github.com/whokilleddb/CVE-2019-17662)	<img alt="forks" src="https://img.shields.io/github/forks/whokilleddb/CVE-2019-17662">	<img alt="stars" src="https://img.shields.io/github/stars/whokilleddb/CVE-2019-17662">
- [MuirlandOracle/CVE-2019-17662](https://github.com/MuirlandOracle/CVE-2019-17662)	<img alt="forks" src="https://img.shields.io/github/forks/MuirlandOracle/CVE-2019-17662">	<img alt="stars" src="https://img.shields.io/github/stars/MuirlandOracle/CVE-2019-17662">
- [bl4ck574r/CVE-2019-17662](https://github.com/bl4ck574r/CVE-2019-17662)	<img alt="forks" src="https://img.shields.io/github/forks/bl4ck574r/CVE-2019-17662">	<img alt="stars" src="https://img.shields.io/github/stars/bl4ck574r/CVE-2019-17662">

---
## CVE-2019-17621 (2019-12-30T17:15:00)
> The UPnP endpoint URL /gena.cgi in the D-Link DIR-859 Wi-Fi router 1.05 and 1.06B01 Beta01 allows an Unauthenticated remote attacker to execute system commands as root, by sending a specially crafted HTTP SUBSCRIBE request to the UPnP service when connecting to the local network.
- [Ler2sq/CVE-2019-17621](https://github.com/Ler2sq/CVE-2019-17621)	<img alt="forks" src="https://img.shields.io/github/forks/Ler2sq/CVE-2019-17621">	<img alt="stars" src="https://img.shields.io/github/stars/Ler2sq/CVE-2019-17621">
- [s1kr10s/D-Link-DIR-859-RCE](https://github.com/s1kr10s/D-Link-DIR-859-RCE)	<img alt="forks" src="https://img.shields.io/github/forks/s1kr10s/D-Link-DIR-859-RCE">	<img alt="stars" src="https://img.shields.io/github/stars/s1kr10s/D-Link-DIR-859-RCE">

---
## CVE-2019-17577 (2019-10-16T18:15:00)
> An issue was discovered in Dolibarr 10.0.2. It has XSS via the "outgoing email setup" feature in the admin/mails.php?action=edit URI via the "Email used for error returns emails (fields 'Errors-To' in emails sent)" field.
- [Live-Hack-CVE/CVE-2019-17577](https://github.com/Live-Hack-CVE/CVE-2019-17577)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17577">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17577">

---
## CVE-2019-17571 (2019-12-20T17:15:00)
> Included in Log4j 1.2 is a SocketServer class that is vulnerable to deserialization of untrusted data which can be exploited to remotely execute arbitrary code when combined with a deserialization gadget when listening to untrusted network traffic for log data. This affects Log4j versions up to 1.2 up to 1.2.17.
- [Live-Hack-CVE/CVE-2019-17571](https://github.com/Live-Hack-CVE/CVE-2019-17571)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17571">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17571">
- [hillu/local-log4j-vuln-scanner](https://github.com/hillu/local-log4j-vuln-scanner)	<img alt="forks" src="https://img.shields.io/github/forks/hillu/local-log4j-vuln-scanner">	<img alt="stars" src="https://img.shields.io/github/stars/hillu/local-log4j-vuln-scanner">
- [HynekPetrak/log4shell-finder](https://github.com/HynekPetrak/log4shell-finder)	<img alt="forks" src="https://img.shields.io/github/forks/HynekPetrak/log4shell-finder">	<img alt="stars" src="https://img.shields.io/github/stars/HynekPetrak/log4shell-finder">
- [Al1ex/CVE-2019-17571](https://github.com/Al1ex/CVE-2019-17571)	<img alt="forks" src="https://img.shields.io/github/forks/Al1ex/CVE-2019-17571">	<img alt="stars" src="https://img.shields.io/github/stars/Al1ex/CVE-2019-17571">
- [shadow-horse/CVE-2019-17571](https://github.com/shadow-horse/CVE-2019-17571)	<img alt="forks" src="https://img.shields.io/github/forks/shadow-horse/CVE-2019-17571">	<img alt="stars" src="https://img.shields.io/github/stars/shadow-horse/CVE-2019-17571">
- [Live-Hack-CVE/CVE-2019-17571](https://github.com/Live-Hack-CVE/CVE-2019-17571)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17571">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17571">

---
## CVE-2019-17570 (2020-01-23T22:15:00)
> An untrusted deserialization was found in the org.apache.xmlrpc.parser.XmlRpcResponseParser:addResult method of Apache XML-RPC (aka ws-xmlrpc) library. A malicious XML-RPC server could target a XML-RPC client causing it to execute arbitrary code. Apache XML-RPC is no longer maintained and this issue will not be fixed.
- [Live-Hack-CVE/CVE-2019-17570](https://github.com/Live-Hack-CVE/CVE-2019-17570)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17570">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17570">
- [r00t4dm/CVE-2019-17570](https://github.com/r00t4dm/CVE-2019-17570)	<img alt="forks" src="https://img.shields.io/github/forks/r00t4dm/CVE-2019-17570">	<img alt="stars" src="https://img.shields.io/github/stars/r00t4dm/CVE-2019-17570">
- [fbeasts/xmlrpc-common-deserialization](https://github.com/fbeasts/xmlrpc-common-deserialization)	<img alt="forks" src="https://img.shields.io/github/forks/fbeasts/xmlrpc-common-deserialization">	<img alt="stars" src="https://img.shields.io/github/stars/fbeasts/xmlrpc-common-deserialization">

---
## CVE-2019-17569 (2020-02-24T22:15:00)
> The refactoring present in Apache Tomcat 9.0.28 to 9.0.30, 8.5.48 to 8.5.50 and 7.0.98 to 7.0.99 introduced a regression. The result of the regression was that invalid Transfer-Encoding headers were incorrectly processed leading to a possibility of HTTP Request Smuggling if Tomcat was located behind a reverse proxy that incorrectly handled the invalid Transfer-Encoding header in a particular manner. Such a reverse proxy is considered unlikely.
- [Live-Hack-CVE/CVE-2019-17569](https://github.com/Live-Hack-CVE/CVE-2019-17569)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17569">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17569">

---
## CVE-2019-17565 (2020-03-23T22:15:00)
> There is a vulnerability in Apache Traffic Server 6.0.0 to 6.2.3, 7.0.0 to 7.1.8, and 8.0.0 to 8.0.5 with a smuggling attack and chunked encoding. Upgrade to versions 7.1.9 and 8.0.6 or later versions.
- [Live-Hack-CVE/CVE-2019-17565](https://github.com/Live-Hack-CVE/CVE-2019-17565)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17565">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17565">

---
## CVE-2019-17563 (2019-12-23T17:15:00)
> When using FORM authentication with Apache Tomcat 9.0.0.M1 to 9.0.29, 8.5.0 to 8.5.49 and 7.0.0 to 7.0.98 there was a narrow window where an attacker could perform a session fixation attack. The window was considered too narrow for an exploit to be practical but, erring on the side of caution, this issue has been treated as a security vulnerability.
- [Live-Hack-CVE/CVE-2019-17563](https://github.com/Live-Hack-CVE/CVE-2019-17563)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17563">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17563">

---
## CVE-2019-17559 (2020-03-23T22:15:00)
> There is a vulnerability in Apache Traffic Server 6.0.0 to 6.2.3, 7.0.0 to 7.1.8, and 8.0.0 to 8.0.5 with a smuggling attack and scheme parsing. Upgrade to versions 7.1.9 and 8.0.6 or later versions.
- [Live-Hack-CVE/CVE-2019-17559](https://github.com/Live-Hack-CVE/CVE-2019-17559)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17559">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17559">

---
## CVE-2019-17519 (2020-02-12T19:15:00)
> The Bluetooth Low Energy implementation on NXP SDK through 2.2.1 for KW41Z devices does not properly restrict the Link Layer payload length, allowing attackers in radio range to cause a buffer overflow via a crafted packet.
- [Live-Hack-CVE/CVE-2019-17519](https://github.com/Live-Hack-CVE/CVE-2019-17519)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17519">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17519">

---
## CVE-2019-17498 (2019-10-21T22:15:00)
> In libssh2 v1.9.0 and earlier versions, the SSH_MSG_DISCONNECT logic in packet.c has an integer overflow in a bounds check, enabling an attacker to specify an arbitrary (out-of-bounds) offset for a subsequent memory read. A crafted SSH server may be able to disclose sensitive information or cause a denial of service condition on the client system when a user connects to the server.
- [Live-Hack-CVE/CVE-2019-17498](https://github.com/Live-Hack-CVE/CVE-2019-17498)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17498">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17498">

---
## CVE-2019-17223 (2019-10-15T12:15:00)
> There is HTML Injection in the Note field in Dolibarr ERP/CRM 10.0.2 via user/note.php.
- [Live-Hack-CVE/CVE-2019-17223](https://github.com/Live-Hack-CVE/CVE-2019-17223)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17223">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17223">

---
## CVE-2019-17060 (2020-02-10T21:51:00)
> The Bluetooth Low Energy (BLE) stack implementation on the NXP KW41Z (based on the MCUXpresso SDK with Bluetooth Low Energy Driver 2.2.1 and earlier) does not properly restrict the BLE Link Layer header and executes certain memory contents upon receiving a packet with a Link Layer ID (LLID) equal to zero. This allows attackers within radio range to cause deadlocks, cause anomalous behavior in the BLE state machine, or trigger a buffer overflow via a crafted BLE Link Layer frame.
- [Live-Hack-CVE/CVE-2019-17060](https://github.com/Live-Hack-CVE/CVE-2019-17060)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17060">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17060">

---
## CVE-2019-17052 (2019-10-01T14:15:00)
> ax25_create in net/ax25/af_ax25.c in the AF_AX25 network module in the Linux kernel 3.16 through 5.3.2 does not enforce CAP_NET_RAW, which means that unprivileged users can create a raw socket, aka CID-0614e2b73768.
- [Live-Hack-CVE/CVE-2019-17052](https://github.com/Live-Hack-CVE/CVE-2019-17052)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17052">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17052">

---
## CVE-2019-17026 (2020-03-02T05:15:00)
> Incorrect alias information in IonMonkey JIT compiler for setting array elements could lead to a type confusion. We are aware of targeted attacks in the wild abusing this flaw. This vulnerability affects Firefox ESR < 68.4.1, Thunderbird < 68.4.1, and Firefox < 72.0.1.
- [Live-Hack-CVE/CVE-2019-17026](https://github.com/Live-Hack-CVE/CVE-2019-17026)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-17026">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-17026">

---
## CVE-2019-16891 (2019-10-04T14:15:00)
> Liferay Portal CE 6.2.5 allows remote command execution because of deserialization of a JSON payload.
- [Live-Hack-CVE/CVE-2019-16891](https://github.com/Live-Hack-CVE/CVE-2019-16891)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-16891">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-16891">

---
## CVE-2019-16770 (2019-12-05T20:15:00)
> In Puma before versions 3.12.2 and 4.3.1, a poorly-behaved client could use keepalive requests to monopolize Puma's reactor and create a denial of service attack. If more keepalive connections to Puma are opened than there are threads available, additional connections will wait permanently if the attacker sends requests frequently enough. This vulnerability is patched in Puma 4.3.1 and 3.12.2.
- [Live-Hack-CVE/CVE-2019-16770](https://github.com/Live-Hack-CVE/CVE-2019-16770)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-16770">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-16770">

---
## CVE-2019-1649 (2019-05-13T19:29:00)
> A vulnerability in the logic that handles access control to one of the hardware components in Cisco's proprietary Secure Boot implementation could allow an authenticated, local attacker to write a modified firmware image to the component. This vulnerability affects multiple Cisco products that support hardware-based Secure Boot functionality. The vulnerability is due to an improper check on the area of code that manages on-premise updates to a Field Programmable Gate Array (FPGA) part of the Secure Boot hardware implementation. An attacker with elevated privileges and access to the underlying operating system that is running on the affected device could exploit this vulnerability by writing a modified firmware image to the FPGA. A successful exploit could either cause the device to become unusable (and require a hardware replacement) or allow tampering with the Secure Boot verification process, which under some circumstances may allow the attacker to install and boot a malicious software image. An attacker will need to fulfill all the following conditions to attempt to exploit this vulnerability: Have privileged administrative access to the device. Be able to access the underlying operating system running on the device; this can be achieved either by using a supported, documented mechanism or by exploiting another vulnerability that would provide an attacker with such access. Develop or have access to a platform-specific exploit. An attacker attempting to exploit this vulnerability across multiple affected platforms would need to research each one of those platforms and then develop a platform-specific exploit. Although the research process could be reused across different platforms, an exploit developed for a given hardware platform is unlikely to work on a different hardware platform.
- [Live-Hack-CVE/CVE-2019-1649](https://github.com/Live-Hack-CVE/CVE-2019-1649)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-1649">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-1649">

---
## CVE-2019-16278 (2019-10-14T17:15:00)
> Directory Traversal in the function http_verify in nostromo nhttpd through 1.9.6 allows an attacker to achieve remote code execution via a crafted HTTP request.
- [keshiba/cve-2019-16278](https://github.com/keshiba/cve-2019-16278)	<img alt="forks" src="https://img.shields.io/github/forks/keshiba/cve-2019-16278">	<img alt="stars" src="https://img.shields.io/github/stars/keshiba/cve-2019-16278">

---
## CVE-2019-16223 (2019-09-11T14:15:00)
> WordPress before 5.2.3 allows XSS in post previews by authenticated users.
- [Live-Hack-CVE/CVE-2019-16223](https://github.com/Live-Hack-CVE/CVE-2019-16223)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-16223">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-16223">

---
## CVE-2019-16167 (2019-09-09T17:15:00)
> sysstat before 12.1.6 has memory corruption due to an Integer Overflow in remap_struct() in sa_common.c.
- [Live-Hack-CVE/CVE-2019-16167](https://github.com/Live-Hack-CVE/CVE-2019-16167)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-16167">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-16167">

---
## CVE-2019-15605 (2020-02-07T15:15:00)
> HTTP request smuggling in Node.js 10, 12, and 13 causes malicious payload delivery when transfer-encoding is malformed
- [Live-Hack-CVE/CVE-2019-15605](https://github.com/Live-Hack-CVE/CVE-2019-15605)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-15605">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-15605">

---
## CVE-2019-15604 (2020-02-07T15:15:00)
> Improper Certificate Validation in Node.js 10, 12, and 13 causes the process to abort when sending a crafted X.509 certificate
- [Live-Hack-CVE/CVE-2019-15604](https://github.com/Live-Hack-CVE/CVE-2019-15604)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-15604">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-15604">

---
## CVE-2019-15514 (2019-08-23T13:15:00)
> The Privacy > Phone Number feature in the Telegram app 5.10 for Android and iOS provides an incorrect indication that the access level is Nobody, because attackers can find these numbers via the Group Info feature, e.g., by adding a significant fraction of a region's assigned phone numbers.
- [graysuit/CVE-2019-15514](https://github.com/graysuit/CVE-2019-15514)	<img alt="forks" src="https://img.shields.io/github/forks/graysuit/CVE-2019-15514">	<img alt="stars" src="https://img.shields.io/github/stars/graysuit/CVE-2019-15514">

---
## CVE-2019-15107 (2019-08-16T03:15:00)
> An issue was discovered in Webmin <=1.920. The parameter old in password_change.cgi contains a command injection vulnerability.
- [psw01/CVE-2019-15107_webminRCE](https://github.com/psw01/CVE-2019-15107_webminRCE)	<img alt="forks" src="https://img.shields.io/github/forks/psw01/CVE-2019-15107_webminRCE">	<img alt="stars" src="https://img.shields.io/github/stars/psw01/CVE-2019-15107_webminRCE">
- [NullBrunk/CVE-2019-15107](https://github.com/NullBrunk/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/NullBrunk/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/NullBrunk/CVE-2019-15107">
- [MuirlandOracle/CVE-2019-15107](https://github.com/MuirlandOracle/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/MuirlandOracle/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/MuirlandOracle/CVE-2019-15107">
- [trhacknon/CVE-2019-15107](https://github.com/trhacknon/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/trhacknon/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/trhacknon/CVE-2019-15107">
- [merlin-ke/CVE_2019_15107](https://github.com/merlin-ke/CVE_2019_15107)	<img alt="forks" src="https://img.shields.io/github/forks/merlin-ke/CVE_2019_15107">	<img alt="stars" src="https://img.shields.io/github/stars/merlin-ke/CVE_2019_15107">
- [f0rkr/CVE-2019-15107](https://github.com/f0rkr/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/f0rkr/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/f0rkr/CVE-2019-15107">
- [hacknotes/CVE-2019-15107-Exploit](https://github.com/hacknotes/CVE-2019-15107-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/hacknotes/CVE-2019-15107-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/hacknotes/CVE-2019-15107-Exploit">
- [Tuz-Wwsd/CVE-2019-15107_detection](https://github.com/Tuz-Wwsd/CVE-2019-15107_detection)	<img alt="forks" src="https://img.shields.io/github/forks/Tuz-Wwsd/CVE-2019-15107_detection">	<img alt="stars" src="https://img.shields.io/github/stars/Tuz-Wwsd/CVE-2019-15107_detection">
- [whokilleddb/CVE-2019-15107](https://github.com/whokilleddb/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/whokilleddb/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/whokilleddb/CVE-2019-15107">
- [darrenmartyn/CVE-2019-15107](https://github.com/darrenmartyn/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/darrenmartyn/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/darrenmartyn/CVE-2019-15107">
- [puckiestyle/CVE-2019-15107](https://github.com/puckiestyle/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/puckiestyle/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/puckiestyle/CVE-2019-15107">
- [cdedmondson/Modified-CVE-2019-15107](https://github.com/cdedmondson/Modified-CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/cdedmondson/Modified-CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/cdedmondson/Modified-CVE-2019-15107">
- [diegojuan/CVE-2019-15107](https://github.com/diegojuan/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/diegojuan/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/diegojuan/CVE-2019-15107">
- [n0obit4/Webmin_1.890-POC](https://github.com/n0obit4/Webmin_1.890-POC)	<img alt="forks" src="https://img.shields.io/github/forks/n0obit4/Webmin_1.890-POC">	<img alt="stars" src="https://img.shields.io/github/stars/n0obit4/Webmin_1.890-POC">
- [squid22/Webmin_CVE-2019-15107](https://github.com/squid22/Webmin_CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/squid22/Webmin_CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/squid22/Webmin_CVE-2019-15107">
- [ruthvikvegunta/CVE-2019-15107](https://github.com/ruthvikvegunta/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/ruthvikvegunta/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/ruthvikvegunta/CVE-2019-15107">
- [ChakoMoonFish/webmin_CVE-2019-15107](https://github.com/ChakoMoonFish/webmin_CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/ChakoMoonFish/webmin_CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/ChakoMoonFish/webmin_CVE-2019-15107">
- [hannob/webminex](https://github.com/hannob/webminex)	<img alt="forks" src="https://img.shields.io/github/forks/hannob/webminex">	<img alt="stars" src="https://img.shields.io/github/stars/hannob/webminex">
- [ianxtianxt/CVE-2019-15107](https://github.com/ianxtianxt/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/ianxtianxt/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/ianxtianxt/CVE-2019-15107">
- [AleWong/WebminRCE-EXP-CVE-2019-15107-](https://github.com/AleWong/WebminRCE-EXP-CVE-2019-15107-)	<img alt="forks" src="https://img.shields.io/github/forks/AleWong/WebminRCE-EXP-CVE-2019-15107-">	<img alt="stars" src="https://img.shields.io/github/stars/AleWong/WebminRCE-EXP-CVE-2019-15107-">
- [Rayferrufino/Make-and-Break](https://github.com/Rayferrufino/Make-and-Break)	<img alt="forks" src="https://img.shields.io/github/forks/Rayferrufino/Make-and-Break">	<img alt="stars" src="https://img.shields.io/github/stars/Rayferrufino/Make-and-Break">
- [g0db0x/CVE_2019_15107](https://github.com/g0db0x/CVE_2019_15107)	<img alt="forks" src="https://img.shields.io/github/forks/g0db0x/CVE_2019_15107">	<img alt="stars" src="https://img.shields.io/github/stars/g0db0x/CVE_2019_15107">
- [Pichuuuuu/CVE-2019-15107](https://github.com/Pichuuuuu/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/Pichuuuuu/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/Pichuuuuu/CVE-2019-15107">
- [ketlerd/CVE-2019-15107](https://github.com/ketlerd/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/ketlerd/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/ketlerd/CVE-2019-15107">
- [AdministratorGithub/CVE-2019-15107](https://github.com/AdministratorGithub/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/AdministratorGithub/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/AdministratorGithub/CVE-2019-15107">
- [HACHp1/webmin_docker_and_exp](https://github.com/HACHp1/webmin_docker_and_exp)	<img alt="forks" src="https://img.shields.io/github/forks/HACHp1/webmin_docker_and_exp">	<img alt="stars" src="https://img.shields.io/github/stars/HACHp1/webmin_docker_and_exp">
- [TheAlpha19/MiniExploit](https://github.com/TheAlpha19/MiniExploit)	<img alt="forks" src="https://img.shields.io/github/forks/TheAlpha19/MiniExploit">	<img alt="stars" src="https://img.shields.io/github/stars/TheAlpha19/MiniExploit">
- [jas502n/CVE-2019-15107](https://github.com/jas502n/CVE-2019-15107)	<img alt="forks" src="https://img.shields.io/github/forks/jas502n/CVE-2019-15107">	<img alt="stars" src="https://img.shields.io/github/stars/jas502n/CVE-2019-15107">
- [kh4sh3i/Webmin-CVE](https://github.com/kh4sh3i/Webmin-CVE)	<img alt="forks" src="https://img.shields.io/github/forks/kh4sh3i/Webmin-CVE">	<img alt="stars" src="https://img.shields.io/github/stars/kh4sh3i/Webmin-CVE">

---
## CVE-2019-14907 (2020-01-21T18:15:00)
> All samba versions 4.9.x before 4.9.18, 4.10.x before 4.10.12 and 4.11.x before 4.11.5 have an issue where if it is set with "log level = 3" (or above) then the string obtained from the client, after a failed character conversion, is printed. Such strings can be provided during the NTLMSSP authentication exchange. In the Samba AD DC in particular, this may cause a long-lived process(such as the RPC server) to terminate. (In the file server case, the most likely target, smbd, operates as process-per-client and so a crash there is harmless).
- [Live-Hack-CVE/CVE-2019-14907](https://github.com/Live-Hack-CVE/CVE-2019-14907)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-14907">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-14907">

---
## CVE-2019-14591 (2019-11-14T20:15:00)
> Improper input validation in the API for Intel(R) Graphics Driver versions before 26.20.100.7209 may allow an authenticated user to potentially enable denial of service via local access.
- [Live-Hack-CVE/CVE-2019-14591](https://github.com/Live-Hack-CVE/CVE-2019-14591)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-14591">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-14591">

---
## CVE-2019-14590 (2019-11-14T20:15:00)
> Improper access control in the API for the Intel(R) Graphics Driver versions before 26.20.100.7209 may allow an authenticated user to potentially enable information disclosure via local access.
- [Live-Hack-CVE/CVE-2019-14590](https://github.com/Live-Hack-CVE/CVE-2019-14590)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-14590">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-14590">

---
## CVE-2019-14574 (2019-11-14T20:15:00)
> Out of bounds read in a subsystem for Intel(R) Graphics Driver versions before 26.20.100.7209 may allow an authenticated user to potentially enable denial of service via local access.
- [Live-Hack-CVE/CVE-2019-14574](https://github.com/Live-Hack-CVE/CVE-2019-14574)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-14574">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-14574">

---
## CVE-2019-14494 (2019-08-01T17:15:00)
> An issue was discovered in Poppler through 0.78.0. There is a divide-by-zero error in the function SplashOutputDev::tilingPatternFill at SplashOutputDev.cc.
- [Live-Hack-CVE/CVE-2019-14494](https://github.com/Live-Hack-CVE/CVE-2019-14494)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-14494">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-14494">

---
## CVE-2019-14287 (2019-10-17T18:15:00)
> In Sudo before 1.8.28, an attacker with access to a Runas ALL sudoer account can bypass certain policy blacklists and session PAM modules, and can cause incorrect logging, by invoking sudo with a crafted user ID. For example, this allows bypass of !root configuration, and USER= logging, for a "sudo -u \#$((0xffffffff))" command.
- [MariliaMeira/CVE-2019-14287](https://github.com/MariliaMeira/CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/MariliaMeira/CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/MariliaMeira/CVE-2019-14287">
- [shyambhanushali/AttackDefendExercise](https://github.com/shyambhanushali/AttackDefendExercise)	<img alt="forks" src="https://img.shields.io/github/forks/shyambhanushali/AttackDefendExercise">	<img alt="stars" src="https://img.shields.io/github/stars/shyambhanushali/AttackDefendExercise">
- [Hasintha-98/Sudo-Vulnerability-Exploit-CVE-2019-14287](https://github.com/Hasintha-98/Sudo-Vulnerability-Exploit-CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/Hasintha-98/Sudo-Vulnerability-Exploit-CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/Hasintha-98/Sudo-Vulnerability-Exploit-CVE-2019-14287">
- [k4u5h41/CVE-2019-14287](https://github.com/k4u5h41/CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/k4u5h41/CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/k4u5h41/CVE-2019-14287">
- [DularaAnushka/Linux-Privilege-Escalation-using-Sudo-Rights](https://github.com/DularaAnushka/Linux-Privilege-Escalation-using-Sudo-Rights)	<img alt="forks" src="https://img.shields.io/github/forks/DularaAnushka/Linux-Privilege-Escalation-using-Sudo-Rights">	<img alt="stars" src="https://img.shields.io/github/stars/DularaAnushka/Linux-Privilege-Escalation-using-Sudo-Rights">
- [edsonjt81/CVE-2019-14287-](https://github.com/edsonjt81/CVE-2019-14287-)	<img alt="forks" src="https://img.shields.io/github/forks/edsonjt81/CVE-2019-14287-">	<img alt="stars" src="https://img.shields.io/github/stars/edsonjt81/CVE-2019-14287-">
- [CashWilliams/CVE-2019-14287-demo](https://github.com/CashWilliams/CVE-2019-14287-demo)	<img alt="forks" src="https://img.shields.io/github/forks/CashWilliams/CVE-2019-14287-demo">	<img alt="stars" src="https://img.shields.io/github/stars/CashWilliams/CVE-2019-14287-demo">
- [M108Falcon/Sudo-CVE-2019-14287](https://github.com/M108Falcon/Sudo-CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/M108Falcon/Sudo-CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/M108Falcon/Sudo-CVE-2019-14287">
- [shallvhack/Sudo-Security-Bypass-CVE-2019-14287](https://github.com/shallvhack/Sudo-Security-Bypass-CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/shallvhack/Sudo-Security-Bypass-CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/shallvhack/Sudo-Security-Bypass-CVE-2019-14287">
- [Tharana/vulnerability-exploitation](https://github.com/Tharana/vulnerability-exploitation)	<img alt="forks" src="https://img.shields.io/github/forks/Tharana/vulnerability-exploitation">	<img alt="stars" src="https://img.shields.io/github/stars/Tharana/vulnerability-exploitation">
- [janod313/-CVE-2019-14287-SUDO-bypass-vulnerability](https://github.com/janod313/-CVE-2019-14287-SUDO-bypass-vulnerability)	<img alt="forks" src="https://img.shields.io/github/forks/janod313/-CVE-2019-14287-SUDO-bypass-vulnerability">	<img alt="stars" src="https://img.shields.io/github/stars/janod313/-CVE-2019-14287-SUDO-bypass-vulnerability">
- [DewmiApsara/CVE-2019-14287](https://github.com/DewmiApsara/CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/DewmiApsara/CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/DewmiApsara/CVE-2019-14287">
- [HussyCool/CVE-2019-14287-IT18030372-](https://github.com/HussyCool/CVE-2019-14287-IT18030372-)	<img alt="forks" src="https://img.shields.io/github/forks/HussyCool/CVE-2019-14287-IT18030372-">	<img alt="stars" src="https://img.shields.io/github/stars/HussyCool/CVE-2019-14287-IT18030372-">
- [thinuri99/Sudo-Security-Bypass-Vulnerability-CVE-2019-14287-](https://github.com/thinuri99/Sudo-Security-Bypass-Vulnerability-CVE-2019-14287-)	<img alt="forks" src="https://img.shields.io/github/forks/thinuri99/Sudo-Security-Bypass-Vulnerability-CVE-2019-14287-">	<img alt="stars" src="https://img.shields.io/github/stars/thinuri99/Sudo-Security-Bypass-Vulnerability-CVE-2019-14287-">
- [ejlevin99/Sudo-Security-Bypass-Vulnerability](https://github.com/ejlevin99/Sudo-Security-Bypass-Vulnerability)	<img alt="forks" src="https://img.shields.io/github/forks/ejlevin99/Sudo-Security-Bypass-Vulnerability">	<img alt="stars" src="https://img.shields.io/github/stars/ejlevin99/Sudo-Security-Bypass-Vulnerability">
- [ShianTrish/sudo-Security-Bypass-vulnerability-CVE-2019-14287](https://github.com/ShianTrish/sudo-Security-Bypass-vulnerability-CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/ShianTrish/sudo-Security-Bypass-vulnerability-CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/ShianTrish/sudo-Security-Bypass-vulnerability-CVE-2019-14287">
- [SachinthaDeSilva-cmd/Exploit-CVE-2019-14287](https://github.com/SachinthaDeSilva-cmd/Exploit-CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/SachinthaDeSilva-cmd/Exploit-CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/SachinthaDeSilva-cmd/Exploit-CVE-2019-14287">
- [Tharana/Exploiting-a-Linux-kernel-vulnerability](https://github.com/Tharana/Exploiting-a-Linux-kernel-vulnerability)	<img alt="forks" src="https://img.shields.io/github/forks/Tharana/Exploiting-a-Linux-kernel-vulnerability">	<img alt="stars" src="https://img.shields.io/github/stars/Tharana/Exploiting-a-Linux-kernel-vulnerability">
- [CMNatic/Dockerized-CVE-2019-14287](https://github.com/CMNatic/Dockerized-CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/CMNatic/Dockerized-CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/CMNatic/Dockerized-CVE-2019-14287">
- [Sindayifu/CVE-2019-14287-CVE-2014-6271](https://github.com/Sindayifu/CVE-2019-14287-CVE-2014-6271)	<img alt="forks" src="https://img.shields.io/github/forks/Sindayifu/CVE-2019-14287-CVE-2014-6271">	<img alt="stars" src="https://img.shields.io/github/stars/Sindayifu/CVE-2019-14287-CVE-2014-6271">
- [axax002/sudo-vulnerability-CVE-2019-14287](https://github.com/axax002/sudo-vulnerability-CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/axax002/sudo-vulnerability-CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/axax002/sudo-vulnerability-CVE-2019-14287">
- [kumar1100/CVE2019-14287](https://github.com/kumar1100/CVE2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/kumar1100/CVE2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/kumar1100/CVE2019-14287">
- [huang919/cve-2019-14287-PPT](https://github.com/huang919/cve-2019-14287-PPT)	<img alt="forks" src="https://img.shields.io/github/forks/huang919/cve-2019-14287-PPT">	<img alt="stars" src="https://img.shields.io/github/stars/huang919/cve-2019-14287-PPT">
- [Sindadziy/cve-2019-14287](https://github.com/Sindadziy/cve-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/Sindadziy/cve-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/Sindadziy/cve-2019-14287">
- [wenyu1999/sudo-](https://github.com/wenyu1999/sudo-)	<img alt="forks" src="https://img.shields.io/github/forks/wenyu1999/sudo-">	<img alt="stars" src="https://img.shields.io/github/stars/wenyu1999/sudo-">
- [5l1v3r1/cve-2019-14287sudoexp](https://github.com/5l1v3r1/cve-2019-14287sudoexp)	<img alt="forks" src="https://img.shields.io/github/forks/5l1v3r1/cve-2019-14287sudoexp">	<img alt="stars" src="https://img.shields.io/github/stars/5l1v3r1/cve-2019-14287sudoexp">
- [gurneesh/CVE-2019-14287-write-up](https://github.com/gurneesh/CVE-2019-14287-write-up)	<img alt="forks" src="https://img.shields.io/github/forks/gurneesh/CVE-2019-14287-write-up">	<img alt="stars" src="https://img.shields.io/github/stars/gurneesh/CVE-2019-14287-write-up">
- [n0w4n/CVE-2019-14287](https://github.com/n0w4n/CVE-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/n0w4n/CVE-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/n0w4n/CVE-2019-14287">
- [FauxFaux/sudo-cve-2019-14287](https://github.com/FauxFaux/sudo-cve-2019-14287)	<img alt="forks" src="https://img.shields.io/github/forks/FauxFaux/sudo-cve-2019-14287">	<img alt="stars" src="https://img.shields.io/github/stars/FauxFaux/sudo-cve-2019-14287">

---
## CVE-2019-13990 (2019-07-26T19:15:00)
> initDocumentParser in xml/XMLSchedulingDataProcessor.java in Terracotta Quartz Scheduler through 2.3.0 allows XXE attacks via a job description.
- [Live-Hack-CVE/CVE-2019-13990](https://github.com/Live-Hack-CVE/CVE-2019-13990)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-13990">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-13990">

---
## CVE-2019-13120 (2019-10-07T22:15:00)
> Amazon FreeRTOS up to and including v1.4.8 lacks length checking in prvProcessReceivedPublish, resulting in untargetable leakage of arbitrary memory contents on a device to an attacker. If an attacker has the authorization to send a malformed MQTT publish packet to an Amazon IoT Thing, which interacts with an associated vulnerable MQTT message in the application, specific circumstances could trigger this vulnerability.
- [Live-Hack-CVE/CVE-2019-13120](https://github.com/Live-Hack-CVE/CVE-2019-13120)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-13120">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-13120">

---
## CVE-2019-12874 (2019-06-18T18:15:00)
> An issue was discovered in zlib_decompress_extra in modules/demux/mkv/util.cpp in VideoLAN VLC media player 3.x through 3.0.7. The Matroska demuxer, while parsing a malformed MKV file type, has a double free.
- [ahaShiyu/CVE-2019-12874](https://github.com/ahaShiyu/CVE-2019-12874)	<img alt="forks" src="https://img.shields.io/github/forks/ahaShiyu/CVE-2019-12874">	<img alt="stars" src="https://img.shields.io/github/stars/ahaShiyu/CVE-2019-12874">

---
## CVE-2019-12836 (2019-06-21T15:15:00)
> The Bobronix JEditor editor before 3.0.6 for Jira allows an attacker to add a URL/Link (to an existing issue) that can cause forgery of a request to an out-of-origin domain. This in turn may allow for a forged request that can be invoked in the context of an authenticated user, leading to stealing of session tokens and account takeover.
- [9lyph/CVE-2019-12836](https://github.com/9lyph/CVE-2019-12836)	<img alt="forks" src="https://img.shields.io/github/forks/9lyph/CVE-2019-12836">	<img alt="stars" src="https://img.shields.io/github/stars/9lyph/CVE-2019-12836">

---
## CVE-2019-12762 (2019-06-06T20:29:00)
> Xiaomi Mi 5s Plus devices allow attackers to trigger touchscreen anomalies via a radio signal between 198 kHz and 203 kHz, as demonstrated by a transmitter and antenna hidden just beneath the surface of a coffee-shop table, aka Ghost Touch.
- [Live-Hack-CVE/CVE-2019-12762](https://github.com/Live-Hack-CVE/CVE-2019-12762)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-12762">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-12762">

---
## CVE-2019-12735 (2019-06-05T14:29:00)
> getchar.c in Vim before 8.1.1365 and Neovim before 0.3.6 allows remote attackers to execute arbitrary OS commands via the :source! command in a modeline, as demonstrated by execute in Vim, and assert_fails or nvim_input in Neovim.
- [st9007a/CVE-2019-12735](https://github.com/st9007a/CVE-2019-12735)	<img alt="forks" src="https://img.shields.io/github/forks/st9007a/CVE-2019-12735">	<img alt="stars" src="https://img.shields.io/github/stars/st9007a/CVE-2019-12735">
- [nickylimjj/cve-2019-12735](https://github.com/nickylimjj/cve-2019-12735)	<img alt="forks" src="https://img.shields.io/github/forks/nickylimjj/cve-2019-12735">	<img alt="stars" src="https://img.shields.io/github/stars/nickylimjj/cve-2019-12735">
- [oldthree3/CVE-2019-12735-VIM-NEOVIM](https://github.com/oldthree3/CVE-2019-12735-VIM-NEOVIM)	<img alt="forks" src="https://img.shields.io/github/forks/oldthree3/CVE-2019-12735-VIM-NEOVIM">	<img alt="stars" src="https://img.shields.io/github/stars/oldthree3/CVE-2019-12735-VIM-NEOVIM">
- [pcy190/ace-vim-neovim](https://github.com/pcy190/ace-vim-neovim)	<img alt="forks" src="https://img.shields.io/github/forks/pcy190/ace-vim-neovim">	<img alt="stars" src="https://img.shields.io/github/stars/pcy190/ace-vim-neovim">

---
## CVE-2019-12257 (2019-08-09T18:15:00)
> Wind River VxWorks 6.6 through 6.9 has a Buffer Overflow in the DHCP client component. There is an IPNET security vulnerability: Heap overflow in DHCP Offer/ACK parsing inside ipdhcpc.
- [Live-Hack-CVE/CVE-2019-12257](https://github.com/Live-Hack-CVE/CVE-2019-12257)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-12257">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-12257">

---
## CVE-2019-12256 (2019-08-09T18:15:00)
> Wind River VxWorks 6.9 and vx7 has a Buffer Overflow in the IPv4 component. There is an IPNET security vulnerability: Stack overflow in the parsing of IPv4 packets’ IP options.
- [Live-Hack-CVE/CVE-2019-12256](https://github.com/Live-Hack-CVE/CVE-2019-12256)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-12256">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-12256">
- [sud0woodo/Urgent11-Suricata-LUA-scripts](https://github.com/sud0woodo/Urgent11-Suricata-LUA-scripts)	<img alt="forks" src="https://img.shields.io/github/forks/sud0woodo/Urgent11-Suricata-LUA-scripts">	<img alt="stars" src="https://img.shields.io/github/stars/sud0woodo/Urgent11-Suricata-LUA-scripts">

---
## CVE-2019-1205 (2019-08-14T21:15:00)
> A remote code execution vulnerability exists in Microsoft Word software when it fails to properly handle objects in memory, aka 'Microsoft Word Remote Code Execution Vulnerability'. This CVE ID is unique from CVE-2019-1201.
- [razordeveloper/CVE-2019-1205](https://github.com/razordeveloper/CVE-2019-1205)	<img alt="forks" src="https://img.shields.io/github/forks/razordeveloper/CVE-2019-1205">	<img alt="stars" src="https://img.shields.io/github/stars/razordeveloper/CVE-2019-1205">

---
## CVE-2019-11823 (2020-05-04T10:15:00)
> CRLF injection vulnerability in Network Center in Synology Router Manager (SRM) before 1.2.3-8017-2 allows remote attackers to cause a denial of service (out-of-bounds read and application crash) via crafted network traffic.
- [Live-Hack-CVE/CVE-2019-11823](https://github.com/Live-Hack-CVE/CVE-2019-11823)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11823">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11823">

---
## CVE-2019-11810 (2019-05-07T14:29:00)
> An issue was discovered in the Linux kernel before 5.0.7. A NULL pointer dereference can occur when megasas_create_frame_pool() fails in megasas_alloc_cmds() in drivers/scsi/megaraid/megaraid_sas_base.c. This causes a Denial of Service, related to a use-after-free.
- [Live-Hack-CVE/CVE-2019-11810](https://github.com/Live-Hack-CVE/CVE-2019-11810)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11810">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11810">

---
## CVE-2019-11510 (2019-05-08T17:29:00)
> In Pulse Secure Pulse Connect Secure (PCS) 8.2 before 8.2R12.1, 8.3 before 8.3R7.1, and 9.0 before 9.0R3.4, an unauthenticated remote attacker can send a specially crafted URI to perform an arbitrary file reading vulnerability .
- [trhacknon/CVE-2019-11510](https://github.com/trhacknon/CVE-2019-11510)	<img alt="forks" src="https://img.shields.io/github/forks/trhacknon/CVE-2019-11510">	<img alt="stars" src="https://img.shields.io/github/stars/trhacknon/CVE-2019-11510">
- [hackingyseguridad/directoriotraversal](https://github.com/hackingyseguridad/directoriotraversal)	<img alt="forks" src="https://img.shields.io/github/forks/hackingyseguridad/directoriotraversal">	<img alt="stars" src="https://img.shields.io/github/stars/hackingyseguridad/directoriotraversal">
- [pwn3z/CVE-2019-11510-PulseVPN](https://github.com/pwn3z/CVE-2019-11510-PulseVPN)	<img alt="forks" src="https://img.shields.io/github/forks/pwn3z/CVE-2019-11510-PulseVPN">	<img alt="stars" src="https://img.shields.io/github/stars/pwn3z/CVE-2019-11510-PulseVPN">
- [cisagov/check-your-pulse](https://github.com/cisagov/check-your-pulse)	<img alt="forks" src="https://img.shields.io/github/forks/cisagov/check-your-pulse">	<img alt="stars" src="https://img.shields.io/github/stars/cisagov/check-your-pulse">
- [andripwn/pulse-exploit](https://github.com/andripwn/pulse-exploit)	<img alt="forks" src="https://img.shields.io/github/forks/andripwn/pulse-exploit">	<img alt="stars" src="https://img.shields.io/github/stars/andripwn/pulse-exploit">
- [aqhmal/pulsexploit](https://github.com/aqhmal/pulsexploit)	<img alt="forks" src="https://img.shields.io/github/forks/aqhmal/pulsexploit">	<img alt="stars" src="https://img.shields.io/github/stars/aqhmal/pulsexploit">
- [BishopFox/pwn-pulse](https://github.com/BishopFox/pwn-pulse)	<img alt="forks" src="https://img.shields.io/github/forks/BishopFox/pwn-pulse">	<img alt="stars" src="https://img.shields.io/github/stars/BishopFox/pwn-pulse">
- [projectzeroindia/CVE-2019-11510](https://github.com/projectzeroindia/CVE-2019-11510)	<img alt="forks" src="https://img.shields.io/github/forks/projectzeroindia/CVE-2019-11510">	<img alt="stars" src="https://img.shields.io/github/stars/projectzeroindia/CVE-2019-11510">
- [sp4rkhunt3r/CVE2019-11510](https://github.com/sp4rkhunt3r/CVE2019-11510)	<img alt="forks" src="https://img.shields.io/github/forks/sp4rkhunt3r/CVE2019-11510">	<img alt="stars" src="https://img.shields.io/github/stars/sp4rkhunt3r/CVE2019-11510">
- [jason3e7/CVE-2019-11510](https://github.com/jason3e7/CVE-2019-11510)	<img alt="forks" src="https://img.shields.io/github/forks/jason3e7/CVE-2019-11510">	<img alt="stars" src="https://img.shields.io/github/stars/jason3e7/CVE-2019-11510">
- [jas502n/CVE-2019-11510-1](https://github.com/jas502n/CVE-2019-11510-1)	<img alt="forks" src="https://img.shields.io/github/forks/jas502n/CVE-2019-11510-1">	<img alt="stars" src="https://img.shields.io/github/stars/jas502n/CVE-2019-11510-1">
- [r00tpgp/http-pulse_ssl_vpn.nse](https://github.com/r00tpgp/http-pulse_ssl_vpn.nse)	<img alt="forks" src="https://img.shields.io/github/forks/r00tpgp/http-pulse_ssl_vpn.nse">	<img alt="stars" src="https://img.shields.io/github/stars/r00tpgp/http-pulse_ssl_vpn.nse">
- [es0/CVE-2019-11510_poc](https://github.com/es0/CVE-2019-11510_poc)	<img alt="forks" src="https://img.shields.io/github/forks/es0/CVE-2019-11510_poc">	<img alt="stars" src="https://img.shields.io/github/stars/es0/CVE-2019-11510_poc">
- [imjdl/CVE-2019-11510-poc](https://github.com/imjdl/CVE-2019-11510-poc)	<img alt="forks" src="https://img.shields.io/github/forks/imjdl/CVE-2019-11510-poc">	<img alt="stars" src="https://img.shields.io/github/stars/imjdl/CVE-2019-11510-poc">
- [nuc13us/Pulse](https://github.com/nuc13us/Pulse)	<img alt="forks" src="https://img.shields.io/github/forks/nuc13us/Pulse">	<img alt="stars" src="https://img.shields.io/github/stars/nuc13us/Pulse">
- [34zY/APT-Backpack](https://github.com/34zY/APT-Backpack)	<img alt="forks" src="https://img.shields.io/github/forks/34zY/APT-Backpack">	<img alt="stars" src="https://img.shields.io/github/stars/34zY/APT-Backpack">

---
## CVE-2019-11498 (2019-04-24T05:29:00)
> WavpackSetConfiguration64 in pack_utils.c in libwavpack.a in WavPack through 5.1.0 has a "Conditional jump or move depends on uninitialised value" condition, which might allow attackers to cause a denial of service (application crash) via a DFF file that lacks valid sample-rate data.
- [Live-Hack-CVE/CVE-2019-11498](https://github.com/Live-Hack-CVE/CVE-2019-11498)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11498">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11498">

---
## CVE-2019-11113 (2019-11-14T20:15:00)
> Buffer overflow in Kernel Mode module for Intel(R) Graphics Driver before version 25.20.100.6618 (DCH) or 21.20.x.5077 (aka15.45.5077) may allow a privileged user to potentially enable information disclosure via local access.
- [Live-Hack-CVE/CVE-2019-11113](https://github.com/Live-Hack-CVE/CVE-2019-11113)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11113">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11113">

---
## CVE-2019-11111 (2019-11-14T20:15:00)
> Pointer corruption in the Unified Shader Compiler in Intel(R) Graphics Drivers before 10.18.14.5074 (aka 15.36.x.5074) may allow an authenticated user to potentially enable escalation of privilege via local access.
- [Live-Hack-CVE/CVE-2019-11111](https://github.com/Live-Hack-CVE/CVE-2019-11111)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11111">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11111">

---
## CVE-2019-11089 (2019-11-14T20:15:00)
> Insufficient input validation in Kernel Mode module for Intel(R) Graphics Driver before version 25.20.100.6519 may allow an authenticated user to potentially enable denial of service via local access.
- [Live-Hack-CVE/CVE-2019-11089](https://github.com/Live-Hack-CVE/CVE-2019-11089)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11089">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11089">

---
## CVE-2019-11050 (2019-12-23T03:15:00)
> When PHP EXIF extension is parsing EXIF information from an image, e.g. via exif_read_data() function, in PHP versions 7.2.x below 7.2.26, 7.3.x below 7.3.13 and 7.4.0 it is possible to supply it with data what will cause it to read past the allocated buffer. This may lead to information disclosure or crash.
- [Live-Hack-CVE/CVE-2019-11050](https://github.com/Live-Hack-CVE/CVE-2019-11050)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11050">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11050">

---
## CVE-2019-11049 (2019-12-23T03:15:00)
> In PHP versions 7.3.x below 7.3.13 and 7.4.0 on Windows, when supplying custom headers to mail() function, due to mistake introduced in commit 78f4b4a2dcf92ddbccea1bb95f8390a18ac3342e, if the header is supplied in lowercase, this can result in double-freeing certain memory locations.
- [Live-Hack-CVE/CVE-2019-11049](https://github.com/Live-Hack-CVE/CVE-2019-11049)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11049">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11049">

---
## CVE-2019-11047 (2019-12-23T03:15:00)
> When PHP EXIF extension is parsing EXIF information from an image, e.g. via exif_read_data() function, in PHP versions 7.2.x below 7.2.26, 7.3.x below 7.3.13 and 7.4.0 it is possible to supply it with data what will cause it to read past the allocated buffer. This may lead to information disclosure or crash.
- [Live-Hack-CVE/CVE-2019-11047](https://github.com/Live-Hack-CVE/CVE-2019-11047)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11047">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11047">

---
## CVE-2019-11046 (2019-12-23T03:15:00)
> In PHP versions 7.2.x below 7.2.26, 7.3.x below 7.3.13 and 7.4.0, PHP bcmath extension functions on some systems, including Windows, can be tricked into reading beyond the allocated space by supplying it with string containing characters that are identified as numeric by the OS but aren't ASCII numbers. This can read to disclosure of the content of some memory locations.
- [Live-Hack-CVE/CVE-2019-11046](https://github.com/Live-Hack-CVE/CVE-2019-11046)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11046">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11046">

---
## CVE-2019-11045 (2019-12-23T03:15:00)
> In PHP versions 7.2.x below 7.2.26, 7.3.x below 7.3.13 and 7.4.0, PHP DirectoryIterator class accepts filenames with embedded \0 byte and treats them as terminating at that byte. This could lead to security vulnerabilities, e.g. in applications checking paths that the code is allowed to access.
- [Live-Hack-CVE/CVE-2019-11045](https://github.com/Live-Hack-CVE/CVE-2019-11045)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11045">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11045">
- [Live-Hack-CVE/CVE-2019-11045](https://github.com/Live-Hack-CVE/CVE-2019-11045)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-11045">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-11045">

---
## CVE-2019-11043 (2019-10-28T15:15:00)
> In PHP versions 7.1.x below 7.1.33, 7.2.x below 7.2.24 and 7.3.x below 7.3.11 in certain configurations of FPM setup it is possible to cause FPM module to write past allocated buffers into the space reserved for FCGI protocol data, thus opening the possibility of remote code execution.
- [trhacknon/CVE-2019-11043](https://github.com/trhacknon/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/trhacknon/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/trhacknon/CVE-2019-11043">
- [HxDDD/CVE-PoC](https://github.com/HxDDD/CVE-PoC)	<img alt="forks" src="https://img.shields.io/github/forks/HxDDD/CVE-PoC">	<img alt="stars" src="https://img.shields.io/github/stars/HxDDD/CVE-PoC">
- [jas9reet/CVE-2019-11043](https://github.com/jas9reet/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/jas9reet/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/jas9reet/CVE-2019-11043">
- [jptr218/php_hack](https://github.com/jptr218/php_hack)	<img alt="forks" src="https://img.shields.io/github/forks/jptr218/php_hack">	<img alt="stars" src="https://img.shields.io/github/stars/jptr218/php_hack">
- [lindemer/CVE-2019-11043](https://github.com/lindemer/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/lindemer/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/lindemer/CVE-2019-11043">
- [jas502n/CVE-2019-11043](https://github.com/jas502n/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/jas502n/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/jas502n/CVE-2019-11043">
- [corifeo/CVE-2019-11043](https://github.com/corifeo/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/corifeo/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/corifeo/CVE-2019-11043">
- [AleWong/PHP-FPM-Remote-Code-Execution-Vulnerability-CVE-2019-11043-](https://github.com/AleWong/PHP-FPM-Remote-Code-Execution-Vulnerability-CVE-2019-11043-)	<img alt="forks" src="https://img.shields.io/github/forks/AleWong/PHP-FPM-Remote-Code-Execution-Vulnerability-CVE-2019-11043-">	<img alt="stars" src="https://img.shields.io/github/stars/AleWong/PHP-FPM-Remote-Code-Execution-Vulnerability-CVE-2019-11043-">
- [kriskhub/CVE-2019-11043](https://github.com/kriskhub/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/kriskhub/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/kriskhub/CVE-2019-11043">
- [alokaranasinghe/cve-2019-11043](https://github.com/alokaranasinghe/cve-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/alokaranasinghe/cve-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/alokaranasinghe/cve-2019-11043">
- [moniik/CVE-2019-11043_env](https://github.com/moniik/CVE-2019-11043_env)	<img alt="forks" src="https://img.shields.io/github/forks/moniik/CVE-2019-11043_env">	<img alt="stars" src="https://img.shields.io/github/stars/moniik/CVE-2019-11043_env">
- [neex/phuip-fpizdam](https://github.com/neex/phuip-fpizdam)	<img alt="forks" src="https://img.shields.io/github/forks/neex/phuip-fpizdam">	<img alt="stars" src="https://img.shields.io/github/stars/neex/phuip-fpizdam">
- [k8gege/CVE-2019-11043](https://github.com/k8gege/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/k8gege/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/k8gege/CVE-2019-11043">
- [0th3rs-Security-Team/CVE-2019-11043](https://github.com/0th3rs-Security-Team/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/0th3rs-Security-Team/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/0th3rs-Security-Team/CVE-2019-11043">
- [MRdoulestar/CVE-2019-11043](https://github.com/MRdoulestar/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/MRdoulestar/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/MRdoulestar/CVE-2019-11043">
- [ypereirareis/docker-CVE-2019-11043](https://github.com/ypereirareis/docker-CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/ypereirareis/docker-CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/ypereirareis/docker-CVE-2019-11043">
- [huowen/CVE-2019-11043](https://github.com/huowen/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/huowen/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/huowen/CVE-2019-11043">
- [theMiddleBlue/CVE-2019-11043](https://github.com/theMiddleBlue/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/theMiddleBlue/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/theMiddleBlue/CVE-2019-11043">
- [shadow-horse/cve-2019-11043](https://github.com/shadow-horse/cve-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/shadow-horse/cve-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/shadow-horse/cve-2019-11043">
- [akamajoris/CVE-2019-11043-Docker](https://github.com/akamajoris/CVE-2019-11043-Docker)	<img alt="forks" src="https://img.shields.io/github/forks/akamajoris/CVE-2019-11043-Docker">	<img alt="stars" src="https://img.shields.io/github/stars/akamajoris/CVE-2019-11043-Docker">
- [fairyming/CVE-2019-11043](https://github.com/fairyming/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/fairyming/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/fairyming/CVE-2019-11043">
- [ianxtianxt/CVE-2019-11043](https://github.com/ianxtianxt/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/ianxtianxt/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/ianxtianxt/CVE-2019-11043">
- [tinker-li/CVE-2019-11043](https://github.com/tinker-li/CVE-2019-11043)	<img alt="forks" src="https://img.shields.io/github/forks/tinker-li/CVE-2019-11043">	<img alt="stars" src="https://img.shields.io/github/stars/tinker-li/CVE-2019-11043">

---
## CVE-2019-10808 (2020-03-11T23:15:00)
> utilitify prior to 1.0.3 allows modification of object properties. The merge method could be tricked into adding or modifying properties of the Object.prototype.
- [Live-Hack-CVE/CVE-2019-10808](https://github.com/Live-Hack-CVE/CVE-2019-10808)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-10808">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-10808">

---
## CVE-2019-10742 (2019-05-07T19:29:00)
> Axios up to and including 0.18.0 allows attackers to cause a denial of service (application crash) by continuing to accepting content after maxContentLength is exceeded.
- [Viniciuspxf/CVE-2019-10742](https://github.com/Viniciuspxf/CVE-2019-10742)	<img alt="forks" src="https://img.shields.io/github/forks/Viniciuspxf/CVE-2019-10742">	<img alt="stars" src="https://img.shields.io/github/stars/Viniciuspxf/CVE-2019-10742">
- [ossf-cve-benchmark/CVE-2019-10742](https://github.com/ossf-cve-benchmark/CVE-2019-10742)	<img alt="forks" src="https://img.shields.io/github/forks/ossf-cve-benchmark/CVE-2019-10742">	<img alt="stars" src="https://img.shields.io/github/stars/ossf-cve-benchmark/CVE-2019-10742">

---
## CVE-2019-10433 (2019-10-01T14:15:00)
> Jenkins Dingding[??] Plugin stores credentials unencrypted in job config.xml files on the Jenkins master where they can be viewed by users with Extended Read permission, or access to the master file system.
- [Live-Hack-CVE/CVE-2019-10433](https://github.com/Live-Hack-CVE/CVE-2019-10433)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-10433">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-10433">

---
## CVE-2019-10220 (2019-11-27T16:15:00)
> Linux kernel CIFS implementation, version 4.9.0 is vulnerable to a relative paths injection in directory entry lists.
- [nidhi7598/linux-4.1.15_CVE-2019-10220](https://github.com/nidhi7598/linux-4.1.15_CVE-2019-10220)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/linux-4.1.15_CVE-2019-10220">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/linux-4.1.15_CVE-2019-10220">
- [nidhi7598/linux-4.1.15_CVE-2019-10220](https://github.com/nidhi7598/linux-4.1.15_CVE-2019-10220)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/linux-4.1.15_CVE-2019-10220">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/linux-4.1.15_CVE-2019-10220">
- [Trinadh465/device_renesas_kernel_AOSP10_r33_CVE-2019-10220](https://github.com/Trinadh465/device_renesas_kernel_AOSP10_r33_CVE-2019-10220)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/device_renesas_kernel_AOSP10_r33_CVE-2019-10220">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/device_renesas_kernel_AOSP10_r33_CVE-2019-10220">
- [Trinadh465/linux-4.19.72_CVE-2019-10220](https://github.com/Trinadh465/linux-4.19.72_CVE-2019-10220)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/linux-4.19.72_CVE-2019-10220">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/linux-4.19.72_CVE-2019-10220">
- [Trinadh465/linux_3.0.35_CVE-2019-10220](https://github.com/Trinadh465/linux_3.0.35_CVE-2019-10220)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/linux_3.0.35_CVE-2019-10220">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/linux_3.0.35_CVE-2019-10220">
- [Trinadh465/linux-3.0.35_CVE-2019-10220](https://github.com/Trinadh465/linux-3.0.35_CVE-2019-10220)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/linux-3.0.35_CVE-2019-10220">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/linux-3.0.35_CVE-2019-10220">

---
## CVE-2019-1010319 (2019-07-11T20:15:00)
> WavPack 5.1.0 and earlier is affected by: CWE-457: Use of Uninitialized Variable. The impact is: Unexpected control flow, crashes, and segfaults. The component is: ParseWave64HeaderConfig (wave64.c:211). The attack vector is: Maliciously crafted .wav file. The fixed version is: After commit https://github.com/dbry/WavPack/commit/33a0025d1d63ccd05d9dbaa6923d52b1446a62fe.
- [ahaShiyu/CVE-2019-1010319](https://github.com/ahaShiyu/CVE-2019-1010319)	<img alt="forks" src="https://img.shields.io/github/forks/ahaShiyu/CVE-2019-1010319">	<img alt="stars" src="https://img.shields.io/github/stars/ahaShiyu/CVE-2019-1010319">

---
## CVE-2019-1010065 (2019-07-18T17:15:00)
> The Sleuth Kit 4.6.0 and earlier is affected by: Integer Overflow. The impact is: Opening crafted disk image triggers crash in tsk/fs/hfs_dent.c:237. The component is: Overflow in fls tool used on HFS image. Bug is in tsk/fs/hfs.c file in function hfs_cat_traverse() in lines: 952, 1062. The attack vector is: Victim must open a crafted HFS filesystem image.
- [Live-Hack-CVE/CVE-2019-1010065](https://github.com/Live-Hack-CVE/CVE-2019-1010065)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-1010065">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-1010065">

---
## CVE-2019-0845 (2019-04-09T21:29:00)
> A remote code execution vulnerability exists when the IOleCvt interface renders ASP webpage content, aka 'Windows IOleCvt Interface Remote Code Execution Vulnerability'.
- [Live-Hack-CVE/CVE-2019-0845](https://github.com/Live-Hack-CVE/CVE-2019-0845)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-0845">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-0845">

---
## CVE-2019-0230 (2020-09-14T17:15:00)
> Apache Struts 2.0.0 to 2.5.20 forced double OGNL evaluation, when evaluated on raw user input in tag attributes, may lead to remote code execution.
- [Live-Hack-CVE/CVE-2019-0230](https://github.com/Live-Hack-CVE/CVE-2019-0230)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2019-0230">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2019-0230">
