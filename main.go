package main

import (
  "log"
  "fmt"
  "net/http"
  "golang.org/x/net/proxy"
  "io/ioutil"
  "encoding/json"
)

// Define type for config.json
type Config struct {
  Url     string    `json:"url"`
  ProxyPorts []string `json:"proxyPorts"`
}

func main() {
  // Load config
  raw, err := ioutil.ReadFile("./config.json")
  if err != nil {
      fmt.Println(err.Error())
      //os.Exit(1)
  }
  var c Config
  json.Unmarshal(raw, &c)

  client := http.DefaultClient

  for i, proxyPort := range c.ProxyPorts {
    if proxyPort != "" {
      p, _ := proxy.SOCKS5("tcp", "127.0.0.1:"+proxyPort, nil, proxy.Direct)
      client.Transport = &http.Transport{
          Dial: p.Dial,
      }  
    }
  
    resp, err := client.Get(c.Url)
  
    if err != nil {
      log.Fatal(err)
    }
  
    defer resp.Body.Close()
  
    byteArray, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(i+1, ": You are accessing from", string(byteArray), "via proxy port",  proxyPort)
  }
}