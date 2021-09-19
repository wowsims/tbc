package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "regexp"
  "strconv"
  "time"
)

const (
  StatInt int = iota
  StatStm
  StatSpellCrit
  StatSpellHit
  StatSpellDmg
  StatHaste
  StatMP5
  StatMana
  StatSpellPen
  StatSpirit

  StatLen
)
type Stats [StatLen]float64
func StatName(s int) string {
  switch s {
  case StatInt:
    return "StatInt"
  case StatStm:
    return "StatStm"
  case StatSpellCrit:
    return "StatSpellCrit"
  case StatSpellHit:
    return "StatSpellHit"
  case StatSpellDmg:
    return "StatSpellDmg"
  case StatHaste:
    return "StatHaste"
  case StatMP5:
    return "StatMP5"
  case StatMana:
    return "StatMana"
  case StatSpellPen:
    return "StatSpellPen"
  case StatSpirit:
    return "StatSpirit"
  }

  return "none"
}

type WowheadItemResponse struct {
  Name string `json:"name"`
  Quality int `json:"quality"`
  Icon string `json:"icon"`
  Tooltip string `json:"tooltip"`
}

func (item WowheadItemResponse) GetTooltipRegexValue(pattern *regexp.Regexp, matchIdx int) int {
  match := pattern.FindStringSubmatch(item.Tooltip)
  if match == nil {
    return 0
  }

  val, err := strconv.Atoi(match[matchIdx])
  if err != nil {
    return 0
  }

  return val
}

func (item WowheadItemResponse) GetStatValue(pattern *regexp.Regexp) int {
  return item.GetTooltipRegexValue(pattern, 1)
}


var armorRegex, _ = regexp.Compile("<!--amr-->([0-9]+) Armor")
var agilityRegex, _ = regexp.Compile("<!--stat3-->\\+([0-9]+) Agility")
var strengthRegex, _ = regexp.Compile("<!--stat4-->\\+([0-9]+) Strength")
var intellectRegex, _ = regexp.Compile("<!--stat5-->\\+([0-9]+) Intellect")
var spiritRegex, _ = regexp.Compile("<!--stat6-->\\+([0-9]+) Spirit")
var staminaRegex, _ = regexp.Compile("<!--stat7-->\\+([0-9]+) Stamina")
var spellPowerRegex, _ = regexp.Compile("Increases damage and healing done by magical spells and effects by up to ([0-9]+).")
var spellCritRegex, _ = regexp.Compile("Improves spell critical strike rating by <!--rtg21-->([0-9]+).")
var healingPowerRegex, _ = regexp.Compile("Increases healing done by up to ([0-9]+) and damage done by up to ([0-9]+) for all magical spells and effects.")
var mp5Regex, _ = regexp.Compile("Restores ([0-9]+) mana per 5 sec.")

func (item WowheadItemResponse) GetStats() Stats {
  spellPower := item.GetStatValue(spellPowerRegex)
  //healingPowerFromHealing := item.GetTooltipRegexValue(healingPowerRegex, 1)
  spellPowerFromHealing := item.GetTooltipRegexValue(healingPowerRegex, 2)

  // Some items have both (e.g. Windhawk Bracers)
  spellPower = spellPower + spellPowerFromHealing
  //healingPower := spellPower + healingPowerFromHealing

  return Stats{
    StatInt: float64(item.GetStatValue(intellectRegex)),
    StatStm: float64(item.GetStatValue(staminaRegex)),
    StatSpellDmg: float64(spellPower),
    StatSpellCrit: float64(item.GetStatValue(spellCritRegex)),
    StatMP5: float64(item.GetStatValue(mp5)),
  }
}

func getWowheadItemResponse(itemId int) WowheadItemResponse {
  url := fmt.Sprintf("https://tbc.wowhead.com/tooltip/item/%d", itemId)

  httpClient := http.Client{
    Timeout: 5 * time.Second,
  }

  request, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.Fatal(err)
  }

  result, err := httpClient.Do(request)
  if err != nil {
    log.Fatal(err)
  }

  defer result.Body.Close()

  resultBody, err := ioutil.ReadAll(result.Body)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf(string(resultBody))
  itemResponse := WowheadItemResponse{}
  err = json.Unmarshal(resultBody, &itemResponse)
  if err != nil {
    log.Fatal(err)
  }

  return itemResponse
}
