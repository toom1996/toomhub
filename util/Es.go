// @Description
// @Author    2020/9/10 9:51
package util

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	//"log"
)

var es *elasticsearch.Client

//初始化
func init() {
	fmt.Println(GetConfig().Es.Host)

}

type Toomhub struct {
	Name string `json:"name"`
}

func Get() {
	//cfg := elasticsearch.Config{
	//	Addresses: GetConfig().Es.Host,
	//}
	//es, _ = elasticsearch.NewClient(cfg)
	//s, err := es.Get(
	//	"toomhub",
	//	"2",
	//)
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}

	//fmt.Println(s)
}

func Set() {

}
