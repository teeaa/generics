## Links

### First iteration

`git checout first`

- [Generics DataType](./internal/models/generics.go)
- [Create route](./internal/api/routes.go)
- [API Create()](./internal/api/ops.go)

`go run main.go`

> See request data

---

### Second iteration

`git checkout second`

- [Generics DataType](./internal/models/generics.go)
- [Address model](./internal/models/address.go)
- [API Create()](./internal/api/ops.go)

`go run main.go`

> See empty base object

- [API Create()](./internal/api/ops.go)

> Uncomment code, SetBase()

- [Generics DataType](./internal/models/generics.go)

> Add SetBase() to DataType interface

---

### Third iteration

`git checkout third`

- [Generics Item](./internal/models/generics.go)
- [API Create()](./internal/api/ops.go)

`go run main.go`

> See SetItem()

- [API Create()](./internal/api/ops.go)

> Partially uncomment code

- [API Create()](./internal/api/ops.go)

`go run main.go`

> See Item

- [API Create()](./internal/api/ops.go)

> Comment out the rest of the code

> Add Ops functionality to service and database layers

- [Service Create()](./internal/service/ops.go)
- [Database Create()](./internal/database/ops.go)

`go run main.go`

---

### Fourth iteration

`git checkout fourth`

- [New Generic types](./internal/models/generics.go)
- [New models for types](./internal/models/address.go)
- [Update routes](./internal/api/routes.go)
- [Update API Ops](./internal/api/ops.go)

`go run main.go`
