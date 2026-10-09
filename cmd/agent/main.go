package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"net/http"
	"runtime"
	"time"
)

var gauges = map[string]float64{}

const pollInterval = 2
const reportInterval = 10

var pollCount = 0

func collectMetrics() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	gauges["Alloc"] = float64(stats.Alloc)
	gauges["BuckHashSys"] = float64(stats.BuckHashSys)
	gauges["Frees"] = float64(stats.Frees)
	gauges["GCCPUFraction"] = float64(stats.GCCPUFraction)
	gauges["GCSys"] = float64(stats.GCSys)
	gauges["HeapAlloc"] = float64(stats.HeapAlloc)
	gauges["HeapIdle"] = float64(stats.HeapIdle)
	gauges["HeapInuse"] = float64(stats.HeapInuse)
	gauges["HeapObjects"] = float64(stats.HeapObjects)
	gauges["HeapReleased"] = float64(stats.HeapReleased)
	gauges["HeapSys"] = float64(stats.HeapSys)
	gauges["LastGC"] = float64(stats.LastGC)
	gauges["Lookups"] = float64(stats.Lookups)
	gauges["MCacheInuse"] = float64(stats.MCacheInuse)
	gauges["MCacheSys"] = float64(stats.MCacheSys)
	gauges["MSpanInuse"] = float64(stats.MSpanInuse)
	gauges["MSpanSys"] = float64(stats.MSpanSys)
	gauges["Mallocs"] = float64(stats.Mallocs)
	gauges["NextGC"] = float64(stats.NextGC)
	gauges["NumForcedGC"] = float64(stats.NumForcedGC)
	gauges["NumGC"] = float64(stats.NumGC)
	gauges["OtherSys"] = float64(stats.OtherSys)
	gauges["PauseTotalNs"] = float64(stats.PauseTotalNs)
	gauges["StackInuse"] = float64(stats.StackInuse)
	gauges["StackSys"] = float64(stats.StackSys)
	gauges["Sys"] = float64(stats.Sys)
	gauges["TotalAlloc"] = float64(stats.TotalAlloc)
	gauges["RandomValue"] = rand.Float64()

	pollCount++
}

func reportMetrics(port int) {
	baseUrl := fmt.Sprintf("http://localhost:%d/update", port)
	for m, v := range gauges {
		_, err := http.Post(fmt.Sprintf("%s/gauge/%s/%f", baseUrl, m, v), "text/plain", nil)

		if err != nil {
			fmt.Println(err)
		}
	}

	_, err := http.Post(fmt.Sprintf("%s/counter/PollCount/%d", baseUrl, pollCount), "text/plain", nil)

	if err != nil {
		fmt.Println(err)
	} else {
		pollCount = 0
	}
}

func main() {
	portPtr := flag.Int("p", 8080, "target server port")
	flag.Parse()

	go func() {
		for {
			collectMetrics()

			time.Sleep(pollInterval * time.Second)
		}
	}()

	go func() {
		for {
			reportMetrics(*portPtr)

			time.Sleep(reportInterval * time.Second)
		}
	}()

	select {}
}
