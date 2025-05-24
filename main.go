package main

import ("net/http"
	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Album One", Artist: "Artist A", Price: 9.99},
	{ID: "2", Title: "Album Two", Artist: "Artist B", Price: 12.99},
	{ID: "3", Title: "Album Three", Artist: "Artist C", Price: 15.99},
}


// Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// New handler to add a new item

func postAlbums(c *gin.Context){
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum) // Created 201
}

// Return specific item

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
// Use Context.Param to retrieve the id path parameter from the URL. When you map this handler to a path, you’ll include a placeholder for the parameter in the path.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID) // New route to get album by ID

	router.Run("localhost:8080")
}