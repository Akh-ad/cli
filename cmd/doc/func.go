package doc

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var funcName string

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
	return result
}

func FuncInfoCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(getInfoCmd)

	getInfoCmd.Flags().StringVarP(&funcName, "name", "n", "", "function name (required)")
	getInfoCmd.MarkFlagRequired("name")
}

