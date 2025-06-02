package chat

import (
	"context"
	"log"

	"github.com/beachrockhotel/chat-server/internal/converter"
	desc "github.com/beachrockhotel/chat-server/pkg/chat_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	authObj, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %s, created_at: %v, updated_at: %v\n", authObj.ID, authObj.Info.Name, authObj.Info.Email, authObj.Info.Role, authObj.CreatedAt, authObj.UpdatedAt)

	return &desc.GetResponse{
		Chat: converter.ToChatFromService(authObj),
	}, nil
}
