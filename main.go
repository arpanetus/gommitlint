package main

import (
	"github.com/arpanetus/gommitlint/parser"
	"os"
)

func main() {
	parser.Parse(os.Stdin)
}
