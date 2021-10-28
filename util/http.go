package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		Errorln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		Errorln(err)
	}
	result := string(body)
	Infoln(result)
	return result
}

func httpPost(url string, fields ...string) string {

	var s string = ""
	for i, v := range fields {
		val := ""
		if i%2 == 0 {
			val = v + "="
		} else {
			if i != len(fields)-1 {
				val = v + "&"
			} else {
				val = v
			}
		}
		s = s + val
	}

	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(s))
	if err != nil {
		Errorln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		Errorln(err)
	}
	return string(body)
}

func httpPostForm(urlstr string, values url.Values) string {
	//resp, err := http.PostForm(urlstr, url.Values{"key": {"Value"}, "id": {"123"}})
	resp, err := http.PostForm(urlstr, values)

	if err != nil {
		// handle error
		Errorln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		Errorln(err)
	}

	return string(body)
}

func httpDo(urlstr string, fields []string, headers []string) string {
	client := &http.Client{}

	var s string = ""
	for i, v := range fields {
		val := ""
		if i%2 == 0 {
			val = v + "="
		} else {
			if i != len(fields)-1 {
				val = v + "&"
			} else {
				val = v
			}
		}
		s = s + val
	}

	req, err := http.NewRequest("POST", urlstr, strings.NewReader(s))
	if err != nil {
		// handle error
		Errorln(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for i, _ := range headers {
		req.Header.Add(headers[i], headers[i+1])
		i = i + 1
	}
	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		Errorln(err)
	}

	return string(body)
}
