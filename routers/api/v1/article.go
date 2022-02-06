package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"CampingNow/models"
	"CampingNow/pkg/e"
	"CampingNow/pkg/setting"
	"CampingNow/pkg/util"
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.INVALID_PARAMS
	var data interface{}
	if models.ExistArticleByID(id) {
		data = models.GetArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticles 获取多个文章, 作业代码 4
func GetArticles(c *gin.Context) {

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	var authId int = -1
	if arg := c.Query("auth_id"); arg != "" {
		authId = com.StrTo(arg).MustInt()
		maps["auth_id"] = authId
	}

	code := e.SUCCESS

	data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
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
	c.HTML(http.StatusOK, "upload.html", gin.H{
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
