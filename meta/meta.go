package meta

type Meta struct {
	Name  string
	Attrs Attributes
}

func NewMeta(name string, isDir bool) Meta {
	return Meta{
		Name:  name,
		Attrs: NewAttributes(isDir),
	}
}
