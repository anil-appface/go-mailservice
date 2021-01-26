package mailproviders

import "github.com/anil-appface/go-mailservice/model"

type AwsSesProvider struct {
}

func NewAwsSesProvider() *AwsSesProvider {
	return &AwsSesProvider{}
}

func (me *AwsSesProvider) SendEmail(emailParams *model.EmailParams) (interface{}, error) {
	return nil, nil
}
