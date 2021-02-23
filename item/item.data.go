package item

import (
	"context"
	"time"

	"github.com/nick1989Gr/OnlineShopService/database"
	log "github.com/sirupsen/logrus"
)

func getAllItems() ([]Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	results, err := database.DbConn.QueryContext(ctx, `SELECT 
	id, 
	manufacturer, 
	itemType, 
	price, 
	quantity
	FROM items`)
	defer results.Close()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var items []Item
	for results.Next() {
		var item Item
		results.Scan(&item.ID, 
					 &item.Manufacturer, 
					 &item.ItemType, 
					 &item.Price, 
					 &item.Quantity)
		items = append(items, item)
	}
	return items, nil

}
