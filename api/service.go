package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAuthToken(),
})

// func to sendOTP to send OTP using twilio
func (app *Config) TwilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(envSERVICEID(), params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

// func to verifyOTP
func (*Config) TwilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICEID(), params)
	if err != nil {
		return err
	}

	if *resp.Status != "approved" {
		return errors.New("Invalid Code")
	}
	return nil
}
