package service

import (
	"github.com/crabshop/admin/config"
	"github.com/crabshop/admin/internal/models"
	"gorm.io/gorm"
)

// MenuService 菜单服务
type MenuService struct {
	db  *gorm.DB
	cfg *config.Config
}

// NewMenuService 创建一个新的菜单服务
func NewMenuService(db *gorm.DB, cfg *config.Config) *MenuService {
	return &MenuService{
		db:  db,
		cfg: cfg,
	}
}

// GetAllMenus 获取所有菜单
func (s *MenuService) GetAllMenus() ([]models.Menu, error) {
	// 先获取所有顶级菜单
	var rootMenus []models.Menu
	if err := s.db.Where("rule_id = 0 AND status = 1").
		Order("`order` asc").
		Find(&rootMenus).Error; err != nil {
		return nil, err
	}

	// 为每个顶级菜单获取子菜单
	for i := range rootMenus {
		var childMenus []models.Menu
		if err := s.db.Where("rule_id = ? AND status = 1", rootMenus[i].ID).
			Order("`order` asc").
			Find(&childMenus).Error; err != nil {
			return nil, err
		}
		rootMenus[i].Child = childMenus
	}

	return rootMenus, nil
} 