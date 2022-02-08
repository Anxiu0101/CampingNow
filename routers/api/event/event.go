package event

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	data["total"] = models.GetEventTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddEvents(c *gin.Context) {

	// 从上下文中获取新备忘事件的信息
	title := c.Query("title")
	memberID := com.StrTo(c.Query("member_id")).MustInt()
	eventType := c.Query("type")
	status := c.Query("status")
	endTime := c.Query("end_time")
	content := c.Query("content")

	/* Validation of input para */
	temp, _ := strconv.Atoi(endTime)
	eTime, _ := time.ParseDuration(strconv.Itoa(24*temp) + "h")
	fmt.Print(eTime)
	code := e.INVALID_PARAMS

	if models.ExistMemberByID(memberID) {
		data := make(map[string]interface{})
		data["title"] = title
		data["member_id"] = memberID
		data["type"] = eventType
		data["status"] = status
		data["end_time"] = time.Now().Add(eTime).Unix()
		data["content"] = content

		if models.AddEvent(data) {
			code = e.SUCCESS
		} else {
			code = e.ERROR
		}
	} else {
		code = e.ERROR_NOT_EXIST_MEMBER
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}
