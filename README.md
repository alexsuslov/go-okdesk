# go-okdesk

## GetIssueCount
```
  sd:=OKD{}
  result:=[]int{}
  err := sd.GetIssueCount(context.Background(), &result)
  if err!=nil{
    panic(err)
  }
```
