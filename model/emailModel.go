package model

//EmailReqModel which should be used to send as request body
type EmailReqModel struct {
	EmailProvider string      `json:"emailProvider"`
	EmailParams   EmailParams `json: "emailParams"`
}

// EmailParams common email params used by all the providers
type EmailParams struct {
	Sender     string   `json:"sender"`
	Recepients []string `json:"recepients"`
	CC         []string `json:"cc"`
	BCC        []string `json:"bcc"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
}
