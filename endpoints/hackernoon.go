package endpoints

//Hackernoon structure
type Hackernoon struct {
}

//Fetch fetches articles from hackernoon
func (h *Hackernoon) Fetch(queries []string, result *[]string) {

	*result = append(*result, "https://hackernoon.com/search?q="+h.formatQuery(queries))

}

func (h *Hackernoon) formatQuery(queries []string) string {
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
