package bean

type EmailRequestSend struct {
	SenderEmail string `json:"sender_email"`
	ReceiverEmail string `json:"receiver_email"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

