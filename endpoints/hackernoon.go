package endpoints

import "fmt"

//Hackernoon structure
type Hackernoon struct {
}

//Fetch fetches articles from hackernoon
func (h *Hackernoon) Fetch() ([]string, error) {

	fmt.Println("fetching!")
	return []string{"yolo"}, nil
}
