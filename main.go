package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

const (
	sampleFilePath = "./Sample.md"
)

func main() {
	dirName, fileName := inputDate()
	fmt.Println("Directory name:", dirName)

	outputPath := strings.Join([]string{"./", dirName}, "")
	fmt.Println("Output path:", outputPath)

	createDirectory(outputPath)
	
	createFile(dirName, fileName)
}

func inputDate() (string, string) {
	var year int
	var month string
	
	fmt.Print("Enter year and month to create directory & file (ex. 2025 JAN, 2025 FEB): ")
	fmt.Scanln(&year, &month)

	if year < 0 {
		fmt.Println("Invalid input(negative value). Please enter a valid year.")
		return inputDate() 
	}

	if month == "" {
		fmt.Println("Invalid input(empty value). Please enter a valid month.")
		return inputDate() 
	}

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter year.month to create path (ex. 2025.01, 2025.12, ...): ")
		
	// year, _ := reader.ReadString('\n')
	// fmt.Print("What month is this? (3-digit month (ex. JAN): ")
	// month, _ := reader.ReadString('\n')

	return strconv.Itoa(year), month
}

func createDirectory(path string) error {
	if !existCheck(path) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("Error creating directory: %w", err)
		}
	}
	return nil
}

func createFile(dirName, fileName string) {

	filePath := strings.Join([]string{"./", dirName, "/", fileName, ".md"}, "")
	fmt.Println("File path:", filePath)

	if !existCheck(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()
		

		writer := bufio.NewWriter(file)
		sampleContent, _ := readSampleFile(sampleFilePath)
		// // 여기서 다시 시작 : for 반복문 써서 fileName에 해당하는 월의 최대일 수 만큼 반복으로 sampleCOntent 내보내기
		// writer.WriteString(sampleContent)
		// writer.Flush()
		// fmt.Printf("File created successfully at %s\n", filePath)
	} else {
		fmt.Printf("File already exists at %s\n", filePath)
	}
}

// Check if the directory or file exists (exist: true, not exist: false)
func existCheck(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Read Sampple file from the given filepath
func readSampleFile(filepath string) (string, error) {

	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("Error reading sample file: %w", err)
	}

	return string(data), nil
}


