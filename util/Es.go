// @Description
// @Author    2020/9/10 9:51
package util

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strings"
	//"log"
)

var es *elasticsearch.Client

//初始化
func EsInit() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.10.207:9200",
		},
	}
	es, _ = elasticsearch.NewClient(cfg)
}

type Toomhub struct {
	Name string `json:"name"`
}

func EsGet(index string, id string) {
	s, err := es.Get(
		"toomhub",
		"2",
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	fmt.Println(s)
}

func EsSet(index string, p string) {

	_, _ = es.Index(
		index,                        // Index name
		strings.NewReader(p),         // Document body
		es.Index.WithRefresh("true"), // Refresh
	)
}
