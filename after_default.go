package mengpo

import (
	`github.com/drone/envsubst`
)

func afterDefault(from string) (string, error) {
	return envsubst.EvalEnv(from)
}
