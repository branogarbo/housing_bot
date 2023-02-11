package main

import (
	"fmt"
	"log"
	"time"
)

func Run() {
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

		time.Sleep(3 * time.Minute)
	}
}
