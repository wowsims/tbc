// +build wasm

package main

import (
	"log"
	"syscall/js"

	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/api/papi"
	"google.golang.org/protobuf/proto"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("computeStats", js.FuncOf(computeStats))
	js.Global().Set("gearList", js.FuncOf(gearList))
	js.Global().Set("individualSim", js.FuncOf(individualSim))
	js.Global().Set("statWeights", js.FuncOf(statWeights))
	js.Global().Call("wasmready")
	<-c
}

func computeStats(this js.Value, args []js.Value) interface{} {
	// Assumes args[0] is a Uint8Array
	data := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(data, args[0])

	csr := &api.ComputeStatsRequest{}
	if err := proto.Unmarshal(data, csr); err != nil {
		log.Printf("Failed to parse request: %s", err)
		return nil
	}
	result := papi.ComputeStats(csr)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		return nil
	}

	outArray := js.Global().Get("Uint8Array").New(len(outbytes))
	js.CopyBytesToJS(outArray, outbytes)

	return outArray
}

func gearList(this js.Value, args []js.Value) interface{} {
	// Assumes args[0] is a Uint8Array
	data := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(data, args[0])

	glr := &api.GearListRequest{}
	if err := proto.Unmarshal(data, glr); err != nil {
		log.Printf("Failed to parse request: %s", err)
		return nil
	}
	result := papi.GetGearList(glr)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		return nil
	}

	outArray := js.Global().Get("Uint8Array").New(len(outbytes))
	js.CopyBytesToJS(outArray, outbytes)

	return outArray
}

func individualSim(this js.Value, args []js.Value) interface{} {
	// Assumes args[0] is a Uint8Array
	data := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(data, args[0])

	isr := &api.IndividualSimRequest{}
	if err := proto.Unmarshal(data, isr); err != nil {
		log.Printf("Failed to parse request: %s", err)
		return nil
	}
	result := papi.RunSimulation(isr)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		return nil
	}

	outArray := js.Global().Get("Uint8Array").New(len(outbytes))
	js.CopyBytesToJS(outArray, outbytes)

	return outArray
}

func statWeights(this js.Value, args []js.Value) interface{} {
	// Assumes args[0] is a Uint8Array
	data := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(data, args[0])

	swr := &api.StatWeightsRequest{}
	if err := proto.Unmarshal(data, swr); err != nil {
		log.Printf("Failed to parse request: %s", err)
		return nil
	}
	result := papi.StatWeights(swr)

	outbytes, err := proto.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal result: %s", err.Error())
		return nil
	}

	outArray := js.Global().Get("Uint8Array").New(len(outbytes))
	js.CopyBytesToJS(outArray, outbytes)

	return outArray
}
