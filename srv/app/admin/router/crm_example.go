package router

import (
	"go-admin/app/admin/apis/crmexample"

	"github.com/gin-gonic/gin"

	"go-admin/app/admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCrmGreetRouter)
}

// 需认证的路由代码
func registerCrmGreetRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r := v1.Group("/crmexample").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", crmexample.List)
		r.GET("/:id", crmexample.View)
		r.POST("", crmexample.Create)
		r.PUT("/:id", crmexample.Update)
		r.DELETE("/:id", crmexample.Delete)
	}
}
