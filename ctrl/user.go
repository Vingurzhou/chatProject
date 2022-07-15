package ctrl

import (
	"chatProject/service"
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
