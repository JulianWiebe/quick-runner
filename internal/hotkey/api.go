package hotkey

type Combo struct {
	Ctrl, Alt, Shift, Super bool
	Key                     rune
}
