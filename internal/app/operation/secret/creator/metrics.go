package creator

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	SecretCreateTotalAttempts prometheus.Counter
}

func NewMetrics(appName string) *Metrics {
	return &Metrics{
		SecretCreateTotalAttempts: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: appName,
			Name:      "creator_secret_create_total_attempts",
			Help:      "SecretCreator: total count of create attempts",
		}),
	}
}

func (m *Metrics) IncSecretCreateAttempts() {
	m.SecretCreateTotalAttempts.Inc()
}

func (m *Metrics) Collectors() []prometheus.Collector {
	return []prometheus.Collector{
		m.SecretCreateTotalAttempts,
	}
}
