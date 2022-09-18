package creator

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	secretCreateTotalAttempts   prometheus.Counter
	secretCreateSuccessAttempts prometheus.Counter
	useAuthFactors              *prometheus.CounterVec
}

func NewMetrics(appName string) *Metrics {
	return &Metrics{
		secretCreateTotalAttempts: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: appName,
			Name:      "creator_secret_create_total_attempts",
			Help:      "SecretCreator: total count of create attempts",
		}),
		secretCreateSuccessAttempts: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: appName,
			Name:      "creator_secret_create_success_attempts",
			Help:      "SecretCreator: success count of create attempts",
		}),
		useAuthFactors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: appName,
			Name:      "creator_secret_create_use_auth_factors",
		}, []string{"factor"}),
	}
}

func (m *Metrics) IncSecretCreateAttempts() {
	m.secretCreateTotalAttempts.Inc()
}

func (m *Metrics) IncSecretCreateSuccessAttempts() {
	m.secretCreateSuccessAttempts.Inc()
}

func (m *Metrics) IncUseAuthFactor(factor string) {
	m.useAuthFactors.WithLabelValues(factor).Inc()
}

func (m *Metrics) Collectors() []prometheus.Collector {
	return []prometheus.Collector{
		m.secretCreateTotalAttempts,
		m.secretCreateSuccessAttempts,
		m.useAuthFactors,
	}
}
