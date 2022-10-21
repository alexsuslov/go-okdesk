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
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
)

type OKD struct {
	host string
}

func New() OKD {
	return OKD{host: os.Getenv("OKDESK_URL")}
}

func (OKD OKD) ValRes(method, URL string) func(ctx context.Context, val map[string][]string, response interface{}) error {
	return func(ctx context.Context, val map[string][]string, response interface{}) error {
		U, err := url.Parse(OKD.host + URL)
		if err != nil {
			return err
		}
		Q := U.Query()
		for k, values := range val {
			for _, v := range values {
				Q.Add(k, v)
			}
		}

		body, err := Request(ctx, method, URL, nil, nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

func (OKD OKD) Res(method, URL string) func(ctx context.Context, response interface{}) error {
	return func(ctx context.Context, response interface{}) error {

		body, err := Request(ctx, method, OKD.host+URL, nil, nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

func (OKD OKD) ReqRes(method, URL string) func(ctx context.Context, request interface{}, response interface{}) error {
	return func(ctx context.Context, request interface{}, response interface{}) error {
		data, err := json.Marshal(request)
		if err != nil {
			return err
		}
		body, err := Request(ctx, method, OKD.host+URL, io.NopCloser(bytes.NewReader(data)), nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

func (OKD OKD) IdReqRes(method, path string) func(ctx context.Context, ID, request, response interface{}) error {
	return func(ctx context.Context, ID, request interface{}, response interface{}) error {
		URL := OKD.host + fmt.Sprintf(path, ID)
		data, err := json.Marshal(request)
		if err != nil {
			return err
		}
		body, err := Request(ctx, method, URL, io.NopCloser(bytes.NewReader(data)), nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

func (OKD OKD) IdIdReqRes(method, path string) func(ctx context.Context, ID, ID1, request, response interface{}) error {
	return func(ctx context.Context, ID, ID1, request interface{}, response interface{}) error {
		URL := OKD.host + fmt.Sprintf(path, ID, ID1)
		data, err := json.Marshal(request)
		if err != nil {
			return err
		}
		body, err := Request(ctx, method, URL, io.NopCloser(bytes.NewReader(data)), nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

func (OKD OKD) IdRes(method, path string) func(ctx context.Context, ID, response interface{}) error {
	return func(ctx context.Context, ID, response interface{}) (err error) {
		URL := OKD.host + fmt.Sprintf(path, ID)
		body, err := Request(ctx, method, URL, nil, nil)
		if err != nil {
			return
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}
func (OKD OKD) IdIdRes(method, path string) func(ctx context.Context, ID, ID1, response interface{}) error {
	return func(ctx context.Context, ID, ID1, response interface{}) (err error) {
		URL := OKD.host + fmt.Sprintf(path, ID, ID1)
		body, err := Request(ctx, method, URL, nil, nil)
		if err != nil {
			return
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

type ResponseError struct {
	Errors *struct {
		CompanyID             string `json:"company_id"`
		Type                  string `json:"type"`
		CustomParametersTotal string `json:"custom_parameters_total"`
	} `json:"errors,omitempty"`
}

func (OKD OKD) ReaderRes(method, path string) func(ctx context.Context, Form io.Reader, response interface{}) error {
	return func(ctx context.Context, Form io.Reader, response interface{}) error {
		URL := OKD.host + path
		body, err := Request(ctx, method, URL, io.NopCloser(Form), nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

func (OKD OKD) IdReaderRes(method, path string) func(ctx context.Context, ID interface{}, Form io.Reader, response interface{}) error {
	return func(ctx context.Context, ID interface{}, Form io.Reader, response interface{}) error {
		URL := OKD.host + path
		body, err := Request(ctx, method, URL, io.NopCloser(Form), nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}
