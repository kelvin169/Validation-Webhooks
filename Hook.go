package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "log"
        "github.com/sendgrid/sendgrid-go"

)

func Parse (w http.ResponseWriter, req *http.Request) {
        //Get Email Values
        to := req.FormValue("from")
        subject := req.FormValue("subject")
        body:= req.FormValue("text")

        //Get Uploaded File
        file, handler, err := req.FormFile("attachment1")
        if err != nil {
                fmt.Println(err)
        }
        data, err := ioutil.ReadAll(file)
        if err != nil {
                fmt.Println(err)
        }
        err = ioutil.WriteFile(handler.Filename, data, 0777)
        if err != nil {
                fmt.Println(err)
        }
}

func main() {
        http.HandleFunc("/upload", Parse)
        err := http.ListenAndServe(":3000", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}