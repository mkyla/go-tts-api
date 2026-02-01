# Go TTS API

A simple Go web API for text-to-speech using Piper.

## Features

- HTTP server with Gin framework
- POST /tts endpoint accepting JSON {"text": "your text"}
- Returns MP3 audio binary
- Uses Piper TTS engine

## Installation

1. Clone the repository
2. Run `go mod tidy`
3. Run `go run main.go`

## Usage

Start the server:

```bash
go run main.go
```

Send a POST request:

```bash
curl -X POST http://localhost:8080/tts -H "Content-Type: application/json" -d '{"text": "Hello, world!"}' --output output.mp3
```

## License

BSD-3-Clause