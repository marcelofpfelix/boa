package core

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func ConfigCli(config Config) {
	var configRoot = &cobra.Command{
		Use:   config.Use,
		Short: config.Short,
		Long:  config.Long,
		//Args:  cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				fmt.Printf("%v\n", f.Value)
			})
		},
	}
	defineFlags(configRoot, config.Flags)

	if err := configRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func defineFlags(configRoot *cobra.Command, flags []flag) {
	// Get flags dynamically from config
	for _, flag := range flags {
		switch flag.Type {
		case "string":
			var flagValue string
			configRoot.Flags().StringVarP(&flagValue, flag.Use, flag.Alias, "", flag.Short)
		case "int":
			var flagValue int
			configRoot.Flags().IntVarP(&flagValue, flag.Use, flag.Alias, 0, flag.Short)
		default: // bool
			var flagValue bool
			configRoot.Flags().BoolVarP(&flagValue, flag.Use, flag.Alias, false, flag.Short)
		}
		configRoot.Flags().SortFlags = false
	}
}
