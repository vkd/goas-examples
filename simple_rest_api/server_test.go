package simple_rest_api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/vkd/goas-examples/simple_rest_api/handlers"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type PetModelTest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func TestPetsGet(t *testing.T) {
	authHeader := "test_token"
	pets := []PetModelTest{
		{1, "bird", "b"},
		{2, "dog", "d"},
		{3, "cat", "c"},
	}
	limit := 2

	fn := func(r handlers.PetsGetRequest) handlers.PetsGetResponser {
		if r.Headers.Authorization != authHeader {
			t.Errorf("Wrong header: %v", r.Headers.Authorization)
		}
		var out handlers.Pets
		for i, p := range pets {
			if int32(i) >= r.Query.Limit {
				break
			}
			out = append(out, handlers.Pet(p))
		}
		return handlers.PetsGetResponseJSON200(out)
	}

	r, err := http.NewRequest("GET", "/pets?limit="+strconv.Itoa(limit), nil)
	if err != nil {
		t.Fatalf("Error on create request: %v", err)
	}
	r.Header.Set("Authorization", authHeader)
	w := httptest.NewRecorder()

	handlers.PetsGetHandlerFunc(fn).ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Wrong status code: %v", w.Code)
	}
	var out []PetModelTest
	err = json.Unmarshal(w.Body.Bytes(), &out)
	if err != nil {
		t.Fatalf("Error on unmarshal body: %v", err)
	}
	if len(out) != limit {
		t.Fatalf("Wrong output count: %d", len(out))
	}
	for i, p := range pets[:limit] {
		if p != PetModelTest(out[i]) {
			t.Errorf("Wrong out pet[%d]: %v", i, out[i])
		}
	}
}
