package mengpo

var _ = New

type builder struct {
	options *options
}

func New() *builder {
	return &builder{
		options: newDefaultOptions(),
	}
}

func (b *builder) Tag(tag string) *builder {
	b.options.tag = tag

	return b
}

func (b *builder) Initialize(initialize bool) *builder {
	b.options.initialize = initialize

	return b
}

func (b *builder) PanicOnError() *builder {
	b.options.errorMod = errorModPanic

	return b
}

func (b *builder) ReturnOnError() *builder {
	b.options.errorMod = errorModReturn

	return b
}

func (b *builder) SilentOnError() *builder {
	b.options.errorMod = errorModSilent

	return b
}

func (b *builder) Getter(getter getter) *builder {
	b.options.getter = getter

	return b
}

func (b *builder) Processor(processor processor) *builder {
	b.options.processors = append(b.options.processors, processor)

	return b
}

func (b *builder) Build() *mengpo {
	return newMengpo(b.options)
}
