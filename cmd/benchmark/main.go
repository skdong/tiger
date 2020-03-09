package main

import (
	"flag"
	"log"
	"os"

	"runtime/pprof"

	"github.com/skdong/tiger/pkg/bytedance"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	bytedance.ReverseWords("the sky is blue")
}
