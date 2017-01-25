package generator

import (
	"blogo/util"
	//"fmt"
	"github.com/gorilla/feeds"
	"time"
	"io/ioutil"
	"os"
)

func generateFeed() {
	feed := &feeds.Feed{
		Title:       "Wxy's Blog",
		Link:        &feeds.Link{Href: util.Config.Domain},
		Description: "Reversing,Web Backend",
		Author:      &feeds.Author{Name: "Wxy", Email: "wangxiyu191@gmail.com"},
		Created:     time.Now(),
	}
	for _, article := range articles {
		item := &feeds.Item{
			Id:      article.ID,
			Title:   article.Title,
			Link:    &feeds.Link{Href: article.ArticleUrl},
			Author:  &feeds.Author{Name: "Wxy", Email: "wangxiyu191@gmail.com"},
			Created: article.Time,
		}
		feed.Add(item)
	}

	rssOut, err := feed.ToRss()
	util.CheckError(err)

	atomOut,err := feed.ToAtom()
	util.CheckError(err)

	workDir, err := os.Getwd()
	ioutil.WriteFile(workDir + "/build/rss.xml",[]byte(rssOut),os.ModePerm)
	ioutil.WriteFile(workDir + "/build/atom.xml",[]byte(atomOut),os.ModePerm)
}
