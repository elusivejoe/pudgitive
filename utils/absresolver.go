package utils

func ResolveAbsolute(pwd string, path *NormPath) *NormPath {
	if !path.IsAbs() && len(pwd) > 0 {
		return NewNormPath("/" + pwd + "/" + path.Path())
	}

	return path
}
