package okdesk

import (
	"encoding/json"
	"net/url"
)

var OKDESK_COUNT = "%s/api/v1/issues/count"

type CountParam struct {
	CustomParameters map[string]interface{} `json:"custom_parameters"`
}

func (CountParam CountParam) Request(vals url.Values) error {
	if CountParam.CustomParameters == nil {
		return nil
	}
	data, err := json.Marshal(CountParam.CustomParameters)
	if err != nil {
		return err
	}
	vals.Set("custom_parameters", string(data))
	return nil
}

type CountResp struct {
	Errors Err `json:"errors"`
}

//func GetIssuesCount(ctx context.Context, req *CountParam, resp *CountResp) (err error) {
//	URL := os.Getenv("OKDESK_URL") + issuesCountPath
//
//	Url := fmt.Sprintf(OKDESK_COUNT, os.Getenv("OKDESK_URL"))
//	U, err := url.Parse(Url)
//	if err != nil {
//		return
//	}
//	Q := url.Values{}
//	err = req.Request(Q)
//	if err != nil {
//		return
//	}
//	U.RawQuery = Q.Encode()
//
//	body, err := Request(ctx, "GET", U.String(), nil, nil)
//	if err != nil {
//		return
//	}
//	data, err := ioutil.ReadAll(body)
//	if err != nil {
//		return
//	}
//	defer body.Close()
//	resp = &CountResp{}
//	return resp, json.Unmarshal(data, resp)
//}
