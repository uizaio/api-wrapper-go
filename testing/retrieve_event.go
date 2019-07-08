package main

import (
  uiza "github.com/uizaio/api-wrapper-go"
  "github.com/uizaio/api-wrapper-go/live"
  "log"
)

func init() {
  uiza.Authorization = "uap-9521cff34e86473095644ba71cbd0e7f"
}

func main() {

  resp, error := live.Retrieve(&uiza.LiveRetrieveParams{
    ID: uiza.String("070187a0-8ea5-4cb3-bcae-9d870dc39d50"),
  })

  if error != nil {
    log.Printf("%v\n", error)
  }

  log.Println("retrieve Live: ", resp)

}