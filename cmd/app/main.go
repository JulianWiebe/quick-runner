package main

import (
	"os"

	"github.com/JulianWiebe/quick-runner/internal/hotkey"
)

func main() {
	/*
		This is the main entry for the runner application.

		It needs to do a few things in order to function correctly:
		1) Check which OS it is running on
		2) Setup the main keyboard hook/hotkey
		3) If present, load configuration files
		4) If not turned off, show that the application is now running
	*/
	println("Hello World")
	combo := hotkey.Combo{Ctrl: true, Alt: true, Key: 'K'}
	if err := hotkey.Register(combo); err != nil {
		os.Exit(-1) // TODO: Implement showing a window with error text and identification code.
	}

	defer hotkey.Unregister()
}
