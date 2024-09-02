package go_tdl

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
	for _, step := range testcase.Setups {
		step.Run(env)
	}

	for _, step := range testcase.Steps {
		step.Run(env)
	}
	// TODO register result

	for _, step := range testcase.Teardowns {
		step.Run(env)
	}
}
