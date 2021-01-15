package meta

type Attributes struct {
	IsDir bool
}

func NewAttributes(isDir bool) Attributes {
	return Attributes{IsDir: isDir}
}
