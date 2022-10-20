package okdesk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Request Request
func Request(ctx context.Context, method string, URL string,
	reader io.ReadCloser,
	header map[string]string) (body io.ReadCloser, err error) {

	U, err := url.Parse(URL)
	if err != nil {
		return
	}

	Q := U.Query()
	Q.Set("api_token", os.Getenv("OKDESK_TOKEN"))
	U.RawQuery = Q.Encode()

	tr := &http.Transport{
		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	URL = U.String()
	req, err := http.NewRequestWithContext(ctx, method, URL, reader)
	if err != nil {
		return
	}
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}

	} else {
		if reader != nil {
			req.Header.Set("Content-Type", "application/json")
		}
	}

	client := &http.Client{Transport: tr}
	r, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do:%v", err)
		return
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		data, _ := io.ReadAll(r.Body)
		err = fmt.Errorf("%v:%v", r.StatusCode, string(data))
		return
	}

	return r.Body, err
}
