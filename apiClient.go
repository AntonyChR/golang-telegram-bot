package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
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
	_, err := post(t.BaseUrl+t.EndPoints["removeWebhook"], "application/json", nil)
	return err
}

func (t *ApiClient) SetCommandsDescription(desc map[string]string) error {
	if len(desc) == 0 {
		return nil
	}
	fmt.Println("[*] Adding command description")
	objBytes := commandDesc2json(desc)
	_, err := post(t.BaseUrl+t.EndPoints["setCommands"], "application/json", bytes.NewBuffer(objBytes))
	return err
}

func (t *ApiClient) SendText(data TextMessage) error {
	dataBytes, _ := json.Marshal(data)
	_, err := post(t.BaseUrl+t.EndPoints["sendText"], "application/json", bytes.NewBuffer([]byte(dataBytes)))
	return err
}

func (t *ApiClient) SendFile(fileType string, relativePath string, text TextMessage) error {

	currentDir, _ := os.Getwd()
	absolutePath := path.Join(currentDir, relativePath)

	file, _ := os.Open(absolutePath)
	defer file.Close()

	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)
	part, _ := writer.CreateFormFile(fileType, path.Base(file.Name()))
	io.Copy(part, file)

	writer.WriteField("chat_id", strconv.Itoa(text.ChatID))
	writer.WriteField("reply_to_message_id", strconv.Itoa(text.ReplyToMessageID))
	writer.WriteField("caption", text.Text)
	writer.Close()

	var endpoint string

	switch fileType {
	case "document":
		endpoint = "sendDoc"
	case "photo":
		endpoint = "sendImg"
	case "audio":
		endpoint = "sendAudio"
	case "video":
		endpoint = "sendVideo"
	}

	req, _ := http.NewRequest("POST", t.BaseUrl+t.EndPoints[endpoint], &body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
	}
	return err
}
