package go_tdl

type TestSuiteId string

type TestCaseFilter struct {
	Names        []int `json:"names"`
	Priorities   []int `json:"priorities"`
	Tags         []int `json:"tags"`
	ExcludeTags  []int `json:"exclude_tags"`
	ExcludeNames []int `json:"exclude_names"`
}

type TestSuite struct {
	Id             TestSuiteId    `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Setups         []Step         `json:"setups"`
	Teardowns      []Step         `json:"teardowns"`
	SuiteSetups    []Step         `json:"suite_setups"`
	SuiteTeardowns []Step         `json:"suite_teardowns"`
	Tests          []TestCase     `json:"tests"`
	Filter         TestCaseFilter `json:"filter"`
}

func (suite *TestSuite) Run(env *Env) {
	for _, step := range suite.SuiteSetups {
		step.Run(env)
	}
	for _, test := range suite.Tests {
		test.Run(env)
	}
	for _, step := range suite.SuiteTeardowns {
		step.Run(env)
	}
}
