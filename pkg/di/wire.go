//go:build wireinject
// +build wireinject

package di

import (
	http "sample/pkg/api"
	"sample/pkg/api/handler"
	config "sample/pkg/config"
	"sample/pkg/db"
	"sample/pkg/repository"
	"sample/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserDataBase, usecase.NewuserUseCase, handler.NewUserHandler, http.NewServerHttp)

	return &http.ServerHTTP{}, nil
}
