package main

import (
	"github.com/0x41gawor/pptx-quiz/internal/handlers"
)

func main() {
	server := handlers.NewServer(":3456")
	server.Run()
}
