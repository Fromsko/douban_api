package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"douban-goserver/api"
	"embed"
)

//go:embed static/*
var fs embed.FS

func main() {
	r := gin.Default()

	// Setup CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"}

	r.Use(cors.New(config))

	// Register routes
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/movie/top/:movie_id", api.GetMovieByID)
		apiGroup.GET("/movie/all", api.GetAllMovies)
		apiGroup.GET("/movie/random", api.GetRandomMovie)
		apiGroup.GET("/movie/random-image", api.GetMovieImage)
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static")
	})

	// Serve static files using r.Any
	r.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(fs))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	// 处理未找到路由的情况
	r.NoRoute(func(c *gin.Context) {
		log.Println("No route found, redirecting to /static")
		c.Redirect(http.StatusFound, "/static")
	})

	// Start the server
	r.Run(":8080")
}
