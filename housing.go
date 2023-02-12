package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func checkHousingPage(s SMSI) error {
	pageURL := os.Getenv("PAGE_URL")
	pageCookie := os.Getenv("PAGE_COOKIE")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cookie", pageCookie)
	req.Header.Set("upgrade-insecure-requests", "1")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	pageHTML := string(resBody)

	err = handlePage(s, res, pageHTML)
	if err != nil {
		return err
	}

	return nil
}

func handlePage(s SMSI, res *http.Response, pageHTML string) error {
	var err error

	if res.StatusCode == 302 {
		err = authNeeded(s)
	} else if res.StatusCode != 200 {
		err = pageErrored(s)
	} else {
		err = checkPageHTML(s, pageHTML)
	}

	return err
}

func authNeeded(s SMSI) error {
	return notifyUser(s, "Housing Bot needs reauthentication!")
}

func pageErrored(s SMSI) error {
	return notifyUser(s, "Housing Bot ran into a problem fetching the housing page!")
}

func checkPageHTML(s SMSI, resBody string) error {
	if strings.Contains(resBody, "find any available rooms. Inventory is ever-changing, and as rooms become available, they will be displayed in the portal in real-time.") {
		fmt.Println("No housing found yet...")

		return nil
	}

	return notifyUser(s, "HOUSING IS AVAILABLE‼️‼️‼️")
}
