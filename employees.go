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
	employeesPath           = "/employees/"
	employeesListPath       = "/employees/list"
	employerPath            = "/employees/%s"
	employerActivationsPath = "/employees/%s/activations"
	employerRolesPath       = "/employees/%s/roles"
	employerGroupsPath      = "/employees/%s/groups"
	employerRoutesPath      = "/employees/routes"
)

func (OKD OKD) CreateEmployer(ctx context.Context, CreateEmployer interface{}, response interface{}) error {
	return OKD.ReqRes("POST", employeesPath)(ctx, CreateEmployer, response)
}

func (OKD OKD) SetEmployer(ctx context.Context,
	EmployerID string, UpdateEmployer interface{}, response interface{}) error {

	return OKD.IdReqRes("PATCH", employerPath)(ctx, EmployerID, UpdateEmployer, response)
}

func (OKD OKD) SetEmployerActivation(ctx context.Context,
	EmployerID string, UpdateActivation UpdateActivation, response interface{}) error {

	return OKD.IdReqRes("PATCH", employerActivationsPath)(ctx, EmployerID, UpdateActivation, response)
}

type UpdateActivation struct {
	Active bool `json:"active"`
}

func (OKD OKD) SetEmployerRoles(ctx context.Context, response interface{}) error {

	return OKD.Res("GET", employerRolesPath)(ctx, response)
}

func (OKD OKD) SetEmployerRoutes(ctx context.Context, response interface{}) error {

	return OKD.Res("GET", employerRoutesPath)(ctx, response)
}

func (OKD OKD) SetEmployerGroups(ctx context.Context, response interface{}) error {

	return OKD.Res("GET", employerGroupsPath)(ctx, response)
}

func (OKD OKD) SetEmployerList(ctx context.Context, query map[string][]string, response interface{}) error {

	return OKD.ValRes("GET", employeesListPath)(ctx, query, response)
}

type ResponseEmployerList []struct {
	ID              int    `json:"id"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	Patronymic      string `json:"patronymic"`
	Position        string `json:"position"`
	Active          bool   `json:"active"`
	Email           string `json:"email"`
	Login           string `json:"login"`
	Phone           string `json:"phone"`
	TelephonyNumber string `json:"telephony_number"`
	Comment         string `json:"comment"`
	Groups          []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
	Roles []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"roles"`
}
