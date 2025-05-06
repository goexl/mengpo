package kernel

type Getter interface {
	Get(string) string
}
