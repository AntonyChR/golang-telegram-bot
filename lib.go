package gtb

import (
	"encoding/json"
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

func postJSON(url string, content io.Reader) ([]byte, error) {

	contentType := "application/json"
	resp, err := http.Post(url, contentType, content)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return bodyBytes, nil
}
