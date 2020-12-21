// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/17 14:09
package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var ZConfig ZawazawaConfig

type Note struct {
	Content string
	Cities  []string
}

type ZawazawaConfig struct {
	App      app
	Server   server
	Database database
	AliSms   aliSms
	Qiniu    qiniu
}

type app struct {
	RunMode   string `ini:"RUN_MODE"`
	JwtSecret string `ini:"JWT_SECRET"`
}

type server struct {
	HttpHost string `ini:"HTTP_HOST"`
	HttpPort string `ini:"HTTP_PORT"`
}

type database struct {
	User     string `ini:"USER"`
	Password string `ini:"PASSWORD"`
	Host     string `ini:"HOST"`
	Name     string `ini:"NAME"`
	Charset  string `ini:"CHARSET"`
}

type aliSms struct {
	AccessKeyId     string `ini:"ACCESS_KEY_ID"`
	AccessKeySecret string `ini:"ACCESS_KEY_SECRET"`
}

type qiniu struct {
	AccessKey string `ini:"ACCESS_KEY"`
	SecretKey string `ini:"SECRET_KEY"`
}

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config': %v", err)
	}

	err = cfg.MapTo(&ZConfig)
	if err != nil {
		log.Fatalf("map config error: %v", err)
	}
}
