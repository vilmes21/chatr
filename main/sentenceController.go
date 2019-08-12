package main

import (
    // "net/http"
    // "fmt"
    // "strconv"
    // "time"
    // "encoding/json"
    // "../common"
    
)

// func CreateSentenceHandler(w http.ResponseWriter, r *http.Request){
//     err := r.ParseForm()
//     if err != nil {
//         fmt.Println(fmt.Errorf("Error: %v", err))
//         w.WriteHeader(http.StatusInternalServerError)
//         return
//     }

//     s := Sentence{Time: time.Now()}
//     s.Content = r.Form.Get("content")

//     id, _ := strconv.Atoi(r.Form.Get("chatSpeakerId"))
//     s.ChatSpeakerId = id

//     res := common.JsonResp{}
    
// 	err = store.CreateSentence(&s)
// 	if err != nil {
//         fmt.Println(err)
//         res.Msg = "Failed"
//     } else {
//         res.Success = true
//     }
    
//     resString, err := json.Marshal(res)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     w.WriteHeader(200)
//     w.Header().Set("Content-Type", "application/json")
//     w.Write(resString)
// }

