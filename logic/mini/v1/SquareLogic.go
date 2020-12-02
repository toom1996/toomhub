// @Description
// @Author    2020/8/26 10:23
package LogicMiniV1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
	"toomhub/service"
	"toomhub/util"
	"toomhub/validatorRules"
)

type SquareLogic struct {
}

// @title	获取图片广场信息
func (logic *SquareLogic) SquareIndex(validator *validatorRules.SquareIndex, c *gin.Context) ([]interface{}, error) {

	query, err := service.GetSquareIndex(validator, c)

	if err != nil {
		fmt.Println("000000")
	}

	return query, nil
}

func (logic *SquareLogic) SquareCreate(validator *validatorRules.SquareCreate) (bool, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(validator.ImageList), &dat)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = service.SquareCreate(validator, dat)

	return true, nil
}

func (logic *SquareLogic) SquareVideoCreate(validator *validatorRules.SquareVideoCreate) (bool, error) {
	_, _ = service.SquareVideoCreate(validator)
	return true, nil
}

func (logic *SquareLogic) SquareLike(validator *validatorRules.LikeValidator) (bool, error) {
	likeKey := "mini:user:liked:" + fmt.Sprintf("%d", util.GetIdentity().MiniId)
	SquareKey := util.SquareCacheKey + fmt.Sprintf("%d", validator.Id)
	ctx := util.Ctx

	fmt.Println("redisKey -> ", likeKey)
	z := redis.Z{
		Member: validator.Id,
		Score:  float64(time.Now().Unix()),
	}

	has, err := util.Rdb.ZRank(ctx, likeKey, fmt.Sprintf("%d", validator.Id)).Result()

	fmt.Println("has -> ", has)
	fmt.Println("err -> ", err)

	createdBy, _ := util.Rdb.HMGet(ctx, SquareKey, "created_by").Result()
	if validator.O == 1 {
		_, _ = util.Rdb.ZAdd(ctx, likeKey, &z).Result()
		//增加说说点赞量
		_, _ = util.Rdb.HIncrBy(ctx, SquareKey, "likes_count", 1).Result()

		//增加发布说说用户点赞量
		_, _ = util.Rdb.HIncrBy(ctx, "mini:user:"+createdBy[0].(string), "likes_count", 1).Result()

	} else {

		_, _ = util.Rdb.ZRem(ctx, likeKey, validator.Id).Result()

		//扣除发布说说用户点赞量
		_, _ = util.Rdb.HIncrBy(ctx, "mini:user:"+createdBy[0].(string), "likes_count", -1).Result()

	}

	return true, nil
}

func (logic *SquareLogic) SquareView(validator *validatorRules.View, c *gin.Context) (interface{}, error) {

	query, err := service.GetSquareView(validator.Id, c)

	if err != nil {
		return nil, err
	}

	return query, nil
}
