package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var help bool
var apiFile string

func init() {
	flag.BoolVar(&help, "help", false, "print help")
	flag.StringVar(&apiFile, "api-file", "api.json",
		"file to read blueprint from")
}

func parse(apiFile string) (*Blueprint, error) {
	bytes, err := ioutil.ReadFile(apiFile)

	if err != nil {
		return nil, err
	}

	b := Blueprint{}
	err = json.Unmarshal(bytes,&b)

	return &b,err
}

func main() {
	flag.Parse()

	b,err := parse(apiFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v",b)
}
