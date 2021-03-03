package item

import (
	"github.com/gocraft/web"
)

type itemContext struct{}


// InitRouter Initializes a router for the item endpoint
func InitRouter(service IService) *web.Router {


	return web.New(itemContext{}).
				Middleware(web.LoggerMiddleware).
				Get("/item", service.GetAll).
				Get("/item/:id", service.GetByID).
				Post("/item", service.Add).
				Delete("item/:id", service.Remove).
				Put("item", service.Update)
}