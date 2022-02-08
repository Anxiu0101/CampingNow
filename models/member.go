package models

import (
	"time"
)

type Member struct {
	Model `gorm:"embedded"`

	Username string `gorm:"unique_index" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"index" json:"email"`
}

func ExistMemberByID(id int) bool {

	var member Member
	db.Select("id").Where("id = ?", id).First(&member)
	if member.ID > 0 {
		return true
	}

	return false

}

func CheckMember(username, password string) bool {
	var member Member
	db.Select("id").Where(Member{Username: username, Password: password}).First(&member)
	if member.ID > 0 {
		return true
	}

	return false
}

func RegisterMember(username, password string) {
	member := Member{Username: username,
		Password: password,
		Model: Model{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}
	db.Create(&member)

}

func ResetPassword(username, newPassword string) bool {
	db.Model(&Member{}).Where("username = ?", username).Update("password", newPassword)

	return true
}
