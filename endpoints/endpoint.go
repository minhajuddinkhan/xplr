package endpoints

var (
	hackernoon = Hackernoon{}
	google     = Google{}
)

//Fetch fetches data from respective endpoints
func Fetch(args []string) ([]string, error) {

	h, _ := hackernoon.Fetch(args)
	g, _ := google.Fetch(args)
	var articles []string

	for _, val := range h {
		articles = append(articles, val)
	}

	for _, val := range g {
		articles = append(articles, val)
	}

	return articles, nil
}
