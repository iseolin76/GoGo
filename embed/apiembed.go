package embed

import (
	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

// 급식정보로 임베드를 만듭니다.
func EmbedNeisMealServiceDietInfo(schul string, ymd string, menus []string) *discordgo.MessageEmbed {
	var genericEmbed *discordgo.MessageEmbed

	//길이에 따라 다르게 임베드를 생성합니다.
	switch len(menus) {
	case (3) :
		genericEmbed = embed.NewEmbed().
		SetTitle(ymd[:4]+"년 "+ymd[5:6]+"월 "+ymd[7:]+"일").
		SetDescription(schul).
		AddField(menus[0], menus[1]+"\n"+menus[2]).
		InlineAllFields().
		Truncate().
		SetFooter(schul).
		SetColor(config.GO_COLOR).MessageEmbed
	case (6) :
		genericEmbed = embed.NewEmbed().
		SetTitle(ymd[:4]+"년 "+ymd[5:6]+"월 "+ymd[7:]+"일").
		SetDescription(schul).
		AddField(menus[0], menus[1]+"\n"+menus[2]).
		AddField(menus[3], menus[4]+"\n"+menus[5]).
		InlineAllFields().
		Truncate().
		SetFooter(schul).
		SetColor(config.GO_COLOR).MessageEmbed
	case (9) :
		genericEmbed = embed.NewEmbed().
		SetTitle(ymd[:4]+"년 "+ymd[5:6]+"월 "+ymd[7:]+"일").
		SetDescription(schul).
		AddField(menus[0], menus[1]+"\n"+menus[2]).
		AddField(menus[3], menus[4]+"\n"+menus[5]).
		AddField(menus[6], menus[7]+"\n"+menus[8]).
		InlineAllFields().
		Truncate().
		SetFooter(schul).
		SetColor(config.GO_COLOR).MessageEmbed
	default:
		return embed.NewEmbed().
		SetTitle("오잉 뭔가 이상한데?").
		AddField("개발자에게 문의해주세요. ㅎㄹㅁㄹ#5146", "급식 길이"+ string(rune(len(menus)))).
		SetColor(config.ERR_COLOR).MessageEmbed
	}
	
	return genericEmbed;
}