package password

type Type int

const (
	_              Type = iota
	Random              // like "swe2dp230-213m2343v##@*(*@*@(*(*@"
	EasyToRemember      // like "easy-to-remember"
)
