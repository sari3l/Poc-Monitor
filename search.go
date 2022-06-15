package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sari3l/notify/notifier/bark"
	"github.com/sari3l/requests"
	"io/ioutil"
	nUrl "net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"time"
)

// 修改部分

// 查询关键字
const cveQuery = "CVE-20"

// 通知函数
var barkToken = os.Getenv("barkToken")

func Notice(updateItems *[]*Item) {
	for _, item := range *updateItems {
		webhook := fmt.Sprintf("https://api.day.app/%s/%s/%s", barkToken, nUrl.QueryEscape(item.Name), nUrl.QueryEscape(item.Description))
		option := bark.Option{Webhook: webhook}
		option.Url = item.HtmlUrl
		option.ToNotifier().Send(nil)
		fmt.Printf("[>] 新增 %s\n", webhook)
	}
}

// 以下勿动

var EmptyError = errors.New("-1")
var UpdateJsonFilePath = fmt.Sprintf("%s/update.json", GetCurrentDirectory())
var NewJsonFilePath = fmt.Sprintf("%s/new.json", GetCurrentDirectory())

type Item struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login     string `json:"login"`
		Id        int    `json:"id"`
		NodeId    string `json:"node_id"`
		AvatarUrl string `json:"avatar_url"`
		Url       string `json:"url"`
		HtmlUrl   string `json:"html_url"`
		Type      string `json:"type"`
		SiteAdmin bool   `json:"site_admin"`
	} `json:"owner"`
	HtmlUrl         string    `json:"html_url"`
	Description     string    `json:"description"`
	Fork            bool      `json:"fork"`
	Url             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	ForksCount      int       `json:"forks_count"`
	Archived        bool      `json:"archived"`
	Disabled        bool      `json:"disabled"`
	AllowForking    bool      `json:"allow_forking"`
	IsTemplate      bool      `json:"is_template"`
	Topics          []string  `json:"topics"`
	Visibility      string    `json:"visibility"`
	Forks           int       `json:"forks"`
	Watchers        int       `json:"watchers"`
	DefaultBranch   string    `json:"default_branch"`
	Score           float64   `json:"score"`
}

var cveExp, _ = regexp.Compile(`CVE-(\d+)-\d+`)

func main() {
	url := fmt.Sprintf("https://api.github.com/search/repositories?q=%s&sort=updated", cveQuery)
	resp := requests.Get(url)
	if resp == nil {
		fmt.Println(fmt.Errorf("[!] 无法访问 %s", url))
	}
	items := resp.Json().Get("items").Array()
	var addItems = make([]*Item, 0)
	var updateItems = make([]*Item, 0)
	for _, data := range items {
		item := new(Item)
		err := json.Unmarshal([]byte(data.Raw), item)
		if err != nil {
			continue
		}
		// 提取CVE信息
		cveInfo := cveExp.FindStringSubmatch(item.Name)
		// 查询本地信息
		cveId, cveYear := cveInfo[0], cveInfo[1]
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
			err, historyItems = ReadJsonFile(cveFilePath)
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
					updateItems = append(updateItems, item)
					fmt.Printf("[>] 更新 %s %d\n", cveId, item.Id)
				}
				needAdd = false
				break
			}
		}
		if needAdd {
			historyItems = append(historyItems, item)
			addItems = append(addItems, item)
		}
		// 更新cve信息
		byteValue, err := json.Marshal(historyItems)
		if err != nil {
			fmt.Println(fmt.Errorf("[!] 转换 %s 内容失败, %s", cveId, err))
		}
		if err = WriteJsonFile(cveFilePath, byteValue); err != nil {
			fmt.Println(fmt.Errorf("[!] 写入 %s 内容失败, %s", cveFilePath, err))
		}
	}
	// 退出前更新new/update内容
	if len(updateItems) != 0 {
		byteValue, err := json.Marshal(updateItems)
		if err != nil {
			fmt.Println(fmt.Errorf("[!] 转换更新内容失败, %s", err))
		}
		if err = WriteJsonFile(UpdateJsonFilePath, byteValue); err != nil {
			fmt.Println(fmt.Errorf("[!] 写入更新内容失败, %s", err))
		}
	}
	if len(addItems) != 0 {
		byteValue, err := json.Marshal(addItems)
		if err != nil {
			fmt.Println(fmt.Errorf("[!] 转换新增内容失败, %s", err))
		}
		if err = WriteJsonFile(NewJsonFilePath, byteValue); err != nil {
			fmt.Println(fmt.Errorf("[!] 写入新增内容失败, %s", err))
		}
		// 新增后通知
		Notice(&addItems)
	}
}

func CheckFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return true
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func CreateDir(dirPath string) error {
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func ReadJsonFile(filepath string) (error, []*Item) {
	items := make([]*Item, 0)
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return err, items
	}
	byteValue, _ := ioutil.ReadAll(file)
	if len(byteValue) == 0 {
		return EmptyError, items
	}
	if err = json.Unmarshal(byteValue, &items); err != nil {
		return err, items
	}
	return nil, items
}

func WriteJsonFile(filepath string, content []byte) error {
	if !CheckFileExists(filepath) {
		_ = CreateFile(filepath)
	}
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentDirectory() string {
	currentDir, _ := os.Executable()
	exPath := filepath.Dir(currentDir)
	return exPath
}
