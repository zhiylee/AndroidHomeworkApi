package models

import (
	"androidHomeworkApi/pkg/setting"
	"androidHomeworkApi/pkg/util"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID int `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Passwd string `json:"-"`
	Avatar string `json:"avatar"`
	Introduction string `json:"introduction"`
}

func (u *User) AfterFind(tx *gorm.DB)  {
	u.Avatar = setting.CdnHost +  u.Avatar

	return
}

func AddUser(user User) User {
	db.Create(&user)

	return user
}

func GetUserById(id int) (user User,isExist bool) {
	db.First(&user,id)

	if user.ID<1 {
		return user,false
	}

	//user.Avatar = setting.CdnHost +  user.Avatar

	return user,true
}

func GetUserByName(name string) (user User,isExist bool) {
	db.Where("name = ?",name).First(&user)

	if user.ID<1 {
		return user,false
	}

	return user,true
}

func CheckUser(name,passwd string) (isPass bool,user User) {
	db.Where(User{
		Name: name,
	}).First(&user)

	if user.ID<1 {
		return false,User{}
	}

	passHash:=util.SHA256Code(passwd)
	if user.Passwd != passHash {
		return false,User{}
	}

	return true,user
}

func UpdateUserInfo(user *User,info map[string]interface{})  {
	db.Model(&user).Updates(info)
}