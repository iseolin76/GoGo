package api

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/embed"
	"github.com/iseolin76/GoGo/types"
)

// 필요인자: 가져올 날짜 ex) 20201203

// 나이스에서 가져온 급식정보를 임베드로 만들어 리턴합니다.
func NeisMealServiceDietInfo(date string) *discordgo.MessageEmbed {
	//json을 MealServiceDietInfoType으로 Unmarshaling
	result := requestApi(mealServiceDietInfo(date))
	var data types.MealServiceDietInfoType
	json.Unmarshal(result, &data)

	//데이터가 비었을 경우 데이터가 없다는 임베드를 리턴합니다.
	if data.MealServiceDietInfo == nil {
		return embed.NoDataEmbed()
	}

	mealInfo := data.MealServiceDietInfo[1].Row

	var menus []string
	//한글만 필터링합니다.
	re := regexp.MustCompile("[가-힣]+")
	
	for _, menu := range mealInfo {
		menus = append(menus, menu.MmealScNm)
		menus = append(menus, strings.Join(re.FindAllString(menu.DdishNm, -1), "\n"))
		menus = append(menus, menu.CalInfo)
	}

	schul := mealInfo[0].SchulNm
	ymd:= mealInfo[0].MlsvYmd

	return embed.EmbedNeisMealServiceDietInfo(schul, ymd, menus)
}

// 필요인자: 가져올 날짜 ex) 20201203

// 나이스에서 가져온 시간표 정보를 임베드로 만들어 리턴합니다.
func NeisHisTimetable(date, grade, classNm string) *discordgo.MessageEmbed {
	//json을 HisTimetableType Unmarshaling
	result := requestApi(hisTimetable(date, grade, classNm))
	var data types.HisTimetableType
	json.Unmarshal(result, &data)


	//데이터가 비었을 경우 데이터가 없다는 임베드를 리턴합니다.
	if data.HisTimetable == nil {
		return embed.NoDataEmbed()
	}

	timeTable := data.HisTimetable[1].Row

	var times = make(map[string]string)

	for _, time := range timeTable {
		times[time.Perio] = time.ItrtCntnt
	}

	schul := timeTable[0].SchulNm
	ymd:= timeTable[0].AllTiYmd
	grade = timeTable[0].Grade
	classNm = timeTable[0].ClassNm

	return embed.EmbedNeisHisTimetable(schul, ymd, grade, classNm, times)
}