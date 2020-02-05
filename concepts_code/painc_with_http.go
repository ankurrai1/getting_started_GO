package main

import (
  "fmt"
  "net/http"
)

// NOTE: So if there is a defer then defer will called before painc 

func main()  {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("Hello world"))
  })
  err := http.ListenAndServe(":8080", nil)
  if err != nil{
    painc(err.Error())   // here I am telling that i don't know this error so I am retrowing this.
  }
}
