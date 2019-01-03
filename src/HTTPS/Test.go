package main

import (
    "net/http"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprintf(w, "hello, https")
}

func main() {
    http.HandleFunc("/",handler)
    http.ListenAndServeTLS(":443","1_ijavascript.club_bundle.crt","2_ijavascript.club.key", nil)
}
