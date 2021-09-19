package main

const (
  EleSham = "elemental_shaman"
)

type ItemDeclaration struct {
  ID int
  Specs []string // Which specs use this item
}

// Keep these in alphabetical order by name (including name as a comment)
var ItemDeclarations = []ItemDeclaration{
  /** Cyclone Faceguard */ { ID: 29035, Specs: []string{EleSham}, },
  { ID: 29097, Specs: []string{EleSham}, },
  { ID: 29055, Specs: []string{EleSham}, },
  { ID: 29523, Specs: []string{EleSham}, },
}
