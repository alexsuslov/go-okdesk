# go-okdesk

## GetIssueCount
```
  sd:=OKD{}
  result:=[]int{}
  err := sd.GetIssueCount(context.Background(), &result)
  panic(err)
```