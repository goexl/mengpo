package mengpo

type getter interface {
	// Get 获取值
	Get(key string) (value string)
}
