package main

import (
	"log"

	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
)

func initTwilio() (SMSI, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: twilioToken,
	})

	params := &twilioAPI.CreateMessageParams{}
	params.SetFrom(twilioPhone)
	params.SetTo(destPhone)

	return SMSI{client, params}, nil
}

func (s SMSI) notifyUser(message string) error {
	s.Params.SetBody(message)

	err := printToLog(message)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Client.Api.CreateMessage(s.Params)
	if err != nil {
		return err
	}

	return nil
}
