package embed

import (
	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

// 필요인자: schul: 학교 이름, ymd: 급식 일자, menus: 메뉴들

// 급식정보로 임베드를 만듭니다.
func EmbedNeisMealServiceDietInfo(schul, ymd string, menus []string) *discordgo.MessageEmbed {
	mealEmbed := embed.NewEmbed().
	SetTitle(ymd[:4]+"년 "+ymd[4:6]+"월 "+ymd[6:]+"일").
	SetDescription(schul)

	//길이에 따라 다르게 임베드를 생성합니다.
	for idx := range menus {
		if len(menus) < idx * 3 + 2 {
			break
		}
		mealEmbed = mealEmbed.
		AddField(menus[idx*3], menus[idx*3+1]+"\n"+menus[idx*3+2])
	}

	mealEmbed = mealEmbed.
	InlineAllFields().
	Truncate().
	SetFooter(schul).
	SetColor(config.GO_COLOR)
	
	return mealEmbed.MessageEmbed;
}

