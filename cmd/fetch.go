package cmd

import (
	"github.com/minhajuddinkhan/xplr/endpoints"
)

func Fetch(args []string) []string {

	articles, _ := endpoints.Fetch(args)
	return articles
}
