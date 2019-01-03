package utils

import (
    "fmt"
    "strconv"
    "net/http"
    "github.com/qiniu/api.v7/auth/qbox"
    "github.com/qiniu/api.v7/storage"
)
type Day struct{
    x int
}
func Gettoken(w http.ResponseWriter,r *http.Request)  {
    r.ParseForm() //解析参数，默认是不会解析的
    d,_:=strconv.Atoi(r.Form["x"][0])
    buckets:=[7]string{"myone","mytwo","mythree","myfour","myfive","mysix","myseven"}
    accessKey := "qEjtRsk220SrirJdrxY6UxBpFW-sbRaLOGzscjT1"
    secretKey := "CTLSJV4Hwe8Xq6gJaUzotYWe5qBv8qly3tWlnxnA"
    putPolicy := storage.PutPolicy{
        Scope: buckets[d],
    }
    mac := qbox.NewMac(accessKey, secretKey)
    upToken := putPolicy.UploadToken(mac)
    fmt.Println(upToken)
    w.Write([]byte(upToken))
}
