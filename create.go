package okdesk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/srvchain/servicechain/pkg/metr"
	"io"
	"io/ioutil"
	"os"
)

var MetricCreate = metr.New("okdesk_create")

// POST https://<account>.okdesk.ru/api/v1/issues/?api_token=3e9a214215f493c4
var OKDESK_CREATE = "%s/api/v1/issues/"

type CreateRequest struct {
	Issue CreateIssue1 `json:"issue"`
}
type CreateIssue1 struct {
	Title                     string                 `json:"title"`
	Description               *string                `json:"description,omitempty"`
	CompanyID                 *string                `json:"company_id,omitempty"`
	ContactID                 *string                `json:"contact_id,omitempty"`
	AgreementID               *string                `json:"agreement_id,omitempty"`
	AssigneeID                *string                `json:"assignee_id,omitempty"`
	GroupID                   *string                `json:"group_id,omitempty"`
	MaintenanceEntityID       *string                `json:"maintenance_entity_id,omitempty"`
	EquipmentID               *string                `json:"equipment_ids,omitempty"`
	Type                      *string                `json:"type,omitempty"`
	Priority                  *string                `json:"priority,omitempty"`
	DeadlineAt                *string                `json:"deadline_at,omitempty"`
	StartExecutionUntil       *string                `json:"start_execution_until,omitempty"`
	PlannedExecutionInMinutes *int                   `json:"planned_execution_in_minutes,omitempty"`
	CustomParameters          map[string]interface{} `json:"custom_parameters,omitempty"`
	ParentID                  *string                `json:"parent_id,omitempty"`
	Author                    []User                 `json:"author,omitempty"`
}

type CreateResp struct {
	ID     int `json:"id"`
	Errors Err `json:"errors"`
}

type Err map[string]interface{}

func (Err Err) String() string {
	return fmt.Sprintf("%#v", Err)
}

func (issue CreateIssue) Request() (r io.ReadCloser, err error) {
	data, err := json.Marshal(CreateRequest{issue})
	if err != nil {
		return
	}
	return ioutil.NopCloser(bytes.NewReader(data)), nil
}

/**
file
curl  -H "Content-Type: multipart/form-data" -F "issue[title]=Не работает ноутбук" -F "issue[priority]=high"
-F "issue[attachments][0][attachment]=@/path/to/file.txt" -F "issue[attachments][0][is_public]=true"
-F "issue[attachments][0][description]=Условия гарантии"
-F "issue[attachments][1][attachment]=@/path/to/file.docx" -F "issue[attachments][1][is_public]=false"
https://<account>.okdesk.ru/api/v1/issues/{?api_token}
*/

//"comment[attachments][0][attachment]": io.ReadCloser, // lets assume its this file

func Create1(ctx context.Context, issue *CreateIssue) (resp *CreateResp, err error) {
	MetricCreate.Last.Inc()
	MetricCreate.Total.Inc()
	defer MetricCreate.ErrInc(err)

	Url := fmt.Sprintf(OKDESK_CREATE, os.Getenv("OKDESK_URL"))
	r, err := issue.Request()
	if err != nil {
		return
	}
	body, err := Request(ctx, "POST", Url, r, nil)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	defer body.Close()
	resp = &CreateResp{}
	return resp, json.Unmarshal(data, resp)
}

/**
file
curl  -H "Content-Type: multipart/form-data" -F "issue[title]=Не работает ноутбук" -F "issue[priority]=high"
-F "issue[attachments][0][attachment]=@/path/to/file.txt" -F "issue[attachments][0][is_public]=true"
-F "issue[attachments][0][description]=Условия гарантии"
-F "issue[attachments][1][attachment]=@/path/to/file.docx" -F "issue[attachments][1][is_public]=false"
https://<account>.okdesk.ru/api/v1/issues/{?api_token}
*/

func CreateWithAttach(ctx context.Context, issue *CreateIssue) (resp *CreateResp, err error) {
	MetricCreate.Last.Inc()
	MetricCreate.Total.Inc()
	defer MetricCreate.ErrInc(err)

	Url := fmt.Sprintf(OKDESK_CREATE, os.Getenv("OKDESK_URL"))
	r, err := issue.Request()
	if err != nil {
		return
	}
	body, err := Request(ctx, "POST", Url, r, nil)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	defer body.Close()
	resp = &CreateResp{}
	return resp, json.Unmarshal(data, resp)
}
