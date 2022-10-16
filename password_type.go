package pm1

type PasswordType int

const (
	_              PasswordType = iota
	Random                      // like "swe2dp230-213m2343v##@*(*@*@(*(*@"
	EasyToRemember              // like "easy-to-remember"
)
