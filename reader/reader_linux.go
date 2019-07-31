package reader

import (
	"io/ioutil"
	"strings"

	"github.com/MovieStoreGuy/intstats/utils"

	inter "github.com/MovieStoreGuy/intstats/interface"
)

func readInterfaces(path string) (map[string]*inter.Stat, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(buff), "\n")
	interfaces := make(map[string]*inter.Stat)

	for line := 2; line < len(lines); line++ {
		fields := strings.Fields(lines[line])
		// removing the colon from the interface name
		name := fields[0][:len(fields[0])-1]
		iface, err := inter.NewInterface(name, fields[1:])
		if err != nil {
			// TODO(Sean Marciniak): Write some neat way to log this as an info level
			//                       but allow for different types of loggers
			continue
		}
		interfaces[name] = iface
	}

	return interfaces, nil
}

func readSockets(path string) (map[string]map[string]uint64, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(buff), "\n")
	protocols := make(map[string]map[string]uint64, 0)

	for line := 0; line < len(lines); line++ {
		fields := strings.Fields(lines[line])
		name := fields[0][:len(fields[0])-1]
		protocols[name] = make(map[string]uint64, 0)
		for defintions := 1; defintions < len(fields); defintions += 2 {
			protocols[name][fields[defintions]] = utils.ParseInt(fields[defintions+1])
		}
	}
	return protocols, nil
}
