package utils

import "strings"

func TrimDbKey(root, pwd string, key string, isAbs bool) string {
	prefix := root

	if !isAbs && len(pwd) > 0 {
		prefix += "/" + pwd
	}

	key = strings.TrimPrefix(key, prefix)

	return key
}
