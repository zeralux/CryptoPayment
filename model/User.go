package model

type User struct {
	Name   string   `json:"name" example:"david"`      // 姓名
	Level  string   `json:"level" example:"A"`         // 職級
	Method []string `json:"method" example:"ApplePay"` // 支付方式
}
