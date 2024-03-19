# golang api wrapper for [procap solver](https://procap.wtf)

## installation

`go get github.com/procapwtf/procapgo`

## usage

```go
package main

import (
	"log"
	"time"

	"github.com/procapwtf/procapgo"
)

func main() {
	user, err := procapgo.GetUser("ab5-xxxxxx...")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("daily limit: %d daily reset %s daily used %d daily remaining %d funds %f expire %s\n",
		user.DailyLimit,
		time.Unix(user.DailyReset, 0).Format("2006-01-02 15:04:05"),
		user.DailyUsed,
		user.DailyRemaining,
		user.Funds,
		time.Unix(user.PlanExpire, 0).Format("2006-01-02 15:04:05"),
	)
	pass, key, err := procapgo.Solve(procapgo.Options{
		RawUrl:  "https://example.com/", // full url in search bar
		Sitekey: "xxxxxx-xxxxx...",
		Apikey:  "ab5-xxxxxxx...",
	})
	if err != nil {
		log.Fatal("failed:", err)
	}
	log.Printf("pass: %s\nchallange key: %s\n", pass, key)
}

```

more captchas coming soon stay updated on our [discord server](https://discord.gg/procap)
