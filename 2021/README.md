# 2021 List

---
## CVE-2021-44228 (2021-12-10T10:15:00)
> Apache Log4j2 2.0-beta9 through 2.15.0 (excluding security releases 2.12.2, 2.12.3, and 2.3.1) JNDI features used in configuration, log messages, and parameters do not protect against attacker controlled LDAP and other JNDI related endpoints. An attacker who can control log messages or log message parameters can execute arbitrary code loaded from LDAP servers when message lookup substitution is enabled. From log4j 2.15.0, this behavior has been disabled by default. From version 2.16.0 (along with 2.12.2, 2.12.3, and 2.3.1), this functionality has been completely removed. Note that this vulnerability is specific to log4j-core and does not affect log4net, log4cxx, or other Apache Logging Services projects.
- [aws-samples/kubernetes-log4j-cve-2021-44228-node-agent](https://github.com/aws-samples/kubernetes-log4j-cve-2021-44228-node-agent)

---
## CVE-2021-40650 (2022-06-14T10:15:00)
> In Connx Version 6.2.0.1269 (20210623), a cookie can be issued by the application and not have the secure flag set.
- [l00neyhacker/CVE-2021-40650](https://github.com/l00neyhacker/CVE-2021-40650)

---
## CVE-2021-40649 (2022-06-14T10:15:00)
> In Connx Version 6.2.0.1269 (20210623), a cookie can be issued by the application and not have the HttpOnly flag set.
- [l00neyhacker/CVE-2021-40649](https://github.com/l00neyhacker/CVE-2021-40649)
