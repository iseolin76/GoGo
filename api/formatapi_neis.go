package api

import (
	"regexp"
	"strings"

	outEmbed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
	"github.com/iseolin76/GoGo/embed"
)

// 필요인자: 가져올 날짜 ex) 20201203
// 나이스에서 가져온 급식정보를 discord 임베드로 만들어 리턴합니다.
func NeisMealServiceDietInfo(date string) *discordgo.MessageEmbed {
	data := requestApi(mealServiceDietInfo(date))

	//데이터가 비었을 경우 급식이 없다는 임베드를 리턴합니다.
	if data.MealServiceDietInfo == nil {
		return outEmbed.NewEmbed().
		SetTitle("급식이 없어요!").
		SetDescription("이 날은 쉬는 날인가??").
		SetColor(config.GO_COLOR).MessageEmbed
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
