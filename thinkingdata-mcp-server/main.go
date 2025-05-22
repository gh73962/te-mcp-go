package main

import (
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go"
	"github.com/myorg/thinkingdata-mcp-server/tools" // Ensure this matches your module path
)

func main() {
	// Create instances of the tools
	userPropertySetTool := &tools.UserPropertySetTool{}
	userPropertyDeleteTool := &tools.UserPropertyDeleteTool{}
	eventUploadBatchTool := &tools.EventUploadBatchTool{}

	// Create a new MCP server
	server := mcp.NewServer()

	// Register the tools with the server
	err := server.Register(userPropertySetTool)
	if err != nil {
		log.Fatalf("Failed to register UserPropertySetTool: %v", err)
	}

	err = server.Register(userPropertyDeleteTool)
	if err != nil {
		log.Fatalf("Failed to register UserPropertyDeleteTool: %v", err)
	}

	err = server.Register(eventUploadBatchTool)
	if err != nil {
		log.Fatalf("Failed to register EventUploadBatchTool: %v", err)
	}

	// Log a message indicating the server is starting
	log.Println("Starting ThinkingData MCP server on :8080, handling /mcp")

	// Start the HTTP server
	http.Handle("/mcp", server)
	log.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
