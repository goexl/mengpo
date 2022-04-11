package mengpo

const (
	// ErrorModPanic 抛出错误
	ErrorModPanic errorMod = 1
	// ErrorModReturn 返回错误
	ErrorModReturn errorMod = 2
	// ErrorModSilent 不处理，继续执行
	ErrorModSilent errorMod = 3
)

type errorMod uint8
