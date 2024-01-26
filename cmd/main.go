package main

import (
	"cli/cmd/class"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "app"}
	name    string

	helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints 'Hello, World!'",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	personalizedCmd = &cobra.Command{
		Use:   "personalized",
		Short: "Prints a personalized greeting",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
		},
	}
)

func init() {
	rootCmd.AddCommand(helloCmd, personalizedCmd)
	class.AddCreateClassCommand(rootCmd)
	class.AddDeleteClassCommand(rootCmd)
	personalizedCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
}

func main() {
	rootCmd.Execute()
}
