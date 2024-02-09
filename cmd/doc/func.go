package doc

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var (
	funcName string
	moreInfo bool
)

var (
	getInfoCmd = &cobra.Command{
		Use:   "get-info",
		Short: "Get information about a PHP function",
		Run: func(cmd *cobra.Command, args []string) {
			if funcName == "" {
				fmt.Println("Please specify the function name with the flag --function")
				os.Exit(1)
			}
			functionInfo := getFunctionInfo(funcName)
			fmt.Println(functionInfo)
		},
	}
)

func getFunctionInfo(funcName string) string {

	docURL := fmt.Sprintf("https://www.php.net/manual/en/function.%s.php", strings.ToLower(funcName))

	doc, err := goquery.NewDocument(docURL)
	if err != nil {
		log.Fatal(err)
	}

	// Use a specific class toextract a content of div
	var result string
	doc.Find("div.refnamediv").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		result += content + "\n"
	})

	if !moreInfo {
		result += getMoreFunctionInfo(doc)
	}

	return result
}

func getMoreFunctionInfo(doc *goquery.Document) string {

	var result string
	// Description
	doc.Find("div.refsect1.description").Each(func(i int, s *goquery.Selection){
		descriptionContent := s.Text()
		result += descriptionContent + "\n"
	})
	// Parameters
	doc.Find("div.refsect1.parameters").Each(func(i int, s *goquery.Selection){
		paramContent := s.Text()
		result += paramContent + "\n"
	})
	// Examples
	doc.Find("div.refsect1.examples").Each(func(i int, s *goquery.Selection){
		paramContent := s.Text()
		result += paramContent + "\n"
	})

	return result
}

func FuncInfoCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getInfoCmd)


	getInfoCmd.Flags().StringVarP(&funcName, "name", "n", "", "function name (required)")
	getInfoCmd.MarkFlagRequired("name")
	getInfoCmd.Flags().BoolVarP(&moreInfo, "moreInfos", "m", false  , "more informations about the function")
}

