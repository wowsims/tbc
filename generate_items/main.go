package main

import (
	"flag"
	"github.com/wowsims/tbc/sim/api"
	//"fmt"
)

func main() {
	outDir := flag.String("outDir", "", "Path to output directory for writing generated .go files.")
	flag.Parse()

	if *outDir == "" {
		panic("outDir flag is required!")
	}

	itemResponses := make([]WowheadItemResponse, len(ItemDeclarations))
	for idx, itemDeclaration := range ItemDeclarations {
		itemResponse := getWowheadItemResponse(itemDeclaration.ID)
		//fmt.Printf("\n\n%+v\n", itemResponse)
		itemResponses[idx] = itemResponse
	}

	for specVal, _ := range api.Spec_name {
		spec := api.Spec(specVal)
		writeItemFile(*outDir, ItemDeclarations, itemResponses, &spec)
	}
	writeItemFile(*outDir, ItemDeclarations, itemResponses, nil)
}
