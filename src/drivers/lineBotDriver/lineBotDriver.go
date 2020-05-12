package lineBotDriver

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"net/http"
)

type LineBotSdkClient interface {
	PushMessage(to string, messages ...linebot.SendingMessage) *linebot.PushMessageCall
	ReplyMessage(replyToken string, messages ...linebot.SendingMessage) *linebot.ReplyMessageCall
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
}
