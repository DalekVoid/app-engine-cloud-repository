package main

import (
        "net/http"
        "fmt"
)

func init() {
        http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "You can view the source code here: https://github.com/DalekVoid/app-engine-cloud-repository/")
}
