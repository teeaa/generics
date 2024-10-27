package api

import (
	"fmt"

	"github.com/go-fuego/fuego"
	"github.com/teeaa/generics/internal/config"
	"github.com/teeaa/generics/internal/service"
)

type Api struct {
	service *service.Service
	api     *fuego.Server
}

func New(conf *config.Config, srv *service.Service) (*Api, error) {
	server := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf("%s:%d", conf.Service.Address, conf.Service.Port)),
	)

	api := &Api{
		service: srv,
		api:     server,
	}

	api.setRoutes()

	err := server.Run()
	if err != nil {
		return nil, err
	}

	return api, nil
}

func (a *Api) Stop() error {
	return a.api.Close()
}
