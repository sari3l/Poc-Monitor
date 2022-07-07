# 2020 List

---
## CVE-2020-8512 (2020-02-01T00:15:00)
> In IceWarp Webmail Server through 11.4.4.1, there is XSS in the /webmail/ color parameter.
- [trhacknon/CVE-2020-8512](https://github.com/trhacknon/CVE-2020-8512)	<img alt="forks" src="https://img.shields.io/github/forks/trhacknon/CVE-2020-8512">	<img alt="stars" src="https://img.shields.io/github/stars/trhacknon/CVE-2020-8512">

---
## CVE-2020-8163 (2020-07-02T19:15:00)
> The is a code injection vulnerability in versions of Rails prior to 5.0.1 that wouldallow an attacker who controlled the `locals` argument of a `render` call to perform a RCE.
- [lucasallan/CVE-2020-8163](https://github.com/lucasallan/CVE-2020-8163)	<img alt="forks" src="https://img.shields.io/github/forks/lucasallan/CVE-2020-8163">	<img alt="stars" src="https://img.shields.io/github/stars/lucasallan/CVE-2020-8163">
- [TKLinux966/CVE-2020-8163](https://github.com/TKLinux966/CVE-2020-8163)	<img alt="forks" src="https://img.shields.io/github/forks/TKLinux966/CVE-2020-8163">	<img alt="stars" src="https://img.shields.io/github/stars/TKLinux966/CVE-2020-8163">
- [h4ms1k/CVE-2020-8163](https://github.com/h4ms1k/CVE-2020-8163)	<img alt="forks" src="https://img.shields.io/github/forks/h4ms1k/CVE-2020-8163">	<img alt="stars" src="https://img.shields.io/github/stars/h4ms1k/CVE-2020-8163">

---
## CVE-2020-7473 (2020-05-07T14:15:00)
> In certain situations, all versions of Citrix ShareFile StorageZones (aka storage zones) Controller, including the most recent 5.10.x releases as of May 2020, allow unauthenticated attackers to access the documents and folders of ShareFile users. NOTE: unlike most CVEs, exploitability depends on the product version that was in use when a particular setup step was performed, NOT the product version that is in use during a current assessment of a CVE consumer's product inventory. Specifically, the vulnerability can be exploited if a storage zone was created by one of these product versions: 5.9.0, 5.8.0, 5.7.0, 5.6.0, 5.5.0, or earlier. This CVE differs from CVE-2020-8982 and CVE-2020-8983 but has essentially the same risk.
- [DimitriNL/CTX-CVE-2020-7473](https://github.com/DimitriNL/CTX-CVE-2020-7473)	<img alt="forks" src="https://img.shields.io/github/forks/DimitriNL/CTX-CVE-2020-7473">	<img alt="stars" src="https://img.shields.io/github/stars/DimitriNL/CTX-CVE-2020-7473">

---
## CVE-2020-6468 (2020-05-21T04:15:00)
> Type confusion in V8 in Google Chrome prior to 83.0.4103.61 allowed a remote attacker to potentially exploit heap corruption via a crafted HTML page.
- [kiks7/CVE-2020-6468-Chrome-Exploit](https://github.com/kiks7/CVE-2020-6468-Chrome-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/kiks7/CVE-2020-6468-Chrome-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/kiks7/CVE-2020-6468-Chrome-Exploit">
- [Goyotan/CVE-2020-6468-PoC](https://github.com/Goyotan/CVE-2020-6468-PoC)	<img alt="forks" src="https://img.shields.io/github/forks/Goyotan/CVE-2020-6468-PoC">	<img alt="stars" src="https://img.shields.io/github/stars/Goyotan/CVE-2020-6468-PoC">

---
## CVE-2020-5844 (2020-03-16T18:15:00)
> index.php?sec=godmode/extensions&sec2=extensions/files_repo in Pandora FMS v7.0 NG allows authenticated administrators to upload malicious PHP scripts, and execute them via base64 decoding of the file location. This affects v7.0NG.742_FIX_PERL2020.
- [UNICORDev/exploit-CVE-2020-5844](https://github.com/UNICORDev/exploit-CVE-2020-5844)	<img alt="forks" src="https://img.shields.io/github/forks/UNICORDev/exploit-CVE-2020-5844">	<img alt="stars" src="https://img.shields.io/github/stars/UNICORDev/exploit-CVE-2020-5844">

---
## CVE-2020-3580 (2020-10-21T19:15:00)
> Multiple vulnerabilities in the web services interface of Cisco Adaptive Security Appliance (ASA) Software and Cisco Firepower Threat Defense (FTD) Software could allow an unauthenticated, remote attacker to conduct cross-site scripting (XSS) attacks against a user of the web services interface of an affected device. The vulnerabilities are due to insufficient validation of user-supplied input by the web services interface of an affected device. An attacker could exploit these vulnerabilities by persuading a user of the interface to click a crafted link. A successful exploit could allow the attacker to execute arbitrary script code in the context of the interface or allow the attacker to access sensitive, browser-based information. Note: These vulnerabilities affect only specific AnyConnect and WebVPN configurations. For more information, see the Vulnerable Products section.
- [nxtexploit/CVE-2020-3580](https://github.com/nxtexploit/CVE-2020-3580)	<img alt="forks" src="https://img.shields.io/github/forks/nxtexploit/CVE-2020-3580">	<img alt="stars" src="https://img.shields.io/github/stars/nxtexploit/CVE-2020-3580">
- [cruxN3T/CVE-2020-3580](https://github.com/cruxN3T/CVE-2020-3580)	<img alt="forks" src="https://img.shields.io/github/forks/cruxN3T/CVE-2020-3580">	<img alt="stars" src="https://img.shields.io/github/stars/cruxN3T/CVE-2020-3580">

---
## CVE-2020-27786 (2020-12-11T05:15:00)
> A flaw was found in the Linux kernel’s implementation of MIDI, where an attacker with a local account and the permissions to issue ioctl commands to midi devices could trigger a use-after-free issue. A write to this specific memory while freed and before use causes the flow of execution to change and possibly allow for memory corruption or privilege escalation. The highest threat from this vulnerability is to confidentiality, integrity, as well as system availability.
- [kiks7/CVE-2020-27786-Kernel-Exploit](https://github.com/kiks7/CVE-2020-27786-Kernel-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/kiks7/CVE-2020-27786-Kernel-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/kiks7/CVE-2020-27786-Kernel-Exploit">

---
## CVE-2020-27368 (2021-01-14T16:15:00)
> Directory Indexing in Login Portal of Login Portal of TOTOLINK-A702R-V1.0.0-B20161227.1023 allows attacker to access /icons/ directories via GET Parameter.
- [swzhouu/CVE-2020-27368](https://github.com/swzhouu/CVE-2020-27368)	<img alt="forks" src="https://img.shields.io/github/forks/swzhouu/CVE-2020-27368">	<img alt="stars" src="https://img.shields.io/github/stars/swzhouu/CVE-2020-27368">

---
## CVE-2020-26733 (2021-01-14T16:15:00)
> Cross Site Scripting (XSS) in Configuration page in SKYWORTH GN542VF Hardware Version 2.0 and Software Version 2.0.0.16 allows authenticated attacker to inject their own script into the page via DDNS Configuration Section.
- [swzhouu/CVE-2020-26733](https://github.com/swzhouu/CVE-2020-26733)	<img alt="forks" src="https://img.shields.io/github/forks/swzhouu/CVE-2020-26733">	<img alt="stars" src="https://img.shields.io/github/stars/swzhouu/CVE-2020-26733">

---
## CVE-2020-26732 (2021-01-14T16:15:00)
> Skyworth GN542VF Boa version 0.94.13 does not set the Secure flag for the session cookie in an HTTPS session, which makes it easier for remote attackers to capture this cookie by intercepting its transmission within an HTTP session.
- [swzhouu/CVE-2020-26732](https://github.com/swzhouu/CVE-2020-26732)	<img alt="forks" src="https://img.shields.io/github/forks/swzhouu/CVE-2020-26732">	<img alt="stars" src="https://img.shields.io/github/stars/swzhouu/CVE-2020-26732">

---
## CVE-2020-26413 (2020-12-11T04:15:00)
> An issue has been discovered in GitLab CE/EE affecting all versions starting from 13.4 before 13.6.2. Information disclosure via GraphQL results in user email being unexpectedly visible.
- [Kento-Sec/GitLab-Graphql-CVE-2020-26413](https://github.com/Kento-Sec/GitLab-Graphql-CVE-2020-26413)	<img alt="forks" src="https://img.shields.io/github/forks/Kento-Sec/GitLab-Graphql-CVE-2020-26413">	<img alt="stars" src="https://img.shields.io/github/stars/Kento-Sec/GitLab-Graphql-CVE-2020-26413">

---
## CVE-2020-17519 (2021-01-05T12:15:00)
> A change introduced in Apache Flink 1.11.0 (and released in 1.11.1 and 1.11.2 as well) allows attackers to read any file on the local filesystem of the JobManager through the REST interface of the JobManager process. Access is restricted to files accessible by the JobManager process. All users should upgrade to Flink 1.11.3 or 1.12.0 if their Flink instance(s) are exposed. The issue was fixed in commit b561010b0ee741543c3953306037f00d7a9f0801 from apache/flink:master.
- [trhacknon/CVE-2020-17519](https://github.com/trhacknon/CVE-2020-17519)	<img alt="forks" src="https://img.shields.io/github/forks/trhacknon/CVE-2020-17519">	<img alt="stars" src="https://img.shields.io/github/stars/trhacknon/CVE-2020-17519">

---
## CVE-2020-1472 (2020-08-17T19:15:00)
> An elevation of privilege vulnerability exists when an attacker establishes a vulnerable Netlogon secure channel connection to a domain controller, using the Netlogon Remote Protocol (MS-NRPC), aka 'Netlogon Elevation of Privilege Vulnerability'.
- [lele8/CVE-2020-1472](https://github.com/lele8/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/lele8/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/lele8/CVE-2020-1472">
- [hadhub/ad-scanner](https://github.com/hadhub/ad-scanner)	<img alt="forks" src="https://img.shields.io/github/forks/hadhub/ad-scanner">	<img alt="stars" src="https://img.shields.io/github/stars/hadhub/ad-scanner">
- [sho-luv/zerologon](https://github.com/sho-luv/zerologon)	<img alt="forks" src="https://img.shields.io/github/forks/sho-luv/zerologon">	<img alt="stars" src="https://img.shields.io/github/stars/sho-luv/zerologon">
- [Nekoox/zerologon](https://github.com/Nekoox/zerologon)	<img alt="forks" src="https://img.shields.io/github/forks/Nekoox/zerologon">	<img alt="stars" src="https://img.shields.io/github/stars/Nekoox/zerologon">
- [Anonymous-Family/Zero-day-scanning](https://github.com/Anonymous-Family/Zero-day-scanning)	<img alt="forks" src="https://img.shields.io/github/forks/Anonymous-Family/Zero-day-scanning">	<img alt="stars" src="https://img.shields.io/github/stars/Anonymous-Family/Zero-day-scanning">
- [Anonymous-Family/CVE-2020-1472](https://github.com/Anonymous-Family/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/Anonymous-Family/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/Anonymous-Family/CVE-2020-1472">
- [TheJoyOfHacking/dirkjanm-CVE-2020-1472](https://github.com/TheJoyOfHacking/dirkjanm-CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/TheJoyOfHacking/dirkjanm-CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/TheJoyOfHacking/dirkjanm-CVE-2020-1472">
- [TheJoyOfHacking/SecuraBV-CVE-2020-1472](https://github.com/TheJoyOfHacking/SecuraBV-CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/TheJoyOfHacking/SecuraBV-CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/TheJoyOfHacking/SecuraBV-CVE-2020-1472">
- [Exploitspacks/CVE-2020-1472](https://github.com/Exploitspacks/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/Exploitspacks/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/Exploitspacks/CVE-2020-1472">
- [SecuraBV/CVE-2020-1472](https://github.com/SecuraBV/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/SecuraBV/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/SecuraBV/CVE-2020-1472">
- [CPO-EH/CVE-2020-1472_ZeroLogonChecker](https://github.com/CPO-EH/CVE-2020-1472_ZeroLogonChecker)	<img alt="forks" src="https://img.shields.io/github/forks/CPO-EH/CVE-2020-1472_ZeroLogonChecker">	<img alt="stars" src="https://img.shields.io/github/stars/CPO-EH/CVE-2020-1472_ZeroLogonChecker">
- [NickSanzotta/zeroscan](https://github.com/NickSanzotta/zeroscan)	<img alt="forks" src="https://img.shields.io/github/forks/NickSanzotta/zeroscan">	<img alt="stars" src="https://img.shields.io/github/stars/NickSanzotta/zeroscan">
- [R0B1NL1N/CVE-2020-1472](https://github.com/R0B1NL1N/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/R0B1NL1N/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/R0B1NL1N/CVE-2020-1472">
- [zeronetworks/zerologon](https://github.com/zeronetworks/zerologon)	<img alt="forks" src="https://img.shields.io/github/forks/zeronetworks/zerologon">	<img alt="stars" src="https://img.shields.io/github/stars/zeronetworks/zerologon">
- [Fa1c0n35/SecuraBV-CVE-2020-1472](https://github.com/Fa1c0n35/SecuraBV-CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/Fa1c0n35/SecuraBV-CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/Fa1c0n35/SecuraBV-CVE-2020-1472">
- [Fa1c0n35/CVE-2020-1472-02-](https://github.com/Fa1c0n35/CVE-2020-1472-02-)	<img alt="forks" src="https://img.shields.io/github/forks/Fa1c0n35/CVE-2020-1472-02-">	<img alt="stars" src="https://img.shields.io/github/stars/Fa1c0n35/CVE-2020-1472-02-">
- [puckiestyle/CVE-2020-1472](https://github.com/puckiestyle/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/puckiestyle/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/puckiestyle/CVE-2020-1472">
- [itssmikefm/CVE-2020-1472](https://github.com/itssmikefm/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/itssmikefm/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/itssmikefm/CVE-2020-1472">
- [NAXG/CVE-2020-1472](https://github.com/NAXG/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/NAXG/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/NAXG/CVE-2020-1472">
- [Udyz/Zerologon](https://github.com/Udyz/Zerologon)	<img alt="forks" src="https://img.shields.io/github/forks/Udyz/Zerologon">	<img alt="stars" src="https://img.shields.io/github/stars/Udyz/Zerologon">
- [hell-moon/ZeroLogon-Exploit](https://github.com/hell-moon/ZeroLogon-Exploit)	<img alt="forks" src="https://img.shields.io/github/forks/hell-moon/ZeroLogon-Exploit">	<img alt="stars" src="https://img.shields.io/github/stars/hell-moon/ZeroLogon-Exploit">
- [cedelasen/OfensivaPoc](https://github.com/cedelasen/OfensivaPoc)	<img alt="forks" src="https://img.shields.io/github/forks/cedelasen/OfensivaPoc">	<img alt="stars" src="https://img.shields.io/github/stars/cedelasen/OfensivaPoc">
- [YossiSassi/ZeroLogon-Exploitation-Check](https://github.com/YossiSassi/ZeroLogon-Exploitation-Check)	<img alt="forks" src="https://img.shields.io/github/forks/YossiSassi/ZeroLogon-Exploitation-Check">	<img alt="stars" src="https://img.shields.io/github/stars/YossiSassi/ZeroLogon-Exploitation-Check">
- [wrathfulDiety/zerologon](https://github.com/wrathfulDiety/zerologon)	<img alt="forks" src="https://img.shields.io/github/forks/wrathfulDiety/zerologon">	<img alt="stars" src="https://img.shields.io/github/stars/wrathfulDiety/zerologon">
- [SaharAttackit/CVE-2020-1472](https://github.com/SaharAttackit/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/SaharAttackit/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/SaharAttackit/CVE-2020-1472">
- [JayP232/The_big_Zero](https://github.com/JayP232/The_big_Zero)	<img alt="forks" src="https://img.shields.io/github/forks/JayP232/The_big_Zero">	<img alt="stars" src="https://img.shields.io/github/stars/JayP232/The_big_Zero">
- [Whippet0/CVE-2020-1472](https://github.com/Whippet0/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/Whippet0/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/Whippet0/CVE-2020-1472">
- [b1ack0wl/CVE-2020-1472](https://github.com/b1ack0wl/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/b1ack0wl/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/b1ack0wl/CVE-2020-1472">
- [maikelnight/zerologon](https://github.com/maikelnight/zerologon)	<img alt="forks" src="https://img.shields.io/github/forks/maikelnight/zerologon">	<img alt="stars" src="https://img.shields.io/github/stars/maikelnight/zerologon">
- [VoidSec/CVE-2020-1472](https://github.com/VoidSec/CVE-2020-1472)	<img alt="forks" src="https://img.shields.io/github/forks/VoidSec/CVE-2020-1472">	<img alt="stars" src="https://img.shields.io/github/stars/VoidSec/CVE-2020-1472">

---
## CVE-2020-1034 (2020-09-11T17:15:00)
> An elevation of privilege vulnerability exists in the way that the Windows Kernel handles objects in memory, aka 'Windows Kernel Elevation of Privilege Vulnerability'.
- [GeorgyFirsov/CVE-2020-1034](https://github.com/GeorgyFirsov/CVE-2020-1034)	<img alt="forks" src="https://img.shields.io/github/forks/GeorgyFirsov/CVE-2020-1034">	<img alt="stars" src="https://img.shields.io/github/stars/GeorgyFirsov/CVE-2020-1034">

---
## CVE-2020-0136 (2020-06-11T15:15:00)
> In multiple locations of Parcel.cpp, there is a possible out-of-bounds write due to an integer overflow. This could lead to local escalation of privilege in the system server with no additional execution privileges needed. User interaction is not needed for exploitation.Product: AndroidVersions: Android-10Android ID: A-120078455
- [Satheesh575555/libhwbinder_AOSP10_r33_CVE-2020-0136](https://github.com/Satheesh575555/libhwbinder_AOSP10_r33_CVE-2020-0136)	<img alt="forks" src="https://img.shields.io/github/forks/Satheesh575555/libhwbinder_AOSP10_r33_CVE-2020-0136">	<img alt="stars" src="https://img.shields.io/github/stars/Satheesh575555/libhwbinder_AOSP10_r33_CVE-2020-0136">
