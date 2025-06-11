package chat

import (
	"context"
	"log"

	"github.com/beachrockhotel/chat-server/internal/converter"
	desc "github.com/beachrockhotel/chat-server/pkg/chat_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	chatObj, err := i.chatService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, title: %s, usernames: %s, created_at: %v, updated_at: %v\n", chatObj.ID, chatObj.Info.Title, chatObj.Info.Usernames, chatObj.CreatedAt, chatObj.UpdatedAt)

	return &desc.GetResponse{
		Chat: converter.ToChatFromService(chatObj),
	}, nil
}
