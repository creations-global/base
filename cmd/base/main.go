package main

import (
	// "fmt"
	"github.com/creations-global/base/internal/base"
)

var (
	buildTimestamp = ""
)

func main() {
	// fmt.Println("Hello, Base! at ", buildTimestamp)
	new(base.Base).Init(buildTimestamp).Execute()
}