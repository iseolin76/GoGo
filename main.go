package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

var (
	Token = os.Getenv("TOKEN")
)

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//봇이 보낸 메시지인지 판별
	if m.Author.ID == s.State.User.ID {
		return
	}

	//명령어 검사
	msg := strings.Split(m.Content, " ")
	if msg[0] == config.ADD_COMMAND || msg[0] == config.DELETE_COMMAND {
		s.ChannelMessageSend(m.ChannelID, "준비 중인 기능입니다.")
	}
	if msg[0] != config.PREFIX  {
		return
	}

	//응답
	fmt.Println(msg)
	if msg[1] == "초대코드" {
		s.ChannelMessageSend(m.ChannelID, config.INVITE_URL)
	}
	if msg[1] == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if msg[1] == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
