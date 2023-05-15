package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
)

func (b Bot) checkAddress(printNoHouse bool) error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(reqMethod, reqURL, bytes.NewBufferString(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cookie", reqCookies)
	req.Header.Set("upgrade-insecure-requests", "1")

	if reqVerificationToken != "" {
		req.Header.Set("__requestverificationtoken", reqVerificationToken)
	}

	if isAPIendpoint {
		req.Header.Set("content-type", "application/json; charset=UTF-8")
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resData := string(resBody)

	err = b.handlePage(res, resData, printNoHouse)
	if err != nil {
		return err
	}

	return nil
}

func (b Bot) handlePage(res *http.Response, resData string, printNoHouse bool) error {
	var err error
	lastResponse = resData

	if res.StatusCode == 302 {
		err = b.authNeeded()
	} else if res.StatusCode != 200 {
		err = b.requestErrored(res)
	} else {
		err = b.checkResponseBody(resData, printNoHouse)
	}

	return err
}

func (b Bot) authNeeded() error {
	return b.notifyUser("Housing Bot needs reauthentication!")
}

func (b Bot) requestErrored(res *http.Response) error {
	return b.notifyUser("Housing Bot ran into a problem! Got a status of " + res.Status + ". Please check housing manually! " + linkPage)
}

func (b Bot) checkResponseBody(resBody string, printNoHouse bool) error {
	if strings.Contains(resBody, searchPattern) && !alertWhenFound {
		err := printToLog("No housing found yet...")
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

	return b.notifyUser("HOUSING IS AVAILABLE‼️‼️‼️ " + linkPage)
}
