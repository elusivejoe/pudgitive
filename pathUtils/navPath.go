package pathUtils

import (
	"errors"
	"strings"
)

type NavPath struct {
	resolvedPaths []*NormPath
}

func NewNavPath(path *NormPath) (*NavPath, error) {
	if path == nil {
		return nil, errors.New("nil path provided")
	}

	resolvedPaths, err := resolvePaths(path)

	if err != nil {
		return nil, err
	}

	return &NavPath{resolvedPaths: resolvedPaths}, nil
}

func (p *NavPath) DestList() []*NormPath {
	return p.resolvedPaths
}

func (p *NavPath) FinalDest() *NormPath {
	return p.resolvedPaths[len(p.resolvedPaths)-1]
}

func resolvePaths(path *NormPath) ([]*NormPath, error) {
	subPaths, err := collectSubPaths(path.Parts())

	if err != nil {
		return nil, err
	}

	subPaths = filterSubPaths(subPaths)

	result, err := normalizeSubPaths(subPaths, path.IsAbs())

	if err != nil {
		return nil, err
	}

	return result, nil
}

func collectSubPaths(parts []string) ([]string, error) {
	var paths []string
	var accumulator []string

	for idx := 0; idx < len(parts); idx++ {
		if parts[idx] == ".." {
			paths = append(paths, strings.Join(accumulator, "/"))

			for parts[idx] == ".." {
				if len(accumulator) == 0 {
					return nil, errors.New("cannot go any higher")
				}
				idx += 1
				accumulator = accumulator[:len(accumulator)-1]
			}
			idx -= 1
		} else {
			accumulator = append(accumulator, parts[idx])
		}
	}

	if len(accumulator) > 0 {
		paths = append(paths, strings.Join(accumulator, "/"))
	}

	return paths, nil
}

func filterSubPaths(subPaths []string) []string {
	duplicated := make(map[int]bool)

	for i := 0; i < len(subPaths); i++ {
		for j := 1; j < len(subPaths); j++ {
			if i != j && strings.HasPrefix(subPaths[j], subPaths[i]) {
				duplicated[i] = true
				break
			}
		}
	}

	var filtered []string

	for i := 0; i < len(subPaths); i++ {
		if duplicated[i] {
			continue
		}

		filtered = append(filtered, subPaths[i])
	}

	return filtered
}

func normalizeSubPaths(subPaths []string, isAbs bool) ([]*NormPath, error) {
	if len(subPaths) == 0 {
		subPaths = []string{""}
	}

	var result []*NormPath

	for _, subPath := range subPaths {
		if isAbs {
			subPath = "/" + subPath
		}
		result = append(result, NewNormPath(subPath))
	}

	return result, nil
}
