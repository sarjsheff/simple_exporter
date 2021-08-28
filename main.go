package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	aliveMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "devmon_alive",
		Help: "Alive.",
	},
		[]string{"url"},
	)
	resStatusMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "devmon_res_status",
		Help: "Response status code.",
	},
		[]string{"url", "code"},
	)
	errorsMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "devmon_errors",
		Help: "Errors.",
	},
		[]string{"url", "error"},
	)
)

func httpChecker(url string) {
	log.Printf("Start http checker [%s]\n", url)
	client := http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}
	for {
		resp, err := client.Get(url)
		if err != nil {
			aliveMetric.With(prometheus.Labels{"url": url}).Set(0)
			if nerr, ok := err.(net.Error); ok {
				if nerr.Timeout() {
					errorsMetric.With(prometheus.Labels{"url": url, "error": "timeout"}).Inc()
				} else {
					var dnsError *net.DNSError
					if errors.As(err, &dnsError) {
						errorsMetric.With(prometheus.Labels{"url": url, "error": "dns"}).Inc()
					} else {
						errorsMetric.With(prometheus.Labels{"url": url, "error": "network"}).Inc()
					}
				}
			} else {
				log.Println(err)
				errorsMetric.With(prometheus.Labels{"url": url, "error": "other"}).Inc()
			}
		} else {
			aliveMetric.With(prometheus.Labels{"url": url}).Set(1)
			resStatusMetric.With(prometheus.Labels{"url": url, "code": strconv.Itoa(resp.StatusCode)}).Inc()
		}
		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}
}

func main() {
	if config() > -1 {
		log.Println(cfg.Urls)
		for _, v := range cfg.Urls {
			u, err := url.Parse(v)
			if err != nil {
				log.Fatal(err)
			}
			switch {
			case u.Scheme == "https" || u.Scheme == "http":
				go httpChecker(u.String())
			default:
				log.Println("unknown")
			}
		}
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(cfg.Listen, nil)
	}
}
