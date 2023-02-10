package main

import "log"

func Run() {
	smsi, err := initTwilio()
	if err != nil {
		log.Fatal(err)
	}

	err = notifyUser(smsi, "AYYYYYYYYY")
	if err != nil {
		log.Fatal(err)
	}
}
