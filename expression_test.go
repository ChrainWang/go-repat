package repat_test

import (
	"math/rand"
	"testing"

	"github.com/chrainwang/go-repat"
)

func TestMakeChar(t *testing.T) {
	var r byte
	for r = byte(33); r != byte(127); r++ {
		c := repat.MakeChar(rune(r))
		t.Log(c.ToString())
	}
}

func TestMakeRange(t *testing.T) {
	from := rand.Int31n(128 - 33)
	from = from + 33
	to := rand.Int31n(128 - from)
	to = from + to
	r := repat.MakeCharRange(from, to)
	t.Logf("pattern: %q", r.ToString())
	re := repat.MustCompile(r, repat.MatchEntier())
	t.Logf("pattern in regexp: %q", re.String())
	for i := from; i <= to; i++ {
		s := string(i)
		if !re.MatchString(s) {
			t.Errorf("character unmatched %c", i)
		}
	}
}

func TestMakeCC(t *testing.T) {
	cc := repat.MakeCharacterCollection(`!@#$%^-&`, repat.MakeCharRange('a', 'f'))
	t.Log(cc.ToString())
}
