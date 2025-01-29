package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type message struct {
	ID 	string `json:"id"`
	Data	string `json:"data"`
	Client	string `json:"client"`
}

var messages = []message{
	{ID: "1", Data: "Test Message", Client: "Test Client"},
	{ID: "2", Data: "This is another test message", Client: "Test Client echo"},
	{ID: "3", Data: "This is the last test message", Client: "Test Client last"},
}

func main() {
	router := gin.Default()
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessages)

	router.Run("localhost:8000")
}

func getMessages(c *gin.Context) { c.IndentedJSON(http.StatusOK, messages) }

func postMessages(c *gin.Context) {
	var newMessage message

	if err := c.BindJSON(&newMessage); err != nil { return }

	messages = append(messages, newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}
