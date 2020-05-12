package lineBot

import (
	"context"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"net/http"
	"time"
)

func (lineBot *LineBotClient) ReplyMessage(r *http.Request, messageReply string, timeout time.Duration) error {
	gg := time.Now()
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	events, err := lineBot.sdkClient.ParseRequest(r)
	if err != nil {
		fmt.Println("err line", err)
	}

	textMessage := linebot.NewTextMessage(messageReply)
	var eventObject *linebot.Event

	ch := make(chan int64)

	go func() {
		timer := time.Now()
		for _, event := range events {
			eventObject = event
			for {
				if ctx.Err() != nil {
					break
				}
				replyToken := event.ReplyToken
				if lineBot.testNotificationAdmin {
					replyToken = "11"
				}

				_, err = lineBot.sdkClient.ReplyMessage(replyToken, textMessage).Do()

				if err == nil {
					break
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}
		duration := time.Since(timer)
		ch <- duration.Milliseconds()
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("timeout in %d millisecend\n", time.Since(gg), err)
		messageError := linebot.NewTextMessage("UserID: " + eventObject.Source.UserID + " ERROR: " + err.Error() + " replyToken: " + eventObject.ReplyToken)
		lineBot.sdkClient.PushMessage(lineBot.adminId, messageError).Do()
	case t := <-ch:
		fmt.Printf("JOB DONE in %d seconds\n", t)
	}
	fmt.Println("success")
	if err != nil {
		return err
	}
	return nil
}
