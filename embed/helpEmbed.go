package embed

import (
	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

func HelpEmbed() *discordgo.MessageEmbed {
	helpEmbed := embed.NewEmbed().
	SetTitle("GoGo 도움말").
	SetDescription("'고'를 부르면 나타난다! GoGo.\n개발자 : ㅎㄹㅁㄹ#5146").
	AddField(config.INVITE_COMMAND, "GoGo를 초대할 수 있는 링크가 나타납니다.").
	AddField(config.MEAL_COMMAND, "거의 핵심인가?\n**광소마고의 급식**을 보여줍니다.").
	AddField(config.ADD_COMMAND, "사용자 커맨드를 추가할 수 있어요!\n사실 귀찮아 아직 안 만듬").
	AddField(config.DELETE_COMMAND, "사용자 커맨드를 삭제할 수 있어요!\n사실 귀찮아 아직 안 만듬").
	AddField(config.HELP_COMMAND, "GoGo의 명령어 목록과 설명을 보여줘요.\n너가 보고 있는 이거").
	SetFooter("농협) 356-1052-3698-23 예금주 이서린 \n<- 후원하면 진짜 열심히 일해볼게요..ㅎㅎ \n후원이라고 적고 보내줘야 함").
	SetColor(config.GO_COLOR).MessageEmbed

	return helpEmbed
}