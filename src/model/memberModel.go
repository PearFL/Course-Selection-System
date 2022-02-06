package model

import (
	types "course_select/src/global"
)

type Member struct {
	UserID   int            `json:"user_id" form:"user_id" gorm:"primary_key"`
	Nickname string         `json:"nickname" form:"nickname"`
	Username string         `json:"username" form:"username"`
	Password string         `json:"password" form:"password"`
	UserType types.UserType `json:"user_type" form:"user_type"`
}

func (Member) TableName() string {
	return "member"
}
