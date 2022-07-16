package main

import (
	"chatProject/ctrl"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("view/**/*")
	matches, _ := filepath.Glob("view/**/*")
	for _, s := range matches {
		newS := strings.Replace(strings.Replace(s, ".html", ".shtml", -1), "view", "", -1)
		router.GET(newS, func(c *gin.Context) {
			c.HTML(http.StatusOK, newS, nil)
		})
	}

	router.POST("/user/login", ctrl.UserLogin)
	router.POST("/user/register", ctrl.UserRegister)
	router.POST("/contact/addfriend", ctrl.AddFriend)
	router.POST("/contact/loadfriend", ctrl.LoadFriend)

	router.Run(":8000")
}
