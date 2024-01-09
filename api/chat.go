package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/josephe44/chatgtp-api-golang/logger"
	"github.com/josephe44/chatgtp-api-golang/openai"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.POST("/api/chat", handleChatRequest)

	go func() {
		if err := r.Run(":3003"); err != nil {
			logger.Error("Failed to start server: ", err)
		}
	}()

	r.Run()
}

func handleChatRequest(c *gin.Context) {
	apiKey := os.Getenv("RENDER_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "RENDER_API_KEY not set"})
		return
	}

	var requestBody struct {
		Message string `json:"message"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	escapedInput := strconv.Quote(requestBody.Message)

	response, err := openai.GetAIResponse(apiKey, escapedInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting AI response"})
		return
	}

	if len(response.Choices) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Empty AI response"})
		return
	}

	message := response.Choices[0].Message
	if message.Content != "" {
		c.JSON(http.StatusOK, gin.H{"response": message.Content})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Empty AI response"})
	}
}
