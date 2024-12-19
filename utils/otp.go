package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

/*
GenerateOTP generates a six-digit one-time password (OTP) and returns it as a string. It uses the
current Unix timestamp as a seed for the random number generator. The OTP is a random integer between
100000 and 999999.

Returns:
  - otp: A random generated six digit numeric string
  - err: Error, if any
*/
func GenerateOTP() (string, error) {
	seed := time.Now().Unix()
	rand.Seed(seed)

	otp := strconv.Itoa(rand.Intn(900000) + 100000)
	return otp, nil
}

/*
SendOTPSMS sends the provided one-time password (OTP) to the specified phone number via SMS.

Parameters:
  - phone: The phone number to which the OTP should be sent.
  - otp: The one-time password to be sent via SMS.

Returns:
  - An error if there is an issue with the SMS sending process or nil if successful.
*/
func SendOTPSMS(phone, otp string) error {
	// Implement your SMS sending logic here
	// Use a third-party service to send the OTP via SMS
	err := SparrowMessage(phone, otp)
	if err != nil {
		return err
	}

	// For simplicity, print the OTP to the console
	log.Printf("OTP sent to %s: %s", phone, otp)
	return nil
}

/*
SendOTPEmail sends an OTP (One-Time Password) via Email to the specified email address.
It takes the email address and OTP as input parameters and returns an error if any.
*/
func SendOTPEmail(email, otp string) error {
	// Implement your SMS sending logic here
	// Use a third-party service to send the OTP via Email

	err := OTPSMTP(email, otp)
	if err != nil {
		return err
	}

	// For simplicity, print the OTP to the console
	log.Printf("OTP sent to %s: %s", email, otp)
	return nil
}

/*
SparrowMessage sends an OTP (One-Time Password) via SMS using the Sparrow API. It takes a phone
number and OTP as input parameters and returns an error if any.
*/
func SparrowMessage(phone string, otp string) error {
	// Construct the SMS text with the OTP
	smsText := "Your OTP for NAME Online app verification is " + otp

	// Construct the payload for the HTTP request
	var payload = map[string]interface{}{
		"token": viper.GetString(`sms.sparrow.sms_key`),
		"from":  "Name",
		"to":    phone,
		"text":  smsText,
	}

	// Marshal the payload into JSON
	params, _ := json.Marshal(payload)
	request, err := http.NewRequest("POST", viper.GetString(`sms.sparrow.api`), bytes.NewBuffer(params))
	if err != nil {
		return err
	}

	// Set the content type header
	request.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check the response status
	if response.Status != "200 OK" {
		return errors.New("error sending message")
	}

	// For simplicity, print the OTP to the console
	log.Printf("OTP sent to %s: %s", phone, otp)

	return nil
}
