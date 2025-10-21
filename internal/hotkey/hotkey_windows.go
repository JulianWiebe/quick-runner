//go:build windows

package hotkey

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var (
	user32               = windows.NewLazySystemDLL("user32.dll")
	procRegisterHotKey   = user32.NewProc("RegisterHotKey")
	procUnregisterHotKey = user32.NewProc("UnregisterHotKey")
)

const (
	mod_ALT      = 0x0001
	mod_CTRL     = 0x0002
	mod_SHIFT    = 0x0004
	mod_WIN      = 0x0008
	mod_NOREPEAT = 0x4000
)

func translateCombo(c Combo) uintptr {
	var m uintptr
	if c.Alt {
		m |= mod_ALT
	}
	if c.Ctrl {
		m |= mod_CTRL
	}
	if c.Shift {
		m |= mod_SHIFT
	}
	if c.Super {
		m |= mod_WIN
	}
	m |= mod_NOREPEAT
	return m
}

// Registers hotkey to show or hide the application
func Register(c Combo) error {
	combo := translateCombo(c)
	key := uintptr(c.Key)
	ok, _, err := procRegisterHotKey.Call(0, 1, combo, key)
	if ok == 0 {
		return fmt.Errorf("RegisterHotKey failed: %v", err)
	}

	return nil
}

func Unregister() {
	procUnregisterHotKey.Call(0, 1)
}
