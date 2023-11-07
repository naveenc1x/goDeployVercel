package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hellooooo",
		})
	})
}

func init() {
	app = gin.New()
	myRoute(app.Group("/api"))

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
