package meta

type Meta struct {
	Name  string
	Attrs Attributes
}

func NewMeta(name string, isDir bool) Meta {
	return Meta{
		Name:  name,
		Attrs: Attributes{IsDir: isDir},
	}
}

func (e *Meta) EntityName() string {
	return e.Name
}

func (e *Meta) Attributes() *Attributes {
	return &e.Attrs
}
