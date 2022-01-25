package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"CampingNow/pkg/e"
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.SUCCESS
	//code := e.INVALID_PARAMS
	//var data interface{}
	//if !valid.HasErrors() {
	//	if models.ExistArticleByID(id) {
	//		data = models.GetArticle(id)
	//		code = e.SUCCESS
	//	} else {
	//		code = e.ERROR_NOT_EXIST_ARTICLE
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
	//	}
	//}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  "response success",
		"error": "",
	})
}

// GetArticles 获取多个文章, 作业代码 4
func GetArticles(c *gin.Context) {

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  "response success",
		"error": "",
	})
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
}

// UploadFileAsArticleContext 上传文件作为文章内容
func UploadFileAsArticleContext(c *gin.Context) {

	code := e.SUCCESS
	file, err := c.FormFile("testFile")
	if err != nil {
		code = e.ERROR_UPLOAD_FILE
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.GetMsg(code),
			"error":   code,
		})
		return
	}

	log.Fatal(file.Filename)
	c.SaveUploadedFile(file, "./"+file.Filename)
	s := fmt.Sprintf("%s uploaded successful", file.Filename)

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  s,
		"error": "",
	})

}
