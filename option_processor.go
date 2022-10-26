package mengpo

var (
	_        = Processor
	_ option = (*optionProcessor)(nil)
)

type optionProcessor struct {
	processor processor
}

// Processor 配置生命周期前的操作
func Processor(processor processor) *optionProcessor {
	return &optionProcessor{
		processor: processor,
	}
}

func (p *optionProcessor) apply(options *options) {
	options.processors = append(options.processors, p.processor)
}
