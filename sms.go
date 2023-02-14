package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
)

func initTwilio() (SMSI, error) {
	accountSID := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	twilioPhone := os.Getenv("TWILIO_PHONE")
	destPhone := os.Getenv("DEST_PHONE")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	params := &twilioAPI.CreateMessageParams{}
	params.SetFrom(twilioPhone)
	params.SetTo(destPhone)

	return SMSI{client, params}, nil
}

func (s SMSI) notifyUser(message string) error {
	s.Params.SetBody(message)

	fmt.Println(message)

	_, err := s.Client.Api.CreateMessage(s.Params)
	if err != nil {
		return err
	}

	return nil
}
