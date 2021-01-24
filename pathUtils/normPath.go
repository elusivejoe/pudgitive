package pathUtils

import (
	"errors"
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

func NewNormPath(path string) (*NormPath, error) {
	normalized, err := normalize(path)

	if err != nil {
		return nil, err
	}

	return &NormPath{
		isAbs:      strings.HasPrefix(path, "/"),
		normalized: normalized,
		parts:      split(normalized),
	}, nil
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

func normalize(path string) (string, error) {
	if len(path) == 0 || path == "." {
		return "", nil
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
		return "/", nil
	}

	if path == "/.." {
		return "", errors.New("cannot go higher than root")
	}

	return path, nil
}
