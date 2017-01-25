package generator

import "net/url"
import (
	"blogo/util"
	"strings"
	"io/ioutil"
	"os"
)

func generateCNAME()  {


	u , err := url.Parse(util.Config.Domain)
	util.CheckError(err)

	workDir, err := os.Getwd()
	util.CheckError(err)

	host := strings.Split(u.Host, ":")
	ioutil.WriteFile(workDir+"/build/CNAME",[]byte(host[0]),os.ModePerm)
}