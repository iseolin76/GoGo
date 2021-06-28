package api

import (
	"fmt"
	"os"
)

var (
	key string = os.Getenv("NEIS_AUTH_KEY")
	defaultValue string = fmt.Sprintf("?KEY=%s&Type=json&pIndex=1&", key)
	defaultUrl string = "https://open.neis.go.kr/hub/"
)

func mealServiceDietInfo(date string) string {
	apiName := "mealServiceDietInfo"
	value := defaultUrl + apiName + defaultValue + "pSize=3&ATPT_OFCDC_SC_CODE=F10&SD_SCHUL_CODE=7380292&MLSV_YMD="+date
	return value
};

func hisTimetable(date, grade, classNm string) string {
	apiName := "hisTimetable"
	value := defaultUrl + apiName + defaultValue + "pSize=10&ATPT_OFCDC_SC_CODE=F10&SD_SCHUL_CODE=7380292&GRADE="+grade+"&CLASS_NM="+classNm+"&ALL_TI_YMD="+date
	return value
};

