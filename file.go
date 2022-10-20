package okdesk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"

	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// /api/v1/issues/{issue_id}/attachments/{attachment_id}{?api_token}

var OKDESK_ATTACHMENT = "%s/api/v1/issues/%s/attachments/%s"

// HandleFileFunc /?issue_id={}&attachment_id={}
func HandleFileFunc(accessToken string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		split := strings.Split(r.Header.Get("Authorization"), " ")

		if len(split) != 2 ||
			strings.ToLower(split[0]) != "bearer" ||
			split[1] != accessToken {
			http.Error(w, "error token", http.StatusUnauthorized)
			return
		}
		err := Attach(w, r.URL.Query())
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func getAttachment(issueID string, attachmentID string) (issueAttachment IssueAttachment, err error) {
	URL := fmt.Sprintf(OKDESK_ATTACHMENT,
		os.Getenv("OKDESK_URL"),
		issueID,
		attachmentID)
	body, err := Request(context.Background(), "GET", URL, nil, nil)
	defer body.Close()
	issueAttachment = IssueAttachment{}
	return issueAttachment, json.NewDecoder(body).Decode(&issueAttachment)
}

func getAttachmentFile(issueID string, attachmentID string) (filename string, body io.ReadCloser, err error) {

	Attachment, err := getAttachment(issueID, attachmentID)
	if err != nil {
		return
	}
	time.Sleep(5 * time.Second)
	filename = Attachment.FileName
	body, err = Request(context.Background(), "GET", Attachment.Url, nil, nil)
	return
}

func getAttachmentFileAnonymous(issueID string, attachmentID string) (filename string, body io.ReadCloser, err error) {

	Attachment, err := getAttachment(issueID, attachmentID)
	if err != nil {
		return
	}
	//time.Sleep(5* time.Second)
	filename = Attachment.FileName
	res, err := http.Get(Attachment.Url)
	if err != nil {
		return
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		data, _ := ioutil.ReadAll(res.Body)
		err = fmt.Errorf("status:%v", res.Status)
		logrus.
			WithField("data", string(data)).
			WithField("Url", Attachment.Url).
			WithField("Size", Attachment.Size).
			Error(fmt.Errorf("getAttachmentFileAnonymous:%v", err.Error()))
		return
	}

	body = res.Body
	return
}

func Attach(w http.ResponseWriter, q url.Values) error {
	FileName, body, err := getAttachmentFileAnonymous(q.Get("issue_id"), q.Get("attachment_id"))
	if err != nil {
		return err
	}

	defer body.Close()
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%s\"", FileName))

	_, err = io.Copy(w, body)
	if err != nil {
		logrus.
			WithField("err", err).
			Error("file upload error")
	}

	return err

}

type IssueAttachment struct {
	ID          int    `json:"id"`
	FileName    string `json:"attachment_file_name"`
	Size        int    `json:"attachment_file_size"`
	Url         string `json:"attachment_url"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
	CreatedAt   string `json:"created_at"`
}

// https://<account>.okdesk.ru/api/v1/issues/{issue_id}/comments/{?api_token}

var OKDESK_CREATE_Comment = "%s/api/v1/issues/%v/comments/"

/**
fields := map[string]io.Reader{
        "comment": strings.NewReader("hello world!"),
		"author_id": strings.NewReader("{id}"),
        "comment[attachments][0][attachment]": io.ReadCloser, // lets assume its this file
        "comment[attachments][0][attachment]": io.ReadCloser, // lets assume its this file
    }

*/

func CreateCommentAttach(issueID string, fields map[string]File) (
	body io.ReadCloser, err error) {

	remoteURL := fmt.Sprintf(OKDESK_CREATE_Comment, os.Getenv("OKDESK_URL"), issueID)

	client := http.Client{Timeout: 120 * time.Second}

	return Upload(&client, remoteURL, fields)

}

type File struct {
	Name string
	Body io.ReadCloser
}

// GetFile GetFile
func GetFile(ctx context.Context, URL string,
	header map[string]string) (body io.ReadCloser, err error) {
	method := "GET"

	tr := &http.Transport{
		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req, err := http.NewRequestWithContext(ctx, method, URL, nil)
	if err != nil {
		return
	}
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{Transport: tr}
	r, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do:%v", err)
		return
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		data, _ := ioutil.ReadAll(r.Body)
		err = fmt.Errorf("status:%v messge=%v", r.Status, string(data))
		return
	}

	return r.Body, err
}
