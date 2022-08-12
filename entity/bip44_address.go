package entity

import (
	"time"
)

type Bip44Address struct {
	Id          uint      `json:"id"`
	chainId     uint      `json:"chainId"`
	Address     string    `json:"address"`
	Index       uint      `json:"index"`
	CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"` // db.NormalTime是自定义类型，解决日期自动格式化
	UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"` // default:now() 使用数据库时间
}
