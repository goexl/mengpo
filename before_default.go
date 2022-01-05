package mengpo

import (
	`strings`

	`github.com/drone/envsubst`
)

func beforeDefault(from string) (to string, err error) {
	hasEnvCount := 0
	to = from
	for {
		if to, err = envsubst.EvalEnv(to); nil != err {
			break
		}

		if hasEnvCount >= 2 {
			break
		}
		if strings.Contains(to, `$`) {
			hasEnvCount++
		}
	}

	return
}
