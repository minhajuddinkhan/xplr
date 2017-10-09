package endpoints

var (
	hackernoon = Hackernoon{}
)

//Fetch fetches data from respective endpoints
func Fetch() ([]string, error) {

	articles, _ := hackernoon.Fetch()
	return articles, nil
}
