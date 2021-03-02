package paginate

import (
	"strconv"
	"strings"
)

type Pagination struct {
	CurrentPage int
	TotalPages  int
	Boundaries  int
	Around      int
}

// Return an array of pages tha shoud be linked
func (p *Pagination) GetPages() string {
	var output string
	threeDots := false

	for i := 1; i <= p.TotalPages; i++ {
		if i == p.CurrentPage || (p.isAroundSelected(i) || p.isBounderies(i)) {
			output += strconv.Itoa(i) + " "
			threeDots = false
		} else if !threeDots {
			output += "... "
			threeDots = true
		}
	}

	return strings.TrimSpace(output)
}

func (p *Pagination) isAroundSelected(page int) bool {
	return page <= (p.CurrentPage+p.Around) && page >= (p.CurrentPage-p.Around)
}

func (p *Pagination) isBounderies(page int) bool {
	return page <= p.Boundaries || page > (p.TotalPages-p.Boundaries)
}
