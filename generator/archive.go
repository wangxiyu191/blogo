package generator

import (
	"blogo/util"
	"html/template"
	"os"
)

func generateArchive() {
	workDir, err := os.Getwd()
	util.CheckError(err)

	t, err := template.ParseFiles(workDir + "/templates/archive.html")
	util.CheckError(err)

	archiveFile, err := os.Create(workDir + "/build/archive.html")
	defer archiveFile.Close()
	util.CheckError(err)

	err = t.Execute(archiveFile, articles)
	util.CheckError(err)

}
