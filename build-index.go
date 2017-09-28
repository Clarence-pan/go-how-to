package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const placeholderBegin = "<!-- EXAMPLES BEGIN -->"
const placeholderEnd = "<!-- EXAMPLES END -->"
const readmePath = "README.md"

func main() {
	cwd, err := os.Getwd()
	dieIfError(err)
	log.Printf("Enter directory: %s", cwd)

	files, err := ioutil.ReadDir(".")
	dieIfError(err)

	examples := make([]string, len(files))[:0]
	for _, f := range files {
		if !f.IsDir() || isInBlackList(f.Name()) {
			continue
		}

		log.Printf("Processing %s", f.Name())
		examples = append(examples, fmt.Sprintf("- [%s](./%s)", humanize(f.Name()), f.Name()))
	}

	readmeContent, err := ioutil.ReadFile(readmePath)
	dieIfError(err)

	readmeContent = insertExamplesToReadmeContent(readmeContent, examples)

	err = ioutil.WriteFile(readmePath, readmeContent, 0644)
	dieIfError(err)

	log.Printf("All done.")
}

func dieIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func humanize(name string) string {
	s := strings.Replace(name, "-", " ", -1)
	s = strings.ToUpper(s[:1]) + s[1:]
	return s
}

func insertExamplesToReadmeContent(readmeContent []byte, examples []string) []byte {
	beginIndex := bytes.Index(readmeContent, []byte(placeholderBegin))
	endIndex := bytes.Index(readmeContent, []byte(placeholderEnd))

	if beginIndex < 0 {
		panic(errors.New("Cannot find placeholderBegin"))
	}

	if endIndex < 0 {
		panic(errors.New("Cannot find placeholderEnd"))
	}

	if endIndex <= beginIndex {
		panic(errors.New("Invalid placeholder"))
	}

	result := [][]byte{
		readmeContent[:beginIndex],
		[]byte(placeholderBegin), []byte("\n"),
		[]byte(strings.Join(examples, "\n")), []byte("\n"),
		readmeContent[endIndex:],
	}

	return bytes.Join(result, []byte{})
}

func isInBlackList(name string) bool {
	return name == ".git" || name == ".vscode"
}
