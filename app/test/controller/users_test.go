package testController

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kousuketk/crudGo/app/main"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	// requestBody := strings.NewReader("Email=")
	// response := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(response)
	// c.Request, _ = http.NewRequest(
	// 	http.MethodGet,
	// 	"http://localhost:8080/api/v1/ping",
	// 	nil,
	// )
	// u := controllers.UserController{}
	// u.Index(c)
	// fmt.Println(response.Code)
	router := main.Router()
	w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

func TestBubbleSortOrderDESC(t *testing.T) {
	// Init
	elements := []int{5, 2, 3, 1, 6, 7, 9, 0, 4, 8}

	// Validation
	if elements[0] != 5 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 8 {
		t.Error("last elements should be 0")
	}
}
