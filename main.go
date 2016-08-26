package main

import (
	"github.com/himu62/echo-testing/route"
)

func main() {
	e := route.Init()
	route.Run(e)
}
