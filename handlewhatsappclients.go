package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
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

var (
	senderJIDChan    = make(chan types.JID)
	usernameChan     = make(chan string)
	messageIDChan    = make(chan []types.MessageID)
	chatJIDChan      = make(chan types.JID)
	senderNumberChan = make(chan string)
	messageChan      = make(chan *waProto.Message)
	users            = make(OmnicronUsers) // Map to store user IDs
	usersLock        sync.Mutex
	wg               sync.WaitGroup // Mutex for concurrent map access
)

func (cfg *waConfig) HandleUsers(ctx context.Context, client *whatsmeow.Client, senderJID types.JID, username string, chatJID types.JID, senderNumber string, messageID []types.MessageID, message *waProto.Message){
	usersLock.Lock()
	defer usersLock.Unlock()

	
	//   _  := <-messageChan

	wg.Add(1)
	go func(messageID []types.MessageID, now time.Time, chatJID types.JID, senderJID types.JID) {
		defer wg.Done()
		err := client.MarkRead(messageID, now, chatJID, senderJID)
		if err != nil {
			log.Fatalf("couldnt mark receipt as read %v", err)
		}
	}(messageID, time.Now(), chatJID, senderJID)
	wg.Wait()
	// Check if the user is already registered
	if _, ok := users[senderNumber]; !ok {
		// User is not registered, add them to the map and store in the database
		users[senderNumber] = true // Use an empty struct as the value for a set-like behavior

		// Use a channel to wait for the database operation to complete
		done := make(chan struct{})

		// Start a goroutine to store the user in the database concurrently
		go func(senderJID types.JID, username, senderNumber string) {
			defer close(done)
			// Store the user in the database
			// Your database logic here
			err := cfg.HandleNewUser(ctx, client, senderJID, username, senderNumber)
			if err != nil {
				log.Fatalf("error saving users to db %v", err)
			}
		}(senderJID, username, senderNumber)

		// Use a select statement to wait for either the database operation to complete
		// or the context to be canceled
		select {
		case <-done:
			// Database operation completed
			log.Println("done registering new users")
		case <-ctx.Done():
			// Context canceled, stop the database operation
			// You may want to roll back the user addition to the map here
			log.Println("operation cancelled")
			return
		}
	} else {
		// User is already registered, no need to store them again
		err := cfg.SendMessage(ctx, client, senderJID, username, senderNumber)
		if err != nil {
			fmt.Printf("unable to send other message %v", err)
			return
		}
		return 
	}

}
