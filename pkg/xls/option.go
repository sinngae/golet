package xls

type optional struct {
	sheetName      string
	resetSheetName bool
}

type OptionFunc func(*optional)

func parseOptionFunc(list []OptionFunc) *optional {
	options := &optional{
		sheetName: DefaultSheetName,
	}
	for _, handle := range list {
		handle(options)
	}
	return options
}

func WithSheetName(name string) OptionFunc {
	return func(o *optional) {
		o.sheetName = name
		if name != DefaultSheetName {
			o.resetSheetName = true
		}
	}
}
