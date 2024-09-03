package go_tdl

import "fmt"

type TestCaseId string

type TestCase struct {
	Id          TestCaseId `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"`
	Tags        []string   `json:"tags"`
	Setups      []Step     `json:"setups"`
	Teardowns   []Step     `json:"teardowns"`
	Steps       []Step     `json:"steps"`
}

func (testcase *TestCase) Run(env *Env) {
	fmt.Printf("执行测试用例: %s\n", testcase.Name)
	for _, step := range testcase.Setups {
		step.Run(env)
	}

	for index, step := range testcase.Steps {
		fmt.Printf("执行: 步骤%d\n", index+1)
		step.Run(env)

	}
	// TODO register result

	for _, step := range testcase.Teardowns {
		step.Run(env)
	}
}
