package main

import (
    "log"
    "net/http"
    "../common"
    "encoding/json"

  )

 func sayHi(w http.ResponseWriter, r *http.Request){
     res := common.JsonResp {Success: true, Msg: "You are cool"}
     resString, _ := json.Marshal(res)
     w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json")
    w.Write(resString)
}

func main() {
    initDB()

    http.HandleFunc("/", sayHi)
    http.HandleFunc("/user/new", newUserHandler)
    http.HandleFunc("/sentence/create", CreateSentenceHandler)

	go pushMsgToClient()

    log.Fatal(http.ListenAndServe(":8081", nil))
}
