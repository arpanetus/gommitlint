package parser

import (
	"fmt"
	"os"
)

func Report(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args)
}

func ReportFatal(format string, args ...interface{}) {
	Report(format, args...)
	os.Exit(1)

}
