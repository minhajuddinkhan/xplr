package endpoints

//Google structure
type Google struct {
}

var (
	str string
)

//Fetch fetches articles from hackernoon
func (h *Google) Fetch(args []string, query *[]string) {

	*query = append(*query, "https://google.com/search?q="+h.formatQuery(args))
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
