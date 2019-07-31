package inter

import (
	"fmt"

	"github.com/MovieStoreGuy/intstats/utils"
)

type Stat struct {
	Name     string
	Received map[string]uint64
	Transmit map[string]uint64
}

func NewInterface(name string, val []string) (*Stat, error) {
	if len(val) < 16 {
		return nil, fmt.Errorf("Not enough values for %s interface", name)
	}
	return &Stat{
		Name: name,
		Received: map[string]uint64{
			"bytes":      parseInt(val[0]),
			"packets":    parseInt(val[1]),
			"errs":       parseInt(val[2]),
			"drop":       parseInt(val[3]),
			"fifo":       parseInt(val[4]),
			"frame":      parseInt(val[5]),
			"compressed": parseInt(val[6]),
			"multicast":  parseInt(val[7]),
		},
		Transmit: map[string]uint64{
			"bytes":      parseInt(val[8]),
			"packets":    parseInt(val[9]),
			"errs":       parseInt(val[10]),
			"drop":       parseInt(val[11]),
			"fifo":       parseInt(val[12]),
			"colls":      parseInt(val[13]),
			"carrier":    parseInt(val[14]),
			"compressed": parseInt(val[15]),
		},
	}, nil
}

func parseInt(value string) uint64 {
	return utils.ParseInt(value)
}
