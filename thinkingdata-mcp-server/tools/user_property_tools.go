package tools

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go"
)

// UserPropertySetRequest defines the request structure for setting user properties.
type UserPropertySetRequest struct {
	AppID        string                   `json:"appid"`
	UserIDList   []string                 `json:"userid_list"`
	PropertyList []map[string]interface{} `json:"property_list"`
}

// UserPropertySetResponse defines the response structure for setting user properties.
type UserPropertySetResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// UserPropertySetTool is the tool for setting user properties.
type UserPropertySetTool struct {
	mcp.ToolBase
}

// Name returns the name of the tool.
func (t *UserPropertySetTool) Name() string {
	return "user_property_set"
}

// Description returns the description of the tool.
func (t *UserPropertySetTool) Description() string {
	return "Sets user properties. Corresponds to POST /user_property/set."
}

// Input returns an instance of UserPropertySetRequest.
func (t *UserPropertySetTool) Input() interface{} {
	return &UserPropertySetRequest{}
}

// Output returns an instance of UserPropertySetResponse.
func (t *UserPropertySetTool) Output() interface{} {
	return &UserPropertySetResponse{}
}

// Execute performs the action of setting user properties.
func (t *UserPropertySetTool) Execute(input interface{}) (interface{}, error) {
	req, ok := input.(*UserPropertySetRequest)
	if !ok {
		return nil, fmt.Errorf("invalid input type for UserPropertySetTool, expected *UserPropertySetRequest")
	}

	log.Printf("Received UserPropertySetRequest: %+v\n", req)

	// Mock implementation
	return &UserPropertySetResponse{Code: 0, Msg: "Successfully processed set request (mock)"}, nil
}

// UserPropertyDeleteRequest defines the request structure for deleting user properties.
type UserPropertyDeleteRequest struct {
	AppID      string   `json:"appid"`
	UserIDList []string `json:"userid_list"`
}

// UserPropertyDeleteResponse defines the response structure for deleting user properties.
type UserPropertyDeleteResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// UserPropertyDeleteTool is the tool for deleting user properties.
type UserPropertyDeleteTool struct {
	mcp.ToolBase
}

// Name returns the name of the tool.
func (t *UserPropertyDeleteTool) Name() string {
	return "user_property_delete"
}

// Description returns the description of the tool.
func (t *UserPropertyDeleteTool) Description() string {
	return "Deletes a user. Corresponds to POST /user_property/delete."
}

// Input returns an instance of UserPropertyDeleteRequest.
func (t *UserPropertyDeleteTool) Input() interface{} {
	return &UserPropertyDeleteRequest{}
}

// Output returns an instance of UserPropertyDeleteResponse.
func (t *UserPropertyDeleteTool) Output() interface{} {
	return &UserPropertyDeleteResponse{}
}

// Execute performs the action of deleting user properties.
func (t *UserPropertyDeleteTool) Execute(input interface{}) (interface{}, error) {
	req, ok := input.(*UserPropertyDeleteRequest)
	if !ok {
		return nil, fmt.Errorf("invalid input type for UserPropertyDeleteTool, expected *UserPropertyDeleteRequest")
	}

	log.Printf("Received UserPropertyDeleteRequest: %+v\n", req)

	// Mock implementation
	return &UserPropertyDeleteResponse{Code: 0, Msg: "Successfully processed delete request (mock)"}, nil
}
