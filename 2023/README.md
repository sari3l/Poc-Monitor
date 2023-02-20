# 2023 List

---
## CVE-2023-44268 ()
> 
- [agathanon/cve-2023-44268](https://github.com/agathanon/cve-2023-44268)	<img alt="forks" src="https://img.shields.io/github/forks/agathanon/cve-2023-44268">	<img alt="stars" src="https://img.shields.io/github/stars/agathanon/cve-2023-44268">

---
## CVE-2023-25194 (2023-02-07T20:15:00)
> A possible security vulnerability has been identified in Apache Kafka Connect. This requires access to a Kafka Connect worker, and the ability to create/modify connectors on it with an arbitrary Kafka client SASL JAAS config and a SASL-based security protocol, which has been possible on Kafka Connect clusters since Apache Kafka 2.3.0. When configuring the connector via the Kafka Connect REST API, an authenticated operator can set the `sasl.jaas.config` property for any of the connector's Kafka clients to "com.sun.security.auth.module.JndiLoginModule", which can be done via the `producer.override.sasl.jaas.config`, `consumer.override.sasl.jaas.config`, or `admin.override.sasl.jaas.config` properties. This will allow the server to connect to the attacker's LDAP server and deserialize the LDAP response, which the attacker can use to execute java deserialization gadget chains on the Kafka connect server. Attacker can cause unrestricted deserialization of untrusted data (or) RCE vulnerability when there are gadgets in the classpath. Since Apache Kafka 3.0.0, users are allowed to specify these properties in connector configurations for Kafka Connect clusters running with out-of-the-box configurations. Before Apache Kafka 3.0.0, users may not specify these properties unless the Kafka Connect cluster has been reconfigured with a connector client override policy that permits them. Since Apache Kafka 3.4.0, we have added a system property ("-Dorg.apache.kafka.disallowed.login.modules") to disable the problematic login modules usage in SASL JAAS configuration. Also by default "com.sun.security.auth.module.JndiLoginModule" is disabled in Apache Kafka 3.4.0. We advise the Kafka Connect users to validate connector configurations and only allow trusted JNDI configurations. Also examine connector dependencies for vulnerable versions and either upgrade their connectors, upgrading that specific dependency, or removing the connectors as options for remediation. Finally, in addition to leveraging the "org.apache.kafka.disallowed.login.modules" system property, Kafka Connect users can also implement their own connector client config override policy, which can be used to control which Kafka client properties can be overridden directly in a connector config and which cannot.
- [ohnonoyesyes/CVE-2023-25194](https://github.com/ohnonoyesyes/CVE-2023-25194)	<img alt="forks" src="https://img.shields.io/github/forks/ohnonoyesyes/CVE-2023-25194">	<img alt="stars" src="https://img.shields.io/github/stars/ohnonoyesyes/CVE-2023-25194">

---
## CVE-2023-25136 (2023-02-03T06:15:00)
> OpenSSH server (sshd) 9.1 introduced a double-free vulnerability during options.kex_algorithms handling. This is fixed in OpenSSH 9.2. The double free can be triggered by an unauthenticated attacker in the default configuration; however, the vulnerability discoverer reports that "exploiting this vulnerability will not be easy."
- [jfrog/jfrog-CVE-2023-25136-OpenSSH_Double-Free](https://github.com/jfrog/jfrog-CVE-2023-25136-OpenSSH_Double-Free)	<img alt="forks" src="https://img.shields.io/github/forks/jfrog/jfrog-CVE-2023-25136-OpenSSH_Double-Free">	<img alt="stars" src="https://img.shields.io/github/stars/jfrog/jfrog-CVE-2023-25136-OpenSSH_Double-Free">
- [ticofookfook/CVE-2023-25136](https://github.com/ticofookfook/CVE-2023-25136)	<img alt="forks" src="https://img.shields.io/github/forks/ticofookfook/CVE-2023-25136">	<img alt="stars" src="https://img.shields.io/github/stars/ticofookfook/CVE-2023-25136">

---
## CVE-2023-24610 (2023-02-01T14:15:00)
> NOSH 4a5cfdb allows remote authenticated users to execute PHP arbitrary code via the "practice logo" upload feature. The client-side checks can be bypassed. This may allow attackers to steal Protected Health Information because the product is for health charting.
- [abbisQQ/CVE-2023-24610](https://github.com/abbisQQ/CVE-2023-24610)	<img alt="forks" src="https://img.shields.io/github/forks/abbisQQ/CVE-2023-24610">	<img alt="stars" src="https://img.shields.io/github/stars/abbisQQ/CVE-2023-24610">

---
## CVE-2023-24055 (2023-01-22T04:15:00)
> ** DISPUTED ** KeePass through 2.53 (in a default installation) allows an attacker, who has write access to the XML configuration file, to obtain the cleartext passwords by adding an export trigger. NOTE: the vendor's position is that the password database is not intended to be secure against an attacker who has that level of access to the local PC.
- [deetl/CVE-2023-24055](https://github.com/deetl/CVE-2023-24055)	<img alt="forks" src="https://img.shields.io/github/forks/deetl/CVE-2023-24055">	<img alt="stars" src="https://img.shields.io/github/stars/deetl/CVE-2023-24055">
- [alt3kx/CVE-2023-24055_PoC](https://github.com/alt3kx/CVE-2023-24055_PoC)	<img alt="forks" src="https://img.shields.io/github/forks/alt3kx/CVE-2023-24055_PoC">	<img alt="stars" src="https://img.shields.io/github/stars/alt3kx/CVE-2023-24055_PoC">
- [Cyb3rtus/keepass_CVE-2023-24055_yara_rule](https://github.com/Cyb3rtus/keepass_CVE-2023-24055_yara_rule)	<img alt="forks" src="https://img.shields.io/github/forks/Cyb3rtus/keepass_CVE-2023-24055_yara_rule">	<img alt="stars" src="https://img.shields.io/github/stars/Cyb3rtus/keepass_CVE-2023-24055_yara_rule">
- [julesbozouklian/PoC_CVE-2023-24055](https://github.com/julesbozouklian/PoC_CVE-2023-24055)	<img alt="forks" src="https://img.shields.io/github/forks/julesbozouklian/PoC_CVE-2023-24055">	<img alt="stars" src="https://img.shields.io/github/stars/julesbozouklian/PoC_CVE-2023-24055">
- [julesbozouklian/PoC_CVE-2023-24055](https://github.com/julesbozouklian/PoC_CVE-2023-24055)	<img alt="forks" src="https://img.shields.io/github/forks/julesbozouklian/PoC_CVE-2023-24055">	<img alt="stars" src="https://img.shields.io/github/stars/julesbozouklian/PoC_CVE-2023-24055">
- [ATTACKnDEFEND/CVE-2023-24055](https://github.com/ATTACKnDEFEND/CVE-2023-24055)	<img alt="forks" src="https://img.shields.io/github/forks/ATTACKnDEFEND/CVE-2023-24055">	<img alt="stars" src="https://img.shields.io/github/stars/ATTACKnDEFEND/CVE-2023-24055">
- [digital-dev/KeePass-TriggerLess](https://github.com/digital-dev/KeePass-TriggerLess)	<img alt="forks" src="https://img.shields.io/github/forks/digital-dev/KeePass-TriggerLess">	<img alt="stars" src="https://img.shields.io/github/stars/digital-dev/KeePass-TriggerLess">
- [PyterSmithDarkGhost/CVE-2023-24055-PoC-KeePass-2.5x-](https://github.com/PyterSmithDarkGhost/CVE-2023-24055-PoC-KeePass-2.5x-)	<img alt="forks" src="https://img.shields.io/github/forks/PyterSmithDarkGhost/CVE-2023-24055-PoC-KeePass-2.5x-">	<img alt="stars" src="https://img.shields.io/github/stars/PyterSmithDarkGhost/CVE-2023-24055-PoC-KeePass-2.5x-">
- [zwlsix/KeePass-CVE-2023-24055](https://github.com/zwlsix/KeePass-CVE-2023-24055)	<img alt="forks" src="https://img.shields.io/github/forks/zwlsix/KeePass-CVE-2023-24055">	<img alt="stars" src="https://img.shields.io/github/stars/zwlsix/KeePass-CVE-2023-24055">
- [zwlsix/KeePass-CVE-2023-24055](https://github.com/zwlsix/KeePass-CVE-2023-24055)	<img alt="forks" src="https://img.shields.io/github/forks/zwlsix/KeePass-CVE-2023-24055">	<img alt="stars" src="https://img.shields.io/github/stars/zwlsix/KeePass-CVE-2023-24055">

---
## CVE-2023-23924 (2023-02-01T00:15:00)
> Dompdf is an HTML to PDF converter. The URI validation on dompdf 2.0.1 can be bypassed on SVG parsing by passing `<image>` tags with uppercase letters. This may lead to arbitrary object unserialize on PHP < 8, through the `phar` URL wrapper. An attacker can exploit the vulnerability to call arbitrary URL with arbitrary protocols, if they can provide a SVG file to dompdf. In PHP versions before 8.0.0, it leads to arbitrary unserialize, that will lead to the very least to an arbitrary file deletion and even remote code execution, depending on classes that are available.
- [motikan2010/CVE-2023-23924](https://github.com/motikan2010/CVE-2023-23924)	<img alt="forks" src="https://img.shields.io/github/forks/motikan2010/CVE-2023-23924">	<img alt="stars" src="https://img.shields.io/github/stars/motikan2010/CVE-2023-23924">

---
## CVE-2023-23752 (2023-02-16T17:15:00)
> An issue was discovered in Joomla! 4.0.0 through 4.2.7. An improper access check allows unauthorized access to webservice endpoints.
- [DanielRuf/CVE-2023-23752](https://github.com/DanielRuf/CVE-2023-23752)	<img alt="forks" src="https://img.shields.io/github/forks/DanielRuf/CVE-2023-23752">	<img alt="stars" src="https://img.shields.io/github/stars/DanielRuf/CVE-2023-23752">
- [YusinoMy/CVE-2023-23752](https://github.com/YusinoMy/CVE-2023-23752)	<img alt="forks" src="https://img.shields.io/github/forks/YusinoMy/CVE-2023-23752">	<img alt="stars" src="https://img.shields.io/github/stars/YusinoMy/CVE-2023-23752">
- [Saboor-Hakimi/CVE-2023-23752](https://github.com/Saboor-Hakimi/CVE-2023-23752)	<img alt="forks" src="https://img.shields.io/github/forks/Saboor-Hakimi/CVE-2023-23752">	<img alt="stars" src="https://img.shields.io/github/stars/Saboor-Hakimi/CVE-2023-23752">
- [WhiteOwl-Pub/CVE-2023-23752](https://github.com/WhiteOwl-Pub/CVE-2023-23752)	<img alt="forks" src="https://img.shields.io/github/forks/WhiteOwl-Pub/CVE-2023-23752">	<img alt="stars" src="https://img.shields.io/github/stars/WhiteOwl-Pub/CVE-2023-23752">
- [Vulnmachines/joomla_CVE-2023-23752](https://github.com/Vulnmachines/joomla_CVE-2023-23752)	<img alt="forks" src="https://img.shields.io/github/forks/Vulnmachines/joomla_CVE-2023-23752">	<img alt="stars" src="https://img.shields.io/github/stars/Vulnmachines/joomla_CVE-2023-23752">

---
## CVE-2023-23488 (2023-01-20T18:15:00)
> The Paid Memberships Pro WordPress Plugin, version < 2.9.8, is affected by an unauthenticated SQL injection vulnerability in the 'code' parameter of the '/pmpro/v1/order' REST route.
- [r3nt0n/CVE-2023-23488-PoC](https://github.com/r3nt0n/CVE-2023-23488-PoC)	<img alt="forks" src="https://img.shields.io/github/forks/r3nt0n/CVE-2023-23488-PoC">	<img alt="stars" src="https://img.shields.io/github/stars/r3nt0n/CVE-2023-23488-PoC">

---
## CVE-2023-23333 (2023-02-06T22:15:00)
> There is a command injection vulnerability in SolarView Compact through 6.00, attackers can execute commands by bypassing internal restrictions through downloader.php.
- [Timorlover/CVE-2023-23333](https://github.com/Timorlover/CVE-2023-23333)	<img alt="forks" src="https://img.shields.io/github/forks/Timorlover/CVE-2023-23333">	<img alt="stars" src="https://img.shields.io/github/stars/Timorlover/CVE-2023-23333">

---
## CVE-2023-232323 ()
> 
- [Shmily-ing/CVE-2023-232323](https://github.com/Shmily-ing/CVE-2023-232323)	<img alt="forks" src="https://img.shields.io/github/forks/Shmily-ing/CVE-2023-232323">	<img alt="stars" src="https://img.shields.io/github/stars/Shmily-ing/CVE-2023-232323">

---
## CVE-2023-23163 (2023-02-10T20:15:00)
> Art Gallery Management System Project v1.0 was discovered to contain a SQL injection vulnerability via the editid parameter.
- [rahulpatwari/CVE-2023-23163](https://github.com/rahulpatwari/CVE-2023-23163)	<img alt="forks" src="https://img.shields.io/github/forks/rahulpatwari/CVE-2023-23163">	<img alt="stars" src="https://img.shields.io/github/stars/rahulpatwari/CVE-2023-23163">

---
## CVE-2023-23162 (2023-02-10T20:15:00)
> Art Gallery Management System Project v1.0 was discovered to contain a SQL injection vulnerability via the cid parameter at product.php.
- [rahulpatwari/CVE-2023-23162](https://github.com/rahulpatwari/CVE-2023-23162)	<img alt="forks" src="https://img.shields.io/github/forks/rahulpatwari/CVE-2023-23162">	<img alt="stars" src="https://img.shields.io/github/stars/rahulpatwari/CVE-2023-23162">

---
## CVE-2023-23161 (2023-02-10T20:15:00)
> A reflected cross-site scripting (XSS) vulnerability in Art Gallery Management System Project v1.0 allows attackers to execute arbitrary web scripts or HTML via a crafted payload injected into the artname parameter under ART TYPE option in the navigation bar.
- [rahulpatwari/CVE-2023-23161](https://github.com/rahulpatwari/CVE-2023-23161)	<img alt="forks" src="https://img.shields.io/github/forks/rahulpatwari/CVE-2023-23161">	<img alt="stars" src="https://img.shields.io/github/stars/rahulpatwari/CVE-2023-23161">

---
## CVE-2023-23132 (2023-02-01T14:15:00)
> Selfwealth iOS mobile App 3.3.1 is vulnerable to Sensitive key disclosure. The application reveals hardcoded API keys.
- [l00neyhacker/CVE-2023-23132](https://github.com/l00neyhacker/CVE-2023-23132)	<img alt="forks" src="https://img.shields.io/github/forks/l00neyhacker/CVE-2023-23132">	<img alt="stars" src="https://img.shields.io/github/stars/l00neyhacker/CVE-2023-23132">

---
## CVE-2023-23131 (2023-02-01T14:15:00)
> Selfwealth iOS mobile App 3.3.1 is vulnerable to Insecure App Transport Security (ATS) Settings.
- [l00neyhacker/CVE-2023-23131](https://github.com/l00neyhacker/CVE-2023-23131)	<img alt="forks" src="https://img.shields.io/github/forks/l00neyhacker/CVE-2023-23131">	<img alt="stars" src="https://img.shields.io/github/stars/l00neyhacker/CVE-2023-23131">

---
## CVE-2023-23130 (2023-02-01T14:15:00)
> Connectwise Automate 2022.11 is vulnerable to Cleartext authentication. Authentication is being done via HTTP (cleartext) with SSL disabled.
- [l00neyhacker/CVE-2023-23130](https://github.com/l00neyhacker/CVE-2023-23130)	<img alt="forks" src="https://img.shields.io/github/forks/l00neyhacker/CVE-2023-23130">	<img alt="stars" src="https://img.shields.io/github/stars/l00neyhacker/CVE-2023-23130">

---
## CVE-2023-23128 (2023-02-01T14:15:00)
> Connectwise Control 22.8.10013.8329 is vulnerable to Cross Origin Resource Sharing (CORS).
- [l00neyhacker/CVE-2023-23128](https://github.com/l00neyhacker/CVE-2023-23128)	<img alt="forks" src="https://img.shields.io/github/forks/l00neyhacker/CVE-2023-23128">	<img alt="stars" src="https://img.shields.io/github/stars/l00neyhacker/CVE-2023-23128">

---
## CVE-2023-23127 (2023-02-01T14:15:00)
> In Connectwise Control 22.8.10013.8329, the login page does not implement HSTS headers therefore not enforcing HTTPS.
- [l00neyhacker/CVE-2023-23127](https://github.com/l00neyhacker/CVE-2023-23127)	<img alt="forks" src="https://img.shields.io/github/forks/l00neyhacker/CVE-2023-23127">	<img alt="stars" src="https://img.shields.io/github/stars/l00neyhacker/CVE-2023-23127">

---
## CVE-2023-23126 (2023-02-01T14:15:00)
> Connectwise Automate 2022.11 is vulnerable to Clickjacking. The login screen can be iframed and used to manipulate users to perform unintended actions.
- [l00neyhacker/CVE-2023-23126](https://github.com/l00neyhacker/CVE-2023-23126)	<img alt="forks" src="https://img.shields.io/github/forks/l00neyhacker/CVE-2023-23126">	<img alt="stars" src="https://img.shields.io/github/stars/l00neyhacker/CVE-2023-23126">

---
## CVE-2023-22960 (2023-01-23T21:15:00)
> Lexmark products through 2023-01-10 have Improper Control of Interaction Frequency.
- [t3l3machus/CVE-2023-22960](https://github.com/t3l3machus/CVE-2023-22960)	<img alt="forks" src="https://img.shields.io/github/forks/t3l3machus/CVE-2023-22960">	<img alt="stars" src="https://img.shields.io/github/stars/t3l3machus/CVE-2023-22960">
- [manas3c/CVE-2023-22960](https://github.com/manas3c/CVE-2023-22960)	<img alt="forks" src="https://img.shields.io/github/forks/manas3c/CVE-2023-22960">	<img alt="stars" src="https://img.shields.io/github/stars/manas3c/CVE-2023-22960">

---
## CVE-2023-22941 (2023-02-14T18:15:00)
> In Splunk Enterprise versions below 8.1.13, 8.2.10, and 9.0.4, an improperly-formatted ‘INGEST_EVAL’ parameter in a [Field Transformation](https://docs.splunk.com/Documentation/Splunk/latest/Knowledge/Managefieldtransforms) crashes the Splunk daemon (splunkd).
- [eduardosantos1989/CVE-2023-22941](https://github.com/eduardosantos1989/CVE-2023-22941)	<img alt="forks" src="https://img.shields.io/github/forks/eduardosantos1989/CVE-2023-22941">	<img alt="stars" src="https://img.shields.io/github/stars/eduardosantos1989/CVE-2023-22941">

---
## CVE-2023-22855 (2023-02-15T21:15:00)
> Kardex Mlog MCC 5.7.12+0-a203c2a213-master allows remote code execution. It spawns a web interface listening on port 8088. A user-controllable path is handed to a path-concatenation method (Path.Combine from .NET) without proper sanitisation. This yields the possibility of including local files, as well as remote files on SMB shares. If one provides a file with the extension .t4, it is rendered with the .NET templating engine mono/t4, which can execute code.
- [patrickhener/CVE-2023-22855](https://github.com/patrickhener/CVE-2023-22855)	<img alt="forks" src="https://img.shields.io/github/forks/patrickhener/CVE-2023-22855">	<img alt="stars" src="https://img.shields.io/github/stars/patrickhener/CVE-2023-22855">

---
## CVE-2023-22809 (2023-01-18T17:15:00)
> In Sudo before 1.9.12p2, the sudoedit (aka -e) feature mishandles extra arguments passed in the user-provided environment variables (SUDO_EDITOR, VISUAL, and EDITOR), allowing a local attacker to append arbitrary entries to the list of files to process. This can lead to privilege escalation. Affected versions are 1.8.0 through 1.9.12.p1. The problem exists because a user-specified editor may contain a "--" argument that defeats a protection mechanism, e.g., an EDITOR='vim -- /path/to/extra/file' value.
- [n3m1dotsys/CVE-2023-22809-sudoedit-privesc](https://github.com/n3m1dotsys/CVE-2023-22809-sudoedit-privesc)	<img alt="forks" src="https://img.shields.io/github/forks/n3m1dotsys/CVE-2023-22809-sudoedit-privesc">	<img alt="stars" src="https://img.shields.io/github/stars/n3m1dotsys/CVE-2023-22809-sudoedit-privesc">

---
## CVE-2023-2232323 ()
> 
- [Shmily-ing/CVE-2023-2232323](https://github.com/Shmily-ing/CVE-2023-2232323)	<img alt="forks" src="https://img.shields.io/github/forks/Shmily-ing/CVE-2023-2232323">	<img alt="stars" src="https://img.shields.io/github/stars/Shmily-ing/CVE-2023-2232323">

---
## CVE-2023-2222111 ()
> 
- [Shmily-ing/CVE-2023-2222111](https://github.com/Shmily-ing/CVE-2023-2222111)	<img alt="forks" src="https://img.shields.io/github/forks/Shmily-ing/CVE-2023-2222111">	<img alt="stars" src="https://img.shields.io/github/stars/Shmily-ing/CVE-2023-2222111">

---
## CVE-2023-21839 (2023-01-18T00:15:00)
> Vulnerability in the Oracle WebLogic Server product of Oracle Fusion Middleware (component: Core).  Supported versions that are affected are 12.2.1.3.0, 12.2.1.4.0 and  14.1.1.0.0. Easily exploitable vulnerability allows unauthenticated attacker with network access via T3, IIOP to compromise Oracle WebLogic Server.  Successful attacks of this vulnerability can result in  unauthorized access to critical data or complete access to all Oracle WebLogic Server accessible data. CVSS 3.1 Base Score 7.5 (Confidentiality impacts).  CVSS Vector: (CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N).
- [fakenews2025/CVE-2023-21839](https://github.com/fakenews2025/CVE-2023-21839)	<img alt="forks" src="https://img.shields.io/github/forks/fakenews2025/CVE-2023-21839">	<img alt="stars" src="https://img.shields.io/github/stars/fakenews2025/CVE-2023-21839">

---
## CVE-2023-21753 (2023-01-10T22:15:00)
> Event Tracing for Windows Information Disclosure Vulnerability. This CVE ID is unique from CVE-2023-21536.
- [timpen432/-Wh0Am1001-CVE-2023-21753](https://github.com/timpen432/-Wh0Am1001-CVE-2023-21753)	<img alt="forks" src="https://img.shields.io/github/forks/timpen432/-Wh0Am1001-CVE-2023-21753">	<img alt="stars" src="https://img.shields.io/github/stars/timpen432/-Wh0Am1001-CVE-2023-21753">

---
## CVE-2023-21608 (2023-01-18T19:15:00)
> Adobe Acrobat Reader versions 22.003.20282 (and earlier), 22.003.20281 (and earlier) and 20.005.30418 (and earlier) are affected by a Use After Free vulnerability that could result in arbitrary code execution in the context of the current user. Exploitation of this issue requires user interaction in that a victim must open a malicious file.
- [hacksysteam/CVE-2023-21608](https://github.com/hacksysteam/CVE-2023-21608)	<img alt="forks" src="https://img.shields.io/github/forks/hacksysteam/CVE-2023-21608">	<img alt="stars" src="https://img.shields.io/github/stars/hacksysteam/CVE-2023-21608">
- [PyterSmithDarkGhost/CVE-2023-21608-EXPLOIT](https://github.com/PyterSmithDarkGhost/CVE-2023-21608-EXPLOIT)	<img alt="forks" src="https://img.shields.io/github/forks/PyterSmithDarkGhost/CVE-2023-21608-EXPLOIT">	<img alt="stars" src="https://img.shields.io/github/stars/PyterSmithDarkGhost/CVE-2023-21608-EXPLOIT">
- [Malwareman007/CVE-2023-21608](https://github.com/Malwareman007/CVE-2023-21608)	<img alt="forks" src="https://img.shields.io/github/forks/Malwareman007/CVE-2023-21608">	<img alt="stars" src="https://img.shields.io/github/stars/Malwareman007/CVE-2023-21608">

---
## CVE-2023-0860 (2023-02-16T10:15:00)
> Improper Restriction of Excessive Authentication Attempts in GitHub repository modoboa/modoboa-installer prior to 2.0.4.
- [0xsu3ks/CVE-2023-0860](https://github.com/0xsu3ks/CVE-2023-0860)	<img alt="forks" src="https://img.shields.io/github/forks/0xsu3ks/CVE-2023-0860">	<img alt="stars" src="https://img.shields.io/github/stars/0xsu3ks/CVE-2023-0860">

---
## CVE-2023-0748 (2023-02-08T15:15:00)
> Open Redirect in GitHub repository btcpayserver/btcpayserver prior to 1.7.6.
- [gonzxph/CVE-2023-0748](https://github.com/gonzxph/CVE-2023-0748)	<img alt="forks" src="https://img.shields.io/github/forks/gonzxph/CVE-2023-0748">	<img alt="stars" src="https://img.shields.io/github/stars/gonzxph/CVE-2023-0748">

---
## CVE-2023-0669 (2023-02-06T20:15:00)
> Fortra (formerly, HelpSystems) GoAnywhere MFT suffers from a pre-authentication command injection vulnerability in the License Response Servlet due to deserializing an arbitrary attacker-controlled object. This issue was patched in version 7.1.2.
- [0xf4n9x/CVE-2023-0669](https://github.com/0xf4n9x/CVE-2023-0669)	<img alt="forks" src="https://img.shields.io/github/forks/0xf4n9x/CVE-2023-0669">	<img alt="stars" src="https://img.shields.io/github/stars/0xf4n9x/CVE-2023-0669">
- [cataiovita/CVE-2023-0669](https://github.com/cataiovita/CVE-2023-0669)	<img alt="forks" src="https://img.shields.io/github/forks/cataiovita/CVE-2023-0669">	<img alt="stars" src="https://img.shields.io/github/stars/cataiovita/CVE-2023-0669">
- [yosef0x01/CVE-2023-0669](https://github.com/yosef0x01/CVE-2023-0669)	<img alt="forks" src="https://img.shields.io/github/forks/yosef0x01/CVE-2023-0669">	<img alt="stars" src="https://img.shields.io/github/stars/yosef0x01/CVE-2023-0669">

---
## CVE-2023-0297 (2023-01-14T03:15:00)
> Code Injection in GitHub repository pyload/pyload prior to 0.5.0b3.dev31.
- [Small-ears/CVE-2023-0297](https://github.com/Small-ears/CVE-2023-0297)	<img alt="forks" src="https://img.shields.io/github/forks/Small-ears/CVE-2023-0297">	<img alt="stars" src="https://img.shields.io/github/stars/Small-ears/CVE-2023-0297">
- [bAuh0lz/CVE-2023-0297_Pre-auth_RCE_in_pyLoad](https://github.com/bAuh0lz/CVE-2023-0297_Pre-auth_RCE_in_pyLoad)	<img alt="forks" src="https://img.shields.io/github/forks/bAuh0lz/CVE-2023-0297_Pre-auth_RCE_in_pyLoad">	<img alt="stars" src="https://img.shields.io/github/stars/bAuh0lz/CVE-2023-0297_Pre-auth_RCE_in_pyLoad">

---
## CVE-2023-0179 ()
> 
- [TurtleARM/CVE-2023-0179-PoC](https://github.com/TurtleARM/CVE-2023-0179-PoC)	<img alt="forks" src="https://img.shields.io/github/forks/TurtleARM/CVE-2023-0179-PoC">	<img alt="stars" src="https://img.shields.io/github/stars/TurtleARM/CVE-2023-0179-PoC">

---
## CVE-2023-0045 ()
> 
- [es0j/CVE-2023-0045](https://github.com/es0j/CVE-2023-0045)	<img alt="forks" src="https://img.shields.io/github/forks/es0j/CVE-2023-0045">	<img alt="stars" src="https://img.shields.io/github/stars/es0j/CVE-2023-0045">
- [es0j/CVE-2023-0045](https://github.com/es0j/CVE-2023-0045)	<img alt="forks" src="https://img.shields.io/github/forks/es0j/CVE-2023-0045">	<img alt="stars" src="https://img.shields.io/github/stars/es0j/CVE-2023-0045">
- [missyes/CVE-2023-0045](https://github.com/missyes/CVE-2023-0045)	<img alt="forks" src="https://img.shields.io/github/forks/missyes/CVE-2023-0045">	<img alt="stars" src="https://img.shields.io/github/stars/missyes/CVE-2023-0045">
