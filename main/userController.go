package main

import (
    "fmt"
    "net/http"
)

func newUserHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "name to sign up: " + r.FormValue("name"))
}