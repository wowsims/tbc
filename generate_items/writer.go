package main

import (
  "github.com/wowsims/tbc/generate_items/api"
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
  file.WriteString("import (\n")
  file.WriteString("\t\"github.com/wowsims/tbc/generate_items/api\"\n")
  file.WriteString(")\n\n")
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

  itemStr += fmt.Sprintf("Phase: %d,", itemResponse.GetPhase())
  itemStr += fmt.Sprintf("Quality: api.ItemQuality_%s,", api.ItemQuality(itemResponse.Quality).String())

  itemStr += fmt.Sprintf("Stats: %s,", statsToGoString(itemResponse.GetStats()))

  gemSockets := itemResponse.GetGemSockets()
  if len(gemSockets) > 0 {
    itemStr += "GemSlots: []GemColor{"
    for _, gemColor := range gemSockets {
      itemStr += fmt.Sprintf("api.GemColor_%s,", gemColor.String())
    }
    itemStr += "},"
  }

  itemStr += fmt.Sprintf("SocketBonus: %s,", statsToGoString(itemResponse.GetSocketBonus()))

  itemStr += "}"
  return itemStr
}

func statsToGoString(stats Stats) string {
  statsStr := "Stats{"

  for stat, value := range stats {
    if value > 0 {
      statsStr += fmt.Sprintf("api.Stat_%s:%.0f,", api.Stat(stat).String(), value)
    }
  }

  statsStr += "}"
  return statsStr
}
