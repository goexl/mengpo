package mengpo

const (
	errorModPanic errorMode = iota
	errorModReturn
	errorModSilent
)

type errorMode uint8
