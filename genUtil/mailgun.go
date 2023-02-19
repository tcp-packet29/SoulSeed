package genUtil

import (
	"context"
	ma "github.com/mailgun/mailgun-go/v4"
	"main/storageUtil"
	"math"
	"strconv"
	"time"
)

func SendTokenMessage(domain string, token storageUtil.Token, sender string, org string, receiver string) (string, string) { //strings instead of models for the sake of simplcity
	val := GetMailgunData()
	mg := ma.NewMailgun(domain, val)
	aaa := strconv.FormatFloat(math.Round(time.Until(token.Expiry).Hours()), 'f', 0, 64)
	tok := mg.NewMessage(
		"Gaurav <bansal22.gaurav@gmail.com>", //tobedecided
		"You have been invited to the organization \""+org+"\" by "+sender,
		"Please use the following token to join the organization: "+token.Access+".\nThis token will expire in "+aaa+" hours.",
		receiver,
	)

	ca, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, _ = mg.Send(ca, tok)
	return "Success", ""
}
