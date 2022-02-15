package model

import (
	"course_select/src/database"
	types "course_select/src/global"
	"errors"
	"strconv"
)

type Member struct {
	UserID    int            `json:"user_id" form:"user_id" gorm:"primary_key"`
	Nickname  string         `json:"nickname" form:"nickname"`
	Username  string         `json:"username" form:"username" gorm:"unique"`
	Password  string         `json:"password" form:"password"`
	UserType  types.UserType `json:"user_type" form:"user_type"`
	IsDeleted bool           `json:"is_deleted" form:"is_deleted"`
}

func (Member) TableName() string {
	return "member"
}

/*
func (member *Member) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	return scope.SetColumn("user_id", uuid)
}*/

func (member *Member) CreateMember() (string, error) {

	if err := db.Create(&member).Error; err != nil {
		return "", err
	}

	return strconv.Itoa(member.UserID), nil
}

func (model *Member) GetMember(user_id string) (Member, error) {
	var result Member
	err := database.MySqlDb.First(&Member{}, "user_id = ?", user_id).Scan(&result).Error
	return result, err
}

// GetAllMembers 返回所有成员
func (member *Member) GetAllMembers(offset, limit int) ([]Member, error) {
	var ans []Member
	err := database.MySqlDb.Limit(limit).Offset(offset).Find(&ans).Error
	if err != nil {
		return ans, err
	}
	return ans, nil
}

/*func GetMemberByUsernameAndPassword(username, password string) (Member, error) {
	var ans = Member{}
	err := db.Where("username = ? AND password = ?", username, password).First(&ans).Error
	return ans, err
}*/

func GetMemberByUsername(username string) (Member, error) {
	var ans = Member{}
	err := db.Where("username = ? ", username).First(&ans).Error
	return ans, err
}

func UpdateMember(user_id string, nickname string) error {
	id, _ := strconv.Atoi(user_id)
	var result = Member{}
	db.Where(&Member{UserID: id}).First(&result)
	if result.Nickname == "" {
		return errors.New("未找到该用户！")
	}
	db.Model(&Member{}).Where("user_id = ?", user_id).Update("nickname", nickname)
	return nil
}

func DeleteMember(user_id string) error {
	var ans = Member{}
	db.Where("user_id = ? ", user_id).First(&ans)
	if ans.Nickname == "" {
		return errors.New("未找到该用户")
	}
	if ans.IsDeleted == true {
		return errors.New("用户已删除")
	}

	id, _ := strconv.Atoi(user_id)
	db.Model(&Member{}).Where("user_id = ?", id).Update("is_deleted", true)
	return nil
}
