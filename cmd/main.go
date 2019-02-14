package main

import (
	uiza "api-wrapper-go"
	"api-wrapper-go/entity"
	_ "api-wrapper-go/testing"
	"log"
)

func main() {
	params := &uiza.EntityRetrieveParams{ID: uiza.String("")}
	response, _ := entity.Retrieve(params)
	log.Printf("%s\n", response)
}
