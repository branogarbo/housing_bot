package main

import (
	"io/ioutil"
	"net/http"
)

func fetchHousingPage() (string, error) {
	req, err := http.NewRequest("GET", "https://asu.starrezhousing.com/StarRezPortalX/7DAB39A4/28/2135/My_Housing-Choose_My_Community?UrlToken=7055E5A5&TermID=169&ClassificationID=5&DateStart=Saturday%2C%20August%2012%2C%202023&DateEnd=Saturday%2C%20May%204%2C%202024", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("authority", "asu.starrezhousing.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", "starrezportalx_si_StarRezPortalX=31ertcohfswmml4h5t1l4czp; __RequestVerificationToken_L1N0YXJSZXpQb3J0YWxY0=9pJYupK8rDxsB2-Fy0Qk_hY24kIE65tSbcR4Iav0Lb_tbms9XGb4mSbnC9SSV6hpq39qS13q8qfBCQgmRqS8A70b7Nch9vUtoavlD5SqdE0zVmmfcYVw05-WiR17_n8_S2iQMRgf-MqsCw4RDtXPrg2; ptx-affinity-starrezportalx=50c6b81ecf84a12a5c874bdb55ad4fea|e6d3b938a55660fc6ab4000278f2e158; _ga=GA1.2.1622416837.1675978372; _gid=GA1.2.1576614861.1675978372; _ga=GA1.3.1622416837.1675978372; _gid=GA1.3.1576614861.1675978372; _gat=1; _gat_PortalX=1; _gat_portalxacc=1")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://asu.starrezhousing.com/StarRezPortalX/71AF9C9C/28/2035/My_Housing-About_Me?UrlToken=7055E5A5&TermID=169&ClassificationID=5")
	req.Header.Set("sec-ch-ua", `Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "macOS")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

// curl 'https://asu.starrezhousing.com/StarRezPortalX/7DAB39A4/28/2135/My_Housing-Choose_My_Community?UrlToken=7055E5A5&TermID=169&ClassificationID=5&DateStart=Saturday%2C%20August%2012%2C%202023&DateEnd=Saturday%2C%20May%204%2C%202024' \
//   -H 'authority: asu.starrezhousing.com' \
//   -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8' \
//   -H 'accept-language: en-US,en;q=0.9' \
//   -H 'cache-control: no-cache' \
//   -H 'cookie: starrezportalx_si_StarRezPortalX=31ertcohfswmml4h5t1l4czp; __RequestVerificationToken_L1N0YXJSZXpQb3J0YWxY0=9pJYupK8rDxsB2-Fy0Qk_hY24kIE65tSbcR4Iav0Lb_tbms9XGb4mSbnC9SSV6hpq39qS13q8qfBCQgmRqS8A70b7Nch9vUtoavlD5SqdE0zVmmfcYVw05-WiR17_n8_S2iQMRgf-MqsCw4RDtXPrg2; ptx-affinity-starrezportalx=50c6b81ecf84a12a5c874bdb55ad4fea|e6d3b938a55660fc6ab4000278f2e158; _ga=GA1.2.1622416837.1675978372; _gid=GA1.2.1576614861.1675978372; _ga=GA1.3.1622416837.1675978372; _gid=GA1.3.1576614861.1675978372; _gat=1; _gat_PortalX=1; _gat_portalxacc=1' \
//   -H 'pragma: no-cache' \
//   -H 'referer: https://asu.starrezhousing.com/StarRezPortalX/71AF9C9C/28/2035/My_Housing-About_Me?UrlToken=7055E5A5&TermID=169&ClassificationID=5' \
//   -H 'sec-ch-ua: "Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "macOS"' \
//   -H 'sec-fetch-dest: document' \
//   -H 'sec-fetch-mode: navigate' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'sec-fetch-user: ?1' \
//   -H 'sec-gpc: 1' \
//   -H 'upgrade-insecure-requests: 1' \
//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36' \
//   --compressed
