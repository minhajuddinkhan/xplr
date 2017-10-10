package endpoints

//Hackernoon structure
type StackOverflow struct {
}

//Fetch fetches articles from hackernoon
func (h *StackOverflow) Fetch(queries []string, result *[]string) {

	*result = append(*result, "https://stackoverflow.com/search?q="+h.formatQuery(queries))

}

func (h *StackOverflow) formatQuery(queries []string) string {
	var query string
	for i, argument := range queries {

		if i == len(queries)-1 {
			query += argument
		} else {
			query += argument + "+"
		}
	}
	return query

}
