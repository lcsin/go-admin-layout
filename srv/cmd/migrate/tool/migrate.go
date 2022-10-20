package tool

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go-admin/app/admin/models"

	"gorm.io/gorm"
)

// GenRootMenu 生成主菜单目录
func GenRootMenu(db *gorm.DB, name, title, icon string, sort int) error {
	if icon == "" {
		icon = "pass"
	}
	menu := models.Menu{
		MenuName:  name,
		Title:     title,
		Icon:      icon,
		Path:      "/" + name,
		MenuType:  "M",
		Action:    "无",
		NoCache:   true,
		Component: "Layout",
		Sort:      sort,
		CreateBy:  "1",
		UpdateBy:  "1",
		Visible:   "0",
		IsFrame:   "1",
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := db.Create(&menu).Error; err != nil {
		return err
	}

	menu.Paths = fmt.Sprintf("/0/%d", menu.MenuId)
	if err := db.Model(&menu).Updates(&menu).Error; err != nil {
		return err
	}
	return nil
}

// GenChildMenu 生成子菜单项
func GenChildMenu(db *gorm.DB, name, title, icon, permission string, sort int, rootMenuName string) error {
	if icon == "" {
		icon = "pass"
	}
	ctx := context.Background()
	var rootMenu models.Menu
	if err := db.WithContext(ctx).Where("menu_name = ?", rootMenuName).First(&rootMenu).Error; err != nil {
		return err
	}

	components := strings.Split(name, "_")
	var component string
	for _, c := range components {
		component += c
	}
	childMenu := models.Menu{
		MenuName:   name,
		Title:      title,
		Icon:       icon,
		Path:       name,
		MenuType:   "C",
		Action:     "无",
		Permission: fmt.Sprintf("%s:%s:list", permission, permission),
		ParentId:   rootMenu.MenuId,
		NoCache:    true,
		Component:  "/" + component + "/index",
		Sort:       sort,
		Visible:    "0",
		CreateBy:   "1",
		UpdateBy:   "1",
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := db.WithContext(ctx).Create(&childMenu).Error; err != nil {
		return err
	}
	childMenu.Paths = fmt.Sprintf("/0/%d/%d", rootMenu.MenuId, childMenu.MenuId)
	if err := db.WithContext(ctx).Model(&childMenu).Updates(&childMenu).Error; err != nil {
		return err
	}

	childMenuFMA := genChildMenuFMA(name, title, permission, childMenu.MenuId, rootMenu.MenuId)
	if err := db.WithContext(ctx).Create(&childMenuFMA).Error; err != nil {
		return err
	}

	return nil
}

func genChildMenuFMA(ChildName, childMenuTitle, perm string, childMenuID, rootMenuID int) []models.Menu {
	var menus []models.Menu
	var menuID = childMenuID

	for i := 1; i <= 4; i++ {
		menuID++
		var title string
		permission := fmt.Sprintf("%s:%s:", perm, perm)
		switch i {
		case 1:
			title = "分页获取" + childMenuTitle
			permission += "query"
		case 2:
			title = "创建" + childMenuTitle
			permission += "add"
		case 3:
			title = "修改" + childMenuTitle
			permission += "edit"
		case 4:
			title = "删除" + childMenuTitle
			permission += "remove"
		}
		menus = append(menus, models.Menu{
			Title:      title,
			Icon:       "pass",
			Path:       ChildName,
			Paths:      fmt.Sprintf("/0/%d/%d/%d", rootMenuID, childMenuID, menuID),
			MenuType:   "F",
			Action:     "无",
			Permission: permission,
			ParentId:   childMenuID,
			NoCache:    true,
			Visible:    "0",
			CreateBy:   "1",
			UpdateBy:   "1",
			BaseModel: models.BaseModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
	}

	var childMenuMID = menuID + 1
	for i := 1; i <= 6; i++ {
		menuID++
		var name, title, path, paths, _type, action string
		var parentID int
		switch i {
		case 1:
			name = ChildName
			title = childMenuTitle
			path = ChildName
			paths = fmt.Sprintf("/0/%d/%d", 63, childMenuMID)
			_type = "M"
			action = "无"
			parentID = 63
		case 2:
			name = ""
			title = "分页获取" + childMenuTitle
			path = "/api/v1/" + perm
			paths = fmt.Sprintf("/0/%d/%d/%d", 63, childMenuMID, menuID)
			_type = "A"
			action = "GET"
			parentID = childMenuMID
		case 3:
			name = ""
			title = "根据id获取" + childMenuTitle
			path = "/api/v1/" + perm + "/:id"
			paths = fmt.Sprintf("/0/%d/%d/%d", 63, childMenuMID, menuID)
			_type = "A"
			action = "GET"
			parentID = childMenuMID
		case 4:
			name = ""
			title = "创建" + childMenuTitle
			path = "/api/v1/" + perm
			paths = fmt.Sprintf("/0/%d/%d/%d", 63, childMenuMID, menuID)
			_type = "A"
			action = "POST"
			parentID = childMenuMID
		case 5:
			name = ""
			title = "修改" + childMenuTitle
			path = "/api/v1/" + perm
			paths = fmt.Sprintf("/0/%d/%d/%d", 63, childMenuMID, menuID)
			_type = "A"
			action = "PUT"
			parentID = childMenuMID
		case 6:
			name = ""
			title = "删除" + childMenuTitle
			path = "/api/v1/" + perm + "/:id"
			paths = fmt.Sprintf("/0/%d/%d/%d", 63, childMenuMID, menuID)
			_type = "A"
			action = "DELETE"
			parentID = childMenuMID
		}
		menus = append(menus, models.Menu{
			MenuName: name,
			Title:    title,
			Icon:     "bug",
			Path:     path,
			Paths:    paths,
			MenuType: _type,
			Action:   action,
			ParentId: parentID,
			NoCache:  true,
			Visible:  "0",
			CreateBy: "1",
			UpdateBy: "1",
			BaseModel: models.BaseModel{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		})
	}

	return menus
}
