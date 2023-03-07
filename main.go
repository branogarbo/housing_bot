package main

import (
	"os"
)

var (
	pageURL        = os.Getenv("PAGE_URL")
	pageCookie     = os.Getenv("PAGE_COOKIE")
	discordToken   = os.Getenv("DISCORD_TOKEN")
	channelID      = os.Getenv("CHANNEL_ID")
	checkInterval  = os.Getenv("CHECK_INTERVAL")
	htmlPattern    = os.Getenv("HTML_PATTERN")
	alertWhenFound = os.Getenv("ALERT_WHEN_FOUND")
	timezone       = os.Getenv("TIMEZONE")
)

func main() {
	Run()
}
