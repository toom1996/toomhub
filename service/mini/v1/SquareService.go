// @Description
// @Author    2020/8/26 10:51
package ServiceMiniV1

import (
	"fmt"
	"time"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/util"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

// @title
func GetSquareIndex(validator *validatorMiniprogramV1.SquareIndex) (interface{}, error) {
	fmt.Println(validator)
	return map[string]interface{}{
		"created_at":     "2020: 01 :08",
		"content":        "测试的那个",
		"created_by":     "admin",
		"likes_count":    "555",
		"argument_count": "111",
		"collect_count":  "88",
		"share_count":    "100",
	}, nil
}

// @title	创建广场信息
func SquareCreate(validator *validatorMiniprogramV1.SquareCreate) (bool, error) {
	createTime := time.Now().Unix()
	db := util.DB
	//开启事务
	transaction := db.Begin()

	//赋值结构体
	squareModel := ModelMiniV1.Square{
		Content:       validator.Content,
		CreatedBy:     10,
		CreatedAt:     createTime,
		LikesCount:    0,
		ArgumentCount: 0,
		CollectCount:  0,
		ShareCount:    0,
	}
	//insert到表
	query := transaction.Create(&squareModel).Scan(&squareModel)
	if query.Error != nil {
		transaction.Rollback()
		return false, query.Error
	}

	return true, nil
}
