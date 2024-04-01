package albums

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Create Handler configuration
func Create(router *gin.Engine, database *sql.DB) {
	db = database
	router.GET("/albums", GetAlbums)
	router.POST("/albums", CreateAlbum)
}

// func SetDB(database *sql.DB) {
// 	db = database
// }

// returns a list of albums from the database
func GetAlbums(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var albums []album
	for rows.Next() {
		var a album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			log.Fatal(err)
		}
		albums = append(albums, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// creates a new album in the database
func CreateAlbum(c *gin.Context) {
	var awesomeAlbum album
	if err := c.BindJSON(&awesomeAlbum); err != nil {
		fmt.Println("error binding json: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO albums (id, title, artist, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(awesomeAlbum.ID, awesomeAlbum.Title, awesomeAlbum.Artist, awesomeAlbum.Price); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, awesomeAlbum)
}
