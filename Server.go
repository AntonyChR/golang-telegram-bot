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

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(fmt.Sprintf(`{"message": "Method %s not allowed for this endpoint"}`, r.Method)))
			return
		}

		message, err := readMessageFromRequest(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message": "Error reading message: %s"}`, err.Error())))
			return
		}

		select {
		case s.IncomingMessage <- message:
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusServiceUnavailable)
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

	if err != nil {
		return Message{}, err
	}

	var body Body

	fmt.Println(string(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return Message{}, err
	}

	message := body.Message

	return message, nil
}
