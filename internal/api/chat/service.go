package chat

import (
	"github.com/beachrockhotel/chat-server/internal/service"
	desc "github.com/beachrockhotel/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
