package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	indexHandler(res, req)

	// if res.Code != http.StatusOK {
	// 	t.Fatal("Failed! ", res.Code)
	// }
	assert := assert.New(t)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello Nobody!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=jimmy", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello jimmy!", string(data))
}

func TestFooPathHandler_WithoutJson(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooPathHandler_WithJson(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"jimmy", "last_name":"yeon", "email":"jimmy.yeon@luxrobo.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert := assert.New(t)
	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("jimmy", user.FirstName)
	assert.Equal("yeon", user.LastName)
}
