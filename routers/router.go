package routers

import (
	"CampingNow/middleware/jwt"
	eventApi "CampingNow/routers/api/event"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	mainApi "CampingNow/routers/api"
	apiV1 "CampingNow/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 加载静态资源和模板文件
	r.StaticFS("/static", http.Dir("./static"))
	r.LoadHTMLGlob("templates/**/*.html")

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

	//Articles api, 作业代码 4 5 6
	api := r.Group("/api")
	api.Use(jwt.JWT())
	{
		// 获取文章列表
		api.GET("/articles", apiV1.GetArticles)
		// 获取指定文章
		api.GET("/articles/:id", apiV1.GetArticle)
		// 新建文章
		api.POST("/articles", apiV1.AddArticle)
		// 更新指定文章
		api.PUT("/articles/:id", apiV1.EditArticle)
		// 删除指定文章
		api.DELETE("/articles/:id", apiV1.DeleteArticle)
		// 上传文件作为文章内容
		api.POST("/articles/upload", apiV1.UploadFileAsArticleContext)
		// 获取文章内容文件的存储地址
		api.GET("/articles/address", apiV1.GetArticleContentFileAddress)
	}

	member := r.Group("/member")
	{
		/* 用户相关功能 */
		// 用户登录
		member.GET("/login", mainApi.MemberLogin)
		// 用户注册
		member.POST("/register", mainApi.RegisterMember)
		// 修改用户密码
		member.PUT("/space", mainApi.ResetMemberPassword)
		// 用户个人主页
		member.GET("/space", mainApi.GetMemberSpace)

		/* 备忘录功能 */
		event := r.Group("/member")
		event.Use(jwt.JWT())
		{
			// 获取备忘事件列表
			event.GET("/events", eventApi.GetEvents)
			// 获取特定状态备忘事件列表 如 101 为 事件进行中
			event.GET("/events/status", eventApi.GetEventByStatus)
			// 获取特定类型备忘录事件列表 如 201 为 生日
			event.GET("/events/type", eventApi.GetEventByType)
			// 获取指定备忘事件
			event.GET("/events/:id", eventApi.GetEvent)
			// 更新备忘事件信息
			event.PUT("/events/:id", eventApi.EditEventByID)
			// 创建新备忘事件
			event.POST("/events", eventApi.AddEvents)
			// 删除指定备忘事件
			event.DELETE("/events/:id", eventApi.DeleteEvent)
		}
	}

	// 加载404错误页面
	r.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"msg": "404",
		})
	})

	return r
}
