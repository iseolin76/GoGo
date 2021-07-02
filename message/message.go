package message

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/api"
	"github.com/iseolin76/GoGo/config"
	"github.com/iseolin76/GoGo/embed"
	"github.com/iseolin76/GoGo/util"
)


func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//봇이 보낸 메시지인지 판별
	if m.Author.ID == s.State.User.ID {
		return
	}

	//명령어 검사
	msg := strings.Split(m.Content, " ")

	if(len(msg[0]) > 1) {
		msg1 := strings.Split(m.Content, "");
		if msg1[0] != config.PREFIX && msg1[0] != config.PREFIX1 {
			return
		} else {
			msg = strings.Split(strings.Join(msg1[1:], ""), " ")
		}
	}

	if msg[0] == config.ADD_COMMAND || msg[0] == config.DELETE_COMMAND {
		s.ChannelMessageSend(m.ChannelID, "준비 중인 기능입니다.")
	}

	//고 응답
	if msg[0] == config.INVITE_COMMAND {
		s.ChannelMessageSend(m.ChannelID, config.INVITE_URL)
	}

	//밥 명령어
	if msg[0] == config.MEAL_COMMAND ||
	msg[0] == "ㅂ" ||
	(len(msg) > 1 && msg[1] == config.MEAL_COMMAND) || 
	(len(msg) > 2 && msg[2] == config.MEAL_COMMAND) {
		var date string;
		now := util.ReturnDate();

		if len(msg) > 1 {
			var message string

			//명령어 판별
			if msg[0] == config.MEAL_COMMAND || msg[0] == "ㅂ" {
				if len(msg) == 3 {
					message = msg[1] + " " + msg[2]
				} else {
				message = msg[1]
				}
			} else if msg[1] == config.MEAL_COMMAND {
				message = msg[0]
			} else if len(msg) > 2 && msg[2] == config.MEAL_COMMAND {
				message = msg[0] + " " + msg[1]
			}
			now = util.CheckDate(message, now)
		}
		date = now.Format("20060102")	

		s.ChannelMessageSendEmbed(m.ChannelID, api.NeisMealServiceDietInfo(date))
	}

	//도움말 명령어
	if msg[0] == config.HELP_COMMAND {
		s.ChannelMessageSendEmbed(m.ChannelID, embed.HelpEmbed())
	}

	//시간표 명령어
	if msg[0] == config.TIMETABLE_COMMAND || 
	(len(msg) > 1 && msg[1] == config.TIMETABLE_COMMAND) ||
	(len(msg) > 2 && msg[2] == config.TIMETABLE_COMMAND) {
		if msg[len(msg)-1] != config.TIMETABLE_COMMAND && len(msg) > 2 {
			now := util.ReturnDate();
			var date string;
			var message string
			var grade, classNm string

			//명령어 판별
			if len(msg) > 3 {
				if len(msg) == 4 && msg[1] == config.TIMETABLE_COMMAND {
					message = msg[0]
					grade = msg[2]
					classNm = msg[3]
					now = util.CheckDate(message, now)
				} else if len(msg) == 5 && msg[2] == config.TIMETABLE_COMMAND {
					message = msg[0]+" "+msg[1]
					grade = msg[3]
					classNm = msg[4]
					now = util.CheckDate(message, now)
				}
			}

			date = now.Format("20060102")
			s.ChannelMessageSendEmbed(m.ChannelID, api.NeisHisTimetable(date, grade, classNm))
		} else {
			s.ChannelMessageSend(m.ChannelID, `형식은 *"고[일자] 시간표 [학년] [반]"* 이에요!`+"\nex) 고내일 시간표 2 1, 고시간표 1 1...")
		}
	}
}