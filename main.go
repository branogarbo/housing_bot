package main

import (
	"os"
	"strconv"
	"time"
)

var (
	discordToken         = os.Getenv("DISCORD_TOKEN")
	channelID            = os.Getenv("CHANNEL_ID")
	checkIntervalStr     = os.Getenv("CHECK_INTERVAL")
	searchPattern        = os.Getenv("SEARCH_PATTERN")
	alertWhenFound, _    = strconv.ParseBool(os.Getenv("ALERT_WHEN_FOUND"))
	timezone             = os.Getenv("TIMEZONE")
	linkPage             = os.Getenv("LINK_PAGE")
  curlCmd = os.Getenv("CURL_CMD")
  foundMessage = os.Getenv("FOUND_MESSAGE")
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
