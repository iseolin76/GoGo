package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

func requestApi(requestUrl string) ([]byte, error) {
	resp, err := http.Get(requestUrl)
	if err != nil {	
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}