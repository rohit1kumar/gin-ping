package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rohit1kumar/pgo/config"
	"github.com/rohit1kumar/pgo/controllers"
	_ "github.com/rohit1kumar/pgo/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GinPing API
// @version 1.0
// @description This is a sample API built with Go and Postgres.
// @contact.name Rohit Kumar
// @contact.url https://github.com/rohit1kumar/gin-ping
// @contact.email <email>
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://localhost:3000
func init() {
	godotenv.Load()
	config.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	docsURL := ginSwagger.URL("/docs/doc.json")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, config.GetRandomJoke())
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, docsURL))

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.PATCH("/posts/:id", controllers.UpdatePostById)
	r.DELETE("/posts/:id", controllers.DeletePostById)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
