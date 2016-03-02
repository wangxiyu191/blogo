package generator

import (
	"blogo/util"
	"os"
)

func generateAbout() {
	workDir, err := os.Getwd()
	util.CheckError(err)

	util.CopyFile(workDir+"/templates/about.html", workDir+"/build/about.html")
}
