package main

import (
    "encoding/json"
	"net/http"
)

func main() {
    http.HandleFunc("/gettest", getFunc)
    http.ListenAndServe(":1234", nil)
}
func getFunc(response http.ResponseWriter, request *http.Request) {
    JsonResponse(response, struct{First string
                                  Second string}{
                                      "username", 
                                      "password",
                                    })
}
func postFunc(response http.ResponseWriter, request *http.Request) {
    request.ParseForm()
    request.Form.Get("")
}
func JsonResponse(w http.ResponseWriter, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        println(err)
    }
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(response)
}
