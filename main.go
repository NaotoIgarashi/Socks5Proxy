package main

import (
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
)


func main() {
  url := "http://inet-ip.info/ip"

  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray)) // htmlをstringで取得 
}