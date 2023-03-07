package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (d Discord) checkHousingPage() error {
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

	err = d.handlePage(res, pageHTML)
	if err != nil {
		return err
	}

	return nil
}

func (d Discord) handlePage(res *http.Response, pageHTML string) error {
	var err error

	if res.StatusCode == 302 {
		err = d.authNeeded()
	} else if res.StatusCode != 200 {
		err = d.pageErrored(res)
	} else {
		err = d.checkPageHTML(pageHTML)
	}

	return err
}

func (d Discord) authNeeded() error {
	return d.notifyUser("Housing Bot needs reauthentication!")
}

func (d Discord) pageErrored(res *http.Response) error {
	return d.notifyUser("Housing Bot ran into a problem fetching the housing page! Got a status of " + res.Status)
}

func (d Discord) checkPageHTML(resBody string) error {
	alertWhenFound, err := strconv.ParseBool(alertWhenFound)
	if err != nil {
		return err
	}

	if strings.Contains(resBody, htmlPattern) && !alertWhenFound {
		err = printToLog("No housing found yet...")
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	return d.notifyUser("HOUSING IS AVAILABLE‼️‼️‼️")
}
