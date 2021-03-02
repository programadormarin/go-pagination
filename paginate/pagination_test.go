package paginate

import (
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

func TestGetPages(t *testing.T) {
	for _, tt := range correctResponses {
		p := Pagination{
			CurrentPage: tt.CurrentPage,
			TotalPages:  tt.TotalPages,
			Boundaries:  tt.Boundaries,
			Around:      tt.Around,
		}

		result := p.GetPages()

		if result != tt.Wanted {
			t.Fatalf("got %q, wanted %q", result, tt.Wanted)
		}
	}
}
