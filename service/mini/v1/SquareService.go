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

	var list []interface{}
	var model []ModelMiniV1.ToomhubSquare
	db.Select("id").Limit(10).Offset(0).Order("created_at desc").Find(&model)

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
		var i map[string]interface{}
		//json 解析成数组
		err := json.Unmarshal([]byte(result["image"]), &i)
		if err != nil {
			fmt.Println(err)
		}

		var tempI []interface{}
		var tempL []interface{}
		var iModel imageModel

		for t := 0; t < len(i); t++ {
			_ = mapstructure.Decode(i[fmt.Sprintf("%d", t)], &iModel)
			tempL = append(tempL, iModel.Host+iModel.Name+"/toomhubWatermarkStyle")
			tempI = append(tempI, iModel)
		}

		intCreatedAt, _ := strconv.ParseInt(result["created_at"], 10, 64)
		createdAt := util.StrTime(intCreatedAt)
		createdBy, _ := rdb.HMGet(util.Ctx, UserCacheKey+result["created_by"], []string{"nick_name", "avatar_url"}...).Result()

		list = append(list, map[string]interface{}{
			"content":        result["content"],
			"image":          tempI,
			"avatar_url":     createdBy[1],
			"created_at":     createdAt,
			"created_by":     createdBy[0], //这里有优化的空间
			"likes_count":    util.ToInt(result["likes_count"]),
			"argument_count": util.ToInt(result["argument_count"]),
			"collect_count":  util.ToInt(result["collect_count"]),
			"tag":            result["tag"],
			"id":             util.ToInt(result["id"]),
			"list":           tempL,
		})
	}
	return list, nil
}

// @title	创建广场信息
func SquareCreate(v *validatorMiniprogramV1.SquareCreate, image map[string]interface{}) (bool, error) {
	identity := util.GetIdentity()

	createTime := time.Now()
	db := util.DB
	//开启事务
	transaction := db.Begin()

	//赋值结构体
	squareModel := ModelMiniV1.ToomhubSquare{
		Content:       v.Content,
		CreatedBy:     identity.MiniId,
		CreatedAt:     createTime.Unix(),
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

	if v.Tag != "" {
		//标签写入ES
		fmt.Println(fmt.Sprintf(`{"name":"%s"}`, v.Tag))

		util.EsSet("toomhub", fmt.Sprintf(`{"name":"%s"}`, v.Tag))

		//入redisZset排序
		util.Rdb.ZIncrBy(util.Ctx, "hotTag:"+createTime.Format("20060102"), 1, v.Tag)
	}

	return true, nil
}
