package mailproviders

import (
	"github.com/anil-appface/go-mailservice/model"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

type MailJetProvider struct {
	apiPublicKey  string
	apiPrivateKey string
	sender        string
}

//NewEmailJetProvider creates the mailjetProvider instance
func NewEmailJetProvider() *MailJetProvider {
	//TODO:: make the below keys.
	return &MailJetProvider{
		apiPublicKey:  "b93870e07a04be0fd063be8058d2647d",
		apiPrivateKey: "af0a2b1265c3a9eee67282fed39badef",
	}
}

//SendEmail sends an email with the provided email params
func (me *MailJetProvider) SendEmail(emailParams *model.EmailParams) (interface{}, error) {

	//creating a new mailjet client
	mailjetClient := mailjet.NewMailjetClient(me.apiPublicKey, me.apiPrivateKey)

	//small utility function to format the email array of strings to mailjet recepients
	formatEmail := func(emails []string) *mailjet.RecipientsV31 {
		var mails mailjet.RecipientsV31
		for _, email := range emails {
			mails = append(mails, mailjet.RecipientV31{
				Email: email,
			})
		}
		return &mails
	}

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: emailParams.Sender,
			},
			To:       formatEmail(emailParams.Recepients),
			Cc:       formatEmail(emailParams.CC),
			Bcc:      formatEmail(emailParams.BCC),
			Subject:  emailParams.Subject,
			TextPart: emailParams.Body,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return nil, err
	}
	return res, nil
}
