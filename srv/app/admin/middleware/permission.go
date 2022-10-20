package middleware

import (
	"context"
	"net/http"

	"go-admin/app/admin/models"

	"github.com/gin-gonic/gin"

	"go-admin/common/global"
	"go-admin/common/log"
	"go-admin/pkg/jwtauth"
	"go-admin/tools"
	"go-admin/tools/app"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		v := data.(jwtauth.MapClaims)
		e := global.CasbinEnforcer
		var res bool
		var err error
		msgID := tools.GenerateMsgIDFromContext(c)
		//检查权限
		if chekRolePermission(v["rolekey"]) {
			res = true
			log.Infof("msgID[%s] info:%s method:%s path:%s", msgID, v["rolekey"], c.Request.Method, c.Request.URL.Path)
		} else {
			res, err = e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
			if err != nil {
				log.Errorf("msgID[%s] error:%s method:%s path:%s", msgID, err, c.Request.Method, c.Request.URL.Path)
				app.Error(c, 500, err, "")
			}
		}

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}

func chekRolePermission(r interface{}) bool {
	var roles []models.SysRole
	if err := global.Eloquent.WithContext(context.Background()).Find(&roles).Error; err != nil {
		log.Errorf("查询角色列表失败: %v", err)
	}
	for _, role := range roles {
		if r == role.RoleKey {
			return true
		}
	}
	return false
}
