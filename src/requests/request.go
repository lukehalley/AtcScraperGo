package requests

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeGetRequestRAW(RequestURL string) string {

	Response, Error := http.Get(RequestURL)

	if Error != nil {
		log.Fatalln(Error)
	}

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	return string(bytes.Replace(ResponseBody, []byte("\r"), []byte("\r\n"), -1))

}

func MakeGetRequestJSON(RequestURL string) []byte {

	Response, Error := http.Get(RequestURL)

	if Error != nil {
		log.Fatalln(Error)
	}

	ResponseBody, _ := ioutil.ReadAll(Response.Body)

	return ResponseBody

}