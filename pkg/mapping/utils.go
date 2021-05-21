package gomapping

type TagParser struct {
	tag string
	attr TagAttr
}

type TagAttr struct {
	required bool
}

func NewTagParser(tag string) *TagParser {
	return &TagParser{
		tag: tag,
		attr: TagAttr{
			required: true,
		},
	}
}

func (tg *TagParser) Parse() *TagAttr {
	//TODO
	return nil
}
