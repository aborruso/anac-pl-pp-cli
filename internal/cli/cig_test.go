package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestWarnCIGNonValido(t *testing.T) {
	casi := []struct {
		query  string
		avvisa bool
	}{
		{"B7E26B1DC8", true},
		{"B7E26B1DC8,", true},
		{"(B7E26B1DC8)", true},
		{"CIG:B7E26B1DC8", true},
		{"lavori cig=B7E26B1DC8 scuola", true},
		{"B7E26B1DC7", false},
		{"CIG:B7E26B1DC7", false},
		{"microsoft office", false},
		{"MANUTENZIO", false},
	}
	for _, c := range casi {
		var buf bytes.Buffer
		warnCIGNonValido(&buf, c.query)
		if got := strings.Contains(buf.String(), "non è valido"); got != c.avvisa {
			t.Errorf("warnCIGNonValido(%q): avviso=%v, atteso %v (%q)", c.query, got, c.avvisa, buf.String())
		}
	}
}
