package endpoints

import "fmt"

//Hackernoon structure
type Hackernoon struct {
}

//Fetch fetches articles from hackernoon
func (h *Hackernoon) Fetch(args []string) ([]string, error) {

	fmt.Println("fetching!")
	return []string{
		"https://google.com?q=" + args[0],
		"https://google.com?q=" + args[0],
	}, nil
}
