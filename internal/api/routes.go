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

	setItemRoutes[models.Address, models.AddressRequest, models.AddressResponse](tx, a.service, "address")
	setItemRoutes[models.Dob, models.DobRequest, models.DobResponse](tx, a.service, "dob")
	setItemRoutes[models.Name, models.NameRequest, models.NameResponse](tx, a.service, "name")
}

func setItemRoutes[T models.DataType, Req models.RequestType[T], Res models.ResponseType[T]](route *fuego.Server, srv *service.Service, itemType string) {
	fuego.Post(route, fmt.Sprintf("/%s", itemType), NewOps[T, Req, Res](srv).Create)
	fuego.Get(route, fmt.Sprintf("/%s", itemType), NewOps[T, Req, Res](srv).Get)
}
