package main

import (
	"encoding/json"
	"flag"
	"os"
	"fmt"
	"io/ioutil"
	"text/template"
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
	err = json.Unmarshal(bytes, &b)

	return &b, err
}

func makeFileFromTemplates(b Blueprint,
	outFile *os.File,
	tmplFilename string) (err error) {


	tmpl := template.New("john")
	tmpl, err = tmpl.ParseFiles(tmplFilename)

	if err != nil {
		return err
	}

	outFile, err = os.Create("out/out.go")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(outFile,"in.go.tmpl",b)

	if err != nil {
		return err
	}

	for _, r := range b.ResourceGroups {
		fmt.Printf("%#v", r.Name)
	}

	return err
}

func main() {
	flag.Parse()

	b, err := parse(apiFile)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Open("out/out.go")

	err = makeFileFromTemplates(*b,outFile,"in/in.go.tmpl")
	if err != nil {
		fmt.Printf("%#v",err)
	}
}
