package api

import (
	"fmt"
	"os"
)

var (
	key string = os.Getenv("NEIS_AUTH_KEY")
	defaultValue string = fmt.Sprintf("?KEY=%s&Type=json&pIndex=1&pSize=10&", key)
	defaultUrl string = "https://open.neis.go.kr/hub/"
)

func mealServiceDietInfo(date string) string {
	apiName := "mealServiceDietInfo"
	value := defaultUrl + apiName + defaultValue + "ATPT_OFCDC_SC_CODE=F10&SD_SCHUL_CODE=7380292&MLSV_YMD="+date
	return value
};
