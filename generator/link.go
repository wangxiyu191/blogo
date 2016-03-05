package generator

import (
	"blogo/util"
	"html/template"
	"os"
)

func generateLink() {
	workDir, err := os.Getwd()
	util.CheckError(err)

	linkFile, err := os.Create(workDir + "/build/links.html")
	defer linkFile.Close()
	util.CheckError(err)

	t, err := template.ParseFiles(workDir + "/templates/links.html")
	util.CheckError(err)
	links := util.Config.Links
	err = t.Execute(linkFile, links)
	util.CheckError(err)
}
