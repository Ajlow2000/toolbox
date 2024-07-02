package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fetch tool name from the specified path
// 
// Expects a README.md with the name of the tool 
// as a markdown h1 for the first line
func GetToolName(path string) string {
    readFile, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }
    defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line := fileScanner.Text()
        if strings.HasPrefix(line, "#") {
            return strings.Split(line, " ")[1]
        }
    }
    return "NAME_NOT_FOUND"
}

// Fetch short description from the specified path
// 
// Expects a README.md with the same formatting as LongDesc
// but parses only the first non header line as the ShortDesc
func GetShortDesc(path string) string {
    readFile, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }
    defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    desc := ""

    for fileScanner.Scan() {
        line := fileScanner.Text()
        if desc != "" {
            break
        }
        if strings.HasPrefix(line, "#") {
            continue
        }
        desc = line
        
    }

    return desc 
}

// Fetch long description from the specified path.
//
// Expects a README.md with the simple formatting of
// a h1 and then plaintext for the long description.
func GetLongDesc(path string) string {
    readFile, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }
    defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    desc := ""

    for fileScanner.Scan() {
        line := fileScanner.Text()
        if strings.HasPrefix(line, "#") {
            continue
        }
        if line == "" {
            line = "\n\n"
        }

        if strings.HasSuffix(line, ".") {
            line = line + " "
        }

        desc = desc + line
    }


    return desc 
}

// Parse version number from the version file
func GetVersion() string {
    readFile, err := os.Open("./VERSION")
    if err != nil {
        fmt.Println(err)
    }
    defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    version := ""

    for fileScanner.Scan() {
        version = fileScanner.Text()
    }

    return version
}

