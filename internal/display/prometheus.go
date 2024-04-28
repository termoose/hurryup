package display

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"hurryup/internal/speed"
	"net/http"
	"time"
)

var gaugeDownloadSpeed = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "hurryup_download_speed",
	Help: "Download speed in kilobytes per second",
}, []string{"location", "country"})

var gaugeUploadSpeed = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "hurryup_upload_speed",
	Help: "Upload speed in kilobytes per second",
}, []string{"location", "country"})

var gaugeLatency = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "hurryup_latency",
	Help: "Latency in microseconds",
}, []string{"location", "country"})

func fetchMetrics(tester speed.Tester) {
	server := tester.GetServerData()
	labels := prometheus.Labels{"location": server.Name, "country": server.Country}

	for {
		metrics := tester.GetTesterData()

		gaugeDownloadSpeed.With(labels).Set(float64(metrics.DownloadRate))
		gaugeUploadSpeed.With(labels).Set(float64(metrics.UploadRate))
		gaugeLatency.With(labels).Set(float64(time.Duration(metrics.Latency) / time.Microsecond))

		time.Sleep(10 * time.Second)
	}
}

func Prometheus(tester speed.Tester) error {
	go fetchMetrics(tester)

	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":2112", nil)
}
