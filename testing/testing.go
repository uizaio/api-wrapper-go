package main

import (
  uiza "github.com/uizaio/api-wrapper-go"
  "github.com/uizaio/api-wrapper-go/live"
  "log"
)

func init() {
  uiza.Authorization = "uap-212f8ac7dcc7471c936babf43a1a252e-57078be1" // test key, input your own
}

func main() {
  params := &uiza.LiveRetrieveParams{ID: uiza.String("471734b1-90b8-44c5-b24b-132ed9d7529b")}
  response, err := live.Retrieve(params)
  if err != nil {
    log.Printf("%v\n", err)
  } else {
    log.Printf("%v\n", response.LinkStream)
  }
}