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
	"time"
)

const (
	issuePath            = "/issues/%s"
	issuesPath           = "/issues/"
	issuesCountPath      = "/issues/count"
	issueListPath        = "/issues/list"
	issuesPrioritiesPath = "/issues/priorities/"
	issuesStatusesPath   = "/issues/statuses"
	issuesParametersPath = "/issues/parameters/list"

	issueAssigneesPath   = "/issues/%v/assignees"
	issueDeadlinePath    = "/issues/%v/deadlines"
	issueTypePath        = "/issues/%v/types"
	issueParameterPath   = "/issues/%v/parameters"
	issueStatusPath      = "/issues/%v/statuses"
	issueCommentsPath    = "/issues/%v/comments"
	issueRatesPath       = "/issues/%v/rates"
	issueServicesPath    = "/issues/%v/services"
	issueTimeEntriesPath = "/issues/%v/time_entries"
	issueAttachmentsPath = "/issues/%v/attachments/%s"
	issueCheckListPath   = "/issues/%v/check_lists/items"
	issueCheckPath       = "/issues/%v/check_lists/items/%v/check"
	issueAvailableePath  = "/issues/%v/available_services"
)

// https://okdesk.ru/apidoc#!sozdanie-zayavki-sozdanie-zayavki

func (OKD OKD) CreateIssue(ctx context.Context, createIssueRequest interface{}, response interface{}) error {
	return OKD.ReqRes("POST", issuesPath)(ctx, createIssueRequest, response)
}

type createIssueRequest struct {
	Issue IssueCreate `json:"issue"`
}

type IssueCreate struct {
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	CompanyID           string   `json:"company_id"`
	ContactID           string   `json:"contact_id"`
	AgreementID         string   `json:"agreement_id"`
	Type                string   `json:"type"`
	Priority            string   `json:"priority"`
	MaintenanceEntityID string   `json:"maintenance_entity_id"`
	EquipmentIds        []string `json:"equipment_ids"`
	DeadlineAt          string   `json:"deadline_at"`
	CustomParameters    struct {
		Address string `json:"address"`
		Checked bool   `json:"checked"`
	} `json:"custom_parameters"`
	ParentID string `json:"parent_id"`
	Author   struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"author"`
	ObserverIds        []int `json:"observer_ids"`
	ObserverGroupIds   []int `json:"observer_group_ids"`
	ContactObserverIds []int `json:"contact_observer_ids"`
}

// https://okdesk.ru/apidoc#!smena-otvetstvennogo-smena-otvetstvennogo-za-zayavku

func (OKD OKD) SetIssueAssignee(ctx context.Context, issueID, createIssueRequest, response interface{}) error {
	return OKD.IdReqRes("PATCH", issuePath)(ctx, issueID, createIssueRequest, response)
}

// https://okdesk.ru/apidoc#!smena-planovoj-daty-resheniya-smena-planovoj-daty-resheniya-zayavki

func (OKD OKD) SetIssueDeadline(ctx context.Context, issueID, IssueDeadlineRequest, response interface{}) error {
	return OKD.IdReqRes("PATCH", issueDeadlinePath)(ctx, issueID, IssueDeadlineRequest, response)
}

//https://okdesk.ru/apidoc#!smena-tipa-zayavki-smena-tipa-zayavki

func (OKD OKD) SetIssueType(ctx context.Context, issueID, IssueTypeRequest, response interface{}) error {
	return OKD.IdReqRes("PATCH", issueTypePath)(ctx, issueID, IssueTypeRequest, response)
}

//https://okdesk.ru/apidoc#!redaktirovanie-dopolnitelnyh-atributov-zayavki-redaktirovanie-dopolnitelnyh-atributov-zayavki

func (OKD OKD) SetIssueStatus(ctx context.Context, issueID, IssueStatusRequest, response interface{}) error {
	return OKD.IdReqRes("PATCH", issueStatusPath)(ctx, issueID, IssueStatusRequest, response)
}

//https://okdesk.ru/apidoc#!kommentarii-dobavlenie-kommentariya

func (OKD OKD) CreateIssueComment(ctx context.Context, issueID, createCommentRequest, response interface{}) error {
	return OKD.IdReqRes("POST", issueCommentsPath)(ctx, issueID, createCommentRequest, response)
}

//https://okdesk.ru/apidoc#!kommentarii-poluchenie-spiska-kommentariev

func (OKD OKD) GetIssueComments(ctx context.Context, issueID, response interface{}) error {
	return OKD.IdRes("GET", issueCommentsPath)(ctx, issueID, response)
}

//https://okdesk.ru/apidoc#!poluchenie-speczifikaczij-zayavki-dobavlenie-speczifikaczii-k-zayavke

func (OKD OKD) CreateIssueServices(ctx context.Context, issueID, createService, response interface{}) error {
	return OKD.IdReqRes("POST", issueServicesPath)(ctx, issueID, createService, response)
}

// https://okdesk.ru/apidoc#!poluchenie-detalizaczii-po-trudozatratam-zayavki-poluchenie-detalizaczii-po-trudozatratam-zayavki

func (OKD OKD) GetIssueTimeEntries(ctx context.Context, issueID, response interface{}) error {
	return OKD.IdRes("GET", issueTimeEntriesPath)(ctx, issueID, response)
}

//https://okdesk.ru/apidoc#!poluchenie-detalizaczii-po-trudozatratam-zayavki-spisanie-trudozatrat-po-zayavke

func (OKD OKD) CreateIssueTimeEntry(ctx context.Context, issueID, IssueTimeEntry, response interface{}) error {
	return OKD.IdReqRes("POST", issueTimeEntriesPath)(ctx, issueID, IssueTimeEntry, response)
}

//https://okdesk.ru/apidoc#!poluchenie-fajla-zayavki-poluchenie-fajla-zayavki

func (OKD OKD) GetIssueAttachment(ctx context.Context, issueID, attachmentID, response interface{}) error {
	return OKD.IdIdRes("GET", issueAttachmentsPath)(ctx, issueID, attachmentID, response)
}

//https://okdesk.ru/apidoc#!poluchenie-chek-lista-zayavki-poluchenie-chek-lista-zayavki

func (OKD OKD) GetIssueCheckList(ctx context.Context, issueID, response interface{}) error {
	return OKD.IdRes("GET", issueCheckListPath)(ctx, issueID, response)
}

//https://okdesk.ru/apidoc#!pometka-o-vypolnenii-stroki-chek-lista-pometka-o-vypolnenii-stroki-chek-lista

func (OKD OKD) SetIssueCheckList(ctx context.Context, issueID, checkID, item, response interface{}) error {
	return OKD.IdIdReqRes("PATCH", issueCheckPath)(ctx, issueID, checkID, item, response)
}

// https://okdesk.ru/apidoc#!informacziya-o-zayavke-informacziya-o-zayavke

func (OKD OKD) GetIssue(ctx context.Context, issueID, response interface{}) error {
	return OKD.IdRes("GET", issuePath)(ctx, issueID, response)
}

func (OKD OKD) GetIssueParameters(ctx context.Context, response interface{}) error {
	return OKD.Res("GET", issuesParametersPath)(ctx, response)
}

func (OKD OKD) GetIssueCount(ctx context.Context, response interface{}) error {
	return OKD.Res("GET", issuesCountPath)(ctx, response)
}

func (OKD OKD) GetIssueList(ctx context.Context, response interface{}) error {
	return OKD.Res("GET", issueListPath)(ctx, response)
}

func (OKD OKD) GetIssuePriorities(ctx context.Context, response interface{}) error {
	return OKD.Res("GET", issuesPrioritiesPath)(ctx, response)
}

func (OKD OKD) GetIssueStatuses(ctx context.Context, response interface{}) error {
	return OKD.Res("GET", issuesStatusesPath)(ctx, response)
}

type RequestIssueAssignee struct {
	AssigneeID string `json:"assignee_id"`
	GroupID    string `json:"group_id"`
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

/**
file
curl  -H "Content-Type: multipart/form-data" -F "issue[title]=Не работает ноутбук" -F "issue[priority]=high"
-F "issue[attachments][0][attachment]=@/path/to/file.txt" -F "issue[attachments][0][is_public]=true"
-F "issue[attachments][0][description]=Условия гарантии"
-F "issue[attachments][1][attachment]=@/path/to/file.docx" -F "issue[attachments][1][is_public]=false"
https://<account>.okdesk.ru/api/v1/issues/{?api_token}
*/
// todo:

func (OKD OKD) CreateIssueWithAttach(ctx context.Context, createIssueRequest interface{}, response interface{}) error {

	return OKD.ReqRes("POST", issuesPath)(ctx, createIssueRequest, response)
}