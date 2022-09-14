package main

import (
	"errors"
	"fmt"
	"github.com/SimFG/milvus-hook/pb/milvuspb"
	"log"
	"os"
	"plugin"
)

type Hook interface {
	Init(params map[string]interface{}) error
	Before(req interface{}) error
	After(result interface{}, err error) error
}

func main() {
	log.Println("main function start")

	p, err := plugin.Open("./hook.so")
	if err != nil {
		exit(err)
	}
	log.Println("plugin open")

	h, err := p.Lookup("Hook")
	if err != nil {
		exit(err)
	}

	hook, ok := h.(Hook)
	if !ok {
		exit(errors.New("err type"))
	}
	cr := &milvuspb.CreateCollectionRequest{
		DbName: "hello",
	}
	hook.Before(cr)
	// expected output: "result hello Hook"
	fmt.Println("result", cr.DbName, cr.CollectionName)

	dr := &milvuspb.DescribeCollectionRequest{}
	hook.Before(dr)
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
