package utils

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/os/gproc"
)

var (
	// gofmtPath is the binary path of command `gofmt`.
	gofmtPath = gproc.SearchBinaryPath("gofmt")
)

// GoFmt formats the source file using command `gofmt -w -s PATH`.
func GoFmt(path string) {
	if gofmtPath != "" {
		gproc.ShellExec(fmt.Sprintf(`%s -w -s %s`, gofmtPath, path))
	}
}

var (
	replace = map[string]string{
		"iD":   "ID",
		"Id":   "ID",
		"Uid":  "UID",
		"Http": "HTTP",
		"Json": "JSON",
		"Url":  "URL",
		"Ip":   "IP",
		"Sql":  "SQL",
	}
)

// NameToLint 把字段名输出符合golint
func NameToLint(origin string) string {
	name := origin
	for from, to := range replace {
		name = strings.ReplaceAll(name, from, to)
	}
	return name
}
