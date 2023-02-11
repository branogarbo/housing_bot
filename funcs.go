package main

import (
	"fmt"
	"log"
	"time"
)

func Run() {
	smsi, err := initTwilio()
	if err != nil {
		log.Fatal(err)
	}

	// defer notifyUser(smsi, "Shutting down") // do later

	err = notifyUser(smsi, "Watching for available housing...")
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = checkHousingPage(smsi)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(20 * time.Second)
		fmt.Println("for loop iterated")
	}
}
