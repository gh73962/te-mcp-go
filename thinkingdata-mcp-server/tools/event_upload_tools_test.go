package tools_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/myorg/thinkingdata-mcp-server/tools"
)

func TestEventUploadBatchTool(t *testing.T) {
	tool := &tools.EventUploadBatchTool{}

	if tool.Name() != "event_upload_batch" {
		t.Errorf("Unexpected name: got %s, want event_upload_batch", tool.Name())
	}

	if tool.Description() == "" {
		t.Error("Description is empty")
	}

	expectedInputType := reflect.TypeOf(&tools.EventUploadBatchRequest{})
	actualInputType := reflect.TypeOf(tool.Input())
	if actualInputType != expectedInputType {
		t.Errorf("Unexpected input type: got %T, want %T", tool.Input(), &tools.EventUploadBatchRequest{})
	}

	expectedOutputType := reflect.TypeOf(&tools.EventUploadBatchResponse{})
	actualOutputType := reflect.TypeOf(tool.Output())
	if actualOutputType != expectedOutputType {
		t.Errorf("Unexpected output type: got %T, want %T", tool.Output(), &tools.EventUploadBatchResponse{})
	}

	event := tools.EventData{
		EventName:  "test_event",
		Time:       time.Now(),
		UserID:     "user1",
		Properties: map[string]interface{}{"p1": "v1"},
	}
	req := &tools.EventUploadBatchRequest{
		AppID:     "testapp",
		EventList: []tools.EventData{event},
	}

	resp, err := tool.Execute(req)
	if err != nil {
		t.Errorf("Execute returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("Execute response is nil")
	}

	actualResp, ok := resp.(*tools.EventUploadBatchResponse)
	if !ok {
		t.Fatalf("Execute response is not *EventUploadBatchResponse: got %T", resp)
	}

	if actualResp.Code != 0 {
		t.Errorf("Unexpected response code: got %d, want 0", actualResp.Code)
	}
}
