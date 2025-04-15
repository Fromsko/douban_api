package api

import (
	"net/http"
	"math/rand"
	"time"
	"strconv"

	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

//go:embed douban-top250.json
var moviesJSON []byte

var moviesData []gjson.Result

func init() {
	rand.Seed(time.Now().UnixNano())
	loadMoviesData()
}

func loadMoviesData() {
	moviesData = gjson.ParseBytes(moviesJSON).Array()
}

func GetMovieByID(c *gin.Context) {
	movieID := c.Param("movie_id")
	id, err := strconv.Atoi(movieID)
	if err != nil || id < 0 || id >= len(moviesData) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, moviesData[id].Map())
}

func GetAllMovies(c *gin.Context) {
	c.JSON(http.StatusOK, moviesData)
}

func GetRandomMovie(c *gin.Context) {
	randomIndex := rand.Intn(len(moviesData))
	c.JSON(http.StatusOK, moviesData[randomIndex].Map())
}

func GetMovieImage(c *gin.Context) {
	randomIndex := rand.Intn(len(moviesData))
	imageURL := moviesData[randomIndex].Get("img").String()
	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}
