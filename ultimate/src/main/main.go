package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	gin.SetMode(gin.DebugMode)
    app := gin.Default()
    route := app.Group("/api/v1")
    {
    	route.GET("/user", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"response":context.Query("name"),
			})
		})
    	route.POST("/user", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"response":context.PostForm("name"),
			})
		})
	}
    app.Run(":7074")

}
