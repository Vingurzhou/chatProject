package ctrl

import (
	"chatProject/model"
	"chatProject/service"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService service.UserService

func UserLogin(c *gin.Context) {
	mobile := c.PostForm("mobile")
	passwd := c.PostForm("passwd")
	user, err := userService.Login(mobile, passwd)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "",
			"data": user,
		})
	}
}

func UserRegister(c *gin.Context) {

	mobile := c.PostForm("mobile")
	plainpwd := c.PostForm("plainpwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "",
			"data": user,
		})
	}

}
