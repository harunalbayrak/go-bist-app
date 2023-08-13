package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func SendMessage(text string) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	fmt.Println("ChatID:", os.Getenv("CHAT_ID"))

	// Send the message
	tokenUrl := fmt.Sprintf("https://api.telegram.org/bot%s", os.Getenv("TELEGRAM_TOKEN"))
	url := fmt.Sprintf("%s/sendMessage", tokenUrl)
	body, _ := json.Marshal(map[string]string{
		"chat_id":           os.Getenv("CHAT_ID"),
		"message_thread_id": os.Getenv("MESSAGE_THREAD_ID"),
		"text":              text,
		// "parse_mode":        "MarkdownV2",
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// Log
	fmt.Printf("Message\n%s\n", text)
	fmt.Println("ResponseJSON", string(body))

	// Return
	return true, nil
}
