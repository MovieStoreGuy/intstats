package main

import (
	"flag"
	"fmt"

	reader2 "github.com/MovieStoreGuy/intstats/reader"
)

var (
	procPath string
)

func init() {
	flag.StringVar(&procPath, "proc-path", "/proc/net", "define the path where the /net directory exists")
}

func main() {
	flag.Parse()
	reader := reader2.NewReader(procPath)
	info, err := reader.GetInterfaceStatistics()
	if err != nil {
		panic(err)
	}
	for _, d := range info {
		fmt.Println(d)
	}
}
