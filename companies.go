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
	companiesPath          = "/companies/"
	companyPath            = "/companies/%s"
	companyListPath        = "/companies/list"
	companyAttachmentsPath = "/companies/%s/attachments/%s"
)

func (OKD OKD) FindCompany(ctx context.Context, query map[string][]string, response interface{}) error {
	return ValRes("GET", OKD.host+companiesPath)(ctx, query, response)
}

type ResponseFindCompany struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	AdditionalName string    `json:"additional_name"`
	Site           string    `json:"site"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Crm1CID        string    `json:"crm_1c_id"`
	Address        string    `json:"address"`
	Comment        string    `json:"comment"`
	Coordinates    []float64 `json:"coordinates"`
	Observers      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"observers"`
	Contacts        []interface{} `json:"contacts"`
	DefaultAssignee struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"default_assignee"`
	Category struct {
		ID    int    `json:"id"`
		Code  string `json:"code"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"category"`
	Attachments []struct {
		ID                 int       `json:"id"`
		AttachmentFileName string    `json:"attachment_file_name"`
		Description        string    `json:"description"`
		AttachmentFileSize int       `json:"attachment_file_size"`
		IsPublic           bool      `json:"is_public"`
		CreatedAt          time.Time `json:"created_at"`
	} `json:"attachments"`
	Parameters []struct {
		Code      string `json:"code"`
		Name      string `json:"name"`
		FieldType string `json:"field_type"`
		Value     string `json:"value"`
	} `json:"parameters"`
}

func (OKD OKD) CreateCompany(ctx context.Context, CreateCompany interface{}, response interface{}) error {
	return ReqRes("POST", companiesPath)(ctx, CreateCompany, response)
}

type CreateCompany struct {
	Company struct {
		Name              string    `json:"name"`
		AdditionalName    string    `json:"additional_name"`
		Site              string    `json:"site"`
		Email             string    `json:"email"`
		Phone             string    `json:"phone"`
		Address           string    `json:"address"`
		Coordinates       []float64 `json:"coordinates"`
		Comment           string    `json:"comment"`
		ObserverIds       []int     `json:"observer_ids"`
		DefaultAssigneeID int       `json:"default_assignee_id"`
		CategoryCode      string    `json:"category_code"`
		Crm1CID           string    `json:"crm_1c_id"`
		CustomParameters  struct {
			Code1 string `json:"code1"`
			Code2 string `json:"code2"`
			Code3 bool   `json:"code3"`
		} `json:"custom_parameters"`
	} `json:"company"`
}

type ResponseCreateCompany struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	AdditionalName string    `json:"additional_name"`
	Site           string    `json:"site"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Comment        string    `json:"comment"`
	Crm1CID        string    `json:"crm_1c_id"`
	Coordinates    []float64 `json:"coordinates"`
	Observers      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"observers"`
	Contacts        []interface{} `json:"contacts"`
	DefaultAssignee struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"default_assignee"`
	Category struct {
		ID    int    `json:"id"`
		Code  string `json:"code"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"category"`
	Attachments []interface{} `json:"attachments"`
	Parameters  []struct {
		Code      string `json:"code"`
		Name      string `json:"name"`
		FieldType string `json:"field_type"`
		Value     string `json:"value"`
	} `json:"parameters"`
}

func (OKD OKD) SetCompany(ctx context.Context,
	CompanyID string, UpdateCompany interface{}, response interface{}) error {

	return IdReqRes("POST", companyPath)(ctx, CompanyID, UpdateCompany, response)
}

type UpdateCompany struct {
	Company struct {
		Name              string    `json:"name"`
		AdditionalName    string    `json:"additional_name"`
		Site              string    `json:"site"`
		Email             string    `json:"email"`
		Phone             string    `json:"phone"`
		Address           string    `json:"address"`
		Coordinates       []float64 `json:"coordinates"`
		Comment           string    `json:"comment"`
		ObserverIds       []int     `json:"observer_ids"`
		DefaultAssigneeID int       `json:"default_assignee_id"`
		CategoryCode      string    `json:"category_code"`
		Crm1CID           string    `json:"crm_1c_id"`
		CustomParameters  struct {
			Code2 string `json:"code2"`
		} `json:"custom_parameters"`
	} `json:"company"`
}

type ResponseUpdateCompany struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	AdditionalName string    `json:"additional_name"`
	Site           string    `json:"site"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Comment        string    `json:"comment"`
	Crm1CID        string    `json:"crm_1c_id"`
	Coordinates    []float64 `json:"coordinates"`
	Observers      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"observers"`
	Contacts        []interface{} `json:"contacts"`
	DefaultAssignee struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"default_assignee"`
	Category struct {
		ID    int    `json:"id"`
		Code  string `json:"code"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"category"`
	Attachments []struct {
		ID                 int       `json:"id"`
		AttachmentFileName string    `json:"attachment_file_name"`
		Description        string    `json:"description"`
		AttachmentFileSize int       `json:"attachment_file_size"`
		IsPublic           bool      `json:"is_public"`
		CreatedAt          time.Time `json:"created_at"`
	} `json:"attachments"`
	Parameters []struct {
		Code      string `json:"code"`
		Name      string `json:"name"`
		FieldType string `json:"field_type"`
		Value     string `json:"value"`
	} `json:"parameters"`
}

func (OKD OKD) GetCompanyList(ctx context.Context,
	Query map[string][]string, response []interface{}) error {

	return ValRes("POST", companyListPath)(ctx, Query, response)
}

type ResponseCompany struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	AdditionalName interface{} `json:"additional_name"`
	Parameters     []struct {
		Code      string `json:"code"`
		Name      string `json:"name"`
		FieldType string `json:"field_type"`
		Value     string `json:"value"`
	} `json:"parameters"`
}

func (OKD OKD) GetCompanyAttachment(ctx context.Context,
	companyID, attachmentID string, response []interface{}) error {

	return IdIdRes("GET", companyAttachmentsPath)(ctx, companyID, attachmentID, response)
}

type ResponseCompanyAttachment struct {
	ID                 int       `json:"id"`
	AttachmentFileName string    `json:"attachment_file_name"`
	Description        string    `json:"description"`
	AttachmentFileSize int       `json:"attachment_file_size"`
	IsPublic           bool      `json:"is_public"`
	CreatedAt          time.Time `json:"created_at"`
	AttachmentURL      string    `json:"attachment_url"`
}
