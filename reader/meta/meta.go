package meta

import (
	"fmt"
	"strings"
)

func Split(input string) (meta string, content string) {
	index := strings.Index(input, "---")
	meta = input[:index-1]
	//fmt.Print(string(meta))

	offest := strings.Index(input[index:], "\n")
	content = input[index+offest+1:]
	//fmt.Print(string(md))

	return meta, content

}

func Prase(input string, meta map[string]string) (err error) {
	lines := strings.SplitAfter(input, "\n")
	if lines[0][0:2] == "# " {
		meta["title"] = lines[0][2:]
	}
	for i := 2; i < len(lines)-1; i++ {
		if !strings.Contains(lines[i], ":") {
			err = fmt.Errorf("Line %d has mistake(s).", i+1)
			return err
		}
		result := strings.SplitN(lines[i][2:], ":", 2)
		meta[result[0]] = strings.Trim(result[1], " \n")
	}
	return nil
}
