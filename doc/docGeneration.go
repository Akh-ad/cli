package doc

import (
	"fmt"
	"github.com/spf13/cobra"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

var (
	sourceFile string
)

var (
	generateDocCmd = &cobra.Command{
		Use: "generate-doc",
		Short: "Generate documentation for php function",
		Run: func(cmd *cobra.Command, args []string){
			if sourceFile == "" {
				fmt.Println("Please specify the source file with the flag --source")
				os.Exit(1)
			}

			functions := getPHPFunctions(sourceFile)
			generateDocumentation(functions)
		},
	}
)

func getPHPFunctions(sourceFile string) []string {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, sourceFile, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	
	functions := make([]string, 0)
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if fn.Doc != nil {
				comment := fn.Doc.Text()
				if strings.Contains(comment, "@phpdoc") {
					functionSig := extractFunctionSignature(fn)
					functions = append(functions, functionSig)
				}
			}
		}
	}
	return functions
}

func extractFunctionSignature(fn *ast.FuncDecl) string {
	// Extract functoin Name
	functionSig := fn.Name.Name + "("
	
	// Extract Parameters
	for i, param := range fn.Type.Params.List {
		for _, name := range param.Names {
			functionSig += name.Name + " "
		}
		functionSig += param.Type.(*ast.Ident).Name
		if i < len(fn.Type.Params.List)-1 {
			functionSig += ", "
		}
	}
	
	// Extract Return Type
	if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 {
		functionSig += ") " + fn.Type.Results.List[0].Type.(*ast.Ident).Name
	}else{
		functionSig += ")"
	}
	return functionSig
}

func generateDocumentation(functions []string) {
	fmt.Println("Doc for php functions:")
	for _, fn := range functions {
		fmt.Printf(" - %s\n", fn)
	}
}

func DocGenerationCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(generateDocCmd)
	
	generateDocCmd.Flags().StringVarP(&sourceFile, "source", "s", "", "source file (required")
	generateDocCmd.MarkFlagRequired("source")
}