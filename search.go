package main

import (
	"encoding/json"
	"fmt"
	"github.com/sari3l/notify/notifier/bark"
	"github.com/sari3l/requests"
	nUrl "net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 修改部分

// CVE关联项目查询
const enableRelatedQuery = true

// 查询关键字
const cveQuery = "CVE-20"

// 通知函数
var barkToken = os.Getenv("barkToken")
var barkGroup = "Poc-Monitor"
var barkMsgLimit = 150

func Notice(updateItems *[]*Item) {
	for _, item := range *updateItems {
		name := nUrl.QueryEscape(item.Name)
		content := nUrl.QueryEscape(item.Description)
		if len(content) >= barkMsgLimit {
			content = content[:barkMsgLimit] + "..."
		}
		webhook := fmt.Sprintf("https://api.day.app/%s/%s/%s", barkToken, name, content)
		option := bark.Option{Webhook: webhook}
		option.Url = &item.HtmlUrl
		option.Group = &barkGroup
		err := option.ToNotifier().Send(nil)
		if err != nil {
			fmt.Printf("[!] 发送失败 %s %s\n", err, webhook)
		}
		fmt.Printf("[>] 新增 %s\n", webhook)
	}
}

// 以下勿动

const (
	blackListFile = "blacklist.yaml"
)

var UpdateJsonFilePath = fmt.Sprintf("%s/update.json", GetCurrentDirectory())
var NewJsonFilePath = fmt.Sprintf("%s/new.json", GetCurrentDirectory())
var cveExp, _ = regexp.Compile(`(?i)CVE-(\d+)-\d+`)
var blackUserMap map[int64]string

func main() {
	var addItems = make([]*Item, 0)
	var updateItems = make([]*Item, 0)
	_ = ReadYamlFile(blackListFile, &blackUserMap)
	cveList := checkLastUpdate(cveQuery, false, &addItems, &updateItems)
	if enableRelatedQuery {
		for _, cveId := range *cveList {
			_ = checkLastUpdate(cveId, true, &addItems, &updateItems)
		}
	}
	saveItems(&addItems, &updateItems)
}

func checkLastUpdate(queryStr string, isCVE bool, addItems *[]*Item, updateItems *[]*Item) *[]string {
	currentYear := time.Now().Year()
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=%s&sort=updated", queryStr)
	resp := requests.Get(url)
	if resp == nil {
		fmt.Println(fmt.Errorf("[!] 无法访问 %s", url))
	}
	items := resp.Json().Get("items").Array()
	var cveList = make([]string, 0)
	for _, data := range items {
		item := new(Item)
		err := json.Unmarshal([]byte(data.Raw), item)
		if err != nil {
			continue
		}
		// 黑名单判断
		if blackUserMap[item.Owner.Id] != "" {
			//fmt.Printf("[!] 检测到黑名单 %s %d\n", item.Owner.Login, item.Owner.Id)
			continue
		}
		// 提取CVE信息
		var cveInfo []string
		if isCVE {
			cveInfo = cveExp.FindStringSubmatch(queryStr)
		} else {
			cveInfo = cveExp.FindStringSubmatch(item.Name)
			if cveInfo == nil {
				if cveInfo = cveExp.FindStringSubmatch(item.Name); cveInfo == nil {
					continue
				}
			}
		}
		cveId, cveYear := strings.ToUpper(cveInfo[0]), cveInfo[1]
		if year, _ := strconv.Atoi(cveYear); year > currentYear {
			fmt.Println(fmt.Errorf("[!] 错误年限，%s", cveId))
			continue
		}
		cveList = append(cveList, cveId)
		// 查询本地信息
		cveYearPath := fmt.Sprintf("%s/%s", GetCurrentDirectory(), cveYear)
		cveFilePath := fmt.Sprintf("%s/%s/%s.json", GetCurrentDirectory(), cveYear, cveId)
		// 检查年限
		if !CheckFileExists(cveYearPath) {
			if err = CreateDir(cveYearPath); err != nil {
				fmt.Println(fmt.Errorf("[!] 创建 %s 失败, %s", cveYearPath, err))
				continue
			}
		}
		// 检查cve信息
		var historyItems = make([]*Item, 0)
		if CheckFileExists(cveFilePath) {
			// 读取历史cve信息
			err = ReadJsonFile(cveFilePath, &historyItems)
			if err != nil && err != EmptyError {
				fmt.Println(fmt.Errorf("[!] 读取 %s 失败, %s", cveFilePath, err))
			}
		}
		needAdd := true
		for index, historyItem := range historyItems {
			if item.Id == historyItem.Id {
				// diff
				if !reflect.DeepEqual(*item, *historyItem) {
					itemsContentValues := historyItems
					itemsContentValues[index] = item
					*updateItems = append(*updateItems, item)
					fmt.Printf("[>] 更新 %s %d\n", cveId, item.Id)
				}
				needAdd = false
				break
			}
		}
		if needAdd {
			historyItems = append(historyItems, item)
			*addItems = append(*addItems, item)
		}
		// 更新cve信息
		byteValue, err := json.Marshal(historyItems)
		if err != nil {
			fmt.Println(fmt.Errorf("[!] 转换 %s 内容失败, %s", cveId, err))
		}
		if err = WriteFile(cveFilePath, byteValue); err != nil {
			fmt.Println(fmt.Errorf("[!] 写入 %s 内容失败, %s", cveFilePath, err))
		}
	}
	return &cveList
}

func saveItems(addItems *[]*Item, updateItems *[]*Item) {
	if len(*addItems) != 0 || len(*updateItems) != 0 {
		// 更新dateLog
		logPath := fmt.Sprintf("%s/%s", GetCurrentDirectory(), LogFilePath)
		if !CheckFileExists(logPath) {
			_ = CreateDir(logPath)
		}
		dateLogFilePath := fmt.Sprintf("%s/%s.json", logPath, time.Now().Format("2006-01-02"))
		dateLogItems := DateLog{}
		if CheckFileExists(dateLogFilePath) {
			// 读取历史cve信息
			err := ReadJsonFile(dateLogFilePath, &dateLogItems)
			if err != nil && err != EmptyError {
				fmt.Println(fmt.Errorf("[!] 读取 %s 失败, %s", dateLogFilePath, err))
			}
		}
		dateLogItems.New = append(dateLogItems.New, *addItems...)
		for _, item := range *updateItems {
			for logIndex, logItem := range dateLogItems.Update {
				if item.Id == logItem.Id {
					if !reflect.DeepEqual(*item, *logItem) {
						dateLogItems.Update[logIndex] = item
					}
					break
				}
			}
			dateLogItems.Update = append(dateLogItems.Update, item)
		}
		byteValue, err := json.Marshal(dateLogItems)
		if err != nil {
			fmt.Println(fmt.Errorf("[!] 转换日志内容失败, %s", err))
		}
		if err = WriteFile(dateLogFilePath, byteValue); err != nil {
			fmt.Println(fmt.Errorf("[!] 写入日志内容失败, %s", err))
		}

		// 更新new/update内容
		if len(*updateItems) != 0 {
			byteValue, err = json.Marshal(updateItems)
			if err != nil {
				fmt.Println(fmt.Errorf("[!] 转换更新内容失败, %s", err))
			}
			if err = WriteFile(UpdateJsonFilePath, byteValue); err != nil {
				fmt.Println(fmt.Errorf("[!] 写入更新内容失败, %s", err))
			}
		}
		if len(*addItems) != 0 {
			byteValue, err = json.Marshal(addItems)
			if err != nil {
				fmt.Println(fmt.Errorf("[!] 转换新增内容失败, %s", err))
			}
			if err = WriteFile(NewJsonFilePath, byteValue); err != nil {
				fmt.Println(fmt.Errorf("[!] 写入新增内容失败, %s", err))
			}
			// 新增后通知
			Notice(addItems)
		}
	}
}
