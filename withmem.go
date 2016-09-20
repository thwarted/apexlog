package log

import (
	"runtime"
)

type logMemStats struct {
	Alloc       uint64
	TotalAlloc  uint64
	HeapAlloc   uint64
	HeapObjects uint64
	StackInuse  uint64
	StackSys    uint64
}

func getMemStats() interface{} {
	var mem runtime.MemStats
	// runtime.Memstats is a pretty big structure, so just grab some key stuff
	runtime.ReadMemStats(&mem)
	return logMemStats{
		Alloc:       mem.Alloc,
		TotalAlloc:  mem.TotalAlloc,
		HeapAlloc:   mem.HeapAlloc,
		HeapObjects: mem.HeapObjects,
		StackInuse:  mem.StackInuse,
		StackSys:    mem.StackSys,
	}
}
