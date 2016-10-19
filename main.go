package main

import (
	gin "github.com/gin-gonic/gin"
	apis "tinyUrl/apis"
	models "tinyUrl/models"
)

func init(){
	var dbCon string = "postgres://user:pass@localhost/bookstore"
	models.InitDB(dbCon)
}

func main(){

	r := gin.New()
	r.Use(gin.Logger())

	/*
	group
	 */
	v1 := r.Group("/api/v1")
	v1.GET("/original", apis.GetUrl)
	v1.POST("/shorten", apis.PostUrl)

	r.Run(":8080")

}
