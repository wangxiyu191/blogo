package generator

import (
	"blogo/util"
	//"github.com/termie/go-shutil"
	"os"
)

func generateStatic() {
	workDir, err := os.Getwd()
	util.CheckError(err)
	util.CopyDir(workDir+"/templates/static", workDir+"/build/static")
	util.CopyDir(workDir+"/source/static", workDir+"/build/static")
}
