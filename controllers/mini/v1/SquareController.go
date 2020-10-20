// @Description
// @Author    2020/8/26 9:18
package ControllersMiniV1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"strings"
	"time"
	LogicMiniV1 "toomhub/logic/mini/v1"
	v1MiniMiddleware "toomhub/middware/mini/v1"
	"toomhub/util"
	"toomhub/validatorRules"
)

type SquareController struct {
}

//当前控制器注册的路由
func (square *SquareController) Register(engine *gin.Engine) {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("countValidate", validatorRules.CountValidate)
	}

	user := engine.Group("/v1/mini/sq")
	// 广场首页接口
	user.GET("/index", square.index)
	user.GET("/view", square.view)
	user.Use(v1MiniMiddleware.CheckIdentity())
	{
		// 发布一条广场信息
		user.POST("/create", square.create)
		//标签搜索
		user.GET("/tag-search", square.TagSearch)
		//广场消息点赞
		user.POST("/like", square.Like)
	}
}

// @title	广场首页
func (square *SquareController) index(Context *gin.Context) {
	//验证器
	formValidator := validatorRules.SquareIndex{}
	err := Context.BindQuery(&formValidator)
	if err != nil {
		Context.String(http.StatusBadRequest, "参数错误:%s", err.Error())
		return
	}

	//logic
	logic := LogicMiniV1.SquareLogic{}
	query, _ := logic.SquareIndex(&formValidator, Context)

	count := 0
	if query != nil {
		count = len(query)
	}

	Context.JSON(200, gin.H{
		"message": "OK",
		"code":    200,
		"data": map[string]interface{}{
			"list":  query,
			"count": count,
		},
	})
}

// @title	创建一条广场消息
func (square *SquareController) create(Context *gin.Context) {
	//var commonValidator validator2.CommonValidator
	//验证器
	formValidator := validatorRules.SquareCreate{}
	err := Context.ShouldBind(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code": 400,
			//"message": commonvalidatorRules.TransError(err.(validatorRules.ValidationErrors)),
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	//logic
	logic := LogicMiniV1.SquareLogic{}
	_, err = logic.SquareCreate(&formValidator)

	if err != nil {
		Context.JSON(200, gin.H{
			"message": err,
			"code":    400,
		})
	}
	Context.JSON(200, gin.H{
		"message": "OK",
		"code":    200,
	})
}

// @title	标签搜索
func (square *SquareController) TagSearch(Context *gin.Context) {

	//空搜索返回热门标签
	keyword := Context.Query("k")
	res := make(map[int]interface{})
	if keyword == "" {

		t := time.Now()

		t1 := t
		t2 := t.Add(-1 * 24 * time.Hour)
		t3 := t.Add(-2 * 24 * time.Hour)
		t4 := t.Add(-3 * 24 * time.Hour)
		t5 := t.Add(-4 * 24 * time.Hour)
		t6 := t.Add(-5 * 24 * time.Hour)
		t7 := t.Add(-6 * 24 * time.Hour)

		fmt.Println("hotTag:" + t1.Format("20060102"))
		s := &redis.ZStore{
			Keys: []string{
				"hotTag:" + t1.Format("20060102"),
				"hotTag:" + t2.Format("20060102"),
				"hotTag:" + t3.Format("20060102"),
				"hotTag:" + t4.Format("20060102"),
				"hotTag:" + t5.Format("20060102"),
				"hotTag:" + t6.Format("20060102"),
				"hotTag:" + t7.Format("20060102"),
			},
		}

		//合并zset
		_, _ = util.Rdb.ZUnionStore(util.Ctx, "hotTag", s).Result()

		r, _ := util.Rdb.ZRangeWithScores(util.Ctx, "hotTag", -10, -1).Result()

		length := len(r)
		for index, item := range r {
			res[length-index-1] = item
		}

		Context.JSON(200, gin.H{
			"code": 200,
			"data": res,
		})
		return
	}

	query := util.EsSearch("toomhub_tag", fmt.Sprintf(`{
  "query": {
    "function_score": { 
      "query": { 
        "match": {
			"tag": "%s"
		}
      },
      "field_value_factor": { 
        "field": "hot" ,
        "modifier":"log1p"
      }
    }
  }
}`, keyword))

	var r map[string]interface{}

	if err := json.NewDecoder(query.Body).Decode(&r); err != nil {
		fmt.Println(fmt.Println("Error parsing the response body: %s", err))
	}
	fmt.Println(r)

	if _, ok := r["hits"]; ok {
		// Print the ID and document source for each hit.
		for index, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			res[index] = map[string]interface{}{
				"Score":  hit.(map[string]interface{})["_source"].(map[string]interface{})["hot"],
				"Member": hit.(map[string]interface{})["_source"].(map[string]interface{})["tag"],
			}
			fmt.Println(fmt.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"]))
		}

		log.Println(strings.Repeat("=", 37))

		Context.JSON(200, gin.H{
			"code": 200,
			"data": res,
		})
		return
	}

	Context.JSON(200, gin.H{
		"code": 200,
		"data": "",
	})
}

// @title	点赞
func (square *SquareController) Like(Context *gin.Context) {

	formValidator := validatorRules.LikeValidator{}
	err := Context.ShouldBind(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	logic := LogicMiniV1.SquareLogic{}
	_, err = logic.SquareLike(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "fail",
		})
		return
	}

	Context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "OK",
	})

	return

}

//详情
func (square *SquareController) view(Context *gin.Context) {

	formValidator := validatorRules.View{}
	err := Context.ShouldBind(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	logic := LogicMiniV1.SquareLogic{}
	_, err = logic.SquareView(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "fail",
		})
		return
	}

	Context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "OK",
	})

	return

}
