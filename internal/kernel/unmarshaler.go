package kernel

type Unmarshaler interface {
	UnmarshalString(string) error
}
