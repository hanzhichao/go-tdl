package go_tdl

import (
	"fmt"
)

type EnvId string

type Env struct {
	Id          EnvId                  `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Info        map[string]string      `json:"info"`
	Libraries   map[string]interface{} `json:"libraries"`
	Config      map[string][]byte      `json:"config"`
	Variables   map[string][]byte      `json:"variables"`
}

func (env *Env) GetVariable(key string) []byte {
	return env.Variables[key]
}

func (env *Env) SetVariable(key string, value []byte) {
	env.Variables[key] = value
}

func (env *Env) InvokeMethod(library, method string, args map[string]interface{}) []byte {
	libraryObj := env.Libraries[library]
	fmt.Println(libraryObj)
	//return libraryObj.Invoke(method, args)
	return nil
}

func (env *Env) RegisterLibrary(library string, libraryObj interface{}) {
	env.Libraries[library] = libraryObj
}
