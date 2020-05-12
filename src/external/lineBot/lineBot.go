package lineBot

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tkhamsila/backendtest/src/drivers/lineBotDriver"
)

type LineBotClient struct {
	sdkClient lineBotDriver.LineBotSdkClient

	chanelSecret          string
	chanelToken           string
	adminId               string
	testNotificationAdmin bool
}

func New(chanelSecret, chanelToken, adminId string, testNotificationAdmin bool) *LineBotClient {
	bot, err := linebot.New(
		chanelSecret,
		chanelToken,
	)
	if err != nil {
		fmt.Println("Unable Setup LineBot")
	}
	return &LineBotClient{
		sdkClient:             bot,
		chanelSecret:          chanelSecret,
		chanelToken:           chanelToken,
		adminId:               adminId,
		testNotificationAdmin: testNotificationAdmin,
	}
}

func (lineBot *LineBotClient) setSdkClinet() *LineBotClient {
	return lineBot
}
