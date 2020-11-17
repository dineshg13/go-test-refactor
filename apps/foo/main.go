package main

import (
	"fmt"

	"github.com/j-vizcaino/go-test-refactor/pkg/app"
)

const appName = "foo"

func main() {
	a := app.NewApp(appName)
	fmt.Println(a.String())
	a.Name = appName + "-updated"
	fmt.Println(a.String())
}

