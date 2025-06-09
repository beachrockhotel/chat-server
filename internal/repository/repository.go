package repository

import (
	"context"

	"github.com/beachrockhotel/chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, info *model.ChatInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Chat, error)
}
