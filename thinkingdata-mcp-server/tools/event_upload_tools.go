package tools

import (
	"fmt"
	"log"
	"time"

	"github.com/mark3labs/mcp-go"
)

// EventData defines the structure for a single event.
type EventData struct {
	EventName  string                 `json:"event_name"`
	Time       time.Time              `json:"time"`
	UserID     string                 `json:"userid"`
	Properties map[string]interface{} `json:"properties"`
}

// EventUploadBatchRequest defines the request structure for batch event upload.
type EventUploadBatchRequest struct {
	AppID     string      `json:"appid"`
	EventList []EventData `json:"event_list"`
}

// EventUploadBatchResponse defines the response structure for batch event upload.
type EventUploadBatchResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// EventUploadBatchTool is the tool for uploading events in batch.
type EventUploadBatchTool struct {
	mcp.ToolBase
}

// Name returns the name of the tool.
func (t *EventUploadBatchTool) Name() string {
	return "event_upload_batch"
}

// Description returns the description of the tool.
func (t *EventUploadBatchTool) Description() string {
	return "Uploads events in batch. Corresponds to POST /event_upload/batch_event."
}

// Input returns an instance of EventUploadBatchRequest.
func (t *EventUploadBatchTool) Input() interface{} {
	return &EventUploadBatchRequest{}
}

// Output returns an instance of EventUploadBatchResponse.
func (t *EventUploadBatchTool) Output() interface{} {
	return &EventUploadBatchResponse{}
}

// Execute performs the action of uploading events in batch.
func (t *EventUploadBatchTool) Execute(input interface{}) (interface{}, error) {
	req, ok := input.(*EventUploadBatchRequest)
	if !ok {
		return nil, fmt.Errorf("invalid input type for EventUploadBatchTool, expected *EventUploadBatchRequest, got %T", input)
	}

	log.Printf("Received EventUploadBatchRequest: %+v\n", req)

	// Mock implementation
	return &EventUploadBatchResponse{Code: 0, Msg: "Successfully processed batch event upload request (mock)"}, nil
}
