package genUtil

import (
	"github.com/gin-gonic/gin"
	ma "github.com/mailgun/mailgun-go/v3"
	"main/storageUtil"
	"time"
)

func sendTokenMessage(c *gin.Context, domain string, token storageUtil.Token, sender storageUtil.User, org storageUtil.Organization, receiver string) (string, string) {
	val := GetMailgunData()
	mg := ma.NewMailgun(domain, val)
	tok := mg.NewMessage(
		"Gaurav <me@fruitbear.com>", //tobedecided
		"You have been invited to the organization "+org.Name+" by "+sender.Username,
		"Please use the following token to join the organization: "+token.Access+".\nThis token will expire in "+time.Until(token.Expiry).String(),
		receiver,
	)

	_, id, err := mg.Send(c, tok)
	return id, err.Error()
}
