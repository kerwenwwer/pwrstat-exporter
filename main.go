package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	gopwrstat "github.com/kerwenwwer/pwrstat-exporter/pwrstat"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("web.listen-address", ":8088", "Address on which to expose metrics and web interface.")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	// Metric Descriptors
	LoadDesc = prometheus.NewDesc(
		"ups_load",
		"UPS power load (Watt)",
		[]string{"device"},
		nil,
	)

	StateDesc = prometheus.NewDesc(
		"ups_state",
		"UPS status (1 -> Normal, 0 -> Not)",
		[]string{"device"},
		nil,
	)

	BatteryDesc = prometheus.NewDesc(
		"ups_battery_capacity",
		"UPS battery capacity(%)",
		[]string{"device"},
		nil,
	)

	RtimeDesc = prometheus.NewDesc(
		"ups_remaining_runtime",
		"UPS Remaining Runtime(min)",
		[]string{"device"},
		nil,
	)

	InVoltageDesc = prometheus.NewDesc(
		"ups_in_voltage",
		"UPS Input Voltage(V): pass-> 1 non_pass -> 0",
		[]string{"device"},
		nil,
	)

	OutVoltageDesc = prometheus.NewDesc(
		"ups_out_voltage",
		"UPS Output Voltage(V): pass-> 1 non_pass -> 0",
		[]string{"device"},
		nil,
	)

	TestDesc = prometheus.NewDesc(
		"ups_test_result",
		"UPS Test Result",
		[]string{"device"},
		nil,
	)
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("Info: Serving metrics on http://localhost%s%s\n", *listenAddress, *metricsPath)

	pwrstatCollector := NewPwrstatCollector()
	prometheus.MustRegister(pwrstatCollector)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head><title>UPS Exporter</title></head>
            <body>
            <h1>UPS Exporter</h1>
            <p><a href="` + *metricsPath + `">Metrics</a></p>
            </body>
            </html>`))
	})

	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}

type PwrstatCollector struct{}

func NewPwrstatCollector() *PwrstatCollector {
	return &PwrstatCollector{}
}

func (collector *PwrstatCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(collector, ch)
}

func (collector *PwrstatCollector) Collect(ch chan<- prometheus.Metric) {
	status, err := gopwrstat.NewFromSystem()
	if err != nil {
		fmt.Printf("Error fetching UPS status: %s\n", err)
		return
	}

	parseAndCollectMetrics(status, ch)
}

func parseAndCollectMetrics(status *gopwrstat.Pwrstat, ch chan<- prometheus.Metric) {
	model := status.Status["Model Name"]
	for k, v := range status.Status {
		switch k {
		case "Load", "Battery Capacity", "Remaining Runtime", "Utility Voltage", "Output Voltage":
			collectNumericMetric(k, v, model, ch)
		case "State":
			collectStateMetric(v, model, ch)
		case "Test Result":
			collectTestResultMetric(v, model, ch)
		}
	}
}

func collectNumericMetric(metricType, valueStr, model string, ch chan<- prometheus.Metric) {
	valueFields := strings.Fields(valueStr)
	if value, err := strconv.ParseFloat(valueFields[0], 64); err == nil {
		desc := getMetricDesc(metricType)
		if desc != nil {
			ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, value, model)
		}
	}
}

func collectStateMetric(value, model string, ch chan<- prometheus.Metric) {
	state := 0.0
	if value == "Normal" {
		state = 1.0
	}
	ch <- prometheus.MustNewConstMetric(StateDesc, prometheus.GaugeValue, state, model)
}

func collectTestResultMetric(value, model string, ch chan<- prometheus.Metric) {
	result := 0.0
	if value == "Passed" {
		result = 1.0
	}
	ch <- prometheus.MustNewConstMetric(TestDesc, prometheus.GaugeValue, result, model)
}

func getMetricDesc(metricType string) *prometheus.Desc {
	switch metricType {
	case "Load":
		return LoadDesc
	case "Battery Capacity":
		return BatteryDesc
	case "Remaining Runtime":
		return RtimeDesc
	case "Utility Voltage":
		return InVoltageDesc
	case "Output Voltage":
		return OutVoltageDesc
	default:
		return nil
	}
}
