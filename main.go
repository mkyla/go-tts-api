package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/tts", ttsHandler)
	r.Run(":8080")
}

func ttsHandler(c *gin.Context) {
	var req struct {
		Text string `json:"text"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "text is required"})
		return
	}

	// Generate temp file
	tmpFile, err := os.CreateTemp("", "tts_*.mp3")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create temp file"})
		return
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Run piper
	cmd := exec.Command("./bin/piper", "--model", "models/en_US-lessac-medium.onnx", "--output_file", tmpFile.Name())
	cmd.Stdin = bytes.NewReader([]byte(req.Text))
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate speech"})
		return
	}

	// Read the file
	tmpFile.Seek(0, 0)
	data, err := io.ReadAll(tmpFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read audio"})
		return
	}

	c.Header("Content-Type", "audio/mpeg")
	c.Data(http.StatusOK, "audio/mpeg", data)
}