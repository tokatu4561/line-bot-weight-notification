package main

import (
	"log"
	line "tokatu4561/line-bot-weight/notification-service/service"

	// "github.com/joho/godotenv" herokuデプロイ様ににコメントアウト
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// _ = godotenv.Load(".env") // herokuデプロイ様ににコメントアウト
	line, err := line.LineConnection()
	if err != nil {
		log.Fatalln(err)
	}

	alertMessage := "今日の体重を教えてください"
	message := linebot.NewTextMessage(alertMessage)
	
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := line.Client.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}

