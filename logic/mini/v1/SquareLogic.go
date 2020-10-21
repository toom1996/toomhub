// @Description
// @Author    2020/8/26 10:23
package LogicMiniV1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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

	fmt.Println(validator)

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(validator.ImageList), &dat)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = service.SquareCreate(validator, dat)

	return true, nil
}

func (logic *SquareLogic) SquareLike(validator *validatorRules.LikeValidator) (bool, error) {
	likeKey := service.SquareLikeKey + fmt.Sprintf("%d", validator.Id)
	SquareKey := util.SquareCacheKey + fmt.Sprintf("%d", validator.Id)
	ctx := util.Ctx
	//先验证redisKey是否存在
	r, err := util.Rdb.Exists(util.Ctx, likeKey).Result()
	if err != nil {
		return false, err
	}

	if r != 1 {
		return false, errors.New("square not found")
	}

	fmt.Println("id -> ", util.GetIdentity().MiniId)
	has, _ := util.Rdb.HExists(ctx, likeKey, fmt.Sprintf("%d", util.GetIdentity().MiniId)).Result()
	createdBy, _ := util.Rdb.HMGet(ctx, SquareKey, "created_by").Result()
	if validator.O == 1 {
		if has == false {
			rr, _ := util.Rdb.HMSet(ctx, likeKey, map[string]interface{}{
				fmt.Sprintf("%d", util.GetIdentity().MiniId): 1,
			}).Result()
			fmt.Println("rr -> ", rr)
			//增加说说点赞量
			_, _ = util.Rdb.HIncrBy(ctx, SquareKey, "likes_count", 1).Result()

			//增加发布说说用户点赞量
			_, _ = util.Rdb.HIncrBy(ctx, "mini:user:"+createdBy[0].(string), "likes_count", 1).Result()

		}
	} else {
		fmt.Println(has)
		if has != false {
			_, _ = util.Rdb.HDel(ctx, likeKey, fmt.Sprintf("%d", util.GetIdentity().MiniId)).Result()
			_, _ = util.Rdb.HIncrBy(ctx, SquareKey, "likes_count", -1).Result()
			//扣除发布说说用户点赞量
			_, _ = util.Rdb.HIncrBy(ctx, "mini:user:"+createdBy[0].(string), "likes_count", -1).Result()
		}
	}

	//fmt.Println("rr -> ", rr)

	return true, nil
}

func (logic *SquareLogic) SquareView(validator *validatorRules.View) (interface{}, error) {

	query, err := service.GetSquareView(validator.Id)

	if err != nil {
		return nil, err
	}

	return query, nil
}
