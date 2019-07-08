package main

import (
  uiza "github.com/uizaio/api-wrapper-go"
  "github.com/uizaio/api-wrapper-go/live"
  "log"
)

func init() {
  uiza.Authorization = "uap-44d29caea2b74d938d0ec11963434dd8-"
}

func main() {
  modeType := uiza.ModeTypePull
  encodeType := uiza.EncodeTypeOne
  dvrType := uiza.DvrTypeOne
  resourceMode := uiza.ResourceModeSingle

  params := &uiza.LiveCreateParams{
    Name: uiza.String("test event"),
    Mode: &modeType,
    Encode: &encodeType,
    Dvr: &dvrType,
    Description: uiza.String("This is for test event"),
    Thumbnail: uiza.String("//image1.jpeg"),
    Region: "vn-southeast-1",
    Poster: uiza.String("//image1.jpeg"),
    LinkStream: []*string{uiza.String("https://playlist.m3u8")},
    ResourceMode: &resourceMode,
  }
  response, err := live.Create(params)
  if err != nil {
    log.Printf("%v\n", err)
  } else {
    log.Printf("%v\n", response)
  }
}