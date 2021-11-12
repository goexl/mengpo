package god

import (
	`errors`
)

var errInvalidType = errors.New(`必须是一个结构体指针`)
