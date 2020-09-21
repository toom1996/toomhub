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
func (logic *SquareLogic) SquareIndex(validator *validatorMiniprogramV1.SquareIndex, c *gin.Context) (interface{}, error) {

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
	k := ServiceMiniV1.SquareLikeKey + fmt.Sprintf("%d", validator.Id)
	//先验证redisKey是否存在
	r, err := util.Rdb.Exists(util.Ctx, k).Result()
	if err != nil {
		return false, err
	}

	if r != 1 {
		return false, errors.New("square not found")
	}

	fmt.Println("id -> ", util.GetIdentity().MiniId)
	rr, _ := util.Rdb.SetBit(util.Ctx, k, util.GetIdentity().MiniId, 1).Result()

	fmt.Println("rr -> ", rr)

	return true, nil
}
