package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"tokatu4561/line-bot-weight/models"
	line "tokatu4561/line-bot-weight/service"

	"github.com/line/line-bot-sdk-go/linebot"
)

func WeightRegist(w http.ResponseWriter, r *http.Request) {
	line, err := line.LineConnection()
	if err != nil {
		log.Fatalln(err)
	}

	lineEvents, err := line.Client.ParseRequest(r)

	for _, event := range lineEvents {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			
			case *linebot.TextMessage:
				log.Println(message)
				// replyMessage := message.Text
				err = recordWeight(line, event)
				if err != nil {
					writeJSON(w, http.StatusOK)
				}
				break
			case *linebot.LocationMessage:
				break		
			default:
			}
		}
	}
}

func recordWeight(line *line.Line, event *linebot.Event) error {
	var userID int
	userID, _ = strconv.Atoi(event.Source.UserID)

	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", 
	os.Getenv("DB_HOST"), 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_DBNAME"))

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	DBModel := &models.DBModel{
		DB: db,
	}
	err = DBModel.AddUser(userID)
	if err != nil {
		return err
	}

	maxWeight, err := DBModel.GetMaxWeight(userID)
	if err != nil {
		return err
	}

	replyMessage := fmt.Sprintf("あなたの最も痩せていた体重は%dkgです", maxWeight)
	_, err = line.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()

	return err
}

func writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)

	return nil
}