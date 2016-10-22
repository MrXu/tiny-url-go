package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	utils "tinyUrl/utils"
	models "tinyUrl/models"
)

type ShortenUrlRequest struct {
	Url	string        `json:"url"`
}

type OriginalUrlRequest struct {
	Short 	string        `json:"short"`
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

	c.JSON(http.StatusOK, original_url)
}

func CreateNewShortenUrl(c *gin.Context) {

	var json ShortenUrlRequest

	err := c.BindJSON(&json)
	if err != nil{
		utils.AbortWithError(c, http.StatusBadRequest, "wrong request format")
		return
	}

	short_url, err := models.CreateUrl(json.Url)
	if err != nil{
		utils.AbortWithError(c, http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(200, short_url)
}
