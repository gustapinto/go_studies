package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Album struct {
	// Using struct tags to define what field name should be when
	// serializeing the struct to json
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Create a slice to seed record data
var albums = []Album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")

	// Perform a linear search to get the album id
	for _, album := range albums {
		if album.Id == id {
			ctx.IndentedJSON(http.StatusOK, album)

			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(ctx *gin.Context) {
	var newAlbum Album

	//Bind the request body to a pointer
	err := ctx.BindJSON(&newAlbum)

	if err != nil {
		return
	}

	albums = append(albums, newAlbum)

	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	// Open a new gin router
	router := gin.Default()

	// Create endpoints
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	// Start the router
	router.Run("0.0.0.0:80")
}
