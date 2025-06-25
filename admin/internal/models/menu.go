package models

import (
	"time"
)

// Menu 菜单模型
type Menu struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	RuleID       uint       `gorm:"column:rule_id;default:0" json:"rule_id"`
	Status       int        `gorm:"default:1" json:"status"`
	CreateTime   time.Time  `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time  `gorm:"column:update_time" json:"update_time"`
	Name         string     `gorm:"size:100;not null" json:"name"`
	Desc         string     `gorm:"size:100" json:"desc"`
	FrontPath    *string    `gorm:"column:frontpath" json:"frontpath"`
	FrontFilePath *string   `gorm:"column:front_file_path" json:"front_file_path"`
	Condition    *string    `json:"condition"`
	Menu         int        `gorm:"default:1" json:"menu"`
	Order        int        `gorm:"default:0" json:"order"`
	Icon         string     `gorm:"size:100" json:"icon"`
	Method       string     `gorm:"size:10" json:"method"`
	Child        []Menu     `gorm:"-" json:"child,omitempty"`
}

// TableName 设置表名
func (Menu) TableName() string {
	return "menu"
} 