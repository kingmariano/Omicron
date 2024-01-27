package main

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"strings"
)

func (cfg *waConfig) SendMessage(client *whatsmeow.Client, jid types.JID, username string) {
	initialMessage := "Hey [User], how are you?"
	fmt.Println("sending normal messages")
	client.SendMessage(context.Background(), jid, &waProto.Message{
		Conversation: proto.String(formatMessage(initialMessage, username)),
	})

}

func (cfg *waConfig) SendOccasionalMessage(client *whatsmeow.Client, jid types.JID, username string) {
	for dat := range cfg.resMessage {
		fmt.Println("sending occasion messages")
		client.SendMessage(context.Background(), jid, &waProto.Message{
			Conversation: proto.String(formatMessage(dat.Body, username)),
		})
	}
}

func formatMessage(messge string, username string) string {
	if username == ""{
		username = "dear"
	}
	formattedMessage := strings.Replace(messge, "[User]", username, -1)
	return formattedMessage
}
