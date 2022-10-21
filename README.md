# go-admin-layout

#### 介绍
基于[go-admin-1.2.2](https://github.com/go-admin-team/go-admin) 进行的一些封装，以及一些bug上的修复。

#### 具体封装
- 添加了数据库迁移的工具类，通过工具类生成sys_menu表下的菜单
- 添加了deploy目录，用于通过docker-compose部署go-admin
- 后端和前端分别更新了Dockerfile、Makefile、Nginx.conf等部署文件
- 后端添加了基于MVC的CRUD编码风格的crmexample示例，摒弃了go-admin提供的接口完全依赖gin框架
  - `srv/app/admin/router/crm_example.go`：crmexample的controller层
  - `srv/app/admin/apis/crmexample/crm_example.go`：crmexample的service层
  - `srv/app/admin/service/dto/crmexample/crm_example.go`：crmexample的dao层
  - `srv/app/admin/models/crm_example.go`：crmexample的实体类
- 前端添加了crmexample的页面示例代码，包含了一个页面CRUD的基本要素，在开发后台管理页面时可以直接复制使用
  - web/src/views/crmexample/index.vue

#### 修复的bug
- 修改角色权限列表时，不显示权限列表的问题：https://gitee.com/lcsin/go-admin-layout/blob/6d75b7d68ec4045c10750a4341e00e9cd18c45cb/srv/app/admin/apis/system/menu.go#L60
- 修改除admin角色外，其它角色无法访问除sys_menu定义的path路径外的其它路径：https://gitee.com/lcsin/go-admin-layout/blob/6d75b7d68ec4045c10750a4341e00e9cd18c45cb/srv/app/admin/middleware/permission.go#L28
#### migrate工具类使用示例
1. 创建文件：cmd/migrate/migration/version/1603516925109_crmexample_migrate.go（1603516925109为毫秒级时间戳）
2. 添加以下代码：
```go
package version

import (
	"runtime"

	"go-admin/cmd/migrate/migration"
	"go-admin/cmd/migrate/migration/tool"
	common "go-admin/common/models"

	"gorm.io/gorm"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(filename), _1603516925109Test)
}

// 1603516925109 取自文件名 1603516925109_crmexample_migrate.go
func _1603516925109Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 目录
		if err := tool.GenRootMenu(tx, "crmExampleManage", "Example管理", "form", 3); err != nil {
			return err
		}
		// 菜单，通过crmExampleManage关联
		if err := tool.GenChildMenu(tx, "crm_example", "Example列表", "list", "crmexample", 1, "crmExampleManage"); err != nil {
			return err
		}
		// 菜单类似

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
```
3. go-admin会根据migrate文件的文件名前缀（时间戳部分）作为数据库迁移的版本存储在sys_migrate表中，并根据sys_migrate表中的数据进行更新
#### 页面开发步骤（以crmexample为例）
1. 编写页面migrate文件，参考：[migrate工具类使用示例](https://github.com/lcsin/go-admin-layout#migrate%E5%B7%A5%E5%85%B7%E7%B1%BB%E4%BD%BF%E7%94%A8%E7%A4%BA%E4%BE%8B)
2. 通过命令：`go build && go-admin.exe migrate -c settings.yml` 生成页面菜单
3. 在`web/src/views`目录下新建页面文件`crmexample/index.vue`（页面文件的命名规则为：`tool.GenChildMenu()`方法中的`permission`参数/index.vue）