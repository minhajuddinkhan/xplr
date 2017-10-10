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

	c := make(chan string)
	go hackernoon.Fetch(args, c)
	go google.Fetch(args, c)
	go stackOverflow.Fetch(args, c)

	return []string{<-c, <-c, <-c}, nil
}
