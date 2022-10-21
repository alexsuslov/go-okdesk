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
	"context"
	"github.com/alexsuslov/godotenv"
	"testing"
)

func load() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func TestOKD_GetIssueCount(t *testing.T) {
	load()
	type fields struct {
		host string
	}
	type args struct {
		ctx      context.Context
		response interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"GetIssueCount",
			fields{godotenv.GetPanic("OKDESK_URL")},
			args{
				context.Background(),
				&[]int{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OKD := OKD{
				host: tt.fields.host,
			}
			err := OKD.GetIssueCount(tt.args.ctx, tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIssueCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
