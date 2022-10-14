package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type server struct {
	Port            string
	IncomingCommand chan Command
}

func (s *server) Start() {

	fmt.Println("[i] Initializating server on: " + s.Port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			command, _ := getCommand(r)
			s.IncomingCommand <- command
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`{"message": "Can't find method requested"}`))
		}
	})
	err := http.ListenAndServe(":"+s.Port, nil)
	fmt.Println(err)
}

func getCommand(r *http.Request) (Command, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return Command{}, err
	}
	defer r.Body.Close()
	var body Body
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return Command{}, err
	}
	textMessage := strings.Split(body.Message.Text, " ")
	cmd := Command{
		Name:             textMessage[0],
		Params:           textMessage[1:],
		ChatId:           body.Message.Chat.ID,
		ReplyToMessageId: body.Message.MessageID,
	}
	return cmd, nil
}
