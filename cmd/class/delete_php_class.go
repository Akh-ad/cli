package class

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var filePathToDelete string

var (
	deleteClassCmd = &cobra.Command{
		Use:   "delete_class",
		Short: "Delete a selected php class/File",
		Run: func(cmd *cobra.Command, args []string) {
			if filePathToDelete == "" {
				fmt.Println("Please specify the path of the file what you want to delete")
				return
			}

			err := os.Remove(filePathToDelete)
			if err != nil {
				fmt.Printf("Error deleting file : %v\n", err)
				return
			}

			fmt.Printf("File deleted with success : %s\n", filePathToDelete)
		},
	}
)

func AddDeleteClassCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(deleteClassCmd)

	deleteClassCmd.Flags().StringVarP(&filePathToDelete, "file", "f", "", "Path of file to delete (required)")
	deleteClassCmd.MarkFlagRequired("file")
}
