package generator

import (
	"blogo/util"
	"github.com/termie/go-shutil"
	"os"
)

func generateStatic() {
	workDir, err := os.Getwd()
	util.CheckError(err)
	shutil.CopyTree(workDir+"/templates/static", workDir+"/build/static", nil)
	util.CopyDir(workDir+"/source/static", workDir+"/build/static")
}
