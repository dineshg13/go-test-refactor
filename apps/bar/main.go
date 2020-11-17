package main

import (
	"fmt"

	bootstrap "github.com/j-vizcaino/go-test-refactor"
)

const appName = "bar"

func main() {
	a := bootstrap.NewApp(appName)
	fmt.Println(a.String())
	a.Name = appName + "-updated"
	fmt.Println(a.String())
}
