package bot

import (
	"fmt"
	"net/http"
)

type TelegramMessage struct {
	Message struct {
		Chat struct {
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Id        int    `json:"id"`
			Type      string `json:"type"`
		}
		From struct {
			FirstName    string `json:"first_name"`
			LanguageCode string `json:"language_code"`
			Username     string `json:"username"`
			IsBot        bool   `json:"is_bot"`
		}
		Text string `json:"text"`
	}
}

func (tm *TelegramMessage) SendMessage(text string, url string) bool {
	url = fmt.Sprintf("%ssendMessage?text=%s&chat_id=%d", url, text, tm.Message.Chat.Id)
	_, err := http.Get(url)
	if err != nil {
		return false
	}
	return true
}
