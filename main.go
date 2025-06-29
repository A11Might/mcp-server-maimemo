package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer("greeter", "v1.0.0", nil)
	server.AddTools(
		SayHiTool(),
		ListNotepadsTool(),
		CreateNotepadTool(),
		GetNotepadTool(),
		UpdateNotepadTool(),
		DeleteNotepadTool(),
	)
	// Run the server over stdin/stdout, until the client disconnects
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}
