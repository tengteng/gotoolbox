package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Checkout link: https://fightingforalostcause.net/content/misc/2006/compare-email-regex.php
func main() {
	// EXP_EMAIL := regexp.MustCompile("<?\\S+@\\S+?>?")

	EXP_EMAIL := regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&\\'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")

	file, err := os.Open("./invalid.lst")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		if EXP_EMAIL.MatchString(line) {
			fmt.Println("Should be invalid but matched:\n", line)
		}
	}
	file.Close()

	file, err = os.Open("./valid.lst")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		if !EXP_EMAIL.MatchString(line) {
			fmt.Println("Should be valid but failed to match:\n", line)
		}
	}
	file.Close()
}
