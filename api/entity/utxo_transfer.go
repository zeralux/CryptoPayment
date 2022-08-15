package entity

import (
	"math/big"
	"time"
)

type UtxoTransfer struct {
	Id          uint      `json:"id"`
	Status      uint      `json:"status" gorm:"default:1"` // 启用：1， 禁用：2
	Asset       string    `json:"asset"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      big.Int   `json:"amount"`
	Memo        string    `json:"remark"`                             // 鏈上備註
	CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"` // db.NormalTime是自定义类型，解决日期自动格式化
	UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"` // default:now() 使用数据库时间
}
