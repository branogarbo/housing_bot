package main

import (
	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSI struct {
	Client *twilio.RestClient
	Params *twilioAPI.CreateMessageParams
}
