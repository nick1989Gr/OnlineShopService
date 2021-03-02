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

// IService encapsulates usecase logic for items.
type IService interface  {
	GetAll(rw web.ResponseWriter, req *web.Request)
	GetByID(rw web.ResponseWriter, req *web.Request)
	Add(rw web.ResponseWriter, req *web.Request)
	Remove(rw web.ResponseWriter, req *web.Request)
	Update(rw web.ResponseWriter, req *web.Request)
}

type service struct {
	repository IRepository
}

// NewService creates a new Item service 
func NewService(repository IRepository) IService {
	return service{repository}
}

func (s service) GetAll(rw web.ResponseWriter, req *web.Request) {
	items, err := s.repository.GetAll()
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

func (s service) GetByID(rw web.ResponseWriter, req *web.Request) {
	idStr := req.PathParams["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("ID has to be a number")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := s.repository.GetByID(id)
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

func (s service) Add(rw web.ResponseWriter, req *web.Request) {
	var item Item
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.repository.Insert(item)
	if err != nil {
		log.Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (s service) Remove(rw web.ResponseWriter, req *web.Request) {
	idStr := req.PathParams["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("ID has to be a number")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = s.repository.RemoveByID(id)
	if err !=nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (s service) Update(rw web.ResponseWriter, req *web.Request) {
	var item Item
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		log.Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.repository.UpdateExisting(item)
	if err != nil {
		log.Error(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}