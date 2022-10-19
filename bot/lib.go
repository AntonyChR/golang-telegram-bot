package bot

import (
	"fmt"
	"io"
	"net/http"
)

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
