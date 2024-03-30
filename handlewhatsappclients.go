package main

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func (cfg *waConfig) HandleUsers(ctx context.Context, client *whatsmeow.Client, senderJID types.JID, username string, chatJID types.JID, senderNumber string, messageID []types.MessageID, message chan *waProto.Message) {

	//done to close connection when user has registered
	// Check if the user is already registered

	if _, err := cfg.DB.GetUserWhatsappNumber(ctx, senderNumber); err != nil {
		// this is a new user
		 cfg.HandleNewUser(ctx, client, senderJID, username, senderNumber)
		 fmt.Print("finished handling new user")
		 return
	}
	for {
		select {
		case message := <-messageChan:
			msg := message.GetConversation()
			fmt.Println(msg)
			// os.Exit(1)

		case <-ctx.Done():
			fmt.Println("context is cancelled from child")
			//send error message to user
			return

		}
	}
	
}
