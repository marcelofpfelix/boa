package core

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func ConfigCli(config Config) {
	// Create the root command
	rootCmd := createCommand(config.command)
	// Define flags for the root command
	defineFlags(rootCmd, config.Flags)
	// Recursively define subcommands
	defineCommands(rootCmd, config.Commands)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Function to create a command from the provided configuration
func createCommand(cmdConfig command) *cobra.Command {
	cmd := &cobra.Command{
		Use:     cmdConfig.Use,
		Short:   cmdConfig.Short,
		Long:    cmdConfig.Long,
		Aliases: cmdConfig.Aliases,
		Run: func(cmd *cobra.Command, args []string) {

			// print command path
			fmt.Println(cmd.CommandPath())
			// print all flags
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				fmt.Printf("%v\n", f.Value)
			})
		},
	}
	return cmd
}

// Function to recursively define subcommands
func defineCommands(parentCmd *cobra.Command, cmds []command) {
	for _, cmdConfig := range cmds {
		// Create the subcommand
		cmd := createCommand(cmdConfig)

		// Define flags for the subcommand
		defineFlags(cmd, cmdConfig.Flags)

		// Recursively define nested subcommands
		if len(cmdConfig.Commands) > 0 {
			defineCommands(cmd, cmdConfig.Commands)
		}
		// Add the subcommand to the parent command
		parentCmd.AddCommand(cmd)
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
