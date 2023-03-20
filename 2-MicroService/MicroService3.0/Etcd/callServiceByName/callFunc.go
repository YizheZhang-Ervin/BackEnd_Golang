package main

import (
	"context"
	"fmt"
	"goetcd/util"
	"log"
)

func main() {
	client := util.NewClient()
	err := client.LoadService()
	if err != nil {
		log.Fatal(err)
	}
	endpoint := client.GetService("productservice", "GET", serivces.ProdEncodeFunc)
	res, err := endpoint(context.Background(), serivces.ProdRequest{ProdId: 106})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
