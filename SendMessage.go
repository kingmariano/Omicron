package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/charlesozo/whisperbot/internal/database"
	"github.com/google/uuid"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

var (
	msg sync.Mutex
	wg  sync.WaitGroup // Mutex for concurrent map access
)

func FormatMessage(message string, username string) string {
	if username == "" {
		username = "dear"
	}
	formattedMessage := strings.Replace(message, "[User]", username, -1)
	return formattedMessage
}

func (cfg *waConfig) HandleNewUser(ctx context.Context, client *whatsmeow.Client, jid types.JID, username string, waNumber string) {

	// first time user  message

	for _, message := range NewUserMessages {
		wg.Add(1)
		msg := message
		go func() {
			defer wg.Done()
			cfg.SendMessage(ctx, client, jid, msg, username)
		}()
	}

	// Start a goroutine to create the user and subscription concurrently
	wg.Add(1)
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

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("This is first time user")

}

func (cfg *waConfig) SendMessage(ctx context.Context, client *whatsmeow.Client, jid types.JID, message, username string)  {
	msg.Lock()
	defer msg.Unlock()
	_, err := client.SendMessage(ctx, jid, &waProto.Message{
		Conversation: proto.String(FormatMessage(message, username)),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func SendCommandInstruction(ctx context.Context, client *whatsmeow.Client, jid types.JID, instruction string) {
	_, err := client.SendMessage(ctx, jid, &waProto.Message{
		Conversation: proto.String(instruction),
	})
	if err != nil {
		log.Fatal(err)
	}

}
func MarkMessageRead(client *whatsmeow.Client, chatJID types.JID, senderJID types.JID, messageID []types.MessageID) {
	err := client.MarkRead(messageID, time.Now(), chatJID, senderJID)
	if err != nil {
		log.Fatalf("couldnt mark receipt as read %v", err)
	}
}
