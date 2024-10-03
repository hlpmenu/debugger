package debugger

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuProfileFile *os.File
var memProfileFile *os.File

// Start starts the CPU profiler.
func Start(cpuProfile string) {
	cpuProfileFile, err := os.Create(cpuProfile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(cpuProfileFile); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	// Ensure that profile data is saved even if a panic occurs
	defer Stop()
}

// Stop stops the CPU profiler and writes the profile to disk.
func Stop() {
	pprof.StopCPUProfile()
	cpuProfileFile.Close()
}

// StartMemProfile starts the memory profiler.
func StartMemProfile(memProfile string) {
	var err error
	memProfileFile, err = os.Create(memProfile)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
}

// StopMemProfile stops the memory profiler and writes the profile to disk.
func StopMemProfile() {
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(memProfileFile); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	memProfileFile.Close()
}

// Timer is a type for timing operations.
type Timer struct {
	StartTime time.Time
	Operation string
}

// Start starts the timer.
func (t *Timer) Start() {

	t.StartTime = time.Now()
}

// Stop stops the timer and logs the elapsed time.
func (t *Timer) Stop() {
	elapsed := time.Since(t.StartTime)
	log.Printf("The operation for %s took %s", t.Operation, elapsed)

}

// Start starts the timer.
func (t *Timer) StartReq(r *http.Request) {
	go func() {
		t.Operation = r.URL.Path
	}()
	t.StartTime = time.Now()
}

// Stop stops the timer and logs the elapsed time.
func (t *Timer) StopReq() {
	elapsed := time.Since(t.StartTime)
	log.Printf("The operation for %s took %s", t.Operation, elapsed)

}

func NewTimer(s string) *Timer {
	t := &Timer{
		Operation: s,
	}
	t.Start()
	return t
}
