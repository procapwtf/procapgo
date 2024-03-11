package procapgo

import (
	"testing"
)

func TestSolve(t *testing.T) {
	user, err := GetUser("apikey...")
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	t.Log(user.DailyLimit, user.DailyRemaining, user.DailyReset, user.DailyUsed, user.Funds, user.PlanExpire)
	pass, key, err := Solve(Options{
		RawUrl:  "https://accounts.hcaptcha.com/demo",
		Sitekey: "a5f74b19-9e45-40e0-b45d-47ff91b7a6c2",
		Apikey:  "apikey...",
	})
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	t.Logf("pass: %s\nkey: %s", pass, key)
}
