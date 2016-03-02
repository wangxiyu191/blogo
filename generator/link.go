package generator

import (
	"blogo/util"
	"os"
)

func generateLink() {
	workDir, err := os.Getwd()
	util.CheckError(err)

	util.CopyFile(workDir+"/templates/links.html", workDir+"/build/links.html")
}
