package event

import (
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
		models.AddEventViews(id)
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

// GetEvents 获取备忘事件列表
func GetEvents(c *gin.Context) {

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	memberId := -1
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

// GetEventByStatus 通过状态码和类型码获取对应事件列表，如 待办事件 可以通过 101 获取
func GetEventByStatus(c *gin.Context) {

	status, err := strconv.Atoi(c.Query("status"))
	code := e.ERROR
	data := make(map[string]interface{})

	// 检查数据
	if status < 200 && status > 100 {
		data["lists"] = models.GetEventByStatus(status, util.GetPage(c), setting.PageSize)

		code = e.SUCCESS
	} else {
		code = e.INVALID_PARAMS
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"error": err,
	})
}

// GetEventByType 通过类型码获取对应事件列表，如 任务 可以通过 203 获取
func GetEventByType(c *gin.Context) {

	types, err := strconv.Atoi(c.Query("types"))
	data := make(map[string]interface{})
	code := e.ERROR

	// 检查数据
	if types < 300 && types > 200 {
		data["lists"] = models.GetEventByType(types, 10, 1)

		code = e.SUCCESS
	} else {
		code = e.INVALID_PARAMS
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  data,
		"error": err,
	})
}

// EditEventByID 更新单个备忘事件信息
func EditEventByID(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	title := c.Query("title")
	memberID := com.StrTo(c.Query("member_id")).MustInt()
	eventType := c.Query("type")
	status := c.Query("status")
	endTime := c.Query("end_time")
	content := c.Query("content")

	/* Validation */
	temp, _ := strconv.Atoi(endTime)
	eTime, _ := time.ParseDuration(strconv.Itoa(24*temp) + "h")

	code := e.INVALID_PARAMS
	if models.ExistEventByID(id) {
		if models.ExistMemberByID(id) {
			data := make(map[string]interface{})
			data["member_id"] = memberID
			data["title"] = title
			data["type"] = eventType
			data["status"] = status
			data["end_time"] = time.Now().Add(eTime).Unix()
			data["content"] = content
			models.EditEvent(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_MEMBER
		}
	} else {
		code = e.ERROR_NOT_EXIST_EVENT
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// AddEvents 创建新备忘事件
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

// DeleteEvent 删除指定文章
func DeleteEvent(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	/* Validation */
	code := e.INVALID_PARAMS
	if models.ExistEventByID(id) {
		models.DeleteEvent(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "Successful delete this event",
	})
}
