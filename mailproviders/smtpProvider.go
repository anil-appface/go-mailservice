package mailproviders

import "github.com/anil-appface/go-mailservice/model"

type SMTPProvider struct {
}

func NewSMTPProvider() *SMTPProvider {
	return &SMTPProvider{}
}

func (me *SMTPProvider) SendEmail(emailParams *model.EmailParams) (interface{}, error) {
	return nil, nil
}
