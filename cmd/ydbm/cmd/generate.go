package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/alexeyco/ydbm/internal/generator"
	"github.com/alexeyco/ydbm/internal/logx"
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("output", "o", "internal/ydbm", "Output directory")
}

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate [info]",
	Short: "Generate migration",
	Example: `ydbm generate add some table
ydbm generate add some table --output=path/to/directory`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := logx.New()

		info := strings.Join(args, " ")
		directory := cmd.Flags().Lookup("output").Value.String()

		if err := generator.Generate(info).To(directory); err != nil {
			log.Fatal(err)
		}
	},
}
