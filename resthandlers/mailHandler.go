package resthandlers

import (
	"net/http"

	"github.com/anil-appface/go-mailservice/mailproviders"
	"github.com/anil-appface/go-mailservice/model"
	"github.com/gin-gonic/gin"
)

type mailHandler struct {
}

//NewMailHandler creates new mailHandler
func NewMailHandler() *mailHandler {
	return &mailHandler{}
}

//SendEmail router to handle the sendEmail function.
func (me *mailHandler) SendEmail(c *gin.Context) {

	var emailReq model.EmailReqModel
	//Bind the POST request body
	err := c.BindJSON(&emailReq)
	//if error while sending an email return error with 500 status
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//TODO:: configure the sender.
	if emailReq.EmailParams.Sender == "" {
		emailReq.EmailParams.Sender = "anilamilineni01@gmail.com"
	}

	emailProvider := mailproviders.GetEmailProvider(emailReq.EmailProvider)
	res, err := emailProvider.SendEmail(&emailReq.EmailParams)

	//if error while sending an email return error with 500 status
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//send the response to caller
	c.JSON(http.StatusOK, res)
}
