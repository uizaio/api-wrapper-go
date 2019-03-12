package main

import (
	"log"

	uiza "github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/entity"
	_ "github.com/uizaio/api-wrapper-go/testing"
)

func main() {
	params := &uiza.EntityDeleteParams{ID: uiza.String("f7c07eae-4911-45d5-b687-2ba063d2b223")}
	response, err := entity.Delete(params)
	log.Printf("%v %v\n", response, err)
}
