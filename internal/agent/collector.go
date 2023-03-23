package agent

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func getMemStat() map[string]float64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return map[string]float64{
		"RandomValue":   rand.Float64(),
		"Alloc":         float64(memStats.Alloc),
		"BuckHashSys":   float64(memStats.BuckHashSys),
		"Frees":         float64(memStats.Frees),
		"GCCPUFraction": memStats.GCCPUFraction,
		"GCSys":         float64(memStats.GCSys),
		"HeapAlloc":     float64(memStats.HeapAlloc),
		"HeapIdle":      float64(memStats.HeapIdle),
		"HeapInuse":     float64(memStats.HeapInuse),
		"HeapObjects":   float64(memStats.HeapObjects),
		"HeapReleased":  float64(memStats.HeapReleased),
		"HeapSys":       float64(memStats.HeapSys),
		"LastGC":        float64(memStats.LastGC),
		"Lookups":       float64(memStats.Lookups),
		"MCacheInuse":   float64(memStats.MCacheInuse),
		"MCacheSys":     float64(memStats.MCacheSys),
		"Mallocs":       float64(memStats.Mallocs),
		"NextGC":        float64(memStats.NextGC),
		"NumForcedGC":   float64(memStats.NumForcedGC),
		"NumGC":         float64(memStats.NumGC),
		"OtherSys":      float64(memStats.OtherSys),
		"PauseTotalNs":  float64(memStats.PauseTotalNs),
		"StackInuse":    float64(memStats.StackInuse),
		"StackSys":      float64(memStats.StackSys),
		"Sys":           float64(memStats.Sys),
		"MSpanInuse":    float64(memStats.MSpanInuse),
		"MSpanSys":      float64(memStats.MSpanSys),
		"TotalAlloc":    float64(memStats.TotalAlloc),
	}
}

func getMemory() map[string]float64 {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return map[string]float64{}
	}

	return map[string]float64{
		"TotalMemory": float64(vm.Total),
		"FreeMemory":  float64(vm.Free),
	}
}

func getCPU() map[string]float64 {
	metrics := make(map[string]float64)
	cpuCount, err := cpu.Percent(0, true)
	if err != nil {
		return metrics
	}

	for i, c := range cpuCount {
		metrics[fmt.Sprintf("CPUutilization%d", i+1)] = c
	}
	return metrics
}
