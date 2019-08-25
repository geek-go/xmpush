package xmpush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func DoPost(uri string, appSecret string, form *url.Values) ([]byte, error) {
	param := ""
	if form != nil {
		param = form.Encode()
	}

	//fmt.Println(strings.NewReader(param))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s", baseHost()+uri), strings.NewReader(param))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", appSecret))

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return body, nil
}

func DoGet(uri string, appSecret string, form *url.Values) ([]byte, error) {
	param := ""
	if form != nil {
		param = form.Encode()
	}

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s?%s", baseHost()+uri, param),
		nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", appSecret))

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return body, nil
}

func baseHost() string {
	return ProductionHost
}
