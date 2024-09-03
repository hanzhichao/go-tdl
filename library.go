package go_tdl

type Library interface {
	Invoke(method string, args map[string]interface{}) []byte
}
