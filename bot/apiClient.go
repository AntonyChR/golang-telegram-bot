package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type TextMessage struct {
	ChatID           int    `json:"chat_id"`
	ReplyToMessageID int    `json:"reply_to_message_id"`
	Text             string `json:"text"`
}

type ApiClient struct {
	BaseUrl   string
	EndPoints map[string]string
	ServerUrl string
}

func (t *ApiClient) RegisterWebhook() error {
	_, err := post(t.BaseUrl+t.EndPoints["setWebhook"]+t.ServerUrl, "application/json", nil)
	return err
}

func (t *ApiClient) RemoveWebhook() error {
	fmt.Println("[!] Removing webhook...")
	resp, err := post(t.BaseUrl+t.EndPoints["removeWebhook"], "application/json", nil)
	fmt.Println(resp)
	return err
}

func (t *ApiClient) SetCommandsDescription(desc map[string]string) error {
	if len(desc) == 0 {
		return nil
	}
	fmt.Println("[i] Adding command description")
	objBytes := commandDesc2json(desc)
	resp, err := post(t.BaseUrl+t.EndPoints["setCommands"], "application/json", bytes.NewBuffer(objBytes))
	fmt.Println(resp)
	return err
}

func (t *ApiClient) SendText(data TextMessage) error {
	dataBytes, _ := json.Marshal(data)
	_, err := post(t.BaseUrl+t.EndPoints["sendText"], "application/json", bytes.NewBuffer([]byte(dataBytes)))
	return err
}

func (t *ApiClient) SendFile(endPoint string, fileType string, fileContent []byte) {
}
