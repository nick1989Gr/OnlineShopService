package item

import "github.com/gocraft/web"

type itemContext struct{}




// InitRouter Initializes a router for the item endpoint
func InitRouter() *web.Router {
	return web.New(itemContext{}).
				Middleware(web.LoggerMiddleware).
				Get("/item", getItems).
				Get("/item/:id", getItem)
}