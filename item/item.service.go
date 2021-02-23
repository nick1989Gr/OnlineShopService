package item

import (
	"encoding/json"
	"net/http"

	"github.com/gocraft/web"
	log "github.com/sirupsen/logrus"
)

func getItems(rw web.ResponseWriter, req *web.Request) {
	// MAYBE HEADERS ARE NEEDED FOR APPLICATION JSON
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