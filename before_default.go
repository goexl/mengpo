package mengpo

import (
	`github.com/drone/envsubst`
)

func beforeDefault(from string) (string, error) {
	return envsubst.EvalEnv(from)
}
