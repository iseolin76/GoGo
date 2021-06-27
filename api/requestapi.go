package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/iseolin76/GoGo/types"
)



func requestApi(requestUrl string) types.MealServiceDietInfoType {
	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	result := types.MealServiceDietInfoType{}

	json.Unmarshal(data, &result)

	return result
}