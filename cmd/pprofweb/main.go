package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func app() {
	for {
		for i := 0; i < 10000000; i++ {

		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go app()
	http.ListenAndServe("0.0.0.0:9091", nil)
}
