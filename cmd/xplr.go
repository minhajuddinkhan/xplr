package cmd

import (
	"github.com/fatih/color"
	"github.com/minhajuddinkhan/xplr/endpoints"
	"github.com/skratchdot/open-golang/open"
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
				c := color.New(color.FgRed).Add(color.Bold)
				c.Println("No Arguments?")
				return
			}
			c := make(chan []string)
			go endpoints.Fetch(args, c)
			articles := <-c
			for _, article := range articles {

				c := color.New(color.FgHiYellow).Add(color.Bold)
				c.Println("endpoints: ", article)
				open.Run(article)
			}

		},
	}

}
