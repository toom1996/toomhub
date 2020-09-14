// @Description
// @Author    2020/8/26 10:51
package ServiceMiniV1

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"time"
	//LogicMiniV1 "toomhub/logic/mini/v1"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/util"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

const SquareCacheKey = "square:id:"

// @title
func GetSquareIndex(validator *validatorMiniprogramV1.SquareIndex) (interface{}, error) {

	db := util.DB
	rdb := util.Rdb

	_ = rdb.Pipeline()

	var model []ModelMiniV1.ToomhubSquare
	db.Select("id").Limit(10).Offset(0).Find(&model)

	m := map[int]map[string]int64{}
	for k, v := range model {
		m[k]["test"] = v.Id
	}
	return m, nil

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
	}).Result()

	transaction.Commit()

	//标签写入ES

	fmt.Println(fmt.Sprintf(`{"name":"%s"}`, v.Tag))
	util.EsSet("toomhub", fmt.Sprintf(`{"name":"%s"}`, v.Tag))

	return true, nil
}
