package main

import (
	"os"
	"time"
)

var (
	pageURL          = os.Getenv("PAGE_URL")
	pageCookie       = os.Getenv("PAGE_COOKIE")
	discordToken     = os.Getenv("DISCORD_TOKEN")
	channelID        = os.Getenv("CHANNEL_ID")
	checkIntervalStr = os.Getenv("CHECK_INTERVAL")
	htmlPattern      = os.Getenv("HTML_PATTERN")
	alertWhenFound   = os.Getenv("ALERT_WHEN_FOUND")
	timezone         = os.Getenv("TIMEZONE")
)

var (
	checkInterval time.Duration
	possibleCmds  = []string{"check", "help", "stop", "start"}
	isChecking    = true
)

func main() {
	Run()
}
