package endpoints

var (
	hackernoon    = Hackernoon{}
	google        = Google{}
	stackOverflow = StackOverflow{}
	result        []string
)

//Fetch fetches data from respective endpoints
func Fetch(queryArgs []string, c chan []string) {

	hackernoon.Fetch(queryArgs, &result)
	google.Fetch(queryArgs, &result)
	stackOverflow.Fetch(queryArgs, &result)
	c <- result
}
