package publisher

import (
	"blogo/util"
)

func Publish() {
	util.Config.Load()
	viaGit()
}
