package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var correctResponses = []struct {
	CurrentPage int
	TotalPages  int
	Boundaries  int
	Around      int
	Wanted      string
}{
	{4, 10, 2, 2, "1 2 3 4 5 6 ... 9 10"},
	{4, 5, 1, 0, "1 ... 4 5"},
	{4, 10, 1, 1, "1 ... 3 4 5 ... 10"},
	{4, 10, 0, 1, "... 3 4 5 ..."},
	{4, 10, 1, 0, "1 ... 4 ... 10"},
	{4, 20, 0, 1, "... 3 4 5 ..."},
	{2, 3, 0, 1, "1 2 3"},
	{1, 1, 0, 0, "1"},
	{5, 10, 1, 3, "1 2 3 4 5 6 7 8 ... 10"},
}

func TestHealthCheckHandler(t *testing.T) {
	for _, tt := range correctResponses {
		requestString := fmt.Sprintf(
			"/?current_page=%d&total_pages=%d&boundaries=%d&around=%d",
			tt.CurrentPage,
			tt.TotalPages,
			tt.Boundaries,
			tt.Around,
		)
		req, err := http.NewRequest("GET", requestString, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		rh := RequestHandler{}
		handler := http.HandlerFunc(rh.getPagination)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if rr.Body.String() != tt.Wanted {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), tt.Wanted)
		}
	}
}
