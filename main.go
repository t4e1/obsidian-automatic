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
	dirName, fileName, order := inputDate()
	// fmt.Println("Directory name:", dirName)

	outputPath := strings.Join([]string{"./", dirName}, "")
	// fmt.Println("Output path:", outputPath)

	createDirectory(outputPath)
	
	createFile(dirName, fileName, order)
}

func inputDate() (string, string, string) {
	var year int
	var month string

	validMonths := map[string]bool{
    "JAN": true, "FEB": true, "MAR": true, "APR": true,
    "MAY": true, "JUN": true, "JUL": true, "AUG": true,
    "SEP": true, "OCT": true, "NOV": true, "DEC": true,
	}

	orders := map[string]string{
    "JAN": "01", "FEB": "02", "MAR": "03", "APR": "04",
    "MAY": "05", "JUN": "06", "JUL": "07", "AUG": "08",
    "SEP": "09", "OCT": "10", "NOV": "11", "DEC": "12",
	}

	// Code refactored to use a loop for input validation(for preventing Stack Overflow) 
	for {
		fmt.Print("Enter year and month to create directory & file (ex. 2025 JAN, 2025 FEB): ")
		fmt.Scanln(&year, &month)

		if year < 0 {
			fmt.Println("Invalid input(negative value). Please enter a valid year.")
			continue
		}

		if month == "" {
			fmt.Println("Invalid input(empty value). Please enter a valid month.")
			continue
		}

		if !validMonths[strings.ToUpper(month)] {
			fmt.Println("Invalid input(month). Please enter a valid month (3-digit month (ex. JAN, FEB, ...)).")
			continue
		}

		return strconv.Itoa(year), strings.ToUpper(month), orders[strings.ToUpper(month)]
	}
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

func createFile(dirName, fileName, order string) {

	filePath := strings.Join([]string{"./", dirName, "/", order, ".", fileName, ".md"}, "")
	fmt.Println("File path:", filePath)

	if !existCheck(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()

		days := daysIn(dirName, fileName)
		writer := bufio.NewWriter(file)
		sampleContent, _ := readSampleFile(sampleFilePath)
		
		for i := 1; i <= days; i++ {
			line := fmt.Sprintf("# %s-%s-%d\n\n%s", dirName, fileName, i, sampleContent)
			writer.WriteString(line)
		}
		writer.Flush()
		fmt.Printf("File created successfully at %s\n", filePath)
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

func daysIn(year string, month string) int {
	
	y, _ := strconv.Atoi(year)
	months := map[string]time.Month{
		"JAN": time.January, "FEB": time.February, "MAR": time.March, "APR": time.April,
		"MAY": time.May, "JUN": time.June, "JUL": time.July, "AUG": time.August,
		"SEP": time.September, "OCT": time.October, "NOV": time.November, "DEC": time.December,
	}

	return time.Date(y, months[month]+1, 0, 0, 0, 0, 0, time.UTC).Day()
}


