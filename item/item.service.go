package item

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gocraft/web"
	log "github.com/sirupsen/logrus"
)


func getItems(rw web.ResponseWriter, req *web.Request) {
	items, err := getAllItems()
	fmt.Println(items)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		log.Error(err)
		return
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

func addItem(rw web.ResponseWriter, req *web.Request) {
	var item Item
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = insertItem(item)
	if err != nil {
		log.Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func removeItem(rw web.ResponseWriter, req *web.Request) {
	idStr := req.PathParams["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("ID has to be a number")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = removeItemByID(id)
	if err !=nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func updateItem(rw web.ResponseWriter, req *web.Request) {
	var item Item
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		log.Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = updateExistingItem(item)
	if err != nil {
		log.Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}