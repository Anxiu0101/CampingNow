package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"CampingNow/models"
	"CampingNow/pkg/e"
	"CampingNow/pkg/setting"
	"CampingNow/pkg/util"
)

// GetEvent 获取单个备忘事件
func GetEvent(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.INVALID_PARAMS
	var data interface{}
	if models.ExistEventByID(id) {
		data = models.GetEvent(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_EVENT
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetEvents 获取事件列表
func GetEvents(c *gin.Context) {

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	var memberId int = -1
	if arg := c.Query("member_id"); arg != "" {
		memberId = com.StrTo(arg).MustInt()
		maps["member_id"] = memberId
	}

	code := e.SUCCESS

	data["lists"] = models.GetEvents(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
