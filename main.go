package main

import (
	
	"net/http"

	"github.com/dglitxh/velvet/common"
	"github.com/dglitxh/velvet/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// "github.com/spf13/viper"
)

func main() {
	common.LoadConfig()
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	
	port := ":8000"
	r := gin.Default()
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "We are LIVE!!!"})
	})
	common.ConnectDb()
	
	util.Log(util.DefaultLogFile, "get me litt")

	r.Run(port)

	
}