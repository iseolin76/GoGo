package message

import (
	"strings"
	"time"

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

	if msg[0] == config.ADD_COMMAND || msg[0] == config.DELETE_COMMAND {
		s.ChannelMessageSend(m.ChannelID, "준비 중인 기능입니다.")
	}

	if msg[0] != config.PREFIX || len(msg) < 2 {
		return
	}

	//고 응답
	if msg[1] == config.INVITE_COMMAND {
		s.ChannelMessageSend(m.ChannelID, config.INVITE_URL)
	}

	//밥 명령어
	if msg[1] == config.MEAL_COMMAND || 
	(len(msg) > 2 && msg[2] == config.MEAL_COMMAND) || 
	(len(msg) > 3 && msg[3] == config.MEAL_COMMAND) {
		now := time.Now();
		var date string;

		if len(msg) > 2 {
			var message string

			//명령어 판별
			if msg[1] == config.MEAL_COMMAND {
				if len(msg) == 4 {
					message = msg[2] + " " + msg[3]
				} else {
				message = msg[2]
				}
			} else if msg[2] == config.MEAL_COMMAND {
				message = msg[1]
			} else if len(msg) > 3 && msg[3] == config.MEAL_COMMAND {
				message = msg[1] + " " + msg[2]
			}
			now = util.CheckDate(message, now)
		}
		date = now.Format("20060102")	

		s.ChannelMessageSendEmbed(m.ChannelID, api.NeisMealServiceDietInfo(date))
	}

	//도움말 명령어
	if msg[1] == config.HELP_COMMAND {
		s.ChannelMessageSendEmbed(m.ChannelID, embed.HelpEmbed())
	}

	//시간표 명령어
	if msg[1] == config.TIMETABLE_COMMAND || 
	(len(msg) > 2 && msg[2] == config.TIMETABLE_COMMAND) ||
	(len(msg) > 3 && msg[3] == config.TIMETABLE_COMMAND) {
		if msg[len(msg)-1] != config.TIMETABLE_COMMAND && len(msg) > 3 {
			now := time.Now();
			var date string;
			var message string
			var grade, classNm string

			//명령어 판별
			if len(msg) > 5 {
				if len(msg) == 5 && msg[2] == config.TIMETABLE_COMMAND {
					message = msg[1]
					grade = msg[3]
					classNm = msg[4]
					now = util.CheckDate(message, now)
				} else if len(msg) == 6 && msg[3] == config.TIMETABLE_COMMAND {
					message = msg[1]+" "+msg[2]
					grade = msg[4]
					classNm = msg[5]
					now = util.CheckDate(message, now)
				}
			}

			date = now.Format("20060102")
			s.ChannelMessageSendEmbed(m.ChannelID, api.NeisHisTimetable(date, grade, classNm))
		} else {
			s.ChannelMessageSend(m.ChannelID, `형식은 *"고 [일자] 시간표 [학년] [반]"* 이에요!`+"\nex) 고 내일 시간표 2 1, 고 시간표 1 1...")
		}
	}
}