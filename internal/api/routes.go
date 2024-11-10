package api

import (
	"fmt"

	"github.com/go-fuego/fuego"
	"github.com/teeaa/generics/internal/models"
)

func (a *Api) setRoutes() {
	v1 := fuego.Group(a.api, "/v1")
	tx := fuego.Group(v1, "/tx/{txid}")

	setItemRoutes[models.Address](tx, "address")
	setItemRoutes[models.Dob](tx, "dob")
	setItemRoutes[models.Name](tx, "name")
}

// func (a *Api) setItemRoutes[T models.DataType](route *fuego.Server, itemType string) {
func setItemRoutes[T models.DataType](route *fuego.Server, itemType string) {
	fuego.Post(route, fmt.Sprintf("/%s", itemType), Create[T])
}
