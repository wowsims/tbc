package main

import (
  //"bufio"
  "fmt"
  "os"
)

func writeItemFiles(outDir string, itemDeclarations []ItemDeclaration, itemResponses []WowheadItemResponse) {

  file, err := os.Create(outDir + "/test.go")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  file.WriteString("package items\n\n")
  file.WriteString("var items = []Item{\n")

  for idx, itemDeclaration := range itemDeclarations {
    itemResponse := itemResponses[idx]
    file.WriteString(fmt.Sprintf("\t%s,\n", itemToGoString(itemDeclaration, itemResponse)))
  }

  file.WriteString("}\n")

  file.Sync()
}

func itemToGoString(itemDeclaration ItemDeclaration, itemResponse WowheadItemResponse) string {
  itemStr := "{"

  itemStr += fmt.Sprintf("Name: \"%s\",", itemResponse.Name)
  itemStr += fmt.Sprintf("ID: %d,", itemDeclaration.ID)

  itemStr += fmt.Sprintf("Stats: %s,", statsToGoString(itemResponse.GetStats()))

  itemStr += "}"
  return itemStr
}

func statsToGoString(stats Stats) string {
  statsStr := "Stats{"

  for stat, value := range stats {
    statsStr += fmt.Sprintf("%s:%.0f,", StatName(stat), value)
  }

  statsStr += "}"
  return statsStr
}
