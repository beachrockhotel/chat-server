package chat

import (
	"github.com/beachrockhotel/chat-server/internal/service"
	desc "github.com/beachrockhotel/chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.ChatService
}

func NewImplementation(authService service.ChatService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
