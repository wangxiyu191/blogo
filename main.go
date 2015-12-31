package main

import (
	"blogo/generator"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	//new      = kingpin.Command("new", "Make a blogo folder")
	generate = kingpin.Command("generate", "Generate the blog")
	//publish  = kingpin.Command("publish", "publish the blog to Github")
)

func main() {
	switch kingpin.Parse() {
	case generate.FullCommand():
		generator.Generate()
	}
}
