package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	// Create a server with a single tool.
	mms, err := NewMaimemoServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	// Run the server over stdin/stdout, until the client disconnects
	if err := mms.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
