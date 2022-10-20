package crmexample

import (
	"net/http"
	"strconv"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto/crmexample"
	"go-admin/common/log"
	"go-admin/tools/app"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var req crmexample.SearchReq
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("参数绑定错误: %v", err)
		app.Error(c, http.StatusBadRequest, err, "参数错误")
		return
	}

	list, count, err := crmexample.List(c, req)
	if err != nil {
		log.Errorf("查询列表失败: %v", err)
		app.Error(c, http.StatusInternalServerError, err, "查询失败")
		return
	}

	app.PageOK(c, list, int(count), req.PageIndex, req.PageSize, "ok")
}

func View(c *gin.Context) {
	param := c.Param("id")
	ID, err := strconv.Atoi(param)
	if err != nil {
		log.Errorf("参数绑定错误: %v", err)
		app.Error(c, http.StatusBadRequest, err, "参数错误")
		return
	}

	greet, err := crmexample.GetByID(c, ID)
	if err != nil {
		log.Errorf("查询失败: %v", err)
		app.Error(c, http.StatusInternalServerError, err, "查询失败")
		return
	}
	app.OK(c, greet, "查询成功")
}

func Create(c *gin.Context) {
	var m models.CrmExample
	if err := c.ShouldBind(m); err != nil {
		log.Errorf("参数绑定错误: %v", err)
		app.Error(c, http.StatusBadRequest, err, "参数错误")
		return
	}

	if err := crmexample.Create(c, &m); err != nil {
		log.Errorf("创建失败: %v", err)
		app.Error(c, http.StatusInternalServerError, err, "添加失败")
		return
	}
	app.OK(c, m, "创建成功")
}

func Delete(c *gin.Context) {
	param := c.Param("id")
	ID, err := strconv.Atoi(param)
	if err != nil {
		log.Errorf("参数绑定错误: %v", err)
		app.Error(c, http.StatusBadRequest, err, "参数错误")
		return
	}

	if err = crmexample.DeleteByID(c, ID); err != nil {
		log.Errorf("删除失败: %v", err)
		app.Error(c, http.StatusInternalServerError, err, "删除失败")
		return
	}
	app.OK(c, nil, "删除成功")
}

func Update(c *gin.Context) {
	param := c.Param("id")
	ID, err := strconv.Atoi(param)
	if err != nil {
		log.Errorf("参数绑定错误: %v", err)
		app.Error(c, http.StatusBadRequest, err, "参数错误")
		return
	}

	var m models.CrmExample
	m.ID = ID
	if err = crmexample.UpdateByID(c, &m); err != nil {
		log.Errorf("更新失败: %v", err)
		app.Error(c, http.StatusInternalServerError, err, "更新失败")
		return
	}
	app.OK(c, nil, "更新成功")
}
