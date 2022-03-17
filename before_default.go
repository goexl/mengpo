package mengpo

import (
	`reflect`
	`strings`

	`github.com/drone/envsubst`
	`github.com/goexl/env`
)

func beforeDefault(from string, _ reflect.StructField) (to string, err error) {
	hasEnvCount := 0
	to = from
	envPrefix := `$`
	for {
		if to, err = envsubst.Eval(to, env.Get); nil != err {
			break
		}

		if hasEnvCount >= 2 || !strings.Contains(to, envPrefix) {
			break
		}
		if strings.Contains(to, envPrefix) {
			hasEnvCount++
		}
	}

	return
}
