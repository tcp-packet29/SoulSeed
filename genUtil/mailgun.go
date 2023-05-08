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
		"",
		receiver,
	)

	tok.SetHtml("<h1>Please use the following token to join the organization: " + token.Access + ".</h1><br>This token will expire in " + aaa + " hours")

	ca, cancel := context.WithTimeout(context.Background(), time.Second*14)
	defer cancel()

	_, _, _ = mg.Send(ca, tok)
	return "Success", ""
}

// edit ast
// hash pasdswrod and send t6oken mg wqith mailgu n
func setResPword(domain string, uName string, tok string, g string) error { //how p[arsed in astrr with return
	val := GetMailgunData()
	mg := ma.NewMailgun(domain, val)
	rPword := mg.NewMessage(
		"Gaurav <bansal22.gaurav@gmail.com>",
		"Reset Password",
		"Please use the following token to reset your password, "+uName+": "+tok+".\n This token will expire in 1 hour.",
		g)
	con, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(con, rPword)
	return err

}
