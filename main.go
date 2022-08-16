package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	line "tokatu4561/line-bot-weight/service"

	"github.com/line/line-bot-sdk-go/linebot"
)

type application struct {
}

func main() {
	app := &application{}
	err := app.serve()
	if err != nil {
		log.Fatalln(err)
	}

	line, err := line.LineConnection()

	alertMessage := "今日の体重を教えてください"
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(alertMessage)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := line.Client.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}


func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	return srv.ListenAndServe()
}