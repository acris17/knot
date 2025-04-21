package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) string {
	stdin := bufio.NewReader(os.Stdin)

	fmt.Printf("%v", prompt)
	line, _ := stdin.ReadString('\n')

	return strings.TrimSpace(line)
}
func Parse(input string) []string {
	return strings.Fields(input)
}
