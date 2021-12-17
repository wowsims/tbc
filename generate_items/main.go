package main

import (
	"flag"
)

func main() {
	outDir := flag.String("outDir", "", "Path to output directory for writing generated .go files.")
	flag.Parse()

	if *outDir == "" {
		// Default to items package so I can debug this file. :D
		*outDir = "sim/core/items"
	}

	itemResponses := make([]WowheadItemResponse, len(ItemDeclarations))
	for idx, itemDeclaration := range ItemDeclarations {
		itemResponse := getWowheadItemResponse(itemDeclaration.ID)
		//fmt.Printf("\n\n%+v\n", itemResponse)
		itemResponses[idx] = itemResponse
	}

	writeItemFile(*outDir, ItemDeclarations, itemResponses)
}
