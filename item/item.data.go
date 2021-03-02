package item

import (
	"context"
	"database/sql"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

// IRepository encapsulates the logic to access items from the data source.
type IRepository interface {
	GetByID(id int) (*Item, error)
	GetAll() ([]Item, error)
	Insert(newItem Item) error
	RemoveByID(id int) error
	UpdateExisting(item Item) error
}

type repository struct {
	db     *sql.DB
}

// NewRepository creates a new Item repository
func NewRepository(db *sql.DB) IRepository {
	return repository{db}
}

func (r repository) GetByID(id int) (*Item, error){
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	row := r.db.QueryRowContext(ctx, `SELECT 
	id, 
	manufacturer, 
	itemType, 
	price, 
	quantity
	FROM items
	WHERE id = ?`, id)

	item := Item{}
	err := row.Scan(&item.ID, 
					&item.Manufacturer, 
					&item.ItemType, 
					&item.Price, 
					&item.Quantity)
	if err != nil {
		return nil, err 
	}
	return &item, nil
}

func (r repository) GetAll() ([]Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	results, err := r.db.QueryContext(ctx, `SELECT 
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

func (r repository) Insert(newItem Item) error{
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `INSERT INTO items 
						(manufacturer, 
						itemType, 
						price, 
						quantity) VALUES (?, ?, ?, ?)`,  
						newItem.Manufacturer, 
						newItem.ItemType, 
						newItem.Price, 
						newItem.Quantity)
    return err
}

func (r repository) RemoveByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err := r.db.ExecContext(ctx, `DELETE FROM items where id = ?`, id)
	return err
}

func (r repository) UpdateExisting(item Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if item.ID == nil || *item.ID == 0 {
		return errors.New("Non Valid ID")
	}

	_, err := r.db.ExecContext(ctx, `UPDATE items SET
	manufacturer=?,
	itemType=?,
	price=?,
	quantity=?
	WHERE id=?`,  
	item.Manufacturer, 
	item.ItemType, 
	item.Price, 
	item.Quantity,
	item.ID)
	return err
}