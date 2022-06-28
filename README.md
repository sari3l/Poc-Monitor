<h1 align="center">Poc-Monitor</h1>

<p align="center">
    <img src="https://github.com/sari3l/CVE-Monitor/actions/workflows/schedule.yml/badge.svg?branch=main" alt="schedule" />
    <img src="https://img.shields.io/github/last-commit/sari3l/CVE-Monitor" alt="GitHub last commit" />
    <a href="https://github.com/sari3l/Poc-Monitor/blob/main/LICENSE"><img src="https://img.shields.io/github/license/sari3l/Poc-Monitor" alt="GitHub license" /></a>
</p>
<p align="center">
    <img src="https://img.shields.io/github/commit-activity/m/sari3l/Poc-Monitor" alt="GitHub commit activity" />
    <img
        src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fsari3l%2FPoc-Monitor&amp;count_bg=%2379C83D&amp;title_bg=%23555555&amp;icon=&amp;icon_color=%23E7E7E7&amp;title=visitors&amp;edge_flat=false"
        alt="Hits"
    />
</p>

## 关于

1. 状态 `failing` 为短期内没有更新
2. 可从 [new.json](https://raw.githubusercontent.com/sari3l/CVE-Monitor/main/new.json) 文件获取最近一次`新增`的CVE项目信息
3. 可从 [update.json](https://raw.githubusercontent.com/sari3l/CVE-Monitor/main/update.json) 文件获取最近一次`更新`的CVE项目信息
4. 可从年限目录内`README.md`获取当年完整信息
5. 可从`dateLog`目录获取当天新增、更新cve内容

## 通知

1. 只有`新增`才会触发通知，具体逻辑可自行修改 
2. 修改`search.go`中通知函数更换通知渠道，具体可看[sari3l/notify](https://github.com/sari3l/notify)项目