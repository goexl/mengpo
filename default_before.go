package mengpo

import (
	`github.com/drone/envsubst`
)

func env(original string) (to string) {
	if eval, err := envsubst.EvalEnv(original); nil != err {
		to = original
	} else {
		to = eval
	}

	return
}
