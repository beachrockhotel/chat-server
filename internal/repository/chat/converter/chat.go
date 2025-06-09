package converter

import (
	"github.com/beachrockhotel/chat-server/internal/model"
	modelRepo "github.com/beachrockhotel/chat-server/internal/repository/chat/model"
)

func ToChatFromRepo(auth *modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID:        auth.ID,
		Info:      ToChatInfoFromRepo(auth.Info),
		CreatedAt: auth.CreatedAt,
		UpdatedAt: auth.UpdatedAt,
	}
}

func ToChatInfoFromRepo(info modelRepo.ChatInfo) model.ChatInfo {
	return model.ChatInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  info.Role,
	}
}
