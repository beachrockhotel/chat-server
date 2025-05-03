package main

import (
	"context"
	"log"
	"time"

	desc "github.com/beachrockhotel/chat-server/pkg/chat_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createResp, err := client.Create(ctx, &desc.CreateRequest{
		Usernames: []string{"user1", "user2"},
	})
	if err != nil {
		log.Fatalf("Failed to create chat: %v", err)
	}
	log.Printf(color.GreenString("Chat created with ID: %d", createResp.GetId()))

	_, err = client.Delete(ctx, &desc.DeleteRequest{
		Id: createResp.GetId(),
	})
	if err != nil {
		log.Fatalf("Failed to delete chat with ID: %d", createResp.GetId())
	}
	log.Println(color.GreenString("Chat deleted successfully"))

	_, err = client.SendMessage(ctx, &desc.SendMessageRequest{
		From:      "user1",
		Text:      "Hello, world!",
		CreatedAt: timestamppb.Now(),
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	log.Println(color.GreenString("Message sent successfully"))
}
