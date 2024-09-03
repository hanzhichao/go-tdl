package tests

import (
	"encoding/json"
	"fmt"
	"go-tdl"
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
		{"method": "Http.Post", "args": {"url": "/post", "data":  "name=Kevin&age=12", "headers": {"Content-Type": "application/x-www-urlencoded"}}}
	  ]
	}`
	testCase := go_tdl.TestCase{}
	err := json.Unmarshal([]byte(data), &testCase)
	if err != nil {
		fmt.Println(err)
	}

	httpLibrary := NewHttpLibrary("http://httpbin.org")
	env := go_tdl.NewEnv(map[string]go_tdl.Library{"Http": httpLibrary}, nil)
	testCase.Run(env)

}
