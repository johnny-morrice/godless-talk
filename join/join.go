package main

import (
	"fmt"

	"github.com/johnny-morrice/godless"
	"github.com/johnny-morrice/godless/api"
	"github.com/johnny-morrice/godless/query"
)

func main() {
	options := godless.Options{

		KeyStore: godless.MakeKeyStore(),
	}

	godless, err := godless.New(options)
	dieOnError(err)

	query, err := query.Compile("join cars rows (key=@car1, driver=?)", "Mr Speedy")
	dieOnError(err)

	response, err := godless.Send(api.MakeQueryRequest(query))
	dieOnError(err)

	fmt.Println(response)
}

func dieOnError(err error) {
	if err != nil {
		panic(err)
	}
}
