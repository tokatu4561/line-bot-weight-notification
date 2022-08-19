package handlers

import (
	"database/sql"
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
	if err != nil {
		log.Fatalln(err)
	}

	for _, event := range lineEvents {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			
			case *linebot.TextMessage:
				log.Println(message)
				// replyMessage := message.Text
				err = recordWeight(line, event)
				if err != nil {
					log.Fatalln(err)
				}
			case *linebot.LocationMessage:
				break		
			default:
			}
		}
	}
}

func recordWeight(line *line.Line, event *linebot.Event) error {
	var lineID int
	lineID, _ = strconv.Atoi(event.Source.UserID)

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

	user, err := DBModel.GetOneUser(lineID)
	if err != nil {
		// ユーザーが存在しなかった場合新規作成
		err = DBModel.AddUser(lineID)
		if err != nil {
			return err
		}	
		user, _ = DBModel.GetOneUser(lineID)
	}

	maxWeight, err := DBModel.GetMaxWeight(user.ID)
	if err != nil {
		return err
	}

	message :=event.Message.(*linebot.TextMessage).Text

	replyMessage := fmt.Sprintf("%skg! あなたの最も痩せていた体重は%dkgです", message ,maxWeight)
	_, err = line.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()

	return err
}

// func writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
// 	out, err := json.MarshalIndent(data, "", "\t")
// 	if err != nil {
// 		return err
// 	}

// 	if len(headers) > 0 {
// 		for k, v := range headers[0] {
// 			w.Header()[k] = v
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	w.Write(out)

// 	return nil
// }