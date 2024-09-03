package tests

import (
	"github.com/hanzhichao/requests"
	"log"
)

type HttpLibrary struct {
	Session *requests.Session `json:"session"`
}

func NewHttpLibrary(baseUrl string) *HttpLibrary {
	config := requests.NewConfig().SetBaseUrl(baseUrl)
	session := requests.NewSession(config)
	return &HttpLibrary{Session: session}
}

func (http *HttpLibrary) Invoke(method string, args map[string]interface{}) []byte {
	if method == "Get" {
		return http.Get(args)
	} else if method == "Post" {
		return http.Post(args)
	} else {
		log.Fatal("No such method")
	}
	return nil
}

func (http *HttpLibrary) Get(args map[string]interface{}) []byte {
	url := args["url"].(string)
	headers := map[string]string{}
	if args["headers"] == nil {
		headers = nil
	} else {
		tmp := args["headers"].(map[string]interface{})
		for key, value := range tmp {
			headers[key] = value.(string)
		}
	}
	resp := http.Session.Get(url, headers)
	return resp.Content
}

func (http *HttpLibrary) Post(args map[string]interface{}) []byte {
	url := args["url"].(string)
	data := args["data"].(string)
	headers := map[string]string{}
	if args["headers"] == nil {
		headers = nil
	} else {
		tmp := args["headers"].(map[string]interface{})
		for key, value := range tmp {
			headers[key] = value.(string)
		}
	}
	resp := http.Session.Post(url, data, headers)
	return resp.Content
}
