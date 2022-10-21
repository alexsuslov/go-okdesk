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

import "context"

const (
	entitiesPath          = "/maintenance_entities/"
	entityPath            = "/maintenance_entities/%1"
	entityListPath        = "/maintenance_entities/list"
	entityAttachmentsPath = "/maintenance_entities/%s/attachments/"
)

// https://okdesk.ru/apidoc#!obekty-obsluzhivaniya-poisk-obekta-obsluzhivaniya

func (OKD OKD) FindMaintenancesEntity(ctx context.Context, query map[string][]string, response interface{}) error {
	return OKD.ValRes("GET", entitiesPath)(ctx, query, response)
}

// https://okdesk.ru/apidoc#!obekty-obsluzhivaniya-sozdanie-obekta-obsluzhivaniya

func (OKD OKD) CreateMaintenanceEntity(ctx context.Context, IssueTimeEntry, response interface{}) error {
	return OKD.ReqRes("POST", entitiesPath)(ctx, IssueTimeEntry, response)
}

func (OKD OKD) SetMaintenanceEntity(ctx context.Context, MaintenanceEntityID, MaintenanceEntity, response interface{}) error {
	return OKD.IdReqRes("PATCH", entityPath)(ctx, MaintenanceEntityID, MaintenanceEntity, response)
}

// https://okdesk.ru/apidoc#!informacziya-ob-obekte-obsluzhivaniya-informacziya-ob-obekte-obsluzhivaniya

func (OKD OKD) GetMaintenanceEntity(ctx context.Context, MaintenanceEntityID, response interface{}) error {
	return OKD.IdRes("GET", entityPath)(ctx, MaintenanceEntityID, response)
}

//https://okdesk.ru/apidoc#!poluchenie-spiska-obektov-obsluzhivaniya-poluchenie-spiska-po-parametram

func (OKD OKD) GetMaintenanceEntityList(ctx context.Context, query map[string][]string, response interface{}) error {
	return OKD.ValRes("GET", entityListPath)(ctx, query, response)
}

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
