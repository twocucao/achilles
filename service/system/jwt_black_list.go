package system

import (
	"context"
	"time"

	"go.uber.org/zap"

	"achilles/global"
	"achilles/models/system"
)

type JwtService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑 jwt
//@param: jwtList model.JWTDenyList
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JWTDenyList) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断 JWT 是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JWTDenyList{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从 redis 取 jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt 存入 redis 并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于 jwt 过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GVA_DB.Model(&system.JWTDenyList{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("加载数据库 jwt 黑名单失败！", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt 黑名单 加入 BlackCache 中
}
