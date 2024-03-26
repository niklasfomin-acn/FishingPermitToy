package main

import (
	presentation "client/ui"
	"time"
)

func main() {
	// Eingabemaske
	presentation.ShowWelcomeMessage()
	time.Sleep(3 * time.Second)
	presentation.ShowOptions()
}
