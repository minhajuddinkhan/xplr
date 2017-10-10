package endpoints

//Hackernoon structure
type StackOverflow struct {
}

//Fetch fetches articles from hackernoon
func (h *StackOverflow) Fetch(args []string, query *[]string) {

	*query = append(*query, "https://stackoverflow.com/search?q="+h.formatQuery(args))
}

func (h *StackOverflow) formatQuery(args []string) string {
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
