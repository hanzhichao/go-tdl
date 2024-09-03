package go_tdl

type EnvId string

type Env struct {
	Id          EnvId                  `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Info        map[string]string      `json:"info"`
	Libraries   map[string]Library     `json:"libraries"`
	Config      map[string][]byte      `json:"config"`
	Variables   map[string]interface{} `json:"variables"`
}

func NewEnv(libraries map[string]Library, variables map[string]interface{}) *Env {
	if variables == nil {
		variables = make(map[string]interface{})
	}
	return &Env{Libraries: libraries, Variables: variables}
}

func (env *Env) GetVariable(key string) interface{} {
	return env.Variables[key]
}

func (env *Env) SetVariable(key string, value interface{}) {
	env.Variables[key] = value
}

func (env *Env) InvokeMethod(library, method string, args map[string]interface{}) []byte {
	libraryObj := env.Libraries[library]
	return libraryObj.Invoke(method, args)
}

func (env *Env) RegisterLibrary(library string, libraryObj Library) {
	env.Libraries[library] = libraryObj
}
