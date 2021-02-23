package item

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gocraft/web"
	log "github.com/sirupsen/logrus"
)

func getItems(rw web.ResponseWriter, req *web.Request) {
	items, err := getAllItems()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write(itemsJSON)
}

func getItem(rw web.ResponseWriter, req *web.Request) {
	idStr := req.PathParams["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("ID has to be a number")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := getItemByID(id)
	if err == sql.ErrNoRows {
		rw.WriteHeader(http.StatusNotFound)
		return
	} 
	if err != nil{
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	dataJSON, err := json.Marshal(item)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write(dataJSON)
}