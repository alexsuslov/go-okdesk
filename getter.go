package okdesk

var OKDESK_ISSUE = "%s/api/v1/issues/%v"

//func GetByID(ctx context.Context, issueID string) (issue *Issue, err error) {
//
//	Url := fmt.Sprintf(OKDESK_ISSUE, os.Getenv("OKDESK_URL"), issueID)
//
//	body, err := Request(ctx, "GET", Url, nil, nil)
//	if err != nil {
//		err = fmt.Errorf("Request:%v", err)
//		return
//	}
//	data, err := ioutil.ReadAll(body)
//	if err != nil {
//		return
//	}
//	defer body.Close()
//
//	issue = &Issue{}
//	err = json.Unmarshal(data, issue)
//	if err != nil {
//		err = fmt.Errorf("GetByID:%v", err)
//	}
//	if issue.Errors != nil {
//		err = fmt.Errorf(*issue.Errors)
//		return
//	}
//
//	if issue.CommentsInfo.Count > 0 {
//		var comments []CommentType
//		comments, err = GetCommentsByID(ctx, issueID)
//		if err != nil {
//			return
//		}
//		issue.Comments = comments
//	}
//
//	return issue, err
//}
