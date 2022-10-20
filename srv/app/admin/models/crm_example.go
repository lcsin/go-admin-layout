package models

type CrmExample struct {
	ID       int    `json:"ID" gorm:"column:id;type:int;comment:主键;primaryKey"`
	Username string `json:"username" gorm:"column:username;type:varchar(64);comment:用户名"`
}

func (CrmExample) TableName() string {
	return "crm_example"
}
