package kernel

type Unmarshaler interface {
	Unmarshal([]byte) error
}
