package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 1. Create and Open a File
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Make sure to close the file at the end

	// 2. Write to a File
	_, err = file.WriteString("Hello, this is a file handling example.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file.")

	// 3. Open and Read the Entire File
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:\n", string(content))

	// 4. Append to a File
	file, err = os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Appending a new line.\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Data appended to file.")

	// 5. Read File Line by Line
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Reading file line by line:")
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		fmt.Println(reader.Text()) // Print each line
	}
	if err := reader.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// 6. Copy File Content to Another File
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create("example_copy.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	fmt.Println("File copied successfully.")

	// 7. Rename the File
	err = os.Rename("example_copy.txt", "renamed_example.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}
	fmt.Println("File renamed successfully.")

	// 8. Delete the File
	err = os.Remove("renamed_example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully.")
}
