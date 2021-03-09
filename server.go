package main

import (
	"net/http"

	"github.com/nick1989Gr/OnlineShopService/database"
	"github.com/nick1989Gr/OnlineShopService/item"

	log "github.com/sirupsen/logrus"
)


func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	
}


func main() {
	database.SetupDatabase()
	connection := database.NewExternalSQLDb(database.DbConn)
	repository := item.NewRepository(connection)
	service := item.NewService(repository)
	router := item.InitRouter(service)
	log.Info("Server Started")
	http.ListenAndServe("localhost:8080", router) 
}