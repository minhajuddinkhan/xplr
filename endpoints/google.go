package endpoints

//Google structure
type Google struct {
}

var (
	str string
)

//Fetch fetches articles from hackernoon
func (h *Google) Fetch(queries []string, result *[]string) {

	*result = append(*result, "https://google.com/search?q="+h.formatQuery(queries))
}

func (h *Google) formatQuery(queries []string) string {

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
