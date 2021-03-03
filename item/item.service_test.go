package item_test

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocraft/web"
	"github.com/nick1989Gr/OnlineShopService/item"
	model "github.com/nick1989Gr/OnlineShopService/item"
	"github.com/stretchr/testify/suite"
)
type ItemServiceTestSuite struct {
    suite.Suite
    repo *mockRepo
	service model.IService
	router *web.Router
}

func (suite *ItemServiceTestSuite) SetupTest() {
    suite.repo = NewMockRepo( []model.Item {
		{1, "Levis", "Trousers", 33.0, 5,},
		{2, "Nike", "Trousers", 34.0, 3,},
	}, "normal")
	suite.service = model.NewService(suite.repo)
	suite.router = item.InitRouter(suite.service)
}

func (suite *ItemServiceTestSuite) TestGetAllSucceeds() {
	// Before 
	response, request := newTestRequest("GET", "/item")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusOK)
	items,err := json.Marshal(suite.repo.items)
	suite.Equal(err, nil)
	suite.Equal(response.Body.String(), string(items))
}

func (suite *ItemServiceTestSuite) TestGetAllFails() {
	// Before 
	suite.repo.status = "error"
	response, request := newTestRequest("GET", "/item")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusInternalServerError)
	suite.Equal(response.Body.String(), "")
}

func (suite *ItemServiceTestSuite) TestGetByIDSucceeds() {
	// Before 
	response, request := newTestRequest("GET", "/item/1")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusOK)
	
	items,err := json.Marshal(suite.repo.items[0])
	suite.Equal(err, nil)
	suite.Equal(response.Body.String(), string(items))
}

func (suite *ItemServiceTestSuite) TestGetByIDFailsBadRequest() {
	// Before 
	response, request := newTestRequest("GET", "/item/sa")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusBadRequest)
}

func (suite *ItemServiceTestSuite) TestGetByIDFailsNotFound() {
	// Before 
	path := fmt.Sprintf("/item/%v", 1+ len(suite.repo.items))
	response, request := newTestRequest("GET", path)
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusNotFound)
}

func (suite *ItemServiceTestSuite) TestGetByIDFailsInternalServerError() {
	// Before 
	suite.repo.status = "error"
	response, request := newTestRequest("GET", "/item/1")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusInternalServerError)
}



func newTestRequest(method, path string) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()
	return recorder, request
}



type mockRepo struct {
	items []model.Item
	// normal for normal operation 
	// error for simulating error situation from DB
	status string 
}

func NewMockRepo(items []model.Item, status string) *mockRepo{
	return &mockRepo{items, status}
}

func (m *mockRepo) GetByID(id int) (*model.Item, error) {
	if m.status == "error" {
		return nil, errors.New("an error occured")
	}
	for _,i := range m.items {
		if i.ID == id {
			return &i, nil
		}
	}
	return nil, sql.ErrNoRows 
}

func (m *mockRepo) GetAll() ([]model.Item, error){
	if m.status == "normal" {
		return m.items, nil
	} 
	return nil, errors.New("an error occured")
}

func (m *mockRepo) Insert(newItem model.Item) error{
	newItem.ID = 3
	m.items = append(m.items, newItem)
	return nil
}

func (m *mockRepo) RemoveByID(id int) error{
	return nil	
}
func (m *mockRepo) UpdateExisting(item model.Item) error{
	return nil
}

func TestItemServiceTestSuite(t *testing.T) {
    suite.Run(t, new(ItemServiceTestSuite))
}