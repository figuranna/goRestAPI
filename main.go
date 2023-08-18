package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID    string `json:"id"`
	Album string `json:"album"`
	IsNew bool   `json:"isnew"`
}

var albums = []album{
	{ID: "1", Album: "THE WORLD EP.2: OUTLAW", IsNew: true},
	{ID: "2", Album: "NOEASY", IsNew: false},
	{ID: "3", Album: "5-STAR", IsNew: true},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(context *gin.Context) {
	id := context.Param("id")
	album, err := getAlbumById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, album)
}

func toggleAlbum(context *gin.Context) {
	id := context.Param("id")
	album, err := getAlbumById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	album.IsNew = !album.IsNew

	context.IndentedJSON(http.StatusOK, album)
}

func getAlbumById(id string) (*album, error) {

	for i, t := range albums {
		if t.ID == id {
			return &albums[i], nil
		}
	}

	return nil, errors.New("album not found")
}

func addAlbum(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	context.IndentedJSON(http.StatusCreated, getAlbums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.PATCH("/albums/:id", toggleAlbum)
	router.POST("/albums", addAlbum)
	router.Run("localhost:9090")
}
