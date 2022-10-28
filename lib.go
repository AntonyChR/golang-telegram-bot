package gtb

import (
	"io"
	"net/http"
)

func postJSON(url string, content io.Reader) ([]byte, error) {

	resp, err := http.Post(url, "application/json", content)

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
