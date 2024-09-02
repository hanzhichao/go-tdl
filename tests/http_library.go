package tests

import (
	"encoding/json"
	"github.com/hanzhichao/requests"
	"log"
)

type HttpLibrary struct {
	Session *requests.Session `json:"session"`
}

func (http *HttpLibrary) Init(baseUrl string) {
	config := requests.NewConfig().SetBaseUrl(baseUrl)
	http.Session = requests.NewSession(config)
}

func (http *HttpLibrary) Invoke(method string, args map[string][]byte) []byte {
	if method == "Get" {
		url := string(args["url"])
		headers := map[string]string{}
		err := json.Unmarshal(args["headers"], &headers)
		if err != nil {
			log.Fatal(err)
		}
		resp := http.Get(url, headers)
		return resp.Content
	} else if method == "Post" {
		url := string(args["url"])
		data := string(args["data"])
		headers := map[string]string{}
		err := json.Unmarshal(args["headers"], &headers)
		if err != nil {
			log.Fatal(err)
		}
		resp := http.Post(url, data, headers)
		return resp.Content
	} else {
		log.Fatal("No such method")
	}
	return nil
}

func (http *HttpLibrary) Get(url string, headers map[string]string) *requests.Response {
	return http.Session.Get(url, headers)
}

func (http *HttpLibrary) Post(url, data string, headers map[string]string) *requests.Response {
	return http.Session.Post(url, data, headers)
}
