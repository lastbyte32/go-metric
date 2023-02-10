package agent

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
)

type gauge struct {
	name  string
	value float64
}

func (g gauge) getValue() []byte {

	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, g.value)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func (g gauge) getParam() string {
	return fmt.Sprintf("gauge/%s/%f", g.name, g.value)
}

type counter struct {
	name  string
	value int64
}

func (c counter) getParam() string {
	return fmt.Sprintf("counter/%s/%d", c.name, c.value)
}

func (c counter) getValue() []byte {
	return []byte(strconv.FormatInt(c.value, 10))
}

func poolCounter(counters *[]counter, cnt int64) {
	fmt.Println("poolCounter")

	*counters = []counter{
		{"PollCount", cnt},
	}
}

func poolGauge(g *[]gauge) {
	fmt.Println("poolGauge")
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	*g = []gauge{
		{"RandomValue", rand.Float64()},
		{"Alloc", float64(memStats.Alloc)},
		{"BuckHashSys", float64(memStats.BuckHashSys)},
		{"Frees", float64(memStats.Frees)},
		{"GCCPUFraction", memStats.GCCPUFraction},
		{"GCSys", float64(memStats.GCSys)},
		{"HeapAlloc", float64(memStats.HeapAlloc)},
		{"HeapIdle", float64(memStats.HeapIdle)},
		{"HeapInuse", float64(memStats.HeapInuse)},
		{"HeapObjects", float64(memStats.HeapObjects)},
		{"HeapReleased", float64(memStats.HeapReleased)},
		{"HeapSys", float64(memStats.HeapSys)},
		{"LastGC", float64(memStats.LastGC)},
		{"Lookups", float64(memStats.Lookups)},
		{"MCacheInuse", float64(memStats.MCacheInuse)},
		{"MCacheSys", float64(memStats.MCacheSys)},
		{"Mallocs", float64(memStats.Mallocs)},
		{"NextGC", float64(memStats.NextGC)},
		{"NumForcedGC", float64(memStats.NumForcedGC)},
		{"NumGC", float64(memStats.NumGC)},
		{"OtherSys", float64(memStats.OtherSys)},
		{"PauseTotalNs", float64(memStats.PauseTotalNs)},
		{"StackInuse", float64(memStats.StackInuse)},
		{"StackSys", float64(memStats.StackSys)},
		{"Sys", float64(memStats.Sys)},
		{"TotalAlloc", float64(memStats.TotalAlloc)},
	}
}
