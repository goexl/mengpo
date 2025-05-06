package constant

const (
	ModPanic Mode = iota + 1
	ModReturn
	ModSilent
)

type Mode uint8
