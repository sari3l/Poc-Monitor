# 2015 List

---
## CVE-2015-9235 (2018-05-29T20:29:00)
> In jsonwebtoken node module before 4.2.2 it is possible for an attacker to bypass verification when a token digitally signed with an asymmetric key (RS/ES family) of algorithms but instead the attacker send a token digitally signed with a symmetric algorithm (HS* family).
- [WinDyAlphA/CVE-2015-9235_JWT_key_confusion](https://github.com/WinDyAlphA/CVE-2015-9235_JWT_key_confusion)	<img alt="forks" src="https://img.shields.io/github/forks/WinDyAlphA/CVE-2015-9235_JWT_key_confusion">	<img alt="stars" src="https://img.shields.io/github/stars/WinDyAlphA/CVE-2015-9235_JWT_key_confusion">
- [aalex954/jwt-key-confusion-poc](https://github.com/aalex954/jwt-key-confusion-poc)	<img alt="forks" src="https://img.shields.io/github/forks/aalex954/jwt-key-confusion-poc">	<img alt="stars" src="https://img.shields.io/github/stars/aalex954/jwt-key-confusion-poc">

---
## CVE-2015-8994 (2017-03-02T06:59:00)
> An issue was discovered in PHP 5.x and 7.x, when the configuration uses apache2handler/mod_php or php-fpm with OpCache enabled. With 5.x after 5.6.28 or 7.x after 7.0.13, the issue is resolved in a non-default configuration with the opcache.validate_permission=1 setting. The vulnerability details are as follows. In PHP SAPIs where PHP interpreters share a common parent process, Zend OpCache creates a shared memory object owned by the common parent during initialization. Child PHP processes inherit the SHM descriptor, using it to cache and retrieve compiled script bytecode ("opcode" in PHP jargon). Cache keys vary depending on configuration, but filename is a central key component, and compiled opcode can generally be run if a script's filename is known or can be guessed. Many common shared-hosting configurations change EUID in child processes to enforce privilege separation among hosted users (for example using mod_ruid2 for the Apache HTTP Server, or php-fpm user settings). In these scenarios, the default Zend OpCache behavior defeats script file permissions by sharing a single SHM cache among all child PHP processes. PHP scripts often contain sensitive information: Think of CMS configurations where reading or running another user's script usually means gaining privileges to the CMS database.
- [Live-Hack-CVE/CVE-2015-8994](https://github.com/Live-Hack-CVE/CVE-2015-8994)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8994">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8994">

---
## CVE-2015-8873 (2016-05-16T10:59:00)
> Stack consumption vulnerability in Zend/zend_exceptions.c in PHP before 5.4.44, 5.5.x before 5.5.28, and 5.6.x before 5.6.12 allows remote attackers to cause a denial of service (segmentation fault) via recursive method calls.
- [Live-Hack-CVE/CVE-2015-8873](https://github.com/Live-Hack-CVE/CVE-2015-8873)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8873">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8873">

---
## CVE-2015-8660 (2015-12-28T11:59:00)
> The ovl_setattr function in fs/overlayfs/inode.c in the Linux kernel through 4.3.3 attempts to merge distinct setattr operations, which allows local users to bypass intended access restrictions and modify the attributes of arbitrary overlay files via a crafted application.
- [nhamle2/CVE-2015-8660](https://github.com/nhamle2/CVE-2015-8660)	<img alt="forks" src="https://img.shields.io/github/forks/nhamle2/CVE-2015-8660">	<img alt="stars" src="https://img.shields.io/github/stars/nhamle2/CVE-2015-8660">
- [whu-enjoy/CVE-2015-8660](https://github.com/whu-enjoy/CVE-2015-8660)	<img alt="forks" src="https://img.shields.io/github/forks/whu-enjoy/CVE-2015-8660">	<img alt="stars" src="https://img.shields.io/github/stars/whu-enjoy/CVE-2015-8660">

---
## CVE-2015-8562 (2015-12-16T21:59:00)
> Joomla! 1.5.x, 2.x, and 3.x before 3.4.6 allow remote attackers to conduct PHP object injection attacks and execute arbitrary PHP code via the HTTP User-Agent header, as exploited in the wild in December 2015.
- [Caihuar/Joomla-cve-2015-8562](https://github.com/Caihuar/Joomla-cve-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/Caihuar/Joomla-cve-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/Caihuar/Joomla-cve-2015-8562">
- [lorenzodegiorgi/setup-cve-2015-8562](https://github.com/lorenzodegiorgi/setup-cve-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/lorenzodegiorgi/setup-cve-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/lorenzodegiorgi/setup-cve-2015-8562">
- [guanjivip/CVE-2015-8562](https://github.com/guanjivip/CVE-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/guanjivip/CVE-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/guanjivip/CVE-2015-8562">
- [VoidSec/Joomla_CVE-2015-8562](https://github.com/VoidSec/Joomla_CVE-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/VoidSec/Joomla_CVE-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/VoidSec/Joomla_CVE-2015-8562">
- [xnorkl/Joomla_Payload](https://github.com/xnorkl/Joomla_Payload)	<img alt="forks" src="https://img.shields.io/github/forks/xnorkl/Joomla_Payload">	<img alt="stars" src="https://img.shields.io/github/stars/xnorkl/Joomla_Payload">
- [paralelo14/JoomlaMassExploiter](https://github.com/paralelo14/JoomlaMassExploiter)	<img alt="forks" src="https://img.shields.io/github/forks/paralelo14/JoomlaMassExploiter">	<img alt="stars" src="https://img.shields.io/github/stars/paralelo14/JoomlaMassExploiter">
- [paralelo14/CVE-2015-8562](https://github.com/paralelo14/CVE-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/paralelo14/CVE-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/paralelo14/CVE-2015-8562">
- [thejackerz/scanner-exploit-joomla-CVE-2015-8562](https://github.com/thejackerz/scanner-exploit-joomla-CVE-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/thejackerz/scanner-exploit-joomla-CVE-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/thejackerz/scanner-exploit-joomla-CVE-2015-8562">
- [atcasanova/cve-2015-8562-exploit](https://github.com/atcasanova/cve-2015-8562-exploit)	<img alt="forks" src="https://img.shields.io/github/forks/atcasanova/cve-2015-8562-exploit">	<img alt="stars" src="https://img.shields.io/github/stars/atcasanova/cve-2015-8562-exploit">
- [RobinHoutevelts/Joomla-CVE-2015-8562-PHP-POC](https://github.com/RobinHoutevelts/Joomla-CVE-2015-8562-PHP-POC)	<img alt="forks" src="https://img.shields.io/github/forks/RobinHoutevelts/Joomla-CVE-2015-8562-PHP-POC">	<img alt="stars" src="https://img.shields.io/github/stars/RobinHoutevelts/Joomla-CVE-2015-8562-PHP-POC">
- [ZaleHack/joomla_rce_CVE-2015-8562](https://github.com/ZaleHack/joomla_rce_CVE-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/ZaleHack/joomla_rce_CVE-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/ZaleHack/joomla_rce_CVE-2015-8562">
- [Caihuar/Joomla-cve-2015-8562](https://github.com/Caihuar/Joomla-cve-2015-8562)	<img alt="forks" src="https://img.shields.io/github/forks/Caihuar/Joomla-cve-2015-8562">	<img alt="stars" src="https://img.shields.io/github/stars/Caihuar/Joomla-cve-2015-8562">

---
## CVE-2015-8467 (2015-12-29T22:59:00)
> The samldb_check_user_account_control_acl function in dsdb/samdb/ldb_modules/samldb.c in Samba 4.x before 4.1.22, 4.2.x before 4.2.7, and 4.3.x before 4.3.3 does not properly check for administrative privileges during creation of machine accounts, which allows remote authenticated users to bypass intended access restrictions by leveraging the existence of a domain with both a Samba DC and a Windows DC, a similar issue to CVE-2015-2535.
- [Live-Hack-CVE/CVE-2015-8467](https://github.com/Live-Hack-CVE/CVE-2015-8467)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8467">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8467">

---
## CVE-2015-8394 (2015-12-02T01:59:00)
> PCRE before 8.38 mishandles the (?(<digits>) and (?(R<digits>) conditions, which allows remote attackers to cause a denial of service (integer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.
- [Live-Hack-CVE/CVE-2015-8394](https://github.com/Live-Hack-CVE/CVE-2015-8394)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8394">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8394">

---
## CVE-2015-8393 (2015-12-02T01:59:00)
> pcregrep in PCRE before 8.38 mishandles the -q option for binary files, which might allow remote attackers to obtain sensitive information via a crafted file, as demonstrated by a CGI script that sends stdout data to a client.
- [Live-Hack-CVE/CVE-2015-8393](https://github.com/Live-Hack-CVE/CVE-2015-8393)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8393">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8393">

---
## CVE-2015-8390 (2015-12-02T01:59:00)
> PCRE before 8.38 mishandles the [: and \\ substrings in character classes, which allows remote attackers to cause a denial of service (uninitialized memory read) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.
- [Live-Hack-CVE/CVE-2015-8390](https://github.com/Live-Hack-CVE/CVE-2015-8390)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8390">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8390">

---
## CVE-2015-8389 (2015-12-02T01:59:00)
> PCRE before 8.38 mishandles the /(?:|a|){100}x/ pattern and related patterns, which allows remote attackers to cause a denial of service (infinite recursion) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.
- [Live-Hack-CVE/CVE-2015-8389](https://github.com/Live-Hack-CVE/CVE-2015-8389)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8389">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8389">

---
## CVE-2015-8387 (2015-12-02T01:59:00)
> PCRE before 8.38 mishandles (?123) subroutine calls and related subroutine calls, which allows remote attackers to cause a denial of service (integer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.
- [Live-Hack-CVE/CVE-2015-8387](https://github.com/Live-Hack-CVE/CVE-2015-8387)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8387">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8387">

---
## CVE-2015-8386 (2015-12-02T01:59:00)
> PCRE before 8.38 mishandles the interaction of lookbehind assertions and mutually recursive subpatterns, which allows remote attackers to cause a denial of service (buffer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.
- [Live-Hack-CVE/CVE-2015-8386](https://github.com/Live-Hack-CVE/CVE-2015-8386)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8386">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8386">

---
## CVE-2015-8383 (2015-12-02T01:59:00)
> PCRE before 8.38 mishandles certain repeated conditional groups, which allows remote attackers to cause a denial of service (buffer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.
- [Live-Hack-CVE/CVE-2015-8383](https://github.com/Live-Hack-CVE/CVE-2015-8383)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8383">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8383">

---
## CVE-2015-8351 (2017-09-11T20:29:00)
> PHP remote file inclusion vulnerability in the Gwolle Guestbook plugin before 1.5.4 for WordPress, when allow_url_include is enabled, allows remote authenticated users to execute arbitrary PHP code via a URL in the abspath parameter to frontend/captcha/ajaxresponse.php.  NOTE: this can also be leveraged to include and execute arbitrary local files via directory traversal sequences regardless of whether allow_url_include is enabled.
- [igruntplay/exploit-CVE-2015-8351](https://github.com/igruntplay/exploit-CVE-2015-8351)	<img alt="forks" src="https://img.shields.io/github/forks/igruntplay/exploit-CVE-2015-8351">	<img alt="stars" src="https://img.shields.io/github/stars/igruntplay/exploit-CVE-2015-8351">
- [Ki11i0n4ir3/CVE-2015-8351](https://github.com/Ki11i0n4ir3/CVE-2015-8351)	<img alt="forks" src="https://img.shields.io/github/forks/Ki11i0n4ir3/CVE-2015-8351">	<img alt="stars" src="https://img.shields.io/github/stars/Ki11i0n4ir3/CVE-2015-8351">

---
## CVE-2015-8103 (2015-11-25T20:59:00)
> The Jenkins CLI subsystem in Jenkins before 1.638 and LTS before 1.625.2 allows remote attackers to execute arbitrary code via a crafted serialized Java object, related to a problematic webapps/ROOT/WEB-INF/lib/commons-collections-*.jar file and the "Groovy variant in 'ysoserial'".
- [r00t4dm/Jenkins-CVE-2015-8103](https://github.com/r00t4dm/Jenkins-CVE-2015-8103)	<img alt="forks" src="https://img.shields.io/github/forks/r00t4dm/Jenkins-CVE-2015-8103">	<img alt="stars" src="https://img.shields.io/github/stars/r00t4dm/Jenkins-CVE-2015-8103">
- [cved-sources/cve-2015-8103](https://github.com/cved-sources/cve-2015-8103)	<img alt="forks" src="https://img.shields.io/github/forks/cved-sources/cve-2015-8103">	<img alt="stars" src="https://img.shields.io/github/stars/cved-sources/cve-2015-8103">

---
## CVE-2015-7547 (2016-02-18T21:59:00)
> Multiple stack-based buffer overflows in the (1) send_dg and (2) send_vc functions in the libresolv library in the GNU C Library (aka glibc or libc6) before 2.23 allow remote attackers to cause a denial of service (crash) or possibly execute arbitrary code via a crafted DNS response that triggers a call to the getaddrinfo function with the AF_UNSPEC or AF_INET6 address family, related to performing "dual A/AAAA DNS queries" and the libnss_dns.so.2 NSS module.
- [Stick-U235/CVE-2015-7547-Research](https://github.com/Stick-U235/CVE-2015-7547-Research)	<img alt="forks" src="https://img.shields.io/github/forks/Stick-U235/CVE-2015-7547-Research">	<img alt="stars" src="https://img.shields.io/github/stars/Stick-U235/CVE-2015-7547-Research">
- [Amilaperera12/Glibc-Vulnerability-Exploit-CVE-2015-7547](https://github.com/Amilaperera12/Glibc-Vulnerability-Exploit-CVE-2015-7547)	<img alt="forks" src="https://img.shields.io/github/forks/Amilaperera12/Glibc-Vulnerability-Exploit-CVE-2015-7547">	<img alt="stars" src="https://img.shields.io/github/stars/Amilaperera12/Glibc-Vulnerability-Exploit-CVE-2015-7547">
- [miracle03/CVE-2015-7547-master](https://github.com/miracle03/CVE-2015-7547-master)	<img alt="forks" src="https://img.shields.io/github/forks/miracle03/CVE-2015-7547-master">	<img alt="stars" src="https://img.shields.io/github/stars/miracle03/CVE-2015-7547-master">
- [bluebluelan/CVE-2015-7547-proj-master](https://github.com/bluebluelan/CVE-2015-7547-proj-master)	<img alt="forks" src="https://img.shields.io/github/forks/bluebluelan/CVE-2015-7547-proj-master">	<img alt="stars" src="https://img.shields.io/github/stars/bluebluelan/CVE-2015-7547-proj-master">
- [eSentire/cve-2015-7547-public](https://github.com/eSentire/cve-2015-7547-public)	<img alt="forks" src="https://img.shields.io/github/forks/eSentire/cve-2015-7547-public">	<img alt="stars" src="https://img.shields.io/github/stars/eSentire/cve-2015-7547-public">
- [jgajek/cve-2015-7547](https://github.com/jgajek/cve-2015-7547)	<img alt="forks" src="https://img.shields.io/github/forks/jgajek/cve-2015-7547">	<img alt="stars" src="https://img.shields.io/github/stars/jgajek/cve-2015-7547">
- [t0r0t0r0/CVE-2015-7547](https://github.com/t0r0t0r0/CVE-2015-7547)	<img alt="forks" src="https://img.shields.io/github/forks/t0r0t0r0/CVE-2015-7547">	<img alt="stars" src="https://img.shields.io/github/stars/t0r0t0r0/CVE-2015-7547">
- [babykillerblack/CVE-2015-7547](https://github.com/babykillerblack/CVE-2015-7547)	<img alt="forks" src="https://img.shields.io/github/forks/babykillerblack/CVE-2015-7547">	<img alt="stars" src="https://img.shields.io/github/stars/babykillerblack/CVE-2015-7547">
- [fjserna/CVE-2015-7547](https://github.com/fjserna/CVE-2015-7547)	<img alt="forks" src="https://img.shields.io/github/forks/fjserna/CVE-2015-7547">	<img alt="stars" src="https://img.shields.io/github/stars/fjserna/CVE-2015-7547">
- [rexifiles/rex-sec-glibc](https://github.com/rexifiles/rex-sec-glibc)	<img alt="forks" src="https://img.shields.io/github/forks/rexifiles/rex-sec-glibc">	<img alt="stars" src="https://img.shields.io/github/stars/rexifiles/rex-sec-glibc">
- [cakuzo/CVE-2015-7547](https://github.com/cakuzo/CVE-2015-7547)	<img alt="forks" src="https://img.shields.io/github/forks/cakuzo/CVE-2015-7547">	<img alt="stars" src="https://img.shields.io/github/stars/cakuzo/CVE-2015-7547">
- [MRC6666/CVE-2015-7547](https://github.com/MRC6666/CVE-2015-7547)	<img alt="forks" src="https://img.shields.io/github/forks/MRC6666/CVE-2015-7547">	<img alt="stars" src="https://img.shields.io/github/stars/MRC6666/CVE-2015-7547">

---
## CVE-2015-7297 (2015-10-29T20:59:00)
> SQL injection vulnerability in Joomla! 3.2 before 3.4.4 allows remote attackers to execute arbitrary SQL commands via unspecified vectors, a different vulnerability than CVE-2015-7858.
- [Cappricio-Securities/CVE-2015-7297](https://github.com/Cappricio-Securities/CVE-2015-7297)	<img alt="forks" src="https://img.shields.io/github/forks/Cappricio-Securities/CVE-2015-7297">	<img alt="stars" src="https://img.shields.io/github/stars/Cappricio-Securities/CVE-2015-7297">
- [areaventuno/exploit-joomla](https://github.com/areaventuno/exploit-joomla)	<img alt="forks" src="https://img.shields.io/github/forks/areaventuno/exploit-joomla">	<img alt="stars" src="https://img.shields.io/github/stars/areaventuno/exploit-joomla">
- [CCrashBandicot/ContentHistory](https://github.com/CCrashBandicot/ContentHistory)	<img alt="forks" src="https://img.shields.io/github/forks/CCrashBandicot/ContentHistory">	<img alt="stars" src="https://img.shields.io/github/stars/CCrashBandicot/ContentHistory">

---
## CVE-2015-6967 (2015-09-16T14:59:00)
> Unrestricted file upload vulnerability in the My Image plugin in Nibbleblog before 4.0.5 allows remote administrators to execute arbitrary code by uploading a file with an executable extension, then accessing it via a direct request to the file in content/private/plugins/my_image/image.php. <a href="http://cwe.mitre.org/data/definitions/434.html">CWE-434: Unrestricted Upload of File with Dangerous Type</a>
- [FredBrave/CVE-2015-6967](https://github.com/FredBrave/CVE-2015-6967)	<img alt="forks" src="https://img.shields.io/github/forks/FredBrave/CVE-2015-6967">	<img alt="stars" src="https://img.shields.io/github/stars/FredBrave/CVE-2015-6967">
- [hadrian3689/nibbleblog_4.0.3](https://github.com/hadrian3689/nibbleblog_4.0.3)	<img alt="forks" src="https://img.shields.io/github/forks/hadrian3689/nibbleblog_4.0.3">	<img alt="stars" src="https://img.shields.io/github/stars/hadrian3689/nibbleblog_4.0.3">
- [0xkasra/CVE-2015-6967](https://github.com/0xkasra/CVE-2015-6967)	<img alt="forks" src="https://img.shields.io/github/forks/0xkasra/CVE-2015-6967">	<img alt="stars" src="https://img.shields.io/github/stars/0xkasra/CVE-2015-6967">
- [dix0nym/CVE-2015-6967](https://github.com/dix0nym/CVE-2015-6967)	<img alt="forks" src="https://img.shields.io/github/forks/dix0nym/CVE-2015-6967">	<img alt="stars" src="https://img.shields.io/github/stars/dix0nym/CVE-2015-6967">

---
## CVE-2015-6764 (2015-12-06T01:59:00)
> The BasicJsonStringifier::SerializeJSArray function in json-stringifier.h in the JSON stringifier in Google V8, as used in Google Chrome before 47.0.2526.73, improperly loads array elements, which allows remote attackers to cause a denial of service (out-of-bounds memory access) or possibly have unspecified other impact via crafted JavaScript code.
- [Live-Hack-CVE/CVE-2015-6764](https://github.com/Live-Hack-CVE/CVE-2015-6764)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-6764">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-6764">

---
## CVE-2015-6668 (2017-10-19T21:29:00)
> The Job Manager plugin before 0.7.25 allows remote attackers to read arbitrary CV files via a brute force attack to the WordPress upload directory structure, related to an insecure direct object reference.
- [jimdiroffii/CVE-2015-6668](https://github.com/jimdiroffii/CVE-2015-6668)	<img alt="forks" src="https://img.shields.io/github/forks/jimdiroffii/CVE-2015-6668">	<img alt="stars" src="https://img.shields.io/github/stars/jimdiroffii/CVE-2015-6668">
- [c0d3cr4f73r/CVE-2015-6668](https://github.com/c0d3cr4f73r/CVE-2015-6668)	<img alt="forks" src="https://img.shields.io/github/forks/c0d3cr4f73r/CVE-2015-6668">	<img alt="stars" src="https://img.shields.io/github/stars/c0d3cr4f73r/CVE-2015-6668">
- [G01d3nW01f/CVE-2015-6668](https://github.com/G01d3nW01f/CVE-2015-6668)	<img alt="forks" src="https://img.shields.io/github/forks/G01d3nW01f/CVE-2015-6668">	<img alt="stars" src="https://img.shields.io/github/stars/G01d3nW01f/CVE-2015-6668">

---
## CVE-2015-5531 (2015-08-17T15:59:00)
> Directory traversal vulnerability in Elasticsearch before 1.6.1 allows remote attackers to read arbitrary files via unspecified vectors related to snapshot API calls.
- [xpgdgit/CVE-2015-5531](https://github.com/xpgdgit/CVE-2015-5531)	<img alt="forks" src="https://img.shields.io/github/forks/xpgdgit/CVE-2015-5531">	<img alt="stars" src="https://img.shields.io/github/stars/xpgdgit/CVE-2015-5531">
- [j-jasson/CVE-2015-5531-POC](https://github.com/j-jasson/CVE-2015-5531-POC)	<img alt="forks" src="https://img.shields.io/github/forks/j-jasson/CVE-2015-5531-POC">	<img alt="stars" src="https://img.shields.io/github/stars/j-jasson/CVE-2015-5531-POC">

---
## CVE-2015-5521 (2015-07-14T16:59:00)
> Cross-site scripting (XSS) vulnerability in BlackCat CMS 1.1.2 allows remote attackers to inject arbitrary web script or HTML via the name in a new group to backend/groups/index.php.
- [Live-Hack-CVE/CVE-2015-5521](https://github.com/Live-Hack-CVE/CVE-2015-5521)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-5521">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-5521">

---
## CVE-2015-5361 (2020-02-28T23:15:00)
> Background For regular, unencrypted FTP traffic, the FTP ALG can inspect the unencrypted control channel and open related sessions for the FTP data channel. These related sessions (gates) are specific to source and destination IPs and ports of client and server. The design intent of the ftps-extensions option (which is disabled by default) is to provide similar functionality when the SRX secures the FTP/FTPS client. As the control channel is encrypted, the FTP ALG cannot inspect the port specific information and will open a wider TCP data channel (gate) from client IP to server IP on all destination TCP ports. In FTP/FTPS client environments to an enterprise network or the Internet, this is the desired behavior as it allows firewall policy to be written to FTP/FTPS servers on well-known control ports without using a policy with destination IP ANY and destination port ANY. Issue The ftps-extensions option is not intended or recommended where the SRX secures the FTPS server, as the wide data channel session (gate) will allow the FTPS client temporary access to all TCP ports on the FTPS server. The data session is associated to the control channel and will be closed when the control channel session closes. Depending on the configuration of the FTPS server, supporting load-balancer, and SRX inactivity-timeout values, the server/load-balancer and SRX may keep the control channel open for an extended period of time, allowing an FTPS client access for an equal duration.? Note that the ftps-extensions option is not enabled by default.
- [Live-Hack-CVE/CVE-2015-5361](https://github.com/Live-Hack-CVE/CVE-2015-5361)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-5361">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-5361">
- [Live-Hack-CVE/CVE-2015-5361](https://github.com/Live-Hack-CVE/CVE-2015-5361)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-5361">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-5361">

---
## CVE-2015-5299 (2015-12-29T22:59:00)
> The shadow_copy2_get_shadow_copy_data function in modules/vfs_shadow_copy2.c in Samba 3.x and 4.x before 4.1.22, 4.2.x before 4.2.7, and 4.3.x before 4.3.3 does not verify that the DIRECTORY_LIST access right has been granted, which allows remote attackers to access snapshots by visiting a shadow copy directory.
- [Live-Hack-CVE/CVE-2015-5299](https://github.com/Live-Hack-CVE/CVE-2015-5299)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-5299">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-5299">

---
## CVE-2015-5296 (2015-12-29T22:59:00)
> Samba 3.x and 4.x before 4.1.22, 4.2.x before 4.2.7, and 4.3.x before 4.3.3 supports connections that are encrypted but unsigned, which allows man-in-the-middle attackers to conduct encrypted-to-unencrypted downgrade attacks by modifying the client-server data stream, related to clidfs.c, libsmb_server.c, and smbXcli_base.c.
- [Live-Hack-CVE/CVE-2015-5296](https://github.com/Live-Hack-CVE/CVE-2015-5296)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-5296">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-5296">

---
## CVE-2015-5252 (2015-12-29T22:59:00)
> vfs.c in smbd in Samba 3.x and 4.x before 4.1.22, 4.2.x before 4.2.7, and 4.3.x before 4.3.3, when share names with certain substring relationships exist, allows remote attackers to bypass intended file-access restrictions via a symlink that points outside of a share.
- [Live-Hack-CVE/CVE-2015-5252](https://github.com/Live-Hack-CVE/CVE-2015-5252)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-5252">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-5252">

---
## CVE-2015-4836 (2015-10-21T23:59:00)
> Unspecified vulnerability in Oracle MySQL Server 5.5.45 and earlier, and 5.6.26 and earlier, allows remote authenticated users to affect availability via unknown vectors related to Server : SP.
- [Live-Hack-CVE/CVE-2015-4836](https://github.com/Live-Hack-CVE/CVE-2015-4836)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-4836">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-4836">
- [Live-Hack-CVE/CVE-2015-4836](https://github.com/Live-Hack-CVE/CVE-2015-4836)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-4836">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-4836">

---
## CVE-2015-3884 (2017-03-17T14:59:00)
> Unrestricted file upload vulnerability in the (1) myAccount, (2) projects, (3) tasks, (4) tickets, (5) discussions, (6) reports, and (7) scheduler pages in qdPM 8.3 allows remote attackers to execute arbitrary code by uploading a file with an executable extension, then accessing it via a direct request to the file in uploads/attachments/ or uploads/users/.
- [Live-Hack-CVE/CVE-2015-3884](https://github.com/Live-Hack-CVE/CVE-2015-3884)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3884">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3884">
- [Live-Hack-CVE/CVE-2015-3884](https://github.com/Live-Hack-CVE/CVE-2015-3884)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3884">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3884">

---
## CVE-2015-3864 (2015-10-01T00:59:00)
> Integer underflow in the MPEG4Extractor::parseChunk function in MPEG4Extractor.cpp in libstagefright in mediaserver in Android before 5.1.1 LMY48M allows remote attackers to execute arbitrary code via crafted MPEG-4 data, aka internal bug 23034759.  NOTE: this vulnerability exists because of an incomplete fix for CVE-2015-3824.
- [Bhathiya404/Exploiting-Stagefright-Vulnerability-CVE-2015-3864](https://github.com/Bhathiya404/Exploiting-Stagefright-Vulnerability-CVE-2015-3864)	<img alt="forks" src="https://img.shields.io/github/forks/Bhathiya404/Exploiting-Stagefright-Vulnerability-CVE-2015-3864">	<img alt="stars" src="https://img.shields.io/github/stars/Bhathiya404/Exploiting-Stagefright-Vulnerability-CVE-2015-3864">
- [HenryVHuang/CVE-2015-3864](https://github.com/HenryVHuang/CVE-2015-3864)	<img alt="forks" src="https://img.shields.io/github/forks/HenryVHuang/CVE-2015-3864">	<img alt="stars" src="https://img.shields.io/github/stars/HenryVHuang/CVE-2015-3864">
- [eudemonics/scaredycat](https://github.com/eudemonics/scaredycat)	<img alt="forks" src="https://img.shields.io/github/forks/eudemonics/scaredycat">	<img alt="stars" src="https://img.shields.io/github/stars/eudemonics/scaredycat">
- [pwnaccelerator/stagefright-cve-2015-3864](https://github.com/pwnaccelerator/stagefright-cve-2015-3864)	<img alt="forks" src="https://img.shields.io/github/forks/pwnaccelerator/stagefright-cve-2015-3864">	<img alt="stars" src="https://img.shields.io/github/stars/pwnaccelerator/stagefright-cve-2015-3864">
- [Cmadhushanka/CVE-2015-3864-Exploitation](https://github.com/Cmadhushanka/CVE-2015-3864-Exploitation)	<img alt="forks" src="https://img.shields.io/github/forks/Cmadhushanka/CVE-2015-3864-Exploitation">	<img alt="stars" src="https://img.shields.io/github/stars/Cmadhushanka/CVE-2015-3864-Exploitation">

---
## CVE-2015-3416 (2015-04-24T17:59:00)
> The sqlite3VXPrintf function in printf.c in SQLite before 3.8.9 does not properly handle precision and width values during floating-point conversions, which allows context-dependent attackers to cause a denial of service (integer overflow and stack-based buffer overflow) or possibly have unspecified other impact via large integers in a crafted printf function call in a SELECT statement.
- [Live-Hack-CVE/CVE-2015-3416](https://github.com/Live-Hack-CVE/CVE-2015-3416)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3416">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3416">

---
## CVE-2015-3306 (2015-05-18T15:59:00)
> The mod_copy module in ProFTPD 1.3.5 allows remote attackers to read and write to arbitrary files via the site cpfr and site cpto commands.
- [hackarada/cve-2015-3306](https://github.com/hackarada/cve-2015-3306)	<img alt="forks" src="https://img.shields.io/github/forks/hackarada/cve-2015-3306">	<img alt="stars" src="https://img.shields.io/github/stars/hackarada/cve-2015-3306">
- [jptr218/proftpd_bypass](https://github.com/jptr218/proftpd_bypass)	<img alt="forks" src="https://img.shields.io/github/forks/jptr218/proftpd_bypass">	<img alt="stars" src="https://img.shields.io/github/stars/jptr218/proftpd_bypass">
- [0xm4ud/ProFTPD_CVE-2015-3306](https://github.com/0xm4ud/ProFTPD_CVE-2015-3306)	<img alt="forks" src="https://img.shields.io/github/forks/0xm4ud/ProFTPD_CVE-2015-3306">	<img alt="stars" src="https://img.shields.io/github/stars/0xm4ud/ProFTPD_CVE-2015-3306">
- [cved-sources/cve-2015-3306](https://github.com/cved-sources/cve-2015-3306)	<img alt="forks" src="https://img.shields.io/github/forks/cved-sources/cve-2015-3306">	<img alt="stars" src="https://img.shields.io/github/stars/cved-sources/cve-2015-3306">
- [cd6629/CVE-2015-3306-Python-PoC](https://github.com/cd6629/CVE-2015-3306-Python-PoC)	<img alt="forks" src="https://img.shields.io/github/forks/cd6629/CVE-2015-3306-Python-PoC">	<img alt="stars" src="https://img.shields.io/github/stars/cd6629/CVE-2015-3306-Python-PoC">
- [cdedmondson/Modified-CVE-2015-3306-Exploit](https://github.com/cdedmondson/Modified-CVE-2015-3306-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/cdedmondson/Modified-CVE-2015-3306-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/cdedmondson/Modified-CVE-2015-3306-Exploit">
- [t0kx/exploit-CVE-2015-3306](https://github.com/t0kx/exploit-CVE-2015-3306)	<img alt="forks" src="https://img.shields.io/github/forks/t0kx/exploit-CVE-2015-3306">	<img alt="stars" src="https://img.shields.io/github/stars/t0kx/exploit-CVE-2015-3306">
- [davidtavarez/CVE-2015-3306](https://github.com/davidtavarez/CVE-2015-3306)	<img alt="forks" src="https://img.shields.io/github/forks/davidtavarez/CVE-2015-3306">	<img alt="stars" src="https://img.shields.io/github/stars/davidtavarez/CVE-2015-3306">
- [nootropics/propane](https://github.com/nootropics/propane)	<img alt="forks" src="https://img.shields.io/github/forks/nootropics/propane">	<img alt="stars" src="https://img.shields.io/github/stars/nootropics/propane">
- [shk0x/cpx_proftpd](https://github.com/shk0x/cpx_proftpd)	<img alt="forks" src="https://img.shields.io/github/forks/shk0x/cpx_proftpd">	<img alt="stars" src="https://img.shields.io/github/stars/shk0x/cpx_proftpd">
- [xchg-rax-rax/CVE-2015-3306-](https://github.com/xchg-rax-rax/CVE-2015-3306-)	<img alt="forks" src="https://img.shields.io/github/forks/xchg-rax-rax/CVE-2015-3306-">	<img alt="stars" src="https://img.shields.io/github/stars/xchg-rax-rax/CVE-2015-3306-">

---
## CVE-2015-3197 (2016-02-15T02:59:00)
> ssl/s2_srvr.c in OpenSSL 1.0.1 before 1.0.1r and 1.0.2 before 1.0.2f does not prevent use of disabled ciphers, which makes it easier for man-in-the-middle attackers to defeat cryptographic protection mechanisms by performing computations on SSLv2 traffic, related to the get_client_master_key and get_client_hello functions.
- [Live-Hack-CVE/CVE-2015-3197](https://github.com/Live-Hack-CVE/CVE-2015-3197)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3197">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3197">
- [Trinadh465/OpenSSL-1_0_1g_CVE-2015-3197](https://github.com/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3197)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3197">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3197">

---
## CVE-2015-3196 (2015-12-06T20:59:00)
> ssl/s3_clnt.c in OpenSSL 1.0.0 before 1.0.0t, 1.0.1 before 1.0.1p, and 1.0.2 before 1.0.2d, when used for a multi-threaded client, writes the PSK identity hint to an incorrect data structure, which allows remote servers to cause a denial of service (race condition and double free) via a crafted ServerKeyExchange message.
- [Live-Hack-CVE/CVE-2015-3196](https://github.com/Live-Hack-CVE/CVE-2015-3196)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3196">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3196">
- [nidhi7598/OPENSSL_1.0.1g_CVE-2015-3196](https://github.com/nidhi7598/OPENSSL_1.0.1g_CVE-2015-3196)	<img alt="forks" src="https://img.shields.io/github/forks/nidhi7598/OPENSSL_1.0.1g_CVE-2015-3196">	<img alt="stars" src="https://img.shields.io/github/stars/nidhi7598/OPENSSL_1.0.1g_CVE-2015-3196">

---
## CVE-2015-3195 (2015-12-06T20:59:00)
> The ASN1_TFLG_COMBINE implementation in crypto/asn1/tasn_dec.c in OpenSSL before 0.9.8zh, 1.0.0 before 1.0.0t, 1.0.1 before 1.0.1q, and 1.0.2 before 1.0.2e mishandles errors caused by malformed X509_ATTRIBUTE data, which allows remote attackers to obtain sensitive information from process memory by triggering a decoding failure in a PKCS#7 or CMS application.
- [Live-Hack-CVE/CVE-2015-3195](https://github.com/Live-Hack-CVE/CVE-2015-3195)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3195">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3195">
- [Live-Hack-CVE/CVE-2015-3195](https://github.com/Live-Hack-CVE/CVE-2015-3195)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3195">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3195">
- [Trinadh465/OpenSSL-1_0_1g_CVE-2015-3195](https://github.com/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3195)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3195">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3195">

---
## CVE-2015-3194 (2015-12-06T20:59:00)
> crypto/rsa/rsa_ameth.c in OpenSSL 1.0.1 before 1.0.1q and 1.0.2 before 1.0.2e allows remote attackers to cause a denial of service (NULL pointer dereference and application crash) via an RSA PSS ASN.1 signature that lacks a mask generation function parameter.
- [Trinadh465/OpenSSL-1_0_1g_CVE-2015-3194](https://github.com/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3194)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3194">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/OpenSSL-1_0_1g_CVE-2015-3194">

---
## CVE-2015-3193 (2015-12-06T20:59:00)
> The Montgomery squaring implementation in crypto/bn/asm/x86_64-mont5.pl in OpenSSL 1.0.2 before 1.0.2e on the x86_64 platform, as used by the BN_mod_exp function, mishandles carry propagation and produces incorrect output, which makes it easier for remote attackers to obtain sensitive private-key information via an attack against use of a (1) Diffie-Hellman (DH) or (2) Diffie-Hellman Ephemeral (DHE) ciphersuite.
- [Live-Hack-CVE/CVE-2015-3193](https://github.com/Live-Hack-CVE/CVE-2015-3193)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3193">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3193">

---
## CVE-2015-3152 (2016-05-16T10:59:00)
> Oracle MySQL before 5.7.3, Oracle MySQL Connector/C (aka libmysqlclient) before 6.1.3, and MariaDB before 5.5.44 use the --ssl option to mean that SSL is optional, which allows man-in-the-middle attackers to spoof servers via a cleartext-downgrade attack, aka a "BACKRONYM" attack.
- [Live-Hack-CVE/CVE-2015-3152](https://github.com/Live-Hack-CVE/CVE-2015-3152)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3152">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3152">

---
## CVE-2015-3145 (2015-04-24T14:59:00)
> The sanitize_cookie_path function in cURL and libcurl 7.31.0 through 7.41.0 does not properly calculate an index, which allows remote attackers to cause a denial of service (out-of-bounds write and crash) or possibly have other unspecified impact via a cookie path containing only a double-quote character.
- [Serz999/CVE-2015-3145](https://github.com/Serz999/CVE-2015-3145)	<img alt="forks" src="https://img.shields.io/github/forks/Serz999/CVE-2015-3145">	<img alt="stars" src="https://img.shields.io/github/stars/Serz999/CVE-2015-3145">

---
## CVE-2015-2305 (2015-03-30T10:59:00)
> Integer overflow in the regcomp implementation in the Henry Spencer BSD regex library (aka rxspencer) alpha3.8.g5 on 32-bit platforms, as used in NetBSD through 6.1.5 and other products, might allow context-dependent attackers to execute arbitrary code via a large regular expression that leads to a heap-based buffer overflow.
- [Live-Hack-CVE/CVE-2015-2305](https://github.com/Live-Hack-CVE/CVE-2015-2305)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-2305">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-2305">

---
## CVE-2015-2301 (2015-03-30T10:59:00)
> Use-after-free vulnerability in the phar_rename_archive function in phar_object.c in PHP before 5.5.22 and 5.6.x before 5.6.6 allows remote attackers to cause a denial of service or possibly have unspecified other impact via vectors that trigger an attempted renaming of a Phar archive to the name of an existing file.
- [Live-Hack-CVE/CVE-2015-2301](https://github.com/Live-Hack-CVE/CVE-2015-2301)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-2301">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-2301">

---
## CVE-2015-2291 (2017-08-09T18:29:00)
> (1) IQVW32.sys before 1.3.1.0 and (2) IQVW64.sys before 1.3.1.0 in the Intel Ethernet diagnostics driver for Windows allows local users to cause a denial of service or possibly execute arbitrary code with kernel privileges via a crafted (a) 0x80862013, (b) 0x8086200B, (c) 0x8086200F, or (d) 0x80862007 IOCTL call.
- [Exploitables/CVE-2015-2291](https://github.com/Exploitables/CVE-2015-2291)	<img alt="forks" src="https://img.shields.io/github/forks/Exploitables/CVE-2015-2291">	<img alt="stars" src="https://img.shields.io/github/stars/Exploitables/CVE-2015-2291">
- [hfiref0x/KDU](https://github.com/hfiref0x/KDU)	<img alt="forks" src="https://img.shields.io/github/forks/hfiref0x/KDU">	<img alt="stars" src="https://img.shields.io/github/stars/hfiref0x/KDU">
- [Tare05/Intel-CVE-2015-2291](https://github.com/Tare05/Intel-CVE-2015-2291)	<img alt="forks" src="https://img.shields.io/github/forks/Tare05/Intel-CVE-2015-2291">	<img alt="stars" src="https://img.shields.io/github/stars/Tare05/Intel-CVE-2015-2291">

---
## CVE-2015-2166 (2015-04-06T15:59:00)
> Directory traversal vulnerability in the Instance Monitor in Ericsson Drutt Mobile Service Delivery Platform (MSDP) 4, 5, and 6 allows remote attackers to read arbitrary files via a ..%2f (dot dot encoded slash) in the default URI.
- [K3ysTr0K3R/CVE-2015-2166-EXPLOIT](https://github.com/K3ysTr0K3R/CVE-2015-2166-EXPLOIT)	<img alt="forks" src="https://img.shields.io/github/forks/K3ysTr0K3R/CVE-2015-2166-EXPLOIT">	<img alt="stars" src="https://img.shields.io/github/stars/K3ysTr0K3R/CVE-2015-2166-EXPLOIT">

---
## CVE-2015-20107 (2022-04-13T16:15:00)
> In Python (aka CPython) up to 3.10.8, the mailcap module does not add escape characters into commands discovered in the system mailcap file. This may allow attackers to inject shell commands into applications that call mailcap.findmatch with untrusted input (if they lack validation of user-provided filenames or arguments). The fix is also back-ported to 3.7, 3.8, 3.9
- [Live-Hack-CVE/CVE-2015-20107](https://github.com/Live-Hack-CVE/CVE-2015-20107)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-20107">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-20107">

---
## CVE-2015-1986 (2015-06-30T15:59:00)
> The server in IBM Tivoli Storage Manager FastBack 6.1 before 6.1.12 allows remote attackers to execute arbitrary commands via unspecified vectors, a different vulnerability than CVE-2015-1938.
- [3t3rn4lv01d/CVE-2015-1986](https://github.com/3t3rn4lv01d/CVE-2015-1986)	<img alt="forks" src="https://img.shields.io/github/forks/3t3rn4lv01d/CVE-2015-1986">	<img alt="stars" src="https://img.shields.io/github/stars/3t3rn4lv01d/CVE-2015-1986">

---
## CVE-2015-1805 (2015-08-08T10:59:00)
> The (1) pipe_read and (2) pipe_write implementations in fs/pipe.c in the Linux kernel before 3.16 do not properly consider the side effects of failed __copy_to_user_inatomic and __copy_from_user_inatomic calls, which allows local users to cause a denial of service (system crash) or possibly gain privileges via a crafted application, aka an "I/O vector array overrun."
- [ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805-May-2022-](https://github.com/ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805-May-2022-)	<img alt="forks" src="https://img.shields.io/github/forks/ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805-May-2022-">	<img alt="stars" src="https://img.shields.io/github/stars/ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805-May-2022-">
- [mobilelinux/iovy_root_research](https://github.com/mobilelinux/iovy_root_research)	<img alt="forks" src="https://img.shields.io/github/forks/mobilelinux/iovy_root_research">	<img alt="stars" src="https://img.shields.io/github/stars/mobilelinux/iovy_root_research">
- [dosomder/iovyroot](https://github.com/dosomder/iovyroot)	<img alt="forks" src="https://img.shields.io/github/forks/dosomder/iovyroot">	<img alt="stars" src="https://img.shields.io/github/stars/dosomder/iovyroot">
- [FloatingGuy/cve-2015-1805](https://github.com/FloatingGuy/cve-2015-1805)	<img alt="forks" src="https://img.shields.io/github/forks/FloatingGuy/cve-2015-1805">	<img alt="stars" src="https://img.shields.io/github/stars/FloatingGuy/cve-2015-1805">
- [panyu6325/CVE-2015-1805](https://github.com/panyu6325/CVE-2015-1805)	<img alt="forks" src="https://img.shields.io/github/forks/panyu6325/CVE-2015-1805">	<img alt="stars" src="https://img.shields.io/github/stars/panyu6325/CVE-2015-1805">
- [ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805](https://github.com/ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805)	<img alt="forks" src="https://img.shields.io/github/forks/ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805">	<img alt="stars" src="https://img.shields.io/github/stars/ireshchaminda1/Android-Privilege-Escalation-Remote-Access-Vulnerability-CVE-2015-1805">

---
## CVE-2015-1794 (2015-12-06T20:59:00)
> The ssl3_get_key_exchange function in ssl/s3_clnt.c in OpenSSL 1.0.2 before 1.0.2e allows remote servers to cause a denial of service (segmentation fault) via a zero p value in an anonymous Diffie-Hellman (DH) ServerKeyExchange message.
- [Live-Hack-CVE/CVE-2015-1794](https://github.com/Live-Hack-CVE/CVE-2015-1794)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1794">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1794">
- [Live-Hack-CVE/CVE-2015-1794](https://github.com/Live-Hack-CVE/CVE-2015-1794)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1794">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1794">

---
## CVE-2015-1792 (2015-06-12T19:59:00)
> The do_free_upto function in crypto/cms/cms_smime.c in OpenSSL before 0.9.8zg, 1.0.0 before 1.0.0s, 1.0.1 before 1.0.1n, and 1.0.2 before 1.0.2b allows remote attackers to cause a denial of service (infinite loop) via vectors that trigger a NULL value of a BIO data structure, as demonstrated by an unrecognized X.660 OID for a hash function.
- [Live-Hack-CVE/CVE-2015-1792](https://github.com/Live-Hack-CVE/CVE-2015-1792)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1792">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1792">
- [Live-Hack-CVE/CVE-2015-1792](https://github.com/Live-Hack-CVE/CVE-2015-1792)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1792">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1792">
- [Trinadh465/OpenSSL-1_0_1g_CVE-2015-1792](https://github.com/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1792)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1792">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1792">

---
## CVE-2015-1791 (2015-06-12T19:59:00)
> Race condition in the ssl3_get_new_session_ticket function in ssl/s3_clnt.c in OpenSSL before 0.9.8zg, 1.0.0 before 1.0.0s, 1.0.1 before 1.0.1n, and 1.0.2 before 1.0.2b, when used for a multi-threaded client, allows remote attackers to cause a denial of service (double free and application crash) or possibly have unspecified other impact by providing a NewSessionTicket during an attempt to reuse a ticket that had been obtained earlier.
- [Live-Hack-CVE/CVE-2015-1791](https://github.com/Live-Hack-CVE/CVE-2015-1791)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1791">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1791">
- [Live-Hack-CVE/CVE-2015-1791](https://github.com/Live-Hack-CVE/CVE-2015-1791)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1791">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1791">
- [Trinadh465/OpenSSL-1_0_1g_CVE-2015-1791](https://github.com/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1791)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1791">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1791">

---
## CVE-2015-1790 (2015-06-12T19:59:00)
> The PKCS7_dataDecodefunction in crypto/pkcs7/pk7_doit.c in OpenSSL before 0.9.8zg, 1.0.0 before 1.0.0s, 1.0.1 before 1.0.1n, and 1.0.2 before 1.0.2b allows remote attackers to cause a denial of service (NULL pointer dereference and application crash) via a PKCS#7 blob that uses ASN.1 encoding and lacks inner EncryptedContent data.
- [Live-Hack-CVE/CVE-2015-1790](https://github.com/Live-Hack-CVE/CVE-2015-1790)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1790">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1790">
- [Live-Hack-CVE/CVE-2015-1790](https://github.com/Live-Hack-CVE/CVE-2015-1790)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1790">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1790">
- [neominds/RIC190022](https://github.com/neominds/RIC190022)	<img alt="forks" src="https://img.shields.io/github/forks/neominds/RIC190022">	<img alt="stars" src="https://img.shields.io/github/stars/neominds/RIC190022">
- [Trinadh465/OpenSSL-1_0_1g_CVE-2015-1790](https://github.com/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1790)	<img alt="forks" src="https://img.shields.io/github/forks/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1790">	<img alt="stars" src="https://img.shields.io/github/stars/Trinadh465/OpenSSL-1_0_1g_CVE-2015-1790">

---
## CVE-2015-1789 (2015-06-12T19:59:00)
> The X509_cmp_time function in crypto/x509/x509_vfy.c in OpenSSL before 0.9.8zg, 1.0.0 before 1.0.0s, 1.0.1 before 1.0.1n, and 1.0.2 before 1.0.2b allows remote attackers to cause a denial of service (out-of-bounds read and application crash) via a crafted length field in ASN1_TIME data, as demonstrated by an attack against a server that supports client authentication with a custom verification callback.
- [Live-Hack-CVE/CVE-2015-1789](https://github.com/Live-Hack-CVE/CVE-2015-1789)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1789">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1789">
- [Live-Hack-CVE/CVE-2015-1789](https://github.com/Live-Hack-CVE/CVE-2015-1789)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1789">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1789">

---
## CVE-2015-1788 (2015-06-12T19:59:00)
> The BN_GF2m_mod_inv function in crypto/bn/bn_gf2m.c in OpenSSL before 0.9.8s, 1.0.0 before 1.0.0e, 1.0.1 before 1.0.1n, and 1.0.2 before 1.0.2b does not properly handle ECParameters structures in which the curve is over a malformed binary polynomial field, which allows remote attackers to cause a denial of service (infinite loop) via a session that uses an Elliptic Curve algorithm, as demonstrated by an attack against a server that supports client authentication.
- [Live-Hack-CVE/CVE-2015-1788](https://github.com/Live-Hack-CVE/CVE-2015-1788)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1788">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1788">
- [Live-Hack-CVE/CVE-2015-1788](https://github.com/Live-Hack-CVE/CVE-2015-1788)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1788">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1788">
- [pazhanivel07/OpenSSL_1_0_1g_CVE-2015-1788](https://github.com/pazhanivel07/OpenSSL_1_0_1g_CVE-2015-1788)	<img alt="forks" src="https://img.shields.io/github/forks/pazhanivel07/OpenSSL_1_0_1g_CVE-2015-1788">	<img alt="stars" src="https://img.shields.io/github/stars/pazhanivel07/OpenSSL_1_0_1g_CVE-2015-1788">

---
## CVE-2015-1787 (2015-03-19T22:59:00)
> The ssl3_get_client_key_exchange function in s3_srvr.c in OpenSSL 1.0.2 before 1.0.2a, when client authentication and an ephemeral Diffie-Hellman ciphersuite are enabled, allows remote attackers to cause a denial of service (daemon crash) via a ClientKeyExchange message with a length of zero.
- [Live-Hack-CVE/CVE-2015-1787](https://github.com/Live-Hack-CVE/CVE-2015-1787)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1787">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1787">
- [Live-Hack-CVE/CVE-2015-1787](https://github.com/Live-Hack-CVE/CVE-2015-1787)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1787">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1787">

---
## CVE-2015-1635 (2015-04-14T20:59:00)
> HTTP.sys in Microsoft Windows 7 SP1, Windows Server 2008 R2 SP1, Windows 8, Windows 8.1, and Windows Server 2012 Gold and R2 allows remote attackers to execute arbitrary code via crafted HTTP requests, aka "HTTP.sys Remote Code Execution Vulnerability."
- [Cappricio-Securities/CVE-2015-1635](https://github.com/Cappricio-Securities/CVE-2015-1635)	<img alt="forks" src="https://img.shields.io/github/forks/Cappricio-Securities/CVE-2015-1635">	<img alt="stars" src="https://img.shields.io/github/stars/Cappricio-Securities/CVE-2015-1635">
- [SkinAir/ms15-034-Scan](https://github.com/SkinAir/ms15-034-Scan)	<img alt="forks" src="https://img.shields.io/github/forks/SkinAir/ms15-034-Scan">	<img alt="stars" src="https://img.shields.io/github/stars/SkinAir/ms15-034-Scan">
- [w01ke/CVE-2015-1635-POC](https://github.com/w01ke/CVE-2015-1635-POC)	<img alt="forks" src="https://img.shields.io/github/forks/w01ke/CVE-2015-1635-POC">	<img alt="stars" src="https://img.shields.io/github/stars/w01ke/CVE-2015-1635-POC">
- [Sp3c73rSh4d0w/CVE-2015-1635](https://github.com/Sp3c73rSh4d0w/CVE-2015-1635)	<img alt="forks" src="https://img.shields.io/github/forks/Sp3c73rSh4d0w/CVE-2015-1635">	<img alt="stars" src="https://img.shields.io/github/stars/Sp3c73rSh4d0w/CVE-2015-1635">
- [Sp3c73rSh4d0w/CVE-2015-1635-POC](https://github.com/Sp3c73rSh4d0w/CVE-2015-1635-POC)	<img alt="forks" src="https://img.shields.io/github/forks/Sp3c73rSh4d0w/CVE-2015-1635-POC">	<img alt="stars" src="https://img.shields.io/github/stars/Sp3c73rSh4d0w/CVE-2015-1635-POC">
- [aedoo/CVE-2015-1635-POC](https://github.com/aedoo/CVE-2015-1635-POC)	<img alt="forks" src="https://img.shields.io/github/forks/aedoo/CVE-2015-1635-POC">	<img alt="stars" src="https://img.shields.io/github/stars/aedoo/CVE-2015-1635-POC">
- [technion/erlvulnscan](https://github.com/technion/erlvulnscan)	<img alt="forks" src="https://img.shields.io/github/forks/technion/erlvulnscan">	<img alt="stars" src="https://img.shields.io/github/stars/technion/erlvulnscan">
- [limkokholefork/CVE-2015-1635](https://github.com/limkokholefork/CVE-2015-1635)	<img alt="forks" src="https://img.shields.io/github/forks/limkokholefork/CVE-2015-1635">	<img alt="stars" src="https://img.shields.io/github/stars/limkokholefork/CVE-2015-1635">
- [bongbongco/MS15-034](https://github.com/bongbongco/MS15-034)	<img alt="forks" src="https://img.shields.io/github/forks/bongbongco/MS15-034">	<img alt="stars" src="https://img.shields.io/github/stars/bongbongco/MS15-034">
- [u0pattern/Remove-IIS-RIIS](https://github.com/u0pattern/Remove-IIS-RIIS)	<img alt="forks" src="https://img.shields.io/github/forks/u0pattern/Remove-IIS-RIIS">	<img alt="stars" src="https://img.shields.io/github/stars/u0pattern/Remove-IIS-RIIS">
- [wiredaem0n/chk-ms15-034](https://github.com/wiredaem0n/chk-ms15-034)	<img alt="forks" src="https://img.shields.io/github/forks/wiredaem0n/chk-ms15-034">	<img alt="stars" src="https://img.shields.io/github/stars/wiredaem0n/chk-ms15-034">
- [xPaw/HTTPsys](https://github.com/xPaw/HTTPsys)	<img alt="forks" src="https://img.shields.io/github/forks/xPaw/HTTPsys">	<img alt="stars" src="https://img.shields.io/github/stars/xPaw/HTTPsys">
- [neu5ron/cve_2015-1635](https://github.com/neu5ron/cve_2015-1635)	<img alt="forks" src="https://img.shields.io/github/forks/neu5ron/cve_2015-1635">	<img alt="stars" src="https://img.shields.io/github/stars/neu5ron/cve_2015-1635">
- [Zx7ffa4512-Python/Project-CVE-2015-1635](https://github.com/Zx7ffa4512-Python/Project-CVE-2015-1635)	<img alt="forks" src="https://img.shields.io/github/forks/Zx7ffa4512-Python/Project-CVE-2015-1635">	<img alt="stars" src="https://img.shields.io/github/stars/Zx7ffa4512-Python/Project-CVE-2015-1635">

---
## CVE-2015-1578 (2015-02-11T19:59:00)
> Multiple open redirect vulnerabilities in u5CMS before 3.9.4 allow remote attackers to redirect users to arbitrary web sites and conduct phishing attacks via a URL in the (1) pidvesa cookie to u5admin/pidvesa.php or (2) uri parameter to u5admin/meta2.php. <a href="http://cwe.mitre.org/data/definitions/601.html">CWE-601: URL Redirection to Untrusted Site ('Open Redirect')</a>
- [Zeppperoni/CVE-2015-1578](https://github.com/Zeppperoni/CVE-2015-1578)	<img alt="forks" src="https://img.shields.io/github/forks/Zeppperoni/CVE-2015-1578">	<img alt="stars" src="https://img.shields.io/github/stars/Zeppperoni/CVE-2015-1578">

---
## CVE-2015-1427 (2015-02-17T15:59:00)
> The Groovy scripting engine in Elasticsearch before 1.3.8 and 1.4.x before 1.4.3 allows remote attackers to bypass the sandbox protection mechanism and execute arbitrary shell commands via a crafted script.
- [xpgdgit/CVE-2015-1427](https://github.com/xpgdgit/CVE-2015-1427)	<img alt="forks" src="https://img.shields.io/github/forks/xpgdgit/CVE-2015-1427">	<img alt="stars" src="https://img.shields.io/github/stars/xpgdgit/CVE-2015-1427">
- [cved-sources/cve-2015-1427](https://github.com/cved-sources/cve-2015-1427)	<img alt="forks" src="https://img.shields.io/github/forks/cved-sources/cve-2015-1427">	<img alt="stars" src="https://img.shields.io/github/stars/cved-sources/cve-2015-1427">
- [h3inzzz/cve2015_1427](https://github.com/h3inzzz/cve2015_1427)	<img alt="forks" src="https://img.shields.io/github/forks/h3inzzz/cve2015_1427">	<img alt="stars" src="https://img.shields.io/github/stars/h3inzzz/cve2015_1427">
- [cyberharsh/Groovy-scripting-engine-CVE-2015-1427](https://github.com/cyberharsh/Groovy-scripting-engine-CVE-2015-1427)	<img alt="forks" src="https://img.shields.io/github/forks/cyberharsh/Groovy-scripting-engine-CVE-2015-1427">	<img alt="stars" src="https://img.shields.io/github/stars/cyberharsh/Groovy-scripting-engine-CVE-2015-1427">
- [t0kx/exploit-CVE-2015-1427](https://github.com/t0kx/exploit-CVE-2015-1427)	<img alt="forks" src="https://img.shields.io/github/forks/t0kx/exploit-CVE-2015-1427">	<img alt="stars" src="https://img.shields.io/github/stars/t0kx/exploit-CVE-2015-1427">

---
## CVE-2015-1421 (2015-03-16T10:59:00)
> Use-after-free vulnerability in the sctp_assoc_update function in net/sctp/associola.c in the Linux kernel before 3.18.8 allows remote attackers to cause a denial of service (slab corruption and panic) or possibly have unspecified other impact by triggering an INIT collision that leads to improper handling of shared-key data.
- [Live-Hack-CVE/CVE-2015-1421](https://github.com/Live-Hack-CVE/CVE-2015-1421)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1421">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1421">

---
## CVE-2015-1397 (2015-04-29T22:59:00)
> SQL injection vulnerability in the getCsvFile function in the Mage_Adminhtml_Block_Widget_Grid class in Magento Community Edition (CE) 1.9.1.0 and Enterprise Edition (EE) 1.14.1.0 allows remote administrators to execute arbitrary SQL commands via the popularity[field_expr] parameter when the popularity[from] or popularity[to] parameter is set.
- [WHOISshuvam/CVE-2015-1397](https://github.com/WHOISshuvam/CVE-2015-1397)	<img alt="forks" src="https://img.shields.io/github/forks/WHOISshuvam/CVE-2015-1397">	<img alt="stars" src="https://img.shields.io/github/stars/WHOISshuvam/CVE-2015-1397">
- [tmatejicek/CVE-2015-1397](https://github.com/tmatejicek/CVE-2015-1397)	<img alt="forks" src="https://img.shields.io/github/forks/tmatejicek/CVE-2015-1397">	<img alt="stars" src="https://img.shields.io/github/stars/tmatejicek/CVE-2015-1397">
- [47Cid/Magento-Shoplift-SQLI](https://github.com/47Cid/Magento-Shoplift-SQLI)	<img alt="forks" src="https://img.shields.io/github/forks/47Cid/Magento-Shoplift-SQLI">	<img alt="stars" src="https://img.shields.io/github/stars/47Cid/Magento-Shoplift-SQLI">
- [Wytchwulf/CVE-2015-1397-Magento-Shoplift](https://github.com/Wytchwulf/CVE-2015-1397-Magento-Shoplift)	<img alt="forks" src="https://img.shields.io/github/forks/Wytchwulf/CVE-2015-1397-Magento-Shoplift">	<img alt="stars" src="https://img.shields.io/github/stars/Wytchwulf/CVE-2015-1397-Magento-Shoplift">

---
## CVE-2015-1352 (2015-03-30T10:59:00)
> The build_tablename function in pgsql.c in the PostgreSQL (aka pgsql) extension in PHP through 5.6.7 does not validate token extraction for table names, which allows remote attackers to cause a denial of service (NULL pointer dereference and application crash) via a crafted name.
- [Live-Hack-CVE/CVE-2015-1352](https://github.com/Live-Hack-CVE/CVE-2015-1352)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1352">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1352">

---
## CVE-2015-1328 (2016-11-28T03:59:00)
> The overlayfs implementation in the linux (aka Linux kernel) package before 3.19.0-21.21 in Ubuntu through 15.04 does not properly check permissions for file creation in the upper filesystem directory, which allows local users to obtain root access by leveraging a configuration in which overlayfs is permitted in an arbitrary mount namespace.
- [poxicity/CVE-2015-1328](https://github.com/poxicity/CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/poxicity/CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/poxicity/CVE-2015-1328">
- [0x1ns4n3/CVE-2015-1328-GoldenEye](https://github.com/0x1ns4n3/CVE-2015-1328-GoldenEye)	<img alt="forks" src="https://img.shields.io/github/forks/0x1ns4n3/CVE-2015-1328-GoldenEye">	<img alt="stars" src="https://img.shields.io/github/stars/0x1ns4n3/CVE-2015-1328-GoldenEye">
- [notlikethis/CVE-2015-1328](https://github.com/notlikethis/CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/notlikethis/CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/notlikethis/CVE-2015-1328">
- [SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328](https://github.com/SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328">
- [poxicity/CVE-2015-1328](https://github.com/poxicity/CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/poxicity/CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/poxicity/CVE-2015-1328">
- [BlackFrog-hub/cve-2015-1328](https://github.com/BlackFrog-hub/cve-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/BlackFrog-hub/cve-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/BlackFrog-hub/cve-2015-1328">

---
## CVE-2015-10057 (2023-01-16T19:15:00)
> A vulnerability was found in Little Apps Little Software Stats. It has been declared as critical. Affected by this vulnerability is an unknown functionality of the file inc/class.securelogin.php of the component Password Reset Handler. The manipulation leads to improper access controls. Upgrading to version 0.2 is able to address this issue. The name of the patch is 07ba8273a9311d1383f3686ac7cb32f20770ab1e. It is recommended to upgrade the affected component. The identifier VDB-218401 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10057](https://github.com/Live-Hack-CVE/CVE-2015-10057)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10057">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10057">

---
## CVE-2015-10056 (2023-01-16T19:15:00)
> A vulnerability was found in 2071174A vinylmap. It has been classified as critical. Affected is the function contact of the file recordstoreapp/views.py. The manipulation leads to sql injection. The name of the patch is b07b79a1e92cc62574ba0492cce000ef4a7bd25f. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-218400.
- [Live-Hack-CVE/CVE-2015-10056](https://github.com/Live-Hack-CVE/CVE-2015-10056)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10056">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10056">

---
## CVE-2015-10055 (2023-01-16T18:15:00)
> A vulnerability was found in PictureThisWebServer and classified as critical. This issue affects the function router.post of the file routes/user.js. The manipulation of the argument username/password leads to sql injection. The name of the patch is 68b9dc346e88b494df00d88c7d058e96820e1479. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-218399.
- [Live-Hack-CVE/CVE-2015-10055](https://github.com/Live-Hack-CVE/CVE-2015-10055)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10055">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10055">

---
## CVE-2015-10054 (2023-01-16T18:15:00)
> A vulnerability, which was classified as critical, was found in githuis P2Manage. This affects the function Execute of the file PTwoManage/Database.cs. The manipulation of the argument sql leads to sql injection. The name of the patch is 717380aba80002414f82d93c770035198b7858cc. It is recommended to apply a patch to fix this issue. The identifier VDB-218397 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10054](https://github.com/Live-Hack-CVE/CVE-2015-10054)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10054">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10054">

---
## CVE-2015-10053 (2023-01-16T12:15:00)
> A vulnerability classified as critical has been found in prodigasistemas curupira up to 0.1.3. Affected is an unknown function of the file app/controllers/curupira/passwords_controller.rb. The manipulation leads to sql injection. Upgrading to version 0.1.4 is able to address this issue. The name of the patch is 93a9a77896bb66c949acb8e64bceafc74bc8c271. It is recommended to upgrade the affected component. VDB-218394 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10053](https://github.com/Live-Hack-CVE/CVE-2015-10053)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10053">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10053">

---
## CVE-2015-10052 (2023-01-15T19:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability, which was classified as problematic, was found in calesanz gibb-modul-151. This affects the function bearbeiten/login. The manipulation leads to open redirect. It is possible to initiate the attack remotely. The name of the patch is 88a517dc19443081210c804b655e72770727540d. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-218379. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2015-10052](https://github.com/Live-Hack-CVE/CVE-2015-10052)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10052">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10052">

---
## CVE-2015-10051 (2023-01-15T18:15:00)
> A vulnerability, which was classified as critical, has been found in bony2023 Discussion-Board. Affected by this issue is the function display_all_replies of the file functions/main.php. The manipulation of the argument str leads to sql injection. The name of the patch is 26439bc4c63632d63ba89ebc0f149b25a9010361. It is recommended to apply a patch to fix this issue. VDB-218378 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10051](https://github.com/Live-Hack-CVE/CVE-2015-10051)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10051">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10051">

---
## CVE-2015-10050 (2023-01-15T18:15:00)
> A vulnerability was found in brandonfire miRNA_Database_by_PHP_MySql. It has been declared as critical. This vulnerability affects the function __construct/select_single_rna/count_rna of the file inc/model.php. The manipulation leads to sql injection. The name of the patch is 307c5d510841e6142ddcbbdbb93d0e8a0dc3fd6a. It is recommended to apply a patch to fix this issue. VDB-218374 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10050](https://github.com/Live-Hack-CVE/CVE-2015-10050)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10050">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10050">

---
## CVE-2015-10049 (2023-01-15T18:15:00)
> A vulnerability was found in Overdrive Eletrônica course-builder up to 1.7.x and classified as problematic. Affected by this issue is some unknown functionality of the file coursebuilder/modules/oeditor/oeditor.html. The manipulation leads to cross site scripting. The attack may be launched remotely. Upgrading to version 1.8.0 is able to address this issue. The name of the patch is e39645fd714adb7e549908780235911ae282b21b. It is recommended to upgrade the affected component. The identifier of this vulnerability is VDB-218372.
- [Live-Hack-CVE/CVE-2015-10049](https://github.com/Live-Hack-CVE/CVE-2015-10049)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10049">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10049">

---
## CVE-2015-10048 (2023-01-15T10:15:00)
> A vulnerability was found in bmattoso desafio_buzz_woody. It has been rated as critical. This issue affects some unknown processing. The manipulation leads to sql injection. The name of the patch is cb8220cbae06082c969b1776fcb2fdafb3a1006b. It is recommended to apply a patch to fix this issue. The identifier VDB-218357 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10048](https://github.com/Live-Hack-CVE/CVE-2015-10048)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10048">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10048">

---
## CVE-2015-10047 (2023-01-15T10:15:00)
> A vulnerability was found in KYUUBl school-register. It has been classified as critical. This affects an unknown part of the file src/DBManager.java. The manipulation leads to sql injection. The name of the patch is 1cf7e01b878aee923f2b22cc2535c71a680e4c30. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-218355.
- [Live-Hack-CVE/CVE-2015-10047](https://github.com/Live-Hack-CVE/CVE-2015-10047)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10047">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10047">

---
## CVE-2015-10046 (2023-01-15T10:15:00)
> A vulnerability has been found in lolfeedback and classified as critical. Affected by this vulnerability is an unknown functionality. The manipulation leads to sql injection. The name of the patch is 6cf0b5f2228cd8765f734badd37910051000f2b2. It is recommended to apply a patch to fix this issue. The identifier VDB-218353 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10046](https://github.com/Live-Hack-CVE/CVE-2015-10046)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10046">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10046">

---
## CVE-2015-10045 (2023-01-15T10:15:00)
> A vulnerability, which was classified as critical, was found in tutrantta project_todolist. Affected is the function getAffectedRows/where/insert/update in the library library/Database.php. The manipulation leads to sql injection. The name of the patch is 194a0411bbe11aa4813f13c66b9e8ea403539141. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-218352.
- [Live-Hack-CVE/CVE-2015-10045](https://github.com/Live-Hack-CVE/CVE-2015-10045)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10045">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10045">

---
## CVE-2015-10044 (2023-01-15T10:15:00)
> A vulnerability classified as critical was found in gophergala sqldump. This vulnerability affects unknown code. The manipulation leads to sql injection. The name of the patch is 76db54e9073b5248b8863e71a63d66a32d567d21. It is recommended to apply a patch to fix this issue. VDB-218350 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10044](https://github.com/Live-Hack-CVE/CVE-2015-10044)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10044">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10044">

---
## CVE-2015-10043 (2023-01-14T21:15:00)
> A vulnerability, which was classified as critical, was found in abreen Apollo. This affects an unknown part. The manipulation of the argument file leads to path traversal. The name of the patch is 6206406630780bbd074aff34f4683fb764faba71. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-218307.
- [Live-Hack-CVE/CVE-2015-10043](https://github.com/Live-Hack-CVE/CVE-2015-10043)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10043">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10043">

---
## CVE-2015-10042 (2023-01-13T21:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability classified as critical was found in Dovgalyuk AIBattle. Affected by this vulnerability is the function registerUser of the file site/procedures.php. The manipulation of the argument postLogin leads to sql injection. The name of the patch is 448e9880aac18ae7832f8d065e03e46ce0f1d3e3. It is recommended to apply a patch to fix this issue. The identifier VDB-218305 was assigned to this vulnerability. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2015-10042](https://github.com/Live-Hack-CVE/CVE-2015-10042)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10042">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10042">

---
## CVE-2015-10041 (2023-01-13T20:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability classified as critical has been found in Dovgalyuk AIBattle. Affected is the function sendComments of the file site/procedures.php. The manipulation of the argument text leads to sql injection. The name of the patch is e3aa4d0900167641d41cbccf53909229f00381c9. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-218304. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2015-10041](https://github.com/Live-Hack-CVE/CVE-2015-10041)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10041">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10041">

---
## CVE-2015-10040 (2023-01-13T20:15:00)
> A vulnerability was found in gitlearn. It has been declared as problematic. This vulnerability affects the function getGrade/getOutOf of the file scripts/config.sh of the component Escape Sequence Handler. The manipulation leads to injection. The attack can be initiated remotely. The name of the patch is 3faa5deaa509012069afe75cd03c21bda5050a64. It is recommended to apply a patch to fix this issue. VDB-218302 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10040](https://github.com/Live-Hack-CVE/CVE-2015-10040)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10040">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10040">

---
## CVE-2015-10035 (2023-01-09T21:15:00)
> A vulnerability was found in gperson angular-test-reporter and classified as critical. This issue affects the function getProjectTables/addTest of the file rest-server/data-server.js. The manipulation leads to sql injection. The name of the patch is a29d8ae121b46ebfa96a55a9106466ab2ef166ae. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-217715.
- [Live-Hack-CVE/CVE-2015-10035](https://github.com/Live-Hack-CVE/CVE-2015-10035)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10035">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10035">

---
## CVE-2015-10034 (2023-01-09T21:15:00)
> A vulnerability has been found in j-nowak workout-organizer and classified as critical. This vulnerability affects unknown code. The manipulation leads to sql injection. The name of the patch is 13cd6c3d1210640bfdb39872b2bb3597aa991279. It is recommended to apply a patch to fix this issue. VDB-217714 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10034](https://github.com/Live-Hack-CVE/CVE-2015-10034)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10034">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10034">
- [andrenasx/CVE-2015-10034-Test](https://github.com/andrenasx/CVE-2015-10034-Test)	<img alt="forks" src="https://img.shields.io/github/forks/andrenasx/CVE-2015-10034-Test">	<img alt="stars" src="https://img.shields.io/github/stars/andrenasx/CVE-2015-10034-Test">
- [andrenasx/CVE-2015-10034](https://github.com/andrenasx/CVE-2015-10034)	<img alt="forks" src="https://img.shields.io/github/forks/andrenasx/CVE-2015-10034">	<img alt="stars" src="https://img.shields.io/github/stars/andrenasx/CVE-2015-10034">
- [hheeyywweellccoommee/CVE-2015-10034-Test-lazmv](https://github.com/hheeyywweellccoommee/CVE-2015-10034-Test-lazmv)	<img alt="forks" src="https://img.shields.io/github/forks/hheeyywweellccoommee/CVE-2015-10034-Test-lazmv">	<img alt="stars" src="https://img.shields.io/github/stars/hheeyywweellccoommee/CVE-2015-10034-Test-lazmv">
- [hheeyywweellccoommee/CVE-2015-10034-posua](https://github.com/hheeyywweellccoommee/CVE-2015-10034-posua)	<img alt="forks" src="https://img.shields.io/github/forks/hheeyywweellccoommee/CVE-2015-10034-posua">	<img alt="stars" src="https://img.shields.io/github/stars/hheeyywweellccoommee/CVE-2015-10034-posua">
- [andrenasx/CVE-2015-10034](https://github.com/andrenasx/CVE-2015-10034)	<img alt="forks" src="https://img.shields.io/github/forks/andrenasx/CVE-2015-10034">	<img alt="stars" src="https://img.shields.io/github/stars/andrenasx/CVE-2015-10034">
- [andrenasx/CVE-2015-10034-Test](https://github.com/andrenasx/CVE-2015-10034-Test)	<img alt="forks" src="https://img.shields.io/github/forks/andrenasx/CVE-2015-10034-Test">	<img alt="stars" src="https://img.shields.io/github/stars/andrenasx/CVE-2015-10034-Test">
- [andrenasx/CVE-2015-10034](https://github.com/andrenasx/CVE-2015-10034)	<img alt="forks" src="https://img.shields.io/github/forks/andrenasx/CVE-2015-10034">	<img alt="stars" src="https://img.shields.io/github/stars/andrenasx/CVE-2015-10034">
- [hheeyywweellccoommee/CVE-2015-10034-akdfu](https://github.com/hheeyywweellccoommee/CVE-2015-10034-akdfu)	<img alt="forks" src="https://img.shields.io/github/forks/hheeyywweellccoommee/CVE-2015-10034-akdfu">	<img alt="stars" src="https://img.shields.io/github/stars/hheeyywweellccoommee/CVE-2015-10034-akdfu">
- [andrenasx/CVE-2015-10034](https://github.com/andrenasx/CVE-2015-10034)	<img alt="forks" src="https://img.shields.io/github/forks/andrenasx/CVE-2015-10034">	<img alt="stars" src="https://img.shields.io/github/stars/andrenasx/CVE-2015-10034">

---
## CVE-2015-10033 (2023-01-09T21:15:00)
> A vulnerability, which was classified as problematic, was found in jvvlee MerlinsBoard. This affects an unknown part of the component Grade Handler. The manipulation leads to improper authorization. The name of the patch is 134f5481e2914b7f096cd92a22b1e6bcb8e6dfe5. It is recommended to apply a patch to fix this issue. The identifier VDB-217713 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10033](https://github.com/Live-Hack-CVE/CVE-2015-10033)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10033">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10033">

---
## CVE-2015-10032 (2023-01-09T09:15:00)
> A vulnerability was found in HealthMateWeb. It has been declared as problematic. Affected by this vulnerability is an unknown functionality of the file createaccount.php. The manipulation of the argument username/password/first_name/last_name/company/phone leads to cross site scripting. The attack can be launched remotely. The name of the patch is 472776c25b1046ecaf962c46fed7c713c72c28e3. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-217663.
- [Live-Hack-CVE/CVE-2015-10032](https://github.com/Live-Hack-CVE/CVE-2015-10032)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10032">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10032">

---
## CVE-2015-10031 (2023-01-08T17:15:00)
> A vulnerability classified as critical was found in purpleparrots 491-Project. This vulnerability affects unknown code of the file update.php of the component Highscore Handler. The manipulation leads to sql injection. The name of the patch is a812a5e4cf72f2a635a716086fe1ee2b8fa0b1ab. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217648.
- [Live-Hack-CVE/CVE-2015-10031](https://github.com/Live-Hack-CVE/CVE-2015-10031)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10031">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10031">

---
## CVE-2015-10030 (2023-01-08T10:15:00)
> A vulnerability has been found in SUKOHI Surpass and classified as critical. This vulnerability affects unknown code of the file src/Sukohi/Surpass/Surpass.php. The manipulation of the argument dir leads to pathname traversal. Upgrading to version 1.0.0 is able to address this issue. The name of the patch is d22337d453a2a14194cdb02bf12cdf9d9f827aa7. It is recommended to upgrade the affected component. VDB-217642 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10030](https://github.com/Live-Hack-CVE/CVE-2015-10030)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10030">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10030">

---
## CVE-2015-10029 (2023-01-07T20:15:00)
> A vulnerability classified as problematic was found in kelvinmo simplexrd up to 3.1.0. This vulnerability affects unknown code of the file simplexrd/simplexrd.class.php. The manipulation leads to xml external entity reference. Upgrading to version 3.1.1 is able to address this issue. The name of the patch is 4c9f2e028523ed705b555eca2c18c64e71f1a35d. It is recommended to upgrade the affected component. VDB-217630 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10029](https://github.com/Live-Hack-CVE/CVE-2015-10029)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10029">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10029">

---
## CVE-2015-10028 (2023-01-07T19:15:00)
> A vulnerability has been found in ss15-this-is-sparta and classified as problematic. This vulnerability affects unknown code of the file js/roomElement.js of the component Main Page. The manipulation leads to cross site scripting. The attack can be initiated remotely. The name of the patch is ba2f71ad3a46e5949ee0c510b544fa4ea973baaa. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217624.
- [Live-Hack-CVE/CVE-2015-10028](https://github.com/Live-Hack-CVE/CVE-2015-10028)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10028">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10028">

---
## CVE-2015-10027 (2023-01-07T17:15:00)
> A vulnerability, which was classified as problematic, has been found in hydrian TTRSS-Auth-LDAP. Affected by this issue is some unknown functionality of the component Username Handler. The manipulation leads to ldap injection. Upgrading to version 2.0b1 is able to address this issue. The name of the patch is a7f7a5a82d9202a5c40d606a5c519ba61b224eb8. It is recommended to upgrade the affected component. VDB-217622 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10027](https://github.com/Live-Hack-CVE/CVE-2015-10027)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10027">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10027">

---
## CVE-2015-10026 (2023-01-07T13:15:00)
> A vulnerability was found in tiredtyrant flairbot. It has been declared as critical. This vulnerability affects unknown code of the file flair.py. The manipulation leads to sql injection. The name of the patch is 5e112b68c6faad1d4699d02c1ebbb7daf48ef8fb. It is recommended to apply a patch to fix this issue. VDB-217618 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10026](https://github.com/Live-Hack-CVE/CVE-2015-10026)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10026">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10026">

---
## CVE-2015-10025 (2023-01-07T13:15:00)
> A vulnerability has been found in luelista miniConf up to 1.7.6 and classified as problematic. Affected by this vulnerability is an unknown functionality of the file miniConf/MessageView.cs of the component URL Scanning. The manipulation leads to denial of service. Upgrading to version 1.7.7 and 1.8.0 is able to address this issue. The name of the patch is c06c2e5116c306e4e1bc79779f0eda2d1182f655. It is recommended to upgrade the affected component. The associated identifier of this vulnerability is VDB-217615.
- [Live-Hack-CVE/CVE-2015-10025](https://github.com/Live-Hack-CVE/CVE-2015-10025)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10025">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10025">

---
## CVE-2015-10024 (2023-01-07T13:15:00)
> A vulnerability classified as critical was found in hoffie larasync. This vulnerability affects unknown code of the file repository/content/file_storage.go. The manipulation leads to path traversal. The name of the patch is 776bad422f4bd4930d09491711246bbeb1be9ba5. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217612.
- [Live-Hack-CVE/CVE-2015-10024](https://github.com/Live-Hack-CVE/CVE-2015-10024)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10024">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10024">

---
## CVE-2015-10023 (2023-01-07T12:15:00)
> A vulnerability classified as critical has been found in Fumon trello-octometric. This affects the function main of the file metrics-ui/server/srv.go. The manipulation of the argument num leads to sql injection. The name of the patch is a1f1754933fbf21e2221fbc671c81a47de6a04ef. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-217611.
- [Live-Hack-CVE/CVE-2015-10023](https://github.com/Live-Hack-CVE/CVE-2015-10023)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10023">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10023">

---
## CVE-2015-10022 (2023-01-07T12:15:00)
> A vulnerability was found in IISH nlgis2. It has been declared as critical. Affected by this vulnerability is an unknown functionality of the file scripts/etl/custom_import.pl. The manipulation leads to sql injection. The name of the patch is 8bdb6fcf7209584eaf1232437f0f53e735b2b34c. It is recommended to apply a patch to fix this issue. The identifier VDB-217609 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10022](https://github.com/Live-Hack-CVE/CVE-2015-10022)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10022">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10022">

---
## CVE-2015-10021 (2023-01-07T12:15:00)
> A vulnerability was found in ritterim definely. It has been classified as problematic. Affected is an unknown function of the file src/database.js. The manipulation leads to cross site scripting. It is possible to launch the attack remotely. The name of the patch is b31a022ba4d8d17148445a13ebb5a42ad593dbaa. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217608.
- [Live-Hack-CVE/CVE-2015-10021](https://github.com/Live-Hack-CVE/CVE-2015-10021)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10021">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10021">

---
## CVE-2015-10020 (2023-01-14T21:15:00)
> A vulnerability has been found in ssn2013 cis450Project and classified as critical. This vulnerability affects the function addUser of the file HeatMapServer/src/com/datformers/servlet/AddAppUser.java. The manipulation leads to sql injection. The name of the patch is 39b495011437a105c7670e17e071f99195b4922e. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-218380.
- [Live-Hack-CVE/CVE-2015-10020](https://github.com/Live-Hack-CVE/CVE-2015-10020)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10020">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10020">

---
## CVE-2015-10019 (2023-01-07T09:15:00)
> A vulnerability, which was classified as problematic, has been found in foxoverflow MySimplifiedSQL. This issue affects some unknown processing of the file MySimplifiedSQL_Examples.php. The manipulation of the argument FirstName/LastName leads to cross site scripting. The attack may be initiated remotely. The name of the patch is 3b7481c72786f88041b7c2d83bb4f219f77f1293. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-217595.
- [Live-Hack-CVE/CVE-2015-10019](https://github.com/Live-Hack-CVE/CVE-2015-10019)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10019">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10019">

---
## CVE-2015-10018 (2023-01-06T13:15:00)
> A vulnerability has been found in DBRisinajumi d2files and classified as critical. Affected by this vulnerability is the function actionUpload/actionDownloadFile of the file controllers/D2filesController.php. The manipulation leads to sql injection. Upgrading to version 1.0.0 is able to address this issue. The name of the patch is b5767f2ec9d0f3cbfda7f13c84740e2179c90574. It is recommended to upgrade the affected component. The identifier VDB-217561 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10018](https://github.com/Live-Hack-CVE/CVE-2015-10018)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10018">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10018">

---
## CVE-2015-10017 (2023-01-06T11:15:00)
> A vulnerability has been found in HPI-Information-Systems ProLOD and classified as critical. This vulnerability affects unknown code. The manipulation of the argument this leads to sql injection. The name of the patch is 3f710905458d49c77530bd3cbcd8960457566b73. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217552.
- [Live-Hack-CVE/CVE-2015-10017](https://github.com/Live-Hack-CVE/CVE-2015-10017)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10017">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10017">

---
## CVE-2015-10016 (2023-01-06T10:15:00)
> A vulnerability, which was classified as critical, has been found in jeff-kelley opensim-utils. Affected by this issue is the function DatabaseForRegion of the file regionscrits.php. The manipulation of the argument region leads to sql injection. The name of the patch is c29e5c729a833a29dbf5b1e505a0553fe154575e. It is recommended to apply a patch to fix this issue. VDB-217550 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10016](https://github.com/Live-Hack-CVE/CVE-2015-10016)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10016">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10016">

---
## CVE-2015-10015 (2023-01-05T15:15:00)
> A vulnerability, which was classified as critical, has been found in glidernet ogn-live. This issue affects some unknown processing. The manipulation leads to sql injection. The name of the patch is bc0f19965f760587645583b7624d66a260946e01. It is recommended to apply a patch to fix this issue. The associated identifier of this vulnerability is VDB-217487.
- [Live-Hack-CVE/CVE-2015-10015](https://github.com/Live-Hack-CVE/CVE-2015-10015)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10015">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10015">

---
## CVE-2015-10014 (2023-01-05T14:15:00)
> A vulnerability classified as critical has been found in arekk uke. This affects an unknown part of the file lib/uke/finder.rb. The manipulation leads to sql injection. The name of the patch is 52fd3b2d0bc16227ef57b7b98a3658bb67c1833f. It is recommended to apply a patch to fix this issue. The identifier VDB-217485 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10014](https://github.com/Live-Hack-CVE/CVE-2015-10014)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10014">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10014">

---
## CVE-2015-10013 (2023-01-05T10:15:00)
> A vulnerability was found in WebDevStudios taxonomy-switcher Plugin up to 1.0.3. It has been classified as problematic. Affected is the function taxonomy_switcher_init of the file taxonomy-switcher.php. The manipulation leads to cross site scripting. It is possible to launch the attack remotely. Upgrading to version 1.0.4 is able to address this issue. It is recommended to upgrade the affected component. VDB-217446 is the identifier assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10013](https://github.com/Live-Hack-CVE/CVE-2015-10013)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10013">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10013">

---
## CVE-2015-10012 (2023-01-03T09:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability was found in sumocoders FrameworkUserBundle up to 1.3.x. It has been rated as problematic. Affected by this issue is some unknown functionality of the file Resources/views/Security/login.html.twig. The manipulation leads to information exposure through error message. Upgrading to version 1.4.0 is able to address this issue. The name of the patch is abe4993390ba9bd7821ab12678270556645f94c8. It is recommended to upgrade the affected component. The identifier of this vulnerability is VDB-217268. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2015-10012](https://github.com/Live-Hack-CVE/CVE-2015-10012)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10012">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10012">

---
## CVE-2015-10011 (2023-01-02T22:15:00)
> A vulnerability classified as problematic has been found in OpenDNS OpenResolve. This affects an unknown part of the file resolverapi/endpoints.py. The manipulation leads to improper output neutralization for logs. The name of the patch is 9eba6ba5abd89d0e36a008921eb307fcef8c5311. It is recommended to apply a patch to fix this issue. The identifier VDB-217197 was assigned to this vulnerability.
- [Live-Hack-CVE/CVE-2015-10011](https://github.com/Live-Hack-CVE/CVE-2015-10011)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10011">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10011">

---
## CVE-2015-10009 (2023-01-02T16:15:00)
> A vulnerability was found in nterchange up to 4.1.0. It has been rated as critical. This issue affects the function getContent of the file app/controllers/code_caller_controller.php. The manipulation of the argument q with the input %5C%27%29;phpinfo%28%29;/* leads to code injection. The exploit has been disclosed to the public and may be used. Upgrading to version 4.1.1 is able to address this issue. The name of the patch is fba7d89176fba8fe289edd58835fe45080797d99. It is recommended to upgrade the affected component. The associated identifier of this vulnerability is VDB-217187.
- [Live-Hack-CVE/CVE-2015-10009](https://github.com/Live-Hack-CVE/CVE-2015-10009)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10009">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10009">

---
## CVE-2015-10008 (2023-01-02T11:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability was found in 82Flex WEIPDCRM. It has been classified as critical. This affects an unknown part. The manipulation leads to sql injection. It is possible to initiate the attack remotely. The name of the patch is 43bad79392332fa39e31b95268e76fbda9fec3a4. It is recommended to apply a patch to fix this issue. The identifier VDB-217185 was assigned to this vulnerability. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2015-10008](https://github.com/Live-Hack-CVE/CVE-2015-10008)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10008">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10008">

---
## CVE-2015-10007 (2023-01-02T11:15:00)
> ** UNSUPPPORTED WHEN ASSIGNED **** UNSUPPORTED WHEN ASSIGNED ** A vulnerability was found in 82Flex WEIPDCRM and classified as problematic. Affected by this issue is some unknown functionality. The manipulation leads to cross site scripting. The attack may be launched remotely. The name of the patch is 43bad79392332fa39e31b95268e76fbda9fec3a4. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217184. NOTE: This vulnerability only affects products that are no longer supported by the maintainer.
- [Live-Hack-CVE/CVE-2015-10007](https://github.com/Live-Hack-CVE/CVE-2015-10007)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10007">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10007">

---
## CVE-2015-10006 (2023-01-01T17:15:00)
> A vulnerability, which was classified as problematic, has been found in admont28 Ingnovarq. Affected by this issue is some unknown functionality of the file app/controller/insertarSliderAjax.php. The manipulation of the argument imagetitle leads to cross site scripting. The attack may be launched remotely. The name of the patch is 9d18a39944d79dfedacd754a742df38f99d3c0e2. It is recommended to apply a patch to fix this issue. The identifier of this vulnerability is VDB-217172.
- [Live-Hack-CVE/CVE-2015-10006](https://github.com/Live-Hack-CVE/CVE-2015-10006)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10006">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10006">

---
## CVE-2015-10005 (2022-12-27T09:15:00)
> A vulnerability was found in markdown-it up to 2.x. It has been classified as problematic. Affected is an unknown function of the file lib/common/html_re.js. The manipulation leads to inefficient regular expression complexity. Upgrading to version 3.0.0 is able to address this issue. The name of the patch is 89c8620157d6e38f9872811620d25138fc9d1b0d. It is recommended to upgrade the affected component. The identifier of this vulnerability is VDB-216852.
- [Live-Hack-CVE/CVE-2015-10005](https://github.com/Live-Hack-CVE/CVE-2015-10005)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10005">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10005">

---
## CVE-2015-0505 (2015-04-16T16:59:00)
> Unspecified vulnerability in Oracle MySQL Server 5.5.42 and earlier, and 5.6.23 and earlier, allows remote authenticated users to affect availability via vectors related to DDL.
- [Live-Hack-CVE/CVE-2015-0505](https://github.com/Live-Hack-CVE/CVE-2015-0505)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0505">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0505">

---
## CVE-2015-0382 (2015-01-21T18:59:00)
> Unspecified vulnerability in Oracle MySQL Server 5.5.40 and earlier and 5.6.21 and earlier allows remote attackers to affect availability via unknown vectors related to Server : Replication, a different vulnerability than CVE-2015-0381.
- [Live-Hack-CVE/CVE-2015-0382](https://github.com/Live-Hack-CVE/CVE-2015-0382)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0382">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0382">

---
## CVE-2015-0381 (2015-01-21T18:59:00)
> Unspecified vulnerability in Oracle MySQL Server 5.5.40 and earlier and 5.6.21 and earlier allows remote attackers to affect availability via unknown vectors related to Server : Replication, a different vulnerability than CVE-2015-0382.
- [Live-Hack-CVE/CVE-2015-0381](https://github.com/Live-Hack-CVE/CVE-2015-0381)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0381">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0381">

---
## CVE-2015-0291 (2015-03-19T22:59:00)
> The sigalgs implementation in t1_lib.c in OpenSSL 1.0.2 before 1.0.2a allows remote attackers to cause a denial of service (NULL pointer dereference and daemon crash) by using an invalid signature_algorithms extension in the ClientHello message during a renegotiation.
- [Live-Hack-CVE/CVE-2015-0291](https://github.com/Live-Hack-CVE/CVE-2015-0291)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0291">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0291">
- [Live-Hack-CVE/CVE-2015-0291](https://github.com/Live-Hack-CVE/CVE-2015-0291)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0291">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0291">

---
## CVE-2015-0290 (2015-03-19T22:59:00)
> The multi-block feature in the ssl3_write_bytes function in s3_pkt.c in OpenSSL 1.0.2 before 1.0.2a on 64-bit x86 platforms with AES NI support does not properly handle certain non-blocking I/O cases, which allows remote attackers to cause a denial of service (pointer corruption and application crash) via unspecified vectors.
- [Live-Hack-CVE/CVE-2015-0290](https://github.com/Live-Hack-CVE/CVE-2015-0290)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0290">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0290">

---
## CVE-2015-0288 (2015-03-19T22:59:00)
> The X509_to_X509_REQ function in crypto/x509/x509_req.c in OpenSSL before 0.9.8zf, 1.0.0 before 1.0.0r, 1.0.1 before 1.0.1m, and 1.0.2 before 1.0.2a might allow attackers to cause a denial of service (NULL pointer dereference and application crash) via an invalid certificate key.
- [Live-Hack-CVE/CVE-2015-0288](https://github.com/Live-Hack-CVE/CVE-2015-0288)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0288">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0288">

---
## CVE-2015-0286 (2015-03-19T22:59:00)
> The ASN1_TYPE_cmp function in crypto/asn1/a_type.c in OpenSSL before 0.9.8zf, 1.0.0 before 1.0.0r, 1.0.1 before 1.0.1m, and 1.0.2 before 1.0.2a does not properly perform boolean-type comparisons, which allows remote attackers to cause a denial of service (invalid read operation and application crash) via a crafted X.509 certificate to an endpoint that uses the certificate-verification feature.
- [Live-Hack-CVE/CVE-2015-0286](https://github.com/Live-Hack-CVE/CVE-2015-0286)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0286">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0286">

---
## CVE-2015-0285 (2015-03-19T22:59:00)
> The ssl3_client_hello function in s3_clnt.c in OpenSSL 1.0.2 before 1.0.2a does not ensure that the PRNG is seeded before proceeding with a handshake, which makes it easier for remote attackers to defeat cryptographic protection mechanisms by sniffing the network and then conducting a brute-force attack.
- [Live-Hack-CVE/CVE-2015-0285](https://github.com/Live-Hack-CVE/CVE-2015-0285)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0285">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0285">

---
## CVE-2015-0278 (2015-05-18T15:59:00)
> libuv before 0.10.34 does not properly drop group privileges, which allows context-dependent attackers to gain privileges via unspecified vectors.
- [Live-Hack-CVE/CVE-2015-0278](https://github.com/Live-Hack-CVE/CVE-2015-0278)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0278">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0278">

---
## CVE-2015-0209 (2015-03-19T22:59:00)
> Use-after-free vulnerability in the d2i_ECPrivateKey function in crypto/ec/ec_asn1.c in OpenSSL before 0.9.8zf, 1.0.0 before 1.0.0r, 1.0.1 before 1.0.1m, and 1.0.2 before 1.0.2a might allow remote attackers to cause a denial of service (memory corruption and application crash) or possibly have unspecified other impact via a malformed Elliptic Curve (EC) private-key file that is improperly handled during import.
- [Live-Hack-CVE/CVE-2015-0209](https://github.com/Live-Hack-CVE/CVE-2015-0209)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0209">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0209">
- [Live-Hack-CVE/CVE-2015-0209](https://github.com/Live-Hack-CVE/CVE-2015-0209)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0209">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0209">

---
## CVE-2015-0208 (2015-03-19T22:59:00)
> The ASN.1 signature-verification implementation in the rsa_item_verify function in crypto/rsa/rsa_ameth.c in OpenSSL 1.0.2 before 1.0.2a allows remote attackers to cause a denial of service (NULL pointer dereference and application crash) via crafted RSA PSS parameters to an endpoint that uses the certificate-verification feature.
- [Live-Hack-CVE/CVE-2015-0208](https://github.com/Live-Hack-CVE/CVE-2015-0208)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0208">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0208">
- [Live-Hack-CVE/CVE-2015-0208](https://github.com/Live-Hack-CVE/CVE-2015-0208)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-0208">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-0208">

---
## CVE-2015-0205 (2015-01-09T02:59:00)
> The ssl3_get_cert_verify function in s3_srvr.c in OpenSSL 1.0.0 before 1.0.0p and 1.0.1 before 1.0.1k accepts client authentication with a Diffie-Hellman (DH) certificate without requiring a CertificateVerify message, which allows remote attackers to obtain access without knowledge of a private key via crafted TLS Handshake Protocol traffic to a server that recognizes a Certification Authority with DH support.
- [saurabh2088/OpenSSL_1_0_1g_CVE-2015-0205](https://github.com/saurabh2088/OpenSSL_1_0_1g_CVE-2015-0205)	<img alt="forks" src="https://img.shields.io/github/forks/saurabh2088/OpenSSL_1_0_1g_CVE-2015-0205">	<img alt="stars" src="https://img.shields.io/github/stars/saurabh2088/OpenSSL_1_0_1g_CVE-2015-0205">
