package item_test

import (
	"bytes"
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

const(
	NEW_ITEM_ID = 3
	MOCK_ERROR_STATUS = "error"
	MOCK_NORMAL_STATUS = "normal"
)

func (suite *ItemServiceTestSuite) SetupTest() {
    suite.repo = NewMockRepo( []model.Item {
		{1, "Levis", "Trousers", 33.0, 5,},
		{2, "Nike", "Trousers", 34.0, 3,},
	}, MOCK_NORMAL_STATUS)
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
	suite.repo.status = MOCK_ERROR_STATUS
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
	suite.repo.status = MOCK_ERROR_STATUS
	response, request := newTestRequest("GET", "/item/1")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusInternalServerError)
}

func (suite *ItemServiceTestSuite) TestAddSucceeds() {
	// Before 
	expected := fmt.Sprintf(`{"id":%v,"manufacturer":"Diesel","itemType":"Trousers","price":45,"quantity":6}`, NEW_ITEM_ID)
	bodyRequest := `{"manufacturer":"Diesel","itemType":"Trousers","price":45,"quantity":6}`
	request, _ := http.NewRequest("POST", "/item/", bytes.NewReader([]byte(bodyRequest)))
	response := httptest.NewRecorder()
	
	// Action
	suite.router.ServeHTTP(response, request)
	
	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusCreated)
	actual,err := suite.repo.GetByID(NEW_ITEM_ID)
	suite.Equal(err, nil)
	actualBytes,err := json.Marshal(actual)
	suite.Equal(err, nil)
	suite.Equal(actualBytes, []byte(expected))
}

func (suite *ItemServiceTestSuite) TestAddFails() {
	// Before 
	suite.repo.status = MOCK_ERROR_STATUS
	bodyRequest := `{"manufacturer":"Diesel","itemType":"Trousers","price":45,"quantity":6}`
	request, _ := http.NewRequest("POST", "/item/", bytes.NewReader([]byte(bodyRequest)))
	response := httptest.NewRecorder()
	
	// Action
	suite.router.ServeHTTP(response, request)
	
	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusInternalServerError)
}

func (suite *ItemServiceTestSuite) TestRemoveSucceeds() {
	// Before 
	response, request := newTestRequest("DELETE", "/item/1")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusOK)
	_,err:=suite.repo.GetByID(1)
	suite.Equal(err, sql.ErrNoRows)
}

func (suite *ItemServiceTestSuite) TestRemoveFailsDueToBadRequest() {
	// Before 
	response, request := newTestRequest("DELETE", "/item/1as")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusBadRequest)
}

func (suite *ItemServiceTestSuite) TestRemoveFailsDueToDbError() {
	// Before 
	response, request := newTestRequest("DELETE", "/item/156")
	
	// Action
	suite.router.ServeHTTP(response, request)

	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusInternalServerError)
}

func (suite *ItemServiceTestSuite) TestUpdateSucceeds() {
	// Before 
	const UPDATED_INDEX = 1
	expected := fmt.Sprintf(`{"id":%v,"manufacturer":"Levis","itemType":"Trousers","price":45,"quantity":6}`, UPDATED_INDEX)
	request, _ := http.NewRequest("PUT", "/item/", bytes.NewReader([]byte(expected)))
	response := httptest.NewRecorder()
	
	// Action
	suite.router.ServeHTTP(response, request)
	
	// Verify 
	suite.Equal(response.Result().StatusCode, http.StatusOK)
	actual,err := suite.repo.GetByID(UPDATED_INDEX)
	suite.Equal(err, nil)
	actualBytes,err := json.Marshal(actual)
	suite.Equal(err, nil)
	suite.Equal(actualBytes, []byte(expected))
}

func (suite *ItemServiceTestSuite) TestUpdateFails() {
	// Before 
	const UPDATED_INDEX = 56
	expected := fmt.Sprintf(`{"id":%v,"manufacturer":"Levis","itemType":"Trousers","price":45,"quantity":6}`, UPDATED_INDEX)
	request, _ := http.NewRequest("PUT", "/item/", bytes.NewReader([]byte(expected)))
	response := httptest.NewRecorder()
	
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
	if m.status == MOCK_ERROR_STATUS {
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
	if m.status == MOCK_NORMAL_STATUS {
		return m.items, nil
	} 
	return nil, errors.New("an error occured")
}

func (m *mockRepo) Insert(newItem model.Item) error{
	if m.status == MOCK_ERROR_STATUS {
		return errors.New("an error occured")
	}
	newItem.ID = NEW_ITEM_ID
	m.items = append(m.items, newItem)
	return nil
}

func (m *mockRepo) RemoveByID(id int) error{
	for i,item := range m.items {
		if item.ID == id {
			m.items = append(m.items[:i], m.items[i+1:]...)
			return nil
		}
	}
	return errors.New("an error occured")	
}
func (m *mockRepo) UpdateExisting(item model.Item) error{
	for ind, i := range m.items {
		if i.ID == item.ID {
			m.items[ind] = item
			return nil
		}
	}
	return errors.New("an error occured")
}

func TestItemServiceTestSuite(t *testing.T) {
    suite.Run(t, new(ItemServiceTestSuite))
}