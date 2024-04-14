package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	// "os"
	"github.com/charlesozo/whisperbot/internal/database"
	"github.com/google/uuid"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

var (
	wg sync.WaitGroup // Mutex for concurrent map access
)

func FormatMessage(message string, username string) string {
	if username == "" {
		username = "dear"
	}
	formattedMessage := strings.Replace(message, "[User]", username, -1)
	return formattedMessage
}
func (cfg *waConfig) databaseHandling(ctx context.Context, username, waNumber string, wg *sync.WaitGroup){
	defer wg.Done()
	ctxt, cancel := context.WithCancel(ctx)
	defer cancel()
    newUser, err := cfg.DB.CreateUser(ctxt, database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		WhatsappNumber: waNumber,
		DisplayName:    username,
	})
	if err != nil {
		 // Pass the error message into the context
		 ctxt = context.WithValue(ctxt, CTXErrMessage, err.Error())
		 cancel() // Cancel the context to stop further operations
		 return
	}
		// Create the user subscription
		err = cfg.DB.CreateUserSubscription(ctxt, database.CreateUserSubscriptionParams{
			SubscriptionID: uuid.New(),
			Userid:         uuid.NullUUID{UUID: newUser.ID, Valid: true},
			ExpiryDate:     time.Now().AddDate(0, 0, 7),
		})
		if err != nil {
			 // Pass the error message into the context
			 ctxt = context.WithValue(ctxt, CTXErrMessage, err.Error())
			 cancel() // Cancel the context to stop further operations
			 return
		}
		fmt.Printf("complete addition of %s to database", waNumber)

}

func (cfg *waConfig) HandleNewUser(ctx context.Context, client *whatsmeow.Client, chatJID, senderJID types.JID, username string, waNumber string){
	ctxt, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start a goroutine to create the user and subscription concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
			// Create the user
			
		newUser, err := cfg.DB.CreateUser(ctxt, database.CreateUserParams{
			ID:             uuid.New(),
			CreatedAt:      time.Now().UTC(),
			WhatsappNumber: waNumber,
			DisplayName:    username,
		})
		if err != nil {
			 // Pass the error message into the context
			 ctxt = context.WithValue(ctxt, CTXErrMessage, err.Error())
			 cancel() // Cancel the context to stop further operations
			 
		}

		// Create the user subscription
		err = cfg.DB.CreateUserSubscription(ctx, database.CreateUserSubscriptionParams{
			SubscriptionID: uuid.New(),
			Userid:         uuid.NullUUID{UUID: newUser.ID, Valid: true},
			ExpiryDate:     time.Now().AddDate(0, 0, 7),
		})
		if err != nil {
			 // Pass the error message into the context
			 ctxt = context.WithValue(ctxt, CTXErrMessage, err.Error())
			 cancel() // Cancel the context to stop further operations
			 
		}
		fmt.Printf("complete addition of %s to database", waNumber)
	}()


	for _, message := range NewUserMessages {
		
		err := cfg.SendMessage(ctxt, client, senderJID, message, username, "conversation")
		if err != nil {
			 // Pass the error message into the context
			 ctxt = context.WithValue(ctxt, CTXErrMessage, err.Error())
			 cancel() // Cancel the context to stop further operations
			 return
		}
		
	}
	
	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("This is first time user")
}

func (cfg *waConfig) SendMessage(ctx context.Context, client *whatsmeow.Client,senderJID types.JID, message, username, messageType string) error {
	switch messageType {
	case "chat":
		chat := &waProto.Chat{
			DisplayName: &username,
			Id: proto.String("id1"),
		}
     sent, err := client.SendMessage(ctx, senderJID, &waProto.Message{
        Chat: chat,
	 })
	 if err != nil {
		return err
	}
	fmt.Print(sent.Timestamp, sent.ID)
	case "templateMessage":
		TemplateMessage := &waProto.TemplateMessage{
			HydratedTemplate: &waProto.TemplateMessage_HydratedFourRowTemplate{
				Title: &waProto.TemplateMessage_HydratedFourRowTemplate_HydratedTitleText{
					HydratedTitleText: "The Title",
				  },
				  TemplateId:          proto.String("template-id"),
				  HydratedContentText: proto.String("The Content"),
				  HydratedFooterText:  proto.String("The Footer"),
				//   MaskLinkedDevices: proto.Bool(true),
				  HydratedButtons: []*waProto.HydratedTemplateButton{
					 // This for URL button
					 {
						Index: proto.Uint32(1),
						HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
						  UrlButton: &waProto.HydratedTemplateButton_HydratedURLButton{
							DisplayText: proto.String("The Link"),
							Url:         proto.String("https://fb.me/this"),
						  },
						},
					  },
					   // This for call button
					   {
						Index: proto.Uint32(2),
						HydratedButton: &waProto.HydratedTemplateButton_CallButton{
						  CallButton: &waProto.HydratedTemplateButton_HydratedCallButton{
							DisplayText: proto.String("Call us"),
							PhoneNumber: proto.String("1234567890"),
						  },
						},
					  },
					  {
						Index: proto.Uint32(3),
						HydratedButton: &waProto.HydratedTemplateButton_QuickReplyButton{
						  QuickReplyButton: &waProto.HydratedTemplateButton_HydratedQuickReplyButton{
							DisplayText: proto.String("Quick reply"),
							Id:          proto.String("quick-id"),
						  },
						},
					  },
				  },
			},
			
		}
          log.Print("four_row_template",TemplateMessage.GetHydratedTemplate() ,"---\n")
		  log.Print("templateid", TemplateMessage.TemplateId,"---\n")
		  log.Print("templateString",TemplateMessage.String(),"---\n")
		TemplateMessage.ProtoMessage()
		
		res, err := client.SendMessage(ctx, senderJID, &waProto.Message{
			TemplateMessage: TemplateMessage,
		}, whatsmeow.SendRequestExtra{
			Peer: true,
		})
		if err != nil {
			return err
		}
		fmt.Print(res.Timestamp, res.ID)
		
		
	case "conversation":
		_, err := client.SendMessage(ctx, senderJID, &waProto.Message{
			Conversation: proto.String(FormatMessage(message, username)),
		})
		if err != nil {
			return err
		}
		
	default:
		log.Fatal("not a message type")
	}
	return nil
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
