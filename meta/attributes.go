package meta

type Attributes struct {
	isDir bool
}

func (a *Attributes) IsDir() bool {
	return a.isDir
}
