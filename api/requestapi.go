package api

import (
	"io/ioutil"
	"net/http"
)

func requestApi(requestUrl string) []byte {
	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return data
}