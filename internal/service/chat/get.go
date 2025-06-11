package chat

import (
	"context"

	"github.com/beachrockhotel/chat/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Chat, error) {
	chat, err := s.chatRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
