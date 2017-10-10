package endpoints

//Hackernoon structure
type Hackernoon struct {
}

//Fetch fetches articles from hackernoon
func (h *Hackernoon) Fetch(args []string) ([]string, error) {
	return []string{
		"https://hackernoon.com/search?q=" + h.formatQuery(args),
	}, nil
}

func (h *Hackernoon) formatQuery(args []string) string {
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
