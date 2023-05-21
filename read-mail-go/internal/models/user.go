package models

type User struct {
	ID    int
	Name  string
	Token string
}

type Response struct {
	Data interface{} `json:"data"`
}

type Message struct {
	Subject string `json:"subject"`
	Body    struct {
		ContentType string `json:"contentType"`
		Content     string `json:"content"`
	} `json:"body"`
}
