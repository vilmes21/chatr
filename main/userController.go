package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    // "time"
    "github.com/dgrijalva/jwt-go"
    "../keys"
)

func newUserHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "name to sign up: " + r.FormValue("name"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Create a new token object, specifying signing method and the claims
    // you would like it to contain.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": 8888,
    })

    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString(keys.JwtSecret)

    fmt.Println(tokenString, err)
}

type TokenStruct struct {
    Token string `json:"token"`
}

func login2Handler(w http.ResponseWriter, r *http.Request) {
    // r.ParseForm()

    // token := r.FormValue("token")

    decoder := json.NewDecoder(r.Body)
    var t TokenStruct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }

    fmt.Println(`token after POST : `, t.Token)
    claim, _ := authenticateJWT(t.Token)

    //test
    resBytes, _ := json.Marshal(claim)
    w.Header().Set("Content-Type", "application/json")
    w.Write(resBytes)
}

type JwtClaim struct {
    UserId int `json:"userId"`
    jwt.StandardClaims
}

func authenticateJWT(tokenString string) (*JwtClaim, error) {
    //func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error)
    claim := JwtClaim{}
    
    token, err := jwt.ParseWithClaims(
        tokenString, 
        &claim, 
        func(token *jwt.Token) (interface{}, error) {
        return keys.JwtSecret, nil
    })

    if claims, ok := token.Claims.(*JwtClaim); ok && token.Valid {
        // fmt.Printf("userId is: %v", claims.UserId)
        return claims, err
    } else {
        fmt.Println("YO err:", err)
        return nil, err
    }

    
}