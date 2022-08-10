package entity

import (
	"time"
)

type Chain struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	AccountModel string    `json:"account_model"` // 0:utxo-model , 1:account-model
	UseMemo      bool      `json:"use_memo"`      // 需要鏈上備註
	CoinType     uint      `json:"coinType"`
	CreatedDate  time.Time `json:"created_date" gorm:"autoCreateTime"` // db.NormalTime是自定义类型，解决日期自动格式化
	UpdatedDate  time.Time `json:"updated_date" gorm:"autoUpdateTime"` // default:now() 使用数据库时间
}
