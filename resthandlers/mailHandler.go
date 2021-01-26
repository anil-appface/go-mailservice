package resthandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// model which should be used to send as request body
type emailRequest struct {
	Recepient string `json:"recepient"`
	CC        string `json:"cc"`
	BCC       string `json:"bcc"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

type mailHandler struct {
}

//NewMailHandler creates new mailHandler
func NewMailHandler() *mailHandler {
	return &mailHandler{}
}

func (me *mailHandler) SendEmail(c *gin.Context) {
	var emailReq emailRequest
	c.BindJSON(&emailReq)
	// DB call to create a todo
	// Config.DB.Create(todo).Error;
	// err := Models.CreateATodo(&todo)
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.JSON(http.StatusOK, emailReq)
	// }

	c.JSON(http.StatusOK, emailReq)
}
