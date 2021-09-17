package embed

import (
	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

//GoGo봇 도움말 임베드를 리턴합니다.
func HelpEmbed() *discordgo.MessageEmbed {
	helpEmbed := embed.NewEmbed().
	SetTitle("GoGo 도움말").
	SetDescription("'고'를 부르면 나타난다! GoGo.\n접두사 : 'ㄱ'또는 '고'\n개발자 : ㅎㄹㅁㄹ#5146").
	AddField(config.INVITE_COMMAND, "GoGo를 초대할 수 있는 링크가 나타납니다.").
	AddField(config.MEAL_COMMAND, "거의 핵심인가?\n광소마고의 급식을 보여줍니다.\n초성도 돼요. ㄱㅂ, ㄱㅂ 내일, ㄱ밥 ...").
	AddField(config.ADD_COMMAND, "사용자 커맨드를 추가할 수 있어요!\n사실 귀찮아 아직 안 만듬").
	AddField(config.DELETE_COMMAND, "사용자 커맨드를 삭제할 수 있어요!\n사실 귀찮아 아직 안 만듬").
	AddField(config.HELP_COMMAND, "GoGo의 명령어 목록과 설명을 보여줘요.\n너가 보고 있는 이거").
	AddField(config.TIMETABLE_COMMAND, "광소마고의 시간표를 보여줘요.\n형식 고[일자(생략가능)] 시간표 [학년] [반]\nex) 고다음주 월요일 시간표 2 1").
	AddField("오류 및 건의사항은", "깃허브 바로 가기 -> https://github.com/iseolin76/GoGo\n~~별을 홀수 번 눌러주기~~").
	SetFooter("농협) 356-1052-3698-23 예금주 이서린 \n<- 후원하면 진짜 열심히 일해볼게요..ㅎㅎ \n후원이라고 적고 보내줘야 함").
	SetColor(config.GO_COLOR).MessageEmbed

	return helpEmbed
}

//데이터가 없다는 임베드를 리턴합니다.
func NoDataEmbed() *discordgo.MessageEmbed {
	return embed.NewEmbed().
	SetTitle("데이터가 없어요!").
	SetDescription("이 날은 쉬는 날인가??").
	SetColor(config.ERR_COLOR).MessageEmbed
}

//데이터가 없다는 임베드를 리턴합니다.
func ErrorEmbed() *discordgo.MessageEmbed {
	return embed.NewEmbed().
	SetTitle("에러가 발생하였습니다!").
	SetDescription("개발자에게 문의하여주세요. ㅎㄹㅁㄹ#5146").
	SetColor(config.ERR_COLOR).MessageEmbed
}