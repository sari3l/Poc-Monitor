package main

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var EmptyError = errors.New("-1")
var LogFilePath = "dateLog"

type Item struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login     string `json:"login"`
		Id        int64  `json:"id"`
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

type DateLog struct {
	New    []*Item `json:"new"`
	Update []*Item `json:"update"`
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

func ReadYamlFile(path string, obj any) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, obj)
	if err != nil {
		return err
	}
	return nil
}

func ReadJsonFile(filepath string, obj any) error {
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	byteValue, _ := ioutil.ReadAll(file)
	if len(byteValue) == 0 {
		return EmptyError
	}
	if err = json.Unmarshal(byteValue, obj); err != nil {
		return err
	}
	return nil
}

func ReadFile(filepath string) []byte {
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return nil
	}
	byteValue, _ := ioutil.ReadAll(file)
	return byteValue
}

func WriteFile(filepath string, content []byte) error {
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
