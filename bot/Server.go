package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type server struct {
	Port            string
	IncomingCommand chan Message
}

func (s *server) Start() {

	fmt.Println("[i] Initializating server on: ", s.Port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			command, _ := readMessage(r)
			s.IncomingCommand <- command
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"message": "Can't find method requested"}`))
		}
	})
	err := http.ListenAndServe(":"+s.Port, nil)
	fmt.Println(err)
}

func readMessage(r *http.Request) (Message, error) {

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		return Message{}, err
	}

	defer r.Body.Close()

	var body Body

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return Message{}, err
	}
	message := body.Message
	message.HasImages = len(message.Photo) > 0
	return message, nil
}
