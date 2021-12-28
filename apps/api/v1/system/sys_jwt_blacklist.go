package system

import (
	"achilles/global"
	"achilles/model/common/response"
	"achilles/model/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct{}

// @Tags Jwt
// @Summary jwt 加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt 作废失败！", zap.Error(err))
		response.FailWithMessage("jwt 作废失败", c)
	} else {
		response.OkWithMessage("jwt 作废成功", c)
	}
}
