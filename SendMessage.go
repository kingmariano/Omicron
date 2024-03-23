package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/charlesozo/whisperbot/internal/database"
	"github.com/google/uuid"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)


func (cfg *waConfig) HandleNewUser(ctx context.Context, client *whatsmeow.Client, jid types.JID, username string, waNumber string)  {
	var (
		mu sync.Mutex
		wg        sync.WaitGroup // Mutex for concurrent map access
	) 
	mu.Lock()
	defer mu.Unlock()
	// first time user introducing message
		wg.Add(2)
		go func () {
			defer wg.Done()
			for _, message := range NewUserMessages {
			_, err := client.SendMessage(ctx, jid, &waProto.Message{
				Conversation: proto.String(FormatMessage(message, username)),
			})
			if err != nil {
				log.Fatal("error sending message")
				
			}
		}
	}()
		
	
		// Start a goroutine to create the user and subscription concurrently
	go func() {
		defer wg.Done()
		// Create the user
		newUser, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
			ID:             uuid.New(),
			CreatedAt:      time.Now().UTC(),
			WhatsappNumber: waNumber,
			DisplayName:    username,
		})
		if err != nil {
			log.Fatal(err)
		}

		// Create the user subscription
		err = cfg.DB.CreateUserSubscription(ctx, database.CreateUserSubscriptionParams{
			SubscriptionID: uuid.New(),
			Userid:         uuid.NullUUID{UUID: newUser.ID, Valid: true},
			ExpiryDate:     time.Now().AddDate(0, 0, 7),
		})
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()
	
	fmt.Println("This is first time user")
}

func (cfg *waConfig) SendMessage(ctx context.Context, client *whatsmeow.Client, jid types.JID, username string, waNumber string) error{
		_, err := client.SendMessage(ctx, jid, &waProto.Message{
		Conversation: proto.String(FormatMessage(MessageUser, username)),
	})
	if err != nil {
		return err
	}
	return nil
}
