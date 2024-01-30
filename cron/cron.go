package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/charlesozo/whisperbot/internal/database"
	"github.com/robfig/cron/v3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const BaseURL = "https://whisper-message-api.onrender.com/api/v1/messages/search"

type Message struct {
	ID        int32     `json:"id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Body      string    `json:"body"`
	ImageName string    `json:"image_name"`
}

// resMessage chan<- Message
func cronMessage(date string, resMessage chan Message) {
	Client := &http.Client{}
	fullURL := fmt.Sprintf("%s/%s", BaseURL, date)
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatal(err)

	}
	resp, err := Client.Do(request)
	if err != nil {
		log.Fatal(err)

	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
	message := Message{}
	err = json.Unmarshal(dat, &message)
	if err != nil {
		log.Fatal(err)

	}
	resMessage <- message

}
func FormatMessage(messge string, username string) string {
	if username == "" {
		username = "dear"
	}
	formattedMessage := strings.Replace(messge, "[User]", username, -1)
	return formattedMessage
}

func RunTask(cli *whatsmeow.Client, db *database.Queries) {

	c := cron.New(cron.WithChain(cron.Recover(cron.DefaultLogger)))
	fmt.Println("Run task would soon start")
	response := make(chan Message)
	c.AddFunc("*/2 * * * *", func() {
		fmt.Println("This job runs every two minutes")
		go cronMessage("2024-02-14", response)
	})

	// Start the cron scheduler
	c.Start()
	defer c.Stop()
	for dat := range response {
		fmt.Println("task started")
		users, err := db.GetUnRegUsers(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		imageBytes, err := os.ReadFile("./assets/" + dat.ImageName + ".jpg")
		if err != nil {
			log.Fatalf("error reading images %v", err)
		}
		resp, err := cli.Upload(context.Background(), imageBytes, whatsmeow.MediaImage)
		if err != nil {
			log.Fatalf("error uploading image %v", err)
		}

		for _, user := range users {

			imageMsg := &waProto.ImageMessage{
				Caption:  proto.String(FormatMessage(dat.Body, user.DisplayName)),
				Mimetype: proto.String("image/png"), // replace this with the actual mime type
				// you can also optionally add other fields like ContextInfo and JpegThumbnail here
				Url:           &resp.URL,
				DirectPath:    &resp.DirectPath,
				MediaKey:      resp.MediaKey,
				FileEncSha256: resp.FileEncSHA256,
				FileSha256:    resp.FileSHA256,
				FileLength:    &resp.FileLength,
			}

			jid := types.NewJID(user.WhatsappNumber, types.DefaultUserServer)
			cli.SendMessage(context.Background(), jid, &waProto.Message{
				ImageMessage: imageMsg,
			})
		}

	}
}
