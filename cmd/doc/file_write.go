// fil_writer.go
package doc 

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	// Define a fixed file name for the temporary file
	tempFileName = "function-info.txt"
)

// WriteFunctionInfoToFile writes the function information to a temporary file and prints the file path.
func WriteFunctionInfoToFile(functionInfo string){
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current working directory:", err)
	}

	// Define the path for the temporary file in the current working directory
	tempFilePath := filepath.Join(dir, tempFileName)

	// Create or open the file (this will overwrite the file if it already exists)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		log.Fatal("Unable to create or open a temporary file:", err)
	}
	defer tempFile.Close() // Ensure the file is closed after writing

	// Write  the function information to the temporary file 
	if _, err := tempFile.WriteString(functionInfo); err != nil {
		log.Fatal("Unable to write to a temporary file:", err)
	} 

	// Print the path to the temporary file
	fmt.Printf("The information has also been saved to: %s\n", tempFilePath)
}