package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (b Bot) checkHousingPage(printNoHouse bool) error {
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

	err = b.handlePage(res, pageHTML, printNoHouse)
	if err != nil {
		return err
	}

	return nil
}

func (b Bot) handlePage(res *http.Response, pageHTML string, printNoHouse bool) error {
	var err error

	if res.StatusCode == 302 {
		err = b.authNeeded()
	} else if res.StatusCode != 200 {
		err = b.pageErrored(res)
	} else {
		err = b.checkPageHTML(pageHTML, printNoHouse)
	}

	return err
}

func (b Bot) authNeeded() error {
	return b.notifyUser("Housing Bot needs reauthentication!")
}

func (b Bot) pageErrored(res *http.Response) error {
	return b.notifyUser("Housing Bot ran into a problem fetching the housing page! Got a status of " + res.Status)
}

func (b Bot) checkPageHTML(resBody string, printNoHouse bool) error {
	alertWhenFound, err := strconv.ParseBool(alertWhenFound)
	if err != nil {
		return err
	}

	if strings.Contains(resBody, htmlPattern) && !alertWhenFound {
		err = printToLog("No housing found yet...")
		if err != nil {
			log.Fatal(err)
		}

		if printNoHouse {
			_, err = b.Session.ChannelMessageSend(channelID, "No housing found yet...")
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}

	return b.notifyUser("HOUSING IS AVAILABLE‼️‼️‼️")
}
