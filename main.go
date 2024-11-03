package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	c14n "github.com/ucarion/c14n"
)

func canonicalizeXML(file string) (string, error) {
	decoder := xml.NewDecoder(strings.NewReader(file))
	out, err := c14n.Canonicalize(decoder)
	if err != nil {
		return "", err
	}
	return string(out), err
}

func removeTextBetweenTags(input string) string {
	re := regexp.MustCompile(`>([^<]+)<`)
	result := re.ReplaceAllString(input, "><")
	return result
}

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", err
	}
	defer file.Close()

	readFileString := ""
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading file: ", err)
			return "", err
		}
		readFileString += string(buffer[:n])
	}

	return readFileString, nil
}

func main() {
	filePathA := flag.String("file_a", "", "filepath file A")
	filePathB := flag.String("file_b", "", "filepath file B")

	flag.Parse()

	filePathAStr := *filePathA
	filePathBStr := *filePathB

	fileAString, err := readFile(filePathAStr)
	if err != nil {
		panic(fmt.Sprint("Failed to read file A:", err))
	}

	fileBString, err := readFile(filePathBStr)
	if err != nil {
		panic(fmt.Sprint("Failed to read file B:", err))
	}

	fileAString = removeTextBetweenTags(fileAString)
	fileBString = removeTextBetweenTags(fileBString)

	fileAString, err = canonicalizeXML(fileAString)
	if err != nil {
		panic(fmt.Sprint("Failed to canonicalize XML for file A:", err))
	}

	fileBString, err = canonicalizeXML(fileBString)
	if err != nil {
		panic(fmt.Sprint("Failed to canonicalize XML for file B:", err))
	}

	fileABeautified := xmlfmt.FormatXML(fileAString, "\t", "  ")
	fileBBeautified := xmlfmt.FormatXML(fileBString, "\t", "  ")

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(fileABeautified, fileBBeautified, false)

	fmt.Println(dmp.DiffPrettyText(diffs))
}
