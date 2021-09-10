// +build wasm
package main

import (
	"log"
	"syscall/js"

	"github.com/wowsims/tbc/api"
	"google.golang.org/protobuf/proto"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("individualSim", js.FuncOf(individualSim))
	js.Global().Call("wasmready")
	<-c
}

func individualSim(this js.Value, args []js.Value) interface{} {
	data := make([]byte, args[0].Call("length").Int())
	// Assumes input is a JSON object as a string
	js.CopyBytesToGo(data, args[0])

	isr := &api.IndividualSimRequest{}
	if err := proto.Unmarshal(data, isr); err != nil {
		log.Printf("Failed to parse request: %s", err)
		return nil
	}
	result := api.RunSimulation(isr)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		return nil
	}

	// TODO: do I need to create a new Uint8Array and js.CopyBytesToJS?
	return outbytes
}
