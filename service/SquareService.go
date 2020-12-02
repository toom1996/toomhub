// @Description
// @Author    2020/8/26 10:51
package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	"toomhub/validatorRules"

	//LogicMiniV1 "toomhub/logic/mini/v1"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/util"
)

const SquareLikeKey = "square:like:"

// @title
func GetSquareIndex(validator *validatorRules.SquareIndex, c *gin.Context) ([]interface{}, error) {
	type imageModel struct {
		Ext  string `json:"ext"`
		Name string `json:"name"`
		Size int64  `json:"size"`
		Host string `json:"host"`
	}
	db := util.DB

	var list []interface{}
	var model []ModelMiniV1.ToomhubSquare
	db.Select("id").Limit(20).Offset(20 * (validator.Page - 1)).Order("created_at desc").Find(&model)
	var squareIds []string
	for _, v := range model {
		squareIds = append(squareIds, util.SquareCacheKey+strconv.Itoa(int(v.Id)))
	}
	squareInfo, _ := util.RedisMulti([]string{"id", "type", "created_by", "created_at", "likes_count", "tag", "content", "image", "video", "cover", "width", "height"}, squareIds...)

	var creatorIds []string
	for _, item := range squareInfo {
		fmt.Println(item.([]interface{})[2].(string))
		creatorIds = append(creatorIds, util.UserCacheKey+item.([]interface{})[2].(string))
	}

	userInfo, _ := util.RedisMulti([]string{"nick_name", "avatar_url", "mini_id", "exp"}, creatorIds...)

	for index, item := range squareInfo {
		response := map[string]interface{}{}
		//square 类型
		squareType := item.([]interface{})[1]
		if squareType == nil {
			//video
			if item.([]interface{})[8] != nil {
				squareType = util.SquareTypeVideo
			} else {
				squareType = util.SquareTypeImage
			}
		}

		if squareType == 0 {
			var i map[string]interface{}
			var tempI []interface{}
			var tempL []interface{}
			//json 解析成数组
			err := json.Unmarshal([]byte(item.([]interface{})[7].(string)), &i)
			if err != nil {
				fmt.Println(err)
			}

			var iModel imageModel
			for t := 0; t < len(i); t++ {
				_ = mapstructure.Decode(i[fmt.Sprintf("%d", t)], &iModel)
				tempL = append(tempL, iModel.Host+iModel.Name)
				tempI = append(tempI, iModel)
				response["image"] = tempI
				response["list"] = tempL
				response["type"] = squareType
			}
		} else {
			response["video"] = item.([]interface{})[8]
			response["cover"] = item.([]interface{})[9]
			if item.([]interface{})[11] != nil {
				response["height"], _ = strconv.Atoi(item.([]interface{})[11].(string))
			} else {
				response["height"] = 0
			}
			if item.([]interface{})[10] != nil {
				response["width"], _ = strconv.Atoi(item.([]interface{})[10].(string))
			} else {
				response["width"] = 0
			}
			response["type"] = squareType
		}
		response["param"] = "imageMogr2/auto-orient/format/webp"
		response["content"] = item.([]interface{})[6]
		response["tag"] = item.([]interface{})[5]
		response["like_count"], _ = strconv.Atoi(item.([]interface{})[4].(string))
		timeStamp, _ := strconv.Atoi(item.([]interface{})[3].(string))
		response["created_at"] = util.StrTime(int64(timeStamp))
		response["avatar_url"] = userInfo[index].([]interface{})[1]
		response["created_by"] = userInfo[index].([]interface{})[0]
		exp := 0
		if userInfo[index].([]interface{})[3] != nil {
			exp, _ = strconv.Atoi(userInfo[index].([]interface{})[3].(string))
		}
		response["level_tag"] = util.GetLevelTag(exp)
		response["host"] = "http://toomhub.image.23cm.cn/"
		response["id"] = item.([]interface{})[0].(string)
		identity := util.Identity(c)
		isLike := 0
		_, err := util.Rdb.ZRank(util.Ctx, util.SquareLikeKey+fmt.Sprintf("%d", identity.MiniId), item.([]interface{})[0].(string)).Result()

		if err == nil {
			isLike = 1
		}

		response["is_like"] = isLike
		list = append(list, response)
	}

	//var commands []*redis.StringStringMapCmd
	//
	//for _, v := range model {
	//	commands = append(commands, pipe.HGetAll(util.Ctx, util.SquareCacheKey+fmt.Sprintf("%d", v.Id)))
	//}
	//
	//_, _ = pipe.Exec(util.Ctx)
	//
	//for _, cmd := range commands {
	//	result, _ := cmd.Result()
	//	//没结果跳过
	//	if len(result) == 0 {
	//		continue
	//	}
	//
	//	mapString := map[string]interface{}{}
	//
	//	intCreatedAt, _ := strconv.ParseInt(result["created_at"], 10, 64)
	//	createdAt := util.StrTime(intCreatedAt)
	//	createdBy, _ := rdb.HMGet(util.Ctx, util.UserCacheKey+result["created_by"], []string{"nick_name", "avatar_url", "exp"}...).Result()
	//
	//	uid := 0
	//	//判断浏览首页的用户是否登录
	//	token := c.GetHeader("Toomhub-Token")
	//	if token != "" {
	//		r, _ := util.ParseToken(c.GetHeader("Toomhub-Token"), c)
	//		uid = int(r.MiniId)
	//	}
	//
	//	isLikeRes, err := util.Rdb.HMGet(util.Ctx, SquareLikeKey+result["id"], fmt.Sprintf("%d", uid)).Result()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	isLike := 0
	//	if len(isLikeRes) == 1 && isLikeRes[0] == "1" {
	//		isLike = 1
	//	}
	//	mapString["content"] = result["content"]
	//	mapString["avatar_url"] = createdBy[1]
	//	mapString["created_at"] = createdAt
	//	mapString["created_by"] = createdBy[0]
	//	mapString["like_count"] = util.ToInt(result["likes_count"])
	//	mapString["argument_count"] = util.ToInt(result["argument_count"])
	//	mapString["collect_count"] = util.ToInt(result["collect_count"])
	//	mapString["tag"] = result["tag"]
	//	mapString["id"] = util.ToInt(result["id"])
	//	mapString["is_like"] = isLike
	//	mapString["param"] = "imageMogr2/auto-orient/format/webp"
	//	if createdBy[2] == nil {
	//		createdBy[2] = 0
	//	} else {
	//		createdBy[2], _ = strconv.Atoi(createdBy[2].(string))
	//	}
	//	mapString["level_tag"] = util.GetLevelTag(createdBy[2].(int))
	//	list = append(list, mapString)
	//}
	return list, nil
}

func calculateVideoSize(width int32, height int32, screenWidth int32, screenHeight int32) {
}

// @title	创建广场信息
func SquareCreate(v *validatorRules.SquareCreate, image map[string]interface{}) (bool, error) {
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
	_, _ = util.Rdb.HMSet(util.Ctx, util.SquareCacheKey+fmt.Sprintf("%d", squareModel.Id), map[string]interface{}{
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

	//设置一个空的hashmap 用来做点赞
	util.Rdb.HMSet(util.Ctx, SquareLikeKey+fmt.Sprintf("%d", squareModel.Id), map[string]interface{}{
		"0": 0,
	})

	transaction.Commit()

	if v.Tag != "" {
		squareModel := &ModelMiniV1.ToomhubSquareTag{}
		dbQuery := db.Table("toomhub_square_tag").Where("tag = ?", v.Tag).Take(&squareModel)
		if gorm.IsRecordNotFoundError(dbQuery.Error) {
			t := db.Begin()
			model := ModelMiniV1.ToomhubSquareTag{
				Tag:   v.Tag,
				Count: 0,
			}

			db.Create(&model)
			//获取插入记录的Id
			var id []int
			db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)

			//因为Pluck函数返回的是一列值，返回结果是slice类型，我们这里只有一个值，所以取第一个值即可。
			squareModel.Id = int64(id[0])
			t.Commit()
		}

		//fmt.Println(fmt.Sprintf(`{"tag":"%s", "content":"%s"}`, v.Tag, v.Content))
		//util.EsSet("toomhub_tag", fmt.Sprintf(`{"tag":"%s", "content":"%s", "created_at":%d, "like":0}`, v.Tag, v.Content, time.Now().Unix()), fmt.Sprintf("%d", squareModel.Id))

		//入redisZset排序
		_, _ = util.Rdb.ZIncrBy(util.Ctx, "hotTag:"+createTime.Format("20060102"), 1, v.Tag).Result()
		fmt.Println("queryId -> ", squareModel.Id)
		fmt.Println(fmt.Sprintf(`{"script" : {"source": "ctx._source.hot += 1"},"upsert" : {"hot" : 1,"tag": "%s"}}`, v.Tag))
		util.SetTag("toomhub_tag", fmt.Sprintf(`{"script" : {"source": "ctx._source.hot += 1"},"upsert" : {"hot" : 1,"tag": "%s", "created_at":%d}}`, v.Tag, time.Now().Unix()), fmt.Sprintf("%d", squareModel.Id))
	}

	return true, nil
}

func SquareVideoCreate(v *validatorRules.SquareVideoCreate) (bool, error) {
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
		Type:          util.SquareTypeVideo,
	}
	_ = transaction.Create(&squareModel)

	model := ModelMiniV1.ToomhubSquareVideo{
		Rid:      squareModel.Id,
		Host:     v.Host,
		Size:     v.Size,
		Name:     v.Name,
		Duration: v.Duration,
	}

	query := transaction.Create(&model)

	if query.Error != nil {
		return false, query.Error
	}

	//写入缓存
	_, _ = util.Rdb.HMSet(util.Ctx, util.SquareCacheKey+fmt.Sprintf("%d", squareModel.Id), map[string]interface{}{
		"id":             squareModel.Id,
		"created_at":     squareModel.CreatedAt,
		"created_by":     squareModel.CreatedBy,
		"likes_count":    squareModel.LikesCount,
		"argument_count": squareModel.ArgumentCount,
		"collect_count":  squareModel.CollectCount,
		"share_count":    squareModel.ShareCount,
		"tag":            squareModel.Tag,
		"content":        v.Content,
		"video":          v.Name,
		"cover":          v.Cover,
		"height":         v.Height,
		"width":          v.Width,
		"host":           v.Host,
	}).Result()

	transaction.Commit()

	//设置一个空的hashmap 用来做点赞
	util.Rdb.HMSet(util.Ctx, SquareLikeKey+fmt.Sprintf("%d", squareModel.Id), map[string]interface{}{
		"0": 0,
	})

	if v.Tag != "" {
		squareModel := &ModelMiniV1.ToomhubSquareTag{}
		dbQuery := db.Table("toomhub_square_tag").Where("tag = ?", v.Tag).Take(&squareModel)
		if gorm.IsRecordNotFoundError(dbQuery.Error) {
			t := db.Begin()
			model := ModelMiniV1.ToomhubSquareTag{
				Tag:   v.Tag,
				Count: 0,
			}

			db.Create(&model)
			//获取插入记录的Id
			var id []int
			db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)

			//因为Pluck函数返回的是一列值，返回结果是slice类型，我们这里只有一个值，所以取第一个值即可。
			squareModel.Id = int64(id[0])
			t.Commit()
		}

		//fmt.Println(fmt.Sprintf(`{"tag":"%s", "content":"%s"}`, v.Tag, v.Content))
		//util.EsSet("toomhub_tag", fmt.Sprintf(`{"tag":"%s", "content":"%s", "created_at":%d, "like":0}`, v.Tag, v.Content, time.Now().Unix()), fmt.Sprintf("%d", squareModel.Id))

		//入redisZset排序
		_, _ = util.Rdb.ZIncrBy(util.Ctx, "hotTag:"+createTime.Format("20060102"), 1, v.Tag).Result()
		util.SetTag("toomhub_tag", fmt.Sprintf(`{"script" : {"source": "ctx._source.hot += 1"},"upsert" : {"hot" : 1,"tag": "%s", "created_at":%d}}`, v.Tag, time.Now().Unix()), fmt.Sprintf("%d", squareModel.Id))
	}

	return true, nil
}

// @title 获取信息详情信息及热门评论
func GetSquareView(id int64, c *gin.Context) (interface{}, error) {

	type imageModel struct {
		Ext   string `json:"ext"`
		Name  string `json:"name"`
		Param string `json:"param"`
		Size  int64  `json:"size"`
		Host  string `json:"host"`
	}

	rdb := util.Rdb
	res, _ := rdb.HMGet(util.Ctx, util.SquareCacheKey+fmt.Sprintf("%d", id), []string{"content", "created_by", "created_at", "collect_count", "likes_count", "argument_count", "created_at", "image", "tag", "id", "video", "cover"}...).Result()
	response := map[string]interface{}{}
	creatorInfo, _ := util.RedisGetUserInfo(res[1].(string), []string{"nick_name", "avatar_url"}...)
	response["created_by"] = creatorInfo[0]
	response["avatar_url"] = creatorInfo[1]
	intCreatedAt, _ := strconv.ParseInt(res[2].(string), 10, 64)
	createdAt := util.StrTime(intCreatedAt)
	response["created_at"] = createdAt

	identity := util.Identity(c)

	isLike := 0
	_, err := rdb.ZRank(util.Ctx, util.SquareLikeKey+fmt.Sprintf("%d", identity.MiniId), fmt.Sprintf("%d", id)).Result()

	fmt.Println("view Err", err)
	fmt.Println("view key", util.SquareLikeKey+res[9].(string))
	if err == nil {
		isLike = 1
	}
	response["is_like"] = isLike
	response["id"] = res[9]
	response["content"] = res[0]
	response["tag"] = res[8]
	response["like_count"] = util.ToInt(res[4].(string))
	response["argument_count"] = util.ToInt(res[5].(string))
	response["collect_count"] = util.ToInt(res[3].(string))

	if res[7] != nil {
		list, _ := util.JsonDecode(res[7].(string))
		var tmp []interface{}
		var tmpL []interface{}
		dat := imageModel{}
		for t := 0; t < len(list); t++ {
			_ = mapstructure.Decode(list[fmt.Sprintf("%d", t)], &dat)
			tmp = append(tmp, dat)
			tmpL = append(tmpL, dat.Host+dat.Name)
		}
		response["list"] = tmpL
		response["image"] = tmp
		response["type"] = 0
	} else {
		response["video"] = res[10]
		response["cover"] = res[11]
		response["type"] = 1
	}

	return response, nil

}
