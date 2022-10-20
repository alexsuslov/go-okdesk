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

type IssueStatus struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Assignee struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
