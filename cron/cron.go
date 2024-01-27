package cron
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
     "context"
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
type CronStack struct{
	cron *cron.Cron
}
//resMessage chan<- Message
func cronMessage(ctx context.Context, date string, resMessage chan Message){
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

	// select {
	// case resMessage <- message:
	// 	// Message sent successfully.
	// 	return nil
	// case <-ctx.Done():
	// 	// Context done, handle cancellation or timeout.
	// 	return ctx.Err()
	// }
	
}
  
func NewCron() *CronStack{
	c := cron.New(cron.WithChain(cron.Recover(cron.DefaultLogger)))
	return &CronStack{cron: c}
}
func (c *CronStack) RunTask(response chan Message) {
	fmt.Println("Run task would soon start")
	ctx := context.Background()
	c.cron.AddFunc("*/2 * * * *", func() {
		fmt.Println("This job runs every two minutes")
		go cronMessage(ctx, "2024-02-14", response)
	})

	// Start the cron scheduler
	c.cron.Start()
	defer c.cron.Stop()
	
}