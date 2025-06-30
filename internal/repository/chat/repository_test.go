package chat_test

import (
	"context"
	"github.com/joho/godotenv"
	"os"
	"testing"

	"github.com/beachrockhotel/chat-server/internal/client/db/pg"
	"github.com/beachrockhotel/chat-server/internal/model"
	chatRepo "github.com/beachrockhotel/chat-server/internal/repository/chat"
	"github.com/stretchr/testify/require"
)

func TestCreateAndGet(t *testing.T) {
	_ = godotenv.Load(".env")

	dsn := os.Getenv("PG_DSN")
	ctx := context.Background()

	client, err := pg.New(ctx, dsn)
	require.NoError(t, err)
	defer client.Close()

	repo := chatRepo.NewRepository(client)

	id, err := repo.Create(ctx, &model.ChatInfo{
		Title:     "Test Chat",
		Usernames: "user1,user2",
	})
	require.NoError(t, err)

	chatObj, err := repo.Get(ctx, id)
	require.NoError(t, err)
	require.Equal(t, "Test Chat", chatObj.Info.Title)
}
