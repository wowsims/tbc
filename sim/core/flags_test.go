package core

import (
	"testing"
)

func TestProcMasks(t *testing.T) {
	if !ProcMaskMeleeMHAuto.Matches(ProcMaskMeleeMHAuto | ProcMaskMeleeMHSpecial) {
		t.Fatalf("Expected mask match but was mismatch")
	}
	if !(ProcMaskMeleeMHAuto | ProcMaskMeleeMHSpecial).Matches(ProcMaskMeleeMHAuto) {
		t.Fatalf("Expected mask match but was mismatch")
	}
}
