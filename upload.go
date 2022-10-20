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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Upload(client *http.Client, remoteURL string, values map[string]File) (
	body io.ReadCloser, err error) {

	U, err := url.Parse(remoteURL)
	if err != nil {
		return
	}

	Q := U.Query()
	Q.Set("api_token", os.Getenv("OKDESK_TOKEN"))
	U.RawQuery = Q.Encode()

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	for key, f := range values {
		var fw io.Writer
		if f.Body != nil {

			defer f.Body.Close()

			fw, err = w.CreateFormFile(key, f.Name)

			_, err = io.Copy(fw, f.Body)

		} else {

			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
			_, err = io.Copy(fw, strings.NewReader(f.Name))
		}

	}

	w.Close()

	req, err := http.NewRequest("POST", U.String(), buf)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		data, _ := ioutil.ReadAll(res.Body)
		err = errors.New(string(data))
		err = fmt.Errorf("status:%v messge=%v", res.Status, string(data))
		return
	}

	return res.Body, err
}
