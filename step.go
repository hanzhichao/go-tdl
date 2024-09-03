package go_tdl

import (
	"fmt"
	"strings"
)

type StepId string

type Step struct {
	Id          StepId                 `json:"id"`
	Description string                 `json:"description"`
	Method      string                 `json:"method"`
	Args        map[string]interface{} `json:"args"`
}

func (step *Step) Run(env *Env) []byte {

	fmt.Printf("调用方法: %s 参数: %s\n", step.Method, step.Args)

	array := strings.Split(step.Method, ".")
	library := array[0]
	method := array[1]
	args := step.Args
	result := env.InvokeMethod(library, method, args)
	fmt.Printf("调用结果: %s\n", string(result))
	return result
}
