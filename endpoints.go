package gtb

var endpoints = map[string]string{
	"setWebhook":    "/setWebhook?url=",
	"removeWebhook": "/setWebhook?remove",

	"setCommands": "/setMyCommands",

	"sendText":  "/sendMessage",
	"sendVideo": "/sendDocument",
	"sendImg":   "/sendPhoto",
	"sendAudio": "/sendAudio",
	"sendDoc":   "/sendDocument",
}
