package generator

import (
	"blogo/util"
	"html/template"
	"os"
)

func generateIndex() {
	workDir, err := os.Getwd()
	util.CheckError(err)
	//util.CheckError(err)

	t, err := template.ParseFiles(workDir + "/templates/index.html")
	indexFile, err := os.Create(workDir + "/build/" + "index.html")
	util.CheckError(err)
	err = t.Execute(indexFile, articles)
	util.CheckError(err)

}
