package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type album  struct{
	ID 		string 	`json:"id"`
	Title	string 	`json:"title"`
	Artist  string 	`json:"artist"`
	Price	float64 `json:"price"`
}

// album slices
var albums = []album{
	{ID:"1",Title:"Hearn",Artist:"I",Price:999999.9},
	{ID:"2",Title:"Equation",Artist:"ME",Price:99999.9},
	{ID:"3",Title:"Eigen",Artist:"nME",Price:999999.9},
	{ID:"4",Title:"Api",Artist:"myCreation",Price:17.99},
}


func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK,albums)
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/id",getAlbumByID)
	router.POST("/albums",postAlbums)
	router.Run("localhost:8080")
}

func postAlbums(c* gin.Context){
	var newAlbum album

	if err:= c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusAccepted, newAlbum)
}

func getAlbumByID(c* gin.Context){
	id:= c.Param("id")

	for _,a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not founs"})
}
