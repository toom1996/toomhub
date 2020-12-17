// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/17 14:09
package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

var ZConfig *ZawazawaConfig

type server struct {
	HttpHost string `ini:"HTTP_HOST"`
	HttpPort string `ini:"HTTP_PORT"`
}
type Note struct {
	Content string
	Cities  []string
}

type ZawazawaConfig struct {
	Server server
	Age    int `ini:"age"`
	Male   bool
	Born   time.Time
	Note
	Created time.Time `ini:"-"`
}

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println(err)
	}
	// ...
	p := new(ZawazawaConfig)
	err = cfg.MapTo(p)
	ZConfig = p
	fmt.Println(ZConfig.Age)
}
