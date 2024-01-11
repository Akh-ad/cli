package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints 'Hello, World!'",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	var name string
	var personalizedCmd = &cobra.Command{
		Use:   "personalized",
		Short: "Prints a personalized greeting",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
		},
	}
	personalizedCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")

	rootCmd.AddCommand(helloCmd, personalizedCmd)
	rootCmd.Execute()
}
