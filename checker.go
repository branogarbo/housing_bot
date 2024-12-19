package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"regexp"
)

func (b Bot) checkAddress(printNoMatch bool) error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

  reqOpts := parseCurlCmd(curlCmd)

	req, err := http.NewRequest(reqOpts.Method, reqOpts.URL, bytes.NewBufferString(reqOpts.Body)) 
	if err != nil {
		return err
	}

  for k,v := range reqOpts.Headers {
    req.Header.Set(k, v)
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

	err = b.handlePage(res, resData, printNoMatch)
	if err != nil {
		return err
	}

	return nil
}

func (b Bot) handlePage(res *http.Response, resData string, printNoMatch bool) error {
	var err error
	lastResponse = resData

	if res.StatusCode == 302 {
		err = b.authNeeded()
	} else if res.StatusCode != 200 {
		err = b.requestErrored(res)
	} else {
		err = b.checkResponseBody(resData, printNoMatch)
	}

	return err
}

func (b Bot) authNeeded() error {
	return b.notifyUser("Bot needs reauthentication!")
}

func (b Bot) requestErrored(res *http.Response) error {
	return b.notifyUser("Bot ran into a problem! Got a status of " + res.Status + ". Please check manually! " + linkPage)
}

func (b Bot) checkResponseBody(resBody string, printNoMatch bool) error {
  re := regexp.MustCompile(searchPattern)

  if re.MatchString(resBody) && !alertWhenFound {
		err := printToLog("Nothing found yet...")
		if err != nil {
			log.Fatal(err)
		}

		if printNoMatch {
			_, err = b.Session.ChannelMessageSend(channelID, "Nothing found yet...")
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}

	return b.notifyUser(foundMessage + " " + linkPage)
}
