package main

import "os"

var (
	accountSID     = os.Getenv("ACCOUNT_SID")
	twilioPhone    = os.Getenv("TWILIO_PHONE")
	destPhone      = os.Getenv("DEST_PHONE")
	pageURL        = os.Getenv("PAGE_URL")
	pageCookie     = os.Getenv("PAGE_COOKIE")
	twilioToken    = os.Getenv("TWILIO_TOKEN")
	checkInterval  = os.Getenv("CHECK_INTERVAL")
	htmlPattern    = os.Getenv("HTML_PATTERN")
	alertWhenFound = os.Getenv("ALERT_WHEN_FOUND")
)

func main() {
	Run()
}
