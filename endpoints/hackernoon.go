package endpoints

import "fmt"

//Hackernoon structure
type Hackernoon struct {
}

//Fetch fetches articles from hackernoon
func (h *Hackernoon) Fetch(args []string) ([]string, error) {

	fmt.Println("fetching!")

	return []string{
		"https://google.com?q=" + formatGoogleQuery(args),
	}, nil
}

func formatGoogleQuery(args []string) string {

	var query string
	for i, argument := range args {

		if i == len(args)-1 {
			query += argument
		} else {
			query += argument + "+"
		}
	}
	return query
}
