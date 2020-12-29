package model
type EmailSendModel struct {
	EmailBucket []string
}

func NewEmailSendModel(bucketSize int) *EmailSendModel  {
	return &EmailSendModel{
		EmailBucket: make([]string, bucketSize),
	}
}

func (m *EmailSendModel) ProcessData(data string) error  {
	return nil
}

func (m *EmailSendModel) FlushData() error  {
	return nil
}