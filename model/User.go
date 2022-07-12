package model

type User struct {
	Id   string `json:"id" example:"123"`     // 識別碼
	Name string `json:"name" example:"david"` // 姓名
}
