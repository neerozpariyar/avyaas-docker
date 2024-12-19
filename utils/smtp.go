package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/spf13/viper"
)

/*
SMTP sends an email containing the OTP (One-Time Password) to the specified email address using SMTP protocol.
It takes the email address and OTP as input parameters and returns an error if any.
*/
func OTPSMTP(email string, otp string) error {
	// Parse the OTP template
	buff, err := ParseOTPTemplate(otp)
	if err != nil {
		return err
	}

	// Define the recipient, sender, password, SMTP host, and SMTP port
	to := []string{
		email,
	}

	from := viper.GetString(`smtp.sender`)
	password := viper.GetString(`smtp.password`)
	host := viper.GetString(`smtp.host`)
	port := viper.GetString(`smtp.port`)

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", from, password, host)

	// Send the email using SMTP
	err = smtp.SendMail(host+":"+port, auth, from, to, buff.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func TeacherAccountCreatedSMTP(email, userPassword string) error {
	// Parse the email template
	buff, err := ParseTeacherAccountCreateTemplate(email, userPassword)
	if err != nil {
		return err
	}

	// Define the recipient, sender, password, SMTP host, and SMTP port
	to := []string{
		email,
	}

	from := viper.GetString(`smtp.sender`)
	password := viper.GetString(`smtp.password`)
	host := viper.GetString(`smtp.host`)
	port := viper.GetString(`smtp.port`)

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", from, password, host)

	// Send the email using SMTP
	err = smtp.SendMail(host+":"+port, auth, from, to, buff.Bytes())
	if err != nil {
		return err
	}

	return nil
}

type OTPMessage struct {
	OTP string	
}

func ParseOTPTemplate(OTP string) (*bytes.Buffer, error) {
	temp, err := template.ParseFiles("./otp_template.html")
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	buff.Write([]byte(fmt.Sprintf("Subject: Verification OTP \n %s\n", mimeHeaders)))

	err = temp.Execute(buff, OTPMessage{
		OTP: OTP,
	})
	if err != nil {
		return nil, err
	}
	return buff, err
}

type TeacherAccountData struct {
	Email    string
	Password string
}

func ParseTeacherAccountCreateTemplate(email, password string) (*bytes.Buffer, error) {
	temp, err := template.ParseFiles("./account_created_template.html")
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	buff.Write([]byte(fmt.Sprintf("Subject: Teacher Account Created \n %s\n", mimeHeaders)))

	err = temp.Execute(buff, TeacherAccountData{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return buff, err
}
