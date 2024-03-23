package main

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"log"
	"sync"
	"time"
)

type WaClient struct {
	SenderJID    types.JID
	Username     string
	MessageJID   []types.MessageID
	ChatJID      types.JID
	SenderNumber string
	Message      *waProto.Message
}
type OmnicronUsers map[string]bool

func (cfg *waConfig) HandleUsers(ctx context.Context, client *whatsmeow.Client, senderJID types.JID, username string, chatJID types.JID, senderNumber string, messageID []types.MessageID, done chan struct{}) {
	// Use ctx for cancellation if needed
	var (
		usersLock sync.Mutex
		wg        sync.WaitGroup // Mutex for concurrent map access
	)

	var users = make(OmnicronUsers)

	defer close(done)
	defer usersLock.Unlock()


	// Mark message as read
	wg.Add(1)
	go func(messageID []types.MessageID, now time.Time, chatJID types.JID, senderJID types.JID) {
		defer wg.Done()
		err := client.MarkRead(messageID, now, chatJID, senderJID)
		if err != nil {
			log.Fatalf("couldnt mark receipt as read %v", err)
		}
	}(messageID, time.Now(), chatJID, senderJID)
 
	usersLock.Lock()
	// Check if the user is already registered
	_, err :=   cfg.DB.GetUserWhatsappNumber(ctx, senderNumber)
	if err != nil {
		wg.Add(1)
		// Start a goroutine to store the user in the database concurrently
		go func(senderJID types.JID, username, senderNumber string) {
			defer wg.Done()
			// Store the user in the database
			cfg.HandleNewUser(ctx, client, senderJID, username, senderNumber)

		}(senderJID, username, senderNumber)
	}else {
		fmt.Println(users)
		err := cfg.SendMessage(ctx, client, senderJID, username, senderNumber)
		if err != nil {
			fmt.Printf("unable to send other message %v", err)
			return
		}
	}

	

	// Wait for all goroutines to finish
	wg.Wait()

	select {
	case <-ctx.Done():
		// Context is canceled, perform cleanup or return early
		log.Println("Context canceled, cleaning up...")
		return
	default:
		// Context is not canceled, continue processing
	}
}
