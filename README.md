# go-okdesk
## Env
```
cp default.env .env
```
## Add okdesk token 
```
micro .env
```

## GetIssueCount
```
  sd:=OKD{}
  result:=[]int{}
  err := sd.GetIssueCount(context.Background(), &result)
  if err!=nil{
    panic(err)
  }
```
