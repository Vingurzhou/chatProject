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
	router.Static("asset", "./asset")
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
	router.POST("/contact/createcommunity", ctrl.CreateCommunity)
	router.POST("/contact/loadcommunity", ctrl.LoadCommunity)
	router.POST("/contact/joincommunity", ctrl.JoinCommunity)
	router.POST("/chat", ctrl.Chat)
	router.POST("/attach/upload", ctrl.Upload)

	router.Run(":8000")
}
