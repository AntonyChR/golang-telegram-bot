package bot

import (
	"bytes"
	"fmt"
)

type ApiClient struct {
	BaseUrl   string
	EndPoints map[string]string
	ServerUrl string
}

func (t *ApiClient) RegisterWebhook() error {
	resp, err := post(t.BaseUrl+"/setWebhook?url="+t.ServerUrl, "application/json", nil)
	fmt.Println(resp)
	return err
}

func (t *ApiClient) RemoveWebhook() error {
	resp, err := post(t.BaseUrl+"/setWebhook?remove", "application/json", nil)
	fmt.Println(resp)
	return err
}

func (t *ApiClient) SetCommandsDescription(desc map[string]string) error {
	if len(desc) == 0 {
		return nil
	}
	fmt.Println("[i] Adding command description")
	objBytes := commandDesc2json(desc)
	resp, err := post(t.BaseUrl+"/setMyCommands", "application/json", bytes.NewBuffer(objBytes))
	fmt.Println(resp)
	return err
}

func (t *ApiClient) SendText(text string) error {
	_, err := post(t.BaseUrl+"/sendText", "application/json", bytes.NewBuffer([]byte(text)))
	return err
}

func (t *ApiClient) SendFile(endPoint string, fileType string, fileContent []byte) {
}
