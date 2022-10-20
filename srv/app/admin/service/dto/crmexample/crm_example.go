package crmexample

import (
	"go-admin/app/admin/models"
	cDto "go-admin/common/dto"
	"go-admin/tools"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchReq struct {
	PageIndex int    `form:"pageIndex"`
	PageSize  int    `form:"pageSize"`
	ID        int    `form:"int"`
	Username  string `form:"username"`
}

func List(c *gin.Context, condition SearchReq) ([]*models.CrmExample, int64, error) {
	// 获取数据库连接
	db, err := tools.GetOrm(c)
	if err != nil {
		return nil, 0, err
	}

	var count int64
	list := make([]*models.CrmExample, 0)
	if err = db.WithContext(c).Scopes(
		searchCondition(condition),
		cDto.Paginate(condition.PageSize, condition.PageIndex),
	).Order("create_time desc").Find(&list).
		Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func searchCondition(req SearchReq) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.ID != 0 {
			db = db.Where("id = ?", req.ID)
		}
		if req.Username != "" {
			db = db.Where("username = ?", req.Username)
		}
		return db
	}
}

func GetByID(c *gin.Context, ID int) (*models.CrmExample, error) {
	db, err := tools.GetOrm(c)
	if err != nil {
		return nil, err
	}

	var m models.CrmExample
	m.ID = ID
	if err = db.WithContext(c).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func DeleteByID(c *gin.Context, ID int) error {
	db, err := tools.GetOrm(c)
	if err != nil {
		return err
	}

	var m models.CrmExample
	m.ID = ID
	return db.WithContext(c).Delete(&m).Error
}

func Create(c *gin.Context, m *models.CrmExample) error {
	db, err := tools.GetOrm(c)
	if err != nil {
		return err
	}

	return db.WithContext(c).Create(&m).Error
}

func UpdateByID(c *gin.Context, m *models.CrmExample) error {
	db, err := tools.GetOrm(c)
	if err != nil {
		return err
	}

	return db.WithContext(c).Model(&m).Updates(&m).Error
}
