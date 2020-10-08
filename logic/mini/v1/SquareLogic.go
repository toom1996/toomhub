// @Description
// @Author    2020/8/26 10:23
package LogicMiniV1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	ServiceMiniV1 "toomhub/service/mini/v1"
	"toomhub/util"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

type SquareLogic struct {
}

// @title	获取图片广场信息
func (logic *SquareLogic) SquareIndex(validator *validatorMiniprogramV1.SquareIndex, c *gin.Context) ([]interface{}, error) {

	query, err := ServiceMiniV1.GetSquareIndex(validator, c)

	if err != nil {
		fmt.Println("000000")
	}

	return query, nil
}

func (logic *SquareLogic) SquareCreate(validator *validatorMiniprogramV1.SquareCreate) (bool, error) {

	fmt.Println(validator)

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(validator.ImageList), &dat)
	if err != nil {
		fmt.Println(err)
	}

	_, _ = ServiceMiniV1.SquareCreate(validator, dat)

	return true, nil
}

func (logic *SquareLogic) SquareLike(validator *validatorMiniprogramV1.LikeValidator) (bool, error) {
	likeKey := ServiceMiniV1.SquareLikeKey + fmt.Sprintf("%d", validator.Id)
	SquareKey := ServiceMiniV1.SquareCacheKey + fmt.Sprintf("%d", validator.Id)
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
