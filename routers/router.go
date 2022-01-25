package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"CampingNow/pkg/setting"
	articleApi "CampingNow/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.LoadHTMLGlob("templates/*")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

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

	// Articles api, 作业代码 4
	api := r.Group("/api")
	{
		// 获取文章列表
		api.GET("/articles", articleApi.GetArticles)
		// 获取指定文章
		api.GET("/articles/:id", articleApi.GetArticle)
		// 新建文章
		api.POST("/articles", articleApi.AddArticle)
		// 更新指定文章
		api.PUT("/articles/:id", articleApi.EditArticle)
		// 删除指定文章
		api.DELETE("/articles/:id", articleApi.DeleteArticle)
		// 上传文件作为文章内容
		api.POST("/articles/upload", articleApi.UploadFileAsArticleContext)
	}

	return r
}
