package model

import (
	types "course_select/src/global"
)

type Member struct {
	UserID   string         `json:"user_id" form:"user_id" gorm:"primary_key"`
	Nickname string         `json:"nickname" form:"nickname"`
	Username string         `json:"username" form:"username"`
	Password string         `json:"password" form:"password"`
	UserType types.UserType `json:"user_type" form:"user_type"`
}

func (Member) TableName() string {
	return "member"
}

/*
func (member *Member) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	return scope.SetColumn("user_id", uuid)
}*/

// func md5V(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }
