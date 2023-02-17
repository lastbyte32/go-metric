package agent

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

func transmitPlainText(url string, value string, timeout time.Duration) error {

	fmt.Println(url)

	client := resty.New().
		SetTimeout(timeout)

	_, err := client.R().
		SetHeader("Content-Type", "text/plain").
		SetBody(value).
		Post(url)
	if err != nil {
		return err
	}

	return nil
}
