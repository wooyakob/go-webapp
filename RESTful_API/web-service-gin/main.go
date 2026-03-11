// standalone program versus a library is always in pkg main
package main

// add pkgs to support endpoint code
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// declaration of an album struct, used to store album data in memory
type album struct {
	ID          string  `json:"id"` // struct's content serialized into JSON
	Title       string  `json:"title"`
	Artist      string  `json:"artist"`
	Price       float64 `json:"price"` // prices include decimals
	Currency    string  `json:"currency"`
	ReleaseYear int     `json:"release_year"` // upper CamelCase, whole number
}

// slice of album structs containing example data
// using prices from here: https://www.discogs.com/sell/release/691253?condition=Very+Good+Plus+%28VG%2B%29&currency=USD
var albums = []album{
	{ID: "1", Title: "Life", Artist: "Simply Red", Price: 6.25, Currency: "USD", ReleaseYear: 1995},
	{ID: "2", Title: "Big Love", Artist: "Simply Red", Price: 11.49, Currency: "USD", ReleaseYear: 2015},
}

// initialize gin router, associate get http method and albums path with handler function
// pass name of function, not result, result would be getAlbums() and include parenthesis
// use run to attach router to http server, start server

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums) // assign handler function to endpoint path
	router.GET("/albums/:id", getAlbumByID) // colon in gin signifies item is a path param!

	router.POST("albums", postAlbums)
	
	router.Run("localhost:8080")
}

// implement first endpoint
// handler to return all items
// client makes a request to GET /albums, return all albums as JSON
// logic to prepare a response
// map request path to logic
// in reverse of execution at runtime
// add dependencies, then code that depends on them

// getAlbums function, create JSON from slice of album structs, write JSON into response

// getAlbums returns list of all albums as JSON
// Context in Gin carries request details, validates and serializes JSON
// indented is easier to work with when debugging and size difference from c.JSON is small

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// POST request to /albums, add the album in request body to existing album's data
// logic to add new album to existing list
func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// handler to return a specific item
// client makes request to GET /albums/[id], return album with matching ID
// add logic to retrieve requested album


// context param to retrieve id path param from URL
// map handler to path, placeholer for path
// loop over album struct, look for one with an id field value that matches ID param
// if found, serialize album struct to JSON and return it as a response with 200 OK HTTP code
// return HTTP 404 error if album not found

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	
	// loop list of albums, identify album with ID value that matches path param
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}






