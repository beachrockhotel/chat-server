package chat

import (
	"github.com/beachrockhotel/chat-server/internal/service"
	desc "github.com/beachrockhotel/chat-server/pkg/chat_v1"
	"sync"
)

type Chat struct {
	streams map[string]desc.ChatV1_ConnectChatServer
	m       sync.RWMutex
}

type Implementation struct {
	desc.UnimplementedChatV1Server

	chatService service.ChatService

	chats  map[string]*Chat
	mxChat sync.RWMutex

	channels  map[string]chan *desc.Message
	mxChannel sync.RWMutex
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
		chats:       make(map[string]*Chat),
		channels:    make(map[string]chan *desc.Message),
	}
}
