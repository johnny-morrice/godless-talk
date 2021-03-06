package main

import (
	"fmt"

	"github.com/johnny-morrice/godless/api"
	"github.com/johnny-morrice/godless/http"
	"github.com/johnny-morrice/godless/query"
)

func main() {
	options := http.ClientOptions{
		ServerAddr: "http://localhost:8085",
	}
	client, err := http.MakeClient(options)
	dieOnError(err)
	joinQuery(client)
}

func joinQuery(client api.Client) {
	query, err := query.Compile("join cars rows (@key=car1, driver=?)", "Mr Speedy")
	dieOnError(err)
	response, err := client.Send(api.MakeQueryRequest(query))
	dieOnError(err)
	fmt.Println(response)
}

func dieOnError(err error) {
	if err != nil {
		panic(err)
	}
}
