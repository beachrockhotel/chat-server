package chat

import (
	"github.com/beachrockhotel/chat/internal/client/db"
	"github.com/beachrockhotel/chat/internal/repository"
	"github.com/beachrockhotel/chat/internal/service"
)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
