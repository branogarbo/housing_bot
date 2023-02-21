package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func Run() {
	minutes, err := strconv.Atoi(checkInterval)
	if err != nil {
		log.Fatal(err)
	}

	checkInterval := time.Duration(minutes) * time.Minute

	err = printToLog("Running Housing Bot...")
	if err != nil {
		log.Fatal(err)
	}

	smsi, err := initTwilio()
	if err != nil {
		log.Fatal(err)
	}

	err = smsi.notifyUser("Watching for available housing...")
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = smsi.checkHousingPage()
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(checkInterval)
	}
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
