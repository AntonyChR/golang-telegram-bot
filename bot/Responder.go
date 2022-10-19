package bot

type Responder struct{
	apiService *ApiClient
}


type Content struct{
	Text string
	Type string
	Data []byte
}

func (r *Responder) Reply(m Message,c Content){
}
