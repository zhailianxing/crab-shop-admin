package controller

import (
	"net/http"

	"github.com/crabshop/admin/internal/service"
	"github.com/gin-gonic/gin"
)

// MenuController 菜单控制器
type MenuController struct {
	menuService *service.MenuService
}

// NewMenuController 创建一个新的菜单控制器
func NewMenuController(menuService *service.MenuService) *MenuController {
	return &MenuController{
		menuService: menuService,
	}
}

// GetMenus 获取所有菜单
func (c *MenuController) GetMenus(ctx *gin.Context) {
	menus, err := c.menuService.GetAllMenus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取菜单失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"menus": menus,
		},
	})
} 