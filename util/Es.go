// @Description
// @Author    2020/9/10 9:51
package util

import (
	"context"
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
			GetConfig().Es.Host,
		},
	}
	es, _ = elasticsearch.NewClient(cfg)
}

type Toomhub struct {
	Name string `json:"name"`
}

func EsSearch(index string, param string) *esapi.Response {

	r, _ := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(strings.NewReader(param)),
	)

	return r
}

func EsSet(index string, p string, id string) bool {

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(p),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println(res)
	//_, err := es.Index(
	//	index,                        // Index name
	//	strings.NewReader(p),         // Document body
	//	es.Index.WithRefresh("true"), // Refresh
	//)
	return true
}

func SetTag(index string, p string, id string) {

	xx, err := es.Update(
		index,
		id,
		strings.NewReader(p),
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("xx -> ", xx)
}
