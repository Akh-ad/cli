package doc

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	funcName := "strlen"

	docURL := fmt.Sprintf("https://www.php.net/manual/en/function.%s.php", strings.ToLower(funcName))

	doc, err := goquery.NewDocument(docURL)
	if err != nil {
		log.Fatal(err)
	}

	description := doc.Find(".refpurpose").First().Text()

	examples := make([]string, 0)
	doc.Find(".example").Each(func(i int, s *goquery.Selection) {
		examples = append(examples, s.Text())
	})

	fmt.Println("Description of the function", funcName, ":", description)
	fmt.Println("Example", funcName, ":", strings.Join(examples, "\n\n"))
}
