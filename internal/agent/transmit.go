package agent

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type metric interface {
	getParam() string
	getValue() []byte
}

func transmitPlainText(m metric) {

	fmt.Println("http: " + m.getParam())

	client := &http.Client{}
	var resp *http.Response
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://127.0.0.1:8080/update/"+m.getParam(), bytes.NewBuffer(m.getValue()))
	if err != nil {
		fmt.Println("Err req: ", err)
	}
	req.Header.Add("Content-Type", "text/plain")
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Err on executing: ", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("Send success:", m.getParam())
	}
}

func sendReport(counters *[]counter, gauges *[]gauge) {
	for _, c := range *counters {
		transmitPlainText(c)
	}
	for _, g := range *gauges {
		transmitPlainText(g)
	}
}
