package tests

import (
	"encoding/json"
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)


func UserHandler(res http.ResponseWriter, req *http.Request) {
	testUser := storage.UserStruct{Name: "Kjetil", Password: "123"}
	data, _ := json.Marshal(testUser)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func TestCreateUser(t *testing.T) {
	request, _ := http.NewRequest("POST", "/user/signup", nil)
	response := httptest.NewRecorder()

	UserHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
	log.Print(response)

}