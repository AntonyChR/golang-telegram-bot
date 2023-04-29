package gtb

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
	"strings"
)

type TextMessage struct {
	ChatID           int    `json:"chat_id"`
	ReplyToMessageID int    `json:"reply_to_message_id"`
	Text             string `json:"text"`
}

type ApiClient struct {
	BaseUrl   string
	EndPoints map[string]string
	Files     string
	ServerUrl string
}

func (t *ApiClient) RegisterWebhook() error {
	_, err := postJSON(t.BaseUrl+t.EndPoints["setWebhook"]+t.ServerUrl, nil)
	return err
}

func (t *ApiClient) RemoveWebhook() error {
	fmt.Println("[!] Removing webhook...")
	_, err := postJSON(t.BaseUrl+t.EndPoints["removeWebhook"], nil)
	return err
}

type obj map[string]string

func (t *ApiClient) SetCommandsDescription(desc map[string]string) error {
	if len(desc) == 0 {
		return nil
	}
	fmt.Println("[*] Adding command description")
	var commands []obj
	for k, v := range desc {
		commands = append(commands, obj{"command": k, "description": v})
	}
	description := map[string][]obj{
		"commands": commands,
	}
	objBytes, _ := json.Marshal(description)

	_, err := postJSON(t.BaseUrl+t.EndPoints["setCommands"], bytes.NewBuffer(objBytes))
	return err
}

func (t *ApiClient) SendText(data TextMessage) error {
	dataBytes, _ := json.Marshal(data)
	_, err := postJSON(t.BaseUrl+t.EndPoints["sendText"], bytes.NewBuffer(dataBytes))
	return err
}

func (t *ApiClient) SendFile(fileType string, relativePath string, text TextMessage) error {

	absPath := relativePath
	if strings.Contains(relativePath, "~/") {
		userHomeDir, _ := os.UserHomeDir()
		absPath = strings.ReplaceAll(relativePath, "~/", userHomeDir+"/")
	}

	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err.Error())
	}
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

type File struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FilePath     string `json:"file_path"`
}

// Get basic information about a file and prepare it for downloading
func (t *ApiClient) GetFileInfo(fileId string) File {

	dataBytes, _ := json.Marshal(map[string]string{"file_id": fileId})

	resp, _ := postJSON(t.BaseUrl+t.EndPoints["getFile"], bytes.NewBuffer(dataBytes))

	var fileData File

	_ = json.Unmarshal(resp, &fileData)

	return fileData

}

func (t *ApiClient) downloadFile(path string, dir string) error {

	file, err := os.Create(dir)
	if err != nil {
		return err
	}

	defer file.Close()

	resp, err := http.Get(t.Files + "/" + path)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	return err
}

func (t *ApiClient) banChatMember(userData map[string]any) error {
	dataBytes, _ := json.Marshal(userData)
	_, err := postJSON(t.BaseUrl+t.EndPoints["ban"], bytes.NewBuffer(dataBytes))
	return err
}

func (t *ApiClient) unbanChatMember(userData map[string]int) error {
	dataBytes, _ := json.Marshal(userData)
	_, err := postJSON(t.BaseUrl+t.EndPoints["unban"], bytes.NewBuffer(dataBytes))
	return err
}
