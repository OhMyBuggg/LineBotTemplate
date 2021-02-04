// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	Content "content"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/", homeHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<H1>Hello, this is Bugbot's homepage!</H1>")
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				text := message.Text
				cmd := strings.Split(text, " ")

				switch cmd[0] {
				case "測試":
					quota, err := bot.GetMessageQuota().Do()
					if err != nil {
						log.Println("Quota err:", err)
					}
					reply := Content.GetTestMessage(event.ReplyToken, message.ID, message.Text, event.Source.UserID, quota.Value)
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do()
					if err != nil {
						log.Print(err)
					}
				case "蟲蟲機器人":
					reply := Content.Status
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do()
					if err != nil {
						log.Print(err)
					}
				case "重複":
					reply := strings.TrimLeft(message.Text, "重複 ")
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do()
					if err != nil {
						log.Print(err)
					}
				case "大頭貼":
					target := strings.TrimLeft(message.Text, "大頭貼 ")
					var sendr *linebot.Sender
					switch target {
					case "熊大":
						sendr = linebot.NewSender("Brown", Content.BrownImage)
					case "兔兔":
						sendr = linebot.NewSender("Cony", Content.ConyImage)
					case "饅頭人":
						sendr = linebot.NewSender("Moon", Content.MoonImage)
					default:
						reply := Content.StickerUsage
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do()
						if err != nil {
							log.Print(err)
						}
					}

					if sendr != nil {
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hi, this is "+target+", Nice to meet you.").WithSender(sendr)).Do() 
						if err != nil {
							log.Print(err)
						}
					}
				default:
					reply := Content.Usage
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply)).Do()
					if err != nil {
						log.Print(err)
					}
				}

			}

		}
	}

}
