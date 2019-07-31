package reader

import (
	inter "github.com/MovieStoreGuy/intstats/interface"
)

func readInterfaces(_ string) (map[string]*inter.Stat, error) {
	return nil, ErrUnsupportedPlatform
}

func readSockets(path string) (map[string]map[string]uint64, error) {
	return nil, ErrUnsupportedPlatform
}
