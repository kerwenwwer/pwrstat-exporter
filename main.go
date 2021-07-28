package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kerwenwwer/gopwrstat"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	LoadDesc = prometheus.NewDesc(
		"ups_load",
		"UPS power load (Watt)",
		[]string{"device"},
		nil)

	StateDesc = prometheus.NewDesc(
		"ups_state",
		"UPS status (1 -> Normal, 0 -> Not)",
		[]string{"device"},
		nil)

	BatteryDesc = prometheus.NewDesc(
		"ups_battery_capacity",
		"UPS battery capacity(%)",
		[]string{"device"},
		nil)

	RtimeDesc = prometheus.NewDesc(
		"ups_remaining_runtime",
		"UPS Remaining Runtime(min)",
		[]string{"device"},
		nil)

	InVoltageDesc = prometheus.NewDesc(
		"ups_in_voltage",
		"UPS Input Voltage(V): pass-> 1 non_pass -> 0",
		[]string{"device"},
		nil)

	OutVoltageDesc = prometheus.NewDesc(
		"ups_out_voltage",
		"UPS Output Voltage(V): pass-> 1 non_pass -> 0",
		[]string{"device"},
		nil)

	TestDesc = prometheus.NewDesc(
		"ups_test_result",
		"UPS Test Result",
		[]string{"device"},
		nil)
)

type (
	PwrstatCollector struct{}
)

func NewPwrstatCollector() *PwrstatCollector {
	return &PwrstatCollector{}
}

func (l *PwrstatCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- LoadDesc
	ch <- StateDesc
	ch <- BatteryDesc
	ch <- RtimeDesc
	ch <- OutVoltageDesc
	ch <- TestDesc
}

func (l *PwrstatCollector) Collect(ch chan<- prometheus.Metric) {
	status, err := gopwrstat.NewFromSystem()
	if err != nil {
		panic(err)
	}
	for k, v := range status.Status {

		if k == "Load" {
			value_arr := strings.Fields(v)
			if value, err := strconv.ParseFloat(value_arr[0], 64); err == nil {
				ch <- prometheus.MustNewConstMetric(LoadDesc,
					prometheus.GaugeValue, value, status.Status["Model Name"])
			}
		} else if k == "State" {
			var state = 0
			if v == "Normal" {
				state = 1
			}
			ch <- prometheus.MustNewConstMetric(StateDesc,
				prometheus.GaugeValue, float64(state), status.Status["Model Name"])
		} else if k == "Battery Capacity" {
			value_arr := strings.Fields(v)
			if value, err := strconv.ParseFloat(value_arr[0], 64); err == nil {
				ch <- prometheus.MustNewConstMetric(BatteryDesc,
					prometheus.GaugeValue, value, status.Status["Model Name"])
			}
		} else if k == "Remaining Runtime" {
			value_arr := strings.Fields(v)
			if value, err := strconv.ParseFloat(value_arr[0], 64); err == nil {
				ch <- prometheus.MustNewConstMetric(RtimeDesc,
					prometheus.GaugeValue, value, status.Status["Model Name"])
			}
		} else if k == "Utility Voltage" {
			value_arr := strings.Fields(v)
			if value, err := strconv.ParseFloat(value_arr[0], 64); err == nil {
				ch <- prometheus.MustNewConstMetric(InVoltageDesc,
					prometheus.GaugeValue, value, status.Status["Model Name"])
			}
		} else if k == "Output Voltage" {
			value_arr := strings.Fields(v)
			if value, err := strconv.ParseFloat(value_arr[0], 64); err == nil {
				ch <- prometheus.MustNewConstMetric(OutVoltageDesc,
					prometheus.GaugeValue, value, status.Status["Model Name"])
			}
		} else if k == "Test Result" {
			value_arr := strings.Fields(v)
			if value_arr[0] == "Passed" {
				ch <- prometheus.MustNewConstMetric(TestDesc,
					prometheus.GaugeValue, 1, status.Status["Model Name"])
			} else {
				ch <- prometheus.MustNewConstMetric(TestDesc,
					prometheus.GaugeValue, 0, status.Status["Model Name"])
			}
		}
	}
}

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":8088", "Address on which to expose metrics and web interface.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	)

	flag.Parse()
	fmt.Printf("Info:: \n Address: http://localhost%v%v \n::", *listenAddress, *metricsPath)

	pwrstatCollector := NewPwrstatCollector()
	prometheus.MustRegister(pwrstatCollector)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Sensor Exporter</title></head>
			<body>
			<h1>Sensor Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})
	http.ListenAndServe(*listenAddress, nil)

}
