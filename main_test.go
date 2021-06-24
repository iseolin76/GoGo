package main_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

var (
  envToken = os.Getenv("TOKEN")
)

func TestMain(m *testing.M) {
  if envToken == "" {
    os.Exit(m.Run());
  }

	_, err := discordgo.New("Bot " + envToken)
  if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

}

func TestMessageCreate(t *testing.T) {
  addCommandEx := "고커추"
  deleteCommandEx := "고커삭"

  if addCommandEx != config.ADD_COMMAND {
    t.Fatalf("고커추 접두사가 맞는데도 검사가 잘 되지 않음")
  }

  if deleteCommandEx != config.DELETE_COMMAND {
    t.Fatalf("고커삭 접두사가 맞는데도 검사가 잘 되지 않음")
  }
}