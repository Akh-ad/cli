package class

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	className     string
	outputDirFlag string
)

// Create a command to create a class
var (
	createClassCmd = &cobra.Command{
		Use:   "create_class",
		Short: "Crée une class PHP basique",
		Run: func(cmd *cobra.Command, args []string) {
			if className == "" {
				fmt.Println("Veuillez spécifier un nom de class")
				return
			}

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

			// Formate le nom de la class avec la première lettre en maj
			className := strings.Title(className)

			namespace := strings.ReplaceAll(relPath, string(filepath.Separator), "\\")

			// php class content
			phpContent := fmt.Sprintf(`<?php

namespace %s;

class %s {
	
}

	`, namespace, className)

			if outputDirFlag == "" {
				outputDirFlag = "."
			}

			// File name
			fileName := filepath.Join(outputDirFlag, fmt.Sprintf("%s.php", strings.ToLower(className)))

			// Write the content in the file
			err = writeToFile(fileName, phpContent)
			if err != nil {
				fmt.Printf("Erreur lors de la création du fichier %v\n", err)
				return
			}

			fmt.Printf("Fichier PHP crée avec succès: %s\n", fileName)
		},
	}
)

func AddCreateClassCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(createClassCmd)

	// Add flags for the class creation command
	createClassCmd.Flags().StringVarP(&className, "name", "n", "", "Class name (required)")
	createClassCmd.MarkFlagRequired("name")

	createClassCmd.Flags().StringVarP(&outputDirFlag, "output", "o", "", "Output directory for the PHP file")
}

func writeToFile(fileName, content string) error {
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
