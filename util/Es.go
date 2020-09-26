// @Description
// @Author    2020/9/10 9:51
package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
	//"log"
)

var es *elasticsearch.Client

//初始化
func EsInit() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
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

func EsSearch(param map[string]interface{}) *esapi.Response {
	var buf bytes.Buffer

	query := param

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	r, _ := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("toomhub"),
		es.Search.WithBody(&buf),
	)

	return r
}

func EsSet(index string, p string) {

	_, _ = es.Index(
		index,                        // Index name
		strings.NewReader(p),         // Document body
		es.Index.WithRefresh("true"), // Refresh
	)
}
