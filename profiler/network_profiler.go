package debugger

import (
	"log"
	"net/http"
	_ "net/http/pprof" // This is an anonymous import
)

type NetworkProfiler struct {
	Addr string
	Port string
}

func NewNetworkProfiler(addr string, port string) *NetworkProfiler {
	return &NetworkProfiler{
		Addr: addr,
		Port: port,
	}
}

func (np *NetworkProfiler) Start() {
	go func() {
		log.Println(http.ListenAndServe(np.Addr+":"+np.Port, nil))
	}()
}

func ProfileNetwork() {

	// Start a new goroutine for the profiler server
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

}
