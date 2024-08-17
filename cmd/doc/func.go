package doc

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

// Global variables to store the options/flags for the command
var (
	funcName string // The name of the PHP function to retrieve information about
	moreInfo bool   // Flag to indicate if more detailed information should be displayed
	writeToFile bool // Flag to indicate if a temporary file will be create to add the doc into
)

// Definition of the `get-info` command
var (
	getInfoCmd = &cobra.Command{
		Use:   "get-info",                           // Command usage
		Short: "Get information about a PHP function", // Short description of the command
		Run: func(cmd *cobra.Command, args []string) {
			// Check if the function name is provided, if not, display an error and exit
			if funcName == "" {
				fmt.Println("Please specify the function name with the flag --function")
				os.Exit(1)
			}
			// Retrieve the function information and print it
			functionInfo := getFunctionInfo(funcName)
			fmt.Println(functionInfo)

			// Conditionally write the information to a temporary file
			if writeToFile {
				WriteFunctionInfoToFile(functionInfo)
			} 
		},
	}
)

// Function to retrieve information about a PHP function
func getFunctionInfo(funcName string) string {

	// Format the URL of the PHP documentation page for the given function name
	docURL := fmt.Sprintf("https://www.php.net/manual/en/function.%s.php", strings.ToLower(funcName))

	// Load the HTML document from the PHP.net website
	doc, err := goquery.NewDocument(docURL)
	if err != nil {
		log.Fatal(err) // Handle the error by stopping the program if the document cannot be loaded
	}

	// Variable to store the extracted information
	var result string

	// Extract the content of the div with the class "refnamediv" which contains the function's description
	doc.Find("div.refnamediv").Each(func(i int, s *goquery.Selection) {
		content := s.Text()      // Get the text inside the div
		result += content + "\n" // Add the text to the final result
	})

	// If the user has requested more information, add additional sections
	if !moreInfo {
		result += getMoreFunctionInfo(doc)
	}

	return result // Return the retrieved information
}

// Function to retrieve additional information about the function (description, parameters, examples)
func getMoreFunctionInfo(doc *goquery.Document) string {

	var result string
	// Retrieve the function's description
	doc.Find("div.refsect1.description").Each(func(i int, s *goquery.Selection) {
		descriptionContent := s.Text()
		result += descriptionContent + "\n"
	})
	// Retrieve information about the function's parameters
	doc.Find("div.refsect1.parameters").Each(func(i int, s *goquery.Selection) {
		paramContent := s.Text()
		result += paramContent + "\n"
	})
	// Retrieve examples of how the function is used
	doc.Find("div.refsect1.examples").Each(func(i int, s *goquery.Selection) {
		paramContent := s.Text()
		result += paramContent + "\n"
	})

	return result // Return the additional information
}

// Function to add the `get-info` command to the root command of the CLI
func FuncInfoCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getInfoCmd) // Add the `get-info` command to the root command

	// Define the flags for the `get-info` command
	getInfoCmd.Flags().StringVarP(&funcName, "name", "n", "", "function name (required)")
	getInfoCmd.MarkFlagRequired("name") // Make the "name" flag required
	getInfoCmd.Flags().BoolVarP(&moreInfo, "moreInfos", "m", false, "more information about the function")
	getInfoCmd.Flags().BoolVarP(&writeToFile, "write-to-file", "w", false, "write information to a temporary file")
}
