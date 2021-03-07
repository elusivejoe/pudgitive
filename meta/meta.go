package meta

type Meta struct {
	Name  string
	IsDir bool
}

func NewMeta(name string, isDir bool) Meta {
	return Meta{
		Name:  name,
		IsDir: isDir,
	}
}
