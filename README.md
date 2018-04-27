# Golang bindings for the Staytus API

[![GoDoc](https://godoc.org/github.com/aprosvetova/go-staytus?status.svg)](https://godoc.org/github.com/aprosvetova/go-staytus)

I'm very new to Go, so I'll be happy if you make some Pull Requests and help me with tests.

Read the Staytus [description](https://github.com/adamcooke/staytus) and [API docs](https://github.com/adamcooke/staytus/tree/master/doc/api) to understand how it works.

## Example

```go
package main

import (
	"fmt"
	"github.com/aprosvetova/go-staytus"
	"os"
)

func main() {
	api, err := staytus.New("https://status.yourcompany.com", "YOUR_TOKEN", "YOUR_SECRET")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	issues, err := api.GetIssues()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(issues)
}
```

## Contact me
If you have any questions about my shitty code, feel free to contact me by Telegram ([@koteeq](https://t.me/koteeq)). I speak English and Russian.