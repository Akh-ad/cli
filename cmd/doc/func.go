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

	description := doc.Find(".refpurpose").First().Text()

	examples := make([]string, 0)
	doc.Find(".example").Each(func(i int, s *goquery.Selection) {
		examples = append(examples, s.Text())
	})

	fmt.Println("Description of the function", funcName, ":", description)
	fmt.Println("Example", funcName, ":", strings.Join(examples, "\n\n"))
}
