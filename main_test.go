package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTTSHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/tts", ttsHandler)

	t.Run("invalid json", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/tts", bytes.NewBuffer([]byte("invalid")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("empty text", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/tts", bytes.NewBuffer([]byte(`{"text":""}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	t.Run("missing text", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/tts", bytes.NewBuffer([]byte(`{}`)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})
}