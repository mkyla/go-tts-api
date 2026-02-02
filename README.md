# Go TTS API

A lightweight and minimalist real-time text-to-speech web API developed in Go. It allows users to convert text to high-quality MP3 audio via HTTP POST requests, using the Piper neural TTS engine for offline synthesis.

## 📖 Overview

Go TTS API is a simple HTTP server built with Gin that provides text-to-speech functionality. It accepts JSON payloads with text and returns MP3 audio, leveraging Piper for natural-sounding speech without cloud dependencies.

## 🎬 Demo

Start server: `go run main.go`

Curl: `curl -X POST http://localhost:8080/tts -H "Content-Type: application/json" -d '{"text": "Hello!"}' --output hello.mp3`

Play hello.mp3.

## ✨ Features

### 🔊 High-Quality TTS
- Piper neural engine for realistic speech.
- MP3 output.
- Offline processing.

### 🚀 Lightweight and Efficient
- Gin for fast routing.
- Low overhead.
- Single binary deployable.

### 🛠️ Developer-Friendly
- RESTful API.
- Easy integration.
- Go-based reliability.

## 📦 Installation

### 📋 Binary Releases

Download from [GitHub Releases](https://github.com/mkyla/go-tts-api/releases).

**Supported Platforms:**
- 🐧 Linux: amd64
- 🪟 Windows: amd64
- 🍎 macOS: amd64, arm64

### 🔧 Compile from Source

# Clone
git clone https://github.com/mkyla/go-tts-api.git
cd go-tts-api

# Build
go mod tidy
go build -o go-tts-api main.go

# Run
./go-tts-api

### 📝 Setup Piper

Download Piper binary and model as in tts-rust README.

## 📋 Usage Guide

- 🌐 Start: `./go-tts-api`
- 🔄 Request: POST /tts with {"text": "your text"}
- 💾 Response: MP3 binary

### API Endpoints

POST /tts
- Body: {"text": "string"}
- Response: MP3 audio

## ⚙️ Configuration

- Port: 8080 (hardcoded)
- Piper path: Assumes ./piper/piper

## 🛠️ Development

### 📁 Project Structure

```
go-tts-api/
├── main.go           # Server logic
├── go.mod            # Modules
├── go.sum            # Checksums
├── LICENSE           # BSD-3-Clause
├── README.md         # Docs
└── models/           # Piper models (optional)
```

### 🧩 Core Components

1. **main.go**: Gin server, /tts handler, Piper spawning.

### 🛠️ Tech Stack

- Backend: Go 1.21+
- Framework: Gin
- TTS: Piper
- Build: Go toolchain

### 💻 Development Setup

# 1. Install Go
# (Assume installed)

# 2. Clone
git clone https://github.com/mkyla/go-tts-api.git
cd go-tts-api

# 3. Run
go run main.go

### 🔨 Building

# Local
go build -o go-tts-api main.go

# Cross-compile
GOOS=linux GOARCH=amd64 go build -o go-tts-api-linux main.go

## 📄 License

BSD-3-Clause

## 📊 Badges

![License](https://img.shields.io/badge/license-BSD--3--Clause-blue)