package bot

import (
	"bytes"
	"fmt"
)

type TelegramApi struct {
	BaseUrl   string
	EndPoints map[string]string
	ServerUrl string
}

func (t *TelegramApi) RegisterWebhook() error {
	resp, err := post(t.BaseUrl+"/setWebhook?url="+t.ServerUrl, "application/json", nil)
	fmt.Println(resp)
	return err
}
func (t *TelegramApi) SetCommandsDescription(desc map[string]string) error {
	if len(desc) == 0 {
		return nil
	}
	fmt.Println("[i] Adding command description")
	objBytes := commandDesc2json(desc)
	resp, err := post(t.BaseUrl+"/setMyCommands", "application/json", bytes.NewBuffer(objBytes))
	fmt.Println(resp)
	return err
}

func (t *TelegramApi) RemoveWebhook() error {
	resp, err := post(t.BaseUrl+"/setWebhook?remove", "application/json", nil)
	fmt.Println(resp)
	return err
}

func (t *TelegramApi) SendText(text string) {

}

func (t *TelegramApi) SendFile(endPoint string, fileType string, fileContent []byte) {
}
