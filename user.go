package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IsOK struct {
	IsOK   bool    `json:"IsOK"`
}

func main() {
	apiServer := gin.Default()

	apiServer.POST("v1/user/create", create)
	apiServer.POST("v1/user/delete", delete)
	apiServer.POST("v1/user/pwd/change", pwd_change)
	apiServer.GET("v1/user/login", login)

	apiServer.Run(":8080")
}

func create(context *gin.Context) {
	Account := context.PostForm("Account")
	Password := context.PostForm("Password")

	message := "ID:" + Account + ", pwd:" + Password
	isOK := true
	context.JSON(http.StatusOK, gin.H{
		"Code":  0,
		"message": message,
		"Result": IsOK{
			IsOK: isOK,
		},
	})
}

func delete(context *gin.Context) {
	Account := context.PostForm("Account")

	message := "ID:" + Account
	isOK := true
	context.JSON(http.StatusOK, gin.H{
		"Code":  0,
		"message": message,
		"Result": IsOK{
			IsOK: isOK,
		},
	})
}

func pwd_change(context *gin.Context) {
	Account := context.PostForm("Account")
	Password := context.PostForm("Password")

	message := "ID:" + Account + ", pwd:" + Password
	isOK := true
	context.JSON(http.StatusOK, gin.H{
		"Code":  0,
		"message": message,
		"Result": IsOK{
			IsOK: isOK,
		},
	})
}

func login(context *gin.Context) {
	Account := context.Query("Account")
	Password := context.Query("Password")

	message := Account + " is " + Password
	context.String(http.StatusOK, message)
}