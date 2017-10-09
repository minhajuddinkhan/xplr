package endpoints

var (
	hackernoon = Hackernoon{}
)

//Fetch fetches data from respective endpoints
func Fetch(args []string) ([]string, error) {

	articles, _ := hackernoon.Fetch(args)
	return articles, nil
}
