package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// レスポンスの形式に合わせて struct を定義する
type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

func GetWeather() (str string, err error) {
	body, err := httpGetBody("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json")
	if err != nil {
		// エラーを呼び出し元へ委譲する
		return str, err
	}
	weather, err := formatWeather(body)
	if err != nil {
		// エラーを呼び出し元へ委譲する
		return str, err
	}

	result := weather.ToS()

	return result, nil
}

func httpGetBody(url string) ([]byte, error) {
	// HTTPリクエストを発行しレスポンスを取得する
	response, err := http.Get(url)
	if err != nil {
		// エラーをラップして返す
		err = fmt.Errorf("Get Http Error: %s", err)
		return nil, err
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// エラーをラップして返す
		err = fmt.Errorf("IO Read Error:: %s", err)
		return nil, err
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()

	return body, nil
}

func formatWeather(body []byte) (*Weather, error) {
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		// エラーをラップして返す
		err = fmt.Errorf("JSON Unmarshal error: %s", err)
		return nil, err
	}
	return weather, nil
}

func (w *Weather) ToS() string {
	area := fmt.Sprintf("%sの天気です。\n", w.Area)
	head := fmt.Sprintf("%s\n", w.HeadLine)
	body := fmt.Sprintf("%s\n", w.Body)
	result := area + head + body

	return result
}