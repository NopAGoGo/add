package main

import (
	"fmt"

	"github.com/sunday9th/add/pkg/addendpoint"
	"github.com/sunday9th/add/pkg/addservice"
)

func main() {

	var (
		service   = addservice.New()
		endpoints = addendpoint.New(service)
	)

	rq := addendpoint.SumRequest{1, 2}

	Sum := endpoints.SumEndpoint

	rp, err := Sum(nil, rq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rp.(addendpoint.SumResponse).V)
}
