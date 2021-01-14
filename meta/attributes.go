package meta

type Attributes struct {
	IsDir bool
}

func (a *Attributes) IsDirectory() bool {
	return a.IsDir
}
