package main

import (
	"example/user/hello/initializers"
	"example/user/hello/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": http.StatusOK,
		})
	})
	r.POST("/", routes.CreateData)
	r.PUT("/update/:movieID", routes.UpdateData)
	r.GET("/readAll", routes.ReadAll)
	r.GET("/readOne/:movieID", routes.ReadOne)
	r.DELETE("/delete/:movieID", routes.DeletePost)

	r.Run()
}
