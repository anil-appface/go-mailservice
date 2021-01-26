package mailproviders

import "github.com/anil-appface/go-mailservice/model"

// Define all the possible constants for mail providers
const (
	MailJet = "MAILJET"
	SMTP    = "SMTP"
	AWSSES  = "AWSSES"
)

type MailProvider interface {
	SendEmail(emailParams *model.EmailParams) (interface{}, error)
}

//GetEmailProvider based on the mailprovider given, it gives the decoupled mail provider
func GetEmailProvider(mailProvider string) MailProvider {
	switch mailProvider {
	case MailJet:
		return NewEmailJetProvider()
	case SMTP:
		return NewSMTPProvider()
	case AWSSES:
		return NewAwsSesProvider()
	default:
		return NewEmailJetProvider()
	}
}
