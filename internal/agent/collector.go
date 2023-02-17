package agent

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type metric interface {
	getURLUpdateParam() string
	toString() string
	sendReport(server string, timeout time.Duration) error
}

type gauge struct {
	name  string
	value float64
}

func (g gauge) toString() string {
	return fmt.Sprintf("%f", g.value)
}

func (g gauge) getURLUpdateParam() string {
	return fmt.Sprintf("gauge/%s/%f", g.name, g.value)
}

type counter struct {
	name  string
	value int64
}

func (c counter) toString() string {
	return fmt.Sprintf("%d", c.value)
}

func (c counter) getURLUpdateParam() string {
	return fmt.Sprintf("counter/%s/%d", c.name, c.value)
}

func (c counter) sendReport(server string, timeout time.Duration) error {
	url := "http://" + server + "/update/" + c.getURLUpdateParam()
	err := transmitPlainText(url, c.toString(), timeout)
	if err != nil {
		return err
	}
	return nil
}

func (g gauge) sendReport(server string, timeout time.Duration) error {
	url := "http://" + server + "/update/" + g.getURLUpdateParam()
	err := transmitPlainText(url, g.toString(), timeout)
	if err != nil {
		return err
	}
	return nil
}

func poolCounter(cnt int64) []counter {
	fmt.Println("poolCounter")
	return []counter{
		{"PollCount", cnt},
	}
}

func poolGauge() []gauge {
	fmt.Println("poolGauge")
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return []gauge{
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
