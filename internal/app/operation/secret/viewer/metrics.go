package viewer

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	viewTotalAttempts prometheus.Counter
	searchFails       prometheus.Counter
	secretNotFound    prometheus.Counter
	secretFound       prometheus.Counter
	authPass          *prometheus.CounterVec
}

func NewMetrics(appName string) *Metrics {
	return &Metrics{
		viewTotalAttempts: prometheus.NewCounter(prometheus.CounterOpts{
			Name:      "viewer_view_total_attempts",
			Help:      "SecretViewer: View total attempts",
			Namespace: appName,
		}),
		searchFails: prometheus.NewCounter(prometheus.CounterOpts{
			Name:      "viewer_search_fails",
			Help:      "SecretViewer: Count of search fails",
			Namespace: appName,
		}),
		secretNotFound: prometheus.NewCounter(prometheus.CounterOpts{
			Name:      "viewer_view_secret_not_found",
			Help:      "SecretViewer: Count of secret not found",
			Namespace: appName,
		}),
		secretFound: prometheus.NewCounter(prometheus.CounterOpts{
			Name:      "viewer_view_secret_found",
			Help:      "SecretViewer: Count of secret found",
			Namespace: appName,
		}),
		authPass: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name:      "viewer_view_auth_pass",
			Help:      "SecretViewer: Auth Pass",
			Namespace: appName,
		}, []string{"factor", "pass"}),
	}
}

func (m *Metrics) IncViewTotalAttempts() {
	m.viewTotalAttempts.Inc()
}

func (m *Metrics) IncSearchFails() {
	m.searchFails.Inc()
}

func (m *Metrics) IncSecretNotFound() {
	m.secretNotFound.Inc()
}

func (m *Metrics) IncSecretFound() {
	m.secretFound.Inc()
}

func (m *Metrics) IncAuthPassOk(factor string) {
	m.authPass.WithLabelValues(factor, "true")
}

func (m *Metrics) IncAuthPassFail(factor string) {
	m.authPass.WithLabelValues(factor, "false")
}

func (m *Metrics) Collectors() []prometheus.Collector {
	return []prometheus.Collector{
		m.viewTotalAttempts,
		m.searchFails,
		m.secretNotFound,
		m.secretFound,
		m.authPass,
	}
}
