package core

import "testing"

func TestColorIntersection(t *testing.T) {
	chryo := GemsByName["Rune Covered Chrysoprase"]

	if !chryo.Color.Intersects(GemColorBlue) {
		t.Fatalf("Chryo intersects blue...")
	}
	if !chryo.Color.Intersects(GemColorGreen) {
		t.Fatalf("Chryo intersects blue...")
	}
}

func TestAuraNames(t *testing.T) {
	for i := int32(0); i < MagicIDLen; i++ {
		if AuraName(i) == "<<Add Aura name to switch!!>>" {
			t.Logf("Missing Name for : %#v", i)
			t.Fail()
		}
	}
}
