package util

import "time"

//广场消息image类型
const SquareTypeImage = 0

//广场消息video类型
const SquareTypeVideo = 1

//小程序用户cacheKey
const UserCacheKey = "mini:user:"

//
const SquareCacheKey = "square:id:"

//
const SquareLikeKey = "mini:user:liked:"

//-------------------- jwt --------------------
// jwt 令牌过期时间
const JwtExpire = 60 * time.Second

//OAuth状态

const OAuthGithub = 0

// 验证码redisKey
const RedisMobileKey = ":mobile_code"
