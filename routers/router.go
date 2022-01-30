package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	apiv1 "CampingNow/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.LoadHTMLGlob("templates/**/**/*")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// 作业代码 2 3
	test := r.Group("/test")
	{
		// 2
		test.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": 200,
				"data":   "response success",
				"msg":    "ok",
				"error":  "",
			})
		})

		// 3
		test.POST("/test", func(c *gin.Context) {
			fmt.Print("Write new data into database")

			c.JSON(200, gin.H{
				"status": 200,
				"data":   "response success",
				"msg":    "ok",
				"error":  "",
			})
		})

		// 3
		test.DELETE("/test", func(c *gin.Context) {
			fmt.Print("Delete data from database")

			c.JSON(200, gin.H{
				"status": 200,
				"data":   "response success",
				"msg":    "ok",
				"error":  "",
			})
		})

		// 3
		test.PUT("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": 200,
				"data":   "response success",
				"msg":    "ok",
				"error":  "",
			})
		})
	}

	// Articles api, 作业代码 4 5 6
	api := r.Group("/api")
	{
		// 获取文章列表
		api.GET("/articles", apiv1.GetArticles)
		// 获取指定文章
		api.GET("/articles/:id", apiv1.GetArticle)
		// 新建文章
		api.POST("/articles", apiv1.AddArticle)
		// 更新指定文章
		api.PUT("/articles/:id", apiv1.EditArticle)
		// 删除指定文章
		api.DELETE("/articles/:id", apiv1.DeleteArticle)
		// 上传文件作为文章内容
		api.POST("/articles/upload", apiv1.UploadFileAsArticleContext)
		// 获取文章内容文件的存储地址
		api.GET("/articles/address", apiv1.GetArticleContentFileAddress)
	}

	login := r.Group("/login")
	{
		login.GET("/", apiv1.MemberLogin)
	}

	return r
}
