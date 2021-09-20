package main

import (
  "github.com/wowsims/tbc/generate_items/api"
)

type ItemDeclaration struct {
  ID int
  Specs []api.Spec // Which specs use this item
}

// Keep these sorted by ID.
var ItemDeclarations = []ItemDeclaration{
  // Just some test items for now to make sure the generator is working
  { ID: 29035, Specs: []api.Spec{api.Spec_SpecElementalShaman}, },
  { ID: 29055, Specs: []api.Spec{api.Spec_SpecElementalShaman}, },
  { ID: 29097, Specs: []api.Spec{api.Spec_SpecElementalShaman}, },
  { ID: 29523, Specs: []api.Spec{api.Spec_SpecElementalShaman}, },
  { ID: 30153, Specs: []api.Spec{api.Spec_SpecElementalShaman}, },
}
