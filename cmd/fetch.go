package cmd

import (
	"github.com/minhajuddinkhan/xplr/endpoints"
)

func Fetch() []string {

	articles, _ := endpoints.Fetch()
	return articles
}
