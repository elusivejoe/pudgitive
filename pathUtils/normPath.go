package pathUtils

import (
	"strings"
)

type NormPath struct {
	isAbs      bool
	normalized string
	parts      []string
}

func (p *NormPath) Parts() []string {
	return p.parts
}

func (p *NormPath) IsAbs() bool {
	return p.isAbs
}

func (p *NormPath) Path() string {
	return p.normalized
}

func NewNormPath(path string) *NormPath {
	normalized := normalize(path)

	return &NormPath{
		isAbs:      strings.HasPrefix(path, "/"),
		normalized: normalized,
		parts:      split(normalized),
	}
}

func split(path string) []string {
	if path == "/" {
		return nil
	}

	splits := strings.Split(path, "/")

	if isAbs := strings.HasPrefix(path, "/"); isAbs {
		return splits[1:]
	}

	return splits
}

func normalize(path string) string {
	if len(path) == 0 || path == "." {
		return ""
	}

	var normalized strings.Builder
	var unique rune = -1

	for _, current := range path {
		if current == '/' && current == unique {
			continue
		}

		normalized.WriteRune(current)
		unique = current
	}

	path = normalized.String()

	for strings.Contains(path, "/./") {
		path = strings.ReplaceAll(path, "/./", "/")
	}

	path = strings.TrimPrefix(path, "./")

	if len(path) > 1 {
		path = strings.TrimRight(path, "/")
	}

	if path == "/." {
		return "/"
	}

	return path
}
