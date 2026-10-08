package kernel

type Defaulter interface {
	Default() error
}
