package utils

import (
	"achilles/global"
	systemReq "achilles/models/system/request"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从 Gin 的 Context 中获取从 jwt 解析信息失败，请检查请求头是否存在 x-token 且 claims 是否为规定结构")
	}
	return claims, err
}

// 从 Gin 的 Context 中获取从 jwt 解析出来的用户 ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.ID
	}
}

// 从 Gin 的 Context 中获取从 jwt 解析出来的用户 UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// 从 Gin 的 Context 中获取从 jwt 解析出来的用户角色 id
func GetUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.AuthorityId
	}
}

// 从 Gin 的 Context 中获取从 jwt 解析出来的用户角色 id
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}
