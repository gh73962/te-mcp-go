package tools_test

import (
	"reflect"
	"testing"

	"github.com/myorg/thinkingdata-mcp-server/tools"
)

func TestUserPropertySetTool(t *testing.T) {
	tool := &tools.UserPropertySetTool{}

	if tool.Name() != "user_property_set" {
		t.Errorf("Unexpected name: got %s, want user_property_set", tool.Name())
	}

	if tool.Description() == "" {
		t.Error("Description is empty")
	}

	expectedInputType := reflect.TypeOf(&tools.UserPropertySetRequest{})
	actualInputType := reflect.TypeOf(tool.Input())
	if actualInputType != expectedInputType {
		t.Errorf("Unexpected input type: got %T, want %T", tool.Input(), &tools.UserPropertySetRequest{})
	}

	expectedOutputType := reflect.TypeOf(&tools.UserPropertySetResponse{})
	actualOutputType := reflect.TypeOf(tool.Output())
	if actualOutputType != expectedOutputType {
		t.Errorf("Unexpected output type: got %T, want %T", tool.Output(), &tools.UserPropertySetResponse{})
	}

	req := &tools.UserPropertySetRequest{
		AppID:      "testapp",
		UserIDList: []string{"user1"},
		PropertyList: []map[string]interface{}{
			{"prop1": "val1"},
		},
	}

	resp, err := tool.Execute(req)
	if err != nil {
		t.Errorf("Execute returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("Execute response is nil")
	}

	actualResp, ok := resp.(*tools.UserPropertySetResponse)
	if !ok {
		t.Fatalf("Execute response is not *UserPropertySetResponse: got %T", resp)
	}

	if actualResp.Code != 0 {
		t.Errorf("Unexpected response code: got %d, want 0", actualResp.Code)
	}
}

func TestUserPropertyDeleteTool(t *testing.T) {
	tool := &tools.UserPropertyDeleteTool{}

	if tool.Name() != "user_property_delete" {
		t.Errorf("Unexpected name: got %s, want user_property_delete", tool.Name())
	}

	if tool.Description() == "" {
		t.Error("Description is empty")
	}

	expectedInputType := reflect.TypeOf(&tools.UserPropertyDeleteRequest{})
	actualInputType := reflect.TypeOf(tool.Input())
	if actualInputType != expectedInputType {
		t.Errorf("Unexpected input type: got %T, want %T", tool.Input(), &tools.UserPropertyDeleteRequest{})
	}

	expectedOutputType := reflect.TypeOf(&tools.UserPropertyDeleteResponse{})
	actualOutputType := reflect.TypeOf(tool.Output())
	if actualOutputType != expectedOutputType {
		t.Errorf("Unexpected output type: got %T, want %T", tool.Output(), &tools.UserPropertyDeleteResponse{})
	}

	req := &tools.UserPropertyDeleteRequest{
		AppID:      "testapp",
		UserIDList: []string{"user1"},
	}

	resp, err := tool.Execute(req)
	if err != nil {
		t.Errorf("Execute returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("Execute response is nil")
	}

	actualResp, ok := resp.(*tools.UserPropertyDeleteResponse)
	if !ok {
		t.Fatalf("Execute response is not *UserPropertyDeleteResponse: got %T", resp)
	}

	if actualResp.Code != 0 {
		t.Errorf("Unexpected response code: got %d, want 0", actualResp.Code)
	}
}
