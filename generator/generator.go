package generator

import (
	"blogo/util"
	"os"
)

func Generate() {
	util.Config.Load()
	loadArticles()
	resetBuild()

	generateStatic()
	generateArticles()
	generateIndex()
}

func resetBuild() {
	workDir, err := os.Getwd()
	util.CheckError(err)

	util.ClearDir(workDir + "/build/")
	util.CheckError(err)

}
