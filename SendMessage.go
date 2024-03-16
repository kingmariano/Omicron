package main

import (
	"context"
	"fmt"
	"github.com/charlesozo/whisperbot/internal/database"
	"github.com/google/uuid"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"time"
)

func (cfg *waConfig) HandleNewUser(ctx context.Context, client *whatsmeow.Client, jid types.JID, username string, waNumber string) error {
	// first time user introducing message
		fmt.Println("This is first time user")
		
		_, err := client.SendMessage(ctx, jid, &waProto.Message{
			Conversation: proto.String(FormatMessage(NewIntroductoryMessage, username)),
		})
		if err != nil {
			fmt.Println("error sending message")
			return err
		}
		newUser, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
			ID:             uuid.New(),
			CreatedAt:      time.Now().UTC(),
			WhatsappNumber: waNumber,
			DisplayName:    username,
		})
		if err != nil {
			fmt.Println("error creating user")
			return err
		}
		err = cfg.DB.CreateUserSubscription(ctx, database.CreateUserSubscriptionParams{
			SubscriptionID: uuid.New(),
			Userid:         uuid.NullUUID{UUID: newUser.ID, Valid: true},
			ExpiryDate:     time.Now().AddDate(0, 0, 7),
		})
		if err != nil {
			fmt.Printf("error creating subcription %v", err)
			return err
		}
	
	return nil
}
func (cfg *waConfig) SendMessage(ctx context.Context, client *whatsmeow.Client, jid types.JID, username string, waNumber string) error{
	fmt.Println("This is second time user")
		_, err := client.SendMessage(ctx, jid, &waProto.Message{
		Conversation: proto.String(FormatMessage(MessageUser, username)),
	})
	if err != nil {
		
		return err
	}
	return nil
}
