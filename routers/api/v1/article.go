package v1

import (
	"fmt"
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

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  "response success",
		"error": "",
	})

	// 煎鱼文档的代码，需要使用 gorm v2 实现 GetArticle 方法，位置在 JSON 前
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
}

// GetArticles 获取多个文章, 作业代码 4
func GetArticles(c *gin.Context) {

	code := e.SUCCESS
	fmt.Println("The data from database pour into the data and return in JSON")

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

// UploadFileAsArticleContext 上传文件作为文章内容, 作业代码 5
func UploadFileAsArticleContext(c *gin.Context) {

	code := e.SUCCESS
	file, error := c.FormFile("file")
	if error != nil {
		code = e.ERROR_UPLOAD_FILE
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.GetMsg(code),
			"error":   code,
		})
		return
	}

	err := c.SaveUploadedFile(file, "./data/"+file.Filename)
	if err != nil {
		return
	}

	address := "127.0.0.1:8000" + "/data/" + file.Filename

	s := fmt.Sprintf("%s uploaded successful", file.Filename)
	c.HTML(http.StatusOK, "api/articles/upload.html", gin.H{
		"code": code,
		"msg":  s,
		"data": file,
	})
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     s,
		"data":    file,
		"address": address,
	})

}

// GetArticleContentFileAddress 获取文章内容文件的地址, 作业代码 6
func GetArticleContentFileAddress(c *gin.Context) {
	filename := c.Query("filename")

	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	address := reqIP + ":" + "8000/" + "data/" + filename

	c.JSON(http.StatusOK, gin.H{
		"address": address,
	})

}
