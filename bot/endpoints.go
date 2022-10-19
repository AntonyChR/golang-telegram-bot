package bot

func endpoints() map[string]string {
	return map[string]string{
		"setWebhook":    "/setWebhook?url=",
		"removeWebhook": "/setWebhook?remove",
		"setCommands":   "/setMyCommands",
		"sendText":      "/sendMessage",
		"sendVideo":     "",
		"sendImg":       "",
		"sendAudio":     "",
		"sendDoc":       "",
	}
}
