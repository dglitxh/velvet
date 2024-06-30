package main

import (
	"net/http"
	"github.com/dglitxh/velvet/util"
	"github.com/gin-gonic/gin"
	// "github.com/spf13/viper"
)

func main() {
	// viper.SetConfigFile("./common/envs/.env")
	// viper.ReadInConfig()


	port := ":8000"
	r := gin.Default()
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "We are LIVE!!!"})
	})

	r.Run(port)

	util.Log(util.DefaultLogFile, "get me litt")
}