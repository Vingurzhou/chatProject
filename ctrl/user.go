package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	// mobile := c.PostForm("mobile")
	// passwd := c.PostForm("passwd")
	c.JSON(http.StatusOK, gin.H{
		Code: code,
	})
}
