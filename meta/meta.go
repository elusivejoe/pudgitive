package meta

type Meta struct {
	name  string
	attrs Attributes
}

func NewMeta(name string, isDir bool) Meta {
	return Meta{
		name:  name,
		attrs: Attributes{isDir: isDir},
	}
}

func (e *Meta) Name() string {
	return e.name
}

func (e *Meta) Attributes() *Attributes {
	return &e.attrs
}
