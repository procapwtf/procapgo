# golang api wrapper for [procap solver](https://procap.wtf)

## installation

`go get github.com/procapwtf/procapgo`

## usage

```go
package main

import (
	"log"

	"github.com/procapwtf/procapgo"
)

func main() {
	bal, err := procapgo.GetBalance("ab5-xxxx....")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("balance: %f", bal)
	pass, key, err := procapgo.Solve(procapgo.Options{
		RawUrl:  "https://example.com/",
		Sitekey: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		Apikey:  "ab5-xxxx....",
	})
	if err != nil {
		log.Fatal("failed:", err)
	}
	log.Printf("pass: %s\nkey: %s\n", pass, key)
}
```

more captchas coming soon stay updated on our [discord server](https://discord.gg/procap)
