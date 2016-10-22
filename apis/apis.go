package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	utils "tinyUrl/utils"
	models "tinyUrl/models"
)

type ShortenUrlRequest struct {
	Url	string        `json:"url" binding:"required"`
}

type OriginalUrlRequest struct {
	Short 	string        `json:"short" binding:"required"`
}

func GetOriginalUrl(c *gin.Context) {

	var json OriginalUrlRequest

	err := c.BindJSON(&json)

	if err != nil {
		utils.AbortWithError(c, http.StatusBadRequest, "wrong request format")
		return
	}

	original_url, err := models.GetFullUrl(json.Short)
	if err != nil {
		utils.AbortWithError(c, http.StatusOK, "not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url":original_url.Full_url,
	})
}

func CreateNewShortenUrl(c *gin.Context) {

	var json ShortenUrlRequest

	err := c.BindJSON(&json)
	if err != nil{
		utils.AbortWithError(c, http.StatusBadRequest, "wrong request format")
		return
	}

	host_url := c.Request.Host
	shorten_url := host_url+"/"+utils.UrlShortener()

	short_url, err := models.CreateUrl(json.Url, shorten_url)
	if err != nil{
		utils.AbortWithError(c, http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short": short_url.Short_url,
	})
}
