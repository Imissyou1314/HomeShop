package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(400,  gin.H{
		"status": "fail",
		"err":   err,
	})
}

func InitRouter() http.Handler {
	router := gin.Default()


}
