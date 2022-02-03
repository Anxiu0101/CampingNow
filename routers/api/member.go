package api

import (
	"CampingNow/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"CampingNow/pkg/e"
)

func MemberLogin(c *gin.Context) {

	nickname := c.Query("nickname")
	password := c.Query("password")

	//data := make(map[string]interface{})
	//code := e.INVALID_PARAMS
	//isExist := models.CheckMember(nickname, password)
	//if isExist {
	//	token, err := util.GenerateToken(nickname, password)
	//	if err != nil {
	//		code = e.ERROR_AUTH_TOKEN
	//	} else {
	//		data["token"] = token
	//
	//		code = e.SUCCESS
	//	}
	//} else {
	//	code = e.ERROR_AUTH
	//}

	code := e.INVALID_PARAMS
	isExist := models.CheckMember(nickname, password)
	if isExist {
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "Success respond",
	})

}
