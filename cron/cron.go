package cron
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
 	"github.com/robfig/cron/v3"
)

const BaseURL = "https://whisper-message-api.onrender.com/api/v1/messages/search"

type Message struct {
	ID        int32     `json:"id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Body      string    `json:"body"`
	ImageName string    `json:"image_name"`
}

//resMessage chan<- Message
func cronMessage(date string, resMessage chan Message){
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
  

func RunTask() {
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
		fmt.Println(dat.Body)
	}
}