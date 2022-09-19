package app

import (
	"fmt"
	"github.com/artarts36/potaynik/internal/app/operation/health"
	"github.com/artarts36/potaynik/internal/app/operation/secret/informer"

	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vrischmann/envconfig"

	"github.com/artarts36/potaynik/internal/app/operation/secret/auth"
	"github.com/artarts36/potaynik/internal/app/operation/secret/creator"
	"github.com/artarts36/potaynik/internal/app/operation/secret/viewer"
	"github.com/artarts36/potaynik/internal/app/repository"
	"github.com/artarts36/potaynik/internal/port/http/handlers"
)

type Application struct {
	AppName  string
	Services struct {
		Http struct {
			Handlers struct {
				SecretCreateHandler *handlers.SecretCreateHandler
				SecretShowHandler   *handlers.SecretShowHandler
				SecretInfoHandler   *handlers.SecretInfoHandler
				HealthCheckHandler  *handlers.HealthCheckHandler
			}
		}
		Repositories struct {
			SecretRepository *repository.RedisSecretRepository
		}
		Operations struct {
			Secret struct {
				Creator  *creator.Creator
				Viewer   *viewer.Viewer
				Informer *informer.Informer
			}
			Auth struct {
				Authorizers map[string]auth.Authorizer
			}
		}
		HealthCheckers []health.Checker
		Redis          *redis.Client
	}
	Environment struct {
		Http struct {
			Public struct {
				Port int
			}
			Health struct {
				Port int
			}
		}
		Redis struct {
			Addr string
		}
	}
	Metrics struct {
		Collectors struct {
			SecretCreatorMetrics *creator.Metrics
			SecretViewerMetrics  *viewer.Metrics
		}
		Registry *prometheus.Registry
	}
}

func NewApplication(appName string) (*Application, error) {
	app := &Application{}
	app.AppName = appName

	err := envconfig.InitWithPrefix(&app.Environment, appName)

	if err != nil {
		return nil, err
	}

	app.registerMetrics()
	app.connectRedis()

	app.Services.Repositories.SecretRepository = repository.NewRedisSecretRepository(app.Services.Redis, fmt.Sprintf("%s_", appName))

	app.Services.Operations.Auth.Authorizers = map[string]auth.Authorizer{
		auth.PasswordAuthorizerKey: &auth.PasswordAuthorizer{},
	}

	app.Services.Operations.Secret.Creator = creator.NewCreator(
		app.Services.Repositories.SecretRepository,
		app.Services.Operations.Auth.Authorizers,
		app.Metrics.Collectors.SecretCreatorMetrics,
	)

	app.Services.Http.Handlers.SecretCreateHandler = handlers.NewSecretCreateHandler(app.Services.Operations.Secret.Creator)

	app.Services.Operations.Secret.Viewer = viewer.New(
		app.Services.Repositories.SecretRepository,
		app.Services.Operations.Auth.Authorizers,
		app.Metrics.Collectors.SecretViewerMetrics,
	)

	app.Services.Http.Handlers.SecretShowHandler = handlers.NewSecretShowHandler(app.Services.Operations.Secret.Viewer)

	app.Services.Operations.Secret.Informer = informer.New(app.Services.Repositories.SecretRepository)

	app.Services.Http.Handlers.SecretInfoHandler = handlers.NewSecretInfoHandler(app.Services.Operations.Secret.Informer)

	app.Services.Http.Handlers.HealthCheckHandler = handlers.NewHealthCheckHandler(app.Services.HealthCheckers)

	return app, nil
}

func (app *Application) registerMetrics() {
	app.spawnMetrics()

	app.Metrics.Registry = prometheus.NewRegistry()

	app.Metrics.Registry.MustRegister(app.Metrics.Collectors.SecretCreatorMetrics.Collectors()...)
	app.Metrics.Registry.MustRegister(app.Metrics.Collectors.SecretViewerMetrics.Collectors()...)
}

func (app *Application) spawnMetrics() {
	app.Metrics.Collectors.SecretCreatorMetrics = creator.NewMetrics(app.AppName)
	app.Metrics.Collectors.SecretViewerMetrics = viewer.NewMetrics(app.AppName)
}

func (app *Application) connectRedis() {
	app.Services.Redis = redis.NewClient(&redis.Options{
		Addr: app.Environment.Redis.Addr,
	})

	app.Services.HealthCheckers = append(app.Services.HealthCheckers, health.NewRedisChecker(app.Services.Redis))
}
