package class

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	className       string
	outputDirFlag   string
	noFunctionsFlag bool
)

// Create a command to create a class
var (
	createClassCmd = &cobra.Command{
		Use:   "create_class",
		Short: "Crée une class PHP basique",
		Run: func(cmd *cobra.Command, args []string) {

			//Dir Path
			outputDir, err := filepath.Abs(outputDirFlag)
			if err != nil {
				fmt.Printf("Erreur lors de la récupération du chemin absolue")
				return
			}

			// current directory
			currentDir, err := os.Getwd()
			if err != nil {
				fmt.Printf("Erreur lors de la récupération du répertoire de travail actuel: %v\n", err)
				return
			}

			// Relatif path
			relPath, err := filepath.Rel(currentDir, outputDir)
			if err != nil {
				fmt.Printf("Erreur lors de la récupération du chemin relatif: %v\n", err)
				return
			}

			// Ask  user if he needs to add default functions in his class (e.g. __construct)
			if !noFunctionsFlag {
				fmt.Print("Do you want add functions | __construct | __toString ? (yes/no)")
				//Retrieve the user answer
				var userResponse string
				fmt.Scanln(&userResponse)

				if userResponse == "yes" {
					options := getFunctionOption()
					namespace := strings.ReplaceAll(relPath, string(filepath.Separator), "\\")
					phpContent := generateClassCodeWithFunctions(options, namespace)
					writeToFile(outputDirFlag, className, phpContent)
				} else {
					// User choose not to add any functions
					fmt.Printf("Creating a class without functions: %s\n", className)
					className := strings.Title(className)
					namespace := strings.ReplaceAll(relPath, string(filepath.Separator), "\\")
					phpContent := generateClassCode(namespace)
					writeToFile(outputDirFlag, className, phpContent)
				}

			} else {
				fmt.Println("The flag --no-functions is defined, the user do not want a default functions")
			}

			if className == "" {
				fmt.Println("Veuillez spécifier un nom de class")
				return
			}
		},
	}
)

func generateClassCode(namespace string) string {

	return fmt.Sprintf(`<?php

namespace %s;

class %s {

`, namespace, className)
}

// call the function who generate the code for the function chosen by the user
func generateClassCodeWithFunctions(optionsStr, namespace string) string {
	phpContent := generateClassCode(namespace)

	options := parseOptions(optionsStr)

	for _, option := range options {
		switch option {
		case "__construct":
			phpContent += generateConstructorCode()
		case "__toString":
			phpContent += generateToStringCode()
		}
	}

	phpContent += "}\n"
	return phpContent
}

// genrate the __construct function into the class
func generateConstructorCode() string {
	return "public function __construct()\n" +
		"	{\n" +
		"	}\n"

}

// generate __toString function into the class
func generateToStringCode() string {
	return "public function __toString()\n" +
		"	{\n" +
		"	}\n"

}

func AddCreateClassCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(createClassCmd)

	// Add flags for the class creation command
	createClassCmd.Flags().StringVarP(&className, "name", "n", "", "Class name (required)")
	createClassCmd.MarkFlagRequired("name")

	createClassCmd.Flags().StringVarP(&outputDirFlag, "output", "o", "", "Output directory for the PHP file")
	createClassCmd.Flags().BoolVarP(&noFunctionsFlag, "no-functions", "f", false, "functions creator for php class")
}

// write to the file that content the class
func writeToFile(outputDir, className, content string) error {
	fileName := filepath.Join(outputDir, fmt.Sprintf("%s.php", strings.ToLower(className)))

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func getFunctionOption() string {
	fmt.Println("Choose a function for the class:")
	fmt.Println("1. __construct")
	fmt.Println("2. __toString")
	
	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		return "__construct"
	case "2":
		return "__toString"
	default:
		fmt.Println("Invalid option. Defaulting to __construct.")
		return "__construct"
	}
}

func parseOptions(input string) []string {
	// Analys the options
	options := make([]string, 0)
	for _, opt := range splitOptions(input) {
		if opt != "" {
			options = append(options, opt)
		}
	}
	return options
}

func splitOptions(input string) []string {
	return strings.Split(input, ",")
}
