package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo [command] (flags)",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
 love by spf13 and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hugo run...")
	},
}

func Execute() {
	path := rootCmd.CommandPath()

	fmt.Printf("command path: %s\n", path)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
