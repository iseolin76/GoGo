package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

func init() { flag.Parse() }

func main() {
	dg, err := discordgo.New("Bot " + config.Token)
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
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

  prefix := string(strings.Fields(m.Content)[0]);
	var msg string = ""

	if len(strings.Fields(m.Content)) > 1 {
		msg = string(strings.Fields(m.Content)[1])
	}
	if prefix == config.AddCommand {
		s.ChannelMessageSend(m.ChannelID, "준비 중인 기능입니다.")
	}
	if prefix != config.Prefix  {
		return
	}
	if msg == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if msg == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
