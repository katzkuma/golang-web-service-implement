package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	apiServer := gin.Default()

	apiServer.GET("v1/user/login", login)

	apiServer.Run(":8080")
}

func login(context *gin.Context) {
	Account := context.Query("account")
	Password := context.Query("password")

	message := Account + " is " + Password
	context.String(http.StatusOK, message)
}