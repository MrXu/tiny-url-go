package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUrl(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"original":"",
	})
}

func PostUrl(c *gin.Context) {

	c.JSON(200, gin.H{
		"short":  "",
	})
}
