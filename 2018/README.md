# 2018 List

---
## CVE-2018-3990 (2019-02-05T23:29:00)
> An exploitable pool corruption vulnerability exists in the 0x8200E804 IOCTL handler functionality of WIBU-SYSTEMS WibuKey.sys Version 6.40 (Build 2400). A specially crafted IRP request can cause a buffer overflow, resulting in kernel memory corruption and, potentially, privilege escalation. An attacker can send an IRP request to trigger this vulnerability.
- [Exploitables/CVE-2018-3990-TALOS-2018-0658](https://github.com/Exploitables/CVE-2018-3990-TALOS-2018-0658)		<img src="https://user-images.githubusercontent.com/45752995/176187524-029a1a42-1c31-4d23-823a-989c40ca8460.svg" alt="fork"/>0 <img src="https://user-images.githubusercontent.com/45752995/176188923-7eb4772f-794e-48da-962b-5f8f69ca184f.svg" alt="star"/>0

---
## CVE-2018-20175 (2019-03-15T18:29:00)
> rdesktop versions up to and including v1.8.3 contains several Integer Signedness errors that lead to Out-Of-Bounds Reads in the file mcs.c and result in a Denial of Service (segfault).
- [ahaShiyu/CVE-2018-20175](https://github.com/ahaShiyu/CVE-2018-20175)		<img src="https://user-images.githubusercontent.com/45752995/176187524-029a1a42-1c31-4d23-823a-989c40ca8460.svg" alt="fork"/>0 <img src="https://user-images.githubusercontent.com/45752995/176188923-7eb4772f-794e-48da-962b-5f8f69ca184f.svg" alt="star"/>0

---
## CVE-2018-19854 (2018-12-04T16:29:00)
> An issue was discovered in the Linux kernel before 4.19.3. crypto_report_one() and related functions in crypto/crypto_user.c (the crypto user configuration API) do not fully initialize structures that are copied to userspace, potentially leaking sensitive memory to user programs. NOTE: this is a CVE-2013-2547 regression but with easier exploitability because the attacker does not need a capability (however, the system must have the CONFIG_CRYPTO_USER kconfig option).
- [ahaShiyu/CVE-2018-19854](https://github.com/ahaShiyu/CVE-2018-19854)		<img src="https://user-images.githubusercontent.com/45752995/176187524-029a1a42-1c31-4d23-823a-989c40ca8460.svg" alt="fork"/>0 <img src="https://user-images.githubusercontent.com/45752995/176188923-7eb4772f-794e-48da-962b-5f8f69ca184f.svg" alt="star"/>0

---
## CVE-2018-18839 (2019-06-18T16:15:00)
> ** DISPUTED ** An issue was discovered in Netdata 1.10.0. Full Path Disclosure (FPD) exists via api/v1/alarms. NOTE: the vendor says "is intentional."
- [ahaShiyu/CVE-2018-18839](https://github.com/ahaShiyu/CVE-2018-18839)		<img src="https://user-images.githubusercontent.com/45752995/176187524-029a1a42-1c31-4d23-823a-989c40ca8460.svg" alt="fork"/>0 <img src="https://user-images.githubusercontent.com/45752995/176188923-7eb4772f-794e-48da-962b-5f8f69ca184f.svg" alt="star"/>0

---
## CVE-2018-16509 (2018-09-05T06:29:00)
> An issue was discovered in Artifex Ghostscript before 9.24. Incorrect "restoration of privilege" checking during handling of /invalidaccess exceptions could be used by attackers able to supply crafted PostScript to execute code using the "pipe" instruction.
- [rhpco/CVE-2018-16509](https://github.com/rhpco/CVE-2018-16509)		<img src="https://user-images.githubusercontent.com/45752995/176187524-029a1a42-1c31-4d23-823a-989c40ca8460.svg" alt="fork"/>0 <img src="https://user-images.githubusercontent.com/45752995/176188923-7eb4772f-794e-48da-962b-5f8f69ca184f.svg" alt="star"/>0

---
## CVE-2018-15856 (2018-08-25T21:29:00)
> An infinite loop when reaching EOL unexpectedly in compose/parser.c (aka the keymap parser) in xkbcommon before 0.8.1 could be used by local attackers to cause a denial of service during parsing of crafted keymap files.
- [ahaShiyu/CVE-2018-15856](https://github.com/ahaShiyu/CVE-2018-15856)		<img src="https://user-images.githubusercontent.com/45752995/176187524-029a1a42-1c31-4d23-823a-989c40ca8460.svg" alt="fork"/>0 <img src="https://user-images.githubusercontent.com/45752995/176188923-7eb4772f-794e-48da-962b-5f8f69ca184f.svg" alt="star"/>0
