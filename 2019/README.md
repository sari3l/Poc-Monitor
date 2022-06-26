# 2019 List

---
## CVE-2019-6447 (2019-01-16T14:29:00)
> The ES File Explorer File Manager application through 4.1.9.7.4 for Android allows remote attackers to read arbitrary files or execute applications via TCP port 59777 requests on the local Wi-Fi network. This TCP port remains open after the ES application has been launched once, and responds to unauthenticated application/json data over HTTP.
- [VinuKalana/CVE-2019-6447-Android-Vulnerability-in-ES-File-Explorer](https://github.com/VinuKalana/CVE-2019-6447-Android-Vulnerability-in-ES-File-Explorer)

---
## CVE-2019-3560 (2019-04-29T16:29:00)
> An improperly performed length calculation on a buffer in PlaintextRecordLayer could lead to an infinite loop and denial-of-service based on user input. This issue affected versions of fizz prior to v2019.03.04.00.
- [ahaShiyu/CVE-2019-3560](https://github.com/ahaShiyu/CVE-2019-3560)

---
## CVE-2019-19945 (2020-03-16T18:15:00)
> uhttpd in OpenWrt through 18.06.5 and 19.x through 19.07.0-rc2 has an integer signedness error. This leads to out-of-bounds access to a heap buffer and a subsequent crash. It can be triggered with an HTTP POST request to a CGI script, specifying both "Transfer-Encoding: chunked" and a large negative Content-Length value.
- [delicateByte/CVE-2019-19945_Test](https://github.com/delicateByte/CVE-2019-19945_Test)

---
## CVE-2019-17662 (2019-10-16T18:15:00)
> ThinVNC 1.0b1 is vulnerable to arbitrary file read, which leads to a compromise of the VNC server. The vulnerability exists even when authentication is turned on during the deployment of the VNC server. The password for authentication is stored in cleartext in a file that can be read via a ../../ThinVnc.ini directory traversal attack vector.
- [Tamagaft/CVE-2019-17662](https://github.com/Tamagaft/CVE-2019-17662)

---
## CVE-2019-16278 (2019-10-14T17:15:00)
> Directory Traversal in the function http_verify in nostromo nhttpd through 1.9.6 allows an attacker to achieve remote code execution via a crafted HTTP request.
- [keshiba/cve-2019-16278](https://github.com/keshiba/cve-2019-16278)

---
## CVE-2019-15107 (2019-08-16T03:15:00)
> An issue was discovered in Webmin <=1.920. The parameter old in password_change.cgi contains a command injection vulnerability.
- [psw01/CVE-2019-15107_webminRCE](https://github.com/psw01/CVE-2019-15107_webminRCE)

---
## CVE-2019-12874 (2019-06-18T18:15:00)
> An issue was discovered in zlib_decompress_extra in modules/demux/mkv/util.cpp in VideoLAN VLC media player 3.x through 3.0.7. The Matroska demuxer, while parsing a malformed MKV file type, has a double free.
- [ahaShiyu/CVE-2019-12874](https://github.com/ahaShiyu/CVE-2019-12874)

---
## CVE-2019-1205 (2019-08-14T21:15:00)
> A remote code execution vulnerability exists in Microsoft Word software when it fails to properly handle objects in memory, aka 'Microsoft Word Remote Code Execution Vulnerability'. This CVE ID is unique from CVE-2019-1201.
- [razordeveloper/CVE-2019-1205](https://github.com/razordeveloper/CVE-2019-1205)

---
## CVE-2019-10742 (2019-05-07T19:29:00)
> Axios up to and including 0.18.0 allows attackers to cause a denial of service (application crash) by continuing to accepting content after maxContentLength is exceeded.
- [Viniciuspxf/CVE-2019-10742](https://github.com/Viniciuspxf/CVE-2019-10742)

---
## CVE-2019-1010319 (2019-07-11T20:15:00)
> WavPack 5.1.0 and earlier is affected by: CWE-457: Use of Uninitialized Variable. The impact is: Unexpected control flow, crashes, and segfaults. The component is: ParseWave64HeaderConfig (wave64.c:211). The attack vector is: Maliciously crafted .wav file. The fixed version is: After commit https://github.com/dbry/WavPack/commit/33a0025d1d63ccd05d9dbaa6923d52b1446a62fe.
- [ahaShiyu/CVE-2019-1010319](https://github.com/ahaShiyu/CVE-2019-1010319)
