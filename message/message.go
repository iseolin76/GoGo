package message

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/api"
	"github.com/iseolin76/GoGo/config"
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

	//응답
	if msg[1] == config.INVITE_COMMAND {
		s.ChannelMessageSend(m.ChannelID, config.INVITE_URL)
	}

	//밥 커맨드
	if msg[1] == config.MEAL_COMMAND || (len(msg) > 2 && msg[2] == config.MEAL_COMMAND) {
	var date string;
	if len(msg) > 2 {
		var message string
		if msg[1] == config.MEAL_COMMAND {
			message = msg[2]
		} else {
			message = msg[1]
		}
		date = util.ReturnDate(message)
	} else {
		now := time.Now();
		date = now.Format("20060102")	
		fmt.Println(int(time.Thursday - now.Weekday()))
	}
		s.ChannelMessageSendEmbed(m.ChannelID, api.NeisMealServiceDietInfo(date))
	}
}