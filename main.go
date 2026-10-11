package main

import (
	_ "time/tzdata"

	"github.com/vucongthanh92/courier-user-svc/startup"
)

func main() {
	startup.Execute()
}
