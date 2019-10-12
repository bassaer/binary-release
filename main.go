package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extract(src string) string {
	items := strings.Split(src, " ")
	if len(items) != 3 {
		return ""
	}
	return items[2]
}

func version(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "## ver") {
			return extract(line)
		}
	}
	return ""
}

func main() {
	fmt.Printf("binay-release! (ver %s)\n", version("CHANGELOG.md"))
}
