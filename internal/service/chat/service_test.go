package chat_test

import (
	"context"
	"testing"

	"github.com/beachrockhotel/chat-server/internal/client/db"
	"github.com/beachrockhotel/chat-server/internal/model"
	"github.com/beachrockhotel/chat-server/internal/service/chat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мок репозитория
type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Create(ctx context.Context, info *model.ChatInfo) (int64, error) {
	args := m.Called(ctx, info)
	return int64(args.Int(0)), args.Error(1)
}

func (m *mockRepo) Get(ctx context.Context, id int64) (*model.Chat, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Chat), args.Error(1)
}

// Мок TxManager
type mockTx struct {
	mock.Mock
}

func (m *mockTx) ReadCommitted(ctx context.Context, f db.Handler) error {
	return f(ctx)
}

func TestCreate(t *testing.T) {
	repo := &mockRepo{}
	tx := &mockTx{}

	s := chat.NewService(repo, tx)

	repo.On("Create", mock.Anything, mock.Anything).Return(1, nil)
	repo.On("Get", mock.Anything, int64(1)).Return(&model.Chat{ID: 1}, nil)

	id, err := s.Create(context.Background(), &model.ChatInfo{Title: "Test"})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}

func TestGet(t *testing.T) {
	repo := &mockRepo{}
	tx := &mockTx{}
	s := chat.NewService(repo, tx)

	repo.On("Get", mock.Anything, int64(1)).Return(&model.Chat{ID: 1}, nil)

	chatObj, err := s.Get(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), chatObj.ID)
}
