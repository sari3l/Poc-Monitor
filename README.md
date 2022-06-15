# CVE-Monitor

<h1 align="center"><img src="https://raw.githubusercontent.com/sari3l/CVE-Monitor/main/static/logo.png" alt="Logo"/></h1>

[![schedule](https://github.com/sari3l/CVE-Monitor/actions/workflows/schedule.yml/badge.svg?branch=main)](https://github.com/sari3l/CVE-Monitor/actions/workflows/schedule.yml)
![GitHub last commit](https://img.shields.io/github/last-commit/sari3l/CVE-Monitor)

## 关于

1. 状态 `failing` 为短期内没有更新
2. 可从 [update.json](https://raw.githubusercontent.com/sari3l/CVE-Monitor/main/update.json) 文件获取最近一次更新的CVE项目信息

## Fork

1. 自行在`Secret、Environments`中添加相关通知配置信息（建议）
2. 修改`search.go`中通知函数，具体可看[sari3l/notify](https://github.com/sari3l/notify)项目