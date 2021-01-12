package api

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/benny1213/etLab_BE/models"
	"github.com/benny1213/etLab_BE/pkg/e"
	"github.com/benny1213/etLab_BE/pkg/util"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth : 获取token
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	ok, _ := valid.Valid(auth{Username: username, Password: password})

	code := e.InvalidParams
	data := make(map[string]string)
	if ok {
		if isExist := models.CheckAuth(username, password); isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ErrorAuthCheckTokenFail
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
