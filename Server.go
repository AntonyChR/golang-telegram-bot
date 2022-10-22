package gtb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type server struct {
	Port            string
	IncomingMessage chan Message
}

func (s *server) Start() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			message, _ := readMessageFromRequest(r)
			s.IncomingMessage <- message
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"message": "Can't find method requested"}`))
		}
	})

	fmt.Println("[*] Listening for messages on port: ", s.Port)
	err := http.ListenAndServe(":"+s.Port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func readMessageFromRequest(r *http.Request) (Message, error) {

	bodyBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Println(string(bodyBytes))
	if err != nil {
		return Message{}, err
	}

	var body Body

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return Message{}, err
	}

	message := body.Message

	return message, nil
}
