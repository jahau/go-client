package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/neovim/go-client/msgpack"
)

var (
	_ json.RawMessage
	_ fmt.State
)

func main() {
	mpack := os.Args[1]
	f, err := os.Open(mpack)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := msgpack.NewDecoder(f)
	var api API
	if err := dec.Decode(&api); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("api.NvimBufDelVar: %s\n", spew.Sdump(api))

	// data, err := json.MarshalIndent(api, "", "    ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Fprint(os.Stdout, string(data))

	// var apis map[string]interface{}
	// if err := json.Unmarshal(data, &apis); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("apis: %s\n", spew.Sdump(apis))
}
