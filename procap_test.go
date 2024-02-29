package procapgo

import (
	"testing"
)

func TestSolve(t *testing.T) {
	bal, err := GetBalance("ab5-xxxx....")
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	t.Logf("balance: %f", bal)
	pass, key, err := Solve(Options{
		RawUrl:  "https://example.com/",
		Sitekey: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		Apikey:  "ab5-xxxx....",
	})
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	t.Logf("pass: %s\nkey: %s", pass, key)
}
