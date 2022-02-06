package models

type Article struct {
	Model `gorm:"embedded"`

	AuthID int    `json:"auth_id" gorm:"index"`
	Auth   Member `json:"auth" gorm:"foreignKey:AuthID"`

	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	State   int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Auth").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(maps interface{}) (id int) {
	db.Model(&Article{}).Where("id = ?", id)

	return
}

//func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
//	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
//
//	return
//}
//
//func GetArticle(id int) (article Article) {
//	db.Where("id = ?", id).First(&article)
//	//db.Model(&article).Related(&article.Member)
//
//	return
//}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

// 煎鱼的文档代码，使用 gorm v1，需要用 gorm v2 重构
//func AddArticle(data map[string]interface{}) bool {
//	db.Create(&Article{
//		Title:   data["title"].(string),
//		Desc:    data["desc"].(string),
//		Content: data["content"].(string),
//		Auth:    data["auth"].(Member),
//		State:   data["state"].(int),
//	})
//
//	return true
//}
//
//func DeleteArticle(id int) bool {
//	db.Where("id = ?", id).Delete(Article{})
//
//	return true
//}
