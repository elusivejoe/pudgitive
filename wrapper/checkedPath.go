package wrapper

import (
	"errors"
	"strings"
)

type checkedPath struct {
	isAbs      bool
	normalized string
	parts      []string
}

func (p *checkedPath) Parts() []string {
	return p.parts
}

func (p *checkedPath) IsAbs() bool {
	return p.isAbs
}

func (p *checkedPath) Path() string {
	return p.normalized
}

func NewCheckedPath(path string) (*checkedPath, error) {
	normalized := normalize(path)

	if len(normalized) == 0 {
		return nil, errors.New("empty path")
	}

	return &checkedPath{
		isAbs:      strings.HasPrefix(path, "/"),
		normalized: normalized,
		parts:      split(normalized),
	}, nil
}

func split(path string) []string {
	splits := strings.Split(path, "/")

	if isAbs := strings.HasPrefix(path, "/"); isAbs {
		return splits[1:]
	}

	return splits
}

func normalize(path string) string {
	path = strings.TrimRight(path, "/")

	var normalized strings.Builder
	var unique rune = -1

	for _, current := range path {
		if current == '/' && current == unique {
			continue
		}

		normalized.WriteRune(current)
		unique = current
	}

	return normalized.String()
}
