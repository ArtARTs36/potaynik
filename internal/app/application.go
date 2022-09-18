package app

import (
	"github.com/artarts36/potaynik/internal/app/operation/secret/auth"
	"github.com/artarts36/potaynik/internal/app/operation/secret/creator"
	"github.com/artarts36/potaynik/internal/app/operation/secret/viewer"
	"github.com/artarts36/potaynik/internal/app/repository"
	"github.com/artarts36/potaynik/internal/port/http/handlers"
	"github.com/vrischmann/envconfig"
)

type Application struct {
	Services struct {
		Http struct {
			Handlers struct {
				SecretCreateHandler *handlers.SecretCreateHandler
				SecretShowHandler   *handlers.SecretShowHandler
			}
		}
		Repositories struct {
			SecretRepository *repository.MemorySecretRepository
		}
		Operations struct {
			Secret struct {
				Creator *creator.Creator
				Viewer  *viewer.Viewer
			}
			Auth struct {
				Authorizers map[string]auth.Authorizer
			}
		}
	}
	Environment struct {
		Http struct {
			Public struct {
				Port int
			}
		}
	}
}

func NewApplication(appName string) (*Application, error) {
	app := &Application{}

	err := envconfig.InitWithPrefix(&app.Environment, appName)

	if err != nil {
		return nil, err
	}

	app.Services.Repositories.SecretRepository = repository.NewMemorySecretRepository()

	app.Services.Operations.Auth.Authorizers = map[string]auth.Authorizer{
		auth.PasswordAuthorizerKey: &auth.PasswordAuthorizer{},
	}

	app.Services.Operations.Secret.Creator = creator.NewCreator(
		app.Services.Repositories.SecretRepository,
		app.Services.Operations.Auth.Authorizers,
	)

	app.Services.Http.Handlers.SecretCreateHandler = handlers.NewSecretCreateHandler(app.Services.Operations.Secret.Creator)

	app.Services.Operations.Secret.Viewer = viewer.New(app.Services.Repositories.SecretRepository, app.Services.Operations.Auth.Authorizers)
	app.Services.Http.Handlers.SecretShowHandler = handlers.NewSecretShowHandler(app.Services.Operations.Secret.Viewer)

	return app, nil
}
