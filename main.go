package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type card struct {
	ID        string `json:"id"`
	Word      string `json:"word"`
	Defintion string `json:"defintion"`
}

var cards = []card{
	{ID: "1", Word: "Go", Defintion: "gopher"},
	{ID: "2", Word: "Python", Defintion: "snake"},
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)

		fmt.Printf("%s %s %s %s\n",
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			latency,
		)
	}
}

func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		c.Next()

		fmt.Printf("%d %s %s\n",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}

func main() {
	router := gin.Default()
	router.GET("/cards/:word", getCardByName)
	router.Use(RequestLogger())
	router.Use(ResponseLogger())
	router.Run(":4000")
}

func getCardByName(c *gin.Context) {
	word := c.Param("word")

	for _, cardValue := range cards {
		if cardValue.Word == word {
			c.IndentedJSON(http.StatusOK, cardValue)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "albums not found"})
}
