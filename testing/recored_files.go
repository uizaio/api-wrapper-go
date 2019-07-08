package main

import (
  uiza "github.com/uizaio/api-wrapper-go"
  "github.com/uizaio/api-wrapper-go/live"
  "log"
)

func init() {
  uiza.Authorization = "uap-9521cff34e864730"
}

func main() {

  params := &uiza.LiveListRecordedParams{
    Limit: 9,
    Page: 140,
    OrderBy: uiza.String("createdAt"), // not supported on server side yet
    OrderType: uiza.Desc, // not supported on server side yet
  }
  response, err := live.ListRecorded(params)
  if err != nil {
    log.Printf("%v\n", err)
  } else {
    log.Println("length: ", len(response))
  }
}