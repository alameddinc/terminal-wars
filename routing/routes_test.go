package routing

import (
	"bytes"
	"encoding/json"
	"github.com/alameddinc/terminal-wars/schemas"
	"net/http"
	"net/http/httptest"
	"testing"
)

var req *http.Request
var respRec *httptest.ResponseRecorder

func setupTests() {
	routes()
	respRec = httptest.NewRecorder()
}

func TestUserReqister(t *testing.T) {
	setupTests()
	bodySchema := schemas.UserRegisterPostSchema{
		Username: "alameddin",
		Password: "123123",
		Email:    "alameddinc@gmail.com",
	}

	body, err := json.Marshal(bodySchema)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(respRec, req)
	if respRec.Code != http.StatusOK {
		t.Fatal("Code Fail!")
	}
}
