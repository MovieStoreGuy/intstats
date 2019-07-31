package reader

import (
	"errors"
	"path"

	inter "github.com/MovieStoreGuy/intstats/interface"
)

var (
	// ErrUnsupportedPlatform is used when the given platform can not support the given statistics
	ErrUnsupportedPlatform = errors.New("not implemented for the given platform")
)

// Reader abstracts the required code in order to obtain interface statistics from a device
type Reader interface {

	// GetInterfaceStatistics returns a map of all interfaces and their current statistics
	GetInterfaceStatistics() (map[string]*inter.Stat, error)

	// GetSocketStatistics returns all the socket statistics
	GetSocketStatistics() (map[string]map[string]uint64, error)
}

type impl struct {
	DevPath    string
	SocketPath string
}

// NewReader returns a reader configured to use the given path to a /proc/net
func NewReader(netpath string) Reader {
	return &impl{
		DevPath:    path.Join(netpath, "dev"),
		SocketPath: path.Join(netpath, "sockstat"),
	}
}

func (i *impl) GetInterfaceStatistics() (map[string]*inter.Stat, error) {
	return readInterfaces(i.DevPath)
}

func (i *impl) GetSocketStatistics() (map[string]map[string]uint64, error) {
	return readSockets(i.SocketPath)
}
