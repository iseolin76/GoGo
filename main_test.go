package main_test

import (
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/iseolin76/GoGo/config"
)

var (
  envToken, ok = os.LookupEnv("TOKEN")
)

func TestMain(m *testing.M) {
  if !ok {
    os.Exit(m.Run());
  }

	dg, err := discordgo.New("Bot " + envToken)
  if err != nil {
    os.Exit(m.Run());
	}

  err = dg.Open()
  if err != nil {
    os.Exit(m.Run());
  }
}

func TestMessageCreate(t *testing.T) {
  addCommandEx := config.ADD_COMMAND
  deleteCommandEx := config.DELETE_COMMAND
  inviteCommandEx := config.INVITE_COMMAND
  mlsvCommandEx := config.MEAL_COMMAND

  if addCommandEx != config.ADD_COMMAND {
    t.Fatalf("고커추 접두사가 맞는데도 검사가 잘 되지 않음")
  }

  if deleteCommandEx != config.DELETE_COMMAND {
    t.Fatalf("고커삭 접두사가 맞는데도 검사가 잘 되지 않음")
  }

  if inviteCommandEx != config.INVITE_COMMAND {
    t.Fatalf("초대코드 명령어가 맞는데도 검사가 잘 되지 않음")
  }

  if mlsvCommandEx != config.MEAL_COMMAND {
    t.Fatalf("밥 명령어가 맞는데도 검사가 잘 되지 않음")
  }
}