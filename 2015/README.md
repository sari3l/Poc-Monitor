# 2015 List

---
## CVE-2015-8994 (2017-03-02T06:59:00)
> An issue was discovered in PHP 5.x and 7.x, when the configuration uses apache2handler/mod_php or php-fpm with OpCache enabled. With 5.x after 5.6.28 or 7.x after 7.0.13, the issue is resolved in a non-default configuration with the opcache.validate_permission=1 setting. The vulnerability details are as follows. In PHP SAPIs where PHP interpreters share a common parent process, Zend OpCache creates a shared memory object owned by the common parent during initialization. Child PHP processes inherit the SHM descriptor, using it to cache and retrieve compiled script bytecode ("opcode" in PHP jargon). Cache keys vary depending on configuration, but filename is a central key component, and compiled opcode can generally be run if a script's filename is known or can be guessed. Many common shared-hosting configurations change EUID in child processes to enforce privilege separation among hosted users (for example using mod_ruid2 for the Apache HTTP Server, or php-fpm user settings). In these scenarios, the default Zend OpCache behavior defeats script file permissions by sharing a single SHM cache among all child PHP processes. PHP scripts often contain sensitive information: Think of CMS configurations where reading or running another user's script usually means gaining privileges to the CMS database.
- [Live-Hack-CVE/CVE-2015-8994](https://github.com/Live-Hack-CVE/CVE-2015-8994)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-8994">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-8994">

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

---
## CVE-2015-6764 (2015-12-06T01:59:00)
> The BasicJsonStringifier::SerializeJSArray function in json-stringifier.h in the JSON stringifier in Google V8, as used in Google Chrome before 47.0.2526.73, improperly loads array elements, which allows remote attackers to cause a denial of service (out-of-bounds memory access) or possibly have unspecified other impact via crafted JavaScript code.
- [Live-Hack-CVE/CVE-2015-6764](https://github.com/Live-Hack-CVE/CVE-2015-6764)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-6764">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-6764">

---
## CVE-2015-5531 (2015-08-17T15:59:00)
> Directory traversal vulnerability in Elasticsearch before 1.6.1 allows remote attackers to read arbitrary files via unspecified vectors related to snapshot API calls.
- [xpgdgit/CVE-2015-5531](https://github.com/xpgdgit/CVE-2015-5531)	<img alt="forks" src="https://img.shields.io/github/forks/xpgdgit/CVE-2015-5531">	<img alt="stars" src="https://img.shields.io/github/stars/xpgdgit/CVE-2015-5531">
- [j-jasson/CVE-2015-5531-POC](https://github.com/j-jasson/CVE-2015-5531-POC)	<img alt="forks" src="https://img.shields.io/github/forks/j-jasson/CVE-2015-5531-POC">	<img alt="stars" src="https://img.shields.io/github/stars/j-jasson/CVE-2015-5531-POC">

---
## CVE-2015-5361 (2020-02-28T23:15:00)
> Background For regular, unencrypted FTP traffic, the FTP ALG can inspect the unencrypted control channel and open related sessions for the FTP data channel. These related sessions (gates) are specific to source and destination IPs and ports of client and server. The design intent of the ftps-extensions option (which is disabled by default) is to provide similar functionality when the SRX secures the FTP/FTPS client. As the control channel is encrypted, the FTP ALG cannot inspect the port specific information and will open a wider TCP data channel (gate) from client IP to server IP on all destination TCP ports. In FTP/FTPS client environments to an enterprise network or the Internet, this is the desired behavior as it allows firewall policy to be written to FTP/FTPS servers on well-known control ports without using a policy with destination IP ANY and destination port ANY. Issue The ftps-extensions option is not intended or recommended where the SRX secures the FTPS server, as the wide data channel session (gate) will allow the FTPS client temporary access to all TCP ports on the FTPS server. The data session is associated to the control channel and will be closed when the control channel session closes. Depending on the configuration of the FTPS server, supporting load-balancer, and SRX inactivity-timeout values, the server/load-balancer and SRX may keep the control channel open for an extended period of time, allowing an FTPS client access for an equal duration.? Note that the ftps-extensions option is not enabled by default.
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

---
## CVE-2015-3884 (2017-03-17T14:59:00)
> Unrestricted file upload vulnerability in the (1) myAccount, (2) projects, (3) tasks, (4) tickets, (5) discussions, (6) reports, and (7) scheduler pages in qdPM 8.3 allows remote attackers to execute arbitrary code by uploading a file with an executable extension, then accessing it via a direct request to the file in uploads/attachments/ or uploads/users/.
- [Live-Hack-CVE/CVE-2015-3884](https://github.com/Live-Hack-CVE/CVE-2015-3884)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3884">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3884">

---
## CVE-2015-3416 (2015-04-24T17:59:00)
> The sqlite3VXPrintf function in printf.c in SQLite before 3.8.9 does not properly handle precision and width values during floating-point conversions, which allows context-dependent attackers to cause a denial of service (integer overflow and stack-based buffer overflow) or possibly have unspecified other impact via large integers in a crafted printf function call in a SELECT statement.
- [Live-Hack-CVE/CVE-2015-3416](https://github.com/Live-Hack-CVE/CVE-2015-3416)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3416">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3416">

---
## CVE-2015-3197 (2016-02-15T02:59:00)
> ssl/s2_srvr.c in OpenSSL 1.0.1 before 1.0.1r and 1.0.2 before 1.0.2f does not prevent use of disabled ciphers, which makes it easier for man-in-the-middle attackers to defeat cryptographic protection mechanisms by performing computations on SSLv2 traffic, related to the get_client_master_key and get_client_hello functions.
- [Live-Hack-CVE/CVE-2015-3197](https://github.com/Live-Hack-CVE/CVE-2015-3197)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3197">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3197">

---
## CVE-2015-3196 (2015-12-06T20:59:00)
> ssl/s3_clnt.c in OpenSSL 1.0.0 before 1.0.0t, 1.0.1 before 1.0.1p, and 1.0.2 before 1.0.2d, when used for a multi-threaded client, writes the PSK identity hint to an incorrect data structure, which allows remote servers to cause a denial of service (race condition and double free) via a crafted ServerKeyExchange message.
- [Live-Hack-CVE/CVE-2015-3196](https://github.com/Live-Hack-CVE/CVE-2015-3196)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3196">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3196">

---
## CVE-2015-3195 (2015-12-06T20:59:00)
> The ASN1_TFLG_COMBINE implementation in crypto/asn1/tasn_dec.c in OpenSSL before 0.9.8zh, 1.0.0 before 1.0.0t, 1.0.1 before 1.0.1q, and 1.0.2 before 1.0.2e mishandles errors caused by malformed X509_ATTRIBUTE data, which allows remote attackers to obtain sensitive information from process memory by triggering a decoding failure in a PKCS#7 or CMS application.
- [Live-Hack-CVE/CVE-2015-3195](https://github.com/Live-Hack-CVE/CVE-2015-3195)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3195">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3195">
- [Live-Hack-CVE/CVE-2015-3195](https://github.com/Live-Hack-CVE/CVE-2015-3195)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-3195">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-3195">

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
## CVE-2015-20107 (2022-04-13T16:15:00)
> In Python (aka CPython) up to 3.10.8, the mailcap module does not add escape characters into commands discovered in the system mailcap file. This may allow attackers to inject shell commands into applications that call mailcap.findmatch with untrusted input (if they lack validation of user-provided filenames or arguments). The fix is also back-ported to 3.7, 3.8, 3.9
- [Live-Hack-CVE/CVE-2015-20107](https://github.com/Live-Hack-CVE/CVE-2015-20107)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-20107">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-20107">

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

---
## CVE-2015-1791 (2015-06-12T19:59:00)
> Race condition in the ssl3_get_new_session_ticket function in ssl/s3_clnt.c in OpenSSL before 0.9.8zg, 1.0.0 before 1.0.0s, 1.0.1 before 1.0.1n, and 1.0.2 before 1.0.2b, when used for a multi-threaded client, allows remote attackers to cause a denial of service (double free and application crash) or possibly have unspecified other impact by providing a NewSessionTicket during an attempt to reuse a ticket that had been obtained earlier.
- [Live-Hack-CVE/CVE-2015-1791](https://github.com/Live-Hack-CVE/CVE-2015-1791)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1791">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1791">
- [Live-Hack-CVE/CVE-2015-1791](https://github.com/Live-Hack-CVE/CVE-2015-1791)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1791">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1791">

---
## CVE-2015-1790 (2015-06-12T19:59:00)
> The PKCS7_dataDecodefunction in crypto/pkcs7/pk7_doit.c in OpenSSL before 0.9.8zg, 1.0.0 before 1.0.0s, 1.0.1 before 1.0.1n, and 1.0.2 before 1.0.2b allows remote attackers to cause a denial of service (NULL pointer dereference and application crash) via a PKCS#7 blob that uses ASN.1 encoding and lacks inner EncryptedContent data.
- [Live-Hack-CVE/CVE-2015-1790](https://github.com/Live-Hack-CVE/CVE-2015-1790)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1790">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1790">
- [Live-Hack-CVE/CVE-2015-1790](https://github.com/Live-Hack-CVE/CVE-2015-1790)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1790">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1790">

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

---
## CVE-2015-1787 (2015-03-19T22:59:00)
> The ssl3_get_client_key_exchange function in s3_srvr.c in OpenSSL 1.0.2 before 1.0.2a, when client authentication and an ephemeral Diffie-Hellman ciphersuite are enabled, allows remote attackers to cause a denial of service (daemon crash) via a ClientKeyExchange message with a length of zero.
- [Live-Hack-CVE/CVE-2015-1787](https://github.com/Live-Hack-CVE/CVE-2015-1787)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1787">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1787">
- [Live-Hack-CVE/CVE-2015-1787](https://github.com/Live-Hack-CVE/CVE-2015-1787)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-1787">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-1787">

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
## CVE-2015-1328 (2016-11-28T03:59:00)
> The overlayfs implementation in the linux (aka Linux kernel) package before 3.19.0-21.21 in Ubuntu through 15.04 does not properly check permissions for file creation in the upper filesystem directory, which allows local users to obtain root access by leveraging a configuration in which overlayfs is permitted in an arbitrary mount namespace.
- [poxicity/CVE-2015-1328](https://github.com/poxicity/CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/poxicity/CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/poxicity/CVE-2015-1328">
- [0x1ns4n3/CVE-2015-1328-GoldenEye](https://github.com/0x1ns4n3/CVE-2015-1328-GoldenEye)	<img alt="forks" src="https://img.shields.io/github/forks/0x1ns4n3/CVE-2015-1328-GoldenEye">	<img alt="stars" src="https://img.shields.io/github/stars/0x1ns4n3/CVE-2015-1328-GoldenEye">
- [notlikethis/CVE-2015-1328](https://github.com/notlikethis/CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/notlikethis/CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/notlikethis/CVE-2015-1328">
- [SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328](https://github.com/SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/SR7-HACKING/LINUX-VULNERABILITY-CVE-2015-1328">
- [poxicity/CVE-2015-1328](https://github.com/poxicity/CVE-2015-1328)	<img alt="forks" src="https://img.shields.io/github/forks/poxicity/CVE-2015-1328">	<img alt="stars" src="https://img.shields.io/github/stars/poxicity/CVE-2015-1328">

---
## CVE-2015-10005 (2022-12-27T09:15:00)
> A vulnerability was found in markdown-it up to 2.x. It has been classified as problematic. Affected is an unknown function of the file lib/common/html_re.js. The manipulation leads to inefficient regular expression complexity. Upgrading to version 3.0.0 is able to address this issue. The name of the patch is 89c8620157d6e38f9872811620d25138fc9d1b0d. It is recommended to upgrade the affected component. The identifier of this vulnerability is VDB-216852.
- [Live-Hack-CVE/CVE-2015-10005](https://github.com/Live-Hack-CVE/CVE-2015-10005)	<img alt="forks" src="https://img.shields.io/github/forks/Live-Hack-CVE/CVE-2015-10005">	<img alt="stars" src="https://img.shields.io/github/stars/Live-Hack-CVE/CVE-2015-10005">

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
