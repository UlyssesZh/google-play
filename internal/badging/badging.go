package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var f struct {
		name    string
		verbose bool
	}
	flag.StringVar(&f.name, "i", "", "input file")
	flag.BoolVar(&f.verbose, "v", false, "verbose")
	flag.Parse()
	if f.name != "" {
		cmd := exec.Command("aapt", "dump", "badging", f.name)
		cmd.Stderr = os.Stderr
		fmt.Println(cmd.Args)
		text, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		lines := strings.FieldsFunc(string(text), func(r rune) bool {
			return r == '\n'
		})
		for _, line := range lines {
			if f.verbose ||
				strings.HasPrefix(line, "  uses-feature:") ||
				strings.HasPrefix(line, "  uses-gl-es:") ||
				strings.HasPrefix(line, "native-code:") ||
				strings.HasPrefix(line, "package:") ||
				strings.HasPrefix(line, "sdkVersion:") ||
				strings.HasPrefix(line, "supports-gl-texture:") ||
				strings.HasPrefix(line, "uses-library:") {
				fmt.Println(line)
			}
		}
	} else {
		flag.Usage()
	}
}
