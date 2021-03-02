package item

import (
	"github.com/gocraft/web"
	"github.com/nick1989Gr/OnlineShopService/database"
)

type itemContext struct{}




// InitRouter Initializes a router for the item endpoint
func InitRouter() *web.Router {

	service := NewService(NewRepository(database.DbConn))

	return web.New(itemContext{}).
				Middleware(web.LoggerMiddleware).
				Get("/item", service.GetAll).
				Get("/item/:id", service.GetByID).
				Post("/item", service.Add).
				Delete("item/:id", service.Remove).
				Put("item", service.Update)
}