package genUtil

import (
	"context"
	ma "github.com/mailgun/mailgun-go/v4"

	"time"
)

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

func SendNotif(domain string, uName string, message string, llvmcompile string) error { //how p[arsed in astrr with return
	val := GetMailgunData()
	mg := ma.NewMailgun(domain, val)
	rPword := mg.NewMessage(
		"Gaurav <bansal22.gaurav@gmail.com>",
		"Notification",
		"Hey "+uName+",\nYou received a notification!!\n"+message+"\n", llvmcompile,
	)

	con, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(con, rPword)
	return err

}
