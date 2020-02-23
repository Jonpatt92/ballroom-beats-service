package main

import (
	"fmt"
	"net/http"
	// "strconv" // Was removing automatically so I commented it out
	// "reflect" // Was removing automatically so I commented it out
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// development
// const (
// 	host   = "localhost"
// 	port   = 5432
// 	dbname = "ballroom_beats_development"
// )

// production
const (
	host   = "ec2-54-197-34-207.compute-1.amazonaws.com"
	port   = 5432
	dbname = "ddpi9katt80n0p"
	dbusername = "hlbjcopbpxwheo"
	dbpassword = "a98b2d22b63444ded26df1986a3351b23f4be1f69bf35d95fa2173c2d6591b0a"
)


var db *gorm.DB

func init() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", host, port, dbname) // development
	psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable user=%s password=%s", host, port, dbname, dbusername, dbpassword)
	var err error

	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// db.DropTable(&songTable{}) // We need to refactor our migrations to account for 'rollbacks' where we drop our tables and rebuild.
	db.AutoMigrate(&songTable{})
}

func main() {

	router := gin.Default()
	songs := router.Group("/api/v1/songs")

	{
		songs.POST("/", addSong)
		songs.GET("/", fetchAllSongs)
		songs.GET("/:id", fetchSong)
		songs.DELETE("/:id", removeSong)
	}
	router.Run()
}

type (
	// Consider renaming table as 'songs'
	songTable struct {
		gorm.Model
		Title          string  `json:"title"`
		SpotifyId      string  `json:"spotify_id"`
		URL            string  `json:"url"`
		Delay          float64 `json:"delay"`
		AvgBarDuration float64 `json:"avg_bar_duration"`
		Duration       float64 `json:"duration"`
		Tempo          float64 `json:"tempo"`
		TimeSignature  int64   `json:"time_signature"`
	}
	songInput struct {
		Title          string  `json:"title"`
		SpotifyId      string  `json:"spotify_id"`
		URL            string  `json:"url"`
		Delay          float64 `json:"delay"`
		AvgBarDuration float64 `json:"avg_bar_duration"`
		Duration       float64 `json:"duration"`
		Tempo          float64 `json:"tempo"`
		TimeSignature  int64   `json:"time_signature"`
	}

	transformedSong struct {
		ID             uint    `json:"id"`
		Title          string  `json:"title"`
		SpotifyId      string  `json:"spotify_id"`
		URL            string  `json:"url"`
		Delay          float64 `json:"delay"`
		AvgBarDuration float64 `json:"avg_bar_duration"`
		Duration       float64 `json:"duration"`
		Tempo          float64 `json:"tempo"`
		TimeSignature  int64   `json:"time_signature"`
	}
)

func addSong(context *gin.Context) {

	var body songInput
	context.BindJSON(&body)

	song := songTable{
		Title:          body.Title,
		SpotifyId:      body.SpotifyId,
		URL:            body.URL,
		Delay:          body.Delay,
		AvgBarDuration: body.AvgBarDuration,
		Duration:       body.Duration,
		Tempo:          body.Tempo,
		TimeSignature:  body.TimeSignature,
	}

	db.Save(&song)

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Song created successfully!", "resourceId": song.ID})
}

func fetchAllSongs(context *gin.Context) {

	var songs []songTable

	db.Find(&songs)

	if len(songs) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No songs found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": songs})
}

func fetchSong(context *gin.Context) {
	var song songTable
	songID := context.Param("id")

	db.First(&song, songID)

	if song.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
		return
	}

	_song := transformedSong{ID: song.ID, Title: song.Title, SpotifyId: song.SpotifyId, URL: song.URL, Delay: song.Delay, AvgBarDuration: song.AvgBarDuration, Duration: song.Duration, Tempo: song.Tempo, TimeSignature: song.TimeSignature}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _song})
}

func removeSong(context *gin.Context) {
	var song songTable

	songID := context.Param("id")

	db.First(&song, songID)

	if song.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
		return
	}

	db.Delete(&song)

	context.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Successfully deleted!"})

}
