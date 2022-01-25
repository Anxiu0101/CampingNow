package routers

import (
	v1 "CampingNow/routers/api/v1"
	"github.com/gin-gonic/gin"

	"CampingNow/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

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
			c.JSON(200, gin.H{
				"status": 200,
				"data":   "response success",
				"msg":    "ok",
				"error":  "",
			})
		})

		// 3
		test.DELETE("/test", func(c *gin.Context) {
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
	apiv1 := r.Group("/api/v1")
	{
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
	}
	//{
	//	//获取文章列表
	//	apiv1.GET("/articles", v1.GetArticles)
	//	//获取指定文章
	//	apiv1.GET("/articles/:id", v1.GetArticle)
	//	//新建文章
	//	//apiv1.POST("/articles", v1.AddArticle)
	//	//更新指定文章
	//	//apiv1.PUT("/articles/:id", v1.EditArticle)
	//	//删除指定文章
	//	//apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	//}

	return r
}
