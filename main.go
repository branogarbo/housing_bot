package main

import (
	"os"
	"strconv"
	"time"
)

var (
	reqURL               = os.Getenv("REQ_URL")
	pageCookie           = os.Getenv("PAGE_COOKIE")
	discordToken         = os.Getenv("DISCORD_TOKEN")
	channelID            = os.Getenv("CHANNEL_ID")
	checkIntervalStr     = os.Getenv("CHECK_INTERVAL")
	searchPattern        = os.Getenv("SEARCH_PATTERN")
	alertWhenFound, _    = strconv.ParseBool(os.Getenv("ALERT_WHEN_FOUND"))
	timezone             = os.Getenv("TIMEZONE")
	reqMethod            = os.Getenv("REQ_METHOD")
	reqBody              = os.Getenv("REQ_BODY")
	isAPIendpoint, _     = strconv.ParseBool(os.Getenv("IS_API_ENDPOINT"))
	reqVerificationToken = os.Getenv("REQ_VERIFICATION_TOKEN")
	linkPage             = os.Getenv("LINK_PAGE")
)

var (
	checkInterval time.Duration
	possibleCmds  = []string{"check", "help", "stop", "start"}
	isChecking    = true
	lastResponse  = "No response yet"
)

func main() {
	Run()
}
