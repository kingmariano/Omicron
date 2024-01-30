package main

import (
	"context"
	"fmt"
	"github.com/charlesozo/whisperbot/cron"
	"github.com/charlesozo/whisperbot/internal/database"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"time"
)

func (cfg *waConfig) SendMessage(client *whatsmeow.Client, jid types.JID, username string, waNumber string) {
	fmt.Println("sending normal messages")
	messageOne := "Hey [User], how are you?"
	messageTwo := "Hey [User] nice to meet you again"
	_, err := cfg.DB.GetUnRegUser(context.Background(), waNumber)
	if err != nil {
		fmt.Println("This is first time user")
		client.SendMessage(context.Background(), jid, &waProto.Message{
			Conversation: proto.String(cron.FormatMessage(messageOne, username)),
		})
		_, err = cfg.DB.CreateUnRegUser(context.Background(), database.CreateUnRegUserParams{
			WhatsappNumber: waNumber,
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      time.Now().UTC(),
			DisplayName:    username,
		})
		if err != nil {
			fmt.Printf("error registering User %v", err)
			return
		}
		return
	}
	fmt.Println("user already registered")
	client.SendMessage(context.Background(), jid, &waProto.Message{
		Conversation: proto.String(cron.FormatMessage(messageTwo, username)),
	})

}
