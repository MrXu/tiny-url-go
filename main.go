package main

import (
	gin "github.com/gin-gonic/gin"
	apis "tinyUrl/apis"
	models "tinyUrl/models"
)

func init(){
	var dbCon string = "tinyurl:tinyurl@/tinyurl_db"
	models.InitDB(dbCon)
}

func main(){

	r := gin.New()
	r.Use(gin.Logger())

	/*
	group
	 */
	v1 := r.Group("/api/v1")
	v1.POST("/original", apis.GetOriginalUrl)
	v1.POST("/shorten", apis.CreateNewShortenUrl)

	r.Run(":8080")

}
