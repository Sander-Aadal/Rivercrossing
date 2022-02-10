// Package for testing events in the riverworld
// Possible events: put(item), getin(), getout(), crossriver(), takeout(item)
// start() or setup(), reset() osv.
package event

import (
	"testing"

	"github.com/Sander-Aadal/Rivercrossing/state"
)

func TestMoveItem(t *testing.T) {
	state.ViewState(0, 0, 0, 0)

	wanted := ""
	Put(wanted)

	bItem := state.GetBoatitem()

	if bItem != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", bItem, wanted)
	}
}

func TestCross(t *testing.T) {
	wanted := "GameOver"
	got := Crossriver()
	t.Log(got)

	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
