package main

import (
	"log"
	"runtime"
	"time"
)

//Monitor objects hold all the performance stats we intend to monitor.
type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64

	NumGC        uint32
	NumGoroutine int
}

/*NewMonitor creates a new monitor object and starts monitoring.
Should be called with the go- prefix, otherwise it will halt flow.*/
func NewMonitor(duration int) {
	var m Monitor
	var rtm runtime.MemStats
	var interval = time.Duration(duration) * time.Second
	for {
		<-time.After(interval)

		// Read full mem stats
		runtime.ReadMemStats(&rtm)

		// Number of goroutines
		m.NumGoroutine = runtime.NumGoroutine()

		// Misc memory stats
		m.Alloc = rtm.Alloc
		m.TotalAlloc = rtm.TotalAlloc
		m.Sys = rtm.Sys
		m.Mallocs = rtm.Mallocs
		m.Frees = rtm.Frees

		// Live objects = Mallocs - Frees
		m.LiveObjects = m.Mallocs - m.Frees

		// GC Stats
		m.PauseTotalNs = rtm.PauseTotalNs
		m.NumGC = rtm.NumGC

		// Just encode to json and print
		log.Printf(`
		%d - number of goroutines
		%f - MB of allocated heap objects
		%f - cumulative MB allocated for heap objects (does not decrease when objects are freed)
		%f - MB memory obtained from the OS reserved by the Go runtime
		%d - cumulative count of heap objects allocated
		%d - cumulative count of heap objects freed
		%d - count of live heap objects 
		%f - cumulative seconds in stop-the-world pauses
		%d - number of completed GC cycles
		`,
			m.NumGoroutine,
			float64(m.Alloc)/1000000,
			float64(m.TotalAlloc)/1000000,
			float64(m.Sys)/1000000,
			m.Mallocs, m.Frees, m.LiveObjects,
			float64(m.PauseTotalNs)/1000000000,
			m.NumGC)
	}
}
