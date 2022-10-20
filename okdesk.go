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
	"time"
)

type OKD struct {
	host                 string
	GetIssue             fnIdRes
	SetIssueAssignee     fnIdReqRes
	CreateIssue          fnReqRes
	GetIssueParameters   fnRes
	GetIssueCount        fnRes
	GetIssueList         fnRes
	GetIssuePriorities   fnRes
	GetIssueStatuses     fnRes
	SetIssueDeadline     fnIdReqRes
	SetIssueType         fnIdReqRes
	SetIssueParameter    fnIdReqRes
	SetIssueStatus       fnIdReqRes
	CreateIssueComment   fnIdReqRes
	GetIssueComments     fnIdRes
	CreateIssueServices  fnIdReqRes
	GetIssueTimeEntries  fnIdRes
	CreateIssueTimeEntry fnIdReqRes
	GetIssueAttachment   fnIdIdRes
	GetIssueCheckList    fnIdRes
	SetIssueCheckList    fnIdIdReqRes

	GetMaintenanceEntity     fnIdRes
	FindMaintenancesEntity   fnValRes
	CreateMaintenanceEntity  fnReqRes
	SetMaintenanceEntity     fnIdReqRes
	GetMaintenanceEntityList fnValRes
}

func New() OKD {
	host := os.Getenv("OKDESK_URL")
	okd := OKD{host: host}
	// https://okdesk.ru/apidoc#!sozdanie-zayavki-sozdanie-zayavki
	okd.CreateIssue = ReqRes("POST", host+issuesPath)

	// https://okdesk.ru/apidoc#!smena-otvetstvennogo-smena-otvetstvennogo-za-zayavku
	okd.SetIssueAssignee = IdReqRes("PATCH", host+issuePath)

	// https://okdesk.ru/apidoc#!smena-planovoj-daty-resheniya-smena-planovoj-daty-resheniya-zayavki
	okd.SetIssueDeadline = IdReqRes("PATCH", host+issueDeadlinePath)

	//https://okdesk.ru/apidoc#!smena-tipa-zayavki-smena-tipa-zayavki
	okd.SetIssueType = IdReqRes("PATCH", host+issueTypePath)

	//https://okdesk.ru/apidoc#!redaktirovanie-dopolnitelnyh-atributov-zayavki-redaktirovanie-dopolnitelnyh-atributov-zayavki
	okd.SetIssueType = IdReqRes("POST", host+issueParameterPath)

	//https://okdesk.ru/apidoc#!status-zayavki-smena-statusa-zayavki
	okd.SetIssueStatus = IdReqRes("POST", host+issueStatusPath)

	//https://okdesk.ru/apidoc#!kommentarii-dobavlenie-kommentariya
	okd.CreateIssueComment = IdReqRes("POST", host+issueCommentsPath)

	//https://okdesk.ru/apidoc#!kommentarii-poluchenie-spiska-kommentariev
	okd.GetIssueComments = IdRes("GET", host+issueCommentsPath)

	//https://okdesk.ru/apidoc#!poluchenie-speczifikaczij-zayavki-dobavlenie-speczifikaczii-k-zayavke
	okd.CreateIssueServices = IdReqRes("POST", host+issueServicesPath)

	// https://okdesk.ru/apidoc#!poluchenie-detalizaczii-po-trudozatratam-zayavki-poluchenie-detalizaczii-po-trudozatratam-zayavki
	okd.GetIssueTimeEntries = IdRes("GET", host+issueTimeEntriesPath)

	//https://okdesk.ru/apidoc#!poluchenie-detalizaczii-po-trudozatratam-zayavki-spisanie-trudozatrat-po-zayavke
	okd.CreateIssueTimeEntry = IdReqRes("POST", host+issueTimeEntriesPath)

	//https://okdesk.ru/apidoc#!poluchenie-fajla-zayavki-poluchenie-fajla-zayavki
	okd.GetIssueAttachment = IdIdRes("GET", host+issueAttachmentsPath)

	//https://okdesk.ru/apidoc#!poluchenie-chek-lista-zayavki-poluchenie-chek-lista-zayavki
	okd.GetIssueCheckList = IdRes("GET", host+issueCheckListPath)

	//https://okdesk.ru/apidoc#!pometka-o-vypolnenii-stroki-chek-lista-pometka-o-vypolnenii-stroki-chek-lista
	okd.SetIssueCheckList = IdIdReqRes("PATCH", host+issueCheckPath)

	// https://okdesk.ru/apidoc#!informacziya-o-zayavke-informacziya-o-zayavke
	okd.GetIssue = IdRes("GET", host+issuePath)

	// https://okdesk.ru/apidoc#!obekty-obsluzhivaniya-poisk-obekta-obsluzhivaniya
	okd.FindMaintenancesEntity = ValRes("GET", host+entitiesPath)

	// https://okdesk.ru/apidoc#!obekty-obsluzhivaniya-sozdanie-obekta-obsluzhivaniya
	okd.CreateMaintenanceEntity = ReqRes("POST", host+entitiesPath)

	// https://okdesk.ru/apidoc#!obekty-obsluzhivaniya-redaktirovanie-obekta-obsluzhivaniya
	okd.SetMaintenanceEntity = IdReqRes("PATCH", host+entityPath)

	//https://okdesk.ru/apidoc#!informacziya-ob-obekte-obsluzhivaniya-informacziya-ob-obekte-obsluzhivaniya
	okd.GetMaintenanceEntity = IdRes("GET", host+entityPath)

	//https://okdesk.ru/apidoc#!poluchenie-spiska-obektov-obsluzhivaniya-poluchenie-spiska-po-parametram
	okd.GetMaintenanceEntityList = ValRes("GET", host+entityListPath)

	// https://okdesk.ru/apidoc#!dobavleniya-fajla-k-obektu-obsluzhivaniya-dobavleniya-fajla-k-obektu-obsluzhivaniya
	/** todo: Добавления файла к объекту обслуживания

	curl  -H "Content-Type: multipart/form-data"
	-F "maintenance_entity[attachments][0][attachment]=@/home/user/file.jpg"
	-F "maintenance_entity[attachments][0][is_public]=true"
	-F "maintenance_entity[attachments][0][description]=Описание"
	https://<account>.okdesk.ru/api/v1/maintenance_entities/{id}/attachments/{?api_token}

	где @/home/user/file.jpg — это путь к локальному файлу
	пример для windows: @"C:/myfile.txt"

	*/

	okd.GetIssueParameters = Res("GET", host+issuesParametersPath)
	okd.GetIssueCount = Res("GET", host+issuesCountPath)
	okd.GetIssueList = Res("GET", host+issueListPath)
	okd.GetIssuePriorities = Res("GET", host+issuesPrioritiesPath)
	okd.GetIssueStatuses = Res("GET", host+issuesStatusesPath)
	return okd
}

type fnValRes func(ctx context.Context, val map[string][]string, response interface{}) error
type fnIdRes func(ctx context.Context, ID string, response interface{}) error
type fnIdIdRes func(ctx context.Context, ID, ID1 string, response interface{}) error
type fnReqRes func(ctx context.Context, request interface{}, response interface{}) error
type fnIdReqRes func(ctx context.Context, ID string, request interface{}, response interface{}) error
type fnIdIdReqRes func(ctx context.Context, ID, ID1 string, request interface{}, response interface{}) error
type fnRes func(ctx context.Context, response interface{}) (err error)

func ValRes(method, URL string) func(ctx context.Context, val map[string][]string, response interface{}) error {
	return func(ctx context.Context, val map[string][]string, response interface{}) error {
		U, err := url.Parse(URL)
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

func Res(method, URL string) func(ctx context.Context, response interface{}) error {
	return func(ctx context.Context, response interface{}) error {

		body, err := Request(ctx, method, URL, nil, nil)
		if err != nil {
			return err
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

func ReqRes(method, URL string) func(ctx context.Context, request interface{}, response interface{}) error {
	return func(ctx context.Context, request interface{}, response interface{}) error {
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

func IdReqRes(method, path string) func(ctx context.Context, ID string, request interface{}, response interface{}) error {
	return func(ctx context.Context, ID string, request interface{}, response interface{}) error {
		URL := fmt.Sprintf(path, ID)
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

func IdIdReqRes(method, path string) func(ctx context.Context, ID, ID1 string, request interface{}, response interface{}) error {
	return func(ctx context.Context, ID, ID1 string, request interface{}, response interface{}) error {
		URL := fmt.Sprintf(path, ID, ID1)
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

func IdRes(method, path string) func(ctx context.Context, ID string, response interface{}) error {
	return func(ctx context.Context, ID string, response interface{}) (err error) {
		URL := fmt.Sprintf(path, ID)
		body, err := Request(ctx, method, URL, nil, nil)
		if err != nil {
			return
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}
func IdIdRes(method, path string) func(ctx context.Context, ID, ID1 string, response interface{}) error {
	return func(ctx context.Context, ID, ID1 string, response interface{}) (err error) {
		URL := fmt.Sprintf(path, ID, ID1)
		body, err := Request(ctx, method, URL, nil, nil)
		if err != nil {
			return
		}
		defer body.Close()
		return json.NewDecoder(body).Decode(response)
	}
}

type RequestIssueAssignee struct {
	AssigneeID string `json:"assignee_id"`
	GroupID    string `json:"group_id"`
}

type ResponseError struct {
	Errors *struct {
		CompanyID             string `json:"company_id"`
		Type                  string `json:"type"`
		CustomParametersTotal string `json:"custom_parameters_total"`
	} `json:"errors,omitempty"`
}

type Issue struct {
	ID                      int           `json:"id"`
	Title                   string        `json:"title"`
	Description             string        `json:"description"`
	CreatedAt               time.Time     `json:"created_at"`
	CompletedAt             time.Time     `json:"completed_at"`
	DeadlineAt              time.Time     `json:"deadline_at"`
	Source                  string        `json:"source"`
	SpentTimeTotal          int           `json:"spent_time_total"`
	StartExecutionUntil     time.Time     `json:"start_execution_until"`
	PlannedExecutionInHours int           `json:"planned_execution_in_hours"`
	PlannedReactionAt       time.Time     `json:"planned_reaction_at"`
	ReactedAt               time.Time     `json:"reacted_at"`
	UpdatedAt               time.Time     `json:"updated_at"`
	DelayedTo               interface{}   `json:"delayed_to"`
	CompanyID               int           `json:"company_id"`
	GroupID                 int           `json:"group_id"`
	Coexecutors             []Executor    `json:"coexecutors"`
	ServiceObjectID         interface{}   `json:"service_object_id"`
	EquipmentIds            []interface{} `json:"equipment_ids"`
	Attachments             []Attachment  `json:"attachments"`
	StatusTimes             StatusTime    `json:"status_times"`
	Parameters              []interface{} `json:"parameters"`
	Comments                CommentStatus `json:"comments"`
	ParentID                interface{}   `json:"parent_id"`
	ChildIds                []interface{} `json:"child_ids"`
	Type                    Type          `json:"type"`
	Priority                Named         `json:"priority"`
	Status                  Named         `json:"status"`
	OldStatus               Named         `json:"old_status"`
	Rate                    Valued        `json:"rate"`
	Observers               []Named       `json:"observers"`
	ObserverGroups          []Named       `json:"observer_groups"`
	Contact                 Named         `json:"contact"`
	Agreement               interface{}   `json:"agreement"`
	Assignee                Named         `json:"assignee"`
	Author                  Named         `json:"author"`
}

type Valued struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Named struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type CommentStatus struct {
	Count  int       `json:"count"`
	LastAt time.Time `json:"last_at"`
}

type StatusTime struct {
	Opened    TimeStatus `json:"opened"`
	Completed TimeStatus `json:"completed"`
}

type TimeStatus struct {
	Total           string `json:"total"`
	OnScheduleTotal string `json:"on_schedule_total"`
}

type Type struct {
	ID                 int    `json:"id"`
	Code               string `json:"code"`
	Name               string `json:"name"`
	AvailableForClient bool   `json:"available_for_client"`
}

type Attachment struct {
	ID                 int       `json:"id"`
	AttachmentFileName string    `json:"attachment_file_name"`
	Description        string    `json:"description"`
	AttachmentFileSize int       `json:"attachment_file_size"`
	IsPublic           bool      `json:"is_public"`
	CreatedAt          time.Time `json:"created_at"`
}

type Executor struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Group Group  `json:"group"`
}

type Group struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

type CommentType struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	Public      bool   `json:"public"`
	PublishedAt string `json:"published_at"`
	Author      User   `json:"author"`
}

type User struct {
	ID   int     `json:"id"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}
