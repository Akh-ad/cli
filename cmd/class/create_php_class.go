package class

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var className string

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

			// Formate le nom de la class avec la première lettre en maj
			className := strings.Title(className)

			// php class content
			phpContent := fmt.Sprintf(`<?php

	namespace App;

	class %s {
	
	}

	`, className)
			// File name
			fileName := fmt.Sprintf("%s.php", strings.ToLower(className))

			// Write the content in the file
			err := writeToFile(fileName, phpContent)
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
	createClassCmd.Flags().StringVarP(&className, "name", "n", "", "Classe name (required)")
	createClassCmd.MarkFlagRequired("name")
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
