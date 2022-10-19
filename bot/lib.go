package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type commands struct {
	Commands []commandDesc `json:"commands"`
}

type commandDesc struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func commandDesc2json(m map[string]string) []byte {
	var cmds []commandDesc
	for k, v := range m {
		cmds = append(cmds, commandDesc{Command: k, Description: v})
	}
	obj := commands{
		Commands: cmds,
	}
	objBytes, _ := json.Marshal(obj)
	return objBytes
}

func post(url string, contentType string, content io.Reader) (string, error) {
	fmt.Println(url)
	resp, err := http.Post(url, contentType, content)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
