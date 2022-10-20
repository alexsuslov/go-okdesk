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

func TestGetIssueParameters(t *testing.T) {
	godotenv.Load(".env")
	type args struct {
		ctx        context.Context
		parameters []IssueParameter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get parametrs",
			args{
				context.Background(),
				[]IssueParameter{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetIssueParameters(tt.args.ctx, tt.args.parameters); (err != nil) != tt.wantErr {
				t.Errorf("GetIssueParameters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetIssues(t *testing.T) {
	godotenv.Load(".env")
	type args struct {
		ctx        context.Context
		Parameters *[]int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get issues",
			args{context.Background(), &[]int{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GetIssueCount(tt.args.ctx, tt.args.Parameters)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIssues() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetIssueList(t *testing.T) {
	godotenv.Load(".env")
	type args struct {
		ctx      context.Context
		Response *[]IssueListItem
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get issues",
			args{context.Background(), &[]IssueListItem{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GetIssueList(tt.args.ctx, tt.args.Response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIssueList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetIssuePriorities(t *testing.T) {
	godotenv.Load(".env")
	type args struct {
		ctx      context.Context
		Response *[]IssuePriority
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get Priorities",
			args{context.Background(), &[]IssuePriority{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GetIssuePriorities(tt.args.ctx, tt.args.Response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIssuePriorities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetIssueStatuses(t *testing.T) {
	godotenv.Load(".env")
	type args struct {
		ctx      context.Context
		Response *[]IssueStatus
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"get Priorities",
			args{context.Background(), &[]IssueStatus{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GetIssueStatuses(tt.args.ctx, tt.args.Response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIssueStatuses() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
