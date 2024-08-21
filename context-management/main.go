package main

import (
	"context-management/contexts"
)

func main() {
	//contexts.WithCancel()
	//contexts.WithTimeoutContext()
	//contexts.WithDeadlineContext()
	contexts.WithValueContext()
}
