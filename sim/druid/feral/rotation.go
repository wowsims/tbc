package feral

 import (
 	"github.com/wowsims/tbc/sim/core"
 )

 func (cat *FeralDruid) OnGCDReady(sim *core.Simulation) {
 	cat.doRotation(sim)
 }

 func (cat *FeralDruid) doRotation(sim *core.Simulation) {
 }
