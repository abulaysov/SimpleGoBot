package main

import (
	"SimpleGoBot/bot"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

var apiUrl string

func webHook(w http.ResponseWriter, req *http.Request) {
	resp := make(map[string]string)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("An error has occurred")
	}
	message := bot.TelegramMessage{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		fmt.Println("Unmarshal didn't work")
	}
	result := message.SendMessage(message.Message.Text, apiUrl)
	if !result {
		fmt.Println("An error has occurred when sending message")
	}
	resp["result"] = "ok"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("JSON doesn't encoding")
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env didn't load")
	}
	apiUrl = os.Getenv("API_URL")
	fmt.Println("Start working")
	http.HandleFunc("/webhook", webHook)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server doesn't work")
	}
}
