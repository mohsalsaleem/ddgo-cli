package utils

import "os"

// WriteStringToFile - Write string to a file
func WriteStringToFile(stringToWrite string, filePath string) {
	f, _ := os.Create(filePath)
	defer f.Close()
	f.WriteString(stringToWrite)
	f.Sync()
}
