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
