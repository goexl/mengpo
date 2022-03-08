package mengpo

var (
	_        = Tag
	_ option = (*optionTag)(nil)
)

type optionTag struct {
	tag string
}

// Tag 配置标签
func Tag(tag string) *optionTag {
	return &optionTag{
		tag: tag,
	}
}

func (t *optionTag) apply(options *options) {
	options.tag = t.tag
}
