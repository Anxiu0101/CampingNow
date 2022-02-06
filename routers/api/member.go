package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"CampingNow/models"
	"CampingNow/pkg/e"
	"CampingNow/pkg/util"
)

func MemberLogin(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	isExist := models.CheckMember(username, password)
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = e.SUCCESS
		}
	} else {
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func RegisterMember(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	code := e.INVALID_PARAMS
	isExist := models.CheckMember(username, password)
	if !isExist {
		code = e.SUCCESS
		models.RegisterMember(username, password)
		c.Request.URL.Path = "./space"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "Welcome to be a member of CampingNow!",
	})

}

func GetMemberSpace(c *gin.Context) {

}
