package models

type Member struct {
	Model
	NickName string `gorm:"unique_index" json:"nickname"`
	Password string `json:"password"`
	Email    string `gorm:"index" json:"email"`
}

func CheckMember(nickname, password string) bool {
	var member Member
	db.Select("id").Where(Member{NickName: nickname, Password: password}).First(&member)
	if member.ID > 0 {
		return true
	}

	return false
}
