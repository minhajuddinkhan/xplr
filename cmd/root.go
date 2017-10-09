package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//Root baseline commander
func Root() *cobra.Command {

	return &cobra.Command{
		Use:   "explorer",
		Short: "xplr is a cli that helps you get cool tech articles",
		Long: `xplr is a cli that helps you get cool tech articles
		* hackernoon
		* medium
		etc
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here

			if len(args) == 0 {
				fmt.Println("No Arguments?")
				return
			}
			articles := Fetch()
			for _, article := range articles {
				fmt.Println(article)
			}

		},
	}

}
