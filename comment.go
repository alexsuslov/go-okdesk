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
	"os"
)

func GetIssueComments(ctx context.Context, issueID string, comments []CommentType) error {
	URL := os.Getenv("OKDESK_URL") + fmt.Sprintf(issueCommentsPath, issueID)
	body, err := Request(ctx, "GET", URL, nil, nil)
	if err != nil {
		return err
	}
	defer body.Close()
	return json.NewDecoder(body).Decode(comments)
}

func CreateIssueComment(ctx context.Context, issueID string, comment IssueComment, response *ResponseComment) error {
	URL := os.Getenv("OKDESK_URL") + fmt.Sprintf(issueCommentsPath, issueID)
	data, err := json.Marshal(struct {
		Comment IssueComment `json:"comment"`
	}{
		comment,
	})
	if err != nil {
		return err
	}

	body, err := Request(ctx, "POST", URL, io.NopCloser(bytes.NewReader(data)), nil)
	if err != nil {
		return err
	}
	defer body.Close()
	return json.NewDecoder(body).Decode(response)
}

type IssueComment struct {
	Content    string  `json:"content"`
	AuthorID   *int    `json:"author_id,omitempty"`
	AuthorType *string `json:"author_type,omitempty"`
	Public     bool    `json:"public,omitempty"`
}

type ResponseComment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Public  bool   `json:"public"`
	Author  User   `json:"author"`
	Errors  *Err   `json:"errors,omitempty"`
}
