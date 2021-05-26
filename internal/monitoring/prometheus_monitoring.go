package prometheus_monitoring

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
	"github.com/RektangularStudios/novellia/internal/config"
)

// https://prometheus.io/docs/guides/go-application/

const (
	microservice_namespace = "novellia"
	status_interval = 30 * time.Second
)

struct PrometheusMetrics {
	microserviceStatusMetric prometheus.Gauge
	cardanoStatusMetric prometheus.Gauge
	productIDsListedMetric prometheus.Gauge
}
var (
	initialized = false
	prometheusMetrics PrometheusMetrics
)

type statusIndicators struct {
	microserviceStatus float64
	cardanoStatus float64
}

func Init(namespace string) {
	n := microservice_namespace
	if namespace != "" {
		n = fmt.Sprintf("%s_%s", namespace, microservice_namespace)
	}

	prometheusMetrics.microserviceStatusMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: n,
		Name: "microservice_status",
		Help: "Health status indicator for the Novellia microservice",
	})
	cardanoStatusMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: n,
		Name: "cardano_status",
		Help: "Health status indicator for Cardano services such as GraphQL and cardano-node",
	})
	productIDsListedMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: n,
		Name: "products_ids_listed",
		Help: "Number of products IDs returned when accessing Novellia",
	})

	initialized = true
}

func getStatus() (*statusIndicators, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config from env")
	}

	req, err := http.NewRequest("GET", config.Monitoring.StatusURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status health check failed: %+v", resp)
	}

	var respBody nvla.Status
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, err
	}

	s := statusIndicators{
		microserviceStatus: 0,
		cardanoStatus: 0,
	}
	if respBody.Status == "UP" {
		s.microserviceStatus = 1
	}
	if respBody.Cardano.Initialized {
		s.cardanoStatus = float64(respBody.Cardano.SyncPercentage)
	}
	fmt.Printf("Checked Status, result: %+v\n", respBody)

	return &s, nil
}

func RecordNumberOfProductIDsListed(count int) {
	if initialized {
		prometheusMetrics.productIDsListedMetric.Set(float64(count))
	}
}

func RecordMetrics() {
	go func() {
		for {
			if initialized {
				indicators, err := getStatus()
				if err != nil {
					indicators = &statusIndicators{
						microserviceStatus: 0,
						cardanoStatus: 0,
					}
					fmt.Printf("Checked status, got error: %+v\n", err)
				}
				
				microserviceStatusMetric.Set(indicators.microserviceStatus)
				cardanoStatusMetric.Set(indicators.cardanoStatus)
			}
			time.Sleep(status_interval)
		}
	}()
}
