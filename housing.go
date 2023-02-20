package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (s SMSI) checkHousingPage() error {
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

	err = s.handlePage(res, pageHTML)
	if err != nil {
		return err
	}

	return nil
}

func (s SMSI) handlePage(res *http.Response, pageHTML string) error {
	var err error

	if res.StatusCode == 302 {
		err = s.authNeeded()
	} else if res.StatusCode != 200 {
		err = s.pageErrored()
	} else {
		err = s.checkPageHTML(pageHTML)
	}

	return err
}

func (s SMSI) authNeeded() error {
	return s.notifyUser("Housing Bot needs reauthentication!")
}

func (s SMSI) pageErrored() error {
	return s.notifyUser("Housing Bot ran into a problem fetching the housing page!")
}

func (s SMSI) checkPageHTML(resBody string) error {
	alertWhenFound, err := strconv.ParseBool(alertWhenFound)
	if err != nil {
		return err
	}

	if strings.Contains(resBody, htmlPattern) && !alertWhenFound {
		fmt.Println("No housing found yet...")

		return nil
	}

	return s.notifyUser("HOUSING IS AVAILABLE‼️‼️‼️")
}
