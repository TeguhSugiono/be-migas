package services

import (
	"BackendEsp32/config"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendWA(message string) error {

	token := config.GetEnv("FONNTE_TOKEN")
	target := config.GetEnv("FONNTE_TARGET")

	form := url.Values{}
	form.Set("target", target)
	form.Set("message", message)

	req, err := http.NewRequest(
		"POST",
		"https://api.fonnte.com/send",
		bytes.NewBufferString(form.Encode()),
	)

	if err != nil {
		return err
	}

	req.Header.Set(
		"Authorization",
		token,
	)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("FONNTE RESPONSE:")
	fmt.Println(string(body))

	return nil
}
