// @Description
// @Author    2020/8/26 10:51
package ServiceMiniV1

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/goinggo/mapstructure"
	"strconv"
	"time"
	//LogicMiniV1 "toomhub/logic/mini/v1"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/util"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

const SquareCacheKey = "square:id:"

// @title
func GetSquareIndex(validator *validatorMiniprogramV1.SquareIndex) (interface{}, error) {
	type imageModel struct {
		Ext   string `json:"ext"`
		Name  string `json:"name"`
		Param string `json:"param"`
		Size  int64  `json:"size"`
		Host  string `json:"host"`
	}

	db := util.DB
	rdb := util.Rdb

	pipe := rdb.Pipeline()

	var iModel imageModel
	var tempI []interface{}
	var tempL []interface{}
	var list []interface{}
	var i map[string]interface{}
	var model []ModelMiniV1.ToomhubSquare
	db.Select("id").Limit(10).Offset(0).Find(&model)

	var commands []*redis.StringStringMapCmd

	for _, v := range model {
		commands = append(commands, pipe.HGetAll(util.Ctx, SquareCacheKey+fmt.Sprintf("%d", v.Id)))
	}

	_, _ = pipe.Exec(util.Ctx)

	for _, cmd := range commands {
		result, _ := cmd.Result()
		//没结果跳过
		if len(result) == 0 {
			continue
		}
		//json 解析成数组
		err := json.Unmarshal([]byte(result["image"]), &i)
		if err != nil {
			fmt.Println(err)
		}

		for _, xx := range i {
			_ = mapstructure.Decode(xx, &iModel)
			tempL = append(tempL, iModel.Host+iModel.Name)
			tempI = append(tempI, iModel)
		}

		intCreatedAt, _ := strconv.ParseInt(result["created_at"], 10, 64)
		fmt.Println(intCreatedAt)
		createdAt := util.StrTime(intCreatedAt)
		list = append(list, map[string]interface{}{
			"content":        result["content"],
			"image":          tempI,
			"created_at":     createdAt,
			"created_by":     util.ToInt(result["created_by"]),
			"likes_count":    util.ToInt(result["likes_count"]),
			"argument_count": util.ToInt(result["argument_count"]),
			"collect_count":  util.ToInt(result["collect_count"]),
			"tag":            result["tag"],
			"id":             util.ToInt(result["id"]),
			"list":           tempL,
		})
	}

	return list, nil

	//return map[string]interface{}{
	//	"created_at":     "2020: 01 :08",
	//	"content":        "测试的那个",
	//	"created_by":     "admin",
	//	"likes_count":    "555",
	//	"argument_count": "111",
	//	"collect_count":  "88",
	//	"share_count":    "100",
	//}, nil
}

// @title	创建广场信息
func SquareCreate(v *validatorMiniprogramV1.SquareCreate, image map[string]interface{}) (bool, error) {
	identity := util.GetIdentity()

	createTime := time.Now().Unix()
	db := util.DB
	//开启事务
	transaction := db.Begin()

	//赋值结构体
	squareModel := ModelMiniV1.ToomhubSquare{
		Content:       v.Content,
		CreatedBy:     identity.MiniId,
		CreatedAt:     createTime,
		LikesCount:    0,
		ArgumentCount: 0,
		CollectCount:  0,
		ShareCount:    0,
		Tag:           v.Tag,
	}
	_ = transaction.Create(&squareModel)
	for k, value := range image {
		fmt.Println(k)
		fmt.Println(value)
		i := ModelMiniV1.ToomhubSquareImage{
			Rid: squareModel.Id,
		}
		_ = mapstructure.Decode(value, &i)
		query := transaction.Create(&i)

		if query.Error != nil {
			return false, query.Error
		}
	}

	imageJson, _ := json.Marshal(image)
	//写入缓存
	_, _ = util.Rdb.HMSet(util.Ctx, SquareCacheKey+fmt.Sprintf("%d", squareModel.Id), map[string]interface{}{
		"id":             squareModel.Id,
		"created_at":     squareModel.CreatedAt,
		"created_by":     squareModel.CreatedBy,
		"likes_count":    squareModel.LikesCount,
		"argument_count": squareModel.ArgumentCount,
		"collect_count":  squareModel.CollectCount,
		"share_count":    squareModel.ShareCount,
		"tag":            squareModel.Tag,
		"image":          imageJson,
		"content":        v.Content,
	}).Result()

	transaction.Commit()

	//标签写入ES

	fmt.Println(fmt.Sprintf(`{"name":"%s"}`, v.Tag))
	util.EsSet("toomhub", fmt.Sprintf(`{"name":"%s"}`, v.Tag))

	return true, nil
}
