package api

import (
	"fmt"

	"github.com/go-fuego/fuego"
	"github.com/teeaa/generics/internal/models"
	"github.com/teeaa/generics/internal/service"
)

func (a *Api) setRoutes() {
	v1 := fuego.Group(a.api, "/v1")
	tx := fuego.Group(v1, "/tx/{txid}")

	// fuego.Get(tx, "/address", NewOps[models.Address](a.service, "address").Get)
	setItemRoutes[models.Address](tx, a.service, "address")
	setItemRoutes[models.Dob](tx, a.service, "dob")
	setItemRoutes[models.Name](tx, a.service, "name")
}

func setItemRoutes[T models.DataType](route *fuego.Server, srv *service.Service, itemType string) {
	fuego.Get(route, fmt.Sprintf("/%s", itemType), NewOps[T](srv, itemType).Get)
}
