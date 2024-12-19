package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Run() {
	minutes, err := strconv.Atoi(checkIntervalStr)
	if err != nil {
		log.Fatal(err)
	}

	checkInterval = time.Duration(minutes) * time.Minute

	err = printToLog("Running Bot...")
	if err != nil {
		log.Fatal(err)
	}

	dg, err := initDiscord()
	if err != nil {
		log.Fatal(err)
	}
	defer dg.Session.Close()

	err = dg.notifyUser("Watching for pattern...")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		dg.initFiber()
	}()

	dg.startCheckingLoop()
}

func printToLog(msg string) error {
	t := time.Now()
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	entry := fmt.Sprintf("%s - (%s)", msg, t.In(loc))

	fmt.Println(entry)

	return nil
}

func (b Bot) startCheckingLoop() {
	for {
		if isChecking {
			go func() {
				err := b.checkAddress(false)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}

		time.Sleep(checkInterval)
	}
}

func parseCurlCmd(curlCmd string) HTTPoptions {
  headers := map[string]string{}
  opts := HTTPoptions{}

  curlLines := strings.Split(curlCmd, `\`)
  re := regexp.MustCompile(`([^\s]*) '(.*)'`)

  for _, line := range curlLines {
    match := re.FindStringSubmatch(line)
  
    switch match[1] {
    case "curl":
      opts.URL = match[2]
    case "--data":
      opts.Body = match[2]
    case "-H":
      headerSplit := strings.Split(match[2], ": ")
      headers[headerSplit[0]] = headerSplit[1]
    case "-X":
      opts.Method = match[2]  
    }
  }

  opts.Headers = headers
  return opts
}

