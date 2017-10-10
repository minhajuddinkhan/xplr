package endpoints

var (
	hackernoon    = Hackernoon{}
	google        = Google{}
	stackOverflow = StackOverflow{}
)

type Endpoint interface {
	Fetch([]string, *string) ([]string, error)
}

//Fetch fetches data from respective endpoints
func Fetch(args []string) ([]string, error) {

	var result []string
	hackernoon.Fetch(args, &result)
	google.Fetch(args, &result)
	stackOverflow.Fetch(args, &result)

	return result, nil
}
