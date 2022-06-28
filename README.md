# Poc-Monitor

[![schedule](https://github.com/sari3l/CVE-Monitor/actions/workflows/schedule.yml/badge.svg?branch=main)](https://github.com/sari3l/CVE-Monitor/actions/workflows/schedule.yml)
![GitHub last commit](https://img.shields.io/github/last-commit/sari3l/CVE-Monitor)

## 关于

1. 状态 `failing` 为短期内没有更新
2. `Github Action`有不定时延迟，不建议此项目用于生产
3. 只有`新增`才会触发通知，具体逻辑可自行修改
4. 可从 [new.json](https://raw.githubusercontent.com/sari3l/CVE-Monitor/main/new.json) 文件获取最近一次`新增`的CVE项目信息
5. 可从 [update.json](https://raw.githubusercontent.com/sari3l/CVE-Monitor/main/update.json) 文件获取最近一次`更新`的CVE项目信息
6. 可从年限下`README.md`获取当年完整信息
7. 可从 `dateLog/`获取当天新增、更新cve内容

## 通知

1. 自行在`Secret、Environments`中添加相关通知配置信息（建议）
2. 修改`search.go`中通知函数，具体可看[sari3l/notify](https://github.com/sari3l/notify)项目