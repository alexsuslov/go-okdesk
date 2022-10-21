// (c) 2022 Alex Suslov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package okdesk

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Request(ctx context.Context, method, URL string,
	reader io.ReadCloser,
	header map[string]string) (body io.ReadCloser, err error) {

	U, err := url.Parse(URL)
	if err != nil {
		return
	}

	Q := U.Query()
	Q.Set("api_token", os.Getenv("OKDESK_TOKEN"))
	U.RawQuery = Q.Encode()

	req, err := http.NewRequestWithContext(ctx, method, U.String(), reader)
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: os.Getenv("OKDESK_INSECURE") == "YES"},
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

type Value struct {
	Name string
	Body io.ReadCloser
}

func MultiPartRequest(ctx context.Context, method, URL string,
	values map[string]Value) (body io.ReadCloser, err error) {

	U, err := url.Parse(URL)
	if err != nil {
		return
	}

	Q := U.Query()
	Q.Set("api_token", os.Getenv("OKDESK_TOKEN"))
	U.RawQuery = Q.Encode()

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	for k, v := range values {
		var fw io.Writer
		if v.Body != nil {
			defer v.Body.Close()
			fw, err = w.CreateFormFile(k, v.Name)
			_, err = io.Copy(fw, v.Body)
		} else {
			if fw, err = w.CreateFormField(k); err != nil {
				return
			}
			_, err = io.Copy(fw, strings.NewReader(v.Name))
		}
	}
	w.Close()
	req, err := http.NewRequestWithContext(ctx, method, U.String(), buf)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: os.Getenv("OKDESK_INSECURE") == "YES"},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		data, _ := io.ReadAll(res.Body)
		err = fmt.Errorf("%v:%v", res.StatusCode, string(data))
		return
	}

	return res.Body, err
}
