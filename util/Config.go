package util

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string   `json:"app_name"`
	AppMode  string   `json:"app_mode"`
	AppHost  string   `json:"app_host"`
	AppPort  string   `json:"app_port"`
	Mini     Mini     `json:"mini"`
	Jwt      Jwt      `json:"jwt"`
	Redis    Redis    `json:"redis"`
	Database Database `json:"database"`
	Qiniu    Qiniu    `json:"qiniu"`
	Es       Es       `json:"es"`
}

type Es struct {
	Host []string `json:"host"`
}

type Qiniu struct {
	Bucket       string `json:"bucket"`       //七牛云bucket
	AccessKey    string `json:"accessKey"`    //七牛云AK
	SecretKey    string `json:"secretKey"`    //七牛云AS
	FileServer   string `json:"fileServer"`   //cdn地址
	ImageMaxSize int8   `json:"imageMaxSize"` //上传最大图片尺寸 M为单位
}

type Mini struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type Jwt struct {
	Secret string `json:"secret"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database int8   `json:"database"`
	Password string `json:"password"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Charset  string `json:"charset"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func Init(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
