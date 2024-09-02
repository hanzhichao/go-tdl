package tests

import (
	"encoding/json"
	"fmt"
	go_tdl "go-tdl"
	"testing"
)

func TestTdlHttpSteps(t *testing.T) {
	data := `{
	  "name": "test_api_demo",
	  "description": "test description",
	  "priority": 1,
	  "tags": ["http", "api-test"],
	  "setups": [],
	  "teardowns": [],
	  "steps": [
		{"method": "Http.Get", "args": {"url": "/get"}},
		{"method": "Http.Post", "args": {"url": "/post", "data":  "name=Kevin&age=12"}}
	  ]
	}`
	testCase := go_tdl.TestCase{}
	err := json.Unmarshal([]byte(data), &testCase)
	if err != nil {
		fmt.Println(err)
	}

}
