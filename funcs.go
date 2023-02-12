package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Run() {
	minutes, err := strconv.Atoi(os.Getenv("CHECK_INTERVAL"))
	if err != nil {
		log.Fatal(err)
	}

	checkInterval := time.Duration(minutes) * time.Minute

	fmt.Println("Running Housing Bot...")

	smsi, err := initTwilio()
	if err != nil {
		log.Fatal(err)
	}

	err = notifyUser(smsi, "Watching for available housing...")
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = checkHousingPage(smsi)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(checkInterval)
	}
}
