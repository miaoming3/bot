package main

import (
	"context"
	"fmt"
	"github.com/miaoming3/bot/global"
	"github.com/miaoming3/bot/initialization"
	"github.com/miaoming3/botgo"
	"github.com/miaoming3/botgo/dto"
	"github.com/miaoming3/botgo/dto/message"
	"github.com/miaoming3/botgo/event"
	"github.com/miaoming3/botgo/token"
	"github.com/miaoming3/botgo/websocket"
	"log"
	"strings"
)

func init() {
	initialization.LoadConfig("config.yaml")
	initialization.InitContent()
}

func main() {
	botToken := token.BotToken(global.BotConfigure.AppId, global.BotConfigure.Token)
	api := botgo.NewSandboxOpenAPI(botToken)
	ctx := context.Background()
	ws, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalf("webSocket error: %v", err)
	}

	var group event.GroupAtMessageEventHandler = func(event *dto.WSPayload, data *dto.WSGroupATMessageData) error {
		fmt.Println(data.GroupId)
		return nil
	}
	var at event.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		fmt.Println(data.ChannelID, data.GuildID)
		return nil
	}

	var messages event.MessageEventHandler = func(event *dto.WSPayload, data *dto.WSMessageData) error {
		input := strings.ToLower(message.ETLInput(data.Content))
		Ctype := message.ParseCommand(input).Cmd
		fmt.Println(Ctype)
		return nil
	}
	intent := websocket.RegisterHandlers(group, at, messages)
	botgo.NewSessionManager().Start(ws, botToken, &intent)
}
