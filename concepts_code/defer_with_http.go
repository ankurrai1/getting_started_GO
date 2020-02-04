// use case of differ to under stand where all defer is useful
package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  "log"
)

func main()  {
  res, err := http.Get("http://www.google.com/robots.txt")
  if err != nil{
    log.Fatal(err)
  }
  defer res.Body.Close() // here it will be helpfull because you never forget to close near open and it will execute at the end of functions
  robots,err := ioutil.ReadAll(res.Body)
  if err != nil{
    log.Fatal(err)
  }
  fmt.Printf("%s",robots)

}
