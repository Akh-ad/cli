package main

import (
	"cli/cmd/class" // Importing the `class` package containing commands related to PHP classes
	"cli/cmd/doc"   // Importing the `doc` package for commands related to PHP documentation
	"fmt"
	"github.com/spf13/cobra" // Importing the Cobra library for creating CLI applications
)

// Global variables for the CLI commands
var (
	rootCmd = &cobra.Command{Use: "app"} // The root command of the CLI
	name    string                        // Variable to store the name used in the `personalized` command

	// Definition of the `hello` command
	helloCmd = &cobra.Command{
		Use:   "hello",                   // Command usage
		Short: "Prints 'Hello, World!'",  // Short description of the command
		Run: func(cmd *cobra.Command, args []string) {
			// Action to execute when the `hello` command is called
			fmt.Println("Hello, World!")
		},
	}

	// Definition of the `personalized` command
	personalizedCmd = &cobra.Command{
		Use:   "personalized",                  // Command usage
		Short: "Prints a personalized greeting", // Short description of the command
		Run: func(cmd *cobra.Command, args []string) {
			// Action to execute when the `personalized` command is called
			// Prints a personalized message using the provided name argument
			fmt.Printf("Hello, %s!\n", name)
		},
	}
)

// Initialization function to add commands to the root command
func init() {
	// Adding the `hello` and `personalized` commands to the root command
	rootCmd.AddCommand(helloCmd, personalizedCmd)
	
	// Adding commands related to classes and documentation
	class.AddCreateClassCommand(rootCmd)
	class.AddDeleteClassCommand(rootCmd)
	doc.FuncInfoCommand(rootCmd)
	
	// Defining the `--name` flag for the `personalized` command
	personalizedCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
}

// Main function to execute the CLI
func main() {
	// Execute the root command, triggering the CLI
	rootCmd.Execute()
}
