package main

import (
  "flag"
  //"fmt"
)

func main() {
  //clientId := flag.String("clientId", "", "Client ID provided by Blizzard API.")
  //clientSecret := flag.String("clientSecret", "", "Client secret provided by Blizzard API.")
  outDir := flag.String("outDir", "", "Path to output directory for writing generated .go files.")
  flag.Parse()

  if /***clientId == "" || *clientSecret == "" ||*/ *outDir == "" {
    panic("outDir flag is required!")
  }

  //accessToken := getAccessToken(*clientId, *clientSecret)
  //fmt.Println("AccessToken: ", accessToken)

  itemResponses := make([]WowheadItemResponse, len(ItemDeclarations))
  for idx, itemDeclaration := range ItemDeclarations {
    itemResponse := getWowheadItemResponse(itemDeclaration.ID)
    //fmt.Printf("\n\n%+v\n", itemResponse)
    itemResponses[idx] = itemResponse
  }

  writeItemFiles(*outDir, ItemDeclarations, itemResponses)
}
