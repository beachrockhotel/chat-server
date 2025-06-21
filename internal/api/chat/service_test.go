package chat_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/beachrockhotel/chat-server/internal/api/chat"
	"github.com/beachrockhotel/chat-server/internal/model"
	desc "github.com/beachrockhotel/chat-server/pkg/chat_v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// Мок сервиса
type mockChatService struct {
	mock.Mock
}

func (m *mockChatService) Create(ctx context.Context, info *model.ChatInfo) (int64, error) {
	args := m.Called(ctx, info)
	return int64(args.Int(0)), args.Error(1)
}

func (m *mockChatService) Get(ctx context.Context, id int64) (*model.Chat, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Chat), args.Error(1)
}

func dialer(srv desc.ChatV1Server) func(context.Context, string) (net.Conn, error) {
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	desc.RegisterChatV1Server(s, srv)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	return func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}
}

func TestCreate(t *testing.T) {
	mockSvc := &mockChatService{}
	impl := chat.NewImplementation(mockSvc)

	mockSvc.On("Create", mock.Anything, mock.Anything).Return(123, nil)

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(impl)))
	require.NoError(t, err)
	client := desc.NewChatV1Client(conn)

	resp, err := client.Create(context.Background(), &desc.CreateRequest{
		Info: &desc.ChatInfo{Title: "Test", Usernames: "User1,User2"},
	})
	require.NoError(t, err)
	assert.Equal(t, int64(123), resp.GetId())
}

func TestGet(t *testing.T) {
	mockSvc := &mockChatService{}
	impl := chat.NewImplementation(mockSvc)

	mockSvc.On("Get", mock.Anything, int64(1)).Return(&model.Chat{ID: 1, Info: model.ChatInfo{Title: "Test"}}, nil)

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(dialer(impl)))
	require.NoError(t, err)
	client := desc.NewChatV1Client(conn)

	resp, err := client.Get(context.Background(), &desc.GetRequest{Id: 1})
	require.NoError(t, err)
	assert.Equal(t, int64(1), resp.GetChat().GetId())
	assert.Equal(t, "Test", resp.GetChat().GetInfo().GetTitle())
}
