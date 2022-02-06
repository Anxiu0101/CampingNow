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
