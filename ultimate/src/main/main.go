package main

import (
	"fmt"
	"ultimate/src/infrastructure"
)

func main() {

	//gin.SetMode(gin.DebugMode)
	//
	//app := gin.Default()
	//
	//router := app.Group("/api/v1")
	//{
	//	router.GET("/user",)
	//}
	//
	//app.Run(":7074")


	h := &infrastructure.HttpModel{}

	h.Set(h,"ok")

	fmt.Println(h.Response)


}
