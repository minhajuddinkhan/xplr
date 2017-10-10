package endpoints

//Google structure
type Google struct {
}

//Fetch fetches articles from hackernoon
func (h *Google) Fetch(args []string) ([]string, error) {

	return []string{
		"https://google.com/search?q=" + h.formatQuery(args),
	}, nil
}

func (h *Google) formatQuery(args []string) string {

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
