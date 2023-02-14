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
