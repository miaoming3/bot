package handler

import (
	"context"
	"github.com/miaoming3/botgo/dto"
	"github.com/miaoming3/botgo/dto/message"
	"github.com/miaoming3/botgo/event"
	"github.com/miaoming3/botgo/openapi"
	"strings"
)

func AtMessageHandler(ctx context.Context, api openapi.OpenAPI) event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		content := message.ETLInput(data.Content)
		if strings.HasPrefix(content, "/") { //去掉/
			content = strings.TrimLeft(content, "/")
		}
		return nil
	}

}
