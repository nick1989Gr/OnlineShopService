package item_test

import (
	"database/sql"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	item "github.com/nick1989Gr/OnlineShopService/item"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)


type ItemDataTestSuite struct {
	suite.Suite
	db *sql.DB
	mock sqlmock.Sqlmock
	repo item.IRepository
}

func (suite *ItemDataTestSuite) SetupTest() {
	suite.db, suite.mock = NewMock()
	suite.repo = item.NewRepository(suite.db)
}


func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestItemDataTestSuite(t *testing.T) {
    suite.Run(t, new(ItemDataTestSuite))
}

func (suite *ItemDataTestSuite) TestGetByIDSucceeds() {
	// Before
	query := `SELECT id, manufacturer, itemType, price, quantity
	FROM items WHERE id = ?`
	expected := item.Item{1, "Levis", "Trousers", 33.0, 5}
	rows := sqlmock.NewRows([]string{"id", "manufacturer", "itemType", "price","quantity"}).
	   AddRow(expected.ID, expected.Manufacturer, expected.ItemType, expected.Price, expected.Quantity)
	suite.mock.ExpectQuery(query).WithArgs(expected.ID).WillReturnRows(rows)
 
	// Action
	actual, err := suite.repo.GetByID(expected.ID)

	// Verify
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), &expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *ItemDataTestSuite) TestGetByIDFails() {
	// Before
	query := `SELECT id, manufacturer, itemType, price, quantity
	FROM items WHERE id = ?`
	rows := sqlmock.NewRows([]string{"id", "manufacturer", "itemType", "price","quantity"})
	suite.mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	
	// Action
	actual, err := suite.repo.GetByID(1)

	// Verify
	assert.Nil(suite.T(), actual)
	assert.Equal(suite.T(), sql.ErrNoRows, err)
}

func (suite *ItemDataTestSuite) TestGetAll() {
	// Before
	query := `SELECT id, manufacturer, itemType, price, quantity FROM items`
	expected := []item.Item {
		{1, "Levis", "Trousers", 33.0, 5,},
		{2, "Nike", "Trousers", 34.0, 3,},
	}
	rows := sqlmock.NewRows([]string{"id", "manufacturer", "itemType", "price","quantity"})
	for i:=0;i< len(expected);i++ {
		rows.AddRow(expected[i].ID, expected[i].Manufacturer, expected[i].ItemType, expected[i].Price, expected[i].Quantity)
	}
	   
	suite.mock.ExpectQuery(query).WillReturnRows(rows)
 
	// Action
	actual, err := suite.repo.GetAll()

	// Verify
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), expected, actual)
	assert.NoError(suite.T(), err)
}

func (suite *ItemDataTestSuite) TestGetAllFails() {
	// Before
	query := `SELECT id, manufacturer, itemType, price, quantity FROM items`
	// e := errors.New("an error occured")
	rows := sqlmock.NewRows([]string{"id", "manufacturer", "itemType", "price","quantity"})
	suite.mock.ExpectQuery(query).WillReturnRows(rows)
 
	// Action
	actual, err := suite.repo.GetAll()

	// Verifys
	assert.Nil(suite.T(), actual)
	assert.Nil(suite.T(), err)
}

func (suite *ItemDataTestSuite) TestInsert() {
	// Before
	query := `INSERT INTO items (manufacturer, itemType, price, quantity) VALUES (?, ?, ?, ?)`
	newItem := item.Item{1, "Levis", "Trousers", 33.0, 5}
	
	suite.mock.ExpectExec(query).
	WithArgs(newItem.Manufacturer, newItem.ItemType, newItem.Price, newItem.Quantity).
	WillReturnResult(sqlmock.NewResult(0, 1))
				
	// Action
	err := suite.repo.Insert(newItem)

	// Verify
	assert.Nil(suite.T(), err)
}

func (suite *ItemDataTestSuite) TestInsertFail() {
	// Before
	query := `INSERT INTO items (manufacturer, itemType, price, quantity) VALUES (?, ?, ?, ?)`
	newItem := item.Item{1, "Levis", "Trousers", 33.0, 5}
	expected := errors.New("an error occured")
	suite.mock.ExpectExec(query).
	WithArgs(newItem.Manufacturer, newItem.ItemType, newItem.Price, newItem.Quantity).
	WillReturnResult(sqlmock.NewResult(0, 1)).
	WillReturnError(expected)
				
	// Action
	actual := suite.repo.Insert(newItem)

	// Verify
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ItemDataTestSuite) TestRemove() {
	// Before
	query := `DELETE FROM items where id = ?`
	
	suite.mock.ExpectExec(query).
	WithArgs(1).
	WillReturnResult(sqlmock.NewResult(0, 1))
				
	// Action
	err := suite.repo.RemoveByID(1)

	// Verify
	assert.Nil(suite.T(), err)
}

func (suite *ItemDataTestSuite) TestRemoveFails() {
	// Before
	query := `DELETE FROM items where id = ?`
	expected := errors.New("an error occured")
	suite.mock.ExpectExec(query).
	WithArgs(1).
	WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(expected)
				
	// Action
	actual := suite.repo.RemoveByID(1)

	// Verify
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ItemDataTestSuite) TestCreate() {
	// Before
	query := `UPDATE items SET manufacturer=?, itemType=?, price=?, quantity=? WHERE id=?`
	updatedItem := item.Item{1, "Levis", "Trousers", 33.0, 5}
	
	suite.mock.ExpectExec(query).
	WithArgs(updatedItem.Manufacturer, updatedItem.ItemType, updatedItem.Price, updatedItem.Quantity, updatedItem.ID).
	WillReturnResult(sqlmock.NewResult(0, 1))
				
	// Action
	err := suite.repo.UpdateExisting(updatedItem)

	// Verify
	assert.Nil(suite.T(), err)
}

func (suite *ItemDataTestSuite) TestCreateFailsDueToInvalidId() {
	// Before
	query := `UPDATE items SET manufacturer=?, itemType=?, price=?, quantity=? WHERE id=?`
	updatedItem := item.Item{0, "Levis", "Trousers", 33.0, 5}
	
	suite.mock.ExpectExec(query).
	WithArgs(updatedItem.Manufacturer, updatedItem.ItemType, updatedItem.Price, updatedItem.Quantity, updatedItem.ID).
	WillReturnResult(sqlmock.NewResult(0, 1))
				
	// Action
	actual := suite.repo.UpdateExisting(updatedItem)

	// Verify
	assert.Equal(suite.T(), errors.New("Non Valid ID"), actual)
}

func (suite *ItemDataTestSuite) TestCreateFails() {
	// Before
	query := `UPDATE items SET manufacturer=?, itemType=?, price=?, quantity=? WHERE id=?`
	updatedItem := item.Item{1, "Levis", "Trousers", 33.0, 5}
	expected := errors.New("Non Valid ID")
	suite.mock.ExpectExec(query).
	WithArgs(updatedItem.Manufacturer, updatedItem.ItemType, updatedItem.Price, updatedItem.Quantity, updatedItem.ID).
	WillReturnResult(sqlmock.NewResult(0, 1)).
	WillReturnError(expected)
				
	// Action
	actual := suite.repo.UpdateExisting(updatedItem)

	// Verify
	assert.Equal(suite.T(), expected, actual )
}