package go_tdl

import "strings"

type StepId string

type Step struct {
	Id          StepId                 `json:"id"`
	Description string                 `json:"description"`
	Method      string                 `json:"method"`
	Args        map[string]interface{} `json:"args"`
}

func (step *Step) Run(env *Env) []byte {
	array := strings.Split(step.Method, ".")
	library := array[0]
	method := array[1]
	args := step.Args
	return env.InvokeMethod(library, method, args)
}
